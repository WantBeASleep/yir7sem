package imagesplitter

import (
	"bytes"
	"errors"
	"fmt"
	"yir/pkg/utils"

	"image"
	"image/png"

	"github.com/chai2010/tiff"
	_ "golang.org/x/image/tiff"
)

func convertImageToPng(img image.Image) ([]byte, error) {
	buff := bytes.Buffer{}
	if err := png.Encode(&buff, img); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}

	return buff.Bytes(), nil
}

func tiffSplitToPng(img []byte) ([][]byte, error) {
	// почему там возвращается двумерный массив никто блять не знает
	splitted2D, errs2D, err := tiff.DecodeAll(bytes.NewBuffer(img))
	if err != nil {
		errs := utils.Flatten2DArray(errs2D)
		return nil, fmt.Errorf("decode .tiff: %w, errs per pages: %w", err, errors.Join(errs...))
	}

	res := [][]byte{}
	splittedImages := utils.Flatten2DArray(splitted2D)

	for _, img := range splittedImages {
		byteImg, err := convertImageToPng(img)
		if err != nil {
			return nil, fmt.Errorf("convert .tiff layer to png: %w", err)
		}

		res = append(res, byteImg)
	}

	return res, nil
}

func pngSplitToPng(img []byte) ([][]byte, error) {
	return [][]byte{img}, nil
}
