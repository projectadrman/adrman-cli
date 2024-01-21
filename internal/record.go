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
	"strings"
	"time"
)

func CreateFirstRecord(path string) error {
	// Create parent dirs
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	// Create template
	filename := filepath.Join(path, FirstRecordName())
	// Create template file
	recordFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	// Creat content
	content := FirstRecordContent()
	now := time.Now()
	date := now.Format(time.DateOnly)
	total := strings.Replace(content, "{{date}}", date, 1)
	// Write template to file
	_, err = io.WriteString(recordFile, total)
	if err != nil {
		return err
	}
	return nil
}

func FirstRecordName() string {
	return "0001-record-architecture-decisions.md"
}

func FirstRecordContent() string {
	return `
# 1. Record architecture decisions

Date: {{date}}

## Status

Accepted

## Context

We need to record the architectural decisions made on this project.

## Decision

We will use Architecture Decision Records,
as described by Michael Nygard in this article:
https://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions

## Consequences

See Michael Nygard's article, linked above.`
}
