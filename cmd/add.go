// Copyright © 2019 Sion Williams
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/willis7/dfm/api"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "convert a file into a dotfile managed by dfm",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if api.Exists(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid file path specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		dfmFolderName := ".dfm"
		home, _ := homedir.Dir()
		dfmHome := filepath.Join(home, dfmFolderName)
		oldPath := args[0]
		api.CreateDfmHome(dfmHome)
		fmt.Printf("adding file %s to dfm", oldPath)

		// split path immediately following the final separator into directory and file name component
		_, file := filepath.Split(oldPath)

		// TODO: check if file is a symlink and fail if it is

		// move the file to dfmHome
		newPath := filepath.Join(dfmHome, file)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("failed to move file: %s", err)
		}

		// create symlink back to original location
		err = os.Symlink(newPath, oldPath)
		if err != nil {
			fmt.Printf("failed to create symlink: %s", err)
		}
		addSymlinkRecord(newPath, oldPath, dfmHome)
	},
}

func addSymlinkRecord(newPath, oldPath, dfmHome string) {
	d1 := []byte(fmt.Sprintf("%s --> %s", newPath, oldPath))
	err := ioutil.WriteFile(filepath.Join(dfmHome, ".dfm"), d1, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
