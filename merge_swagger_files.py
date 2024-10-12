import yaml
import argparse
import os

# Функция для чтения YAML файла
def load_yaml(file_path):
    try:
        with open(file_path, 'r') as f:
            return yaml.safe_load(f)
    except Exception as e:
        print(f"Ошибка при чтении файла {file_path}: {e}")
        return None

# Функция для объединения двух Swagger файлов
def merge_swagger(file1_path, file2_path, output_dir):
    swagger1 = load_yaml(file1_path)
    swagger2 = load_yaml(file2_path)

    if not swagger1 or not swagger2:
        print("Ошибка: не удалось загрузить один или оба файла.")
        return

    # Объединить пути
    swagger1['paths'].update(swagger2.get('paths', {}))

    # Объединить компоненты (если есть)
    if 'components' in swagger1 and 'components' in swagger2:
        swagger1['components'].update(swagger2['components'])

    # Определить имя и путь для сохранения файла
    output_path = os.path.join(output_dir, "merged_swagger.yaml")
    
    # Сохранить результат
    try:
        with open(output_path, 'w') as fout:
            yaml.dump(swagger1, fout, default_flow_style=False, allow_unicode=True)
        print(f"Файлы успешно объединены и сохранены в {output_path}")
    except Exception as e:
        print(f"Ошибка при сохранении файла {output_path}: {e}")

# Основная функция с использованием аргументов командной строки
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Объединение двух Swagger файлов.")
    parser.add_argument("file1", help="Путь к первому Swagger файлу")
    parser.add_argument("file2", help="Путь ко второму Swagger файлу")
    parser.add_argument("output_dir", help="Директория для сохранения объединенного файла")

    args = parser.parse_args()

    # Выполнить объединение
    merge_swagger(args.file1, args.file2, args.output_dir)
