package subsystem

type ResourceConfig struct {
	MemLimit string
	CpuShare string
	CpuSet   string
}

type Subsystem interface {
	// e.g. cpu/memory
	GetName() string
	// 设置某个cgroup在这个subsystem中的资源限制
	Set(cgroupPath string, res *ResourceConfig) error
	// 将进程添加到subsystem中
	Apply(cgroupPath string, pid int) error
	// 移除某个cgroup
	Remove(cgroupPath string) error
}

var (
	SubsystemIns = []Subsystem{
		&CpusetSubsystem{},
		&MemorySubsystem{},
		&CpuSubsystem{},
	}
)
