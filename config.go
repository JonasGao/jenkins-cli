package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Domain   string `yaml:"domain"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var (
	readingConfig sync.Once
	setting       Config
)

func readConfig() (Config, error) {
	readingConfig.Do(func() {
		dirname, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		config, err := ioutil.ReadFile(dirname + "/.jk.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(config, &setting)
		if err != nil {
			panic(err)
		}
	})
	return setting, nil
}
