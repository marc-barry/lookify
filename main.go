package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
)

var (
	Log = logrus.New()
)

func main() {
	// Set the log output to stderr since communication with ExaBGP is done via stdout.
	Log.Out = os.Stderr

	path, err := osext.ExecutableFolder()
	if err != nil {
		Log.WithField("error", err).Fatal(err)
	}

	opts := ParseOptions()

	bgp, err := NewBGP(os.Stdin)
	if err != nil {
		Log.WithField("error", err).Errorf("BGP initialization error")
		os.Exit(1)
	}

	go withPanicLogging(bgp.ReadMessages)

	if *opts.HTTPEnable {
		go func() {
			Log.Infof("Starting HTTP server on port %d", *opts.HTTPPort)
			if err := <-StartHTTPServer(*opts.HTTPPort, path); err != nil {
				Log.WithField("error", err).Fatal("Error starting HTTP server")
			}
		}()
	}

	waitForSignals()
}
