package navigationpreloadmanager

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented               = errors.New("Browser not implemented NavigationPreloadManager")
	ErrNotANavigationPreloadManager = errors.New("Object is not a NavigationPreloadManager")
)
