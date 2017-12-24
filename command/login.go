package command

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)


func init() {
	AddCommand(&Command{
		Name: "login",
		Help: "Log in to your account",
		Handle: func(r *Request) error {
			if len(r.Args) > 0 {
				return errors.New("this does not accept any additional arguments")
			}

			validate := func(input string) error {
				if len(input) < 1 {
					return errors.New("Please enter something")
				}
				return nil
			}

			prompt := promptui.Prompt{
				Label:    "Username",
				Validate: validate,
				Default:  "",
			}

			username, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			prompt = promptui.Prompt{
				Label:    "Password",
				Validate: validate,
				Mask:     '*',
			}

			password, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			fmt.Println("Trying to log in using", username, password)

			return nil
		},
	})
}