package command

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func init() {
	AddCommand(&Command{
		Name: "shell",
		Help: "Drops into a shell",
		Handle: func(r *Request) error {
			cmd := strings.TrimSpace(strings.Join(r.Args, " "))
			if len(cmd) < 1 {
				return errors.New("no shell command provided")
			}
			out, err := exec.Command("bash", "-c", cmd).Output()
			if err == nil {
				fmt.Println(string(out))
				return nil
			}
			return errors.New("failed to execute command, " + err.Error())
		},
	})
}

