// Copyright © 2021 Hedzr Yeh.
// All Rights Reserved.
// These codes and documentations are reserved for
// non-commercial and private purpose.

package main

//go:generate swag init --output ./swaggerdocs

import (
	_ "cmdr-starter/cli/app/cli/app/swaggerdocs"

	_ "github.com/alecthomas/template"
	_ "github.com/swaggo/swag"
)
