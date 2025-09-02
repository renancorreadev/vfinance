package models

import "math/big"

// =============================================================
//                    BLOCKCHAIN RESPONSE MODELS
// =============================================================

// BlockchainContractResponse representa resposta de contrato on-chain
type BlockchainContractResponse struct {
	Success bool                   `json:"success"`
	Data    BlockchainContractData `json:"data"`
}

// BlockchainContractData representa dados do contrato on-chain
type BlockchainContractData struct {
	TokenId        string                   `json:"tokenId"`
	ContractRecord BlockchainContractRecord `json:"contractRecord"`
	VehicleCore    BlockchainVehicleCore    `json:"vehicleCore"`
}

// BlockchainContractRecord representa dados do contrato on-chain
type BlockchainContractRecord struct {
	RegistryId     string `json:"registryId"`
	ContractNumber string `json:"contractNumber"`
	ContractDate   string `json:"contractDate"`
	MetadataHash   string `json:"metadataHash"`
	Timestamp      uint64 `json:"timestamp"`
	RegisteredBy   string `json:"registeredBy"`
	Active         bool   `json:"active"`
}

// BlockchainVehicleCore representa dados do veículo on-chain
type BlockchainVehicleCore struct {
	Chassis      string `json:"chassis"`
	LicensePlate string `json:"licensePlate"`
	TotalValue   string `json:"totalValue"`
	BrandId      uint64 `json:"brandId"`
	ModelId      uint64 `json:"modelId"`
}

// BlockchainActiveContractsResponse representa resposta de contratos ativos on-chain
type BlockchainActiveContractsResponse struct {
	Success bool                          `json:"success"`
	Data    BlockchainActiveContractsData `json:"data"`
}

// BlockchainActiveContractsData representa dados de contratos ativos on-chain
type BlockchainActiveContractsData struct {
	TokenIds []*big.Int `json:"tokenIds"`
	Total    int        `json:"total"`
	Offset   uint64     `json:"offset"`
	Limit    uint64     `json:"limit"`
}

// BlockchainTotalSupplyResponse representa resposta do total supply on-chain
type BlockchainTotalSupplyResponse struct {
	Success bool                      `json:"success"`
	Data    BlockchainTotalSupplyData `json:"data"`
}

// BlockchainTotalSupplyData representa dados do total supply on-chain
type BlockchainTotalSupplyData struct {
	TotalSupply string `json:"totalSupply"`
}

// BlockchainExistsResponse representa resposta de verificação de existência on-chain
type BlockchainExistsResponse struct {
	Success bool                 `json:"success"`
	Data    BlockchainExistsData `json:"data"`
}

// BlockchainExistsData representa dados de verificação de existência on-chain
type BlockchainExistsData struct {
	RegistryId string `json:"registryId"`
	Exists     bool   `json:"exists"`
}

// BlockchainHashExistsResponse representa resposta de verificação de hash on-chain
type BlockchainHashExistsResponse struct {
	Success bool                     `json:"success"`
	Data    BlockchainHashExistsData `json:"data"`
}

// BlockchainHashExistsData representa dados de verificação de hash on-chain
type BlockchainHashExistsData struct {
	Hash   string `json:"hash"`
	Exists bool   `json:"exists"`
}

// BlockchainBrandResponse representa resposta de marca on-chain
type BlockchainBrandResponse struct {
	Success bool                `json:"success"`
	Data    BlockchainBrandData `json:"data"`
}

// BlockchainBrandData representa dados de marca on-chain
type BlockchainBrandData struct {
	BrandId   uint64 `json:"brandId"`
	BrandName string `json:"brandName"`
}

// BlockchainModelResponse representa resposta de modelo on-chain
type BlockchainModelResponse struct {
	Success bool                `json:"success"`
	Data    BlockchainModelData `json:"data"`
}

// BlockchainModelData representa dados de modelo on-chain
type BlockchainModelData struct {
	ModelId   uint64 `json:"modelId"`
	ModelName string `json:"modelName"`
}

// BlockchainMetadataUrlResponse representa resposta de URL de metadados on-chain
type BlockchainMetadataUrlResponse struct {
	Success bool                      `json:"success"`
	Data    BlockchainMetadataUrlData `json:"data"`
}

// BlockchainMetadataUrlData representa dados de URL de metadados on-chain
type BlockchainMetadataUrlData struct {
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

// BlockchainVersionResponse representa resposta de versão do contrato on-chain
type BlockchainVersionResponse struct {
	Success bool                  `json:"success"`
	Data    BlockchainVersionData `json:"data"`
}

// BlockchainVersionData representa dados de versão do contrato on-chain
type BlockchainVersionData struct {
	Version string `json:"version"`
}
