package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"vfinance-api/internal/blockchain"
	"vfinance-api/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type BlockchainHandler struct {
	blockchainClient *blockchain.Client
}

func NewBlockchainHandler(blockchainClient *blockchain.Client) *BlockchainHandler {
	return &BlockchainHandler{blockchainClient: blockchainClient}
}

// @Summary Get Contract by Token ID
// @Description Busca contrato on-chain por Token ID do ERC721
// @Tags blockchain
// @Accept json
// @Produce json
// @Param tokenId path string true "Token ID do contrato"
// @Success 200 {object object "Contrato encontrado"
// @Failure 400 {object object "Token ID inválido"
// @Failure 404 {object object "Contrato não encontrado"
// @Router /api/blockchain/contract/token/{tokenId} [get]
func (h *BlockchainHandler) GetContractByTokenId(c *gin.Context) {
	tokenIdStr := c.Param("tokenId")
	if tokenIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token ID é obrigatório"})
		return
	}

	tokenId, ok := new(big.Int).SetString(tokenIdStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token ID inválido"})
		return
	}

	contractInfo, err := h.blockchainClient.GetContractByTokenId(tokenId)
	if err != nil {
		fmt.Printf("Erro ao buscar contrato tokenId %s: %v\n", tokenIdStr, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Contrato não encontrado na blockchain",
			"details": err.Error(),
		})
		return
	}

	// Converter para modelo da API
	onChainData := h.blockchainClient.ContractInfoToBlockchainModel(contractInfo)

	response := &models.BlockchainContractResponse{
		Success: true,
		Data: models.BlockchainContractData{
			TokenId:        tokenId.String(),
			ContractRecord: *onChainData,
			VehicleCore: models.BlockchainVehicleCore{
				Chassis:      h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis),
				LicensePlate: h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate),
				TotalValue:   contractInfo.Vehicle.TotalValue.String(),
				BrandId:      contractInfo.Vehicle.BrandId,
				ModelId:      contractInfo.Vehicle.ModelId,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Contract by Registry ID
// @Description Busca contrato on-chain por Registry ID
// @Tags blockchain
// @Accept json
// @Produce json
// @Param registryId path string true "Registry ID do contrato"
// @Success 200 {object object "Contrato encontrado"
// @Failure 400 {object object "Registry ID inválido"
// @Failure 404 {object object "Contrato não encontrado"
// @Router /api/blockchain/contract/registry/{registryId} [get]
func (h *BlockchainHandler) GetContractByRegistryId(c *gin.Context) {
	registryId := c.Param("registryId")
	if registryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registry ID é obrigatório"})
		return
	}

	contractInfo, err := h.blockchainClient.GetContractByRegistryId(registryId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado na blockchain"})
		return
	}

	// Converter para modelo da API
	onChainData := h.blockchainClient.ContractInfoToBlockchainModel(contractInfo)

	response := &models.BlockchainContractResponse{
		Success: true,
		Data: models.BlockchainContractData{
			TokenId:        "0", // Não temos token ID neste endpoint
			ContractRecord: *onChainData,
			VehicleCore: models.BlockchainVehicleCore{
				Chassis:      h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis),
				LicensePlate: h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate),
				TotalValue:   contractInfo.Vehicle.TotalValue.String(),
				BrandId:      contractInfo.Vehicle.BrandId,
				ModelId:      contractInfo.Vehicle.ModelId,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetContractByHash obtém dados on-chain por metadata hash
func (h *BlockchainHandler) GetContractByHash(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	// Converter string para bytes32
	hashBytes, err := h.blockchainClient.HexStringToBytes32(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash inválido"})
		return
	}

	contractInfo, err := h.blockchainClient.GetContractByHash(hashBytes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado na blockchain"})
		return
	}

	// Converter para modelo da API
	onChainData := h.blockchainClient.ContractInfoToBlockchainModel(contractInfo)

	response := &models.BlockchainContractResponse{
		Success: true,
		Data: models.BlockchainContractData{
			TokenId:        "0", // Não temos token ID neste endpoint
			ContractRecord: *onChainData,
			VehicleCore: models.BlockchainVehicleCore{
				Chassis:      h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis),
				LicensePlate: h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate),
				TotalValue:   contractInfo.Vehicle.TotalValue.String(),
				BrandId:      contractInfo.Vehicle.BrandId,
				ModelId:      contractInfo.Vehicle.ModelId,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetContractByChassis obtém dados on-chain por chassi
func (h *BlockchainHandler) GetContractByChassis(c *gin.Context) {
	chassis := c.Param("chassis")
	if chassis == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chassi é obrigatório"})
		return
	}

	contractInfo, err := h.blockchainClient.GetContractByChassis(chassis)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado na blockchain"})
		return
	}

	// Converter para modelo da API
	onChainData := h.blockchainClient.ContractInfoToBlockchainModel(contractInfo)

	response := &models.BlockchainContractResponse{
		Success: true,
		Data: models.BlockchainContractData{
			TokenId:        "0", // Não temos token ID neste endpoint
			ContractRecord: *onChainData,
			VehicleCore: models.BlockchainVehicleCore{
				Chassis:      h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.Chassis),
				LicensePlate: h.blockchainClient.Bytes32ToString(contractInfo.Vehicle.LicensePlate),
				TotalValue:   contractInfo.Vehicle.TotalValue.String(),
				BrandId:      contractInfo.Vehicle.BrandId,
				ModelId:      contractInfo.Vehicle.ModelId,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Active Contracts
// @Description Lista contratos ativos on-chain com paginação
// @Tags blockchain
// @Accept json
// @Produce json
// @Param offset query int false "Offset para paginação" default(0)
// @Param limit query int false "Limite de resultados" default(10)
// @Success 200 {object object "Lista de contratos ativos"
// @Failure 400 {object object "Parâmetros inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/contracts/active [get]
func (h *BlockchainHandler) GetActiveContracts(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Offset inválido"})
		return
	}

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit inválido"})
		return
	}

	tokenIds, err := h.blockchainClient.GetActiveContracts(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &models.BlockchainActiveContractsResponse{
		Success: true,
		Data: models.BlockchainActiveContractsData{
			TokenIds: tokenIds,
			Total:    len(tokenIds),
			Offset:   offset,
			Limit:    limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Total Supply
// @Description Obtém total de contratos on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Success 200 {object object "Total de contratos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/contracts/total [get]
func (h *BlockchainHandler) GetTotalSupply(c *gin.Context) {
	totalSupply, err := h.blockchainClient.GetTotalSupply()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &models.BlockchainTotalSupplyResponse{
		Success: true,
		Data: models.BlockchainTotalSupplyData{
			TotalSupply: totalSupply.String(),
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Check Contract Exists
// @Description Verifica se contrato existe on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param registryId path string true "Registry ID do contrato"
// @Success 200 {object object "Status de existência"
// @Failure 400 {object object "Registry ID inválido"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/contract/exists/{registryId} [get]
func (h *BlockchainHandler) DoesContractExist(c *gin.Context) {
	registryId := c.Param("registryId")
	if registryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registry ID é obrigatório"})
		return
	}

	exists, err := h.blockchainClient.DoesContractExist(registryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &models.BlockchainExistsResponse{
		Success: true,
		Data: models.BlockchainExistsData{
			RegistryId: registryId,
			Exists:     exists,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Check Hash Exists
// @Description Verifica se hash existe on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param hash path string true "Hash para verificar"
// @Success 200 {object object "Status de existência"
// @Failure 400 {object object "Hash inválido"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/hash/exists/{hash} [get]
func (h *BlockchainHandler) DoesHashExist(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	// Converter string para bytes32
	hashBytes, err := h.blockchainClient.HexStringToBytes32(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash inválido"})
		return
	}

	exists, err := h.blockchainClient.DoesHashExist(hashBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &models.BlockchainHashExistsResponse{
		Success: true,
		Data: models.BlockchainHashExistsData{
			Hash:   hash,
			Exists: exists,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Brand Name
// @Description Obtém nome da marca por ID on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param brandId path int true "ID da marca"
// @Success 200 {object object "Nome da marca"
// @Failure 400 {object object "Brand ID inválido"
// @Failure 404 {object object "Marca não encontrada"
// @Router /api/blockchain/brand/{brandId} [get]
func (h *BlockchainHandler) GetBrandName(c *gin.Context) {
	brandIdStr := c.Param("brandId")
	if brandIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand ID é obrigatório"})
		return
	}

	brandId, err := strconv.ParseUint(brandIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand ID inválido"})
		return
	}

	brandName, err := h.blockchainClient.GetBrandName(brandId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Marca não encontrada"})
		return
	}

	response := &models.BlockchainBrandResponse{
		Success: true,
		Data: models.BlockchainBrandData{
			BrandId:   brandId,
			BrandName: brandName,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Model Name
// @Description Obtém nome do modelo por ID on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param modelId path int true "ID do modelo"
// @Success 200 {object object "Nome do modelo"
// @Failure 400 {object object "Model ID inválido"
// @Failure 404 {object object "Modelo não encontrado"
// @Router /api/blockchain/model/{modelId} [get]
func (h *BlockchainHandler) GetModelName(c *gin.Context) {
	modelIdStr := c.Param("modelId")
	if modelIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Model ID é obrigatório"})
		return
	}

	modelId, err := strconv.ParseUint(modelIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Model ID inválido"})
		return
	}

	modelName, err := h.blockchainClient.GetModelName(modelId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modelo não encontrado"})
		return
	}

	response := &models.BlockchainModelResponse{
		Success: true,
		Data: models.BlockchainModelData{
			ModelId:   modelId,
			ModelName: modelName,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Metadata URL by Hash
// @Description Obtém URL dos metadados por hash on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param hash path string true "Hash dos metadados"
// @Success 200 {object object "URL dos metadados"
// @Failure 400 {object object "Hash inválido"
// @Failure 404 {object object "URL não encontrada"
// @Router /api/blockchain/metadata-url/{hash} [get]
func (h *BlockchainHandler) GetMetadataUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	// Converter string para bytes32
	hashBytes, err := h.blockchainClient.HexStringToBytes32(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash inválido"})
		return
	}

	url, err := h.blockchainClient.GetMetadataUrl(hashBytes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"})
		return
	}

	response := &models.BlockchainMetadataUrlResponse{
		Success: true,
		Data: models.BlockchainMetadataUrlData{
			Hash: hash,
			Url:  url,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Metadata URL by Registry ID
// @Description Obtém URL dos metadados por registry ID on-chain
// @Tags blockchain
// @Accept json
// @Produce json
// @Param registryId path string true "Registry ID do contrato"
// @Success 200 {object object "URL dos metadados"
// @Failure 400 {object object "Registry ID inválido"
// @Failure 404 {object object "URL não encontrada"
// @Router /api/blockchain/metadata-url/registry/{registryId} [get]
func (h *BlockchainHandler) GetMetadataUrlByRegistryId(c *gin.Context) {
	registryId := c.Param("registryId")
	if registryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registry ID é obrigatório"})
		return
	}

	url, err := h.blockchainClient.GetMetadataUrlByRegistryId(registryId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"})
		return
	}

	response := &models.BlockchainMetadataUrlResponse{
		Success: true,
		Data: models.BlockchainMetadataUrlData{
			Hash: "", // Não temos hash neste endpoint
			Url:  url,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Contract Version
// @Description Obtém versão do contrato inteligente
// @Tags blockchain
// @Accept json
// @Produce json
// @Success 200 {object object "Versão do contrato"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/version [get]
func (h *BlockchainHandler) GetVersion(c *gin.Context) {
	version, err := h.blockchainClient.GetContractVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &models.BlockchainVersionResponse{
		Success: true,
		Data: models.BlockchainVersionData{
			Version: version,
		},
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update Metadata Hash
// @Description Atualiza o hash de metadados de um contrato existente
// @Tags write-operations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object true "Dados para atualização"
// @Success 200 {object object "Hash atualizado com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/contract/metadata [put]
func (h *BlockchainHandler) UpdateMetadataHash(c *gin.Context) {
	var req struct {
		TokenId         string `json:"tokenId" binding:"required"`
		NewMetadataHash string `json:"newMetadataHash" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converter tokenId para big.Int
	tokenId, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token ID inválido"})
		return
	}

	// Converter hash hex para bytes32
	hashBytes, err := h.blockchainClient.HexStringToBytes32(req.NewMetadataHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash de metadados inválido"})
		return
	}

	// Executar transação
	txHash, err := h.blockchainClient.UpdateMetadataHash(tokenId, hashBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar hash de metadados: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Hash de metadados atualizado com sucesso",
		"txHash":  txHash,
		"tokenId": req.TokenId,
	})
}

// @Summary Update Contract Status
// @Description Ativa ou desativa um contrato existente
// @Tags write-operations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object true "Dados para atualização"
// @Success 200 {object object "Status atualizado com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/contract/status [put]
func (h *BlockchainHandler) UpdateStatus(c *gin.Context) {
	var req struct {
		TokenId string `json:"tokenId" binding:"required"`
		Active  bool   `json:"active"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	// Converter tokenId para big.Int
	tokenId, ok := new(big.Int).SetString(req.TokenId, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token ID inválido"})
		return
	}

	// Executar transação
	txHash, err := h.blockchainClient.UpdateStatus(tokenId, req.Active)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar status: " + err.Error()})
		return
	}

	status := "inativo"
	if req.Active {
		status = "ativo"
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Status do contrato atualizado para " + status,
		"txHash":  txHash,
		"tokenId": req.TokenId,
		"active":  req.Active,
	})
}

// @Summary Update Server Config
// @Description Atualiza configurações do servidor (apenas admin)
// @Tags write-operations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object true "Configurações do servidor"
// @Success 200 {object object "Configuração atualizada com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/admin/config [put]
func (h *BlockchainHandler) UpdateServerConfig(c *gin.Context) {
	var req struct {
		ServerAddress   string `json:"serverAddress" binding:"required"`
		MetadataBaseUrl string `json:"metadataBaseUrl" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Converter endereço de string para common.Address
	serverAddr := common.HexToAddress(req.ServerAddress)

	// Chamar função do cliente blockchain - ordem correta dos parâmetros
	txHash, err := h.blockchainClient.UpdateServerConfig(req.MetadataBaseUrl, serverAddr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update server config",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Server config updated successfully",
		"txHash":  txHash,
	})
}

// @Summary Register Brand
// @Description Registra uma nova marca de veículo (apenas admin)
// @Tags write-operations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object true "Dados da marca"
// @Success 200 {object object "Marca registrada com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/admin/brand [post]
func (h *BlockchainHandler) RegisterBrand(c *gin.Context) {
	var req struct {
		BrandName string `json:"brandName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Chamar função do cliente blockchain
	txHash, err := h.blockchainClient.RegisterBrand(req.BrandName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to register brand",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Brand registered successfully",
		"txHash":  txHash,
	})
}

// @Summary Register Model
// @Description Registra um novo modelo de veículo (apenas admin)
// @Tags write-operations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object true "Dados do modelo"
// @Success 200 {object object "Modelo registrado com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/blockchain/admin/model [post]
func (h *BlockchainHandler) RegisterModel(c *gin.Context) {
	var req struct {
		BrandId   uint64 `json:"brandId" binding:"required"`
		ModelName string `json:"modelName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Chamar função do cliente blockchain
	txHash, err := h.blockchainClient.RegisterModel(req.BrandId, req.ModelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to register model",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Model registered successfully",
		"txHash":  txHash,
	})
}
