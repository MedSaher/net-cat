package net_cat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func (serv *Server) Builder(port string, maxConn int) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	serv.Server = listener
	serv.MaxConn = maxConn
	serv.Connections = make(map[net.Conn]string)
	serv.Names = make(map[string]bool)
	return nil
}

func (s *Server) RegistrNewUser(conn net.Conn) {
	if s.isRoomFull() {
		fmt.Fprintf(conn, "%sThe Room chat is full, come back later%s", colorGray, colorReset)
		conn.Close()
		return
	}
	fmt.Fprint(conn, GreetingsMsg)
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Registre User : %v", err)
	}
	name = strings.Replace(name, "\n", "", 1)
	err = s.addConnection(name, conn)
	if err != nil {
		fmt.Fprintf(conn, "%s%v%s", colorRed, err, colorReset)
		conn.Close()
		return
	}
	log.Printf("Connected %v", conn.RemoteAddr())
	if len(s.History) != 0 {
		s.loadArchive(conn)
	}
	s.informeOthers(conn, 0)
	s.startConv(conn)

}

func (s *Server) isRoomFull() bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return len(s.Connections) >= s.MaxConn
}

func (s *Server) addConnection(name string, conn net.Conn) error {
	if s.Names[name] {
		return fmt.Errorf("name already in use try again")
	}
	if name == "" {
		return fmt.Errorf("an empty name cannot be used")
	}
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Connections[conn] = name
	s.Names[name] = true

	return nil
}

func (s *Server) informeOthers(conn net.Conn, Mode int) {
	for con := range s.Connections {
		if con == conn {
			continue
		}
		msg := formatMsg("", conn, s, Mode) + s.getMsg("", con)
		fmt.Fprint(con, msg)
	}
	if Mode == 2 {
		s.removeConn(conn)
	}
}

func (s *Server) startConv(conn net.Conn) {
	for {
		fmt.Fprint(conn, s.getMsg("", conn))
		message, err := bufio.NewReader(conn).ReadString('\n')
		message = strings.Replace(message, "\n", "", 1)
		if err != nil {
			break
		}
		if message == "" {
			fmt.Fprintln(conn, "a message cannot be empty!")
			continue
		}
		s.sendMsg(message, conn)

	}
	log.Printf("left %v", conn.RemoteAddr())
	s.informeOthers(conn, 2)
}

func (s *Server) getMsg(message string, conn net.Conn) string {
	name := s.Connections[conn]
	time := time.Now().Format(TimeDefault)
	return fmt.Sprintf(msgPattern, time, name, message)
}

func (s *Server) sendMsg(message string, conn net.Conn) {
	msg := formatMsg(message, conn, s, ModeSendMessage)
	s.Mutex.Lock()
	for con := range s.Connections {
		if con == conn {
			continue
		}
		fmt.Fprintf(con, "\033[K")  // Clears the current line
        fmt.Fprintf(con, "\033[F")  // Move cursor up one line   
		fmt.Fprintln(con, "\n"+msg)
		fmt.Fprintf(con, "%v", s.getMsg("", con))
	}
	s.Mutex.Unlock()
	s.archiveMsg(msg)
}

func (s *Server) archiveMsg(msg string) {
	s.Mutex.Lock()
	s.History = append(s.History, msg)
	s.Mutex.Unlock()
}

func (s *Server) loadArchive(conn net.Conn) {
	s.Mutex.Lock()
	for _, msg := range s.History {
		fmt.Fprintln(conn, msg)
	}
	s.Mutex.Unlock()
}

func (s *Server) CloseServer() {
	log.Println("Closing Server")
	s.Mutex.Lock()
	for conn := range s.Connections {
		fmt.Fprintf(conn, "\n%sServer Was Closed!%s", colorRed, colorReset)
		conn.Close()
	}
	s.Mutex.Unlock()
	s.Server.Close()
	log.Println("Server Closed")
}

func (s *Server) removeConn(conn net.Conn) {
	s.Mutex.Lock()
	delete(s.Names, s.Connections[conn])
	delete(s.Connections, conn)
	s.Mutex.Unlock()

}
