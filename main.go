package main

import (
	"fmt"
	"flag"
	"net"
)

func main() {
	var listen bool

	flag.BoolVar(&listen, "listen", false, "Listen?")

	flag.Parse()

	if listen {
		udpListen()
	} else {
		udpDial()
	}
}

func udpListen() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:7777")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	for {
		pkt := make([]byte, 512)
		nn, _, flags, sadr, err := conn.ReadMsgUDP(pkt, nil)
		if err != nil {
			panic(err)
		}

		fmt.Println("Got packet")
		fmt.Println("nn    =", nn)
		fmt.Println("flags =", flags)
		fmt.Println("sadr  =", sadr)
		fmt.Println("pkt   =", pkt)


	}
}

func udpDial() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:7777")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", nil)

	pkt := []byte("hallo")
	nn, _, err := conn.WriteMsgUDP(pkt, nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("sent bytes:", nn)
}
