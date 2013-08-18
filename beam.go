// Beam is a protocol and library for service-oriented communication,
// with an emphasis on real-world patterns, simplicity and not reinventing the wheel.
//
// See http://github.com/dotcloud/beam.

package beam

import (
	redis "github.com/dotcloud/go-redis-server"
)

type Server struct {
}


// NewServer initializes a new beam server.
func NewServer() srv *Server {

}

// RegisterJob exposes the function <h> as a remote job to be invoked by clients
// under the name <name>.
func (srv *Server) RegisterJob(name string, h JobHandler) {
}

// A JobHandler is a function which can be invoked as a job by beam clients.
// The API for invoking jobs resembles that of unix processes:
//  - A job is invoked under a certain <name>.
//  - It may receive arguments as a string array (<args>).
//  - It may receive an environment as a map of key-value pairs (<env>).
//  - It may read from, and write to, streams of binary data. (<streams>).
//  - It returns value which can either indicate "success" or a variety of error conditions.
//
// Additionally, a job may modify the server's database, which is shared with all other jobs.
// This is similar to how multiple unix processes share access to the same filesystem.
//
type JobHandler func(name string, args []string, env map[string]string, streams Streamer, db DB) error

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

// ListenAndSerce listens on the address <addr> at protocol <proto> and then
// handles incoming requests on incoming following the beam protocol.
func (srv *Server) ListenAndServe(proto, addr string) error {

}
