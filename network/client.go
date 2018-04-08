package network

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)


func encodeRequestParams(params url.Values) string {
	if params == nil {
		return ""
	}

	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, key := range keys {
		value := params.Get(key)
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(url.QueryEscape(value))
	}
	return buf.String()
}


func MakeRequest(api string, args []string) (map[string]interface{}, error) {
	fmt.Println("running api:", api, args)

	params := make(url.Values)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) == 2 {
			params.Add(parts[0], parts[1])
		}
	}

	// Test hardcoded
	apiKey := "IgrUOA_46IVoBNzAR_Th2JbdbgIs2lMW1kGe9A80F9X0uOnfGO0Su23IqOSqbdzZW3To95PNrcdWsk60ieXYBQ"
	secretKey := "E7NRSv5d_1VhqXUHJEqvAsm7htR_V_vtPJZsCPkgPKSgkiS3sh4SOrIqMm_eWhSFoL6RHRIlxtA_viQAt7EDVA"

	params.Add("response", "json")
	params.Add("command", api)
	params.Add("apiKey", apiKey)
	encodedParams := encodeRequestParams(params)

	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(strings.Replace(strings.ToLower(encodedParams), "+", "%20", -1)))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	encodedParams = encodedParams + fmt.Sprintf("&signature=%s", url.QueryEscape(signature))

	url := fmt.Sprintf("http://orangebox.yadav.cloud/client/api?%s", encodedParams)

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