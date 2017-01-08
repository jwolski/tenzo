package main

import (
	"errors"
	"net"
	"time"
)

// A Connector represents a thing that can be connected to.
type Connector interface {
	Close() error
	Connect() error
}

// A Reconnector is a type that implements a reconnection scheme.
type Reconnector interface {
	Reconnect() <-chan error
}

// A TextUpdater is a type that can transmit text updates.
type TextUpdater interface {
	Update(text string) error
}

// A TextWatcher is a type that can watch for text updates.
type TextWatcher interface {
	Watch() chan string
	WatchOne(expiry time.Duration) (string, error)
}

// A Server is a Connector, TextSender, and TextWatcher.
type Server struct {
	conn        *net.TCPConn
	isConnected bool
}

func (s *Server) Close() error {
	return s.conn.Close()
}

func (s *Server) Connect() error {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 3333,
	})
	if err != nil {
		return err
	}
	s.conn = conn
	s.isConnected = true
	return nil
}

// TODO: Re-establish conn to server, if possible.
// Updates document on remote server.
func (s *Server) Send(text string) error {
	writeBuffer := []byte(text)
	_, err := s.conn.Write(writeBuffer)
	return err
}

// Sets watch on document server for document updates from
// remote clients.
func (s *Server) Watch() (<-chan string, error) {
	// If we're not connected to the server, do not allow a watch to be set.
	if !s.isConnected {
		return nil, errors.New("not connected. call Connect() first.")
	}

	// Start read loop against server
	watchChan := make(chan string)
	go func() {
		for {
			update, err := s.readOne()
			// If an error has occurred reading from the server, attempt to
			// reconnect. If that too fails, abort.
			if err != nil {
				reconnector := newReconnector(s)
				reconnChan := reconnector.reconnect()
				if err := <-reconnChan; err != nil {
					abort("failed to reconnect watch on server")
					break
				}
			}

			watchChan <- update
		}
	}()

	return watchChan, nil
}

func (s *Server) WatchOne(expiry time.Duration) (string, error) {
	// Create temporary channel to deliver a single message
	// to the client.
	tempChan := make(chan string)
	go func() {
		if update, err := s.readOne(); err == nil {
			tempChan <- update
		}
		close(tempChan)
	}()

	// Wait for update from server until timeout
	select {
	case update := <-tempChan:
		return update, nil
	case <-time.After(expiry):
		return "", errors.New("watch expired")
	}
}

func (s *Server) readOne() (string, error) {
	buffer := make([]byte, 4096)
	numBytesRead, err := s.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:numBytesRead]), nil
}
