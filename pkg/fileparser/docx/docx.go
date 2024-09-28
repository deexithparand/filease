package docx

import "io"

type docx struct{}

func (d docx) Parse(file io.Reader, size int64) (string, error) {
	return "docx file", nil
}
