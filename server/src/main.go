package main

import (
	"fmt"
	"os"
	"sms/constants"
)

func main() {
	res, err := os.ReadFile(constants.EnvFile)

	if err != nil {
		fmt.Println("Unable to read env file", err)
	}

	fmt.Println(string(res))

}
