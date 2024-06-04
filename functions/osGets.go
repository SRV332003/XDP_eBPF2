package functions

import (
	"bytes"
	"errors"
	"fmt"
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

func GetPIDByName(name string) (string, error) {
	cmd := exec.Command("pgrep", name)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	pidStr := string(bytes.TrimSpace(output))

	return pidStr, nil
}

func GetPortByPID(pid string) ([]int, error) {
	cmd := exec.Command("lsof", "-Pan", "-p", pid, "-i")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("command execution failed: %v, stderr: %s", err, stderr.String())
	}

	re := regexp.MustCompile(`\s\d+.\d+.\d+.\d+:(\d+)(->|\s)`)

	portstr := re.FindAllSubmatch(output, -1)

	var ports []int

	for _, port := range portstr {
		iport, err := strconv.Atoi(string(port[1]))
		if err != nil {
			return nil, fmt.Errorf("Failed to convert port to int: %v", err)
		}
		ports = append(ports, iport)
	}
	log.Println(ports)

	return ports, nil
}
