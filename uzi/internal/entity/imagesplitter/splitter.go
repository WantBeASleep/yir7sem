package imagesplitter

import (
	"fmt"
	"yir/uzi/internal/entity"
)

// здесь нужно передават файл и возвращать тоже, но пока что не требуется
type splitter func([]byte) ([][]byte, error)

// надо разбить на сплиттеры + конвертеры в .png но пока что не требуется
var (
	contentTypeSplitters = map[string]splitter{
		"image/tiff": tiffSplitToPng,
		"image/png":  pngSplitToPng,
	}
)

func SplitToPng(img *File) ([]File, error) {
	splitter, ok := contentTypeSplitters[img.FileMeta.ContentType]
	if !ok {
		return nil, entity.ErrUnsupportedFormat
	}

	splitted, err := splitter(img.FileBin)
	if err != nil {
		return nil, fmt.Errorf("split image: %w", err)
	}

	resp := make([]File, 0, len(splitted))
	for _, s := range splitted {
		resp = append(resp, File{
			FileMeta: FileMeta{
				ContentType: "image/png",
			},
			FileBin: s,
		})
	}

	return resp, nil
}

// 	case "tiff":
// 		return "image/tiff", nil
// 	case "png":
// 		return "image/png", nil
// 	case "dicom":
// 		return "application/dicom", nil
