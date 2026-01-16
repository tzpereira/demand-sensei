package deps

import (
	"demand-sensei/backend/internal/events/producer"
	"demand-sensei/backend/internal/storage"
)

type Deps struct {
	Producer *producer.Producer
	Storage  storage.Storage
}
