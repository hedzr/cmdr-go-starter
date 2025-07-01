package cmd

import (
	"context"
	"fmt"

	"github.com/hedzr/cmdr/v2/cli"
	"github.com/hedzr/is"
	logz "github.com/hedzr/logg/slog"
)

// var lastMultiCmd *multiCmd

type multiCmd struct {
	root root
}

func (s *multiCmd) Add(app cli.App) {
	// lastMultiCmd = s

	// just for debugging, removing this 'if' branch is safe.
	if is.DebuggerAttached() {
		logz.SetLevel(logz.TraceLevel)
		// app.WithOpts(cli.WithArgs(os.Args[0], "~~tree"))
		logz.Trace(fmt.Sprintf("multiCmd.root.ptr = %p, .val = %+v\n    multiCmd.A.ptr = %p\n    multiCmd.A.D.ptr = %p\n    multiCmd.A.D.E.ptr = %p",
			&s.root, s.root,
			&s.root.A,
			&s.root.A.D,
			&s.root.A.E,
		))
		fmt.Printf("    >> A.Fa3 (positional): %p\n", &s.root.A.Fa3)
	}

	// add 'multi' subcmd, and add more subcmds and flags from `root` struct
	app.Cmd("multi", "mu", "").
		Description("multi-level test and imported form struct").
		// Group("Test").
		TailPlaceHolders("[text1, text2, ...]").
		FromStruct(&s.root).
		OnAction(s.Action).
		With(func(b cli.CommandBuilder) {
			// b.FromStruct(&root{})
		})
}

func (s *multiCmd) postAction(ctx context.Context, cmd cli.Cmd, args []string) (err error) {
	logz.Println(fmt.Sprintf("multiCmd.root: .ptr = %p, .val = %+v\n    multiCmd.A.ptr = %p\n    multiCmd.A.D.ptr = %p\n    multiCmd.A.D.E.ptr = %p",
		&s.root, s.root,
		&s.root.A,
		&s.root.A.D,
		&s.root.A.E,
	))
	logz.OK("postAction done", "s.root", s.root)
	_, _, _ = ctx, cmd, args
	return
}

func (s *multiCmd) Action(ctx context.Context, cmd cli.Cmd, args []string) (err error) {
	logz.Println(".   - multiCmd.Action() invoked.", "cmd", cmd, "args", args)
	_, err = cmd.App().DoBuiltinAction(ctx, cli.ActionDefault, anyArrayToAnyArray(args)...)
	fmt.Printf("root: %+v\n", s.root)
	_ = s.postAction(ctx, cmd, args)
	return
}

type root struct {
	b   bool // unexported values ignored
	Int int  `cmdr:"-"` // ignored
	A   `title:"a-cmd" shorts:"a,a1,a2" alias:"a1-cmd,a2-cmd" desc:"A command for demo" required:"true"`
	B   `env:"B"`
	C
	F1 int
	F2 string
}

type A struct {
	D
	Fa1 int
	Fa2 string
	Fa3 []string `cmdr:"positional"`
}
type B struct {
	F2 int
	F3 string
}
type C struct {
	F3 bool
	F4 string
}
type D struct {
	E
	FromNowOn F
	F3        bool
	F4        string
}
type E struct {
	F3 bool `title:"f3" shorts:"ff" alias:"f3ff" desc:"A flag for demo" required:"true"`
	F4 string
}
type F struct {
	F5    uint
	F6    byte
	Files []string `cmdr:"positional"`
}

// a --f1 1 --f2 str
// --a.f1 1 --a.f2 str

func (A) With(cb cli.CommandBuilder) {
	// customize for A command, for instance: fb.ExtraShorts("ff")
	logz.Info(".   - A.With() invoked.", "cmdbuilder", cb)
}
func (A) F1With(fb cli.FlagBuilder) {
	// customize for A.F1 flag, for instance: fb.ExtraShorts("ff")
	logz.Info(".   - A.F1With() invoked.", "flgbuilder", fb)
}

// Action method will be called if end-user type subcmd for it (like `app a d e --f3`).
func (s E) Action(ctx context.Context, cmd cli.Cmd, args []string) (err error) {
	logz.Info(".   - E.Action() invoked.", "cmd", cmd, "args", args)
	_, err = cmd.App().DoBuiltinAction(ctx, cli.ActionDefault, anyArrayToAnyArray(args)...)
	fmt.Printf("E: %+v\n", s)

	// if lastMultiCmd != nil {
	// 	fmt.Printf("D: %+v\n", lastMultiCmd.root.A.D)
	// 	fmt.Printf("A: %+v\n", lastMultiCmd.root.A)
	// 	fmt.Printf("A.Fa3 (positional): %p\n", &lastMultiCmd.root.A.Fa3)
	// 	_ = lastMultiCmd.postAction(ctx, cmd, args)
	// }
	return
}

// Action method will be called if end-user type subcmd for it (like `app a d f --f5=7`).
func (s F) Action(ctx context.Context, cmd cli.Cmd, args []string) (err error) {
	(&s).Inc()
	logz.Info(".   - F.Action() invoked.", "cmd", cmd, "args", args, "F5", s.F5)
	_, err = cmd.App().DoBuiltinAction(ctx, cli.ActionDefault, anyArrayToAnyArray(args)...)
	fmt.Printf("F: %+v\n", s)

	// if lastMultiCmd != nil {
	// 	fmt.Printf("D: %+v\n", lastMultiCmd.root.A.D)
	// 	fmt.Printf("A: %+v\n", lastMultiCmd.root.A)
	// 	_ = lastMultiCmd.postAction(ctx, cmd, args)
	// }
	return
}

func (s *F) Inc() {
	s.F5++
}

func anyArrayToAnyArray[T any](args []T) (ret []any) {
	ret = make([]any, 0, len(args))
	for _, it := range args {
		ret = append(ret, it)
	}
	return
}
