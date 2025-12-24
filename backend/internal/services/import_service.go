package services

import (
	"context"
	"encoding/json"
	"mime/multipart"

	"demand-sensei/backend/internal/events/producer"
	"demand-sensei/backend/internal/storage"
)

type ImportService struct {
	storage  storage.Storage
	producer *producer.Producer
}

func NewImportService(
	st storage.Storage,
	prod *producer.Producer,
) *ImportService {
	return &ImportService{
		storage:  st,
		producer: prod,
	}
}

func (s *ImportService) Import(file *multipart.FileHeader) (*storage.UploadResult, error) {
	result, err := s.storage.Save(file)
	if err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(result)

	_ = s.producer.Produce(
		context.Background(),
		"imports.created",
		[]byte(result.Filename),
		payload,
	)

	return result, nil
}
