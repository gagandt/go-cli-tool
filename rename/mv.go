package rename

import (
	"fmt"
	"os"
	"log"
)

func RenameFile(oldName, newName string) {
	err := os.Rename("./" + oldName, "./" + newName)
	
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(oldName, "renamed to", newName, ".")
	}	
}