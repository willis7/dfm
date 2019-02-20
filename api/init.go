package api

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// FolderSetup will intitialise the dfmhome directory
func FolderSetup(cmd *cobra.Command, args []string) {
	dfmFoldername := ".dfm"
	home, _ := homedir.Dir()
	dfmHome := filepath.Join(home, dfmFoldername)

	if !exists(dfmHome) {
		fmt.Printf("Creating %s directory in %s\n", dfmFoldername, home)
		os.MkdirAll(dfmHome, os.ModePerm)
	} else {
		fmt.Printf("%s already exists in %s\n", dfmFoldername, home)
		os.Exit(1)
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
