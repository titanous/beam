package beam

import (
	"net"
	"fmt"
)

type Server struct {
	handlers map[string]JobHandler
}


// NewServer initializes a new beam server.
func NewServer() *Server {
	return &Server{
		handlers: make(map[string]JobHandler),
	}
}


// RegisterJob exposes the function <h> as a remote job to be invoked by clients
// under the name <name>.
func (srv *Server) RegisterJob(name string, h JobHandler) {
	srv.handlers[name] = h
}

// ServeJob is the server's default job handler. It is called every time a new job is created.
// It looks up a handler registered at <name>, and calls it with the same arguments. If no handler
// is registered, it returns an error.
func (srv *Server) ServeJob(name string, args []string, env map[string]string, streams Streamer, db DB) error {
	h, exists := srv.handlers[name]
	if !exists {
		return fmt.Errorf("No such job: %s", name)
	}
	return h(name, args, env, streams, db)
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

// Serve serves request on the connection conn.
func (srv *Server) ServeConn(conn net.Conn) error {
	return nil
}
