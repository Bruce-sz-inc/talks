package main

import (
	"encoding/json"
	"log"
	"os"
)

const defaultConfigPath = "./golang-tokyo.json"

func main() {
	var cfgPath string

	if len(cfgPath) == 0 {
		cfgPath = defaultConfigPath
	}

	config, err := LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello,", config.Name)
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	config := &Config{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, err
}

type Config struct {
	Name string `json:"name"`
}
