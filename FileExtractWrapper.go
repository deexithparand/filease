package FileEase

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/DeexithParand2k2/FileEase/fileExtractor"
)

type FileStatus struct {
	AcceptableExt bool
	Extension     string
}

func checkFileExtension(file_path string) FileStatus {

	extension := filepath.Ext(file_path)

	resStatus := FileStatus{
		AcceptableExt: false,
		Extension:     extension,
	}

	if extension == ".docx" {
		resStatus.AcceptableExt = true
	} else if extension == ".txt" {
		resStatus.AcceptableExt = true
	} else if extension == ".pdf" {
		resStatus.AcceptableExt = false
	} else {
		resStatus.AcceptableExt = false
	}

	return resStatus
}

func getFileContents(file_path string, extension string) (string, error) {

	if extension == ".txt" {

		bytesContent, err := os.ReadFile(file_path)
		if err != nil {
			return "", err
		}

		return string(bytesContent), nil

	} else if extension == ".docx" {

		file, err := os.Open(file_path)
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
			return "", err
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			log.Fatalf("Error getting file info: %s", err)
			return "", err
		}

		text, err := fileExtractor.DOCX2Text(file, stat.Size())
		if err != nil {
			log.Fatalf("Error extracting text: %s", err)
			return "", err
		}
		return text, nil

	}

	return "", nil
}

// main func of module
// input : path (string)
func FileExtractWrapper(file_path string) (string, string) {

	var checkFileStatus FileStatus

	if checkFileStatus = checkFileExtension(file_path); !checkFileStatus.AcceptableExt {
		log.Fatalf("Extension Not Accessible Error: %s", checkFileStatus.Extension)
		return "", "Extension Not Accessible Error"
	}

	contents, err := getFileContents(file_path, checkFileStatus.Extension)
	if err != nil {
		log.Fatalf("File Reading Error: %s", err)
		return "", "File Reading Error"
	}

	return contents, "No Error"
}

func init() {
	fmt.Println("[FileExtractWrapper-Package] initializing ...")
}
