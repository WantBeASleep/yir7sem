from minio import Minio

class S3():
# пока так, бакеты прям сюда вставляем, как найдем решение лучше - перепишем
# MVP
    def __init__(self, minio_client, bucket, access_key, secret_key):
        self.minio_client = minio_client
        self.bucket = bucket
        self.access_key = access_key
        self.secret_key = secret_key

    
