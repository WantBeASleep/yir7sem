import numpy as np
import cv2
from confluent_kafka import Producer
import json
import imageio
import matplotlib.pyplot as plt
from ml_service.internal.ml_model.neuro_class import ModelABC
from ml_service.internal.s3.s3 import S3
from ml_service.internal.utils import image_parser
import ml_service.internal.events.kafka_pb2 as pb_event

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

        print_lengths_return_ndarray_list(rois)
        print("\n\n\n\n")
        print_lengths_return_ndarray_list(indv)
        print("\n\n\n\n")
        print_lengths_return_ndarray_list(tracked)

    def segmentClassificateSave(self, uzi_id, pages_id):
        # фул похуй поехали
        print("в с3 кабаном")
        data = self.store.load(uzi_id + "/" + uzi_id)

        masks, rois = self.segmentUzi(data)
        indv, tracked = self.classificateUzi(rois)

        formations = dict()
        for k in tracked:
            tirads = pb_event.Tirads(
                tirads_1=0,
                tirads_2=0,
                tirads_3=tracked[k][0],
                tirads_4=tracked[k][0],
                tirads_5=tracked[k][0]
            )
            formations[k] = pb_event.MlFormation(
                tirads=tirads
            )


        for i in range(len(rois)):
            print("РАЗЪЕБНАЯ И:", i)
            for j in range(len(rois[i])):
                # бинаризуем
                mask = rois[i][j][2]
                mask = mask.astype(np.uint8)
                mask = (mask * 255).astype(np.uint8)
                print(mask.shape, type(mask))
                print(np.unique(mask))

                contours, _ = cv2.findContours(mask, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
                print("КОЛИЧЕСТВО КОНТУТУРОВ: ", len(contours))
                contour = contours[0].squeeze()
                contour_points = [{"x": int(point[0]), "y": int(point[1])} for point in contour]
                contourJSSSSSSSSS = json.dumps(contour_points)

                tirads = pb_event.Tirads(
                    tirads_1=0,
                    tirads_2=0,
                    tirads_3=indv[i][j][0],
                    tirads_4=indv[i][j][1],
                    tirads_5=indv[i][j][2]
                )
                ml_segment = pb_event.MlSegment(
                    image_id=pages_id[i],
                    contor=contourJSSSSSSSSS,
                    tirads = tirads
                )
                
                formations[rois[i][j][1]].segments.append(ml_segment)

        msg_event = pb_event.uziProcessed(
            formations=list(formations.values())
        )

        content = msg_event.SerializeToString()

        producer_config = {
            'bootstrap.servers': 'localhost:9092'  # Адрес вашего Kafka брокера
        }
        producer = Producer(producer_config)

        producer.produce('uziProcessed', content)
        producer.flush()

        

        




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
