package sigint

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSignalHandler handles SIGINT
func WithSignalHandler(pCtx context.Context) context.Context {
	ctx, cancel := context.WithCancel(pCtx)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	return ctx
}

// OnTriggerExit exit on signal trigger
func OnTriggerExit() {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		os.Exit(1)
	}()
}
