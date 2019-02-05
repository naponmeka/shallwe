package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	var base string
	var foldersFlag string

	app := cli.NewApp()
	app.Name = "shallwe"
	app.Usage = "Shall we build?"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "base",
			Value:       "HEAD~1",
			Usage:       "the commit in which you want to compare",
			Destination: &base,
		},
		cli.StringFlag{
			Name:        "folders",
			Value:       "",
			Usage:       "the folders/files in which you want to watch (seperate with ',')",
			Destination: &foldersFlag,
		},
	}

	app.Action = func(c *cli.Context) error {
		needToBuild := false
		folders := strings.Split(foldersFlag, ",")
		if len(folders) == 0 {
			folders = append(folders, "")
		}
		outputFromDiff, err := exec.Command("bash", "-c", fmt.Sprintf("git --no-pager diff --name-only %s", base)).Output()
		if err != nil {
			log.Fatal(err)
		}
		diffedFiles := strings.Split(string(outputFromDiff[:]), "\n")
		for _, diffedFile := range diffedFiles {
			if diffedFile != "" {
				for _, folder := range folders {
					if strings.HasPrefix(diffedFile, folder) {
						needToBuild = true
						break
					}
				}
			}
		}
		if needToBuild {
			return errors.New("Need to build")
		}
		fmt.Println("Need not to build")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
