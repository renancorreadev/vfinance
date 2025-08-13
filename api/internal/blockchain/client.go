package blockchain

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"vfinance-api/internal/config"
	"vfinance-api/internal/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Constantes para o contrato
const (
	ContractRegisteredEventName = "ContractRegistered"
	VehicleTrackedEventName     = "VehicleTracked"
	StatusChangedEventName      = "StatusChanged"
	MetadataUpdatedEventName    = "MetadataUpdated"
)

// Erros customizados do contrato
var (
	ErrContractNotFound     = fmt.Errorf("ContractNotFound")
	ErrContractExists       = fmt.Errorf("ContractAlreadyExists")
	ErrDuplicateHash        = fmt.Errorf("DuplicateHash")
	ErrUnauthorizedAccess   = fmt.Errorf("UnauthorizedAccess")
	ErrInvalidInput         = fmt.Errorf("InvalidInput")
	ErrMaxSupplyExceeded    = fmt.Errorf("MaxSupplyExceeded")
	ErrTokenNotFound        = fmt.Errorf("TokenNotFound")
)

type Client struct {
	ethClient       *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	contractABI     abi.ABI
	chainID         *big.Int
}

// Estruturas que correspondem ao contrato
type ContractRecord struct {
	RegistryId     [32]byte `json:"registryId"`
	ContractNumber [32]byte `json:"contractNumber"`
	ContractDate   uint32   `json:"contractDate"`
	MetadataHash   [32]byte `json:"metadataHash"`
	Timestamp      uint32   `json:"timestamp"`
	RegisteredBy   common.Address `json:"registeredBy"`
	Active         bool     `json:"active"`
}

type VehicleCore struct {
	Chassis      [32]byte `json:"chassis"`
	LicensePlate [32]byte `json:"licensePlate"`
	TotalValue   *big.Int `json:"totalValue"` // uint128
	BrandId      uint64   `json:"brandId"`
	ModelId      uint64   `json:"modelId"`
}

type ContractInfo struct {
	Contract ContractRecord `json:"contract"`
	Vehicle  VehicleCore    `json:"vehicle"`
}

type RegistrationResult struct {
	TokenId      *big.Int `json:"tokenId"`
	MetadataHash [32]byte `json:"metadataHash"`
	TxHash       string   `json:"txHash"`
}

func NewClient(rpcURL, contractAddr, privateKeyHex string, chainID int64) (*Client, error) {
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial ethereum client: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(config.ContractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &Client{
		ethClient:       ethClient,
		contractAddress: common.HexToAddress(contractAddr),
		privateKey:      privateKey,
		contractABI:     contractAbi,
		chainID:         big.NewInt(chainID),
	}, nil
}

// Função auxiliar para converter string para bytes32
func stringToBytes32(s string) [32]byte {
	var result [32]byte
	copy(result[:], s)
	return result
}

// Função auxiliar para converter bytes32 para string
func bytes32ToString(b [32]byte) string {
	return strings.TrimRight(string(b[:]), "\x00")
}

// Função auxiliar para calcular hash keccak256 de uma string
func keccak256Hash(data string) [32]byte {
	hash := crypto.Keccak256Hash([]byte(data))
	var result [32]byte
	copy(result[:], hash[:])
	return result
}

// RegisterContract registra um novo contrato no blockchain
func (c *Client) RegisterContract(registryId, contractNumber, chassis, licensePlate, brandName, modelName string, contractDate uint32, totalValue *big.Int) (*RegistrationResult, error) {
	if registryId == "" || contractNumber == "" || chassis == "" || licensePlate == "" || brandName == "" || modelName == "" {
		return nil, ErrInvalidInput
	}

	// Verificar se o contrato já existe
	exists, err := c.DoesContractExist(registryId)
	if err != nil {
		return nil, fmt.Errorf("failed to check if contract exists: %w", err)
	}
	if exists {
		return nil, ErrContractExists
	}

	auth, err := c.createTransactor()
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Preparar dados para transação
	data, err := c.contractABI.Pack("registerContract",
		registryId,
		contractNumber,
		contractDate,
		chassis,
		licensePlate,
		totalValue,
		brandName,
		modelName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack transaction data: %w", err)
	}

	// Estimar gas
	gasLimit, err := c.ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From: auth.From,
		To:   &c.contractAddress,
		Data: data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	// Adicionar margem de segurança ao gas
	gasLimit = gasLimit * 120 / 100

	auth.GasLimit = gasLimit

	// Criar transação
	nonce, err := c.ethClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := c.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	tx := types.NewTransaction(nonce, c.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Assinar transação
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.chainID), c.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Enviar transação
	err = c.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	// Decodificar eventos para obter tokenId e metadataHash
	result := &RegistrationResult{
		TxHash: signedTx.Hash().Hex(),
	}

	for _, log := range receipt.Logs {
		if log.Address == c.contractAddress && len(log.Topics) > 0 {
			switch log.Topics[0] {
			case c.contractABI.Events[ContractRegisteredEventName].ID:
				event, err := c.parseContractRegisteredEvent(log)
				if err == nil {
					result.TokenId = event.TokenId
					result.MetadataHash = event.MetadataHash
				}
			}
		}
	}

	return result, nil
}

// GetContractByTokenId obtém os dados do contrato pelo token ID
func (c *Client) GetContractByTokenId(tokenId *big.Int) (*ContractInfo, error) {
	data, err := c.contractABI.Pack("getContract", tokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var unpacked []interface{}
	err = c.contractABI.UnpackIntoInterface(&unpacked, "getContract", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	if len(unpacked) != 2 {
		return nil, fmt.Errorf("unexpected result length: %d", len(unpacked))
	}

	// Parse ContractRecord
	contractData := unpacked[0].(struct {
		RegistryId     [32]byte       `json:"registryId"`
		ContractNumber [32]byte       `json:"contractNumber"`
		ContractDate   uint32         `json:"contractDate"`
		MetadataHash   [32]byte       `json:"metadataHash"`
		Timestamp      uint32         `json:"timestamp"`
		RegisteredBy   common.Address `json:"registeredBy"`
		Active         bool           `json:"active"`
	})

	// Parse VehicleCore
	vehicleData := unpacked[1].(struct {
		Chassis      [32]byte `json:"chassis"`
		LicensePlate [32]byte `json:"licensePlate"`
		TotalValue   *big.Int `json:"totalValue"`
		BrandId      uint64   `json:"brandId"`
		ModelId      uint64   `json:"modelId"`
	})

	return &ContractInfo{
		Contract: ContractRecord{
			RegistryId:     contractData.RegistryId,
			ContractNumber: contractData.ContractNumber,
			ContractDate:   contractData.ContractDate,
			MetadataHash:   contractData.MetadataHash,
			Timestamp:      contractData.Timestamp,
			RegisteredBy:   contractData.RegisteredBy,
			Active:         contractData.Active,
		},
		Vehicle: VehicleCore{
			Chassis:      vehicleData.Chassis,
			LicensePlate: vehicleData.LicensePlate,
			TotalValue:   vehicleData.TotalValue,
			BrandId:      vehicleData.BrandId,
			ModelId:      vehicleData.ModelId,
		},
	}, nil
}

// GetContractByRegistryId obtém os dados do contrato pelo registry ID
func (c *Client) GetContractByRegistryId(registryId string) (*ContractInfo, error) {
	data, err := c.contractABI.Pack("getContractByRegistryId", registryId)
	if err != nil {
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var unpacked []interface{}
	err = c.contractABI.UnpackIntoInterface(&unpacked, "getContractByRegistryId", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	if len(unpacked) != 2 {
		return nil, fmt.Errorf("unexpected result length: %d", len(unpacked))
	}

	// Parse similar to GetContractByTokenId
	contractData := unpacked[0].(struct {
		RegistryId     [32]byte       `json:"registryId"`
		ContractNumber [32]byte       `json:"contractNumber"`
		ContractDate   uint32         `json:"contractDate"`
		MetadataHash   [32]byte       `json:"metadataHash"`
		Timestamp      uint32         `json:"timestamp"`
		RegisteredBy   common.Address `json:"registeredBy"`
		Active         bool           `json:"active"`
	})

	vehicleData := unpacked[1].(struct {
		Chassis      [32]byte `json:"chassis"`
		LicensePlate [32]byte `json:"licensePlate"`
		TotalValue   *big.Int `json:"totalValue"`
		BrandId      uint64   `json:"brandId"`
		ModelId      uint64   `json:"modelId"`
	})

	return &ContractInfo{
		Contract: ContractRecord{
			RegistryId:     contractData.RegistryId,
			ContractNumber: contractData.ContractNumber,
			ContractDate:   contractData.ContractDate,
			MetadataHash:   contractData.MetadataHash,
			Timestamp:      contractData.Timestamp,
			RegisteredBy:   contractData.RegisteredBy,
			Active:         contractData.Active,
		},
		Vehicle: VehicleCore{
			Chassis:      vehicleData.Chassis,
			LicensePlate: vehicleData.LicensePlate,
			TotalValue:   vehicleData.TotalValue,
			BrandId:      vehicleData.BrandId,
			ModelId:      vehicleData.ModelId,
		},
	}, nil
}

// GetContractByChassis obtém os dados do contrato pelo chassi
func (c *Client) GetContractByChassis(chassis string) (*ContractInfo, error) {
	data, err := c.contractABI.Pack("getContractByChassis", chassis)
	if err != nil {
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var unpacked []interface{}
	err = c.contractABI.UnpackIntoInterface(&unpacked, "getContractByChassis", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	if len(unpacked) != 2 {
		return nil, fmt.Errorf("unexpected result length: %d", len(unpacked))
	}

	contractData := unpacked[0].(struct {
		RegistryId     [32]byte       `json:"registryId"`
		ContractNumber [32]byte       `json:"contractNumber"`
		ContractDate   uint32         `json:"contractDate"`
		MetadataHash   [32]byte       `json:"metadataHash"`
		Timestamp      uint32         `json:"timestamp"`
		RegisteredBy   common.Address `json:"registeredBy"`
		Active         bool           `json:"active"`
	})

	vehicleData := unpacked[1].(struct {
		Chassis      [32]byte `json:"chassis"`
		LicensePlate [32]byte `json:"licensePlate"`
		TotalValue   *big.Int `json:"totalValue"`
		BrandId      uint64   `json:"brandId"`
		ModelId      uint64   `json:"modelId"`
	})

	return &ContractInfo{
		Contract: ContractRecord{
			RegistryId:     contractData.RegistryId,
			ContractNumber: contractData.ContractNumber,
			ContractDate:   contractData.ContractDate,
			MetadataHash:   contractData.MetadataHash,
			Timestamp:      contractData.Timestamp,
			RegisteredBy:   contractData.RegisteredBy,
			Active:         contractData.Active,
		},
		Vehicle: VehicleCore{
			Chassis:      vehicleData.Chassis,
			LicensePlate: vehicleData.LicensePlate,
			TotalValue:   vehicleData.TotalValue,
			BrandId:      vehicleData.BrandId,
			ModelId:      vehicleData.ModelId,
		},
	}, nil
}

// GetActiveContracts obtém a lista de contratos ativos (tokenIds)
func (c *Client) GetActiveContracts(offset, limit uint64) ([]*big.Int, error) {
	data, err := c.contractABI.Pack("getActiveContracts", big.NewInt(int64(offset)), big.NewInt(int64(limit)))
	if err != nil {
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var tokenIds []*big.Int
	err = c.contractABI.UnpackIntoInterface(&tokenIds, "getActiveContracts", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return tokenIds, nil
}

// DoesContractExist verifica se um contrato existe pelo registry ID
func (c *Client) DoesContractExist(registryId string) (bool, error) {
	data, err := c.contractABI.Pack("doesContractExist", registryId)
	if err != nil {
		return false, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return false, fmt.Errorf("failed to call contract: %w", err)
	}

	var exists bool
	err = c.contractABI.UnpackIntoInterface(&exists, "doesContractExist", result)
	if err != nil {
		return false, fmt.Errorf("failed to unpack result: %w", err)
	}

	return exists, nil
}

// GetMetadataUrl obtém a URL dos metadados pelo hash
func (c *Client) GetMetadataUrl(metadataHash [32]byte) (string, error) {
	data, err := c.contractABI.Pack("getMetadataUrl", metadataHash)
	if err != nil {
		return "", fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var url string
	err = c.contractABI.UnpackIntoInterface(&url, "getMetadataUrl", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %w", err)
	}

	return url, nil
}

// GetMetadataUrlByRegistryId obtém a URL dos metadados pelo registry ID
func (c *Client) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	data, err := c.contractABI.Pack("getMetadataUrlByRegistryId", registryId)
	if err != nil {
		return "", fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var url string
	err = c.contractABI.UnpackIntoInterface(&url, "getMetadataUrlByRegistryId", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %w", err)
	}

	return url, nil
}

// GetBrandName obtém o nome da marca pelo ID
func (c *Client) GetBrandName(brandId uint64) (string, error) {
	data, err := c.contractABI.Pack("getBrandName", brandId)
	if err != nil {
		return "", fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var brandName string
	err = c.contractABI.UnpackIntoInterface(&brandName, "getBrandName", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %w", err)
	}

	return brandName, nil
}

// GetModelName obtém o nome do modelo pelo ID
func (c *Client) GetModelName(modelId uint64) (string, error) {
	data, err := c.contractABI.Pack("getModelName", modelId)
	if err != nil {
		return "", fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var modelName string
	err = c.contractABI.UnpackIntoInterface(&modelName, "getModelName", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %w", err)
	}

	return modelName, nil
}

// GetTotalSupply obtém o total de tokens emitidos
func (c *Client) GetTotalSupply() (*big.Int, error) {
	data, err := c.contractABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var total *big.Int
	err = c.contractABI.UnpackIntoInterface(&total, "totalSupply", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %w", err)
	}

	return total, nil
}

// GetContractVersion obtém a versão do contrato
func (c *Client) GetContractVersion() (string, error) {
	data, err := c.contractABI.Pack("getVersion")
	if err != nil {
		return "", fmt.Errorf("failed to pack call data: %w", err)
	}

	result, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &c.contractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var version string
	err = c.contractABI.UnpackIntoInterface(&version, "getVersion", result)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %w", err)
	}

	return version, nil
}

// Métodos auxiliares privados

func (c *Client) createTransactor() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	nonce, err := c.ethClient.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := c.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	return auth, nil
}

// Estrutura para o evento ContractRegistered
type ContractRegisteredEvent struct {
	TokenId         *big.Int
	RegistryIdHash  [32]byte
	ChassisHash     [32]byte
	MetadataHash    [32]byte
	Timestamp       uint32
}

func (c *Client) parseContractRegisteredEvent(log *types.Log) (*ContractRegisteredEvent, error) {
	event := &ContractRegisteredEvent{}

	// Parse indexed topics (tokenId, registryIdHash, chassisHash)
	if len(log.Topics) >= 4 {
		event.TokenId = log.Topics[1].Big()
		copy(event.RegistryIdHash[:], log.Topics[2][:])
		copy(event.ChassisHash[:], log.Topics[3][:])
	}

	// Parse non-indexed data (metadataHash, timestamp)
	// O evento tem 2 parâmetros não-indexados: metadataHash e timestamp
	var unpacked []interface{}
	err := c.contractABI.UnpackIntoInterface(&unpacked, ContractRegisteredEventName, log.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack event data: %w", err)
	}

	// Verificar se temos os 2 parâmetros esperados
	if len(unpacked) >= 2 {
		// metadataHash é o primeiro parâmetro não-indexado
		if metadataHash, ok := unpacked[0].([32]byte); ok {
			event.MetadataHash = metadataHash
		} else {
			return nil, fmt.Errorf("failed to parse metadataHash from event data")
		}

		// timestamp é o segundo parâmetro não-indexado
		if timestamp, ok := unpacked[1].(uint32); ok {
			event.Timestamp = timestamp
		} else {
			return nil, fmt.Errorf("failed to parse timestamp from event data")
		}
	} else {
		return nil, fmt.Errorf("expected 2 non-indexed parameters, got %d", len(unpacked))
	}

	return event, nil
}

// Converter ContractInfo para o modelo usado pela API
func (c *Client) ContractInfoToModel(info *ContractInfo) *models.ContractRecord {
	return &models.ContractRecord{
		RegConId:       bytes32ToString(info.Contract.RegistryId),
		NumeroContrato: bytes32ToString(info.Contract.ContractNumber),
		DataContrato:   strconv.FormatUint(uint64(info.Contract.ContractDate), 10),
		MetadataHash:   hex.EncodeToString(info.Contract.MetadataHash[:]),
		Timestamp:      uint64(info.Contract.Timestamp),
		RegisteredBy:   info.Contract.RegisteredBy.Hex(),
		Active:         info.Contract.Active,
	}
}

// Converter RegistrationResult para modelo da resposta
func (c *Client) RegistrationResultToResponse(result *RegistrationResult, regConId string) *models.ContractRegistrationResponse {
	return &models.ContractRegistrationResponse{
		Success:      true,
		Message:      "Contrato registrado com sucesso",
		RegConId:     regConId,
		MetadataHash: hex.EncodeToString(result.MetadataHash[:]),
		TxHash:       result.TxHash,
	}
}
