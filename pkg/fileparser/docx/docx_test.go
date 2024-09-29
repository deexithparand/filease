package docx

import (
	"testing"

	"github.com/deexithparand/FileEase/pkg/fileparser"
)

func TestDocxParser(t *testing.T) {

	var expectedFileContents string = `hello world from docx`
	var filepath string = `./testdata/sample.docx`

	docxParser := Docx{}

	fileContents, err := fileparser.ParseFile(docxParser, filepath)
	if err != nil {
		t.Fatalf("Error parsing the file %v", err)
	}

	if fileContents != expectedFileContents {
		t.Fatalf("Mismatching File contents : Expected %s, got %s", expectedFileContents, fileContents)
	}

}
