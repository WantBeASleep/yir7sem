package uzi

import (
	"encoding/json"
	"net/http"
	"yir/gateway/internal/uzi/models"
	pb "yir/gateway/rpc/uzi"

	s3client "yir/s3upload/pkg/client"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

type Ctrl struct {
	json jsoniter.API

	uziClient pb.UziAPIClient
	s3Client  s3client.S3Client
}

func NewCtrl() *Ctrl {
	return &Ctrl{
		json: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

// GetDeviceList возвращает список доступных устройств.
// @Summary Получить список устройств
// @Description Возвращает список доступных устройств для УЗИ.
// @Tags Uzi
// @Produce json
// @Success 200 {array} models.Device "Список устройств"
// @Failure 500 {string} string "Ошибка обработки на сервере"
// @Router /uzi/device/list [get]
func (c *Ctrl) GetDeviceList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp, err := c.uziClient.GetDeviceList(ctx, &empty.Empty{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := c.json.NewEncoder(w).Encode(models.PBGetDeviceListToDevices(resp)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetFormationWithSegments возвращает узла с сегментами по ID.
// @Summary Получить узла с сегментами
// @Description Возвращает узла с сегментами по указанному ID формации.
// @Tags Uzi
// @Produce json
// @Param formation_id path string true "ID узла"
// @Success 200 {object} models.FormationWithSegments "Данные формации с сегментами"
// @Failure 500 {string} string "Ошибка обработки на сервере"
// @Router /formation/segments/{formation_id} [get]
func (c *Ctrl) GetFormationWithSegments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	req := vars["formation_id"]
	resp, err := c.uziClient.GetFormationWithSegments(ctx, models.IdToPbId(models.Id{Id: req}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	segments := make([]models.SegmentResp, 0, len(resp.Segments))
	for _, pbseg := range resp.Segments {
		binContor, err := c.s3Client.GetFullFileByStream(ctx, pbseg.ContorUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var contor []models.Point
		if err = json.Unmarshal(binContor.FileBin, &contor); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		segments = append(segments, models.PBSegmentToSegmentResp(pbseg, contor))
	}

	reqResp := models.FormationWithSegments{
		Formation: models.PBFormationRespToFormationResp(resp.Formation),
		Segments:  segments,
	}

	if err := c.json.NewEncoder(w).Encode(reqResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// PostFormationWithSegments добавляет узел с сегментами для УЗИ.
// @Summary Добавить узел с сегментами
// @Description Создает узел с сегментами, привязанную к указанному uziID.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param uzi_id path string true "ID УЗИ"
// @Param formation body models.FormationWithNestedSegments true "Узел с сегментами"
// @Success 200 {object} models.FormationWithSegmentsIDs true "Успешное выполнение"
// @Failure 500 {string} string "Ошибка обработки на сервере"
// @Router /formation/segments/{uzi_id} [post]
func (c *Ctrl) PostFormationWithSegments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	uziID := vars["uzi_id"]
	// проверка что передали норм idшник вообще

	var req models.FormationWithNestedSegmentsReq
	if err := c.json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := c.uziClient.CreateFormationWithSegments(ctx, &pb.CreateFormationWithSegmentsRequest{
		UziId:     uziID,
		Formation: models.FormationWithNestedSegmentsReqToPB(&req),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqResp := models.CreateFormationWithSegmentsRespToFormationWithSegmentsIDs(resp)
	if err := c.json.NewEncoder(w).Encode(reqResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PatchFormation обновляет информацию о узеле.
// @Summary Обновить узел
// @Description Обновляет существующую узел по переданному ID.
// @Tags Uzi
// @Accept json
// @Produce json
// @Param formation_id path string true "ID узел"
// @Param formation body models.FormationPatch true "Данные узела"
// @Success 200 {object} models.FormationResp "Успешное выполнение"
// @Failure 500 {string} string "Ошибка обработки на сервере"
// @Router /uzi/formation/{formation_id} [patch]
func (c *Ctrl) PatchFormation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	uziID := vars["formation_id"]

	var req models.FormationReq
	if err := c.json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := c.uziClient.UpdateFormation(ctx, &pb.UpdateFormationRequest{
		Id: uziID,
		Formation: models.FormationReqToPB(&req),
	})
	if  err != nil {
		http.Error(w, "Failed to update formation. Please try again later.", http.StatusBadGateway)
		return
	}

	reqResp := models.PBFormationRespToFormationResp(resp)
	if err := c.json.NewEncoder(w).Encode(reqResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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