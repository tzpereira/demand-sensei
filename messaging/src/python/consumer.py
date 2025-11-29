import threading
from confluent_kafka import Consumer, KafkaException
from typing import Callable, Dict, Any, List, Optional

def new_consumer(brokers: str, group_id: str, topics: List[str], config: Optional[Dict[str, Any]] = None) -> Consumer:
	"""
	Creates and subscribes a Kafka consumer using the provided brokers, group_id, topics, and config.
	Returns a configured Consumer instance.
	"""
	base_config = {
		'bootstrap.servers': brokers,
		'group.id': group_id,
		'auto.offset.reset': 'earliest',
	}
 
	if config:
		base_config.update(config)
  
	consumer = Consumer(base_config)
	consumer.subscribe(topics)
 
	return consumer

def start_consumers(num: int, new_consumer_fn: Callable[[], Consumer], handler: Callable[[Any], None]):
	"""
	Starts multiple Kafka consumers in separate threads.
	Each consumer is created using new_consumer_fn and processes messages using the provided handler.
	Blocks until all consumers have exited.
	"""
	threads = []
	stop_event = threading.Event()

	def run():
		consumer = new_consumer_fn()
		try:
			consume(stop_event, consumer, handler)
		finally:
			consumer.close()

	for _ in range(num):
		t = threading.Thread(target=run)
		t.start()
		threads.append(t)
  
	try:
		for t in threads:
			t.join()
	except KeyboardInterrupt:
		stop_event.set()
		for t in threads:
			t.join()

def consume(stop_event: threading.Event, consumer: Consumer, handler: Callable[[Any], None]):
	"""
	Reads messages from the Kafka consumer and calls the handler for each message.
	Blocks until stop_event is set or a fatal error occurs.
	"""
	while not stop_event.is_set():
		try:
			msg = consumer.poll(timeout=1.0)
   
			if msg is None:
				continue

			if msg.error():
				if msg.error().retriable():
					print(f"Transient error: {msg.error()}")
					continue
				else:
					raise KafkaException(msg.error())
 
			handler(msg)
		except Exception as e:
			print(f"Consumer error: {e}")
			break

def consume_batch(stop_event: threading.Event, consumer: Consumer, handler: Callable[[List[Any]], None], batch_size: int = 100, poll_timeout: float = 0.1):
	"""
	Reads messages from the Kafka consumer in batches and calls the handler with a list of messages.
	Blocks until stop_event is set.
	"""
	batch = []
 
	while not stop_event.is_set():
		try:
			msg = consumer.poll(timeout=poll_timeout)
   
			if msg is None:
				if batch:
					handler(batch)
					batch = []
				continue

			if msg.error():
				print(f"Consumer error: {msg.error()}")
				continue
			batch.append(msg)
   
			if len(batch) >= batch_size:
				handler(batch)
				batch = []
		except Exception as e:
			print(f"Batch consumer error: {e}")
			break
