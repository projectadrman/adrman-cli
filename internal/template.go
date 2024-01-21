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
	"io"
	"os"
	"path/filepath"
)

func CreateTemplateFile(path string) error {
	// Create parent dirs
	dirPath := filepath.Dir(path)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	// Create template
	template := DefaultTemplate()
	// Create template file
	tmpFile, err := os.Create(path)
	if err != nil {
		return err
	}
	// Write template to file
	_, err = io.WriteString(tmpFile, template)
	if err != nil {
		return err
	}
	return nil
}

func DefaultTemplate() string {
	return `# {{name}}

Date: {{date}}

## Status

Your status...

## Context

Your context...

## Decision

Your decision...

## Consequences

Your consequences...`
}
