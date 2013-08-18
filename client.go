package beam

import (
	"io"
)

type Client struct {
	Transport io.ReadWriteCloser
}

func (c *Client) Job(name string, args []string) *Job {
	return &Job{
		Name:	name,
		Args:	args,
		client:	c,
	}
}

type Job struct {
	Name		string
	Args		[]string
	Env		[]string
	Streams		Streamer
	Running		bool
	ExitStatus	int
	client		*Client
}

func (job *Job) Start() error {
	return nil
}

func (job *Job) Wait() error {
	return nil
}
