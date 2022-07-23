package main

import (
	"encoding/base64"
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		panic(err)
	}

	whatIsIt = reverse(string(sd))

	fmt.Println(whatIsIt)
}
