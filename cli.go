package main

import (
	"os"
	"log"
	"fmt"	
	"io/ioutil"
	"cli-tool/ls"
	"cli-tool/rename"
	"cli-tool/utilities"
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
		{
			Name: "rename",
			Usage: "Renames a file or a folder. `old-name` and `new-name` follows `rename` respectively and they specify relative paths.", 
			Flags: flags,
			Action: func(c *cli.Context) error {
				indexOfRename := utilities.FindIndexOfCommandInOsArgs(os.Args, "rename")

				if utilities.OsArgsHaveSufficientNumberOfArgs(len(os.Args), indexOfRename, 2) {
					oldName := os.Args[indexOfRename + 1]
					newName := os.Args[indexOfRename + 2]					

					if utilities.FileExists(oldName) {
						if newName != "" {
							rename.RenameFile(oldName, newName)
						}
					} else {
						fmt.Println(oldName, "doesn't exist! Please specify the paths correctly.")
					}
				} else {
					fmt.Println("Please specify `old-name` and `new-name` in terms of specify relative paths.")
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