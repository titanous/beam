package beam

import (
	"net"
)

type Server struct {
}


// NewServer initializes a new beam server.
func NewServer() *Server {
	return &Server{}
}


// RegisterJob exposes the function <h> as a remote job to be invoked by clients
// under the name <name>.
func (srv *Server) RegisterJob(name string, h JobHandler) {
}

// ServeJob is the server's default job handler. It is called every time a new job is created.
func (srv *Server) ServeJob(name string, args []string, env map[string]string, streams Streamer, db DB) error {
	return nil
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


// ListenAndServe listens on the address <addr> at protocol <proto> and then
// handles incoming requests following the beam protocol.
func (srv *Server) ListenAndServe(proto, addr string) error {
	return nil
}

// Serve accepts incoming Beam connections on the listener l, and then
// serves requests on them.
func (srv *Server) Serve(l net.Listener) error {
	return nil
}
