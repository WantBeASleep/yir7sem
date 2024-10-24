package imagesplitter

import (
	"bytes"
	"fmt"
	"image"
	"yir/uzi/internal/entity"
)

type splitter func([]byte) ([][]byte, error)

var (
	splitters = map[string]splitter{
		"tiff": tiffSplitToPng,
		"png":  pngSplitToPng,
	}
)

func SplitToPng(img []byte) ([][]byte, error) {
	_, format, err := image.DecodeConfig(bytes.NewBuffer(img))
	if err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}

	splitter, ok := splitters[format]
	if !ok {
		return nil, entity.ErrUnsupportedFormat
	}

	splitted, err := splitter(img)
	if err != nil {
		return nil, fmt.Errorf("split image: %w", err)
	}

	return splitted, nil
}

// 	case "tiff":
// 		return "image/tiff", nil
// 	case "png":
// 		return "image/png", nil
// 	case "dicom":
// 		return "application/dicom", nil
