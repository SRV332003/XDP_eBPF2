package functions

import (
	"bytes"
	"errors"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strconv"
)

func GetIfaceIdex(ifaceName string) (int, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}

	var ifaceIndex int
	for _, iface := range ifaces {
		if iface.Name == ifaceName {
			ifaceIndex = iface.Index
			break
		}
	}

	if ifaceIndex == 0 {
		return 0, errors.New("Interface not found")
	}

	return ifaceIndex, nil
}

func GetPIDByName(name string) (int, error) {
	cmd := exec.Command("pgrep", name)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	pidStr := string(bytes.TrimSpace(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func GetPortByPID(pid int) ([]int, error) {
	cmd := exec.Command("lsof", "-Pan", "-p", strconv.Itoa(pid), "-i")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`\s\d+.\d+.\d+.\d+:(\d+)(->|\s)`)

	portstr := re.FindAllSubmatch(output, -1)

	var ports []int

	for _, port := range portstr {
		iport, err := strconv.Atoi(string(port[1]))
		if err != nil {
			log.Println("Failed to get port from lsof output", err)
			return nil, err
		}
		ports = append(ports, iport)
	}
	log.Println(ports)

	return ports, nil
}
