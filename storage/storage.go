package storage

import (
	"fmt"
	"log"

	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"
	"github.com/edznux/wonderxss/storage/sqlite"
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
	CreateExecution(execution models.Execution, payloadIDOrAlias string) (models.Execution, error)
	CreateCollector(collector models.Collector) (models.Collector, error)

	// Read
	GetUser(id string) (models.User, error)

	GetPayloads() ([]models.Payload, error)
	GetPayload(id string) (models.Payload, error)
	GetPayloadByAlias(short string) (models.Payload, error)

	GetAliases() ([]models.Alias, error)
	GetAlias(id string) (models.Alias, error)
	GetAliasByID(id string) (models.Alias, error)
	GetAliasByPayloadID(payloadID string) (models.Alias, error)

	GetExecutions() ([]models.Execution, error)
	GetExecution(id string) (models.Execution, error)

	GetCollectors() ([]models.Collector, error)
	GetCollector(id string) (models.Collector, error)

	// Update
	UpdatePayload(models.Payload) error
	UpdateUser(models.User) error

	// Delete
	DeletePayload(models.Payload) error
	DeleteUser(models.User) error
}

func InitStorage(cfg config.Config) {
	backend = map[string]Storage{}
	fmt.Printf("Init databases: %+v\n", cfg)
	s, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Error while initializing storage:", err)
	}
	backend["sqlite"] = s
}

func GetDB() Storage {
	selectedBackend := "sqlite"
	currentStorage = backend[selectedBackend]
	return currentStorage
}
