package data

import "os"

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {

		return err == nil && !info.IsDir()
	}
}
