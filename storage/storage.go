package storage

import (
	"fmt"
	"log"

	"github.com/edznux/wonder-xss/config"
	"github.com/edznux/wonder-xss/storage/models"
	"github.com/edznux/wonder-xss/storage/sqlite"
)

var backend map[string]Storage
var currentStorage Storage

// Storage represents the interface for all storage adapters.
// It should be easy to plug in new storage type like SQL Databases or plain files.
type Storage interface {
	// First time setup (create tables, create file...)
	Setup() error
	// Open the connection, open the file...
	Init(config.Config) error
	// CRUD interface
	// Create
	CreatePayload(models.Payload) (models.Payload, error)
	CreateUser(models.User) (models.User, error)
	CreateAlias(models.Alias) (models.Alias, error)
	CreateLoot(loot models.Loot, payloadIDOrAlias string) (models.Loot, error)

	// Read
	GetPayloads() ([]models.Payload, error)
	GetPayload(id string) (models.Payload, error)
	GetAliases() ([]models.Alias, error)
	GetAlias(id string) (models.Alias, error)
	GetLoot(id string) (models.Loot, error)
	GetLoots() ([]models.Loot, error)
	GetPayloadByAlias(short string) (models.Payload, error)
	GetUser(id string) (models.User, error)

	// Update
	UpdatePayload(models.Payload) error
	UpdateUser(models.User) error

	// Delete
	DeletePayload(models.Payload) error
	DeleteUser(models.User) error
}

func init() {
	backend = map[string]Storage{}
	fmt.Println("Init databases:")
	s, err := sqlite.New(config.Config{Domain: "localhost", Database: "db.sqlite", StandaloneHTTPS: true})
	if err != nil {
		log.Fatal(err)
	}
	backend["sqlite"] = s
	fmt.Println(backend)
}

func GetDB() Storage {
	selectedBackend := "sqlite"
	currentStorage = backend[selectedBackend]
	return currentStorage
}
