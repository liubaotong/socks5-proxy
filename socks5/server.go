package socks5

import (
	"fmt"
	"net"
	"sync"
)

type Config struct {
	Username string
	Password string
	Port     int
}

type Server struct {
	config *Config
	logger *Logger
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		logger: NewLogger(),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	s.logger.Info("服务器启动在端口 %d", s.config.Port)

	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("接受连接错误: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			s.handleConnection(conn)
		}()
	}
} 