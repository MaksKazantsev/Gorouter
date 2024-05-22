package main

import (
	"fmt"
	"time"
)

const (
	// PORT_DEFAULT is a default value of port
	PORT_DEFAULT = ":3000"
	// WRITE_DEFAULT is a default value of WriteTimeout field of HTTP server
	WRITE_DEFAULT = 3 * time.Second
	// READ_DEFAULT is a default value of WriteTimeout field of HTTP server
	READ_DEFAULT = 3 * time.Second

	TIME_SECOND      = time.Second
	TIME_MILLISECOND = time.Millisecond
)

// OptionsStruct describes app options, which can be changed by user
type OptionsStruct struct {
	Addr   string
	WriteT time.Duration
	ReadT  time.Duration
}

// WithAddress changes default port of app while init
func WithAddress(addr string) Options {
	return func(o *OptionsStruct) {
		o.Addr = fmt.Sprintf(":%s", addr)
	}
}

// WithTimeout changes default timeout of app while init
func WithTimeout(writeDur, readDur time.Duration, timeT ...time.Duration) Options {
	for _, timeType := range timeT {
		switch timeType {
		case TIME_MILLISECOND:
			return func(o *OptionsStruct) {
				o.WriteT = writeDur * time.Millisecond
				o.ReadT = readDur * time.Millisecond
			}
		case TIME_SECOND:
			return func(o *OptionsStruct) {
				o.WriteT = writeDur * time.Second
				o.ReadT = readDur * time.Second
			}
		}
	}
	return func(o *OptionsStruct) {
		o.WriteT = writeDur
		o.ReadT = readDur
	}
}

// Options is a func which can be used to change default init settings
type Options func(o *OptionsStruct)
type TimeT time.Duration
