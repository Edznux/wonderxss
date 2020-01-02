package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Current is a global variables storing the current configuration of the application.
// This is not very pretty, but it's avoiding us to send the config in every function.
// Moving to a singleton or something might be a nice refactor
// TODO: auto-refresh the configuration (watching file ? every x seconds?)
var Current Config

type Config struct {

	// The domain this application should respond
	Domain string `mapstructure:"domain"`
	// DatabaseFile represents the filename of the storage system.
	// It depends on your database type:
	// It can be a connection string (postgres, mysql)
	// or a simple filename (sqlite, json...)
	Database string `mapstructure:"database"`
	// This enable the HTTPs webserver
	// This will allow this webserver to run by itself, without any reverse proxy
	// doing the HTTPS decryption. If you are using a cloud provider of some kind,
	// with auto-managed https, it's probably best to disable it.
	StandaloneHTTPS bool `mapstructure:"standalone_https"`
	// HTTPPort is the port number for the HTTP listenner
	HTTPPOrt int `mapstructure:"http_port"`
	// HTTPsPort is the port number for the HTTPS listenner. Only used if StandaloneHTTPS is set to true
	HTTPSPOrt int `mapstructure:"https_port"`
	// Notifications represents all the configurations for the differents notification systems
	Notifications map[string]Notification `mapstructure:"notifications"`
	// Storage is the list of all the storages providers available.
	// We might add other in the future to be able to integrate more easily to existing systems.
	Storages map[string]Storage `mapstructure:"storages"`
	// JWTToken is the *SECRET* JWT Token.
	JWTToken string `mapstructure:"jwt_token"`
}

// Notifications represents the configuration for every notification systems
// Some field may be redundant depending on the notification system.
// For example, slack web hooks will only use the token field, but emails will need
// the SMTP server, user & pass, and the destination email.
type Notification struct {
	Enabled     bool   `toml:"enabled"`
	Server      string `toml:"server"`
	Name        string `toml:"name"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	Token       string `toml:"token"`
	Destination string `toml:"destination"`
}

type Storage struct {
	Adapter  string `toml:"adapter"`
	File     string `toml:"file"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Server   string `toml:"server"`
}

func Setup() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/wonderxss/")
	viper.SetConfigName("wonderxss")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&Current)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Config file loaded !")
	fmt.Println("=====================================")
	fmt.Println("Database : ", Current.Database)
	fmt.Println("HTTPPOrt : ", Current.HTTPPOrt)
	fmt.Println("=====================================")
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
