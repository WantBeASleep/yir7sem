from concurrent import futures

import grpc
from minio import Minio

import ml_service.api.ml_api_pb2_grpc as pb_grpc
import ml_service.internal.api.ml_controller as ctrl
import ml_service.internal.events.events as kafkaevents
import ml_service.internal.ml_model.classification_efficientnet as cla
import ml_service.internal.ml_model.segmentation as seg
import ml_service.internal.s3.s3 as mys3
import ml_service.internal.usecases.uzi.uzi as usecaseuzi
from ml_service.config.default import get_settings


def run_server():
    settings = get_settings()

    minio_client = Minio(
        endpoint=settings.s3_endpoint,
        access_key=settings.s3_access_key,
        secret_key=settings.s3_secret_key.get_secret_value(),
        secure=False,  # Установите True, если используете HTTPS
    )

    bucket = settings.s3_bucket_name

    s3 = mys3.S3(minio_client, bucket)

    segmdl = seg.SegmentationModel(model_type=settings.segmentation_model_type)
    claml = cla.EfficientNetModel(model_type=settings.classification_model_type)

    usecase = usecaseuzi.uziUseCase(segmdl, claml, s3)

    controller = ctrl.MlController(usecase)
    kafka = kafkaevents.EventsYo(usecase)
    print("Kafka started...")
    kafka.run()

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb_grpc.add_MLAPIServicer_to_server(controller, server)

    server_port = str(settings.grpc_port)

    server.add_insecure_port("[::]:" + server_port)
    server.start()  # Запускаем сервер
    print(f"Server started on port {server_port}")
    server.wait_for_termination()


if __name__ == "__main__":
    run_server()
