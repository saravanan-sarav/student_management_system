package config

import (
	"encoding/json"
	"io"
	"log"
)

type Config struct {
	Env                string `json:"env"`
	Port               string `json:"port"`
	AppName            string `json:"app_name"`
	DatabaseURL        string `json:"database_url"`
	MaxDBConn          int    `json:"max_db_conn"`
	TokenSize          int    `json:"token_size"`
	AccessTokenExpiry  int    `json:"access_token_expiry"`
	RefreshTokenExpiry int    `json:"refresh_token_expiry"`
}

const (
	JSONType = "json"
	envType = "env"
)

var Conf *Config

func Parse(fileType string, file io.Reader) error {
	switch fileType {
	case JSONType:
		return ParseJson(file)
	case envType:
		return parseEnv(file)
	}
	return nil
}

func ParseJson(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		log.Println("Unable to parse json config file. Err: ", err)
		return err
	}

	Conf = &Config{}
	return json.Unmarshal(data, Conf)
}

func parseEnv(r io.Reader) error {
	log.Println("This feature will implement soon!!!", r)
	return nil
}