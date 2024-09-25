package test

import (
	"strings"
	"testing"

	"github.com/DeexithParand2k2/FileEase"
)

func TestFileExtractWrapper_TxtFile(t *testing.T) {
	a := "./database/testtxt.txt"
	expected := `A text file (sometimes spelled textfile; an old alternative name is flatfile) is a kind of computer file that is structured as a sequence of lines of electronic text. A text file exists stored as data within a computer file system.`

	if got := strings.TrimSpace(FileEase.FileExtractWrapper(a)); got != expected {
		t.Errorf("FileExtractWrapper, didn't return expected file content.\nActual: %s\nExpected: %s", got, expected)
	} else {
		t.Logf("Test succeeded: Your success message here")
	}
}

func TestFileExtractWrapper_DocxFile(t *testing.T) {
	a := "./database/testdocx.docx"
	expected := `A .DOCX file`

	if got := strings.TrimSpace(FileExtractWrapper(a)); got != expected {
		t.Errorf("FileExtractWrapper, didn't return expected file content.\nActual: %s\nExpected: %s", got, expected)
	} else {
		t.Logf("Test succeeded: Your success message here")
	}
}
