package cgroups

import (
	"github.com/Danceiny/Marxist/libcontainerd/cgroups/subsystem"
	"github.com/sirupsen/logrus"
)

type CgroupManager struct {
	// cgroup在hierarchy中的路径 相当于创建的cgroup目录相对于root cgroup目录的路径
	Path string
	// 资源配置
	Resource *subsystem.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{Path: path}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Apply(c.Path, pid); err != nil {
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystem.ResourceConfig) error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Set(c.Path, res); err != nil {
			return err
		}
	}
	return nil
}

/*
释放 cgroup
*/
func (c *CgroupManager) Clear() error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup failed %v", err)
		}
	}
	return nil
}
