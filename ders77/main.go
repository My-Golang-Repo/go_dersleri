package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
		{
			"data":{
				"object":"card",
				"id":"card_42343424234234",
				"first_name":"Sadagat",
				"last_name":"Asgarov",
				"balance":"45.233"
			}
		}
	`

	var m map[string]map[string]interface{}

	if err:=json.Unmarshal([]byte(jsonStr),&m); err != nil {
		panic(err)
	} 
	fmt.Println(m["data"]["balance"])

	b, err := json.Marshal(m)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(b))


}
