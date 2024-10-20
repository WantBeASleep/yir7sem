package entity

type ImageDataWithFormat struct {
	Format string `json:"format" validate:"required"`
	Image  []byte `json:"image" validate:"required"`
}
