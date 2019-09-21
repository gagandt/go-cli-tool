package main

import (
	"fmt"
	"os"
	"log"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Testing")
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}