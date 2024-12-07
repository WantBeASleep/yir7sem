from abc import ABC, abstractmethod
import numpy as np

# Веса модели
# relative 
settings = {
    'classification': {'all': './ml_service/internal/ml_model/models/all/efficientnet.pth', 'cross': './ml_service/internal/ml_model/models/cross/efficientnet.pth', 'long': './ml_service/internal/ml_model/models/long/efficientnet.pth'},
    'segmentation': {'all': './ml_service/internal/ml_model/models/all/deeplabv3plus.pkl', 'cross': './ml_service/internal/ml_model/models/cross/deeplabv3plus.pkl', 'long': './ml_service/internal/ml_model/models/long/deeplabv3plus.pkl'},
    'tracking': './models/tracking_model.pkl',
}


class ModelABC(ABC):

    def __init__(self):
        self._model = None

    @abstractmethod
    def load(self, path: str) -> None:
        """
        Функция, в которой обределяется структура NN и
        происходит загрузка весов модели в self._model

        params:
          path - путь к файлу, в котором содержатся веса модели
        """
        ...

    @abstractmethod
    def preprocessing(self, image: np.ndarray) -> object:
        """
        Функция, котороя предобрабатывает изображение к виду, 
        с которым можеn взаимодействовать модель из self._model

        params:
          image - numpy_array изображения

        return - возвращает предобработанное изображение 
        """
        ...

    @abstractmethod
    def predict(self, group: list) -> object:
        """
        Функция, в которой предобработанное изображение подается
        на входы NN (self._model) и возвращается результат работы NN 

        params:
          group - группа сников которые обрабатываются, это list из np.ndarray

        return - результаты предсказания
        """
        ...