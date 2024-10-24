from minio import Minio
from minio.error import S3Error
import time

from concurrent import futures
import grpc
import sys
sys.path.append("/home/wantbeasleep/yir/ml_service")
sys.path.append("/home/wantbeasleep/yir/ml_service/api")
import ml_service.api.ml_api_pb2 as pb
import ml_service.api.ml_api_pb2_grpc as pb_grpc


sys.path.append("/home/wantbeasleep/yir/ml_service/internal/s3")
import ml_service.internal.s3.s3 as mys3
sys.path.append("/home/wantbeasleep/yir/ml_service/internal/ml_model")
import ml_service.internal.ml_model.segmentation as seg
import ml_service.internal.ml_model.classification_efficientnet as cla
sys.path.append("/home/wantbeasleep/yir/ml_service/internal/usecases/uzi")
sys.path.append("/home/wantbeasleep/yir/ml_service/internal/utils")
import ml_service.internal.usecases.uzi.uzi as usecaseuzi
import ml_service.internal.api.ml_controller as ctrl

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

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
pb_grpc.add_MLAPIServicer_to_server(controller, server)
server.add_insecure_port('[::]:50055')
server.start()  # Запускаем сервер
print("Server started on port 50055")

try:
    while True:
            time.sleep(86400)  # Сервер работает бесконечно
except KeyboardInterrupt:
    server.stop(0)  # Остановка сервера при прерывании