package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/xamust/qtimTestQuiz/internal/app/server"
	"log"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "configs-path", "configs/config.toml", "Path to config files...")
	flag.Parse()
}

func main() {
	config := server.NewConfig()
	meta, err := toml.DecodeFile(confPath, config)
	if err != nil {
		log.Fatalln(err)
	}

	if len(meta.Undecoded()) != 0 {
		log.Fatalln("Undecoded configs param: ", meta.Undecoded())
	}

	//start server...
	server := server.NewServer(config)
	if err = server.Start(); err != nil {
		log.Fatalln("Error on start:", err)
	}

}
