package handlers

import (
	"math/big"
	"net/http"
	"strconv"
	"vfinance-api/internal/models"
	"vfinance-api/internal/services"

	"github.com/gin-gonic/gin"
)

type ContractHandler struct {
	contractService *services.ContractService
}

func NewContractHandler(contractService *services.ContractService) *ContractHandler {
	return &ContractHandler{contractService: contractService}
}

func (h *ContractHandler) GetContract(c *gin.Context) {
	regConId := c.Param("regConId")
	if regConId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "regConId é obrigatório"})
		return
	}

	contractData, err := h.contractService.GetCompleteContract(regConId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}

	c.JSON(http.StatusOK, contractData)
}

func (h *ContractHandler) GetActiveContracts(c *gin.Context) {
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

	response, err := h.contractService.GetActiveContracts(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContractHandler) GetContractByHash(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	contractData, err := h.contractService.GetContractByHash(hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}

	c.JSON(http.StatusOK, contractData)
}

func (h *ContractHandler) GetStats(c *gin.Context) {
	response, err := h.contractService.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ContractHandler) RegisterContract(c *gin.Context) {
	var request models.ContractRegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar campos obrigatórios
	if request.RegConId == "" || request.NumeroContrato == "" || request.DataContrato == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "regConId, numeroContrato e dataContrato são obrigatórios"})
		return
	}

	// Registrar contrato
	response, err := h.contractService.RegisterContract(
		request.RegConId,
		request.NumeroContrato,
		request.DataContrato,
		request.VehicleData,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetContractByTokenId busca contrato pelo token ID
func (h *ContractHandler) GetContractByTokenId(c *gin.Context) {
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

	contractData, err := h.contractService.GetContractByTokenId(tokenId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}

	c.JSON(http.StatusOK, contractData)
}

// GetContractByChassis busca contrato pelo chassi
func (h *ContractHandler) GetContractByChassis(c *gin.Context) {
	chassis := c.Param("chassis")
	if chassis == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chassi é obrigatório"})
		return
	}

	contractData, err := h.contractService.GetContractByChassis(chassis)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}

	c.JSON(http.StatusOK, contractData)
}

// GetMetadataUrl obtém a URL dos metadados pelo hash
func (h *ContractHandler) GetMetadataUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	url, err := h.contractService.GetMetadataUrl(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MetadataUrlResponse{
		Success: true,
		Url:     url,
	})
}

// GetMetadataUrlByRegistryId obtém a URL dos metadados pelo registry ID
func (h *ContractHandler) GetMetadataUrlByRegistryId(c *gin.Context) {
	registryId := c.Param("registryId")
	if registryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registry ID é obrigatório"})
		return
	}

	url, err := h.contractService.GetMetadataUrlByRegistryId(registryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MetadataUrlResponse{
		Success: true,
		Url:     url,
	})
}
