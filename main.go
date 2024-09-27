package main

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

type Config struct {
	Meta
	Dotfiles []Dotfile `toml:"dotfiles"`
}

type Meta struct {
	Version int `toml:"version"`
}

type Dotfile struct {
	IsEnabled bool   `toml:"enable"`
	Source    string `toml:"source"`
	Target    string `toml:"target"`
}

func main() {
	data, err := os.ReadFile("./testdata/config/dotfiles.toml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	for _, v := range cfg.Dotfiles {
		if v.IsEnabled == false {
			continue
		}

		if err := link(v.Source, v.Target); err != nil {
			log.Fatal(err)
		}
	}
}
