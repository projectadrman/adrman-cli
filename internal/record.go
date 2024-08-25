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
	"os"
	"path/filepath"
	"text/template"
	"unicode"
)

type Record struct {
	Serial int
	Title  string
	Fields map[string]string
}

func createFirstRecordFile(dirPath string, templatePath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	templateTool, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	record := createFirstRecord()
	//record.Fields["Date"] = time.Now().Format(time.DateOnly)
	recordPath := filepath.Join(dirPath, createRecordFilename(record))
	recordFile, err := os.Create(recordPath)
	if err != nil {
		return err
	}
	err = templateTool.Execute(recordFile, record.Fields)
	if err != nil {
		return err
	}
	return nil
}

func createFirstRecord() Record {
	return Record{
		Serial: 0,
		Title:  "Record architecture decisions",
		Fields: map[string]string{
			"Date":    "foo",
			"Satus":   "Accepted",
			"Context": "We need to record the architectural decisions made on this project.",
			"Decision": `We will use Architecture Decision Records,
as described by Michael Nygard in this article:
https://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions`,
			"Consequences": "See Michael Nygard's article, linked above.",
		},
	}
}

func createRecordFilename(record Record) string {
	numberStr := fmt.Sprintf("%04d", record.Serial)
	kebabTitle := toKebabCase(record.Title)
	return fmt.Sprintf("%s-%s.md", numberStr, kebabTitle)
}

func toKebabCase(input string) string {
	var result []rune
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '-')
			}
			result = append(result, unicode.ToLower(r))
		} else if r == ' ' || r == '_' || r == '-' {
			result = append(result, '-')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
