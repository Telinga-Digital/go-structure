package seed

import (
	"github.com/Telinga-Digital/go-structure/database/seeders"
	"github.com/urfave/cli/v2"
)

var (
	Cmd = cli.Command{
		Name:   "seed",
		Usage:  "Seed the database with test data",
		Action: Execute,
	}
)

func Execute(c *cli.Context) error {
	seeders.Execute()
	return nil
}
