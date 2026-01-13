package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
)

// GetCPUTemp 获取 CPU 温度
func GetCPUTemp() string {
	if runtime.GOOS == "windows" {
		return "45.5°C (Simulated)"
	}

	path := "/sys/class/thermal/thermal_zone0/temp"
	data, err := os.ReadFile(path)
	if err != nil {
		return "N/A"
	}

	raw := strings.TrimSpace(string(data))
	if len(raw) >= 3 {
		return raw[:2] + "." + raw[2:3] + "°C"
	}
	return raw + "°C"
}

// GetCPUUsage 获取 CPU 使用率
func GetCPUUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return 0.0
	}
	return percent[0]
}