package splitter

import (
	"fmt"

	"github.com/WantBeASleep/goooool/slicer"

	"uzi/internal/domain"

	"github.com/chai2010/tiff"
)

type tiffSplitter struct{}

func (tiffSplitter) splitToPng(f domain.File) ([]domain.File, error) {
	imgs2D, errs, err := tiff.DecodeAll(f.Buf)
	if err != nil {
		return nil, &SplittError{errMain: err, errImages: errs}
	}

	imgs := slicer.Flatten2DArray(imgs2D)
	res := make([]domain.File, 0, len(imgs))
	for _, img := range imgs {
		reader, err := convertToPng(img)
		if err != nil {
			return nil, fmt.Errorf("convert to png: %w", err)
		}
		res = append(res, reader)
	}

	return res, nil
}
