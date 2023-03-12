package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
)

type config struct {
	DBpath string   `json:"db_path"`
	Addr   string   `json:"addr"`
	Music  []string `json:"music"`
}

func readConf() config {
	var conf config
	path := fmt.Sprintf("%v/yampd.json", xdg.ConfigHome)
	b, err := os.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Panic(err)
	}
	return conf
}
