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
	Use:     "ls",
	Aliases: []string{"list"},
	Example: "adr list",
	Short:   "List ADRs.",
	Long:    "List Any (Architectural) Decision Records (ADRs).",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
