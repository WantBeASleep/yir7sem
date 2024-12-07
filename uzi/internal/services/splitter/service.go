package splitter

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"uzi/internal/domain"
)

type Service interface {
	SplitToPng(file domain.File) ([]domain.File, error)
}

type splitter interface {
	splitToPng(f domain.File) ([]domain.File, error)
}

const (
	Png  = "image/png"
	Tiff = "image/tiff"
)

var splitters = map[string]splitter{
	Png:  pngSplitter{},
	Tiff: tiffSplitter{},
}

type service struct{}

func New() Service { return &service{} }

func (s *service) SplitToPng(file domain.File) ([]domain.File, error) {
	splitter, ok := splitters[file.Format]
	if !ok {
		return nil, ErrUnsupportedFormat
	}

	bufs, err := splitter.splitToPng(file)
	if err != nil {
		return nil, err
	}

	return bufs, nil
}

func convertToPng(img image.Image) (domain.File, error) {
	b := new(bytes.Buffer)
	if err := png.Encode(b, img); err != nil {
		return domain.File{}, fmt.Errorf("encode png: %w", err)
	}

	return domain.File{Format: Png, Size: int64(b.Len()), Buf: b}, nil
}
