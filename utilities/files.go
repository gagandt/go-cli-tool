package utilities

import (	
	"os"
	"log"
	"io/ioutil"
)

func ReturnFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir("./")
	return files, err
}

// @ toDo
// func ReturnFilesInRelativePathProvided() {

// }

// @ toDo
// func PathIsNested(fileName string) bool {
	
// }

func FileExists(fileName string) bool {
	files, err := ReturnFiles()

	if err != nil {
		log.Fatal(err)
		return false
	}
	
	for _, file := range(files) {
		if fileName == file.Name() {
			return true
		}
	}
	return false
}