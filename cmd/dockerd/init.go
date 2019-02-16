package dockerd

import (
	"github.com/Danceiny/Marxist/cli"
	"github.com/Danceiny/Marxist/container"
	"github.com/sirupsen/logrus"
)

var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in a container. DO NOT call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
