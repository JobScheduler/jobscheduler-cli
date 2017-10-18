package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	configDirectory = ".jobscheduler"
	configFile      = "config.toml"
)

type Config struct {
	ApiToken string `toml:"api_token",json:"api_token"`
}

// FatalIfErr log.Fatal if err is different than nil
func FatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfigDirectoryPath() (path string, err error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", errors.New("Missing env variable: $HOME")
	}

	path = home + "/" + configDirectory
	return path, nil
}

func GetConfigFilePath() (path string, err error) {
	directoryPath, err := GetConfigDirectoryPath()
	if err != nil {
		return "", err
	}

	path = directoryPath + "/" + configFile
	return path, nil
}

func GetAPIToken() string {
	var config Config
	// check the existence
	path, err := GetConfigFilePath()
	if err != nil {
		return ""
	}

	_, err = toml.DecodeFile(path, &config)
	if err != nil {
		return ""
	}
	return config.ApiToken
}

// SetAPIToken Set the api_token in the configuration file
func SetAPIToken(apiToken string) (err error) {
	directory, err := GetConfigDirectoryPath()
	if err != nil {
		return err
	}
	doesDirectoryExists, err := exists(directory)
	if err != nil {
		return err
	}
	if doesDirectoryExists == false {
		err = os.Mkdir(directory, 0744)
		if err != nil {
			return err
		}
	}

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}
	var config = map[string]interface{}{
		"api_token": apiToken}
	data := new(bytes.Buffer)
	err = toml.NewEncoder(data).Encode(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilePath, data.Bytes(), 0600)
	return err
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
