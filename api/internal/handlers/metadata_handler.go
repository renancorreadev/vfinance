package handlers

import (
	"net/http"
	"vfinance-api/internal/models"
	"vfinance-api/internal/services"

	"github.com/gin-gonic/gin"
)

type MetadataHandler struct {
	metadataService *services.MetadataService
}

func NewMetadataHandler(metadataService *services.MetadataService) *MetadataHandler {
	return &MetadataHandler{metadataService: metadataService}
}

// @Summary Store Metadata
// @Description Armazena metadados de um veículo
// @Tags metadata
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param hash path string true "Hash dos metadados"
// @Param request body object true "Dados do veículo"
// @Success 201 {object object "Metadados armazenados com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/metadata/{hash} [post]
func (h *MetadataHandler) StoreMetadata(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	var vehicleData models.VehicleData
	if err := c.ShouldBindJSON(&vehicleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.metadataService.StoreMetadata(hash, vehicleData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Metadados armazenados com sucesso"})
}

// @Summary Get Metadata
// @Description Busca metadados de um veículo pelo hash
// @Tags metadata
// @Accept json
// @Produce json
// @Param hash path string true "Hash dos metadados"
// @Success 200 {object object "Metadados encontrados"
// @Failure 400 {object object "Hash inválido"
// @Failure 404 {object object "Metadados não encontrados"
// @Router /api/metadata/{hash} [get]
func (h *MetadataHandler) GetMetadata(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	vehicleData, err := h.metadataService.GetMetadata(hash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metadados não encontrados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": vehicleData})
}

// @Summary Update Metadata
// @Description Atualiza metadados de um veículo
// @Tags metadata
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param hash path string true "Hash dos metadados"
// @Param request body object true "Dados do veículo"
// @Success 200 {object object "Metadados atualizados com sucesso"
// @Failure 400 {object object "Dados inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/metadata/{hash} [put]
func (h *MetadataHandler) UpdateMetadata(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	var vehicleData models.VehicleData
	if err := c.ShouldBindJSON(&vehicleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.metadataService.UpdateMetadata(hash, vehicleData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Metadados atualizados com sucesso"})
}

// @Summary Delete Metadata
// @Description Remove metadados de um veículo
// @Tags metadata
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param hash path string true "Hash dos metadados"
// @Success 200 {object object "Metadados removidos com sucesso"
// @Failure 400 {object object "Hash inválido"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/metadata/{hash} [delete]
func (h *MetadataHandler) DeleteMetadata(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash é obrigatório"})
		return
	}

	if err := h.metadataService.DeleteMetadata(hash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Metadados removidos com sucesso"})
}
