package command

import (
	"fmt"

)

type Command struct {
	Name        string
	Help        string
	SubCommands []string
	CustomCompleter func(input string, position int)
	Handle      func(*Request) error
}

var commands []*Command
var commandMap map[string]*Command

func FindCommand(name string) *Command {
	return commandMap[name]
}

func AllCommands() []*Command {
	return commands
}

func AddCommand(cmd *Command) {
	commands = append(commands, cmd)
	if commandMap == nil {
		commandMap = make(map[string]*Command)
	}
	commandMap[cmd.Name] = cmd
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
