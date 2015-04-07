package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func runCommand(cmd string) bool {
	c := exec.Command(cmd)
	output, err := c.CombinedOutput()

	logCommandOutput(cmd, string(output))

	if err != nil {
		Log.WithField("cmd", cmd).Warn("Command failed")
		return false
	}

	Log.Debug("Command returned successfully")
	return true
}

func logCommandOutput(cmd string, output string) {
	pathParts := strings.Split(cmd, "/")
	filename := pathParts[len(pathParts)-1]

	if output == "" {
		return
	}

	logLines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range logLines {
		Log.Infof("[%s] %s", filename, line)
	}
}

func withPanicLogging(f func()) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("Recovered from panic(%+v)", r)

			Log.WithField("error", err).Panicf("Stopped with panic: %s", err.Error())
		}
	}()

	f()
}

func waitForSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for sig := range c {
		Log.WithField("signal", sig).Infof("Signalled.")
		switch sig {
		case syscall.SIGTERM, os.Interrupt:
			Log.Infof("Shutting down.")
			os.Exit(0)
		default:
			Log.Warnf("Unknown signal %s", sig.String())
		}
	}
}
