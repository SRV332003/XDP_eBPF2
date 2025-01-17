package handlers

import (
	"errors"
	"fmt"

	"github.com/SRV332003/XDP_eBPF/functions"
)

func HandleInput() (int, int, []int, error) {
	//take user input
	var inputport int
	fmt.Print("Enter the port number to block (press enter to pickup from .env): ")
	fmt.Scanln(&inputport)
	if inputport < 0 || inputport > 65535 {
		return 0, 0, nil, errors.New("Invalid port number")
	}
	if inputport == 0 {
		inputport = functions.EnvPort()
	}

	//take user input
	var ifaceName string
	fmt.Print("Enter the interface name (press enter to pickup from .env): ")
	fmt.Scanln(&ifaceName)
	if ifaceName == "" {
		ifaceName = functions.EnvIFace()
	}

	// Get the interface index.
	ifaceIndex, err := functions.GetIfaceIdex(ifaceName)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("Failed to get interface index: %v", err)
	}

	//take user input
	var process string
	fmt.Print("Enter the process name (press enter to pickup from .env): ")
	fmt.Scanln(&process)
	if process == "" {
		process = functions.EnvProcess()
	}

	//Get the process ID
	processID, err := functions.GetPIDByName(process)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("Failed to get process ID: %v", err)
	}

	//Get the ports used by the process
	ports, err := functions.GetPortByPID(processID)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("Failed to get ports for process id: %s\n%v", processID, err)
	}

	fmt.Println("------------------------------")
	fmt.Println("Interface index:", ifaceIndex, "\nInterface name:", ifaceName, "\nProcess ID:", processID, process)

	return ifaceIndex, inputport, ports, nil

}
