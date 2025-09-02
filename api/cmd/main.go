// @title MobX API - VFinance Registry
// @version 3.0.0
// @description API para registro e consulta de contratos de financiamento automotivo utilizando blockchain Hyperledger Besu e contratos inteligentes UUPS
// @termsOfService http://swagger.io/terms/

// @contact.name MobX API Support
// @contact.email support@mobx-api.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host 144.22.179.183:3000
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer" seguido de um espaço e o token JWT

// @tag.name health
// @tag.description Health check da API

// @tag.name auth
// @tag.description Autenticação JWT

// @tag.name blockchain
// @tag.description Consultas diretas ao contrato inteligente - apenas dados on-chain

// @tag.name contracts
// @tag.description Operações híbridas que combinam dados on-chain e off-chain

// @tag.name metadata
// @tag.description Operações CRUD para metadados de veículos

// @tag.name write-operations
// @tag.description Operações de escrita no contrato inteligente - requer autenticação

package main

import (
	"log"
	"vfinance-api/internal/config"
	"vfinance-api/internal/database"
	"vfinance-api/internal/server"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar configurações:", err)
	}

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	// Executar migrações
	if err := database.Migrate(db); err != nil {
		log.Fatal("Erro nas migrações:", err)
	}

	// Iniciar servidor
	srv := server.New(cfg, db)
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := srv.Start(); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
