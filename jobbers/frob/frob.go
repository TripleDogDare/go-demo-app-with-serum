package frob

import (
	"fmt"
	"math/rand"

	"github.com/serum-errors/go-demo-app-with-serum/jobbers"
)

// You'll notice that the snoz Jobber is a little more reliably than
// the Jobber implementation in the frob package -- the snoz Jobber doesn't
// return any errors regarding foo (but it can still encounter natch!).
type FrobJobber struct{}

// This assignment asserts that SnozJobber implements jobbers.Jobber.
//
// (This is a pretty common golang pattern to make sure interfaces are checked...
// You need this for the Serum error codes on that interface to be checked too!)
var _ jobbers.Jobber = FrobJobber{}

// TheJob is the method required by the Jobber interface.
//
// Errors:
//
//   - jobber-error-foo -- for errors regarding foo.
//   - jobber-error-natch -- if a natch appears!
func (FrobJobber) TheJob() error {
	n := rand.Int()
	if n%3 == 0 {
		otherError := fmt.Errorf("some other error, to demo cause chaining")
		return jobbers.NewErrFoo("the foo frobnosticated: %w", otherError)
	}
	if n%4 == 0 {
		return jobbers.NewErrNatch(n)
	}
	return nil
}
