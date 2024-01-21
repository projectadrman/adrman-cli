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
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirFlagName = "directory"
	initCmd     = &cobra.Command{
		Use:        "init",
		Aliases:    []string{"initialize", "initialise"},
		SuggestFor: []string{"create", "start"},
		Example: `  adr init
  adr init -d relative/path/to/docs/adr
  adr init -d ~/absolute/path/to/docs/adr`,
		Short: "Initialize an ADRs directory.",
		Long: `Initialize an Any (Architectural) Decision Records (ADRs) directory:
By default is docs/adr`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("must not be args: %q", args)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			currentDirPath := getCurrentDirPath()
			fmt.Println("Current directory path: " + currentDirPath)
			argDirPath := cmd.Flag(dirFlagName).Value.String()
			fmt.Println("Directory path by argument: " + argDirPath)
			var targetDirPath string
			if argDirPath != "" {
				switch {
				case filepath.IsAbs(argDirPath):
					targetDirPath = argDirPath
				default:
					targetDirPath = filepath.Join(currentDirPath, argDirPath)
				}
			} else {
				switch {
				case strings.HasSuffix(currentDirPath, "adr"):
					targetDirPath = currentDirPath
				case strings.HasSuffix(currentDirPath, "docs"):
					targetDirPath = filepath.Join(currentDirPath, "adr")
				default:
					targetDirPath = filepath.Join(currentDirPath, "docs", "adr")
				}
			}
			fmt.Println("Target directory path: " + targetDirPath)
			if _, err := os.Stat(targetDirPath); os.IsNotExist(err) {
				err := os.MkdirAll(targetDirPath, os.ModePerm)
				cobra.CheckErr(err)
			}
			cfgFile = filepath.Join(targetDirPath, ".adrman.yml")
			err := CreateConfigFile(cfgFile)
			cobra.CheckErr(err)
			tmpFile := filepath.Join(targetDirPath, ".adrtemplate.md")
			err = CreateTemplateFile(tmpFile)
			cobra.CheckErr(err)
		},
	}
)

func getCurrentDirPath() string {
	dir, err := os.Getwd()
	cobra.CheckErr(err)
	return filepath.ToSlash(dir)
}

func init() {
	initCmd.Flags().StringP(dirFlagName, "d", "", "path to ADRs directory (default is docs/adr)")
	rootCmd.AddCommand(initCmd)
}
