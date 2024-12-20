package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var (
		suffix string
	)

	app := cli.NewApp()
	app.Name = "cudair xxxx"
	app.Usage = "cudair init"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "suffix, s",
			Value:       "!!!",
			Usage:       "text after speaking something",
			Destination: &suffix,
			EnvVar:      "SUFFIX",
		},
	}

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
	}

	app.Run(os.Args)
}
