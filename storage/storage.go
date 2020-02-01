package storage

import (
	"fmt"

	log "github.com/sirupsen/logrus"

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
	Init() error

	// CRUD interface
	// Create
	CreatePayload(models.Payload) (models.Payload, error)
	CreateUser(models.User) (models.User, error)
	CreateAlias(models.Alias) (models.Alias, error)
	CreateExecution(execution models.Execution, payloadIDOrAlias string) (models.Execution, error)
	CreateCollector(collector models.Collector) (models.Collector, error)
	CreateInjection(injection models.Injection) (models.Injection, error)
	CreateOTP(user models.User, TOTPSecret string) (models.User, error)
	// Read
	GetUser(id string) (models.User, error)
	GetUserByName(name string) (models.User, error)

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

	GetInjections() ([]models.Injection, error)
	GetInjection(id string) (models.Injection, error)
	// Update
	UpdatePayload(models.Payload) error
	UpdateUser(models.User) error

	// Delete
	DeleteExecution(models.Execution) error
	DeleteInjection(models.Injection) error
	DeleteCollector(models.Collector) error
	DeletePayload(models.Payload) error
	DeleteAlias(models.Alias) error
	DeleteUser(models.User) error
	RemoveOTP(models.User) (models.User, error)
}

func LoadStorageBackends() {
	fmt.Println("LoadStorageBackends")
	backend = map[string]Storage{}
	s, err := sqlite.New()
	if err != nil {
		log.Fatal("Error while initializing storage:", err)
	}
	s.Init()
	backend["sqlite"] = s
	fmt.Printf("Initialiazed storage backends: %+v\n", backend)
}

func GetDB() Storage {
	log.Println("GetDB:", backend)
	currentStorage = backend[config.Current.Database]
	return currentStorage
}
