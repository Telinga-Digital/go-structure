package version

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Cmd = cli.Command{
	Name:   "version",
	Usage:  "Print the version number of recome",
	Action: Execute,
}

func Execute(c *cli.Context) error {
	fmt.Println("v1.0.0")
	return nil
}
