package states

import "github.com/pkg/errors"

// AlreadyExistsErr is throw when you try to register two identical states
var AlreadyExistsErr = errors.Errorf("this state already exists")
