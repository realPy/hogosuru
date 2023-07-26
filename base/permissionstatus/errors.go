package permissionstatus

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented  = errors.New("Browser not implemented PermissionStatus")
	ErrNotAPermissions = errors.New("Object is not a PermissionStatus")
)
