package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeRequest(api string, args []string) (map[string]interface{}, error) {
	fmt.Println("running api:", api, args)
	url := fmt.Sprintf("http://localhost:8096/client/api?command=%s&response=json&%s", api, strings.Join(args, "&"))
	fmt.Println("requesting: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}
	_ = json.Unmarshal([]byte(body), &data)


	for k := range data {
		if strings.HasSuffix(k, "response") {
			return data[k].(map[string]interface{}), nil
		}
	}
	return nil, errors.New("failed to decode response")
}