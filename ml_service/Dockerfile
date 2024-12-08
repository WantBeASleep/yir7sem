FROM pytorch/pytorch:2.5.0-cuda12.4-cudnn9-runtime

RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    ffmpeg libsm6 libxext6 libgl1 \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install 
    
WORKDIR /ml_service

COPY ./requirements.txt /ml_service/

RUN  pip install --no-cache-dir -r requirements.txt

COPY ./ /ml_service/

CMD ["python", "main.py"]
