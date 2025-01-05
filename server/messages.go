package net_cat

import (
	"fmt"
	"net"
	"time"
)

func formatMsg(message string, conn net.Conn, s *Server, mod int) string {
	var msg string
	switch mod {
	case ModeJoinChat:
		msg = fmt.Sprintf("\n%s%v has joined our chat...%s\n", colorGreen, s.Connections[conn], colorReset)
	case ModeLeftChat:
		msg = fmt.Sprintf("\n%s%v has left our chat...%s\n", colorRed, s.Connections[conn], colorReset)
	case ModeSendMessage:
		s.Mutex.Lock()
		name := s.Connections[conn]
		s.Mutex.Unlock()
		time := time.Now().Format(TimeDefault)
		msg = fmt.Sprintf(msgPattern, time, name, message)
	}
	return msg
}
