package api

import (
	"fmt"
	"os"
)

// CreateDfmHome if not exists, else do nothing
func CreateDfmHome(basename string) {
	if !Exists(basename) {
		fmt.Printf("Creating directory %s\n", basename)
		os.MkdirAll(basename, os.ModePerm)
	}
}

// Exists returns true or false based on file existence
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
