package ports

import "errors"

// Port errors.
//
// These are *not* users domain errors.
// These are just common sentinel errors for all adapter implementations to use
// so that adapter specific details do not leak into the users service.
var (
	ErrPortSomething     = errors.New("something")
	ErrPortSomethingElse = errors.New("something else")
)
