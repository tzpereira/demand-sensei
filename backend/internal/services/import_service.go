package services

import (
	"mime/multipart"

	"demand-sensei/backend/internal/storage"
)

type ImportService struct {
	storage storage.Storage
}

func NewImportService(st storage.Storage) *ImportService {
	return &ImportService{
		storage: st,
	}
}

func (s *ImportService) Import(file *multipart.FileHeader) (*storage.UploadResult, error) {
	return s.storage.Save(file)
}
