package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pborges/huejack"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	huejack.SetLogger(os.Stdout)
	huejack.Handle("test", func(req huejack.Request, res *huejack.Response) {
		fmt.Println("im handling test from", req.RemoteAddr, req.RequestedOnState)
		res.OnState = req.RequestedOnState
		// res.ErrorState = true //set ErrorState to true to have the echo respond with "unable to reach device"
		return
	})

	panic(huejack.ListenAndServe(GetOutboundIP().String() + ":5000"))
}
