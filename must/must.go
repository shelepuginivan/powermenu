// Package must provides utility function Must.
package must

import "log"

// Must is a helper wrapper function for calls returning (any, error). If error
// is not nil, [log.Fatal] is called.
func Must[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}
