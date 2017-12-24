package command

import (
	"fmt"

	"../config"

	"github.com/chzyer/readline"
)

var apiCommand *Command

func GetAPIHandler() *Command {
	return apiCommand
}

func init() {
	apiCommand = &Command{
		Name: "api",
		Help: "Runs API",
		CustomCompleter: func(parent *readline.PrefixCompleter, cfg *config.Config) {
			apiCache := cfg.LoadCache()
			fmt.Println("Reading api cache:", apiCache)

			var testMap = map[string]string{"list": "virtualmachines", "stop": "virtualmachines"}

			for k := range testMap {
				pc := readline.PcItem(k)
				pc.SetChildren(append(pc.GetChildren(), readline.PcItem(testMap[k])))
				parent.SetChildren(append(parent.GetChildren(), pc))
			}
		},
		Handle: func(r *Request) error {
			fmt.Println("Running API: ", r.Args)
			return nil
		},
	}
	AddCommand(apiCommand)
}
