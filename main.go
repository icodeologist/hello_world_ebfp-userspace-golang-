package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	err := rlimit.RemoveMemlock()
	ERR(err)
	//load the compiled ebpf file
	spec, err := ebpf.LoadCollectionSpec("./ebpf_code/xdp_kern.o")
	ERR(err)

	objs := struct {
		Prog *ebpf.Program `ebpf:"xdp_hello"`
	}{}

	err = spec.LoadAndAssign(&objs, nil)
	ERR(err)

	defer objs.Prog.Close()

	iface, err := net.InterfaceByName("wlan0")
	ERR(err)

	l, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.Prog,
		Interface: iface.Index,
	})
	ERR(err)
	defer l.Close()

	fmt.Println("XDP program successfully loaded and attached")
	select {}
}

func ERR(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
