package uzi

// MVP MODE
// DICOM AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

// на REFACTOR
// Очень плохо написано, 0 переиспользуемости, продекмозировать все, вынести в разные класс
// open-close из solid

// import (
// 	"bytes"
// 	"fmt"

// 	"yir/pkg/utils"
// 	"yir/s3upload/internal/entity"

// 	"image"
// 	_ "image/jpeg"
// 	"image/png"

// 	"github.com/chai2010/tiff"
// 	_ "golang.org/x/image/tiff"
// )

// func convertFormatToContentType(format string) (string, error) {
// 	switch format {
// 	case "tiff":
// 		return "image/tiff", nil
// 	case "png":
// 		return "image/png", nil
// 	case "dicom":
// 		return "application/dicom", nil
// 	default:
// 		return "", entity.ErrUnsupportedFormat
// 	}
// }

// func convertToPng(img image.Image) ([]byte, error) {
// 	buff := bytes.Buffer{}
// 	if err := png.Encode(&buff, img); err != nil {
// 		return nil, fmt.Errorf("encode png: %w", err)
// 	}

// 	return buff.Bytes(), nil
// }

// // абсолютно конченое название но я не придумал лучше
// func addMetaToImageData(img []byte) (*entity.ImageDataWithFormat, error) {
// 	_, format, err := image.DecodeConfig(bytes.NewBuffer(img))
// 	if err != nil {
// 		return nil, fmt.Errorf("decode img format: %w", err)
// 	}

// 	contentType, err := convertFormatToContentType(format)
// 	if err != nil {
// 		return nil, fmt.Errorf("convert format: %w", err)
// 	}

// 	return &entity.ImageDataWithFormat{
// 		Format:      format,
// 		ContentType: contentType,
// 		Image:       img,
// 	}, nil
// }

// func convertToPngSliceWithMeta(imgs []image.Image) ([]entity.ImageDataWithFormat, error) {
// 	res := make([]entity.ImageDataWithFormat, 0, len(imgs))
// 	for i, img := range imgs {
// 		pngEncode, err := convertToPng(img)
// 		if err != nil {
// 			return nil, fmt.Errorf("convert to png img [index %q]: %w", i, err)
// 		}

// 		contentType, _ := convertFormatToContentType("png")

// 		res = append(res, entity.ImageDataWithFormat{
// 			Format:      "png",
// 			ContentType: contentType,
// 			Image:       pngEncode,
// 		})
// 	}

// 	return res, nil
// }

// func splitImageWithMeta(img []byte) ([]entity.ImageDataWithFormat, error) {
// 	_, format, err := image.DecodeConfig(bytes.NewBuffer(img))
// 	if err != nil {
// 		return nil, fmt.Errorf("decode img format: %w", err)
// 	}

// 	switch format {
// 	case "tiff":
// 		// почему там возвращается двумерный массив никто блять не знает
// 		splitted2D, errs2D, err := tiff.DecodeAll(bytes.NewBuffer(img))
// 		if err != nil {
// 			errs := utils.Flatten2DArray(errs2D)
// 			return nil, fmt.Errorf("decode .tiff: %w", entity.NewIndexedError(err, errs))
// 		}

// 		splittedImages := utils.Flatten2DArray(splitted2D)
// 		splitted, err := convertToPngSliceWithMeta(splittedImages)
// 		if err != nil {
// 			return nil, fmt.Errorf("convert to png splitted imgs: %w", err)
// 		}

// 		return splitted, nil

// 	case "png":
// 		contentType, _ := convertFormatToContentType("png")

// 		img := entity.ImageDataWithFormat{Format: "png", ContentType: contentType, Image: img}
// 		return []entity.ImageDataWithFormat{img}, nil

// 	default:
// 		return nil, entity.ErrUnsupportedFormat
// 	}
// }
