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

// @Summary Get Contract by Registry ID
// @Description Busca contrato híbrido por Registry ID
// @Tags contracts
// @Accept json
// @Produce json
// @Param regConId path string true "Registry ID do contrato"
// @Success 200 {object object "Contrato encontrado"
// @Failure 400 {object object "Registry ID inválido"
// @Failure 404 {object object "Contrato não encontrado"
// @Router /api/contracts/{regConId} [get]
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

// @Summary Get Active Contracts
// @Description Lista contratos ativos (híbrido)
// @Tags contracts
// @Accept json
// @Produce json
// @Param offset query int false "Offset para paginação" default(0)
// @Param limit query int false "Limite de resultados" default(10)
// @Success 200 {object object "Lista de contratos ativos"
// @Failure 400 {object object "Parâmetros inválidos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/contracts/active [get]
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

// @Summary Get Contract by Hash
// @Description Busca contrato híbrido por Metadata Hash
// @Tags contracts
// @Accept json
// @Produce json
// @Param hash path string true "Hash dos metadados"
// @Success 200 {object object "Contrato encontrado"
// @Failure 400 {object object "Hash inválido"
// @Failure 404 {object object "Contrato não encontrado"
// @Router /api/contracts/hash/{hash} [get]
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

// @Summary Get Contract Stats
// @Description Estatísticas dos contratos (híbrido)
// @Tags contracts
// @Accept json
// @Produce json
// @Success 200 {object object "Estatísticas dos contratos"
// @Failure 500 {object object "Erro interno do servidor"
// @Router /api/contracts/stats [get]
func (h *ContractHandler) GetStats(c *gin.Context) {
	response, err := h.contractService.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Register Contract
// @Description Registra novo contrato de financiamento automotivo (híbrido - blockchain + banco local)
// @Tags contracts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.ContractRegistrationExample true "Dados do contrato e veículo" example({"regConId":"b8d899ac-4a02-4e2d-a3d5-cada74d30927","numeroContrato":"668","dataContrato":"2024-01-15","vehicleData":{"regConId":"b8d899ac-4a02-4e2d-a3d5-cada74d30927","numeroContrato":"521","dataContrato":"2024-01-15","cnpjAgenteFinanceiro":"03817784000133","nomeAgenteFinanceiro":"NATAL CAR VEICULOS LTDA","enderecoAgenteFinanceiro":"Rua das Flores, 123","numeroEnderecoAgenteFinanceiro":"123","complementoEnderecoAgenteFinanceiro":"Sala 101","bairroEnderecoAgenteFinanceiro":"Centro","nomeMunicipioEnderecoAgenteFinanceiro":"Natal","ufEnderecoAgenteFinanceiro":"RN","cepEnderecoAgenteFinanceiro":"59000-000","telefoneAgenteFinanceiro":"(84) 3333-4444","emailAgenteFinanceiro":"contato@natalcar.com.br","cpfCnpjProprietario":"12345678901","nomeProprietario":"João da Silva","enderecoProprietario":"Av. Principal, 456","numeroEnderecoProprietario":"456","bairroEnderecoProprietario":"Lagoa Nova","nomeMunicipioProprietario":"Natal","ufEnderecoProprietario":"RN","cepEnderecoProprietario":"59075-000","telefoneProprietario":"(84) 9999-8888","emailProprietario":"joao@email.com","veiculoZeroKm":true,"chassiVeiculo":"n609l","chassiRemarcadoVeiculo":"","placaVeiculo":"ABC55","tipoPlacaVeiculo":"MERCOSUL","ufPlacaVeiculo":"RN","renavamVeiculo":"485","anoFabricacaoVeiculo":"2024","anoModeloVeiculo":"2024","numeroRestricaoVeiculo":"715","especieVeiculo":"AUTOMOVEL","marcaVeiculo":"TOYOTA","modeloVeiculo":"COROLLA ALTIS 2.0","tipoRestricacaoContrato":"ALIENACAO FIDUCIARIA","ufRegistroContrato":"RN","cnpjResponsavelPeloRegistro":"03817784000133","valorTotalContrato":"85000.0","valorParcelaContrato":"1416.67","quantidadeParcelasContrato":"60","taxaJurosMesContrato":"1.2","taxaJurosMesAnoContrato":"15.39","possuiJurosMultaContrato":"S","taxaJurosMultaContrato":"2.0","possuiJurosMoraDiaContrato":"S","taxaJurosMoraDiaContrato":"0.033","valorCustoRegistroContrato":"150.00","valorIofContrato":"425.00","dataVencimentoPrimeiraParcelaContrato":"2024-02-15","dataVencimentoUltimaParcelaContrato":"2029-01-15","dataLiberacaoCreditoContrato":"2024-01-15","cidadeLiberacaoCreditoContrato":"Natal","ufLiberacaoCreditoContrato":"RN","indiceCorrecaoContrato":"IPCA","numeroGrupoConsorcioContrato":"","numeroCotaConsorcioContrato":"","indicativoPenalidadeContrato":"S","penalidadeContrato":"3.0","indicativoComissaoContrato":"S","comissaoContrato":"2.5","categoriaVeiculo":"PASSEIO"}})
// @Success 201 {object} models.ContractRegistrationResponse "Contrato registrado com sucesso"
// @Failure 400 {object} object{error=string} "Dados inválidos ou campos obrigatórios ausentes"
// @Failure 401 {object} object{error=string} "Token de autenticação inválido ou ausente"
// @Failure 500 {object} object{error=string} "Erro interno do servidor ou falha no blockchain"
// @Router /api/contracts [post]
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

// SyncBlockchainData sincroniza dados da blockchain com o banco local
func (h *ContractHandler) SyncBlockchainData(c *gin.Context) {
	if err := h.contractService.SyncBlockchainData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sincronização concluída com sucesso",
	})
}
