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
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dfm",
	Short: "CLI for managing your dotfiles",
}

// Execute is the CLI entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dfm.yml)")
	dfmFolderName := ".dfm"
	home, _ := homedir.Dir()
	dfmHome := filepath.Join(home, dfmFolderName)
	viper.SetDefault("home", dfmHome)
}

func initConfig() {
	//	DOnt forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		//	 Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//	Search config in home directory with name ".dfm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dfm")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("WARNING: no config specified, using default values")
	}
}
