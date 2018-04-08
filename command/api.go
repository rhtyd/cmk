package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"../network"
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
			apiCache := cfg.GetCache()
			apiMap := make(map[string][]*config.Api)
			for api := range apiCache {
				apiMap[apiCache[api].Verb] = append(apiMap[apiCache[api].Verb], apiCache[api])
			}

			for verb := range apiMap {
				pc := readline.PcItem(verb)
				for _, api := range apiMap[verb] {
					argName := strings.ToLower(strings.TrimPrefix(api.Name, verb))
					argPc := readline.PcItem(argName)
					//FIXME: dynamic args?
					for _, param := range api.Args {
						paramName := param.Name + "="
						paramPc := readline.PcItem(paramName)
						argPc.SetChildren(append(argPc.GetChildren(), paramPc))
					}
					pc.SetChildren(append(pc.GetChildren(), argPc))
				}
				parent.SetChildren(append(parent.GetChildren(), pc))
			}
		},
		Handle: func(r *Request) error {
			if len(r.Args) < 2 {
				return errors.New("please provide full API name")
			}
			apiName := strings.ToLower(strings.Join(r.Args[:2], ""))
			apiArgs := r.Args[2:]

			api := r.Config.GetCache()[apiName]
			if api == nil {
				//prompt?
				return errors.New("unknown or unauthorized API")
			}

			fmt.Println("Running API: ", api.Name, apiArgs)
			b, _ := network.MakeRequest(api.Name, apiArgs)
			response, _ := json.MarshalIndent(b, "", "  ")
			fmt.Println(string(response))
			return nil
		},
	}
	AddCommand(apiCommand)
}
