package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/player"
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
)

// FlushHandler define the basic oprations that flush image to remote server
type FlushHandler interface {
	Init() error
	HandlerFunc() func(w http.ResponseWriter, r *http.Request)
}

// supportedFlushHandlerMatchers register the supported flush handler
// and if the Match function is return true, just call the constructor
// function to build the flusher handler.
var supportedFlushHandlerMatchers = []struct {
	Match       func(string) bool
	Constructor func(string, *convert.Options) FlushHandler
}{
	{
		Match:       util.IsGif,
		Constructor: NewGifFlushHandler,
	},
}

// NewFlushHandler is factory method to create flush handler
func NewFlushHandler(filename string, options *convert.Options) (handler FlushHandler, supported bool) {
	for _, matcher := range supportedFlushHandlerMatchers {
		if matcher.Match(filename) {
			return matcher.Constructor(filename, options), true
		}
	}
	return nil, false
}

// BaseFlushHandler is a basic flush handler that define some basic opration
type BaseFlushHandler struct {
}

// Init doing nothing in base flush handler
func (handler *BaseFlushHandler) Init() error {
	return nil
}

// HandlerFunc return a empty hadnler function
func (handler *BaseFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Flush flush the string to remote client immediately
func (handler *BaseFlushHandler) Flush(w http.ResponseWriter, s string) error {
	_, err := fmt.Fprintf(w, s)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Fprintf(w, player.ClearScreen)

	// flush to the remote immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
		return nil
	}
	return errors.New("can not flush to invalid writer")
}
