package mappers

import (
	"yir/uzi/internal/repositories/db/models"
	"yir/uzi/internal/entity"
)

//experimental
type EntityToModel struct {
}

func (EntityToModel) Device(ent *entity.Device) *models.Device {
	return &models.Device{
		Id: uint64(ent.Id),
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

type ModelToEntity struct {
}

func (ModelToEntity) Device(mdl *models.Device) *entity.Device {
	return &entity.Device{
		Id: int(mdl.Id),
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