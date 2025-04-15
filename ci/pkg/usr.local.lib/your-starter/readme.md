# `ext` folder

The scripts inside `ext` folder will be loaded as subcmd. This feature is called [`DynCommand`](https://cmdr-docs.vercel.app/docs/cmdr.v2/concepts/g40-cmd-dyncmd/).

## Try it

To list the dynamic commands, try running `app jump --help`.
To try invoking a dynamic command, try running `app jump cpu` or `app jump mem`, etc..

```bash
$ go run ./cli/your-starter --help jump
your-starter v1.0.0 ~ Copyright © 2025 by The Examples Authors ~ All Rights Reserved.

Usage:

  $ your-starter jump [Options...][files...]

Description:

  jump command

Examples:

  jump example

Commands:
  cpu                                         ci/pkg/usr.local.lib/your-starter/ext/cpu
  disk                                        ci/pkg/usr.local.lib/your-starter/ext/disk
  mem                                         ci/pkg/usr.local.lib/your-starter/ext/mem/mem
  memory                                      ci/pkg/usr.local.lib/your-starter/ext/memory
  to                                          to command [Since: v0.1.1]

Global Flags:

  [Misc]
    -h, --help,--info,--usage                 Show this help screen (-?) [Env: HELP] (Default: true)

Type '-h'/'-?' or '--help' to get command help screen. 
More: '-D'/'--debug', '-V'/'--version', '-#'/'--build-info', '--no-color'...

$ go run ./cli/your-starter jump cpu
129.7%
$
```
