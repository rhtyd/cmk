// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package config

import (
	"os"
	"path"

)

type Config struct {
	Dir           string
	ConfigFile    string
	HistoryFile   string
	CacheFile     string
	LogFile       string

	Prompt        string
	ActiveProfile string
}

func NewConfig() *Config {
	return refreshConfig()
}

func refreshConfig() *Config {
	cfgDir := getHomeDirectory()
	cfg := &Config{
		Dir:           cfgDir,
		ConfigFile:    path.Join(cfgDir, "config"),
		HistoryFile:   path.Join(cfgDir, "history"),
		CacheFile:     path.Join(cfgDir, "cache"),
		LogFile:       path.Join(cfgDir, "log"),
		ActiveProfile: "local",
	}

	if _, err := os.Stat(cfg.Dir); err != nil {
		os.Mkdir(cfg.Dir, 0700)
	}

	if _, err := os.Stat(cfg.ConfigFile); err != nil {
		// FIXME: write default cfg
	} else {
		//load config?
	}

	LoadCache(cfg)

	return cfg
}