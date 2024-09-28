package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/DeexithParand2k2/FileEase/pkg/fileparser"
	"github.com/DeexithParand2k2/FileEase/pkg/fileparser/docx"
)

func main() {

	// setting relative path
	// testfilename := "example.docx"
	testfilename := "example.docx"

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Unable to get current file information")
	}

	sourceDir := filepath.Dir(filename)

	filePath := filepath.Join(sourceDir, testfilename)

	fmt.Println("Current File Path : ", filePath)

	// library usage for Docx
	docxParser := docx.Docx{}

	fileContents, err := fileparser.ParseFile(docxParser, filePath)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	// library usage for Txt file
	// txtParser := txt.Txt{}

	// fileContents, err := fileparser.ParseFile(txtParser, filePath)
	// if err != nil {
	// 	log.Fatalf("Error : %v", err)
	// }

	fmt.Println("Extracted file contents : ")
	fmt.Println(fileContents)

}
