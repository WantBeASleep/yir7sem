package splitter

import (
	"yir/uzi/internal/domain"
)

type pngSplitter struct{}

func (pngSplitter) splitToPng(f domain.File) ([]domain.File, error) {
	return []domain.File{f}, nil
}
