// Copyright 2016 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package clog provides a logging interface for cayley packages.
package log

import "log"

// Logger is the clog logging interface.
type Logger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	V(int) bool
	SetV(level int)
}

var logger Logger = &std{
	verbosity: 0,
}

// SetLogger set the clog logging implementation.
func SetLogger(l Logger) { logger = l }

// V returns whether the current clog verbosity is above the specified level.
func V(level int) bool {
	if logger == nil {
		return false
	}
	return logger.V(level)
}

// SetV sets the clog verbosity level.
func SetV(level int) {
	if logger != nil {
		logger.SetV(level)
	}
}

// Info logs information level messages.
func Info(format string, args ...interface{}) {
	if logger != nil {
		logger.Info(format, args...)
	}
}

// Warn logs warning level messages.
func Warn(format string, args ...interface{}) {
	if logger != nil {
		logger.Warn(format, args...)
	}
}

// Error logs error level messages.
func Error(format string, args ...interface{}) {
	if logger != nil {
		logger.Error(format, args...)
	}
}

// Fatal logs fatal messages and terminates the program.
func Fatal(format string, args ...interface{}) {
	if logger != nil {
		logger.Fatal(format, args...)
	}
}

// std wraps the standard library logger.
type std struct {
	verbosity int
}

func (that *std) Info(format string, args ...interface{})  { log.Printf(format, args...) }
func (that *std) Warn(format string, args ...interface{})  { log.Printf("WARN: "+format, args...) }
func (that *std) Error(format string, args ...interface{}) { log.Printf("ERROR: "+format, args...) }
func (that *std) Fatal(format string, args ...interface{}) { log.Fatalf("FATAL: "+format, args...) }
func (that *std) V(level int) bool                         { return that.verbosity >= level }
func (that *std) SetV(level int)                           { that.verbosity = level }
