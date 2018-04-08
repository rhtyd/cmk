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
			response, err := network.MakeRequest("listApis", []string{"listall=true"})
			if err != nil {
				return err
			}
			fmt.Printf("Discovered %v APIs\n", r.Config.UpdateCache(response))
			r.Config.SaveCache(response)
			return nil
		},
	})
}
