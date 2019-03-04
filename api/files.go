package api

import (
	"fmt"
	"os"
)

// CreateDfmHome will intitialise the dfmhome directory
func CreateDfmHome(basename string) {
	if !Exists(basename) {
		fmt.Printf("Creating directory %s\n", basename)
		os.MkdirAll(basename, os.ModePerm)
	} else {
		fmt.Printf("%s already exists\n", basename)
		os.Exit(1)
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
