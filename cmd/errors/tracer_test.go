package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestStackTracer(t *testing.T) {
	defaultEntry := logrus.WithField("Blah", "blah")

	defaultErr := fmt.Errorf("boguse")
	entry := AddStackTrace(defaultEntry, defaultErr)
	require.Equal(t, defaultEntry, entry)

	err2 := errors.Wrap(defaultErr, "wrapped")
	entry = AddStackTrace(defaultEntry, err2)
	require.NotEqual(t, defaultEntry, entry)

	err3 := errors.Wrap(defaultErr, "wrapped")
	entry = AddStackTrace(defaultEntry, err3)
	require.NotEqual(t, defaultEntry, entry)

}
