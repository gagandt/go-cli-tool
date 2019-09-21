package mv

import (
	"fmt"
	"os"
	"log"
)

func MoveFile(oldName, newName string) {
	err := os.Rename("./" + oldName, "./" + newName)
	
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(oldName, "renamed to", newName, ".")
	}	
}