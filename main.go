package main

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

type MyConfig struct {
	Version string
}

func main() {
	data, err := os.ReadFile("./testdata/config/dotfiles.toml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg MyConfig
	if err := toml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Version)
}
