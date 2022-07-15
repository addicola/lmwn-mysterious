package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		panic(err)
	}

	whatIsIt = string(sd)

	fmt.Println(whatIsIt)
}
