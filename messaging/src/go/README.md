# Go Messaging SDK

This project is a Go SDK for Kafka messaging, including:
- Asynchronous Producer abstraction
- Multi-goroutine Consumer utilities
- Batch and single-message consumption

## Requirements
- Go 1.18+
- [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)

Install dependencies:
```sh
go get github.com/confluentinc/confluent-kafka-go/kafka
```

## Usage

### Producer
```go
import (
    "context"
    "github.com/tzpereira/demand-sensei/messaging/src/go/messaging"
    ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

producer, err := messaging.NewProducer([]string{"kafka1:9092", "kafka2:9092"}, nil)
if err != nil {
    // handle error
}

ctx, cancel := context.WithCancel(context.Background())
messaging.StartDeliveryHandler(ctx, producer, func(msg *ckafka.Message) {
    if msg.TopicPartition.Error != nil {
        // handle delivery error
    } else {
        // handle successful delivery
    }
})

data := []byte("payload")
err = messaging.Produce(context.Background(), producer, "my-topic", data)
if err != nil {
    // handle produce error
}

// On shutdown:
cancel()
producer.Close()
```

### Consumer
```go
import (
    "context"
    "github.com/tzpereira/demand-sensei/messaging/src/go/messaging"
)

consumer, err := messaging.NewConsumer(
    []string{"kafka1:9092", "kafka2:9092"},
    "my-group",
    []string{"my-topic"},
    nil,
)
if err != nil {
    // handle error
}

ctx, cancel := context.WithCancel(context.Background())
err = messaging.Consume(ctx, consumer, func(msg *ckafka.Message) {
    // process message
})
if err != nil {
    // handle consume error
}

// On shutdown:
cancel()
consumer.Close()
```

## Notes
- No brokers or topics are hardcoded; provide them via parameters.
- See the code for more advanced usage (batch consumption, custom configs, etc).
