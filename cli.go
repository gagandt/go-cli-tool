package main

import (
	"fmt"
	"os"
	"log"	
	"io/ioutil"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "A terminal built in GoLang!"
	app.Usage = "Provides basic functionalities like in a Terminal."

	flags := []cli.Flag {
		cli.StringFlag {
			Name: "view-options",
			Value: "a",
		},
	}


	app.Commands = []cli.Command {
		{
			Name: "ls",
			Usage: "Lists out the files and folders in the current directory.",
			Flags: flags,
			Action: func(c *cli.Context) error {
				files, err := ioutil.ReadDir("./")
				if err != nil {
					log.Fatal(err)
				}
				for _, file := range(files) {
					fmt.Println(file.Name())
				}
				fmt.Println(c.String("view-options") == "l");
				return nil
			},
		},
	}
	

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}