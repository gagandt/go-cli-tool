package main

import (
	"os"
	"log"
	"fmt"	
	"io/ioutil"
	"cli-tool/ls"
	"cli-tool/mv"
	"cli-tool/rm"
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
			Name: "mv",
			Usage: "Moves and Renames a file or a folder. `old-name` and `new-name` follows `rename` respectively and they specify relative paths.", 
			Flags: flags,
			Action: func(c *cli.Context) error {
				indexOfCommand := utilities.FindIndexOfCommandInOsArgs(os.Args, "mv")

				if utilities.OsArgsHaveSufficientNumberOfArgs(len(os.Args), indexOfCommand, 2) {
					oldName := os.Args[indexOfCommand + 1]
					newName := os.Args[indexOfCommand + 2]					

					if utilities.FileExists(oldName) {
						if newName != "" {
							mv.MoveFile(oldName, newName)
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
		{
			Name: "rm",
			Usage: "Deletes files and folders.",
			Flags: flags,
			Action: func(c *cli.Context) error {
				indexOfCommand := utilities.FindIndexOfCommandInOsArgs(os.Args, "rm")

				if utilities.OsArgsHaveSufficientNumberOfArgs(len(os.Args), indexOfCommand, 1) {
					fileName := os.Args[indexOfCommand + 1]
					rm.Delete(fileName)
				} else {
					fmt.Println("Please provide the file to be deleted.")
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