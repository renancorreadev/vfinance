package services

import (
	"encoding/hex"
	"encoding/json"
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
	// Primeiro verificar se o regConId é um hash hexadecimal (começa com 0x)
	if strings.HasPrefix(regConId, "0x") {
		// Se for um hash, tentar buscar na blockchain por registryId primeiro
		if contractInfo, err := s.blockchainClient.GetContractByRegistryId(regConId); err == nil {
			// Encontrou na blockchain
			onChainData := *s.blockchainClient.ContractInfoToModel(contractInfo)

			// Buscar metadados off-chain
			offChainData, err := s.metadataService.GetMetadata(onChainData.MetadataHash)
			if err != nil {
				// Se não encontrar metadados, criar dados básicos
				offChainData = &models.VehicleData{
					RegConId:       regConId,
					NumeroContrato: onChainData.NumeroContrato,
					DataContrato:   onChainData.DataContrato,
				}
			}

			response := &models.CompleteContractData{
				Success: true,
			}
			response.Data.OnChain = onChainData
			response.Data.OffChain = *offChainData

			return response, nil
		}

		// Se não encontrou por registryId, tentar buscar por tokenId
		// Vamos tentar buscar nos primeiros 100 tokens para encontrar o contrato
		for tokenId := 1; tokenId <= 100; tokenId++ {
			if contractInfo, err := s.blockchainClient.GetContractByTokenId(big.NewInt(int64(tokenId))); err == nil {
				// Converter o registryId de bytes32 para string
				registryId := s.blockchainClient.Bytes32ToString(contractInfo.Contract.RegistryId)

				// Se o registryId convertido for igual ao regConId fornecido
				if registryId == regConId {
					// Encontrou o contrato correto
					onChainData := *s.blockchainClient.ContractInfoToModel(contractInfo)

					// Buscar metadados off-chain
					offChainData, err := s.metadataService.GetMetadata(onChainData.MetadataHash)
					if err != nil {
						// Se não encontrar metadados, criar dados básicos
						offChainData = &models.VehicleData{
							RegConId:       regConId,
							NumeroContrato: onChainData.NumeroContrato,
							DataContrato:   onChainData.DataContrato,
						}
					}

					response := &models.CompleteContractData{
						Success: true,
					}
					response.Data.OnChain = onChainData
					response.Data.OffChain = *offChainData

					return response, nil
				}
			}
		}
	}

	// Se não for hash ou não encontrou na blockchain, buscar no banco local
	var registry models.ContractRegistry
	if err := s.db.Where("reg_con_id = ?", regConId).First(&registry).Error; err != nil {
		return nil, fmt.Errorf("contrato não encontrado: %w", err)
	}

	// Se encontrou na tabela de contratos, buscar metadados
	offChainData, err := s.metadataService.GetMetadata(registry.MetadataHash)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados: %w", err)
	}

	// Tentar buscar dados on-chain da blockchain usando o tokenId
	var onChainData models.ContractRecord
	if registry.TokenId != "" && registry.TokenId != "0" {
		// Se temos tokenId, buscar na blockchain
		if tokenId, ok := new(big.Int).SetString(registry.TokenId, 10); ok {
			if contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId); err == nil {
				// Se conseguiu buscar na blockchain, usar esses dados
				onChainData = *s.blockchainClient.ContractInfoToModel(contractInfo)
				// Garantir que o tokenId seja preenchido
				onChainData.TokenId = registry.TokenId
				onChainData.RegConId = registry.RegConId
				// ✅ IMPORTANTE: Usar o metadataHash do banco local, NÃO da blockchain!
				onChainData.MetadataHash = registry.MetadataHash
			} else {
				// Se não conseguir, criar dados básicos do banco local
				onChainData = models.ContractRecord{
					TokenId:        registry.TokenId,
					RegConId:       registry.RegConId,
					NumeroContrato: offChainData.NumeroContrato,
					DataContrato:   offChainData.DataContrato,
					MetadataHash:   registry.MetadataHash,
					Timestamp:      uint64(registry.CreatedAt.Unix()),
					RegisteredBy:   "blockchain",
					Active:         registry.Status == "active",
				}
			}
		}
	} else {
		// Se não temos tokenId, criar dados básicos do banco local
		onChainData = models.ContractRecord{
			TokenId:        registry.TokenId,
			RegConId:       registry.RegConId,        // ✅ Usar regConId do banco local
			NumeroContrato: offChainData.NumeroContrato,
			DataContrato:   offChainData.DataContrato,
			MetadataHash:   registry.MetadataHash,
			Timestamp:      uint64(registry.CreatedAt.Unix()),
			RegisteredBy:   "local",
			Active:         registry.Status == "active",
		}
	}

	// Criar resposta
	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetContractByTokenId(tokenId *big.Int) (*models.CompleteContractData, error) {
	// Tentar buscar dados on-chain da blockchain primeiro
	var onChainData models.ContractRecord
	var offChainData *models.VehicleData

	if contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId); err == nil {
		// Se conseguiu buscar na blockchain, usar esses dados
		onChainData = *s.blockchainClient.ContractInfoToModel(contractInfo)

		// Adicionar o tokenId que estava faltando
		onChainData.TokenId = tokenId.String()

		// Buscar metadados off-chain
		var err2 error
		offChainData, err2 = s.metadataService.GetMetadata(onChainData.MetadataHash)
		if err2 != nil {
			return nil, fmt.Errorf("erro ao buscar metadados: %w", err2)
		}

		// Buscar informações de marca e modelo da blockchain
		brandName, _ := s.blockchainClient.GetBrandName(contractInfo.Vehicle.BrandId)
		modelName, _ := s.blockchainClient.GetModelName(contractInfo.Vehicle.ModelId)

		// Adicionar os nomes aos dados off-chain
		offChainData.BrandName = brandName
		offChainData.ModelName = modelName

		// Formatar dados do veículo para melhor legibilidade
		offChainData.Chassis = s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis)
		offChainData.LicensePlate = s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate)
		offChainData.TotalValue = contractInfo.Vehicle.TotalValue.String()
	} else {
		// Se não conseguir na blockchain, buscar no banco local
		// Buscar por token_id que corresponda ao tokenId
		var registry models.ContractRegistry
		if err := s.db.Where("token_id = ?", tokenId.String()).First(&registry).Error; err != nil {
			return nil, fmt.Errorf("contrato não encontrado: %w", err)
		}

		// Buscar metadados
		var err2 error
		offChainData, err2 = s.metadataService.GetMetadata(registry.MetadataHash)
		if err2 != nil {
			return nil, fmt.Errorf("erro ao buscar metadados: %w", err2)
		}

		// Para dados locais, usar os nomes já presentes nos metadados
		offChainData.BrandName = offChainData.MarcaVeiculo
		offChainData.ModelName = offChainData.ModeloVeiculo

		// Criar dados on-chain básicos
		onChainData = models.ContractRecord{
			RegConId:       registry.RegConId,
			NumeroContrato: offChainData.NumeroContrato,
			DataContrato:   offChainData.DataContrato,
			MetadataHash:   registry.MetadataHash,
			Timestamp:      uint64(registry.CreatedAt.Unix()),
			RegisteredBy:   "local",
			Active:         registry.Status == "active",
		}
	}

	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetContractByChassis(chassis string) (*models.CompleteContractData, error) {
	// Primeiro tentar buscar dados on-chain da blockchain
	var onChainData models.ContractRecord
	var offChainData *models.VehicleData

	if contractInfo, err := s.blockchainClient.GetContractByChassis(chassis); err == nil {
		// Se conseguiu buscar na blockchain, usar esses dados
		onChainData = *s.blockchainClient.ContractInfoToModel(contractInfo)

		// Buscar metadados off-chain
		var err2 error
		offChainData, err2 = s.metadataService.GetMetadata(onChainData.MetadataHash)
		if err2 != nil {
			// Se não encontrar metadados, criar dados básicos
			offChainData = &models.VehicleData{
				RegConId:       onChainData.RegConId,
				NumeroContrato: onChainData.NumeroContrato,
				DataContrato:   onChainData.DataContrato,
				ChassiVeiculo:  chassis,
			}
		}

		// Buscar informações de marca e modelo da blockchain
		brandName, _ := s.blockchainClient.GetBrandName(contractInfo.Vehicle.BrandId)
		modelName, _ := s.blockchainClient.GetModelName(contractInfo.Vehicle.ModelId)

		// Adicionar os nomes aos dados off-chain
		offChainData.BrandName = brandName
		offChainData.ModelName = modelName

		// Formatar dados do veículo para melhor legibilidade
		offChainData.Chassis = s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis)
		offChainData.LicensePlate = s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate)
		offChainData.TotalValue = contractInfo.Vehicle.TotalValue.String()
	} else {
		// Se não conseguir na blockchain, buscar no banco local
		// Buscar por chassi nos metadados
		var metadata models.VehicleMetadata
		if err := s.db.Where("vehicle_data->>'chassiVeiculo' = ?", chassis).First(&metadata).Error; err != nil {
			return nil, fmt.Errorf("contrato não encontrado: %w", err)
		}

		// Extrair dados do JSON
		var vehicleData models.VehicleData
		if err := json.Unmarshal(metadata.VehicleData, &vehicleData); err != nil {
			return nil, fmt.Errorf("erro ao deserializar metadados: %w", err)
		}

		offChainData = &vehicleData

		// Verificar se existe na tabela de contratos
		var registry models.ContractRegistry
		if err := s.db.Where("metadata_hash = ?", metadata.Hash).First(&registry).Error; err == nil {
			// Encontrou na tabela de contratos
			if registry.TokenId != "" && registry.TokenId != "0" {
				// Se temos tokenId, tentar buscar na blockchain
				if tokenId, ok := new(big.Int).SetString(registry.TokenId, 10); ok {
					if contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId); err == nil {
						// Se conseguiu buscar na blockchain, usar esses dados
						onChainData = *s.blockchainClient.ContractInfoToModel(contractInfo)
						onChainData.TokenId = registry.TokenId
					} else {
						// Se não conseguir, criar dados básicos do banco local
						onChainData = models.ContractRecord{
							TokenId:        registry.TokenId,
							RegConId:       registry.RegConId,
							NumeroContrato: vehicleData.NumeroContrato,
							DataContrato:   vehicleData.DataContrato,
							MetadataHash:   registry.MetadataHash,
							Timestamp:      uint64(registry.CreatedAt.Unix()),
							RegisteredBy:   "blockchain",
							Active:         registry.Status == "active",
						}
					}
				}
			} else {
				// Se não temos tokenId, criar dados básicos do banco local
				onChainData = models.ContractRecord{
					TokenId:        registry.TokenId,
					RegConId:       registry.RegConId,
					NumeroContrato: vehicleData.NumeroContrato,
					DataContrato:   vehicleData.DataContrato,
					MetadataHash:   registry.MetadataHash,
					Timestamp:      uint64(registry.CreatedAt.Unix()),
					RegisteredBy:   "local",
					Active:         registry.Status == "active",
				}
			}
		} else {
			// Não encontrou na tabela de contratos, criar dados básicos
			onChainData = models.ContractRecord{
				RegConId:       vehicleData.RegConId,
				NumeroContrato: vehicleData.NumeroContrato,
				DataContrato:   vehicleData.DataContrato,
				MetadataHash:   metadata.Hash,
				Timestamp:      uint64(metadata.CreatedAt.Unix()),
				RegisteredBy:   "local",
				Active:         true,
			}
		}
	}

	// Criar resposta
	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetActiveContracts(offset, limit uint64) (*models.ActiveContractsResponse, error) {
	// Primeiro tentar buscar da blockchain
	var contracts []models.ContractSummary

	if tokenIds, err := s.blockchainClient.GetActiveContracts(offset, limit); err == nil {
		// Se conseguiu buscar da blockchain
		for _, tokenId := range tokenIds {
			// Buscar informações completas do contrato
			if contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId); err == nil {
				// Converter para modelo da API
				onChainData := s.blockchainClient.ContractInfoToModel(contractInfo)

				// Buscar informações de marca e modelo
				brandName, _ := s.blockchainClient.GetBrandName(contractInfo.Vehicle.BrandId)
				modelName, _ := s.blockchainClient.GetModelName(contractInfo.Vehicle.ModelId)

				// Converter o registryId de bytes32 para string legível
				registryId := s.blockchainClient.Bytes32ToString(contractInfo.Contract.RegistryId)

				summary := models.ContractSummary{
					TokenId:        tokenId.String(),
					RegConId:       registryId, // Usar o registryId real, não o hash
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
		}
	}

	// Se não conseguiu buscar da blockchain ou não retornou nada, buscar do banco local
	if len(contracts) == 0 {
		var registries []models.ContractRegistry
		if err := s.db.Where("status = ?", "active").Offset(int(offset)).Limit(int(limit)).Find(&registries).Error; err == nil {
			for _, registry := range registries {
				// Buscar metadados
				if metadata, err := s.metadataService.GetMetadata(registry.MetadataHash); err == nil {
					summary := models.ContractSummary{
						TokenId:        registry.TokenId,
						RegConId:       registry.RegConId,
						NumeroContrato: metadata.NumeroContrato,
						DataContrato:   metadata.DataContrato,
						MetadataHash:   registry.MetadataHash,
						TotalValue:     metadata.ValorTotalContrato,
						BrandName:      metadata.MarcaVeiculo,
						ModelName:      metadata.ModeloVeiculo,
						Active:         registry.Status == "active",
						Timestamp:      uint64(registry.CreatedAt.Unix()),
					}

					contracts = append(contracts, summary)
				}
			}
		}
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
	// Primeiro tentar buscar na blockchain por metadata hash
	hashBytes, err := s.blockchainClient.HexStringToBytes32(hash)
	if err != nil {
		return nil, fmt.Errorf("hash inválido: %w", err)
	}

	if contractInfo, err := s.blockchainClient.GetContractByHash(hashBytes); err == nil {
		// Encontrou na blockchain
		onChainData := *s.blockchainClient.ContractInfoToModel(contractInfo)

		// Buscar metadados off-chain
		offChainData, err := s.metadataService.GetMetadata(hash)
		if err != nil {
			// Se não encontrar metadados, criar dados básicos
			offChainData = &models.VehicleData{
				RegConId:       onChainData.RegConId,
				NumeroContrato: onChainData.NumeroContrato,
				DataContrato:   onChainData.DataContrato,
			}
		}

		response := &models.CompleteContractData{
			Success: true,
		}
		response.Data.OnChain = onChainData
		response.Data.OffChain = *offChainData

		return response, nil
	}

	// Se não encontrou na blockchain, buscar no banco local
	offChainData, err := s.metadataService.GetMetadata(hash)
	if err != nil {
		return nil, fmt.Errorf("metadados não encontrados: %w", err)
	}

	// Buscar no banco local se existir
	var registry models.ContractRegistry
	var onChainData models.ContractRecord

	if err := s.db.First(&registry, "metadata_hash = ?", hash).Error; err == nil {
		// Encontrou na tabela de contratos
		if registry.TokenId != "" && registry.TokenId != "0" {
			// Se temos tokenId, tentar buscar na blockchain
			if tokenId, ok := new(big.Int).SetString(registry.TokenId, 10); ok {
				if contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId); err == nil {
					// Se conseguiu buscar na blockchain, usar esses dados
					onChainData = *s.blockchainClient.ContractInfoToModel(contractInfo)
					onChainData.TokenId = registry.TokenId
				} else {
					// Se não conseguir, criar dados básicos do banco local
					onChainData = models.ContractRecord{
						TokenId:        registry.TokenId,
						RegConId:       registry.RegConId,
						NumeroContrato: offChainData.NumeroContrato,
						DataContrato:   offChainData.DataContrato,
						MetadataHash:   registry.MetadataHash,
						Timestamp:      uint64(registry.CreatedAt.Unix()),
						RegisteredBy:   "blockchain",
						Active:         registry.Status == "active",
					}
				}
			}
		} else {
			// Se não temos tokenId, criar dados básicos do banco local
			onChainData = models.ContractRecord{
				TokenId:        registry.TokenId,
				RegConId:       registry.RegConId,
				NumeroContrato: offChainData.NumeroContrato,
				DataContrato:   offChainData.DataContrato,
				MetadataHash:   registry.MetadataHash,
				Timestamp:      uint64(registry.CreatedAt.Unix()),
				RegisteredBy:   "local",
				Active:         registry.Status == "active",
			}
		}
	} else {
		// Não encontrou, criar dados básicos
		onChainData = models.ContractRecord{
			RegConId:       offChainData.RegConId,
			NumeroContrato: offChainData.NumeroContrato,
			DataContrato:   offChainData.DataContrato,
			MetadataHash:   hash,
			Timestamp:      uint64(time.Now().Unix()),
			RegisteredBy:   "local",
			Active:         true,
		}
	}

	response := &models.CompleteContractData{
		Success: true,
	}
	response.Data.OnChain = onChainData
	response.Data.OffChain = *offChainData

	return response, nil
}

func (s *ContractService) GetStats() (*models.StatsResponse, error) {
	// Tentar buscar total de contratos na blockchain
	var totalContracts uint64
	if totalSupply, err := s.blockchainClient.GetTotalSupply(); err == nil {
		totalContracts = totalSupply.Uint64()
	} else {
		// Se não conseguir, contar do banco local
		var count int64
		s.db.Model(&models.ContractRegistry{}).Count(&count)
		totalContracts = uint64(count)
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
			TotalContracts:  totalContracts,
			ActiveContracts: uint64(activeCount),
			ContractVersion: version,
		},
	}, nil
}

func (s *ContractService) RegisterContract(regConId, numeroContrato, dataContrato string, vehicleData models.VehicleData) (*models.ContractRegistrationResponse, error) {
	// Verificar se o regConId é válido (não pode ser vazio)
	if regConId == "" {
		return nil, fmt.Errorf("regConId não pode ser vazio")
	}

	// Verificar se o contrato já existe
	exists, err := s.blockchainClient.DoesContractExist(regConId)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar existência do contrato: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("contrato com regConId %s já existe", regConId)
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

	// Registrar contrato na blockchain usando o regConId original
	result, err := s.blockchainClient.RegisterContract(
		regConId,        // registryId
		numeroContrato,  // contractNumber
		vehicleData.ChassiVeiculo,    // chassis
		vehicleData.PlacaVeiculo,     // licensePlate
		vehicleData.MarcaVeiculo,     // brandName
		vehicleData.ModeloVeiculo,    // modelName
		contractDate,    // contractDate
		totalValue,      // totalValue
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar contrato na blockchain: %w", err)
	}

	// Converter metadataHash retornado pelo contrato para string hexadecimal
	metadataHash := hex.EncodeToString(result.MetadataHash[:])

	// Armazenar metadados no banco de dados utilizando o hash gerado on-chain
	err = s.metadataService.StoreMetadata(metadataHash, vehicleData)
	if err != nil {
		return nil, fmt.Errorf("erro ao armazenar metadados: %w", err)
	}

	// Criar registro no banco local com tokenId da blockchain
	tokenIdStr := "0"
	if result.TokenId != nil {
		tokenIdStr = result.TokenId.String() // ✅ Armazenar tokenId da blockchain
	}

	// Criar o registro local vinculando regConId, tokenId e metadataHash
	registry := models.ContractRegistry{
		RegConId:     regConId,        // ✅ regConId original (qualquer string que você enviar)
		TokenId:      tokenIdStr,      // ✅ tokenId da blockchain
		MetadataHash: metadataHash,    // ✅ metadataHash gerado pelo smart contract (NÃO gerar localmente!)
		BlockchainTx: result.TxHash,   // ✅ hash da transação
		Status:       "active",
	}

	err = s.db.Create(&registry).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao criar registro local: %w", err)
	}

	// Retornar resposta com todos os dados vinculados
	return &models.ContractRegistrationResponse{
		Success:      true,
		Message:      "Contrato registrado com sucesso",
		RegConId:     regConId,        // ✅ regConId original
		MetadataHash: metadataHash,    // ✅ metadataHash da blockchain
		TxHash:       result.TxHash,   // ✅ hash da transação
	}, nil
}

func (s *ContractService) GetMetadataUrl(metadataHash string) (string, error) {
	// Converter string para bytes32
	hash, err := s.blockchainClient.HexStringToBytes32(metadataHash)
	if err != nil {
		return "", fmt.Errorf("hash inválido: %w", err)
	}

	return s.blockchainClient.GetMetadataUrl(hash)
}

func (s *ContractService) GetMetadataUrlByRegistryId(registryId string) (string, error) {
	return s.blockchainClient.GetMetadataUrlByRegistryId(registryId)
}

// SyncBlockchainData sincroniza dados da blockchain com o banco local
func (s *ContractService) SyncBlockchainData() error {
	// Buscar todos os contratos ativos na blockchain
	tokenIds, err := s.blockchainClient.GetActiveContracts(0, 1000) // Buscar até 1000 contratos
	if err != nil {
		return fmt.Errorf("erro ao buscar contratos ativos na blockchain: %w", err)
	}

	for _, tokenId := range tokenIds {
		// Buscar dados completos do contrato
		contractInfo, err := s.blockchainClient.GetContractByTokenId(tokenId)
		if err != nil {
			continue
		}

		// Converter metadataHash para string
		metadataHash := hex.EncodeToString(contractInfo.Contract.MetadataHash[:])

		// Verificar se os metadados já existem no banco local
		_, err = s.metadataService.GetMetadata(metadataHash)
		if err != nil {
			// Metadados não existem, criar dados básicos
			vehicleData := models.VehicleData{
				RegConId:       s.blockchainClient.Bytes32ToString(contractInfo.Contract.RegistryId),
				NumeroContrato: s.blockchainClient.Bytes32ToString(contractInfo.Contract.ContractNumber),
				DataContrato:   strconv.FormatUint(uint64(contractInfo.Contract.ContractDate), 10),
				ChassiVeiculo:  s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis),
				PlacaVeiculo:   s.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate),
				ValorTotalContrato: contractInfo.Vehicle.TotalValue.String(),
				MarcaVeiculo:   "N/A", // Será preenchido depois
				ModeloVeiculo:  "N/A", // Será preenchido depois
			}

			// Buscar nomes de marca e modelo
			if brandName, err := s.blockchainClient.GetBrandName(contractInfo.Vehicle.BrandId); err == nil {
				vehicleData.MarcaVeiculo = brandName
			}
			if modelName, err := s.blockchainClient.GetModelName(contractInfo.Vehicle.ModelId); err == nil {
				vehicleData.ModeloVeiculo = modelName
			}

			// Salvar metadados básicos
			err = s.metadataService.StoreMetadata(metadataHash, vehicleData)
			if err != nil {
				fmt.Printf("Erro ao salvar metadados para hash %s: %v\n", metadataHash, err)
				continue
			}

			// Verificar se já existe registro local
			var existingRegistry models.ContractRegistry
			if err := s.db.First(&existingRegistry, "metadata_hash = ?", metadataHash).Error; err != nil {
				// Criar registro local se não existir
				registry := models.ContractRegistry{
					RegConId:     vehicleData.RegConId,
					TokenId:      tokenId.String(),
					MetadataHash: metadataHash,
					BlockchainTx: "synced",
					Status:       "active",
				}

				if err := s.db.Create(&registry).Error; err != nil {
					fmt.Printf("Erro ao criar registro local para hash %s: %v\n", metadataHash, err)
				}
			}
		}
	}

	return nil
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
			timestamp := uint32(t.Unix())
			return timestamp, nil
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
		result := big.NewInt(intVal)
		return result, nil
	}

	// Tentar conversão direta para inteiro
	if intVal, ok := new(big.Int).SetString(cleanValue, 10); ok {
		return intVal, nil
	}

	return nil, fmt.Errorf("não foi possível converter valor: %s", valueStr)
}
