package config

import (
	"os"
	"os/user"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

const (
	CONFIG_NAME = ".wonderxss"
)

var (
	HOME_DIR   string
	configFile string
)

type Client struct {
	Version string `toml:"version"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Token   string `toml:"token"`
}

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	HOME_DIR = usr.HomeDir
	configFile = filepath.Join(HOME_DIR, CONFIG_NAME)
}

func configFileExist() bool {
	info, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadClientConfig() (Client, error) {
	var config Client
	if !configFileExist() {
		SaveClientConfig(Client{})
	}
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return Client{}, err
	}

	return config, nil
}

func SaveClientConfig(config Client) error {
	r, err := os.OpenFile(configFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer r.Close()

	err = toml.NewEncoder(r).Encode(config)
	if err != nil {
		return err
	}
	return nil
}
