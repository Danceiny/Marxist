package dockerd

import (
	"github.com/Danceiny/Marxist/cli"
	"github.com/Danceiny/Marxist/container"
	"github.com/cloudflare/cfssl/log"
	"github.com/sirupsen/logrus"
)

var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in a container. DO NOT call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
