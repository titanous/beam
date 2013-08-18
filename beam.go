// Beam is a protocol and library for service-oriented communication,
// with an emphasis on real-world patterns, simplicity and not reinventing the wheel.
//
// See http://github.com/dotcloud/beam.
//
// github.com/dotcloud/beam/common holds functions and data structures common to the
// client and server implementations. It is typically not used directly by external
// programs.

package beam

import (
	_ "github.com/dotcloud/go-redis-server"
	"io"
)


type DB interface {

}

type Streamer interface {
	// OpenRead returns a read-only interface to receive data on the stream <name>.
	// If the stream hasn't been open for read access before, it is advertised as such to the peer.
	OpenRead(name string) io.Reader

	// ReadFrom opens a read-only interface on the stream <name>, and copies data
	// to that interface from <src> until EOF or error.
	// The return value n is the number of bytes read.
	// Any error encountered during the write is also returned.
	ReadFrom(src io.Reader, name string) (int64, error)

	// OpenWrite returns a write-only interface to send data on the stream <name>.
	// If the stream hasn't been open for write access before, it is advertised as such to the peer.
	OpenWrite(name string) io.Writer

	// WriteTo opens a write-only interface on the stream <name>, and copies data
	// from that interface to <dst> until there's no more data to write or when an error occurs.
	// The return value n is the number of bytes written.
	// Any error encountered during the write is also returned.
	WriteTo(dst io.Writer, name string) (int64, error)

	// OpenReadWrite returns a read-write interface to send and receive on the stream <name>.
	// If the stream hasn't been open for read or write access before, it is advertised as such to the peer.
	OpenReadWrite(name string) io.ReadWriter

	// Close closes the stream <name>. All future reads will return io.EOF, and writes will return
	// io.ErrClosedPipe
	Close(name string)
}
