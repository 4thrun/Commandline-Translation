package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
)

// Youdao returns JSON
// JSON can be parsed directly without using a struct
func YoudaoParse(data []byte) {
	var json, err = simplejson.NewJson(data)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	query, err := json.Get("query").String()
	if err != nil {
		query = ""
	}
	var phoneticStr string
	phonetic_0, err0 := json.Get("basic").Get("phonetic").String()
	phonetic_1, err1 := json.Get("basic").Get("us-phonetic").String()
	phonetic_2, err2 := json.Get("basic").Get("uk-phonetic").String()
	if err0 != nil {
		phoneticStr = ""
	} else {
		phoneticStr = fmt.Sprintf("[ %s ] ", phonetic_0)
	}
	if err1 != nil {
		phoneticStr += ""
	} else {
		phoneticStr += fmt.Sprintf("en-US[ %s ] ", phonetic_1)
	}
	if err2 != nil {
		phoneticStr += ""
	} else {
		phoneticStr += fmt.Sprintf("en-UK[ %s ] ", phonetic_2)
	}
	fmt.Printf("%s %s %s \n\n", query, phoneticStr, " ~ from fanyi.youdao.com ")
	explains, _ := json.Get("basic").Get("explains").Array()
	for _, val := range explains {
		fmt.Printf("%s %s \n", "-", val)
	}
	fmt.Print("\n")
	web, _ := json.Get("web").Array()
	for i, val := range web {
		value := val.(map[string]interface{})
		fmt.Printf(" %s %s\n", strconv.Itoa(i+1)+".", value["key"].(string))
		valuelen := len(value["value"].([]interface{}))
		valArr := make([]string, valuelen)
		for i, value := range value["value"].([]interface{}) {
			valArr[i] = value.(string)
		}
		valueStr := strings.Join(valArr, ", ")
		fmt.Printf("    %s\n", valueStr)
	}
	fmt.Println()
	fmt.Println("   --------")
	fmt.Println()
}
