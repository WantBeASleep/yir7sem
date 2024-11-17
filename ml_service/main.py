from minio import Minio
from minio.error import S3Error
import time

from concurrent import futures
import grpc
import sys

import ml_service.api.ml_api_pb2 as pb
import ml_service.api.ml_api_pb2_grpc as pb_grpc



import ml_service.internal.events.events as kafkaevents



import ml_service.internal.s3.s3 as mys3

import ml_service.internal.ml_model.segmentation as seg
import ml_service.internal.ml_model.classification_efficientnet as cla

import ml_service.internal.usecases.uzi.uzi as usecaseuzi
import ml_service.internal.api.ml_controller as ctrl

def run_server():

    minio_client = Minio(
        "localhost:9000",  # Например, "localhost:9000" или "minio.example.com:9000"
        access_key="z1qTDuqQH233Di9sWyiV",
        secret_key="OXgQ5UKoVvejYCEEpUb3xBZ8rvW4dwHNNO7O2b2L",
        secure=False  # Установите True, если используете HTTPS
    )

    bucket = "uzi"

    s3 = mys3.S3(minio_client, bucket)

    segmdl = seg.SegmentationModel('cross')
    claml = cla.EfficientNetModel('cross')

    usecase = usecaseuzi.uziUseCase(segmdl, claml, s3)

    controller = ctrl.MlController(usecase)
    kafka = kafkaevents.EventsYo(usecase)
    print("РАЗЪЕБНАЯ КАФКА ПОЕХАЛА")
    kafka.run()

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb_grpc.add_MLAPIServicer_to_server(controller, server)
    server.add_insecure_port('[::]:50057')
    server.start()  # Запускаем сервер
    print("Server started on port 50057")
    server.wait_for_termination()

    # try:
    #     while True:
    #             time.sleep(86400)  # Сервер работает бесконечно
    # except KeyboardInterrupt:
    #     server.stop(0)  # Остановка сервера при прерывании


if __name__ == "__main__":
    run_server()