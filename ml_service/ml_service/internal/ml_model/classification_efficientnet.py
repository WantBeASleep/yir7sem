import torch
from torch import nn
from torchvision.models import efficientnet_b6 as efficientnet_b6
from torch.nn import functional as F
from torchvision import transforms as T
from imgaug import augmenters as iaa

# from model_template_class import ModelABC, settings
from ml_service.internal.ml_model.neuro_class import ModelABC

# relative 
settings = settings = {
    'classification': {'all': './ml_service/internal/ml_model/models/all/efficientnet.pth', 'long': './ml_service/internal/ml_model/models/long/efficientnet.pth', 'cross': './ml_service/internal/ml_model/models/cross/efficientnet.pth'}
}


class EfficientNetModel(ModelABC):

    def __init__(self, model_type: str):  # model_type='all'/'long'/'cross'
        super().__init__()
        self.model_type = model_type
        self.device = torch.device('cuda') if torch.cuda.is_available() else torch.device('cpu')
        # self.load(path=settings['classification'][self.model_type])
        self.load(path=settings['classification']['all'])
        self.transform = T.Compose([
            iaa.Sequential([
            iaa.Resize({"height": 224, "width": 224})
            ]).augment_image,
            T.ToTensor(),
            T.Normalize((0.24), (0.12))
        ])

    def load(self, path: str) -> None:
        self._model = efficientnet_b6()
        self._model.features[0][0] = nn.Conv2d(1, 56, kernel_size=(3, 3), stride=(2, 2), padding=(1, 1), bias=False)
        self._model.classifier[1] = nn.Linear(2304, 3) # Текущее количество классов равно 3: 0 - tirads 2-3, 1 - tirads 4, 2 - tirads 5
        self._model.to(self.device)
        self._model.load_state_dict(torch.load(path, map_location=self.device))
        self._model.eval()

    def preprocessing(self, image_array: object) -> object:
        image_tensor = self.transform(image_array)
        image_tensor = torch.unsqueeze(image_tensor, 0)
        image_tensor = image_tensor.to(self.device)
        return image_tensor

    def predict(self, rois: list) -> tuple:
        """
        args:
            rois: 
                len(rois) - кол-во изображений
                len(rois[i]) - количество узлов на i-ой картинке
                [ <-- массив для изображений
                    [ <-- массив для изображения
                        [scaled_pic_of_roi, roi_uniq_idx, segment_binary_masks] <-- увеличенное изображения узла + уникальный индекс + бинарная маска узла
                    ]
                ]

        # update 22.10.24 - теперь в тупол еще засовываем бин маску для каждого узла
        # вместо [binary_mask, roi_uniq_idx]
        # получаем [scaled_pic_of_roi, roi_uniq_idx, segment_binary_masks]

        return:
            individual_probs - лист в вероятностями tirads по сегментам (структура такая же как rois tiff->изображение->узел)
            tracked_nodules_probs - мапа с вероятностями tirads по узлам (uniq tirads для физического узла, id как в rois) (none для одного изображения)
                tirads передаются как numpy массив с итоговыми вероятностями

        original:
            Аргумент rois - [[[roi[0], 1], ..., [roi[n], m]], [...], ..., [...]] - список списков с rois и соответствующих им индексов узлов (для всего tif),
            len(rois) - количество изображений tif, len(rois[i]) - количество rois на i-ой картинке

            Возвращает: (individual_probs, tracked_nodules_probs)
            individual_probs - [[...], [...], ..., [...]] - список списков с вероятностями отнесения каждого сегмента к каждому классу для каждого изображения tif
            len(individual_probs) - количество изображений tif, len(individual_probs[i]) - количество numpy массивов с вероятностями отнесения к каждому классу для каждого roi на i-ой картинке
            tracked_nodules_probs - None для одного изображения, словарь для кинопетли (нескольких изображений) - принадлежность сегмента к formation, 
            у которого ключ - индекс узла, отслеживаемого на кинопетле, значение - numpy массив с итоговыми вероятностями отнесения к классам
        """

        print(f'Device: {self.device}')
        print('Class inference...')

        tracked_nodules_logits = {}
        tracked_nodules_probs = {}
        tracked_nodules_counts = {}

        individual_probs = []

        for r in rois:
            probs_for_image = []
            for nd in r:
                roi_tensor = self.preprocessing(nd[0])
                with torch.no_grad():
                    logits = self._model(roi_tensor)
                    new_probs = F.softmax(logits, dim=1)
                    new_probs = new_probs.cpu().numpy()[0]
                    probs_for_image.append(new_probs)
                    if len(rois) > 1:
                        if nd[1] in tracked_nodules_logits:
                            tracked_nodules_logits[nd[1]] += logits
                            tracked_nodules_counts[nd[1]] += 1
                        else:
                            tracked_nodules_logits[nd[1]] = logits
                            tracked_nodules_counts[nd[1]] = 1
            individual_probs.append(probs_for_image)

        if len(rois) > 1:
            for nodule in tracked_nodules_logits:
                tracked_nodules_logits[nodule] = tracked_nodules_logits[nodule] / tracked_nodules_counts[nodule]
                tracked_nodules_probs[nodule] = F.softmax(tracked_nodules_logits[nodule], dim=1).cpu().numpy()[0]

        if len(rois) == 1:
            print(individual_probs)
        print('Done!')

        return individual_probs, tracked_nodules_probs