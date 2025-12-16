package services

import (
	"mime/multipart"

	"demand-sensei/backend/internal/storage"
)

type ImportService struct {
	storage *storage.LocalStorage
}

func NewImportService(storage *storage.LocalStorage) *ImportService {
	return &ImportService{storage: storage}
}

func (s *ImportService) Import(file *multipart.FileHeader) (*storage.UploadResult, error) {
	return s.storage.Save(file)
}
