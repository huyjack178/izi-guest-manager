package main

import (
	"flag"
	"server"
	"util/load-config"
	"util/logs"
)

var (
	flConfigFile = flag.String("config-file", "src/config-default.json", "Load config from file")

	log = logs.New("server")
)

func main() {
	flag.Parse()

	var cfg server.Config

	err := loadConfig.FromFileAndEnv(&cfg, *flConfigFile)

	if err != nil {
		log.Fatalln("Error loading config: ", err)
	}

	server.Start(cfg)
}
