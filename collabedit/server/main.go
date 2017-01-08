package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
)

const (
	host           = "localhost"
	port           = 3333
	readBufferSize = 4096
)

type Server struct {
	listener              *net.TCPListener
	connections           []*net.TCPConn
	seedContent           string
	documentContent       string
	isDocumentInitialized bool
}

func newServer(seedContent string) *Server {
	return &Server{
		listener:    nil,
		connections: make([]*net.TCPConn, 0),
		seedContent: seedContent,
	}
}

func (s *Server) handleConnection(connection *net.TCPConn) {
	// Capture state used throughout function
	clientAddr := connection.RemoteAddr().String()

	defer connection.Close()

	// If there is seed data to send down to the client, send it. If writing
	// the seed data fails, assume the worst and return.
	if s.isDocumentInitialized || s.seedContent != "" {
		_, err := connection.Write([]byte(s.seedContent))
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"Failed to write seed content to client, client: %s, err: %s\n",
				clientAddr, err)
			return
		}
	}

	// Read from the connection until the end of time or until the connection
	// fails.
	for {
		buffer := make([]byte, readBufferSize)
		numBytesRead, err := connection.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"Failed read on client connection, client: %s, err: %s",
				clientAddr, err)
			break
		}

		// Update contents of shared document upon reading data from connection
		fmt.Fprintf(os.Stdout,
			"Read data from client, client: %s, numBytes: %d\n",
			clientAddr, numBytesRead)
		// Client input is expected to have a trailing \n character not meant
		// to be included in the document's contents.
		contents := strings.TrimSuffix(string(buffer[:numBytesRead]), "\n")
		s.updateDoc(clientAddr, contents)
	}
}

func (s *Server) acceptConnections() {
	for {
		connection, err := s.listener.AcceptTCP()
		// Ignore errors, for now...
		if err != nil {
			continue
		}

		fmt.Println("Client connected @", connection.RemoteAddr().String())
		// Catalog open connections.
		s.connections = append(s.connections, connection)

		// Start reading data from connections
		go s.handleConnection(connection)
	}
}

func (s *Server) start() error {
	// Create the listening socket
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	})
	if err != nil {
		return err
	}

	// Assign the server's listener for bookkeeping and connection mgmt.
	s.listener = listener

	go s.acceptConnections()
	return nil
}

func (s *Server) stop() error {
	// Close all of the client connections
	for _, connection := range s.connections {
		connection.Close()
	}

	return s.listener.Close()
}

func (s *Server) updateDoc(updaterAddr string, updateContents string) {
	// Document becomes initialized after first client updates it.
	if !s.isDocumentInitialized {
		s.isDocumentInitialized = true
	}

	// TODO: Diff update with existing contents (and merge)
	s.documentContent = updateContents
	s.seedContent = s.documentContent

	for _, client := range s.connections {
		clientAddr := client.RemoteAddr().String()
		// Do not update the contents of the client that made the update.
		if clientAddr == updaterAddr {
			continue
		}

		numBytesWritten, err := client.Write([]byte(s.documentContent))
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"Failed to write updated contents to client, client: %s\n",
				clientAddr)
		}

		fmt.Fprintf(os.Stdout,
			"Wrote data to client, client: %s, numBytes: %d\n",
			clientAddr, numBytesWritten)
	}
}

func setupSignalHandler(server *Server) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	// Wait for interrupt
	_ = <-sigChan
	fmt.Println("Server process received interrupt signal")

	// Cleanup the server after receiving interrupt
	server.stop()
	os.Exit(0)
}

func main() {
	// Setup and start server
	// TODO: Make seed content a command-line flag
	server := newServer("Here's some seed content for ya!")
	err := server.start()
	if err != nil {
		fmt.Println("Error: failed to start server")
		os.Exit(1)
	}

	// Setup signal handler for graceful server shutdown
	go setupSignalHandler(server)

	serverAddr := server.listener.Addr().String()
	fmt.Println("Server listening on " + serverAddr)

	// Block process. Let the signal handler teardown the server.
	select {}
}
