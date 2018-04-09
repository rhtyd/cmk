package command

import (
	"os"

	"github.com/manifoldco/promptui"
)

func init() {
	AddCommand(&Command{
		Name: "exit",
		Help: "Exits",
		Handle: func(r *Request) error {
			prompt := promptui.Prompt{
				Label:     "Do you really want to exit ([y]/n)?",
				IsConfirm: true,
			}

			if result, _ := prompt.Run(); result == "y" {
				os.Exit(0)
			}
			return nil
		},
	})
}
