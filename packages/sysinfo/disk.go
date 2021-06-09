package sysinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"syscall"
)

type DiskUsage struct {
	Device         string `json:"device"`
	Type           string `json:"type"`
	MountPoint     string `json:"mount_point"`
	FreeSpace      uint64 `json:"free_space"`
	AvailableSpace uint64 `json:"available_space"`
	DiskSize       uint64 `json:"disk_size"`
}

func getDiskUsage() ([]*DiskUsage, error) {
	data, err := ioutil.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return nil, fmt.Errorf("detecting mounted filesystems: %s", err)
	}
	mountInfos := strings.Split(string(data), "\n")

	var res []*DiskUsage
	for _, mountInfo := range mountInfos {
		fields := strings.Fields(mountInfo)
		if len(fields) < 5 {
			continue
		}

		mount := fields[4]

		stats := &syscall.Statfs_t{}
		if err := syscall.Statfs(mount, stats); err != nil {
			// ignore error
			continue
		}

		fs := &DiskUsage{
			MountPoint:     mount,
			Type:           fields[8],
			Device:         fields[9],
			FreeSpace:      stats.Bfree * uint64(stats.Bsize),
			AvailableSpace: stats.Bavail * uint64(stats.Bsize),
			DiskSize:       stats.Blocks * uint64(stats.Bsize),
		}
		res = append(res, fs)
	}

	return res, nil
}
