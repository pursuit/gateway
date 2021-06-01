package config

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Prefix string   `yaml:"prefix"`
	Routes []string `yaml:"routes"`
}

func Instance(base string) map[string]string {
	urls := make(map[string]string)

	files, err := ioutil.ReadDir(base)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var config Config
		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		yamlFile, _ := ioutil.ReadFile(base + "/" + file.Name())
		if err := yaml.Unmarshal(yamlFile, &config); err != nil {
			panic(err)
		}

		key := file.Name()[:len(file.Name())-5]
		target := os.Getenv("URL_" + strings.ToUpper(key))

		for _, route := range config.Routes {
			urls[config.Prefix+route] = target
			if route == "/" {
				urls[config.Prefix] = target
			} else {
				urls[config.Prefix+route+"/"] = target
			}
		}
	}

	return urls
}
