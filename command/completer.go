package interpretor

import (
	"../command"
)

type CliCompleter struct{
}

var completer *CliCompleter
var commands []*command.Command

func (t *CliCompleter) Do(input []rune, pos int) ([][]rune, int) {

	var commands [][]rune
	for _, cmd := range command.AllCommands() {
		commands = append(commands, []rune(cmd.Name))
	}

	return commands, 0

	/*
	cmdName := strings.Split(string(input), " ")[0]
	cmd := command.FindCommand(cmdName)
	if cmd == nil {
		return [][]rune{}, 0
	}

	line := string(input)

	if cmd.CustomCompleter != nil {
		cmd.CustomCompleter(line, pos)
	}



	fmt.Println("\nline=", string(line), "int=", pos)
	fmt.Println("")
//	return [][]rune{[]rune("list")}, 0
	*/
}

func NewCompleter() *CliCompleter {
	completer = &CliCompleter{}
	return completer
}