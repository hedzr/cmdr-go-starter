package main

import (
	"cmdr-starter/cli/your-starter/cmd"
	"github.com/hedzr/cmdr"
	"testing"
)

func Test1(t *testing.T) {
	cmdr.Set("app.testing", true)
	cmd.Entry()
}
