package main

import (
	"os"
	"log"	
	"io/ioutil"
	"cli-tool/ls"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "A terminal built in GoLang!"
	app.Usage = "Provides basic functionalities like in a Terminal."

	flags := []cli.Flag {
		cli.StringFlag {
			Name: "view-options",
			Value: "",
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
				
				mode := c.String("view-options")
				if mode == "" {
					ls.ShowNonHiddedFilesAndFolders(files)
				} else if mode == "a" {
					ls.ShowAllFilesAndFolders(files)
				}				
				
				return nil
			},
		},
	}
	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}