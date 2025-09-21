package models

// AuthTokenResponse representa a resposta do endpoint de geração de token
type AuthTokenResponse struct {
	Token     string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Success   bool   `json:"success" example:"true"`
	ExpiresIn int64  `json:"expires_in" example:"86400"`
}

// AuthValidateResponse representa a resposta do endpoint de validação de token
type AuthValidateResponse struct {
	Valid   bool   `json:"valid" example:"true"`
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Token válido"`
	UserID  string `json:"user_id" example:"admin"`
	Role    string `json:"role" example:"admin"`
	Exp     int64  `json:"exp" example:"1756217832"`
}

// AuthErrorResponse representa a resposta de erro dos endpoints de autenticação
type AuthErrorResponse struct {
	Valid   bool   `json:"valid" example:"false"`
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Token inválido ou expirado"`
	Error   string `json:"error" example:"Unauthorized"`
}

// Error representa uma resposta de erro genérica
type Error struct {
	Error   string `json:"error" example:"Erro interno do servidor"`
	Details string `json:"details,omitempty" example:"Detalhes do erro"`
}

// VehicleDataExample representa um exemplo de dados do veículo
type VehicleDataExample struct {
	RegConId                              string `json:"regConId" example:"b8d899ac-4a02-4e2d-a3d5-cada74d30927"`
	NumeroContrato                        string `json:"numeroContrato" example:"521"`
	DataContrato                          string `json:"dataContrato" example:"2024-01-15"`
	CnpjAgenteFinanceiro                  string `json:"cnpjAgenteFinanceiro" example:"03817784000133"`
	NomeAgenteFinanceiro                  string `json:"nomeAgenteFinanceiro" example:"NATAL CAR VEICULOS LTDA"`
	EnderecoAgenteFinanceiro              string `json:"enderecoAgenteFinanceiro" example:"Rua das Flores, 123"`
	NumeroEnderecoAgenteFinanceiro        string `json:"numeroEnderecoAgenteFinanceiro" example:"123"`
	ComplementoEnderecoAgenteFinanceiro   string `json:"complementoEnderecoAgenteFinanceiro" example:"Sala 101"`
	BairroEnderecoAgenteFinanceiro        string `json:"bairroEnderecoAgenteFinanceiro" example:"Centro"`
	NomeMunicipioEnderecoAgenteFinanceiro string `json:"nomeMunicipioEnderecoAgenteFinanceiro" example:"Natal"`
	UfEnderecoAgenteFinanceiro            string `json:"ufEnderecoAgenteFinanceiro" example:"RN"`
	CepEnderecoAgenteFinanceiro           string `json:"cepEnderecoFinanceiro" example:"59000-000"`
	TelefoneAgenteFinanceiro              string `json:"telefoneAgenteFinanceiro" example:"(84) 3333-4444"`
	EmailAgenteFinanceiro                 string `json:"emailAgenteFinanceiro" example:"contato@natalcar.com.br"`
	CpfCnpjProprietario                   string `json:"cpfCnpjProprietario" example:"12345678901"`
	NomeProprietario                      string `json:"nomeProprietario" example:"João da Silva"`
	EnderecoProprietario                  string `json:"enderecoProprietario" example:"Av. Principal, 456"`
	NumeroEnderecoProprietario            string `json:"numeroEnderecoProprietario" example:"456"`
	BairroEnderecoProprietario            string `json:"bairroEnderecoProprietario" example:"Lagoa Nova"`
	NomeMunicipioProprietario             string `json:"nomeMunicipioProprietario" example:"Natal"`
	UfEnderecoProprietario                string `json:"ufEnderecoProprietario" example:"RN"`
	CepEnderecoProprietario               string `json:"cepEnderecoProprietario" example:"59075-000"`
	TelefoneProprietario                  string `json:"telefoneProprietario" example:"(84) 9999-8888"`
	EmailProprietario                     string `json:"emailProprietario" example:"joao@email.com"`
	VeiculoZeroKm                         bool   `json:"veiculoZeroKm" example:"true"`
	ChassiVeiculo                         string `json:"chassiVeiculo" example:"n609l"`
	ChassiRemarcadoVeiculo                string `json:"chassiRemarcadoVeiculo" example:""`
	PlacaVeiculo                          string `json:"placaVeiculo" example:"ABC55"`
	TipoPlacaVeiculo                      string `json:"tipoPlacaVeiculo" example:"MERCOSUL"`
	UfPlacaVeiculo                        string `json:"ufPlacaVeiculo" example:"RN"`
	RenavamVeiculo                        string `json:"renavamVeiculo" example:"485"`
	AnoFabricacaoVeiculo                  string `json:"anoFabricacaoVeiculo" example:"2024"`
	AnoModeloVeiculo                      string `json:"anoModeloVeiculo" example:"2024"`
	NumeroRestricaoVeiculo                string `json:"numeroRestricaoVeiculo" example:"715"`
	EspecieVeiculo                        string `json:"especieVeiculo" example:"AUTOMOVEL"`
	MarcaVeiculo                          string `json:"marcaVeiculo" example:"TOYOTA"`
	ModeloVeiculo                         string `json:"modeloVeiculo" example:"COROLLA ALTIS 2.0"`
	TipoRestricacaoContrato               string `json:"tipoRestricacaoContrato" example:"ALIENACAO FIDUCIARIA"`
	UfRegistroContrato                    string `json:"ufRegistroContrato" example:"RN"`
	CnpjResponsavelPeloRegistro           string `json:"cnpjResponsavelPeloRegistro" example:"03817784000133"`
	ValorTotalContrato                    string `json:"valorTotalContrato" example:"85000.0"`
	ValorParcelaContrato                  string `json:"valorParcelaContrato" example:"1416.67"`
	QuantidadeParcelasContrato            string `json:"quantidadeParcelasContrato" example:"60"`
	TaxaJurosMesContrato                  string `json:"taxaJurosMesContrato" example:"1.2"`
	TaxaJurosMesAnoContrato               string `json:"taxaJurosMesAnoContrato" example:"15.39"`
	PossuiJurosMultaContrato              string `json:"possuiJurosMultaContrato" example:"S"`
	TaxaJurosMultaContrato                string `json:"taxaJurosMultaContrato" example:"2.0"`
	PossuiJurosMoraDiaContrato            string `json:"possuiJurosMoraDiaContrato" example:"S"`
	TaxaJurosMoraDiaContrato              string `json:"taxaJurosMoraDiaContrato" example:"0.033"`
	ValorCustoRegistroContrato            string `json:"valorCustoRegistroContrato" example:"150.00"`
	ValorIofContrato                      string `json:"valorIofContrato" example:"425.00"`
	DataVencimentoPrimeiraParcelaContrato string `json:"dataVencimentoPrimeiraParcelaContrato" example:"2024-02-15"`
	DataVencimentoUltimaParcelaContrato   string `json:"dataVencimentoUltimaParcelaContrato" example:"2029-01-15"`
	DataLiberacaoCreditoContrato          string `json:"dataLiberacaoCreditoContrato" example:"2024-01-15"`
	CidadeLiberacaoCreditoContrato        string `json:"cidadeLiberacaoCreditoContrato" example:"Natal"`
	UfLiberacaoCreditoContrato            string `json:"ufLiberacaoCreditoContrato" example:"RN"`
	IndiceCorrecaoContrato                string `json:"indiceCorrecaoContrato" example:"IPCA"`
	NumeroGrupoConsorcioContrato          string `json:"numeroGrupoConsorcioContrato" example:""`
	NumeroCotaConsorcioContrato           string `json:"numeroCotaConsorcioContrato" example:""`
	IndicativoPenalidadeContrato          string `json:"indicativoPenalidadeContrato" example:"S"`
	PenalidadeContrato                    string `json:"penalidadeContrato" example:"3.0"`
	IndicativoComissaoContrato            string `json:"indicativoComissaoContrato" example:"S"`
	ComissaoContrato                      string `json:"comissaoContrato" example:"2.5"`
	CategoriaVeiculo                      string `json:"categoriaVeiculo" example:"PASSEIO"`
}
