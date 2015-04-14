package main

import (
	"flag"
)

const (
	DebugFlag      = "debug"
	HTTPEnableFlag = "http-enable"
	HTTPPortFlag   = "http-port"
)

type Options struct {
	Debug      *bool
	HTTPEnable *bool
	HTTPPort   *int
}

func ParseOptions() *Options {
	opts := &Options{
		Debug:      flag.Bool(DebugFlag, false, "Enable debug logging."),
		HTTPEnable: flag.Bool(HTTPEnableFlag, false, "Enable HTTP server."),
		HTTPPort:   flag.Int(HTTPPortFlag, 8001, "HTTP server listening port."),
	}

	flag.Parse()

	if opts.valid() {
		return opts
	}

	return nil
}

func (o *Options) valid() bool {
	return true
}
