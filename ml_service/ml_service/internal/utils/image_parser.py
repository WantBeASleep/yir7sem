from PIL import Image
import numpy as np
from io import BytesIO


def read_image(data) -> list:
    images = []
    image = Image.open(BytesIO(data))
    i = 0
    while True:
        try:
            image.seek(i)
            image_array = np.array(image)
            images.append(image_array)
            i += 1
        except EOFError:
            break
    return images
