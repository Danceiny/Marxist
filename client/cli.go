package client

import (
	"fmt"
	"github.com/Danceiny/Marxist/cli"
	"github.com/Danceiny/Marxist/libcontainerd"
	"github.com/Danceiny/Marxist/libcontainerd/cgroups/subsystem"
	"github.com/sirupsen/logrus"
	"os"
)

import "github.com/Danceiny/Marxist/cmd/dockerd"

const USAGE = `cli client for the dance.club`

func Start() {
	var app = cli.NewApp()
	app.Name = "marxist-cli"
	app.Usage = USAGE
	var runCmd = dockerd.RunCommand
	// 直接调用libcontainerd
	runCmd.Action = func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArr []string
		for _, arg := range context.Args() {
			cmdArr = append(cmdArr, arg)
		}
		tty := context.Bool("ti")
		resConf := &subsystem.ResourceConfig{
			MemLimit: context.String("m"),
			CpuSet:   context.String("cpuset"),
			CpuShare: context.String("cpushare"),
		}
		libcontainerd.Run(tty, cmdArr, resConf)
		return nil
	}
	app.Commands = []cli.Command{
		dockerd.InitCommand, // 目前的架构，只能放这里了
		runCmd,
	}
	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}
