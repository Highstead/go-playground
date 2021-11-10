package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	entry := logrus.WithField("hello", "world")

	// We expect no stacktrace here
	AddStackTrace(entry, regularError()).Errorln("1")
	AddStackTrace(entry, errorsWrap()).Errorln("2")

	//We expect a stack trace here
	AddStackTrace(entry, errorPkg()).Errorln("3")
	// the wrapping of this error 'removes' the previous stack trace.  Since the stack is unique to each error
	AddStackTrace(entry, errorPkgWrap()).Errorln("4")
}

func regularError() error {
	return fmt.Errorf("error occured here")
}

func errorsWrap() error {
	return fmt.Errorf("failed: %w", regularError())
}

func errorPkg() error {
	return errors.New("error package here")
}

func errorPkgWrap() error {
	return errors.Wrap(errorPkg(), "blah")
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func AddStackTrace(entry *log.Entry, err error) *log.Entry {
	if stackErr, ok := err.(stackTracer); ok {
		entry = log.WithField("stacktrace", fmt.Sprintf("%+v", stackErr.StackTrace()))
	}
	return entry
}
