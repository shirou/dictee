package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/shirou/dictee"
)

const version string = "0.0.1"

var commands = []cli.Command{
	{
		Name:        "index",
		Usage:       "",
		Description: "make index",
		Action:      doIndex,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "c",
				Value:  "",
				Usage:  "config file",
				EnvVar: "DICTEE_CONFIG",
			},
		},
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "dictee"
	app.Version = version
	app.Usage = `dictee search <word>`
	app.Author = "WAKAYAMA Shirou"
	app.Email = "shirou.faw@gmail.com"
	app.Compiled = time.Now()
	app.Action = doSearch
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "c",
			Value:  "",
			Usage:  "config file",
			EnvVar: "DICTEE_CONFIG",
		},
	}

	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed: %s", err)
	}

}

func doSearch(c *cli.Context) error {
	configPath := c.String("c")
	d, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	word := c.Args().First()
	d.Search(word)

	return nil
}

func doIndex(c *cli.Context) error {
	configPath := c.String("c")
	d, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	d.MakeIndex()
	return nil
}
