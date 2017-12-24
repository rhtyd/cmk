package command

import "fmt"

func init() {
	AddCommand(&Command{
		Name: "sync",
		Help: "Discovers and updates APIs",
		Handle: func(r *Request) error {
			fmt.Println("Discovering APIs for you")

			ConfigurePrefixCompleter(r.Completer, r.Config)
			return nil
		},
	})
}
