package storage

import (
	"mime/multipart"
	"os"
	"path/filepath"
)

const UploadDir = "/app/data/uploads"

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		BasePath: UploadDir,
	}
}

func (s *LocalStorage) Save(file *multipart.FileHeader) (*UploadResult, error) {
	if err := os.MkdirAll(s.BasePath, 0755); err != nil {
		return nil, err
	}

	dstPath := filepath.Join(s.BasePath, file.Filename)

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
		Filename: file.Filename,
		Size:     file.Size,
		Path:     dstPath,
	}, nil
}
