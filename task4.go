package main

import (
	"context"
	"errors"
	"fmt"
)

func NewPerson(name string, age int) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	return p
}

type Server struct {
	Port int
	Host string
	TLS  bool
}

func NewServer(port int, host string, TLS bool) *Server {
	return &Server{Port: port, Host: host, TLS: TLS}
}

// ///////////////////////////////////////
type Worker interface {
	Do(ctx context.Context) error
}

type FileProcessor struct {
	FilePath string
}

func (fp *FileProcessor) Do(ctx context.Context) error {
	if fp.FilePath == "/path/to/file" {
		return nil
	} else {
		return errors.New("file not found")
	}
}

type NetworkFetcher struct {
	URL string
}

func (nf *NetworkFetcher) Do(ctx context.Context) error {
	if nf.URL == "https/ok" {
		return nil
	} else {
		return errors.New("not ok")
	}
}

// /////////////////////////////////
type Logger interface {
	Log(msg string)
}

type StdoutLogger struct{}

func (l StdoutLogger) Log(msg string) {
	fmt.Println(msg)
}

type MockLogger struct {
	CallLoog bool
}

func (m MockLogger) Log(msg string) {
	m.CallLoog = true
	fmt.Println(msg)
}

func Work(l Logger) {
	l.Log("test")
}
func Test() {
	test := &MockLogger{}
	Work(test)
	if !test.CallLoog {
		fmt.Println("Fail")
	}

}
