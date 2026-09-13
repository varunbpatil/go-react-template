package users

import "errors"

// Users domain errors.
//
// These are the only sentinel errors that can be returned by the users service
// and which other domains can depend on.
var (
	ErrUserNotFound = errors.New("not found")
)
