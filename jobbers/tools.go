package jobbers

import (
	"github.com/serum-errors/go-serum"
)

/*

In this file, you see how having error codes can be handy!

The RetryWhenNatch function only returns some of the errors it receives.
This demonstrates how you can handle errors by code,
and end up with a function that produces fewer error codes than
the functions it has called.

*/

// RetryWhenNatch attempts to call TheJob of a given Jobber...
// and simply retries for all natch errors.
//
// Errors:
//
//   - jobber-error-foo -- any errors regarding foo are returned immediately.
//
func RetryWhenNatch(theJobber Jobber) error {
	err := theJobber.TheJob()
	switch serum.Code(err) {
	case ErrNatch:
		return RetryWhenNatch(theJobber)
	default:
		// Error Codes -= jobber-error-natch
		return err
	}
}
