package interpretor

import (
	"fmt"
	"io"
	"strings"

	"../command"
	"../config"
	"github.com/mattn/go-shellwords"

	"github.com/chzyer/readline"
)

func ExecShell(cfg *config.Config) {
	shell, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m»\033[0m ", //cfg.GetPrompt(),
		HistoryFile:       cfg.HistoryFile,
		AutoComplete:      command.NewCompleter(cfg),
		InterruptPrompt:   "^C",
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
			continue
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) < 1 {
			continue
		}

		shellwords.ParseEnv = true
		parser := shellwords.NewParser()
		args, err := parser.Parse(line)
		if err != nil {
			fmt.Println("Failed to parse line:", err)
			continue
		}

		if parser.Position > 0 {
			line = fmt.Sprintf("shell %s %v", cfg.Name(), line)
			args = strings.Split(line, " ")
		}

		fmt.Println("args=", strings.Join(args, ", "), "pos=", parser.Position)

		err = ExecCmd(cfg, args, shell)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
