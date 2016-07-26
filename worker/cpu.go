package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type CPUInfo struct {
	NumCPUCore int
	OSType     string
	Util       float32
}

func getCPUInfo() (*CPUInfo, error) {
	info := &CPUInfo{NumCPUCore: runtime.NumCPU(), OSType: runtime.GOOS}

	var err error
	switch info.OSType {
	case "darwin":
		err = getMacCPUInfo(info)
	case "":
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
	var util float32 = 0
	for _, item := range utils {
		if val, err := strconv.ParseFloat(strings.TrimSpace(item), 32); err == nil {
			util += float32(val)
		}
	}
	info.Util = util / float32(info.NumCPUCore)
	return nil
}

func getLinuxCPUInfo(info *CPUInfo) error {
	return nil
}
