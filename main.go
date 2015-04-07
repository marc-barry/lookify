package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	Log = logrus.New()
)

func main() {
	// Set the log output to stderr since communication with ExaBGP is done via stdout.
	Log.Out = os.Stderr

	opts := ParseOptions()

	go withPanicLogging(readStdin)

	if *opts.HTTPEnable {
		go func() {
			Log.Infof("Starting HTTP server on port %d", *opts.HTTPPort)
			if err := <-StartHTTPServer(*opts.HTTPPort); err != nil {
				Log.WithField("error", err).Fatal("Error starting HTTP server")
			}
		}()
	}

	waitForSignals()
}
