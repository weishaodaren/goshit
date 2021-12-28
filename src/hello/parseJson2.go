package main

import (
	"https://github.com/mitchellh/mapstructure"
	"encoding/json"
	"fmt"
)

type MobileInfo struct {
	Resultcode string `json:"resultcode"`
}

func main() {
	jsonStr := `
		{
			"resultcode: 200
		}
	`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	var mobile MobileInfo
	err = mapstructure.WeakDecode(result, &mobile)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(mobile.Resultcode)
}
