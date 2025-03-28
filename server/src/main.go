package main

import (
	"fmt"
	"log"
	"os"
	"sms/config"
	"sms/constants"
	"sms/db"
	"sms/utils/router"
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

	db.Init()
	router.Init()

}
