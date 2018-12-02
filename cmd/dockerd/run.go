package dockerd

import (
	"github.com/Danceiny/Marxist/cli"
)

var RunCommand cli.Command = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			marx run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
}
