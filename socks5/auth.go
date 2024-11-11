package socks5

import (
	"errors"
	"net"
)

func (s *Server) authenticate(conn net.Conn) error {
	// 读取版本和认证方法数量
	buf := make([]byte, 2)
	if _, err := conn.Read(buf); err != nil {
		return err
	}

	if buf[0] != 0x05 { // SOCKS5
		return errors.New("不支持的协议版本")
	}

	methodCount := int(buf[1])
	methods := make([]byte, methodCount)
	if _, err := conn.Read(methods); err != nil {
		return err
	}

	// 发送认证方法选择消息
	if _, err := conn.Write([]byte{0x05, 0x02}); err != nil {
		return err
	}

	// 验证用户名密码
	if err := s.verifyCredentials(conn); err != nil {
		return err
	}

	return nil
}

func (s *Server) verifyCredentials(conn net.Conn) error {
	// 实现用户名密码验证逻辑
	// ...
	return nil
} 