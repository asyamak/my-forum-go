package config

import (
	"encoding/json"
	"os"
	"time"
)

const (
	// var add string
	filename      = "config/config.json"
	maxHeader     = 1 >> 20
	writeTO       = 10 * time.Second
	shutDownTO    = 3 * time.Second
	dbNameandPath = "./forum.db"
	dbDriver      = "sqlite3"
)

type Config struct {
	Port            string `json:"port"`
	MaxHeaderBytes  int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeOut time.Duration
	DbNameAndPath   string
	DbDriver        string
}

func NewConfig() (*Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}
	defer configFile.Close()
	return &Config{
		Port:            config.Port,
		MaxHeaderBytes:  maxHeader,
		ReadTimeout:     writeTO,
		WriteTimeout:    writeTO,
		ShutdownTimeOut: shutDownTO,
		DbNameAndPath:   dbNameandPath,
		DbDriver:        dbDriver,
	}, nil
}
