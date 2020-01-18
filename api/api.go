package api

type API interface {
	GetHealth() (string, error)

	GetAliases() ([]Alias, error)
	GetAlias(id string) (Alias, error)
	GetAliasByID(id string) (Alias, error)
	GetAliasByPayloadID(id string) (Alias, error)
	AddAlias(name string, payloadId string) (Alias, error)
	DeleteAlias(id string) error

	GetCollectors() ([]Collector, error)
	GetCollector(id string) (Collector, error)
	AddCollector(data string) (Collector, error)
	DeleteCollector(id string) error

	GetExecutions() ([]Execution, error)
	GetExecution(id string) (Execution, error)
	AddExecution(payloadID string, aliasID string) (Execution, error)
	DeleteExecution(id string) error

	GetInjections() ([]Injection, error)
	GetInjection(id string) (Injection, error)
	AddInjection(name string, content string) (Injection, error)
	DeleteInjection(id string) error

	GetPayloads() ([]Payload, error)
	ServePayload(idOrAlias string) (string, error)
	GetPayload(id string) (Payload, error)
	AddPayload(name string, content string, contentType string) (Payload, error)
	DeletePayload(id string) error

	//Login doesn't return a user but a JWT token if the auth is successful
	Login(loginParam, passwordParam, otp string) (string, error)
	GetUserByName(name string) (User, error)
	GetUser(id string) (User, error)
	CreateOTP(userID string, secret string, otp string) (User, error)
	CreateUser(username, password string) (User, error)
	DeleteUser(id string) error
}
