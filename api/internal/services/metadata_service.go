package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"vfinance-api/internal/models"

	"errors"
	"gorm.io/gorm"
)

type MetadataService struct {
	db *gorm.DB
}

func NewMetadataService(db *gorm.DB) *MetadataService {
	return &MetadataService{db: db}
}

func (s *MetadataService) StoreMetadata(hash string, vehicleData models.VehicleData) error {
	// Debug: verificar se hash já existe
	var existing models.VehicleMetadata
	if err := s.db.First(&existing, "hash = ?", hash).Error; err == nil {
		// Hash já existe - vamos ver o que tem lá
		return fmt.Errorf("hash já existe no banco: %s (dados existentes: %s)", hash, string(existing.VehicleData))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Erro diferente de "não encontrado"
		return fmt.Errorf("erro ao verificar hash existente: %w", err)
	}

	// Hash não existe, podemos inserir
	// Converter para JSON
	jsonData, err := json.Marshal(vehicleData)
	if err != nil {
		return err
	}

	// Criar registro (hash já foi validado pelo contrato)
	metadata := models.VehicleMetadata{
		Hash:        hash,
		VehicleData: jsonData,
	}

	return s.db.Create(&metadata).Error
}

func (s *MetadataService) GetMetadata(hash string) (*models.VehicleData, error) {
	var metadata models.VehicleMetadata
	if err := s.db.First(&metadata, "hash = ?", hash).Error; err != nil {
		return nil, err
	}

	var vehicleData models.VehicleData
	if err := json.Unmarshal(metadata.VehicleData, &vehicleData); err != nil {
		return nil, err
	}

	return &vehicleData, nil
}

func (s *MetadataService) UpdateMetadata(hash string, vehicleData models.VehicleData) error {
	jsonData, err := json.Marshal(vehicleData)
	if err != nil {
		return err
	}

	return s.db.Model(&models.VehicleMetadata{}).
		Where("hash = ?", hash).
		Update("vehicle_data", jsonData).Error
}

func (s *MetadataService) DeleteMetadata(hash string) error {
	return s.db.Delete(&models.VehicleMetadata{}, "hash = ?", hash).Error
}

func (s *MetadataService) GenerateHash(vehicleData models.VehicleData) (string, error) {
	jsonData, err := json.Marshal(vehicleData)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:]), nil
}
