package main

import (
	"fmt"
	"os"

	"github.com/lukejoshuapark/mq/cmd/text"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mq",
		Usage: "Generates slick mock implementations of interfaces for testing",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generates mock implementations for all interfaces in a suppled .go input file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Usage:   "the name of the input .go file",
					},
				},
				Action: func(c *cli.Context) error {
					return text.ProcessFile(c.String("input"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
