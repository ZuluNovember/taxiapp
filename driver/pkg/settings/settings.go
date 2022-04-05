package settings

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Address string `yaml:"address"`
		Name    string `yaml:"name"`
	} `yaml:"database"`

	Server struct {
		EndPoint      string `yaml:"endPoint"`
		RunMode       string `yaml:"runMode"`
		AcceptAddress string `yaml:"acceptAddress"`
	} `yaml:"server"`
}

var Configuration Config

func Setup() {
	f, err := os.Open("conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	decoder.Decode(&Configuration)
}
