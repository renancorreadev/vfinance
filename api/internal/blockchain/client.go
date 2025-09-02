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

	"vfinance-api/internal/models"
	"vfinance-api/internal/blockchain/bindings"

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
	ErrContractNotFound   = fmt.Errorf("ContractNotFound")
	ErrContractExists     = fmt.Errorf("ContractAlreadyExists")
	ErrDuplicateHash      = fmt.Errorf("DuplicateHash")
	ErrUnauthorizedAccess = fmt.Errorf("UnauthorizedAccess")
	ErrInvalidInput       = fmt.Errorf("InvalidInput")
	ErrMaxSupplyExceeded  = fmt.Errorf("MaxSupplyExceeded")
	ErrTokenNotFound      = fmt.Errorf("TokenNotFound")
)

type Client struct {
	ethClient       *ethclient.Client
	contractAddress common.Address
	privateKey      *ecdsa.PrivateKey
	contract        *bindings.Bindings
	chainID         *big.Int
}

// Estruturas que correspondem ao contrato (usando as do binding)
type ContractRecord = bindings.VFinanceRegistryContractRecord
type VehicleCore = bindings.VFinanceRegistryVehicleCore

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

	contractAddress := common.HexToAddress(contractAddr)

	// Criar instância do binding do contrato
	contract, err := bindings.NewBindings(contractAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract binding: %w", err)
	}

	client := &Client{
		ethClient:       ethClient,
		contractAddress: contractAddress,
		privateKey:      privateKey,
		contract:        contract,
		chainID:         big.NewInt(chainID),
	}

	return client, nil
}

// Função auxiliar para converter bytes32 para string
func bytes32ToString(b [32]byte) string {
	// Tenta como string primeiro
	str := strings.TrimRight(string(b[:]), "\x00")

	// Verifica se é uma string ASCII válida
	isValidString := true
	for _, char := range str {
		if char < 32 || char > 126 {
			isValidString = false
			break
		}
	}

	if isValidString && len(str) > 0 {
		return str
	}

	// Se não é string válida, retorna como hex (sem zeros à direita)
	trimmed := strings.TrimRight(hex.EncodeToString(b[:]), "0")
	if len(trimmed) == 0 {
		return ""
	}
	return "0x" + trimmed
}

// Bytes32ToString é a versão pública da função auxiliar
func (c *Client) Bytes32ToString(b [32]byte) string {
	return bytes32ToString(b)
}

// HexStringToBytes32 converte string hexadecimal para bytes32
func (c *Client) HexStringToBytes32(hexStr string) ([32]byte, error) {
	// Remover prefixo 0x se presente
	if len(hexStr) > 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	// Garantir que tenha 64 caracteres (32 bytes)
	if len(hexStr) != 64 {
		return [32]byte{}, fmt.Errorf("hash deve ter exatamente 64 caracteres hexadecimais")
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return [32]byte{}, fmt.Errorf("hash hexadecimal inválido: %w", err)
	}

	var result [32]byte
	copy(result[:], bytes)
	return result, nil
}

// GetContractByHash obtém contrato por metadata hash
func (c *Client) GetContractByHash(metadataHash [32]byte) (*ContractInfo, error) {
	contractRecord, vehicleCore, err := c.contract.GetContractByHash(&bind.CallOpts{}, metadataHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract by hash: %w", err)
	}

	return &ContractInfo{
		Contract: contractRecord,
		Vehicle:  vehicleCore,
	}, nil
}

// DoesHashExist verifica se hash existe
func (c *Client) DoesHashExist(metadataHash [32]byte) (bool, error) {
	exists, err := c.contract.DoesHashExist(&bind.CallOpts{}, metadataHash)
	if err != nil {
		return false, fmt.Errorf("failed to check if hash exists: %w", err)
	}

	return exists, nil
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

	// Validar totalValue
	if totalValue == nil || totalValue.Sign() <= 0 {
		return nil, fmt.Errorf("totalValue deve ser maior que zero")
	}

	// Verificar se totalValue cabe em uint128 (2^128 - 1)
	maxUint128 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	if totalValue.Cmp(maxUint128) > 0 {
		return nil, fmt.Errorf("totalValue excede o limite de uint128")
	}

	auth, err := c.createTransactor()
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Criar objeto ABI a partir da string para simulação
	contractABI, err := abi.JSON(strings.NewReader(bindings.BindingsABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Fazer simulação ANTES da transação para obter valores de retorno
	callData, err := contractABI.Pack("registerContract",
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
		return nil, fmt.Errorf("failed to pack call data: %w", err)
	}

	// Simular a chamada para obter valores de retorno
	callResult, err := c.ethClient.CallContract(context.Background(), ethereum.CallMsg{
		From: auth.From,
		To:   &c.contractAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to simulate contract call: %w", err)
	}

	// Decodificar os valores de retorno da simulação
	unpacked, err := contractABI.Unpack("registerContract", callResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack return values from simulation: %w", err)
	}

	// Verificar se temos os valores esperados
	if len(unpacked) != 2 {
		return nil, fmt.Errorf("expected 2 return values, got %d", len(unpacked))
	}

	// Extrair tokenId
	var tokenId *big.Int
	if unpacked[0] != nil {
		if tokenIdValue, ok := unpacked[0].(*big.Int); ok {
			tokenId = tokenIdValue
		} else {
			return nil, fmt.Errorf("failed to parse tokenId from return values")
		}
	}

	// Extrair metadataHash
	var metadataHash [32]byte
	if unpacked[1] != nil {
		if metadataHashBytes, ok := unpacked[1].([32]byte); ok {
			metadataHash = metadataHashBytes
		} else {
			return nil, fmt.Errorf("failed to parse metadataHash from return values")
		}
	}

	// Agora executar a transação real
	tx, err := c.contract.RegisterContract(auth, registryId, contractNumber, contractDate, chassis, licensePlate, totalValue, brandName, modelName)
	if err != nil {
		return nil, fmt.Errorf("failed to register contract: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	// Extrair tokenId e metadataHash dos logs da transação
	tokenId, metadataHash, err = c.extractValuesFromLogs(receipt)
	if err != nil {
		return nil, fmt.Errorf("failed to extract tokenId and metadataHash from logs: %w", err)
	}

	// Retornar resultado com valores obtidos da simulação
	return &RegistrationResult{
		TokenId:      tokenId,
		MetadataHash: metadataHash,
		TxHash:       tx.Hash().Hex(),
	}, nil
}

// UpdateMetadataHash atualiza o hash de metadados de um contrato
func (c *Client) UpdateMetadataHash(tokenId *big.Int, newMetadataHash [32]byte) (string, error) {
	if tokenId == nil {
		return "", ErrInvalidInput
	}

	auth, err := c.createTransactor()
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Usar o binding do contrato
	tx, err := c.contract.UpdateMetadataHash(auth, tokenId, newMetadataHash)
	if err != nil {
		return "", fmt.Errorf("failed to update metadata hash: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return tx.Hash().Hex(), nil
}

// UpdateStatus atualiza o status de um contrato (ativo/inativo)
func (c *Client) UpdateStatus(tokenId *big.Int, active bool) (string, error) {
	if tokenId == nil {
		return "", ErrInvalidInput
	}

	auth, err := c.createTransactor()
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Usar o binding do contrato
	tx, err := c.contract.UpdateStatus(auth, tokenId, active)
	if err != nil {
		return "", fmt.Errorf("failed to update status: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return tx.Hash().Hex(), nil
}

// GetContractByTokenId obtém os dados do contrato pelo token ID
func (c *Client) GetContractByTokenId(tokenId *big.Int) (*ContractInfo, error) {
	contractRecord, vehicleCore, err := c.contract.GetContract(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract by token ID: %w", err)
	}

	return &ContractInfo{
		Contract: contractRecord,
		Vehicle:  vehicleCore,
	}, nil
}

// GetContractByRegistryId obtém os dados do contrato pelo registry ID
func (c *Client) GetContractByRegistryId(registryId string) (*ContractInfo, error) {
	contractRecord, vehicleCore, err := c.contract.GetContractByRegistryId(&bind.CallOpts{}, registryId)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract by registry ID: %w", err)
	}

	return &ContractInfo{
		Contract: contractRecord,
		Vehicle:  vehicleCore,
	}, nil
}

// GetContractByChassis obtém os dados do contrato pelo chassi
func (c *Client) GetContractByChassis(chassis string) (*ContractInfo, error) {
	contractRecord, vehicleCore, err := c.contract.GetContractByChassis(&bind.CallOpts{}, chassis)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract by chassis: %w", err)
	}

	return &ContractInfo{
		Contract: contractRecord,
		Vehicle:  vehicleCore,
	}, nil
}

// GetActiveContracts obtém a lista de contratos ativos (tokenIds)
func (c *Client) GetActiveContracts(offset, limit uint64) ([]*big.Int, error) {
	tokenIds, err := c.contract.GetActiveContracts(&bind.CallOpts{}, big.NewInt(int64(offset)), big.NewInt(int64(limit)))
	if err != nil {
		return nil, fmt.Errorf("failed to get active contracts: %w", err)
	}

	return tokenIds, nil
}

// DoesContractExist verifica se um contrato existe pelo registry ID
func (c *Client) DoesContractExist(registryId string) (bool, error) {
	exists, err := c.contract.DoesContractExist(&bind.CallOpts{}, registryId)
	if err != nil {
		return false, fmt.Errorf("failed to check if contract exists: %w", err)
	}

	return exists, nil
}

// GetMetadataUrl obtém a URL dos metadados pelo hash
func (c *Client) GetMetadataUrl(metadataHash [32]byte) (string, error) {
	url, err := c.contract.GetMetadataUrl(&bind.CallOpts{}, metadataHash)
	if err != nil {
		return "", fmt.Errorf("failed to get metadata URL: %w", err)
	}

	return url, nil
}

// GetMetadataUrlByRegistryId obtém a URL dos metadados pelo registry ID
func (c *Client) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	url, err := c.contract.GetMetadataUrlByRegistryId(&bind.CallOpts{}, registryId)
	if err != nil {
		return "", fmt.Errorf("failed to get metadata URL by registry ID: %w", err)
	}

	return url, nil
}

// GetBrandName obtém o nome da marca pelo ID
func (c *Client) GetBrandName(brandId uint64) (string, error) {
	brandName, err := c.contract.GetBrandName(&bind.CallOpts{}, brandId)
	if err != nil {
		return "", fmt.Errorf("failed to get brand name: %w", err)
	}

	return brandName, nil
}

// GetModelName obtém o nome do modelo pelo ID
func (c *Client) GetModelName(modelId uint64) (string, error) {
	modelName, err := c.contract.GetModelName(&bind.CallOpts{}, modelId)
	if err != nil {
		return "", fmt.Errorf("failed to get model name: %w", err)
	}

	return modelName, nil
}

// GetTotalSupply obtém o total de tokens emitidos
func (c *Client) GetTotalSupply() (*big.Int, error) {
	total, err := c.contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: %w", err)
	}

	return total, nil
}

// GetContractVersion obtém a versão do contrato
func (c *Client) GetContractVersion() (string, error) {
	version, err := c.contract.GetVersion(&bind.CallOpts{})
	if err != nil {
		return "", fmt.Errorf("failed to get contract version: %w", err)
	}

	return version, nil
}

// UpdateServerConfig atualiza a configuração do servidor (apenas admin)
func (c *Client) UpdateServerConfig(newMetadataBaseUrl string, newServerAddress common.Address) (string, error) {
	// Criar transactor
	auth, err := c.createTransactor()
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Usar o binding do contrato
	tx, err := c.contract.UpdateServerConfig(auth, newMetadataBaseUrl, newServerAddress)
	if err != nil {
		return "", fmt.Errorf("failed to update server config: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return tx.Hash().Hex(), nil
}

// RegisterBrand registra uma nova marca (apenas admin)
func (c *Client) RegisterBrand(brandName string) (string, error) {
	// Criar transactor
	auth, err := c.createTransactor()
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Usar o binding do contrato
	tx, err := c.contract.RegisterBrand(auth, brandName)
	if err != nil {
		return "", fmt.Errorf("failed to register brand: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return tx.Hash().Hex(), nil
}

// RegisterModel registra um novo modelo (apenas admin)
func (c *Client) RegisterModel(brandId uint64, modelName string) (string, error) {
	// Criar transactor
	auth, err := c.createTransactor()
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	// Usar o binding do contrato
	tx, err := c.contract.RegisterModel(auth, modelName)
	if err != nil {
		return "", fmt.Errorf("failed to register model: %w", err)
	}

	// Aguardar confirmação da transação
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return "", fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return tx.Hash().Hex(), nil
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
	TokenId        *big.Int
	RegistryIdHash [32]byte
	ChassisHash    [32]byte
	MetadataHash   [32]byte
	Timestamp      uint32
}

func (c *Client) parseContractRegisteredEvent(log *types.Log) (*ContractRegisteredEvent, error) {
	event := &ContractRegisteredEvent{}

	// Parse indexed topics (tokenId, registryIdHash, chassisHash, metadataHash)
	if len(log.Topics) >= 4 {
		event.TokenId = log.Topics[1].Big()
		copy(event.RegistryIdHash[:], log.Topics[2][:])
		copy(event.ChassisHash[:], log.Topics[3][:])

		// Se temos 5 topics, o metadataHash está no Topic 4
		if len(log.Topics) >= 5 {
			copy(event.MetadataHash[:], log.Topics[4][:])
		}
	}

	// Se não conseguimos extrair o metadataHash dos topics, tentar do log.Data
	if event.MetadataHash == [32]byte{} {
		var unpacked []interface{}

		// Criar objeto ABI a partir da string
		contractABI, err := abi.JSON(strings.NewReader(bindings.BindingsABI))
		if err != nil {
			// Se não conseguir criar o ABI, usar fallback
			copy(event.MetadataHash[:], event.RegistryIdHash[:])
		} else {
			err = contractABI.UnpackIntoInterface(&unpacked, ContractRegisteredEventName, log.Data)
			if err != nil {
				// Se não conseguir desempacotar, vamos gerar o hash baseado nos dados que temos
				// Usar o registryIdHash como base para o metadataHash
				copy(event.MetadataHash[:], event.RegistryIdHash[:])
			} else {
				if len(unpacked) >= 1 {
					if metadataHash, ok := unpacked[0].([32]byte); ok {
						event.MetadataHash = metadataHash
					}
				}
			}
		}
	}

	// Definir timestamp padrão se não conseguirmos extrair
	if event.Timestamp == 0 {
		event.Timestamp = uint32(time.Now().Unix())
	}

	return event, nil
}

// Converter ContractInfo para o modelo usado pela API
func (c *Client) ContractInfoToModel(info *ContractInfo) *models.ContractRecord {
	metadataHashHex := hex.EncodeToString(info.Contract.MetadataHash[:])

	return &models.ContractRecord{
		TokenId:        "",
		RegConId:       bytes32ToString(info.Contract.RegistryId),
		NumeroContrato: bytes32ToString(info.Contract.ContractNumber),
		DataContrato:   strconv.FormatUint(uint64(info.Contract.ContractDate), 10),
		MetadataHash:   metadataHashHex,
		Timestamp:      uint64(info.Contract.Timestamp),
		RegisteredBy:   info.Contract.RegisteredBy.Hex(),
		Active:         info.Contract.Active,
	}
}

// Converter ContractInfo para BlockchainContractRecord
func (c *Client) ContractInfoToBlockchainModel(info *ContractInfo) *models.BlockchainContractRecord {
	metadataHashHex := hex.EncodeToString(info.Contract.MetadataHash[:])

	return &models.BlockchainContractRecord{
		RegistryId:     bytes32ToString(info.Contract.RegistryId),
		ContractNumber: bytes32ToString(info.Contract.ContractNumber),
		ContractDate:   strconv.FormatUint(uint64(info.Contract.ContractDate), 10),
		MetadataHash:   metadataHashHex,
		Timestamp:      uint64(info.Contract.Timestamp),
		RegisteredBy:   info.Contract.RegisteredBy.Hex(),
		Active:         info.Contract.Active,
	}
}

// Converter RegistrationResult para modelo da resposta
func (c *Client) RegistrationResultToResponse(result *RegistrationResult, regConId string) *models.ContractRegistrationResponse {
	metadataHashHex := hex.EncodeToString(result.MetadataHash[:])

	return &models.ContractRegistrationResponse{
		Success:      true,
		Message:      "Contrato registrado com sucesso",
		RegConId:     regConId,
		MetadataHash: metadataHashHex,
		TxHash:       result.TxHash,
	}
}

// Converter ContractInfo para resposta formatada
func (c *Client) ContractInfoToFormattedResponse(info *ContractInfo, tokenId *big.Int) *models.BlockchainContractResponse {
	metadataHashHex := hex.EncodeToString(info.Contract.MetadataHash[:])

	return &models.BlockchainContractResponse{
		Success: true,
		Data: models.BlockchainContractData{
			TokenId: tokenId.String(),
			ContractRecord: models.BlockchainContractRecord{
				RegistryId:     bytes32ToString(info.Contract.RegistryId),
				ContractNumber: bytes32ToString(info.Contract.ContractNumber),
				ContractDate:   strconv.FormatUint(uint64(info.Contract.ContractDate), 10),
				MetadataHash:   metadataHashHex,
				Timestamp:      uint64(info.Contract.Timestamp),
				RegisteredBy:   info.Contract.RegisteredBy.Hex(),
				Active:         info.Contract.Active,
			},
			VehicleCore: models.BlockchainVehicleCore{
				Chassis:      bytes32ToString(info.Vehicle.Chassis),
				LicensePlate: bytes32ToString(info.Vehicle.LicensePlate),
				TotalValue:   info.Vehicle.TotalValue.String(),
				BrandId:      info.Vehicle.BrandId,
				ModelId:      info.Vehicle.ModelId,
			},
		},
	}
}

// GetTokenIdByRegistryId obtém o token ID pelo registry ID
func (c *Client) GetTokenIdByRegistryId(registryId string) (*big.Int, error) {
	// Converter registryId para bytes32
	var registryIdBytes [32]byte
	copy(registryIdBytes[:], []byte(registryId))

	tokenId, err := c.contract.RegistryIdHashToTokenId(&bind.CallOpts{}, registryIdBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get token ID by registry ID: %w", err)
	}

	return tokenId, nil
}

// GetTokenIdByChassis obtém o token ID pelo chassi
func (c *Client) GetTokenIdByChassis(chassis string) (*big.Int, error) {
	// Converter chassis para bytes32
	var chassisBytes [32]byte
	copy(chassisBytes[:], []byte(chassis))

	tokenId, err := c.contract.ChassisHashToTokenId(&bind.CallOpts{}, chassisBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get token ID by chassis: %w", err)
	}

	return tokenId, nil
}

// GetTokenIdByHash obtém o token ID pelo metadata hash
func (c *Client) GetTokenIdByHash(metadataHash [32]byte) (*big.Int, error) {
	tokenId, err := c.contract.MetadataHashToTokenId(&bind.CallOpts{}, metadataHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get token ID by hash: %w", err)
	}

	return tokenId, nil
}

func (c *Client) extractValuesFromLogs(receipt *types.Receipt) (*big.Int, [32]byte, error) {
	var tokenId *big.Int
	var metadataHash [32]byte

	// Procurar pelo evento ContractRegistered nos logs
	for _, vLog := range receipt.Logs {
		// Verificar se o log é do nosso contrato
		if vLog.Address != c.contractAddress {
			continue
		}

		// Usar os bindings para fazer parse do evento
		event, err := c.contract.ParseContractRegistered(*vLog)
		if err != nil {
			continue // Tentar próximo log
		}

		// Extrair os valores do evento
		tokenId = event.TokenId
		metadataHash = event.MetadataHash
		break
	}

	if tokenId == nil {
		return nil, [32]byte{}, fmt.Errorf("tokenId not found in transaction logs")
	}

	return tokenId, metadataHash, nil
}
