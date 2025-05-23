package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{}

	app := &cli.App{
		Name:  "Core Service",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			srv := newService()

			if err := srv.start(); err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
