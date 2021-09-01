package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/echodiv/simple_blog/internal/app/apiserver"
)

var (
	configPath string
)

// Vars for params parsing:
// - config-path: path to config file
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// Parse params
// Get config and decode it
// Create and run API server
func main() {
	flag.Parse()
	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
