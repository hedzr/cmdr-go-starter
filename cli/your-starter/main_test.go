package main

import (
	"testing"

	"github.com/hedzr/cmdr/v2"
)

func Test1(t *testing.T) {
	cmdr.Store().Set("app.testing", true)
	mainRun()
}
