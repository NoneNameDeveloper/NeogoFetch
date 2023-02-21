package info

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func GetHost() string {
	host, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return host
}

func GetUserName() string {
	name, err := user.Current()

	if err != nil {
		panic(err)
	}

	return name.Name
}

func CompareNameHost() string {
	return GetUserName() + "@" + GetHost()
}

func GetOsName() string {
	cmd := exec.Command("grep", "^NAME", "/etc/os-release")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	res := strings.Replace(string(stdout), "NAME", "", -1)

	res = strings.Replace(res, "\"", "", -1)
	res = strings.Replace(res, "=", "", -1)

	return res

}

func GetMachineType() string {
	cmd := exec.Command("uname", "-m")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return string(stdout)
}

func GetKernel() string {
	cmd := exec.Command("uname", "-r")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return string(stdout)

}

func GetUptime() string {
	cmd := exec.Command("uptime", "-p")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	var res = strings.Replace(string(stdout), "up ", "", -1)

	return res
}

func GetDesktopSessionName() string {
	stdout := os.Getenv("XDG_CURRENT_DESKTOP")

	return stdout
}

func GetResolution() string {
	cmd := exec.Command("bash", "-c", "xdpyinfo | awk '/dimensions/ {print $2}'")

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return string(stdout)
}

func GetCPUInfo() string {
	cmd := exec.Command("bash", "-c", "cat /proc/cpuinfo | grep 'model name' -m1")

	stdout, err := cmd.Output()

	if err != nil {
		return "Undefined"
	}

	return strings.TrimSpace(strings.Trim(string(stdout), "model name"))[2:]
}

func GetGPUInfo() string {
	cmd := exec.Command("bash", "-c", "lspci | grep VGA")

	stdout, err := cmd.Output()

	if err != nil {
		return "Undefined"
	}

	out := strings.Split(string(stdout), ":")

	return string(out[2])
}

func processMemory(stdout string) []string {
	splitOut := strings.Split(stdout, "\n")

	memTotalStr_ := strings.Split(strings.TrimSpace(splitOut[0]), ":")[1]
	memTotalStr := strings.Split(strings.TrimSpace(memTotalStr_), " ")[0]

	memFreeStr_ := strings.Split(strings.TrimSpace(splitOut[1]), ":")[1]
	memFreeStr := strings.Split(strings.TrimSpace(memFreeStr_), " ")[0]

	return []string{memTotalStr, memFreeStr}
}

func MemoryInfo() string {
	cmd := exec.Command("bash", "-c", "cat /proc/meminfo | grep 'MemTotal\\|MemAvailable'")

	stdout, err := cmd.Output()

	if err != nil {
		return "Undefined"
	}

	memTotalStr, memFreeStr := processMemory(string(stdout))[0], processMemory(string(stdout))[1]

	memTotal, _ := strconv.ParseFloat(memTotalStr, 64)
	memTotal = memTotal / 1000

	memFree, _ := strconv.ParseFloat(memFreeStr, 64)
	memFree = memFree / 1000

	memPercent := memFree / memTotal * 100

	res := fmt.Sprintf("%.1f MB / %.1f MB (%.1f", memFree, memTotal, memPercent) + " %)"

	return res
}
