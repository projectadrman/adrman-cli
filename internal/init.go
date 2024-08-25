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
		Use: "init",
		Example: `  adr init
  adr init -d relative/path/to/docs/adr
  adr init -d ~/absolute/path/to/docs/adr`,
		Short: "Init an ADRs directory",
		Long: `Init Any (Architectural) Decision Records (ADRs) directory:
By default is docs/adr`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("must not be args: %q", args)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			dirPath := cmd.Flag(dirFlagName).Value.String()
			initADRsDirectory(dirPath)
		},
	}
)

func initADRsDirectory(dirPath string) {
	adrDirPath := getADRsDirectory(dirPath)
	fmt.Println("Target directory path: " + adrDirPath)
	if _, err := os.Stat(adrDirPath); os.IsNotExist(err) {
		err := os.MkdirAll(adrDirPath, 0755) // only user can write
		cobra.CheckErr(err)
	}
	configPath = filepath.Join(adrDirPath, ".adrconfig.yml")
	err := createDefaultLocalConfigFile(configPath)
	cobra.CheckErr(err)
	templatePath := filepath.Join(adrDirPath, ".adrtemplate.md")
	err = createDefaultTemplateFile(templatePath)
	cobra.CheckErr(err)
	err = createFirstRecordFile(adrDirPath, templatePath)
	cobra.CheckErr(err)
}

func getADRsDirectory(dirPath string) string {
	currDirPath := getCurrentDirectory()
	fmt.Println("Current directory path: " + currDirPath)
	if dirPath != "" {
		fmt.Println("Directory path: " + dirPath)
		switch {
		case filepath.IsAbs(dirPath):
			return dirPath
		default:
			return filepath.Join(currDirPath, dirPath)
		}
	} else {
		switch {
		case strings.HasSuffix(currDirPath, "adr"):
			return currDirPath
		case strings.HasSuffix(currDirPath, "docs"):
			return filepath.Join(currDirPath, "adr")
		default:
			return filepath.Join(currDirPath, "docs", "adr")
		}
	}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	cobra.CheckErr(err)
	return filepath.ToSlash(dir)
}

func init() {
	initCmd.Flags().StringP(dirFlagName, "d", "", "path to ADRs directory (default is docs/adr)")
	rootCmd.AddCommand(initCmd)
}
