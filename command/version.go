package command

import "fmt"

func init() {
	AddCommand(&Command{
		Name: "version",
		Help: "Version info",
		Handle: func(r *Request) error {
			fmt.Println(r.Config.Version())
			return nil
		},
	})
}
