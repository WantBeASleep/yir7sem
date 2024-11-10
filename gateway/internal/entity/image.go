package entity

type Image struct {
	Format string
	Data   []byte
}

type ImageRequest struct {
	Format string
	Path   string
}
