package main

import (
	"NeogoFetch/info"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

// create splitter
func splitter(len_ int, char string) (splitValue string) {

	for i := 0; i < len_; i++ {
		splitValue += char
	}

	return splitValue
}

// prepare parameters for out
func output() []string {

	head := color.GreenString(strings.TrimSpace(info.CompareNameHost()))
	splitter := splitter(len(strings.TrimSpace(info.CompareNameHost())), "-")
	osName := color.GreenString("OS: ") + strings.TrimSpace(info.GetOsName())
	machineType := color.GreenString("Type: ") + strings.TrimSpace(info.GetMachineType())
	kernel := color.GreenString("Kernel: ") + strings.TrimSpace(info.GetKernel())
	uptime := color.GreenString("Uptime: ") + strings.TrimSpace(info.GetUptime())
	desktopSession := color.GreenString("DE: ") + strings.TrimSpace(info.GetDesktopSessionName())
	resolution := color.GreenString("Resolution: ") + strings.TrimSpace(info.GetResolution())
	cpuInfo := color.GreenString("CPU: ") + strings.TrimSpace(info.GetCPUInfo())
	gpuInfo := color.GreenString("GPU: ") + strings.TrimSpace(info.GetGPUInfo())
	memoryInfo := color.GreenString("RAM: ") + strings.TrimSpace(info.MemoryInfo())

	return []string{head, splitter, osName, machineType, kernel, uptime, desktopSession, resolution, cpuInfo, gpuInfo, memoryInfo}
}

// entry point
func main() {
	file, _ := os.ReadFile("logo.txt")

	info := output()

	splitFileData := strings.Split(string(file), "\n")

	curLen := 0

	for i := 0; i < len(info); i++ {
		if i < len(splitFileData) {
			fmt.Print(splitFileData[i])
			curLen = len(splitFileData[i])
		}
		fmt.Println(splitter(50-curLen, " "), info[i])
		curLen = 0
	}
}
