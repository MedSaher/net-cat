package net_cat

import (
	"net"
	"sync"
)

type Server struct {
	Server      net.Listener
	MaxConn     int
	Connections map[net.Conn]string
	Names       map[string]bool
	History     []string
	Mutex       sync.Mutex
}
