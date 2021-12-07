package main

import (
	"log"
	"os"

	"github.com/Telinga-Digital/go-structure/cli/migrate"
	"github.com/Telinga-Digital/go-structure/cli/seed"
	"github.com/Telinga-Digital/go-structure/cli/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Application CLI",
		Usage: "Command line list of application",
		Commands: []*cli.Command{
			&version.Cmd,
			&migrate.Cmd,
			&seed.Cmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
