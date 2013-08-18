package main

import (
	beamserver "github.com/dotcloud/beam/server"
	beam "github.com/dotcloud/beam/common"
	"os/exec"
	"log"
	"fmt"
)

func main() {
	if err := RunServer(); err != nil {
		log.Fatal(err)
	}
}

func RunServer() error {
	srv := beamserver.New()
	srv.RegisterJob("hello", JobHelloWorld)
	srv.RegisterJob("exec", JobExec)
	if err := srv.ListenAndServe("unix", "./beam.sock"); err != nil {
		return err
	}
	return nil
}


func JobHello(name string, args []string, env map[string]string, streams beam.Streamer, db beam.DB) error {
	 // FIXME: can we get away with not returning errors? Maybe the server just kills the handler on error?
	stdout := streams.OpenWrite("stdout")
	fmt.Fprintf(stdout, "Hello, %s!", strings.Join(args, " "))
	return nil
}

func JobExec(name string, args []string, env map[string]string, streams beam.Streamer, db beam.DB) error {
	var (
		cmdName string
		cmdArgs []string
	)
	if len(args) >= 1 {
		cmdName = args[0]
	} else {
		return fmt.Errorf("Not enough arguments")
	}
	if len(args) > 1 {
		cmdArgs = args[1:]
	}
	p := exec.Command(cmdName, cmdArgs...)
	p.Stdin = streams.OpenRead("stdin")
	p.Stdout = streams.OpenWrite("stdout")
	p.Stderr = streams.OpenWrite("stderr")
}
