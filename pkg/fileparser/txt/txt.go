package txt

import "io"

type txt struct{}

func (t txt) Parse(file io.Reader, size int64) (string, error) {
	return "txt file", nil
}
