package command

import (
	"sort"
	"strings"
	"unicode"

	"../config"
	"github.com/chzyer/readline/runes"
)

type CliCompleter struct{
	Config *config.Config
}

var completer *CliCompleter


func TrimSpaceLeft(in []rune) []rune {
	firstIndex := len(in)
	for i, r := range in {
		if unicode.IsSpace(r) == false {
			firstIndex = i
			break
		}
	}
	return in[firstIndex:]
}

func doInternal(line []rune, pos int, lineLen int, argName []rune) (newLine [][]rune, offset int) {
	//fmt.Println("do internal inline=", string(line), ", len=", lineLen, ", argname=", string(argName) )
	offset = lineLen
	if lineLen >= len(argName) {
		if runes.HasPrefix(line, argName) {
			if lineLen == len(argName) {
				newLine = append(newLine, []rune{' '})
			} else {
				newLine = append(newLine, argName)
			}
			offset = offset - len(argName) - 1
		}
	} else {
		if runes.HasPrefix(argName, line) {
			newLine = append(newLine, argName[offset:])
		}
	}
	return
}


func (t *CliCompleter) Do(line []rune, pos int) (newLine [][]rune, offset int) {

	line = TrimSpaceLeft(line[:pos])
	lineLen := len(line)

	apiCache := t.Config.GetCache()
	apiMap := make(map[string][]*config.Api)
	for api := range apiCache {
		verb := apiCache[api].Verb
		apiMap[verb] = append(apiMap[verb], apiCache[api])
	}

	for _, cmd := range AllCommands() {
		verb := cmd.Name
		if cmd.SubCommands != nil && len(cmd.SubCommands) > 0 {
			for _, scmd := range cmd.SubCommands {
				dummyApi := &config.Api{
					Name: scmd,
					Verb: verb,
				}
				apiMap[verb] = append(apiMap[verb], dummyApi)
			}
		} else {
			dummyApi := &config.Api{
				Name: "",
				Verb: verb,
			}
			apiMap[verb] = append(apiMap[verb], dummyApi)
		}
	}

	var verbs []string
	for verb := range apiMap {
		verbs = append(verbs, verb)
		sort.Slice(apiMap[verb], func(i, j int) bool {
			return apiMap[verb][i].Name < apiMap[verb][j].Name
		})
	}
	sort.Strings(verbs)

	var verbsFound []string
	for _, verb := range verbs {
		search := verb + " "
		if !runes.HasPrefix(line, []rune(search)) {
			sLine, sOffset := doInternal(line, pos, lineLen, []rune(search))
			newLine = append(newLine, sLine...)
			offset = sOffset
		} else {
			verbsFound = append(verbsFound, verb)
		}
	}

	for _, verbFound := range verbsFound {
		search := verbFound + " "

		nLine := TrimSpaceLeft(line[len(search):])
		offset = lineLen - len(verbFound)

		for _, api := range apiMap[verbFound] {
			resource := strings.TrimPrefix(strings.ToLower(api.Name), verbFound)
			search = resource + " "

			if runes.HasPrefix(nLine, []rune(search)) {
				// FIXME: handle params to API here with = stuff
			} else {
				sLine, _ := doInternal(nLine, pos, len(nLine), []rune(search))
				newLine = append(newLine, sLine...)
			}
		}
	}

	return
}

func NewCompleter(cfg *config.Config) *CliCompleter {
	completer = &CliCompleter{
		Config: cfg,
	}
	return completer
}