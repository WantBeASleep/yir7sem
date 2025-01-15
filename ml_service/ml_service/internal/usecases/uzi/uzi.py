import numpy as np
import cv2
from confluent_kafka import Producer
import json
import imageio
import matplotlib.pyplot as plt
from ml_service.internal.ml_model.neuro_class import ModelABC
from ml_service.internal.s3.s3 import S3
from ml_service.internal.utils import image_parser
from ml_service.config.default import get_settings
import ml_service.internal.events.kafka_pb2 as pb_event
import uuid

res_for_recursive = []
settings = get_settings()


class uziUseCase:
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

    def segmentClassificateSave(self, uzi_id, pages_id):
        print("Going to S3...")
        print(pages_id)
        data = self.store.load(uzi_id + "/" + uzi_id)

        masks, rois = self.segmentUzi(data)
        indv, tracked = self.classificateUzi(rois)
        print(type(tracked))
        # indv - probs по segments
        # tracked - probs по formations
        # tirads=probs

        nodes = dict()
        # k - это formation_id из модели
        formation_ids = {}
        for k in tracked:
            print(k)

            formation_uuid = str(uuid.uuid4())

            nodes[k] = pb_event.UziProcessed.Node(
                id=formation_uuid,
                uzi_id=uzi_id,
                tirads_23=tracked[k][0],
                tirads_4=tracked[k][1],
                tirads_5=tracked[k][2],
            )
            formation_ids[k] = formation_uuid

        # Это мы запихнули в словарик dct[formation] = probs
        # print_lengths_return_ndarray_list(tracked)

        segment_ids = []

        # Далее бежим по всем картинкам и сегментам с целью отдать

        segments = []

        for i in range(len(rois)):
            print("ROI num: ", i)
            for j in range(len(rois[i])):
                formation_id_from_model = rois[i][j][1]  # Это то же самое, что k
                print("FORMATION ID FROM MODEL: ", formation_id_from_model)
                if formation_id_from_model not in formation_ids:
                    formation_ids[formation_id_from_model] = str(uuid.uuid4())
                formation_id = formation_ids.get(formation_id_from_model)
                # бинаризуем
                mask = rois[i][j][2]
                mask = mask.astype(np.uint8)
                mask = (mask * 255).astype(np.uint8)
                print(mask.shape, type(mask))
                print(np.unique(mask))

                contours, _ = cv2.findContours(
                    mask, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE
                )
                print("КОЛИЧЕСТВО КОНТУТУРОВ: ", len(contours))
                contour = contours[0].squeeze()

                contour_points = [
                    {"x": int(point[0]), "y": int(point[1])} for point in contour
                ]
                contour = json.dumps(contour_points)

                segment_id = str(uuid.uuid4())
                segment_ids.append(segment_id)

                segment = pb_event.UziProcessed.Segment(
                    id=segment_id,
                    node_id=formation_id,
                    image_id=pages_id[i],
                    contor=contour,
                    tirads_23=indv[i][j][0],
                    tirads_4=indv[i][j][1],
                    tirads_5=indv[i][j][2],
                )

                segments.append(segment)

        msg_event = pb_event.UziProcessed(nodes=list(nodes.values()), segments=segments)

        content = msg_event.SerializeToString()

        producer_config = {"bootstrap.servers": settings.kafka_host + ":" + str(settings.kafka_port)}
        producer = Producer(producer_config)

        producer.produce("uziprocessed", content)
        producer.flush()


def print_lengths_return_ndarray_list(data, level=0):
    """Рекурсивная функция для вывода длин всех вложенных массивов/list."""
    # Проверяем, что переданные данные являются списком
    if isinstance(data, list):
        print(f"{' ' * level}Length of list: {len(data)}")
        for item in data:
            print_lengths_return_ndarray_list(
                item, level + 1
            )  # Рекурсивный вызов для каждого элемента
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
    # relative
    path = "../../../../trashlog/" + name + ".tiff"
    if len(arr) > 1:
        imageio.mimwrite(path, arr)
    elif len(arr) == 1:
        plt.imsave(path, np.array(arr[0]))
    print("Result saved")
