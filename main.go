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
	file, err := os.Open("./testdata/config/dotfiles.toml")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	var cfg MyConfig
	if err := toml.Unmarshal(data[:count], &cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Version)
}
