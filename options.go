package main

import (
	"flag"
)

const (
	HTTPEnableFlag = "http-enable"
	HTTPPortFlag   = "http-port"
)

type Options struct {
	HTTPEnable *bool
	HTTPPort   *int
}

func ParseOptions() *Options {
	opts := &Options{
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
