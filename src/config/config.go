package config

import (
	"os"
	"encoding/json"
	"log"
	"github.com/xsank/EasyProxy/src/structure"
)

type Config struct {
	Service      string `json:"service"`
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	WebPort      uint16 `json:"webport"`
	Strategy     string `json:"strategy"`
	Heartbeat    int `json:"heartbeat"`
	MaxProcessor int `json:"maxprocessor"`
	Backends     []structure.Backend `json:"backends"`
}

func Load(filename string) (*Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		log.Println("load config failed:", err)
	} else {
		buff := make([]byte, 1024)
		end, _ := file.Read(buff)
		err = json.Unmarshal(buff[:end], &config)
		if err != nil {
			log.Println("decode json config failed:", err)
		}
	}
	log.Println("success load config file:", filename)
	return &config, err
}

