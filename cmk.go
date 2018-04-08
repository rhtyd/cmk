package main

import (
	"os"

	"./config"
	"./interpretor"
)

func main() {
	args := os.Args[1:]
	cfg := config.NewConfig()
	if len(args) > 0 {
		interpretor.ExecCmd(cfg, args, nil)
	} else {
		interpretor.ExecShell(cfg)
	}
}
