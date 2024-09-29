package txt

import (
	"testing"

	"github.com/deexithparand/FileEase/pkg/fileparser"
)

func TestTxtParser(t *testing.T) {
	var expectedFileContents string = `hello world from text file`
	var filepath string = `./testdata/sample.txt`

	txtParser := Txt{}

	fileContents, err := fileparser.ParseFile(txtParser, filepath)
	if err != nil {
		t.Fatalf("Error parsing txt file : %v", err)
	}

	if fileContents != expectedFileContents {
		t.Errorf("Mismatching File Contents : Expected %s, Got %s", fileContents, expectedFileContents)
	}
}
