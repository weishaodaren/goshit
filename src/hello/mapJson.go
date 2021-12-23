package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"usename": "Tom",
		"age":     "30",
		"hobby":   []string{"读书", "爬山"},
	}
	fmt.Println("map data:", res)

	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal err:", errs)
	}
	fmt.Println(string(jsons))

	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil{
		fmt.Println("json unmarshal error:", errs)
	}

	fmt.Println(res2)
}
