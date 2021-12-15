package main

import (
	"errors"
	"log"
	"net/http"
)

// ServerCommand define the server command that responsible for serving a http server
// for ASCII image.
type ServerCommand struct {
	Options
	Delay float64
	Host  string
	Port  string
}

// server for ServerCommand setup a http server that can share the ASCII image to remote client
func (serverCommand *ServerCommand) Serve(args []string) error {
	filename := args[0]
	flushHandler, supported := NewFlushHandler(filename, &serverCommand.Options)
	if !supported {
		return errors.New("not supported file type")
	}

	err := flushHandler.Init()
	if err != nil {
		return err
	}

	http.HandleFunc("/", flushHandler.HandlerFunc())

	addr := serverCommand.Host + ":" + serverCommand.Port
	log.Println("Going to listen and serve on " + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}
