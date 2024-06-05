# Process Traffic Controller (Problem 2)

This is a Go project that uses eBPF to allow a given process to recieve TCP packets on the given port and droping all other network traffic. This is the assessment for recruitment process at Accuknox. Current Readme.md describes the solution of **2nd** problem statement. For other problems, visit [`AccuknowTest Repository`](https://github.com/SRV332003/AccuknoxTest "All Problems")

## Usage
This program needs to be executed as super user or previledged-user as it needs to lock memory for `ebpf maps`.
To run this project, install dependencies and execute the [`main.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/main.go "main.go") file:

```bash
git clone https://github.com/SRV332003/XDP_eBPF.git
cd XDP_eBPF
go mod tidy
sudo go run main.go
```

You will be prompted to enter the allowed port number, the interface name and the process name. If you press enter without providing any input, the values will be picked up from the environment variables.

## Environment Variables
This project uses the following environment variables:

- `PORT` : The port number to block. If not provided, the default value is `5173`.
- `IFACE` : The name of the network interface. If not provided, the default value is `wlp3s1`.
- `PROCESS` : The name of the process whose traffic is to be blocked except given port. If not provided, the default value is  `brave-browser`

```env
# .env
PORT=8080               # drop packets on port 8080
IFACE=wlp3s0            # attach program to wlp3s0 network interface
PROCESS=brave-browser   # blocks all other ports of this process
 ```
## Code Structure
- **[`main.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/main.go "main.go")** : This is the entry point of the application. It handles user input, loads the eBPF program into the kernel, and attaches the eBPF program to the given network interface.
- **[`handlers`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/inputHandler.go "handlers module") :** This module container higher level abstraction allowing operations such as handling `input` and creating `ebpf` program from the information. This `go module` contains the following files:
    - **[`handlers/inputHandler.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/inputHandler.go "handlers/inputHandler.go") :** This file contains the HandleInput function which handles `user input` or `.env` for the `port` number, the `interface` name and the `process` name. This uses the `functions module` to process info.
    - **[`handlers/xdpHandler.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/xdpHandler.go "handlers/inputHandler.go") :** This file contains the `GetXDPProgram` function which generates the program based on the given process's ports returns an object `ebpf.ProgramSpec` that contains instructions as the part of `ebpf program`. This file is made to separate the assembly code written using `github.com/cilium/ebpf/asm` module from the rest of linking code and also to handle it according to the information.
- **[`functions`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions "functions module") :**
    - **[`functions/envGets.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions/envGets.go "functions/envGets.go") :** This file contains functions to get environment variables.
    - **[`functions/osGets.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions/osGets.go "functions/envGets.go") :** This file contains functions to get OS-related information.

## Control Flow
The control flow of this application is straightforward and easy to follow. 
- It starts in the `main` function, where it handles input, removes the memlock limit, and gets the XDP program according to the input. 
- The `HandleInput` function is responsible for parsing the `input` and returning the necessary data like ports to be blocked, port to be passed, and the network interace. 
- The `GetXDPProgram` function is responsible for getting the `XDP` program based on the allowed port and blocked ports of that process.
- Finally, it gets the XDP program with the `GetXDPProgram` function.

## Dependencies
This project uses the following dependencies:

- `github.com/cilium/ebpf`: A package to work with eBPF programs in Go.
- `github.com/joho/godotenv`: A package to load environment variables from a .env file.