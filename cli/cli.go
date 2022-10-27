package main

/*
In this file we demonstrate a real-world CLI application, together with Serum.

Most CLI libraries involve callbacks as part of how they handle subcommand processing
before routing control back to it, in the callback.
This can be a little tricky to wrap the analyzer around.

We use a pretty simple hack to get to a happy outcome:
we make an "AllCmds" interface, and wrap all real business logic in one of those.
This gives us a place to put Serum error annotations.

Then, we have an all-in-one-place statement of every possible error
that the whole program can produce, and it's verifiable with static analysis.
Nice!

You wouldn't have to make an "AllCmds" all in one, of course.
Maybe you want to do that per subcommand, or... whatever you want.
You can do that!  We're leaning on Serum's features around interfaces here,
not doing any special magic that's unique to CLIs nor to this particular CLI library.

*/

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/serum-errors/go-demo-app-with-serum/jobbers"
	"github.com/serum-errors/go-demo-app-with-serum/jobbers/frob"
	"github.com/serum-errors/go-demo-app-with-serum/jobbers/snoz"

	cli "github.com/jawher/mow.cli"
	"github.com/serum-errors/go-serum"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond())) // We used rand in the demo to produce errors.  Shhhh.

	app := cli.App("demodemo", "It's just a demo")
	app.Command("do", "does a thing", func(cmd *cli.Cmd) {
		var (
			jobberName = cmd.String(cli.StringOpt{
				Name:  "jobber",
				Value: "frob",
				Desc:  "Which jobber implementation to use -- frob or snoz",
			})
		)
		cmd.Action = ExitControl(CmdDo{jobberName})
	})

	app.Run(os.Args)
}

// ExitControl is a function we use to translate errors into exit codes,
// since setting exit codes requires calling a function.
//
// We have it return a `func()` because that's the callback our chosen CLI library wants.
//
// It's also a good point for us to ensure we know what all the errors are
// that can possibly be expected at the top level of our program.
// We use the AllCmds interface on the parameter to constrain this!
func ExitControl(cmd AllCmds) func() {
	return func() {
		err := cmd.Do()
		switch serum.Code(err) {
		case "":
			return
		case "jobber-error-foo":
			defer cli.Exit(12)
		case "jobber-error-natch":
			defer cli.Exit(13)
		default:
			defer cli.Exit(99)
		}
		errLogger := json.NewEncoder(os.Stderr)
		errLogger.SetIndent("", "\t")
		errLogger.Encode(err)
	}
}

type AllCmds interface {
	// This is where we check that we know every error code the whole application can emit!
	//
	// Try deleting one of the lines below naming an error code,
	// and see what happens when you try to run go-serum-analyzer without it!
	//
	// Errors:
	//
	//   - jobber-error-foo -- for errors regarding foo.
	//   - jobber-error-natch -- if a natch appears!
	Do() error
}

/*
Note that there is a feature request in go-serum-analyzer to make this
process possible on named function types, rather than just interfaces.
That'd make this possible in a few less lines :)
*/

type CmdDo struct {
	jobberName *string
}

// Errors:
//
//   - jobber-error-foo -- for errors regarding foo.
//   - jobber-error-natch -- if a natch appears!
func (x CmdDo) Do() error {
	var jobber jobbers.Jobber
	switch *x.jobberName {
	case "frob":
		jobber = frob.FrobJobber{}
	case "snoz":
		jobber = snoz.SnozJobber{}
	}
	return jobber.TheJob()
}
