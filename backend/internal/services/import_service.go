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

func (s *ImportService) Import(
	file *multipart.FileHeader,
	importType string,
) (*storage.UploadResult, error) {

	result, err := s.storage.Save(file)
	if err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"import_type": importType,
		"filename":    result.Filename,
		"path":        result.Path,
		"size":        result.Size,
	})

	topic := "import." + importType + ".created"

	_ = s.producer.Produce(
		context.Background(),
		topic,
		[]byte(result.Filename),
		payload,
	)

	return result, nil
}

