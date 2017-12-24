package interpretor

import (
	"fmt"
	"io"
	"strings"

	"../command"
	"../config"

	"github.com/chzyer/readline"
)

func ExecShell(cfg *config.Config) {
	completer := readline.NewPrefixCompleter()
	command.ConfigurePrefixCompleter(completer, cfg)

	shell, err := readline.NewEx(&readline.Config{
		Prompt:            cfg.GetPrompt(),
		HistoryFile:       cfg.HistoryFile,
		AutoComplete:      completer,
		EOFPrompt:         "exit",
		VimMode:           false,
		HistorySearchFold: true,
		FuncFilterInputRune: func(r rune) (rune, bool) {
			switch r {
			case readline.CharCtrlZ:
				return r, false
			}
			return r, true
		},
	})

	if err != nil {
		panic(err)
	}
	defer shell.Close()

	cfg.PrintHeader()

	for {
		line, err := shell.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) < 1 {
			continue
		}

		err = ExecCmd(cfg, strings.Split(line, " "), shell, completer)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
