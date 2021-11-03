package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"
)

type Config struct {
}

type Server struct {
	Addr string
	Port int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns 
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server {
		Addr: addr,
		Port: port,
		Protocol: "tcp",
		Timeout: 30 * time.Second,
		MaxConns: 1000,
		TLS: nil,
	}

	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func main() {
	s1, _ := NewServer("127.0.0.1", 3306)
	fmt.Printf("s1: %+v\n", s1)
	s2, _ := NewServer("localhost", 6379, Protocol("udp"))
	fmt.Printf("s2: %+v\n", s2)
	s3, _ := NewServer("0.0.0.0", 80, MaxConns(99999), TLS(&tls.Config{
		RootCAs:                     &x509.CertPool{},
		NextProtos:                  []string{},
		ServerName:                  "",
		ClientAuth:                  0,
		ClientCAs:                   &x509.CertPool{},
		InsecureSkipVerify:          false,
		CipherSuites:                []uint16{'a', 'b'},
		PreferServerCipherSuites:    false,
		SessionTicketsDisabled:      false,
		SessionTicketKey:            [32]byte{'A','B'},
		ClientSessionCache:          nil,
		MinVersion:                  0,
		MaxVersion:                  0,
		CurvePreferences:            []tls.CurveID{},
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}))
	fmt.Printf("s3: %+v\n", s3)
}






