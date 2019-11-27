package sigint

import (
	"os"
	"os/signal"
	"syscall"
)

// OnTriggerExit exit on signal trigger
// Do not use this in PROD! Instead use context.Done() and propagate it
func OnTriggerExit() {
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		os.Exit(1)
	}()
}
