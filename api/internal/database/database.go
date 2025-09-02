package database

import (
	"fmt"
	"vfinance-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// Executar migração automática
	err := db.AutoMigrate(
		&models.VehicleMetadata{},
		&models.ContractRegistry{},
	)
	if err != nil {
		return err
	}

	// Verificar se a coluna token_id existe, se não, adicioná-la
	if !db.Migrator().HasColumn(&models.ContractRegistry{}, "token_id") {
		err = db.Migrator().AddColumn(&models.ContractRegistry{}, "token_id")
		if err != nil {
			return fmt.Errorf("erro ao adicionar coluna token_id: %w", err)
		}
	}

	// Criar índice na coluna token_id se não existir
	if !db.Migrator().HasIndex(&models.ContractRegistry{}, "idx_contract_registries_token_id") {
		err = db.Migrator().CreateIndex(&models.ContractRegistry{}, "token_id")
		if err != nil {
			return fmt.Errorf("erro ao criar índice token_id: %w", err)
		}
	}

	// Verificar se a coluna token_id existe, se não, adicioná-la
	if !db.Migrator().HasColumn(&models.ContractRegistry{}, "token_id") {
		err = db.Migrator().AddColumn(&models.ContractRegistry{}, "token_id")
		if err != nil {
			return fmt.Errorf("erro ao adicionar coluna token_id: %w", err)
		}
	}

	// Criar índice na coluna token_id se não existir
	if !db.Migrator().HasIndex(&models.ContractRegistry{}, "idx_contract_registries_token_id") {
		err = db.Migrator().CreateIndex(&models.ContractRegistry{}, "token_id")
		if err != nil {
			return fmt.Errorf("erro ao criar índice token_id: %w", err)
		}
	}

	return nil
}
