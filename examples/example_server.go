package main

import (
	"github.com/dotcloud/beam"
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
	srv := beam.NewServer()
	srv.RegisterJob("hello", JobHelloWorld)
	srv.RegisterJob("exec", JobExec)
	if err := srv.ListenAndServe("unix", "./beam.sock"); err != nil {
		return err
	}
	return nil
}


func JobHello(name string, args []string, env map[string]string, db beam.DB, streams beam.Streamer) error {
	 // FIXME: can we get away with not returning errors? Maybe the server just kills the handler on error?
	stdout := streams.Open("stdout", beam.O_WRONLY)
	fmt.Fprintf(stdout, "Hello, %s!", strings.Join(args, " "))
	return nil
}

func JobExec(name string, args []string, env map[string]string, db beam.DB, streams beam.Streamer) error {
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
	p.Stdin = streams.Open("stdin", beam.O_RDONLY)
	p.Stdout = streams.Open("stdout", beam.O_WRONLY)
	p.Stderr = streams.Open("stderr", beam.O_WRONLY)
	return p.Run()
}
