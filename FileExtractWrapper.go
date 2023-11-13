package fileExtractWrapper

import (
	"path/filepath"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"github.com/deexithparand2k2/go-filereader/fileExtractor"
)


type FileStatus struct{
	AcceptableExt bool
	Extension string
}

func checkFileExtension(file_path string) FileStatus {

	extension := filepath.Ext(file_path)

	resStatus := FileStatus{
		AcceptableExt : false,
		Extension : extension,
	}

	if extension==".docx"{
		resStatus.AcceptableExt = true
	} else if extension==".txt"{
		resStatus.AcceptableExt = true
	} else if extension==".pdf"{
		resStatus.AcceptableExt = false
	} else {
		resStatus.AcceptableExt = false
	}

	return resStatus
}

func getFileContents(file_path string,extension string) (string,error){

	if extension==".txt" {
		
		bytesContent,err := ioutil.ReadFile(file_path)
		if err != nil {
			return "", err
		}
		
		return string(bytesContent),nil

	} else if extension==".docx" {

		file, err := os.Open(file_path)
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
			return "",err
		}
		defer file.Close()


		stat, err := file.Stat()
		if err != nil {
			log.Fatalf("Error getting file info: %s", err)
			return "",err
		}

		text, err := fileExtractor.DOCX2Text(file, stat.Size())
		if err != nil {
			log.Fatalf("Error extracting text: %s", err)
			return "",err
		}
		return text, nil

	} else if extension==".pdf" {

		//no implementation yet

	} else {

	}
	
	return "",nil;
}


// main func of module
// input : path (string)
func FileExtractWrapper(file_path string) string{

	var checkFileStatus FileStatus

	if checkFileStatus=checkFileExtension(file_path); !checkFileStatus.AcceptableExt {
		log.Fatalf("Extension Not Accessible Error: %s", checkFileStatus.Extension)
	}

	contents,err := getFileContents(file_path,checkFileStatus.Extension)
	if err!=nil{
		log.Fatalf("File Reading Error: %s",err)
	}

	return contents
}


func init(){
	fmt.Println("[FileExtractWrapper-Package] initializing ...")
}