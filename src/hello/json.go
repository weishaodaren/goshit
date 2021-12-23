package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	res := Result{Code: 200, Message: "success"}

	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data: ", string(jsons))

	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("res: ", res2)
}
