// Copyright © 2021 Hedzr Yeh.
// All Rights Reserved.
// These codes and documentations are reserved for
// non-commercial and private purpose.

package main

import (
	cmdrrel "cmdr-starter/cli/app/cli/app/cmdr"
)

func init() {
	// build.New(build.NewLoggerConfigWith(true, "logrus", "debug"))
}

func main() {
	cmdrrel.Entry()
	return
}
