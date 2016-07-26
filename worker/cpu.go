package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CPUInfo struct {
	NumCPUCore int
	OSType     string
	Util       float64
}

func getCPUInfo() (*CPUInfo, error) {
	info := &CPUInfo{NumCPUCore: runtime.NumCPU(), OSType: runtime.GOOS}

	var err error
	switch info.OSType {
	case "darwin":
		err = getMacCPUInfo(info)
	case "linux":
		err = getLinuxCPUInfo(info)
	default:
		return nil, fmt.Errorf("Unknown system type %v", info.OSType)
	}

	if err != nil {
		return nil, err
	}
	return info, nil
}

func getMacCPUInfo(info *CPUInfo) error {
	cmd := exec.Command("ps", "-A", "-o %cpu")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	var utils = strings.Split(string(stdout), "\n")
	var util float64 = 0
	for _, item := range utils {
		if val, err := strconv.ParseFloat(strings.TrimSpace(item), 64); err == nil {
			util += val
		} else {
			return fmt.Errorf("parse ps result err: %v", err)
		}
	}
	info.Util = util / float64(info.NumCPUCore)
	return nil
}

func parseStatSample() (float64, float64, error) {
	if _, err := os.Stat("/proc/stat"); err != nil {
		return 0, 0, fmt.Errorf("/proc/stat file not exists.")
	}

	cmd := exec.Command("cat", "/proc/stat")
	stdout, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	var lines = strings.Split(string(stdout), "\n")
	if len(lines) < 1 {
		return 0, 0, fmt.Errorf("Something wrong with /proc/stat file.")
	}

	vals := strings.Split(lines[0], " ")
	var idle, total float64 = 0, 0
	for i, item := range vals[2:] {
		if val, err := strconv.ParseFloat(strings.TrimSpace(item), 64); err == nil {
			if i == 3 {
				idle = val
			}
			total += val
		} else {
			return 0, 0, fmt.Errorf("parse cpu stat err: %v", err)
		}
	}
	return idle, total, nil
}

func getLinuxCPUInfo(info *CPUInfo) error {
	idle0, total0, err := parseStatSample()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	idle1, total1, err := parseStatSample()
	if err != nil {
		return err
	}

	idleTick := (idle1 - idle0)
	totalTick := (total1 - total0)
	info.Util = 100. * (totalTick - idleTick) / totalTick
	return nil
}
