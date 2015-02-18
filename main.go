package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Sirupsen/logrus"
)

const (
	HTTPEnableFlag  = "http-enable"
	HTTPPortFlag    = "http-port"
	HTTPLogPathFlag = "http-log-path"
)

var (
	Log = logrus.New()

	httpEnable = flag.Bool(HTTPEnableFlag, false, "Enable HTTP server.")
	httpPort   = flag.Int(HTTPPortFlag, 8001, "HTTP server listening port.")

	stopOnce sync.Once
	wg       sync.WaitGroup
)

func withLogging(f func()) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("Recovered from panic(%+v)", r)

			Log.WithField("error", err).Panicf("Stopped with panic: %s", err.Error())
		}
	}()

	f()
}

func main() {
	// Set the log output to stderr since communication with ExaBGP is done via stdout.
	Log.Out = os.Stderr

	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			Log.WithField("signal", sig).Infof("Signalled. Shutting down.")

			stopOnce.Do(func() { shutdown(0) })
		}
	}()

	// If HTTP is enable start the HTTP server (which blocks).
	if *httpEnable {
		if err := <-StartHTTPServer(*httpPort); err != nil {
			Log.WithField("error", err).Fatal("Error starting HTTP server.")
		}
	} else {
		// If HTTP is not enabled we need to block with a wait on a WaitGroup.
		wg.Add(1)
		wg.Wait()
	}
}

func shutdown(code int) {
	Log.WithField("code", code).Infof("Stopping.")

	// If HTTP is enabled we must exit in order to cause the HTTP server to shutdown.
	if *httpEnable {
		os.Exit(0)
	} else {
		wg.Done()
	}
}
