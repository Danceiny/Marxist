package libcontainerd

import (
	"github.com/Danceiny/Marxist/container"
	"github.com/Danceiny/Marxist/libcontainerd/cgroups"
	"github.com/Danceiny/Marxist/libcontainerd/cgroups/subsystem"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Run(tty bool, commandArr []string, res *subsystem.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logrus.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
	cgroupManager := cgroups.NewCgroupManager("marxist-cgroup")
	defer cgroupManager.Clear()
	if err := cgroupManager.Set(res); err != nil {
		logrus.Error(err)
		os.Exit(-2)
	}
	if err := cgroupManager.Apply(parent.Process.Pid); err != nil {
		logrus.Error(err)
		os.Exit(-3)
	}
	sendInitCommand(commandArr, writePipe)
	if err := parent.Wait(); err != nil {
		logrus.Errorf("wait parent: %v", err)
	}
	os.Exit(0)
}

func sendInitCommand(commandArr []string, writePipe *os.File) {
	command := strings.Join(commandArr, " ")
	logrus.Infof("command all is %s", command)
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}
