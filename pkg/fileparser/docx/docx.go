package docx

import (
	"os"
	"strings"
)

type Docx struct{}

func (d Docx) Parse(filepath string) (string, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	// opening a file using os package gives a file of type io.Reader or io.ReaderAt
	fileContent, err := DOCX2Text(file, fileInfo.Size())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(fileContent), err

}
