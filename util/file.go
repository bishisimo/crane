// Describe:
package util

import (
	"os"
)

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func IsRegularFile(filePath string) bool {
	stat, err := os.Lstat(filePath)
	if err != nil {
		return false
	}
	return stat.Mode().Type().IsRegular()
}
