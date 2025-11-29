# Python Messaging SDK

This project is a Python SDK for Kafka messaging, including:
- Asynchronous Producer abstraction
- Multi-threaded Consumer utilities
- Batch and single-message consumption

## Requirements
- Python 3.7+
- [confluent-kafka](https://pypi.org/project/confluent-kafka/)

Install dependencies:
```sh
pip install -r requirements.txt
```

## Usage

### Producer
```python
from producer import KafkaProducer

producer = KafkaProducer(brokers="kafka1:9092,kafka2:9092")

def delivery_report(err, msg):
    if err:
        print(f"Delivery failed: {err}")
    else:
        print(f"Delivered to {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")

producer.start_delivery_handler(delivery_report)
producer.produce(topic="my-topic", value=b"payload", key=b"user-1")
producer.close()
```

### Consumer
```python
from consumer import new_consumer, start_consumers

def handler(msg):
    print(f"Received: {msg.value()} from {msg.topic()} [{msg.partition()}] offset {msg.offset()}")

consumer_fn = lambda: new_consumer(
    brokers="kafka1:9092,kafka2:9092",
    group_id="my-group",
    topics=["my-topic"]
)

start_consumers(num=1, new_consumer_fn=consumer_fn, handler=handler)
```

## Notes
- No brokers or topics are hardcoded; provide them via parameters.
- See the code for more advanced usage (batch consumption, custom configs, etc).
