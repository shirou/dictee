package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/shirou/dictee"
)

const version string = "0.0.1"

var cmdAddDictionary = cli.Command{
	Name:        "add",
	Usage:       "",
	Description: "add dictionary",
	Action:      doAddDictionary,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "conf, c",
			Value:  "",
			Usage:  "Load configuration from `FILE`",
			EnvVar: "DICTEE_CONFIG",
		},
		cli.StringFlag{
			Name:  "name, n",
			Value: "",
			Usage: "Dictionary name",
		},
		cli.StringFlag{
			Name:  "file, f",
			Value: "",
			Usage: "Dictionary file",
		},
		cli.StringFlag{
			Name:  "type, t",
			Value: "apple",
			Usage: "Dictionary type (apple or stardict)",
		},
	},
}

var cmdRemoveDictionary = cli.Command{
	Name:        "remove",
	Usage:       "",
	Description: "remove dictionary",
	Action:      doRemoveDictionary,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "conf, c",
			Value:  "",
			Usage:  "Load configuration from `FILE`",
			EnvVar: "DICTEE_CONFIG",
		},
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "Dictionary name",
		},
	},
}

var cmdListDictionary = cli.Command{
	Name:        "list",
	Usage:       "",
	Description: "list dictionary",
	Action:      doListDictionary,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "conf, c",
			Value:  "",
			Usage:  "Load configuration from `FILE`",
			EnvVar: "DICTEE_CONFIG",
		},
	},
}

var commands = []cli.Command{
	cli.Command{
		Name:        "dict",
		Usage:       "dict add/remove",
		Description: "Manage dictionary",
		Subcommands: []cli.Command{
			cmdAddDictionary,
			cmdRemoveDictionary,
			cmdListDictionary,
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

	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.DebugLevel)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed: %s", err)
	}

}

func doSearch(c *cli.Context) error {
	configPath := c.String("c")
	_, d, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	word := c.Args().First()
	d.Search(word)

	return nil
}

func doAddDictionary(c *cli.Context) error {
	configPath := c.String("c")
	conf, _, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	f := c.String("f")
	if f == "" {
		return fmt.Errorf("please specify dictonary path")
	}
	name := c.String("name")
	if name == "" {
		return fmt.Errorf("please specify dictonary name")
	}

	dictType := c.String("t")
	if dictType == "" {
		return fmt.Errorf("please specify dictonary type")
	}

	if err := conf.AddDictionary(name, dictType, f); err != nil {
		return errors.Wrap(err, "doAddDictionary")
	}

	return conf.Dump()
}

func doRemoveDictionary(c *cli.Context) error {
	configPath := c.String("c")
	conf, _, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	name := c.String("name")

	if err := conf.RemoveDictionary(name); err != nil {
		return errors.Wrap(err, "doRemoveDictionary")
	}
	return conf.Dump()
}

func doListDictionary(c *cli.Context) error {
	configPath := c.String("c")
	conf, _, err := dictee.ReadConfig(configPath)
	if err != nil {
		return err
	}
	return conf.ListDictionary()
}
