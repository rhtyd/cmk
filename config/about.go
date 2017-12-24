package config

import "fmt"

var version = "🐵 cloudmonkey 6.0.0-alpha1"

func (c *Config) Version() string {
	return version
}

func (c *Config) PrintHeader() {
	fmt.Printf("Welcome to cmk! %s\n", version)
	fmt.Printf("Author: Rohit Yadav <rohit@yadav.cloud>\n\n")
}