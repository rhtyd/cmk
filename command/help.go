package command

import (
	"fmt"

)

func init() {
	AddCommand(&Command{
		Name: "help",
		Help: "Help",
		Handle: func(r *Request) error {
			if len(r.Args) < 1 {
				PrintUsage()
				return nil
			}
			fmt.Println("FIXME: add cmd help docs")
			return nil
		},
	})
}

