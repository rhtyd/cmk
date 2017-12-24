package command

import (
	"fmt"
	"strings"
)

func init() {
	AddCommand(&Command{
		Name:        "set",
		Help:        "Configures options for cmk",
		SubCommands: []string{"block", "display", "profile", "prompt"},
		Handle: func(r *Request) error {
			if len(r.Args) < 1 {
				fmt.Println("Please select one of the sub-commands: ", strings.Join(r.Command.SubCommands, ", "))
				return nil
			}
			subCommand := r.Args[0]
			value := strings.Join(r.Args[1:], " ")
			r.Config.UpdateGlobalConfig(subCommand, value)

			if subCommand == "prompt" && r.Shell != nil {
				r.Shell.SetPrompt(r.Config.GetPrompt())
			}
			return nil
		},
	})
}
