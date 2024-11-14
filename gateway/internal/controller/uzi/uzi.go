package uzi

import (
	"context"
	"net/http"
	"strconv"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity"
	"yir/gateway/internal/entity/uzimodel"
	"yir/gateway/internal/entity/uzimodel/uzidto"
	"yir/gateway/repository"

	validator "github.com/go-playground/validator"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

type UziService interface {
	CreateUzi(ctx context.Context, in *uzimodel.Uzi) (string, error)
	UpdateUzi(ctx context.Context, in *uzimodel.Uzi) error
	CreateFormationWithSegments(ctx context.Context, uziID string, formations *uzidto.FormationWithSegments) error
	UpdateFormation(ctx context.Context, formation *uzidto.Formation) error
	GetUziByID(ctx context.Context, uziID string) (*uzimodel.Uzi, error)
	GetImageWithFormationsSegments(ctx context.Context, imageID string) (*uzidto.ImageWithSegmentsFormations, error)
	GetFormationWithSegments(ctx context.Context, formationID string) (*uzidto.FormationWithSegments, error)
	GetDeviceList(ctx context.Context) ([]uzimodel.Device, error)
}

type UziController struct {
	Service UziService
	S3      *repository.S3Repo
	Kafka   *repository.Producer
}

// PostUzi добавляет метаданные УЗИ и изображение в S3, отправляет uziID в Kafka.
// @Summary Добавить УЗИ
// @Description Создает запись УЗИ, загружает изображение в S3 и отправляет uziID в Kafka.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param image formData file true "Изображение УЗИ"
// @Param uzi body uzimodel.Uzi true "Метаданные УЗИ"
// @Success 200 {string} string "Успешное выполнение"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/add [post]
func (c *UziController) PostUzi(w http.ResponseWriter, r *http.Request) {
	req := &uzimodel.Uzi{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	uziID, err := c.Service.CreateUzi(ctx, req)
	if err != nil {
		http.Error(w, "Failed to get uzi info. Please try again later.", http.StatusBadGateway)
		return
	}
	// Получаем файл из поля "image"
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Проверяем формат изображения
	format := handler.Header.Get("Content-Type")
	if format != "image/png" && format != "image/tiff" {
		http.Error(w, "Unsupported image format. Only PNG and TIFF are allowed.", http.StatusBadRequest)
		return
	}

	// загружаем в S3
	metadata := &entity.FileMeta{
		ContentType: format,
		Path:        req.URL,
	}

	if err := c.S3.Upload(ctx, metadata, file); err != nil {
		custom.Logger.Error(
			"failed to send image to S3",
			zap.Error(err),
		)
		http.Error(w, "Failed to send image to S3. Please try again later.", http.StatusBadGateway)
		return
	}
	// отправляем в кафку uzi_id
	if err := c.Kafka.Send("1", uziID); err != nil {
		custom.Logger.Error(
			"failed to send uzi_id to kafka",
			zap.Error(err),
		)
		http.Error(w, "Failed to send uzi_id to Kafka. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateUzi обновляет информацию УЗИ.
// @Summary Обновить УЗИ
// @Description Обновляет существующую запись УЗИ по переданным данным.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param uzi body uzimodel.Uzi true "Метаданные УЗИ"
// @Success 200 {string} string "Успешное выполнение"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/{uzi_id} [put]
func (c *UziController) UpdateUzi(w http.ResponseWriter, r *http.Request) {
	req := &uzimodel.Uzi{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	if err := c.Service.UpdateUzi(ctx, req); err != nil {
		http.Error(w, "Failed to get uzi info. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// PostFormationWithSegments добавляет формацию с сегментами для УЗИ.
// @Summary Добавить формацию с сегментами
// @Description Создает формацию с сегментами, привязанную к указанному uziID.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param uzi_id path string true "ID УЗИ"
// @Param formation body uzidto.FormationWithSegments true "Формация с сегментами"
// @Success 200 {string} string "Успешное выполнение"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /formation/segments/{uzi_id} [post]
func (c *UziController) PostFormationWithSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uziID := vars["uzi_id"]
	// проверка что передали норм idшник вообще
	_, err := strconv.ParseUint(uziID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	req := &uzidto.FormationWithSegments{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	if err := c.Service.CreateFormationWithSegments(ctx, uziID, req); err != nil {
		http.Error(w, "Failed to insert formation with segments. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateFormation обновляет информацию о формации.
// @Summary Обновить формацию
// @Description Обновляет существующую формацию по переданному ID.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param formation_id path string true "ID формации"
// @Param formation body uzidto.Formation true "Данные формации"
// @Success 200 {string} string "Успешное выполнение"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/formation/{formation_id} [put]
func (c *UziController) UpdateFormation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// проверка что передали норм idшник вообще
	_, err := strconv.ParseUint(vars["formation_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	req := &uzidto.Formation{}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	if err := c.Service.UpdateFormation(ctx, req); err != nil {
		http.Error(w, "Failed to update formation. Please try again later.", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUziByID возвращает метаданные УЗИ по ID.
// @Summary Получить УЗИ по ID
// @Description Возвращает информацию о УЗИ по указанному ID.
// @Tags Uzi
// @Produce json
// @Param uzi_id path string true "ID УЗИ"
// @Success 200 {object} uzimodel.Uzi "Данные УЗИ"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/{uzi_id} [get]
func (c *UziController) GetUziByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := vars["uzi_id"]
	// проверка что передали норм idшник вообще
	_, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	data, err := c.Service.GetUziByID(ctx, req)
	if err != nil {
		http.Error(w, "Failed to get uzi by id. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUziInfo возвращает информацию об УЗИ без изображения.
// @Summary Получить информацию УЗИ
// @Description Возвращает данные УЗИ по ID, без изображения.
// @Tags Uzi
// @Produce json
// @Param uzi_id path string true "ID УЗИ"
// @Success 200 {object} uzidto.Report "Информация об УЗИ"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/info/{uzi_id} [get]
func (c *UziController) GetUziInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := vars["uzi_id"]
	// проверка что передали норм idшник вообще
	_, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	data, err := c.Service.GetUziByID(ctx, req)
	if err != nil {
		http.Error(w, "Failed to get uzi info. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetFormationWithSegments возвращает формацию с сегментами по ID.
// @Summary Получить формацию с сегментами
// @Description Возвращает формацию с сегментами по указанному ID формации.
// @Tags Uzi
// @Produce json
// @Param formation_id path string true "ID формации"
// @Success 200 {object} uzidto.FormationWithSegments "Данные формации с сегментами"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /formation/segments/{formation_id} [get]
func (c *UziController) GetFormationWithSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := vars["formation_id"]
	// проверка что передали норм idшник вообще
	_, err := strconv.ParseUint(req, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payload request", http.StatusBadRequest)
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	ctx := context.Background()
	data, err := c.Service.GetFormationWithSegments(ctx, req)
	if err != nil {
		http.Error(w, "Failed to get formation with segments. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetUziImageWithSegments возвращает изображение УЗИ с сегментами по ID.
// @Summary Получить изображение УЗИ с сегментами
// @Description Возвращает изображение УЗИ и информацию о сегментах по указанному ID изображения.
// @Tags Uzi
// @Produce image/png, image/tiff
// @Param image_id path string true "ID изображения"
// @Success 200 {string} string "Изображение с сегментами"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/image/segments/{image_id} [get]
func (c *UziController) GetUziImageWithSegments(w http.ResponseWriter, r *http.Request) {
	// Извлекаем image_id из URL
	vars := mux.Vars(r)
	imageID := vars["image_id"]
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	// Проверяем, что ID является корректным числом
	if _, err := strconv.ParseUint(imageID, 10, 64); err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Получаем изображение и метаинформацию из Uzi
	data, err := c.Service.GetImageWithFormationsSegments(ctx, imageID)
	if err != nil {
		http.Error(w, "Failed to retrieve uzi-image data", http.StatusInternalServerError)
		return
	}
	image, err := c.S3.Get(ctx, data.Image.URL)
	if err != nil {
		http.Error(w, "Failed to get image from S3", http.StatusInternalServerError)
		return
	}
	// данные об узи записываем
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}

	// Устанавливаем заголовки
	w.Header().Set("Content-Type", getExtension(image.FileMeta.ContentType))

	// Пишем изображение в ответ
	if _, err := w.Write(image.FileBin); err != nil { // тут тоже stream?
		http.Error(w, "Failed to write image data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Вспомогательная функция для определения расширения по MIME-типу
func getExtension(format string) string {
	switch format {
	case "image/png":
		return "png"
	case "image/jpeg":
		return "jpeg"
	case "image/tiff":
		return "tiff"
	default:
		return "bin"
	}
}

// GetDeviceList возвращает список доступных устройств.
// @Summary Получить список устройств
// @Description Возвращает список доступных устройств для УЗИ.
// @Tags Uzi
// @Produce json
// @Success 200 {array} uzimodel.Device "Список устройств"
// @Failure 502 {string} string "Ошибка обработки на сервере"
// @Router /uzi/device/list [get]
func (c *UziController) GetDeviceList(w http.ResponseWriter, r *http.Request) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	ctx := context.Background()
	data, err := c.Service.GetDeviceList(ctx)
	if err != nil {
		http.Error(w, "Failed to get uzi image with segments. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
}
