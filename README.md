# Beam

Beam is a protocol and library for service-oriented communication,
with an emphasis on real-world patterns, simplicity and not reinventing the wheel.

It was created by Solomon Hykes and is maintained by the Docker team at dotCloud, inc.

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

### Wire protocol

### Reading and writing data

### Data synchronization

### Jobs

### Service addressing

