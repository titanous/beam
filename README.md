# Beam

Beam is a protocol and library for service-oriented communication,
with an emphasis on real-world patterns, simplicity and not reinventing the wheel.

It was created by Solomon Hykes and is maintained by Guillaume Charmes and
the Docker team at dotCloud, inc.

## Design principles

### Real-world patterns

Modern applications are made of many loosely coupled components which communicate with each other
over the network. Typically there are 4 patterns of communication: crud, sync, jobs and streams.

1) *CRUD*: read and write structured data representing the state of another service.

2) *Sync*: same as crud, except the state is pushed continuously by the service as it changes.

3) *Jobs*: remotely execute "tasks" exposed by another service and interact with their inputs and outputs.

4) *Streams*. Send and receive continuous streams of data produced or consumed by another service.

Most applications combine more than one pattern, and typically have to use a different tool for each pattern,
as communication tend to be biased towards one of them.


### Simplicity

Beam doesn't try to magically solve all the communication problems in every application.
Rather, it offers a robust and low-level primitive which can be used as a foundation to solve
higher-level communication problems, like service discovery, job scheduling, queueing,
latency management, caching, etc.


### Don't reinvent the wheel

Beam relies heavily on existing standards and conventions.

Specifically it uses the *Redis* protocol for both *CRUD* and *sync*. In other words,
Beam is a subset of the Redis protocol, and yes, every redis client is also a valid Beam client.

Instead of the traditional (and obsolete) RPC metaphor, Beam exposes *jobs* using the familiar
unix process API. And just like processes, *jobs* can produce and consume *streams*.

## Specification


### Underlying transport

Beam is a point-to-point protocol designed for structure communication over a reliable, bi-directional octet stream.
TCP, Unix domain sockets and TLS sessions are all good transports.


### Wire protocol

The beam protocol is a *strict subset* of the Redis wire protocol, with additional semantics. You could say it's the "ReST of Redis".
In other words, all Beam commands are map to a sequence of 1 or more valid Redis commands, but not all Redis commands are valid Beam commands.

For a reference of the Redis wire protocol, see http://redis.io/topics/protocol.

### Reading and writing data

Once a session is established by the underlying transport, a *context* is exposed to the client as a redis database. The full set of redis
commands is available to read and write values on that context. Beam does not specify which keys must be accessible for read or write - that
is up to the context to implement.

Multiple clients may access the context data concurrently, just like they can access a redis database
concurrently. All commands follow the redis specification.

Keys beginning with "/beam" are reserved for beam usage, and should not be used.



### Data synchronization

A beam client can subscribe to a continuous feed of updates to the context data,
in effect maintaining a local copy of the entire context database.

This can be used to watch for changes on particular fields, for example.

Synchronization is implemented using the SYNC redis command. This means that synchronization
is available "for free" without extra work in the application, and without re-inventing the
wheel in the Beam protocol.

### Jobs

A Beam context is basically a redis database which can run jobs. A job is a mechanism for executing code remotely
within the context of the context. The implementation of the job is entirely up to the context.

Jobs can be used as an alternative to RPC. The main difference is that jobs are modeled after unix processes,
not function calls. This means the API for manipulating jobs is both unfamiliar (if you're used to rpc)
and very familiar (if you've ever used a unix shell).

A job has the following inputs:

* Command-line arguments. This is an array of strings describing the job's arguments. For example: ["ping", "www.google.com"]

* Environment. This is a set of key-value pairs organized in a hash. They are a form of configuration. For example: {"DEBUG": "1", "retries": "42}.

* Input streams. Once the job is started it may read data from any number of optional binary streams (similar to stdin in unix processes).
The streams are framed, so they can be used for discrete messages as well as continuous octet streams.

A job can produce the following outputs:

* Output streams. Symetrically to input streams, a job may write data to any number of optional binary streams (similar to stdout and stderr).
The streams are framed, so they can be used for discrete messages as well as continuous octet streams.

* Exit status. When a job exits, it yields an integer as its exit status. Just like unix processes, 0 indicates successful completion. Any other
number indicates failure. Similar failure codes should indicate similar kinds of errors.

* Changes to context. Unlike rpc calls, beam jobs can't return arbitrary data. Instead,
they have shared access to the context.

### Creating a job

Pseudo-code for creating a job from a client, using redis commands:

```
# Run the job ["exec", "ls", "-l", "/foobar"] with DEBUG=1 in the environment,
# then stream the result.

# Create a new job and get its id
nJobs = RPUSH /jobs ls
id = nJobs - 1

# Set arguments
RPUSH /jobs/$id/args ls -l /foobar

# Set environment
HSET /jobs/$id/env DEBUG 1

# Start the job
SET /jobs/$id ""

# Send stdin on an input stream, then close it
for line in readlines(stdin) {
	RPUSH /jobs/$id/in "0:$line"
}
RPUSH /jobs/$id/in "0:"

# Read output streams
while true {
	frame = BLPOP /jobs/$id/out
	id, data = split(frame, ":")
	if data == "" {
		print "Stream $id is closed"
	} else if id == "1" {
		print(stdout, $data)
	} else if id == "2" {
		print(stderr, $data)
	} else {
		print "Received data on extra channel $id"
	}
}
```


### Service addressing

Beam does not define how to map a particular context to a particular client connection. That is
the responsibility of the server. For example, all clients connecting to a certain tcp port might
share the same context (see notes on multi-tenancy), or each new connection might get its own
context generated on the fly, or maybe an HTTP gateway selects the right context from the request
url, then upgrades the connection to websocket.
