package permissions

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented  = errors.New("Browser not implemented Permissions")
	ErrNotAPermissions = errors.New("Object is not a Permissions")
)
