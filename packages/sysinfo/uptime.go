package sysinfo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getUptime() (time.Duration, error) {
	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return 0, fmt.Errorf("reading /proc/uptime: %s", err)
	}
	timeString := strings.Fields(string(data))[0]
	t, err := strconv.ParseFloat(timeString, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing /proc/uptime: %s", err)
	}
	res := time.Millisecond * time.Duration(t*1000)
	return res, nil
}
