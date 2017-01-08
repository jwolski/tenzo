package main

import (
	"errors"
	"math"
	"time"
)

const (
	defaultDelay  = 200 * time.Millisecond
	defaultMaxAtt = 10
)

// reconnector reconnects to the server
type reconnector struct {
	baseDelay time.Duration
	maxAtt    uint8
	connector Connector
}

// Creates a new reconnector with default options.
func newReconnector(connector Connector) *reconnector {
	r := &reconnector{connector: connector}
	r.baseDelay = defaultDelay
	r.maxAtt = defaultMaxAtt
	return r
}

// Calculates reconnect delay by increasing the value exponentially based on the
// number of reconnect attempts (`attNo`) that have been made thus far.
func (r *reconnector) calcDelay(attNo uint8) time.Duration {
	return r.baseDelay * time.Duration(math.Pow(2, float64(attNo)))
}

// Reconnects to server in background up to `maxAtt` number of times.
func (r *reconnector) reconnect() <-chan error {
	reconnChan := make(chan error)
	go r.reconnectLoop(reconnChan)
	return reconnChan
}

func (r *reconnector) reconnectLoop(reconnChan chan error) {
	attNo := uint8(1)
	for {
		// If we've exhausted allowable reconnect attempts, abort.
		if attNo > r.maxAtt {
			reconnChan <- errors.New("exceeded max reconnect attempts")
			break
		}

		// Attempt to re-establish a connection to the server. If a
		// connection has been re-established abort the reconnect loop.
		if err := r.connector.Connect(); err == nil {
			_ = r.connector.Close()
			reconnChan <- nil
			break
		}

		delay := r.calcDelay(attNo)
		time.Sleep(delay)
		attNo++
	}
}
