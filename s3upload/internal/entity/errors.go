package entity

import (
	"fmt"
	"strings"
)

var (
	ErrUnsupportedFormat = fmt.Errorf("unsupported format, use .png or .tiff")
)

type IndexedError struct {
	mainError  error
	listErrors []error
}

func NewIndexedError(mainErr error, listErrs []error) *IndexedError {
	return &IndexedError{
		mainError:  mainErr,
		listErrors: listErrs,
	}
}

func (e *IndexedError) Error() string {
	indexedErrors := []string{}
	for i, idxE := range e.listErrors {
		if idxE == nil {
			continue
		}

		indexedErrors = append(indexedErrors, fmt.Sprintf("%d - %s", i, idxE.Error()))
	}

	return fmt.Sprintf("%s:\n%s", e.mainError, strings.Join(indexedErrors, "\n"))
}
