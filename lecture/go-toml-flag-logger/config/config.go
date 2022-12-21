package config

import (
	"os"
	"github.com/naoina/toml"
)

// config.toml
/*
type Work struct {
	Name string
	Desc string
	Excute string
	Duration int
	Args string
}

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]map[string]interface{}

	Work []Work
}

func GetConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		// toml파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
*/

// log-config.toml
type Config struct {
	Log struct {
		Level string
		Fpath string
		Msize int
		Mage int
		Mbackup int
	}
}

func GetConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		// toml파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}