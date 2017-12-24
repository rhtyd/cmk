package command

import (
	"fmt"

	"../config"
	"github.com/chzyer/readline"
)

type Command struct {
	Name        string
	Help        string
	SubCommands []string
	CustomCompleter func(*readline.PrefixCompleter, *config.Config)
	Handle      func(*Request) error
}

var commands []*Command
var commandMap map[string]*Command

func FindCommand(name string) *Command {
	return commandMap[name]
}

func AddCommand(cmd *Command) {
	commands = append(commands, cmd)
	if commandMap == nil {
		commandMap = make(map[string]*Command)
	}
	commandMap[cmd.Name] = cmd
}

func ConfigurePrefixCompleter(parent *readline.PrefixCompleter, cfg *config.Config) {
	parent.SetChildren(nil)
	for _, cmd := range commands {
		pc := readline.PcItem(cmd.Name)
		parent.SetChildren(append(parent.GetChildren(), pc))
		if cmd.SubCommands != nil {
			for _, name := range cmd.SubCommands {
				spc := readline.PcItem(name)
				pc.SetChildren(append(pc.GetChildren(), spc))
			}
		}
		if cmd.CustomCompleter != nil {
			cmd.CustomCompleter(parent, cfg)
		}
	}
}

func PrintUsage() {
	commandHelp := ""
	for _, cmd := range commands {
		commandHelp += fmt.Sprintf("%s\t\t%s\n", cmd.Name, cmd.Help)
	}
	fmt.Printf(`usage: cmk [options] [commands]

Command Line Interface for Apache CloudStack

default commands:
%s

Try cmk [help]`, commandHelp)
}
