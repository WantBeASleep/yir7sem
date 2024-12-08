from minio import Minio
from minio.error import S3Error
from io import BytesIO


class S3:
    def __init__(self, minio_client: Minio, bucket):
        self.minio_client = minio_client
        self.bucket = bucket
        if not minio_client.bucket_exists(bucket):
            self.minio_client.make_bucket(bucket)

    def load(self, path):
        print("траим с3 вызов")
        try:
            response = self.minio_client.get_object(self.bucket, path)
            data = response.read()
            response.close()
            response.release_conn()  # Освобождаем соединение
            return data
        except S3Error as e:
            print("Ошибка при загрузке файла:", e)

    def store(self, obj, path, content_type: str):
        """Сохранить файл (бинарные) в S3."""
        try:
            # Используем BytesIO для передачи данных
            file_stream = BytesIO(obj)
            self.minio_client.put_object(
                self.bucket, path, file_stream, len(obj), content_type
            )
        except S3Error as e:
            print("Ошибка при загрузке файла:", e)
