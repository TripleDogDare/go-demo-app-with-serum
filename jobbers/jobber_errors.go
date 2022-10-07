package jobbers

import (
	"strconv"

	"github.com/serum-errors/go-serum"
)

// You don't have to make constants for the error codes when using Serum...
// but it's often a good convention.
const (
	ErrFoo   = "jobber-error-foo"
	ErrNatch = "jobber-error-natch"
)

// NewErrFoo makes a foo error, with a formatting string.
//
// Errors:
//
//    - jobber-error-foo -- this is a constructor for it!
func NewErrFoo(fmtPattern string, args ...interface{}) error {
	// This constructor is a bit boring!  Maybe you didn't even need this one, in fact.
	// You could also use an Errorf constructor like this, inline, wherever you want.
	return serum.Errorf(ErrFoo, fmtPattern, args...)
}

// NewErrNatch makes a natch error, which records an int for the flavor.
// (We just made up "flavors".  This is a demo.  It's silly, okay?)
//
// Errors:
//
//    - jobber-error-natch -- this is a constructor for it!
func NewErrNatch(flavor int) error {
	return serum.Error(
		ErrNatch,
		serum.WithMessageTemplate("flavor {{flavor}}"),
		serum.WithDetail("flavor", strconv.Itoa(flavor)),
	)
}
