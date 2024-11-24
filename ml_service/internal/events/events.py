# ПАЙТОН
from confluent_kafka import Consumer, KafkaException
import ml_service.internal.events.kafka_pb2 as pb

class EventsYo():
    def __init__(self, uzi):
        self.uzi = uzi

    def run(self):
        consumer_config = {
            'bootstrap.servers': 'localhost:9092',  # Адрес Kafka-брокера
            'group.id': 'mlService',        # Имя consumer group
            'auto.offset.reset': 'earliest'         # Начинать с самого начала, если оффсет не найден
        }

        consumer = Consumer(consumer_config)
        consumer.subscribe(['uziSplitted'])
        while True:
            msg = consumer.poll(timeout=1.0)
            # continue
            if msg is None:
                continue  # Если сообщения нет, то пропускаем итерацию

            uzi_splitted_event = pb.uziSplitted()
            uzi_splitted_event.ParseFromString(msg.value())

            print("РАЗЪЕБАЛОВО АЙДИ: ", uzi_splitted_event.uzi_id)

            self.uzi.segmentClassificateSave(uzi_splitted_event.uzi_id, uzi_splitted_event.pages_id)
            consumer.commit(msg)
