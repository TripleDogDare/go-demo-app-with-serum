package jobbers

// This interface does jobbery.
// It's a plugin system -- there are several implementations.
//
// Implementations of it can be found in subpackages.
// Check out "frob" and "snoz" -- both provide a Jobber implementation.
type Jobber interface {
	// TheJob does the job.
	//
	// Errors:
	//
	//   - jobber-error-foo -- for errors regarding foo.
	//   - jobber-error-natch -- if a natch appears!
	TheJob() error
}
