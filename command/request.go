package command

import (
	"../config"
	"github.com/chzyer/readline"
)

type Request struct {
	Command   *Command
	Config    *config.Config
	Shell     *readline.Instance
	Args      []string
}

func NewRequest(cmd *Command, cfg *config.Config, shell *readline.Instance, args []string) *Request {
	return &Request{
		Command: cmd,
		Config: cfg,
		Shell: shell,
		Args: args,
	}
}