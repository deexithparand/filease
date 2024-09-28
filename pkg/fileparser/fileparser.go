package FileEase

import "io"

type Parser interface {
	Parse(file io.Reader, size int64)
}
