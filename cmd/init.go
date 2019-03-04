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
	"github.com/spf13/cobra"
	"github.com/willis7/dfm/api"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new dfm folder",
	Long: `intitialise a new dfm folder. Creates a .dfm directory
		in the $HOME directory`,
	Args: cobra.NoArgs,
	Run:  api.FolderSetup,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
