package mappers

import (
	"yir/uzi/internal/entity"
	"yir/uzi/internal/utils"

	"github.com/google/uuid"
)

func HttpImagesToImages(http []entity.HttpImage, uziID uuid.UUID) []entity.Image {
	images := utils.MustTransformSlice[entity.HttpImage, entity.Image](http)
	for i := range images {
		images[i].UziID = uziID
	}

	return images
}

func UziDeviceImagesToUziMeta(uzi *entity.Uzi, device *entity.Device, images []entity.Image) *entity.GetMetaUziResponse {
	resp := entity.GetMetaUziResponse{}
	resp.Uzi = *utils.MustTransformObj[entity.Uzi, entity.HttpUziWithDevice](uzi)
	resp.Uzi.Device = *device
	
	httpImages := utils.MustTransformSlice[entity.Image, entity.HttpImage](images)
	resp.Images = httpImages
	return &resp
}
