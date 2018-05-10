package config

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type BuildConfig struct {
	Project string `json:"project"`
	Target []string `json:"targets"`
	Path string `json:"source_dir"`
}

var BuildConfigs []BuildConfig

func init() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err == nil {
		err = json.Unmarshal(config, &BuildConfigs)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}