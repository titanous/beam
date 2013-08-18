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

	// OpenWrite returns a write-only interface to send data on the stream <name>.
	// If the stream hasn't been open for write access before, it is advertised as such to the peer.
	OpenWrite(name string) io.Writer

	// OpenReadWrite returns a read-write interface to send and receive on the stream <name>.
	// If the stream hasn't been open for read or write access before, it is advertised as such to the peer.
	OpenReadWrite(name string) io.ReadWriter

	// Close closes the stream <name>. All future reads will return io.EOF, and writes will return
	// io.ErrClosedPipe
	Close(name string)
}
