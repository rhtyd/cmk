package config

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

func getHomeDirectory() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path.Join(home, ".cmk")
}

func (c *Config) GetPrompt() string {
	return fmt.Sprintf("(%s) %s ", c.ActiveProfile, c.Prompt)
}

func (c *Config) UpdateGlobalConfig(key string, value string) {
	c.UpdateConfig("", key, value)
}

func (c *Config) UpdateConfig(namespace string, key string, value string) {
	fmt.Println("Updating for key", key, ", value=", value, ", in ns=", namespace)
	if key == "prompt" {
		c.Prompt = value
	}
}