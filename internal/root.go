/*
 * Copyright 2024 The Project ADRMAN Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:     "adr",
		Version: "0.1.0",
		Short:   "ADRMAN CLI",
		Long: `ADRMAN CLI
is a CLI tool for managing Any (Architectural) Decision Records (ADRs).`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	//cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.adrman.yml)")
}

func initConfig() {
	fmt.Println("Init config")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".adrman" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yml")
		viper.SetConfigName(".adrman")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
