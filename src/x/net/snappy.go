package net

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"syscall"
	"time"

	"github.com/golang/snappy"
)

var snappyStreamID = []byte{0xff, 0x06, 0x00, 0x00, 0x73, 0x4e, 0x61, 0x50, 0x70, 0x59}

// SnappyConn wraps net.Conn with Snappy compression
type SnappyConn struct {
	conn          net.Conn
	reader        *snappy.Reader
	writer        *snappy.Writer
	initialized   bool
	snappyEnabled bool
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (sc *SnappyConn) Read(b []byte) (n int, err error) {
	if !sc.initialized {
		fmt.Println(">>> Initialize snappy connection")
		// Try to read the snappy stream header
		streamHeader := make([]byte, 10)
		if n, err := sc.conn.Read(streamHeader); err == nil {
			reader := io.MultiReader(bytes.NewReader(streamHeader), sc.conn)
			if bytes.Equal(streamHeader, snappyStreamID) {
				fmt.Println(">>> This is snappy connection")
				sc.snappyEnabled = true
				sc.reader = snappy.NewReader(reader)
				if sc.writer == nil { // Could be initialized if Write has already been called
					sc.writer = snappy.NewBufferedWriter(sc.conn)
				}
			} else {
				fmt.Println(">>> This is not snappy connection")
				sc.snappyEnabled = false
				sc.reader = nil
			}
			sc.initialized = true
		} else {
			return n, err
		}
	}

	if sc.snappyEnabled {
		return sc.reader.Read(b)
	} else {
		return sc.conn.Read(b)
	}
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (sc *SnappyConn) Write(b []byte) (n int, err error) {
	if !sc.initialized {
		// By default we send out snappy compressed data
		sc.writer = snappy.NewBufferedWriter(sc.conn)
	}
	if sc.snappyEnabled || !sc.initialized {
		n, err = sc.writer.Write(b)
		ferr := sc.writer.Flush()
		if ferr != nil {
			return 0, ferr
		}
	} else {
		n, err = sc.conn.Write(b)
	}
	return n, err
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (sc *SnappyConn) Close() error {
	if sc.writer != nil {
		err := sc.writer.Close()
		if err != nil {
			sc.conn.Close()
			return err
		}
	}
	return sc.conn.Close()
}

// LocalAddr returns the local network address.
func (sc *SnappyConn) LocalAddr() net.Addr {
	return sc.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (sc *SnappyConn) RemoteAddr() net.Addr {
	return sc.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
//
// Note that if a TCP connection has keep-alive turned on,
// which is the default unless overridden by Dialer.KeepAlive
// or ListenConfig.KeepAlive, then a keep-alive failure may
// also return a timeout error. On Unix systems a keep-alive
// failure on I/O can be detected using
// errors.Is(err, syscall.ETIMEDOUT).
func (sc *SnappyConn) SetDeadline(t time.Time) error {
	return sc.conn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (sc *SnappyConn) SetReadDeadline(t time.Time) error {
	return sc.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (sc *SnappyConn) SetWriteDeadline(t time.Time) error {
	return sc.conn.SetWriteDeadline(t)
}

// SyscallConn from the underlying connection
func (sc *SnappyConn) SyscallConn() (syscall.RawConn, error) {
	tcpConn := sc.conn.(*net.TCPConn)
	return tcpConn.SyscallConn()
}

// NewSnappyConnection creates a new Snappy compressed connection
func NewSnappyConnection(conn net.Conn) net.Conn {
	return &SnappyConn{conn: conn, writer: nil, reader: nil}
}
