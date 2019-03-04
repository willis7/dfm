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
		dfmFoldername := ".dfm"
		home, _ := homedir.Dir()
		dfmHome := filepath.Join(home, dfmFoldername)
		oldpath := args[0]
		api.CreateDfmHome(dfmHome)
		fmt.Printf("adding file %s to dfm", oldpath)

		// split path immediately following the final seperator into directory and file name component
		_, file := filepath.Split(oldpath)

		// move the file to dfmHome
		newpath := filepath.Join(dfmHome, file)
		os.Rename(oldpath, newpath)

		// create symlink back to original location
		os.Symlink(newpath, oldpath)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
