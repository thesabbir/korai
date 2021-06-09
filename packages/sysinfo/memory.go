package sysinfo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type MemoryStats struct {
	FreeMem      uint64 `json:"free_mem"`
	TotalMem     uint64 `json:"total_mem"`
	AvailableMem uint64 `json:"available_mem"`
	Buffers      uint64 `json:"buffers"`
	Cached       uint64 `json:"cached"`
	FreeSwapMem  uint64 `json:"free_swap_mem"`
	TotalSwapMem uint64 `json:"total_swap_mem"`
}

func getMemory() (*MemoryStats, error) {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("detecting mounted filesystems: %s", err)
	}
	memInfos := strings.Split(string(data), "\n")

	res := &MemoryStats{}
	for _, l := range memInfos {
		fields := strings.Fields(l)
		if len(fields) < 2 {
			continue
		}
		tag := fields[0]
		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}
		switch tag {
		case "MemTotal:":
			res.TotalMem = val
		case "MemFree:":
			res.FreeMem = val
		case "MemAvailable:":
			res.AvailableMem = val
		case "Buffers:":
			res.Buffers = val
		case "Cached:":
			res.Cached = val
		case "SwapTotal:":
			res.TotalSwapMem = val
		case "SwapFree:":
			res.FreeSwapMem = val
		}
	}
	return res, nil
}
