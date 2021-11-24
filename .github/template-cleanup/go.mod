module github.com/%REPOSITORY%

go 1.16

//replace github.com/hedzr/log => ../10.log

//replace github.com/hedzr/logex => ../15.logex

//replace github.com/hedzr/cmdr => ../50.cmdr

//replace github.com/hedzr/cmdr-addons => ../53.cmdr-addons

require (
	github.com/hedzr/cmdr v1.9.3
	github.com/hedzr/cmdr-addons v1.9.3
	github.com/hedzr/log v1.3.22
	github.com/hedzr/logex v1.3.22
	gopkg.in/hedzr/errors.v2 v2.1.5
)
