package entity

type ImageDataWithFormat struct {
	Format      string
	ContentType string
	Image       []byte
}

type ImageMetaData struct {
	ContentType string
}
