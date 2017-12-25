package command

import (
	"fmt"

	"../network"
)

func init() {
	AddCommand(&Command{
		Name: "sync",
		Help: "Discovers and updates APIs",
		Handle: func(r *Request) error {
			fmt.Println("Discovering APIs for you")

			response, err := network.MakeRequest("listApis", []string{"listall=true"})
			if err != nil {
				return err
			}

			r.Config.UpdateCache(response)
			ConfigurePrefixCompleter(r.Completer, r.Config)
			return nil
		},
	})
}
