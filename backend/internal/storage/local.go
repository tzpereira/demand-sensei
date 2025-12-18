package storage

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		BasePath: basePath,
	}
}

func (s *LocalStorage) Save(file *multipart.FileHeader) (*UploadResult, error) {
	if err := os.MkdirAll(s.BasePath, 0755); err != nil {
		return nil, err
	}

	filename := uuid.New().String() + "_" + file.Filename
	dstPath := filepath.Join(s.BasePath, filename)

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return nil, err
	}

	return &UploadResult{
		Filename: filename,
		Size:     file.Size,
		Path:     dstPath,
	}, nil
}
