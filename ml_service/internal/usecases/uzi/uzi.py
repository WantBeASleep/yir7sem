import uuid
from ml_service.internal.ml_model.neuro_class import ModelABC
from ml_service.internal.s3.s3 import S3
from ml_service.internal.utils import image_parser

class uziUseCase():
    def __init__(self, segmentModel: ModelABC, efficientModel: ModelABC, store: S3):
        self.segmentationModel = segmentModel
        self.efficientModel = efficientModel
        self.store = store

    def segmentUzi(self, data):
        parsed = image_parser.read_image(data)
        
        return self.segmentationModel.predict(parsed)

    def classificateUzi(self, rois):
        indv, tracked = self.efficientModel.predict(rois)
        return indv, tracked
    
    def segmentAndClassificateByID(self, uzi_id):
        print("в с3 кабаном")
        data = self.store.load(uzi_id)

        size_in_bytes = len(data)
        size_in_kb = size_in_bytes/1024
        size_in_mb = size_in_kb/1024    
        print(f"Размер загруженных данных: {size_in_bytes} байт, {size_in_kb:.2f} Кб, {size_in_mb:.2f} Mb")
        
        masks, rois = self.segmentUzi(data)
        indv, tracked = self.classificateUzi(rois)

        print(rois)
        print(indv)
        print(tracked)

        # дальше код - насрать в S3

    




