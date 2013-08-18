// Beam is a protocol and library for service-oriented communication,
// with an emphasis on real-world patterns, simplicity and not reinventing the wheel.
//
// See http://github.com/dotcloud/beam.
//
// github.com/dotcloud/beam/server is a server-side implementation of the Beam protocol.

package client

import (
	beam "github.com/dotcloud/beam/common"
)


type Server struct {
}


// New initializes a new beam server.
func New() srv *Server {

}


// RegisterJob exposes the function <h> as a remote job to be invoked by clients
// under the name <name>.
func (srv *Server) RegisterJob(name string, h JobHandler) {
}

// ServeJob is the server's default job handler. It is called every time a new job is created.
func (srv *Server) ServeJob(name string, args []string, env map[string]string, streams beam.Streamer, db beam.DB) error {
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
type JobHandler func(name string, args []string, env map[string]string, streams beam.Streamer, db beam.DB) error


// ListenAndSerce listens on the address <addr> at protocol <proto> and then
// handles incoming requests on incoming following the beam protocol.
func (srv *Server) ListenAndServe(proto, addr string) error {

}
