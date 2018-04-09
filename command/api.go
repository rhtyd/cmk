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
		Help: "Runs a provided API",
		Handle: func(r *Request) error {
			if len(r.Args) == 0 {
				return errors.New("please provide an API to execute")
			}

			apiName := strings.ToLower(r.Args[0])
			apiArgs := r.Args[1:]
			if r.Config.GetCache()[apiName] == nil && len(r.Args) > 1 {
				apiName = strings.ToLower(strings.Join(r.Args[:2], ""))
				apiArgs = r.Args[2:]
			}

			api := r.Config.GetCache()[apiName]
			if api == nil {
				return errors.New("unknown or unauthorized API: " + apiName)
			}

			fmt.Println("Running API: ", api.Name, apiArgs)
			b, _ := network.MakeRequest(api.Name, apiArgs)
			response, _ := json.MarshalIndent(b, "", "  ")

			// Implement various output formats
			fmt.Println(string(response))
			return nil
		},
	}
	AddCommand(apiCommand)
}
