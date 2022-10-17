package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/src/utils"
	"zinx/src/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       uint32
	msgHandler ziface.IMsgHandler
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.IP,
		Port:       utils.GlobalObject.Port,
		msgHandler: NewMsgHandler(),
	}

	return s
}

func (s *Server) Start() {
	fmt.Println("Server listening at IP:", s.IP, ", Port: ", s.Port)
	fmt.Println("MaxConn: ", utils.GlobalObject.MaxConn)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr error: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenIP error: ", err)
			return
		}

		fmt.Println("Start zinx server ", s.Name, "successfully, now listening...")

		// TODO: generate connid randomly
		var connid uint32
		connid = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP error: ", err)
				continue
			}

			dealConn := NewConnection(conn, connid, s.msgHandler)
			connid++

			go dealConn.Start()
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("Stop zinx server ", s.Name)
}

func (s *Server) Serve() {
	fmt.Println("Serve zinx server ", s.Name)
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}
