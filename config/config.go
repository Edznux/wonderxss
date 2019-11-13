package config

import "os"

const ENV_PREFIX = "WONDERXSS_"

type Config struct {

	// The domain this application should respond
	Domain string
	// DatabaseFile represents the filename of the storage system.
	// It depends on your database type:
	// It can be a connection string (postgres, mysql)
	// or a simple filename (sqlite, json...)
	Database string
	// This enable the HTTPs webserver
	// This will allow this webserver to run by itself, without any reverse proxy
	// doing the HTTPS decryption. If you are using a cloud provider of some kind,
	// with auto-managed https, it's probably best to disable it.
	StandaloneHTTPS bool
}

func Load() *Config {
	standaloneHTTPS := false
	envHTTPS := os.Getenv(ENV_PREFIX + "HTTPS")
	if envHTTPS == "true" {
		standaloneHTTPS = true
	}
	cfg := Config{
		Domain:          os.Getenv(ENV_PREFIX + "DOMAIN"),
		Database:        os.Getenv(ENV_PREFIX + "STORE"),
		StandaloneHTTPS: standaloneHTTPS,
	}
	return &cfg
}
