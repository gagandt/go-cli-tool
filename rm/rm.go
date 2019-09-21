package rm

import (
	"os"
	"fmt"
)

func Delete(fileName string) {
	os.Remove("./" + fileName)
	fmt.Println(fileName, "deleted.")
}
