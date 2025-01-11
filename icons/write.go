package icons

import "github.com/shelepuginivan/fsutil"

// ExistsOrWrite creates a file and writes data to it if the given path does
// not exist.
func ExistsOrWrite(path string, data []byte) error {
	if fsutil.FileExists(path) {
		return nil
	}

	return fsutil.WriteAll(path, data)
}
