package sysinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CPUInfo struct {
	Usage   []float64      `json:"usage"`
	Details []cpu.InfoStat `json:"details"`
}

func getCPUInfo() (*CPUInfo, error){
	info := &CPUInfo{}
	usage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu usage")
	}

	info.Usage = usage
	cpuInfo, err := cpu.Info()

	if err != nil {
		return nil, fmt.Errorf("failed to get cpu info")
	}
	info.Details = cpuInfo

	return info, nil
}

