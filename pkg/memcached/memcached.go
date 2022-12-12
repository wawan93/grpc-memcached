package memcached

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/wawan93/grpc-memcached/pkg/errors"
)

type Memcached struct {
	conn net.Conn
}

func NewMemcached(addr string) (*Memcached, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Memcached{
		conn: conn,
	}, nil
}

func (m *Memcached) Get(key string) (string, error) {
	_, err := m.conn.Write([]byte("get " + key + "\r\n"))
	if err != nil {
		return "", err
	}

	buf := bufio.NewReader(m.conn)

	lines := make([]string, 0)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return "", err
		}
		if line == "END\r\n" {
			break
		}
		lines = append(lines, strings.TrimSpace(line))
	}

	if len(lines) == 0 {
		return "", errors.ErrNotFound
	}

	// the first line is "VALUE <key> <flags> <bytes>\r\n"
	// ignore the first line
	lines = lines[1:]

	if err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func (m *Memcached) Set(key, value string) (err error) {
	_, err = m.conn.Write([]byte("set " + key + " 0 0 " + strconv.Itoa(len(value)) + "\r\n"))
	if err != nil {
		return err
	}
	_, err = m.conn.Write([]byte(value + "\r\n"))

	buf := bufio.NewReader(m.conn)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "STORED\r\n" {
			break
		}
		if line == "ERROR\r\n" {
			break
		}
		if strings.HasPrefix(line, "CLIENT_ERROR\r\n") {
			err = fmt.Errorf(line)
			continue
		}
		if strings.HasPrefix(line, "SERVER_ERROR\r\n") {
			err = fmt.Errorf(line)
			continue
		}
	}
	return err
}

func (m *Memcached) Delete(key string) error {
	_, err := m.conn.Write([]byte("delete " + key + "\r\n"))

	buf := bufio.NewReader(m.conn)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "DELETED\r\n" {
			break
		}
		if line == "NOT_FOUND\r\n" {
			break
		}
	}
	return err
}
