package subsystem

import (
	"io/ioutil"
	"os"
)
import "path"
import "fmt"

type MemorySubsystem struct {
}

func (s *MemorySubsystem) GetName() string {
	return "memory"
}

func (s *MemorySubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsystemCgroupPath, err := GetCgroupPath(s.GetName(), cgroupPath, true); err != nil {
		return err
	} else if res.MemLimit != "" {
		if err := ioutil.WriteFile(path.Join(subsystemCgroupPath, "memory.limit_in_bytes"), []byte(res.MemLimit), 0644); err != nil {
			return fmt.Errorf("set cgroup memory fail %v", err)
		}
	}
	return nil
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	if subSysCgroupPath, err := GetCgroupPath(s.GetName(), cgroupPath, false); err != nil {
		return err
	} else {
		return os.RemoveAll(subSysCgroupPath)
	}
}
