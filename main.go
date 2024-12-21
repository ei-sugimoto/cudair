package main

import (
	"log"
	"os"

	"github.com/ei-sugimoto/cudair/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cudair"
	app.Usage = "cudair init"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("cudair errors:", err)
	}
}
