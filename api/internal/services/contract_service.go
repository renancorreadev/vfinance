package services

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
	"vfinance-api/internal/blockchain"
	"vfinance-api/internal/models"

	"gorm.io/gorm"
)

type ContractService struct {
	db               *gorm.DB
	blockchainClient *blockchain.Client
	metadataService  *MetadataService
}

func NewContractService(db *gorm.DB, blockchainClient *blockchain.Client, metadataService *MetadataService) *ContractService {
	return &ContractService{
		db:               db,
		blockchainClient: blockchainClient,
		metadataService:  metadataService,
	}
}

func (s *ContractService) GetCompleteContract(regConId string) (*models.CompleteContractData, error) {
	// Buscar dados on-chain pelo registryId
	contractInfo, err := s.blockchainClient.GetContractByRegistryId(regConId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar contrato na blockchain: %w", err)
	}

	// Converter para modelo da API
	onChainData := s.blockchainClient.ContractInfoToModel(contractInfo)

	// Buscar metadados off-chain
	offChainData, err := s.metadataService.GetMetadata(onChainData.MetadataHash)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados: %w", err)
	}

	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = *onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetContractByTokenId(tokenId *big.Int) (*models.CompleteContractData, error) {
	// Buscar dados on-chain pelo token ID
	contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar contrato por token ID: %w", err)
	}

	// Converter para modelo da API
	onChainData := s.blockchainClient.ContractInfoToModel(contractInfo)

	// Buscar metadados off-chain
	offChainData, err := s.metadataService.GetMetadata(onChainData.MetadataHash)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados: %w", err)
	}

	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = *onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetContractByChassis(chassis string) (*models.CompleteContractData, error) {
	// Buscar dados on-chain pelo chassi
	contractInfo, err := s.blockchainClient.GetContractByChassis(chassis)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar contrato por chassi: %w", err)
	}

	// Converter para modelo da API
	onChainData := s.blockchainClient.ContractInfoToModel(contractInfo)

	// Buscar metadados off-chain
	offChainData, err := s.metadataService.GetMetadata(onChainData.MetadataHash)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados: %w", err)
	}

	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = *onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetActiveContracts(offset, limit uint64) (*models.ActiveContractsResponse, error) {
	// Buscar token IDs ativos
	tokenIds, err := s.blockchainClient.GetActiveContracts(offset, limit)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar contratos ativos: %w", err)
	}

	var contracts []models.ContractSummary
	for _, tokenId := range tokenIds {
		// Buscar informações básicas do contrato
		contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId)
		if err != nil {
			continue // Pular contratos com erro
		}

		// Converter para modelo da API
		onChainData := s.blockchainClient.ContractInfoToModel(contractInfo)

		// Buscar informações de marca e modelo
		brandName, _ := s.blockchainClient.GetBrandName(contractInfo.Vehicle.BrandId)
		modelName, _ := s.blockchainClient.GetModelName(contractInfo.Vehicle.ModelId)

		summary := models.ContractSummary{
			TokenId:        tokenId.String(),
			RegConId:       onChainData.RegConId,
			NumeroContrato: onChainData.NumeroContrato,
			DataContrato:   onChainData.DataContrato,
			MetadataHash:   onChainData.MetadataHash,
			TotalValue:     contractInfo.Vehicle.TotalValue.String(),
			BrandName:      brandName,
			ModelName:      modelName,
			Active:         onChainData.Active,
			Timestamp:      onChainData.Timestamp,
		}

		contracts = append(contracts, summary)
	}

	return &models.ActiveContractsResponse{
		Success: true,
		Data: models.ActiveContractsData{
			Contracts: contracts,
			Total:     len(contracts),
			Offset:    offset,
			Limit:     limit,
		},
	}, nil
}

func (s *ContractService) GetContractByHash(hash string) (*models.CompleteContractData, error) {
	// Buscar no banco local primeiro
	var registry models.ContractRegistry
	if err := s.db.First(&registry, "metadata_hash = ?", hash).Error; err != nil {
		return nil, fmt.Errorf("contrato não encontrado no banco local: %w", err)
	}

	return s.GetCompleteContract(registry.RegConId)
}

func (s *ContractService) GetStats() (*models.StatsResponse, error) {
	// Buscar total de contratos na blockchain
	totalSupply, err := s.blockchainClient.GetTotalSupply()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar total de contratos: %w", err)
	}

	// Buscar contratos ativos no banco local
	var activeCount int64
	s.db.Model(&models.ContractRegistry{}).Where("status = ?", "active").Count(&activeCount)

	// Buscar versão do contrato
	version, err := s.blockchainClient.GetContractVersion()
	if err != nil {
		version = "unknown"
	}

	return &models.StatsResponse{
		Success: true,
		Data: models.StatsData{
			TotalContracts:  totalSupply.Uint64(),
			ActiveContracts: uint64(activeCount),
			ContractVersion: version,
		},
	}, nil
}

func (s *ContractService) RegisterContract(regConId, numeroContrato, dataContrato string, vehicleData models.VehicleData) (*models.ContractRegistrationResponse, error) {
	// Verificar se o contrato já existe
	exists, err := s.blockchainClient.DoesContractExist(regConId)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar existência do contrato: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("contrato com regConId %s já existe", regConId)
	}

	// Gerar hash dos metadados
	metadataHash, err := s.metadataService.GenerateHash(vehicleData)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar hash dos metadados: %w", err)
	}

	// Converter data do contrato para timestamp
	contractDate, err := parseContractDate(dataContrato)
	if err != nil {
		return nil, fmt.Errorf("formato de data inválido: %w", err)
	}

	// Converter valor total para big.Int
	totalValue, err := parseValueFromString(vehicleData.ValorTotalContrato)
	if err != nil {
		return nil, fmt.Errorf("valor total inválido: %w", err)
	}

	// Registrar contrato na blockchain
	result, err := s.blockchainClient.RegisterContract(
		regConId,
		numeroContrato,
		vehicleData.ChassiVeiculo,
		vehicleData.PlacaVeiculo,
		vehicleData.MarcaVeiculo,
		vehicleData.ModeloVeiculo,
		contractDate,
		totalValue,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar contrato na blockchain: %w", err)
	}

	// Armazenar metadados no banco de dados
	err = s.metadataService.StoreMetadata(metadataHash, vehicleData)
	if err != nil {
		return nil, fmt.Errorf("erro ao armazenar metadados: %w", err)
	}

	// Criar registro no banco local
	registry := models.ContractRegistry{
		RegConId:     regConId,
		MetadataHash: metadataHash,
		BlockchainTx: result.TxHash,
		Status:       "active",
	}

	err = s.db.Create(&registry).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao criar registro local: %w", err)
	}

	// Usar a função de conversão do client
	return s.blockchainClient.RegistrationResultToResponse(result, regConId), nil
}

func (s *ContractService) GetMetadataUrl(metadataHash string) (string, error) {
	// Converter string para bytes32
	var hash [32]byte
	hashBytes, err := hexStringToBytes32(metadataHash)
	if err != nil {
		return "", fmt.Errorf("hash inválido: %w", err)
	}
	copy(hash[:], hashBytes)

	return s.blockchainClient.GetMetadataUrl(hash)
}

func (s *ContractService) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	return s.blockchainClient.GetMetadataUrlByRegistryId(registryId)
}

// Funções auxiliares
func parseContractDate(dateStr string) (uint32, error) {
	// Tentar diferentes formatos de data
	formats := []string{
		"2006-01-02",
		"02/01/2006",
		"2006-01-02T15:04:05Z",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return uint32(t.Unix()), nil
		}
	}

	// Se não conseguir parsear, tentar como timestamp Unix
	if timestamp, err := strconv.ParseInt(dateStr, 10, 64); err == nil {
		return uint32(timestamp), nil
	}

	return 0, fmt.Errorf("formato de data não suportado: %s", dateStr)
}

func parseValueFromString(valueStr string) (*big.Int, error) {
	// Remover caracteres não numéricos exceto ponto e vírgula
	cleanValue := ""
	for _, char := range valueStr {
		if (char >= '0' && char <= '9') || char == '.' || char == ',' {
			cleanValue += string(char)
		}
	}

	// Substituir vírgula por ponto (padrão brasileiro)
	cleanValue = strings.Replace(cleanValue, ",", ".", 1)

	// Converter para float e depois para big.Int (assumindo 2 casas decimais)
	if floatVal, err := strconv.ParseFloat(cleanValue, 64); err == nil {
		// Multiplicar por 100 para preservar centavos
		intVal := int64(floatVal * 100)
		return big.NewInt(intVal), nil
	}

	// Tentar conversão direta para inteiro
	if intVal, ok := new(big.Int).SetString(cleanValue, 10); ok {
		return intVal, nil
	}

	return nil, fmt.Errorf("não foi possível converter valor: %s", valueStr)
}

func hexStringToBytes32(hexStr string) ([]byte, error) {
	// Remover prefixo 0x se presente
	if len(hexStr) > 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	// Garantir que tenha 64 caracteres (32 bytes)
	if len(hexStr) != 64 {
		return nil, fmt.Errorf("hash deve ter exatamente 64 caracteres hexadecimais")
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hash hexadecimal inválido: %w", err)
	}

	return bytes, nil
}
