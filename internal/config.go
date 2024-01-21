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
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	Version      string `yaml:"version"`
	TemplatePath string `yaml:"template-path"`
}

func CreateConfigFile(path string) error {
	// Create parent dirs
	dirPath := filepath.Dir(path)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	// Create config
	config := DefaultConfig()
	// Create yaml from config
	cfgYaml, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	// Create yaml file
	yamlFile, err := os.Create(path)
	if err != nil {
		return err
	}
	// Write yaml to file
	_, err = io.WriteString(yamlFile, string(cfgYaml))
	if err != nil {
		return err
	}
	return nil
}

func DefaultConfig() Config {
	return Config{
		Version:      "0.1.0",
		TemplatePath: ".adrtemplate.md",
	}
}
