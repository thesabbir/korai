package sysinfo

import (
	"encoding/json"
	"fmt"
	"time"
)

type SysInfo struct {
	Uptime time.Duration `json:"uptime"`
	Memory *MemoryStats `json:"memory"`
	Disk []*DiskUsage `json:"disk_usage"`
	Cpu *CPUInfo `json:"cpu"`
}


func Info() ([]byte, error) {

	uptime, err := getUptime()
	if err != nil {
		return nil, fmt.Errorf("failed to get uptime")
	}
	memory, err := getMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory")
	}

	disk, err := getDiskUsage()
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage")
	}

	cpu, err := getCPUInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu usage and info")
	}

	info := &SysInfo{}
	info.Uptime = uptime
	info.Memory = memory
	info.Disk = disk
	info.Cpu = cpu

	data, err := json.MarshalIndent(&info, "", "")
	return data, err
}
