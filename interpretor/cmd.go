package interpretor

import (
	"fmt"

	"../command"
	"../config"
	"github.com/chzyer/readline"
)

func ExecCmd(cfg *config.Config, args []string, shell *readline.Instance, pc *readline.PrefixCompleter) error {
	fmt.Println("[debug] executing line =", args)

	if len(args) < 1 {
		return nil
	}

	cmd := command.FindCommand(args[0])
	if cmd != nil {
		return cmd.Handle(command.NewRequest(cmd, cfg, shell, pc, args[1:]))
	}

	apiHandler := command.GetAPIHandler()
	return apiHandler.Handle(command.NewRequest(apiHandler, cfg, shell, pc, args))
}

