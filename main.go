package main

import (
	"fmt"
	"os"

	"github.com/ei-sugimoto/cudair/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cudair xxxx"
	app.Usage = "cudair init"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "hello",
			Usage: "if use set -t or --text",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "text, t",
					Value: "world",
					Usage: "hello xxx text",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("Hello %s\n", c.String("text"))
				return nil
			},
		},
		{
			Name:  "init",
			Usage: "cudair init",
			Action: func(c *cli.Context) error {
				if err := cmd.Initialize(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "run",
			Usage: "if use set -c or --config. default '.cudair.toml'",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Value: ".cudair.toml",
					Usage: "",
				},
			},
			Action: func(c *cli.Context) error {
				if err := cmd.Run(c.String("config")); err != nil {
					return err
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
