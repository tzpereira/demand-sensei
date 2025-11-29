import threading
from confluent_kafka import Producer
from typing import Callable, Dict, Any, Optional

class KafkaProducer:
    """
    KafkaProducer provides a reusable, production-ready abstraction for producing messages to Kafka.
    - Brokers and configuration must be provided by the application.
    - Does not hardcode brokers or topics.
    - Handles asynchronous delivery reports in a background thread.
    """
    def __init__(self, brokers: str, config: Optional[Dict[str, Any]] = None):
        base_config = {
            'bootstrap.servers': brokers,
            'acks': 'all',
        }
        if config:
            base_config.update(config)
        self.producer = Producer(base_config)
        self._delivery_thread = None
        self._stop_event = threading.Event()
        self.report_func = None

    def start_delivery_handler(self, report_func: Callable):
        """
        Starts a background thread to handle delivery reports.
        Should be called once after creating the producer.
        The report_func will be called for each delivery report.
        """
        self.report_func = report_func
        def delivery_loop():
            while not self._stop_event.is_set():
                self.producer.poll(0.1)
        self._delivery_thread = threading.Thread(target=delivery_loop, daemon=True)
        self._delivery_thread.start()

    def produce(self, topic: str, value: bytes, key: Optional[bytes] = None):
        """
        Sends a message asynchronously to the specified Kafka topic.
        Returns immediately unless the message could not be queued for delivery.
        """
        self.producer.produce(
            topic=topic,
            value=value,
            key=key,
            callback=self.report_func
        )

    def close(self):
        """
        Stops the delivery handler and flushes all pending messages.
        Should be called before shutting down the application to ensure all messages are delivered.
        """
        self._stop_event.set()
        if self._delivery_thread:
            self._delivery_thread.join()
        self.producer.flush()