package main

import (
	cmdrrel "githubc.com/%REPOSITORY%/cli/%NAME%/cmdr"
)

func init() {
	// build.New(build.NewLoggerConfigWith(true, "logrus", "debug"))
}

func main() {
	cmdrrel.Entry()
	return
}
