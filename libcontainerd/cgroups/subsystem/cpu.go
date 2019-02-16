package subsystem

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
)

type CpuSubsystem struct {
}

func (*CpuSubsystem) GetName() string {
	return "cpu"
}

func (s *CpuSubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.GetName(), cgroupPath, true); err != nil {
		return err
	} else {
		if res.CpuShare != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpu.shares"), []byte(res.CpuShare), 0644); err != nil {
				return fmt.Errorf("set cgroup cpu share fail %v", err)
			}
		}
		return nil
	}
}

func (s *CpuSubsystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.GetName(), cgroupPath, false); err != nil {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	} else {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	}
}

func (s *CpuSubsystem) Remove(path string) error {
	panic("implement me")
}
