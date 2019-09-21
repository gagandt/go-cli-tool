package ls

import (
	"fmt"
	"os"
)

func ShowAllFilesAndFolders(files []os.FileInfo) {
	for _, file := range(files) {
		fmt.Println(file.Name())
	}
}

func ShowNonHiddedFilesAndFolders(files []os.FileInfo) {
	for _, file := range(files) {
		if file.Name()[0] != '.' {
			fmt.Println(file.Name())
		}
	}
}
