package config

import (
	"fmt"
	"strings"
	"unicode"
)

type ApiArg struct {
	Name string
	Description string
	Required bool
	Length int
	Type string
	Related []string
}

type Api struct {
	Name string
	ResponseName string
	Description string
	Async bool
	Related []string
	Args []*ApiArg
	RequiredArgs []*ApiArg
	Verb string
}

var apiCache map[string]*Api

func (c *Config) GetCache() map[string]*Api {
	if apiCache == nil {
		// read from disk?
		return make(map[string]*Api)
	}
	return apiCache
}

func (c *Config) UpdateCache(response map[string]interface{}) {
	apiCache = make(map[string]*Api)

	count := response["count"]
	apiList := response["api"].([]interface{})

	for _, node := range apiList {
		api, valid := node.(map[string]interface{})
		if !valid {
			fmt.Println("Errro, moving on")
			continue
		}
		fmt.Println("Name=", api["name"])
		apiName := api["name"].(string)
		isAsync := api["isasync"].(bool)

		idx := 0
		for _, chr := range apiName {
			if unicode.IsLower(chr) {
				idx++
			} else {
				break
			}
		}
		verb := apiName[:idx]
		fmt.Println(verb)

		var apiArgs []*ApiArg
		for _, argNode := range api["params"].([]interface{}) {
			apiArg, _ := argNode.(map[string]interface{})
			apiArgs = append(apiArgs, &ApiArg{
				Name: apiArg["name"].(string),
				Type: apiArg["type"].(string),
				Required: apiArg["required"].(bool),
			} )
		}

		apiCache[strings.ToLower(apiName)] = &Api{
			Name: apiName,
			Async: isAsync,
			Args: apiArgs,
			Verb: verb,
		}
	}
	fmt.Printf("Discovered %v APIs\n", count)
}
