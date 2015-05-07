package helpers

import "os"

func FileDirExists(fp string) bool {
	if _, err := os.Stat(fp); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	} else {
		return true
	}
}
