package main

import (
	"github.com/dotcloud/beam"
	"log"
	"fmt"
)


func main() {
	if err := RunClient(); err != nil {
		log.Fatal(err)
	}
}

func RunClient() error {
	client := &beam.Client{}
	if err := client.Connect("unix", "./beam.sock"); err != nil {
		return err
	}

	// Use case 1: crud on service state
	dbUrl, err := client.Get("/connect/db")
	if err != nil {
		return err
	}
	fmt.Printf("DB url is %s\n", dbUrl)

	// Use case 2: run a job and get streams from it
	job := client.Job("exec", "/bin/echo", "hello", "world")
	os.Stdout = job.Streams.OpenRead("stdout")
	os.Stderr = job.Streams.OpenRead("stderr")
	if err := job.Create(); err != nil {
		return err
	}
	if err := job.Start(); err != nil {
		return err
	}

	// Use case 3: synchronize service state
	db, err := client.Sync()
	if err != nil {
		return err
	}
	db.Watch("/connect/db", func(a, b beam.DB) {
		oldValue, err := a.Get("/connect/db")
		if err != nil {
			return err
		}
		newValue, err := b.Get("/connect/db")
		if err != nil {
			return err
		}
	}
}
