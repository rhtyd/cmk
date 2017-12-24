package command

import (
	"../config"
	"github.com/chzyer/readline"
)

type Request struct {
	Command   *Command
	Config    *config.Config
	Shell     *readline.Instance
	Completer *readline.PrefixCompleter
	Args      []string
}

func NewRequest(cmd *Command, cfg *config.Config, shell *readline.Instance, pc *readline.PrefixCompleter, args []string) *Request {
	return &Request{
		Command: cmd,
		Config: cfg,
		Shell: shell,
		Completer: pc,
		Args: args,
	}
}