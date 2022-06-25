package model

import (
	"errors"
	"fmt"
	"strings"
)

type Server struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func NewServer(ip, port string) *Server {
	return &Server{
		IP:   ip,
		Port: port,
	}
}

func (s *Server) String() string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s", s.IP, s.Port)
}

func StringToServer(str string) (*Server, error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {
		return nil, errors.New("err")
	}
	return &Server{
		IP:   list[0],
		Port: list[1],
	}, nil
}
