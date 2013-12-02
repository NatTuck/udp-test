package main

import (
	"fmt"
	"flag"
	"net"
)

func main() {
	var connect string
	var port    string

	flag.StringVar(&connect, "connect", "", "Connect to remote address.")
	flag.StringVar(&port,    "port", "7777", "Listen on local port.")

	flag.Parse()

	if connect != "" {
		udpDial(connect, ":" + port)
	} else {
		udpListen(":" + port)
	}
}

func udpListen(localAddr string) {
	fmt.Println("Listen on", localAddr)

	addr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	for {
		pkt := make([]byte, 2048)
		nn, _, flags, sadr, err := conn.ReadMsgUDP(pkt, nil)
		if err != nil {
			panic(err)
		}

		fmt.Println("Got packet")
		fmt.Println("nn    =", nn)
		fmt.Println("flags =", flags)
		fmt.Println("sadr  =", sadr)
		fmt.Println("pkt   =", string(pkt))
	}
}

func udpDial(remoteAddr string, localAddr string) {
	fmt.Println("Connect to", remoteAddr)

	addr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		panic(err)
	}

	sendAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", sendAddr)

	pkt := []byte("hallo")
	nn, _, err := conn.WriteMsgUDP(pkt, nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("sent bytes:", nn)

	for {
		pkt := make([]byte, 2048)
		nn, _, flags, sadr, err := conn.ReadMsgUDP(pkt, nil)
		if err != nil {
			panic(err)
		}

		fmt.Println("Got packet")
		fmt.Println("nn    =", nn)
		fmt.Println("flags =", flags)
		fmt.Println("sadr  =", sadr)
		fmt.Println("pkt   =", string(pkt))
	}

}
