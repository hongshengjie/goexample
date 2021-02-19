package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v, _ := json.Marshal(nil)
	fmt.Println(string(v))
	var f []string
	json.Unmarshal(v, &v)
	fmt.Println(f)
}
