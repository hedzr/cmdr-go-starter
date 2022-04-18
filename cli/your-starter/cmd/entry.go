package cmd

import (
	"fmt"

	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"gopkg.in/hedzr/errors.v3"
)

func Entry() {
	root := buildRootCmd()
	if err := cmdr.Exec(root, options...); err != nil {
		log.Fatalf("error occurs in app running: %+v\n", err)
	}
}

func onUnhandledErrorHandler(err interface{}) {
	if cmdr.GetBoolR("enable-ueh") {
		dumpStacks()
		// return
	}

	// do post-processing to shut down app resources if unhandled error
	// occurs to avoid resources leak. NOTE that internal.App() will
	// be closed automatically when app terminate normally.
	// internal.App().Close()

	panic(err) // re-throw it
}

func dumpStacks() {
	fmt.Printf("\n\n=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n\n", errors.DumpStacksAsString(true))
}
