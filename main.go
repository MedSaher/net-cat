package main

import (
	"fmt"
	"log"
	"net"
	net_cat "net_cat/server"
	"os"
	"os/signal"
	"syscall"
)

var (
	maxConn = 3
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Execution : ./netcat [Port]")
	}
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	server := &net_cat.Server{}
	err := server.Builder(port, maxConn)
	if err != nil {
		log.Fatal(err)
	}
	addr := server.Server.Addr().String()

	// Extract the port using net.SplitHostPort
	_, port, _ = net.SplitHostPort(addr)
	fmt.Println("Server started on port " + port)
	closeSignal := make(chan os.Signal, 1)
	go closeServer(server, closeSignal)
	for {
		conn, err := server.Server.Accept()
		if err != nil {
			break
		}
		go server.RegistrNewUser(conn)

	}
	<-closeSignal
}

func closeServer(server *net_cat.Server, closeSignal chan os.Signal) {
	signal.Notify(closeSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-closeSignal
	server.CloseServer()
	os.Exit(0)
}
