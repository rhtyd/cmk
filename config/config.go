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
		Prompt:        "🐵 >",
	}

	if _, err := os.Stat(cfg.Dir); err != nil {
		os.Mkdir(cfg.Dir, 0700)
	}

	if _, err := os.Stat(cfg.ConfigFile); err != nil {
		// FIXME: write default cfg
	} else {
		//load config?
	}

	return cfg
}