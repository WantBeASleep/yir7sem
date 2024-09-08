package mappers

import (
	"yir/uzi/internal/entity"
	"yir/uzi/internal/repositories/db/models"

	"github.com/google/uuid"
)

// experimental
type EntityToModel struct {
}

func (EntityToModel) Device(ent *entity.Device) *models.Device {
	return &models.Device{
		Id:   uint64(ent.Id),
		Name: ent.Name,
	}
}

func (EntityToModel) Tirads(ent *entity.Tirads) *models.Tirads {
	return &models.Tirads{
		Tirads1: ent.Tirads1,
		Tirads2: ent.Tirads2,
		Tirads3: ent.Tirads3,
		Tirads4: ent.Tirads4,
		Tirads5: ent.Tirads5,
	}
}

func (EntityToModel) Uzi(ent *entity.Uzi) *models.Uzi {
	return &models.Uzi{
		Uuid:        ent.Uuid.String(),
		Url:         ent.Url,
		Projection:  ent.Projection,
		PatientUUID: ent.PatientUUID.String(),
		DeviceID:    uint64(ent.DeviceID),
	}
}

func (EntityToModel) Image(ent *entity.Image) *models.Image {
	return &models.Image{
		Uuid:    ent.Uuid.String(),
		Url:     ent.Url,
		Page:    int64(ent.Page),
		UziUUID: ent.UziUUID.String(),
	}
}

type ModelToEntity struct {
}

func (ModelToEntity) Device(mdl *models.Device) *entity.Device {
	return &entity.Device{
		Id:   int(mdl.Id),
		Name: mdl.Name,
	}
}

func (ModelToEntity) Tirads(mdl *models.Tirads) *entity.Tirads {
	return &entity.Tirads{
		Tirads1: mdl.Tirads1,
		Tirads2: mdl.Tirads2,
		Tirads3: mdl.Tirads3,
		Tirads4: mdl.Tirads4,
		Tirads5: mdl.Tirads5,
	}
}

func (ModelToEntity) Uzi(mdl *models.Uzi) *entity.Uzi {
	return &entity.Uzi{
		Uuid:        uuid.MustParse(mdl.Uuid),
		Url:         mdl.Url,
		Projection:  mdl.Projection,
		PatientUUID: uuid.MustParse(mdl.PatientUUID),
		DeviceID:    int(mdl.DeviceID),
	}
}

func (ModelToEntity) Image(mdl *models.Image) *entity.Image {
	return &entity.Image{
		Uuid:    uuid.MustParse(mdl.Uuid),
		Url:     mdl.Url,
		Page:    int(mdl.Page),
		UziUUID: uuid.MustParse(mdl.UziUUID),
	}
}
