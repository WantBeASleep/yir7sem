FROM pytorch/pytorch:2.5.0-cuda12.4-cudnn9-runtime

# Устанавливаем зависимости для сборки Python-пакетов
RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    ffmpeg libsm6 libxext6 libgl1 \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install 
    
# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл requirements.txt для установки зависимостей
COPY ./requirements.txt /app/

# Устанавливаем зависимости с помощью pip
RUN  pip install --no-cache-dir --upgrade -r requirements.txt
# Копируем весь проект
COPY ./ /app/
RUN python -c "import sys; print(sys.path)"     
# Команда для запуска приложения
CMD ["python", "main.py"]
