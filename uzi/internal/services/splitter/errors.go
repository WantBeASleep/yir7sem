package splitter

import (
	"errors"
	"fmt"
	"strings"
)

var ErrUnsupportedFormat = errors.New("format unsupported")

type SplittError struct {
	errMain   error
	errImages [][]error
}

func (e *SplittError) Error() string {
	b := new(strings.Builder)
	b.WriteString(fmt.Sprintf("main error: %v\n", e.errMain))
	for i, v := range e.errImages {
		b.WriteString(fmt.Sprintf("%d image: %v", i, errors.Join(v...)))
	}

	return b.String()
}
