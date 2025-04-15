# %REPOSITORY%

![Go](https://github.com/%REPOSITORY%/workflows/Go/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/%REPOSITORY%)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/%REPOSITORY%.svg?label=release)](https://github.com/%REPOSITORY%/releases)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/%REPOSITORY%)
[![Go Report Card](https://goreportcard.com/badge/github.com/%REPOSITORY%)](https://goreportcard.com/report/github.com/%REPOSITORY%)<!--
[![codecov](https://codecov.io/gh/%REPOSITORY%/branch/master/graph/badge.svg)](https://codecov.io/gh/%REPOSITORY%)
[![Coverage Status](https://coveralls.io/repos/github/%REPOSITORY%/badge.svg?branch=master)](https://coveralls.io/github/%REPOSITORY%?branch=master) -->
[![DocSite](https://img.shields.io/badge/Doc-Site-blue)](https://docs.hedzr.com/docs/cmdr.v2/)

A template repository to build your first Golang Open Source app based [cmdr](https://github.com/hedzr/cmdr).

> powered by [cmdr](https://github.com/hedzr/cmdr) v2.1+.

## fast guide

1. All you have to do is click the <kbd>Use this template</kbd> button upon this page.
2. run! (`go run ./cli/app/cli/app/main.go`)

## with command-line

### new repo

1. clone cmdr-go-starter as a template

   ```bash
   # clone cmdr-go-starter as a template
   git clone https://github.com/hedzr/cmdr-go-starter.git new-repo
   cd new-repo
   git push git@github.com:yourname/new-repo.git +master:master
   ```

2. clone the `new-repo` to your working directory:

   ```bash
   # in the working directory of your new-repo
   cd ~/work
   git clone git@github.com:yourname/new-repo.git
   cd new-repo
   ```

3. do rename stuffs ...

## Getting Started (For the generated golang project)

To run the CLI app:

```bash
# go install -v github.com/swaggo/swag/cmd/swag
go generate ./...               # run it once at least, for gen the swagger-doc files from skeletons
go run ./cli/%NAME%/      # build the CLI app
```

For a first sight, you can check if the config files under `./ci/etc/%NAME%/` are loaded properly. A sample result should be:

```bash
$ HELP=1 go run ./cli/%NAME% ~~debug	# `HELP=1` is equivalent with `--help`

Store:
  app.                          <B>
    cmd.                        <B>
      generate.                 <B>
        manual.                 <B>
          dir                   <L> app.cmd.generate.manual.dir => 
          type                  <L> app.cmd.generate.manual.type => 1
        doc.dir                 <L> app.cmd.generate.doc.dir => 
        shell.                  <B>
          Shell                 <L> app.cmd.generate.shell.Shell => auto
          dir                   <L> app.cmd.generate.shell.dir => 
          output                <L> app.cmd.generate.shell.output => 
          auto                  <L> app.cmd.generate.shell.auto => true
          zsh                   <L> app.cmd.generate.shell.zsh => false
          bash                  <L> app.cmd.generate.shell.bash => false
          fi                    <B>
            sh                  <L> app.cmd.generate.shell.fish => false
            g                   <L> app.cmd.generate.shell.fig => false
          powershell            <L> app.cmd.generate.shell.powershell => false
          elvish                <L> app.cmd.generate.shell.elvish => false
      strict-mode               <L> app.cmd.strict-mode => false
      no-                       <B>
        env-overrides           <L> app.cmd.no-env-overrides => false
        color                   <L> app.cmd.no-color => false
      v                         <B>
        er                      <B>
          bose                  <L> app.cmd.verbose => false
          sion                  <L> app.cmd.version => false
            -sim                <L> app.cmd.version-sim => 
        alue-type               <L> app.cmd.value-type => false
      quiet                     <L> app.cmd.quiet => false
      debug                     <L> app.cmd.debug => true
        -output                 <L> app.cmd.debug-output => 
      env                       <L> app.cmd.env => false
      m                         <B>
        ore                     <L> app.cmd.more => false
        anual                   <L> app.cmd.manual => false
      raw                       <L> app.cmd.raw => false
      built-info                <L> app.cmd.built-info => false
      help                      <L> app.cmd.help => true
      keep                      <L> app.cmd.keep => false
      tree                      <L> app.cmd.tree => false
      config                    <L> app.cmd.config => 
    server.                     <B>
      a                         <B>
        uto-cert.               <B>
          enabled               <L> app.server.auto-cert.enabled => false
          domains               <L> app.server.auto-cert.domains => [example.com]
            -2nd-level          <L> app.server.auto-cert.domains-2nd-level => [aurora api home res]
        ddr                     <L> app.server.addr => 
          esses.                <B>
            rpc                 <L> app.server.addresses.rpc => 
            advertise           <L> app.server.addresses.advertise => 
      mux                       <L> app.server.mux => gin
      port                      <L> app.server.port => 9477
        s.                      <B>
          tls                   <L> app.server.ports.tls => 9479
          de                    <B>
            bug                 <L> app.server.ports.debug => 9478
            fault               <L> app.server.ports.default => 9477
      t                         <B>
        ls.certs.               <B>
          c                     <B>
            a-cert              <L> app.server.tls.certs.ca-cert => 
            ert                 <L> app.server.tls.certs.cert => 
            lient-auth          <L> app.server.tls.certs.client-auth => false
          server-cert           <L> app.server.tls.certs.server-cert => 
          key                   <L> app.server.tls.certs.key => 
          min-tls-version       <L> app.server.tls.certs.min-tls-version => 772
        ype                     <L> app.server.type => 
      statics.                  <B>
        assets.                 <B>
          url                   <L> app.server.statics.assets.url => /assets
          path                  <L> app.server.statics.assets.path => ./static/assets
        templates.              <B>
          pat                   <B>
            h                   <L> app.server.statics.templates.path => ./static/templates
            tern                <L> app.server.statics.templates.pattern => *.tmpl
          left                  <L> app.server.statics.templates.left => {{
          right                 <L> app.server.statics.templates.right => }}
    base.param                  <B>
      1                         <L> app.base.param1 => false
      2                         <L> app.base.param2 => true
      3                         <L> app.base.param3 => 8
      4                         <L> app.base.param4 => 7m2s13us
    logger.                     <B>
      level                     <L> app.logger.level => DEBUG
      format                    <L> app.logger.format => text
      target                    <L> app.logger.target => journal
      dir                       <L> app.logger.dir => /var/log/app

Matched flags:
- 1. debug (+1) Flg{'.debug'} /TILDE/ | [owner: Cmd{''}]

ACTIONS:
- ShowDebug

```

In the result above, `app.server.xxx` identifys the loading is okay. And `app.base.xxx` identifys `cond.d` also loaded ok.

## Status

- cmdr v2.1.x and higher

## LICENSE

Apache 2.0
