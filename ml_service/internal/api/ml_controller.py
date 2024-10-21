import ml_service.api.ml_api_pb2 as pb
import ml_service.api.ml_api_pb2_grpc as pb_grpc
import ml_service.internal.usecases.uzi.uzi as uziusecase

class MlController(pb_grpc.MLAPIServicer):
    def __init__(self, uzi_usecase: uziusecase.uziUseCase):
        super().__init__()
        self.uzi_usecase = uzi_usecase
        


    def SegmentAndClassification(self, request, context):
        uzi_id = request.uzi_id

        try:
            self.uzi_usecase.segmentAndClassificateByID(uzi_id)
            return
        except Exception as e:
            context.set_details(f'Error processing request: {str(e)}')
            context.set_code(pb_grpc.StatusCode.INTERNAL)  # устанавливаем код ошибки
            return None  # Возвращаем None в случае ошибки, чтобы завершить вызов с ошибкой


        