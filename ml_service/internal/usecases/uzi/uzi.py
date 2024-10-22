import numpy as np
import imageio
import matplotlib.pyplot as plt
from ml_service.internal.ml_model.neuro_class import ModelABC
from ml_service.internal.s3.s3 import S3
from ml_service.internal.utils import image_parser

res_for_recursive = []

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
        global res_for_recursive
        print("в с3 кабаном")
        data = self.store.load(uzi_id)

        size_in_bytes = len(data)
        size_in_kb = size_in_bytes/1024
        size_in_mb = size_in_kb/1024    
        print(f"Размер загруженных данных: {size_in_bytes} байт, {size_in_kb:.2f} Кб, {size_in_mb:.2f} Mb")
        
        masks, rois = self.segmentUzi(data)
        indv, tracked = self.classificateUzi(rois)

        # print(f"размер indv {len(indv)}")
        # print(indv)

        print(f"размер tracked {len(tracked)}")
        print(tracked)


        # дальше код - насрать в S3

def print_lengths_return_ndarray_list(data, level=0):
    """Рекурсивная функция для вывода длин всех вложенных массивов/list."""
    # Проверяем, что переданные данные являются списком
    if isinstance(data, list):
        print(f"{' ' * level}Length of list: {len(data)}")
        for item in data:
            print_lengths_return_ndarray_list(item, level + 1)  # Рекурсивный вызов для каждого элемента
    else:
        if type(data) is np.ndarray:
            res_for_recursive.append(data)
            print(f"{' ' * level}Item: NUMPY ARRAY: размерность(shape): {data.shape}")
            # print(f"{' ' * (level + 1)}Размерность (shape): {data.shape}")  # Вывод: (2, 3), т.е. 2 строки и 3 столбца
            # print(f"{' ' * (level + 1)}Общее количество элементов (size): {data.size}")  # Вывод: 6
            # print(f"{' ' * (level + 1)}Количество измерений (ndim): {data.ndim}")  # Вывод: 2
        else:
            print(f"{' ' * level}Item: {data} (type: {type(data)})")
    
def save_result(arr: list, name):
        path = "/home/wantbeasleep/yir/trashlog/" + name + ".tiff"
        if len(arr) > 1:
            imageio.mimwrite(path, arr)
        elif len(arr) == 1:
            plt.imsave(path, np.array(arr[0]))
        print('Result saved')
