package terror

import (
	"github.com/serum-errors/go-demo-app-with-serum/jobbers"
	"github.com/serum-errors/go-serum"
)

// MinionFunction just returns an error
// Errors:
//
//    - jobber-error-natch -- always
func MinionFunction() error {
	return jobbers.NewErrNatch(-1)
}

// InducePanic will cause go-serum to panic by reaching the unreachable
// Errors:
//
//    - jobber-error-foo -- if our minions aren't natchy
//    - jobber-error-natch -- our minions are natchy
func InducePanic() error {
	if err := MinionFunction(); err != nil {
		if err, ok := err.(serum.ErrorInterface); ok {
			return err
		}
		return jobbers.NewErrFoo("is it okay? %w", err)
	}
	return nil
}
