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

// отдаем метаданные узи в uziService + отдаем S3 картинку
// в кафку пишем uziID(из того вернул uziService), возвращаем клиенту 200
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

// Картинку не возвращать!
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
