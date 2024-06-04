package handlers

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
)

func getInetBPF(port int, pid int) *ebpf.ProgramSpec {
	return &ebpf.ProgramSpec{
		Type:         ebpf.Kprobe,
		License:      "GPL",
		Instructions: asm.Instructions{},
	}

}
