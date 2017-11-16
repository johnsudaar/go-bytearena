package transport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type Transport interface {
	Connect() (chan interface{}, error)
	Send(interface{}) error
}

type TCPTransport struct {
	Host string
	Port int
	Conn net.Conn
}

func FromEnv() (Transport, error) {
	transport := &TCPTransport{}
	var err error

	if os.Getenv("PORT") == "" {
		return nil, errors.New("no port specified")
	} else {
		transport.Port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			return nil, errors.Wrap(err, "invalid port")
		}
	}

	if os.Getenv("HOST") == "" {
		return nil, errors.New("no host specified")
	}
	transport.Host = os.Getenv("HOST")
	return transport, nil
}

func (t *TCPTransport) Connect() (chan interface{}, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port))
	if err != nil {
		return nil, errors.Wrap(err, "fail to connect to host")
	}
	t.Conn = conn

	respChan := make(chan interface{}, 5)
	go t.start(respChan)

	return respChan, nil
}

func (t *TCPTransport) start(c chan interface{}) {
	reader := bufio.NewReader(t.Conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			c <- err
		} else {
			c <- line
		}
	}
}

func (t *TCPTransport) Send(value interface{}) error {
	err := json.NewEncoder(t.Conn).Encode(value)
	if err != nil {
		return errors.Wrap(err, "fail to send json")
	}
	return nil
}
