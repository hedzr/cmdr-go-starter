package main

//goland:noinspection GoNameStartsWithPackageName
const (
	appName = "your-starter"
	// version  = "1.0.0"
	author   = `The Examples Authors`
	desc     = `a good blueprint for you.`
	longDesc = `a good blueprint for you.

To get help for your-starter building options, run 
'your-starter --help', or 'your-starter -h'.
`
	examples = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} help
  entering help subsystem within interactive mode
$ {{.AppName}} --help
  show help screen.
`
	overview = `` //nolint:varcheck

	zero = 0

	defaultTraceEnabled  = true //nolint:varcheck
	defaultDebugEnabled  = false
	defaultLoggerLevel   = "info"
	defaultLoggerBackend = "logrus"
)
