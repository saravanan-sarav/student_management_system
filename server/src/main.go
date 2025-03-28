package main

import (
	"fmt"
	"log"
	"os"
	"sms/config"
	"sms/constants"
	"sms/db"
	"sms/utils/token"
)

func main() {
	res, err := os.Open(constants.EnvFile)
	if err != nil {
		fmt.Println("Unable to read env file", err)
	}

	err = config.Parse(config.JSONType, res)
	if err != nil {
		log.Fatal("unable to parse json", err)
	}

	fmt.Println(config.Conf)

	db.Init()

	fmt.Println(token.GetAuthCode(config.Conf.TokenSize))

}
