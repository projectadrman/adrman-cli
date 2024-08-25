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
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Example: `  adr list
  adr list -d relative/path/to/docs/adr
  adr list -d ~/absolute/path/to/docs/adr`,
	Short: "Show ADRs",
	Long:  `Show Any (Architectural) Decision Records (ADRs)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("must not be args: %q", args)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		showADRs()
	},
}

func showADRs() {
	fmt.Println("ADRs:")
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		if res, _ := regexp.MatchString("\\d{4}-.+\\.md", filename); res {
			fmt.Println(filename)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
