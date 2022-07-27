package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json2map()
}

//json转map
func json2map(){
	var a = "{\n  \"status\": 1,\n  \"msg\": \"success\",\n  \"data\": [{\n    \"name\": \"sskcal\",\n    \"age\": 16\n  },{\n    \"name\": \"zxq\",\n    \"age\": 18\n  }]\n}"
	result := make(map[string]interface{})
	json.Unmarshal([]byte(a),&result)
	fmt.Println(result["data"])
}