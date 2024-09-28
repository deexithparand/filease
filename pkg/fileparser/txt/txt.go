package txt

import (
	"os"
	"strings"
)

type Txt struct{}

func (t Txt) Parse(filepath string) (string, error) {
	fileContents, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(fileContents)), nil
}
