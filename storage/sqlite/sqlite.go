package sqlite

import (
	"database/sql"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// Sqlite struct represent the base Sqlite object
// It implements the storage.Storage interface
type Sqlite struct {
	file string
	db   *sql.DB
}

// New return a new Sqlite object
func New() (*Sqlite, error) {
	cfg := config.Current
	file := cfg.Storages["sqlite"].File
	log.Debugf("Setup SQLite, using file: %+v\n", file)
	s := Sqlite{file: file}

	return &s, nil
}

// Init open the SQLite3 database
func (s *Sqlite) Init() error {
	var err error
	s.db, err = sql.Open("sqlite3", s.file)
	if err != nil {
		return err
	}
	return nil
}

// Setup create all the tables for the database
func (s *Sqlite) Setup() error {
	//return last error, but keep executing all instruction
	var lastErr error

	log.Debugln("Creating users' table")
	_, err := s.db.Exec(CREATE_TABLE_USERS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}

	log.Debugln("Creating payloads' table")
	_, err = s.db.Exec(CREATE_TABLE_PAYLOADS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}

	log.Debugln("Creating aliases' table")
	_, err = s.db.Exec(CREATE_TABLE_ALIASES)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}

	log.Debugln("Creating Executions' table")
	_, err = s.db.Exec(CREATE_TABLE_EXECUTIONS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}

	log.Debugln("Creating Injections' table")
	_, err = s.db.Exec(CREATE_TABLE_INJECTIONS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}

	log.Debugln("Creating Loots' table")
	_, err = s.db.Exec(CREATE_TABLE_LOOTS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}
	return lastErr
}

// CreatePayload create a payload based on models.Payload
// It also returns it if sucessfuly stored
func (s *Sqlite) CreatePayload(payload models.Payload) (models.Payload, error) {
	_, err := s.db.Exec(INSERT_PAYLOAD, payload.ID, payload.Name, payload.Hashes.String(), payload.Content, payload.ContentType)
	if err != nil {
		log.Errorln(err)
		return models.Payload{}, err
	}

	return s.GetPayload(payload.ID)
}

// CreateUser create a user based on models.User
// It also returns it if sucessfuly stored
func (s *Sqlite) CreateUser(user models.User) (models.User, error) {
	_, err := s.db.Exec(INSERT_USER, user.ID, user.Username, user.Password)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

// CreateOTP create an OTP token
// It returns the User if sucessfuly stored
func (s *Sqlite) CreateOTP(user models.User, TOTPSecret string) (models.User, error) {
	_, err := s.db.Exec(UPDATE_ADD_TOTP, 1, TOTPSecret, user.ID)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

// RemoveOTP remove the OTP for the user provided
func (s *Sqlite) RemoveOTP(user models.User) (models.User, error) {
	_, err := s.db.Exec(UPDATE_ADD_TOTP, 0, "", user.ID)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

// CreateAlias create an alias token based on models.Alias
// It also returns it if sucessfuly stored
func (s *Sqlite) CreateAlias(alias models.Alias) (models.Alias, error) {
	_, err := s.db.Exec(INSERT_ALIAS, alias.ID, alias.PayloadID, alias.Short)
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			log.Errorln(err)
			return models.Alias{}, models.AlreadyExist
		}
	}
	if err != nil {
		log.Errorln(err)
		return models.Alias{}, err
	}

	return s.GetAlias(alias.ID)
}

// CreateLoot create a loot token based on models.Loot
// It also returns it if sucessfuly stored
func (s *Sqlite) CreateLoot(loot models.Loot) (models.Loot, error) {
	_, err := s.db.Exec(INSERT_LOOT, loot.ID, loot.Data)
	if err != nil {
		log.Errorln(err)
		return models.Loot{}, err
	}

	return s.GetLoot(loot.ID)
}

// CreateInjection create an injection token based on models.Injection
// It also returns it if sucessfuly stored
func (s *Sqlite) CreateInjection(injection models.Injection) (models.Injection, error) {
	_, err := s.db.Exec(INSERT_INJECTION, injection.ID, injection.Name, injection.Content)
	if err != nil {
		log.Errorln("CreateInjection failed:", err)
		return models.Injection{}, err
	}

	return s.GetInjection(injection.ID)
}

// CreateExecution create an execution token based on models.Execution
// It also returns it if sucessfuly stored
func (s *Sqlite) CreateExecution(execution models.Execution, payloadIDOrAlias string) (models.Execution, error) {
	// id, payload_id, alias_id
	// TODO : store the alias_ID and not the alias directly
	_, err := s.db.Exec(INSERT_EXECUTION, execution.ID, execution.PayloadID, payloadIDOrAlias)
	if err != nil {
		log.Errorln(err)
		return models.Execution{}, err
	}

	return s.GetExecution(execution.ID)
}

// GetPayloads returns all the payloads stored in the database
func (s *Sqlite) GetPayloads() ([]models.Payload, error) {

	res := []models.Payload{}

	rows, err := s.db.Query(SELECT_ALL_PAYLOADS)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	var tmpRes models.Payload
	var hashes string
	var contentType sql.NullString

	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.Name, &hashes, &tmpRes.Content, &contentType, &tmpRes.CreatedAt, &tmpRes.ModifiedAt)
		err := json.Unmarshal([]byte(hashes), &tmpRes.Hashes)
		if err != nil {
			log.Warnln(err)
		}
		if contentType.Valid {
			tmpRes.ContentType = contentType.String
		}
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return nil, err
	}

	return res, nil
}

// GetPayload returns the select payload based on its ID
func (s *Sqlite) GetPayload(id string) (models.Payload, error) {

	row := s.db.QueryRow(SELECT_PAYLOAD_BY_ID, id)

	var res models.Payload
	var hashes string
	var contentType sql.NullString
	err := row.Scan(&res.ID, &res.Name, &hashes, &res.Content, &contentType, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Payload{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Payload{}, err
	}

	if contentType.Valid {
		res.ContentType = contentType.String
	}

	err = json.Unmarshal([]byte(hashes), &res.Hashes)
	if err != nil {
		log.Warnln(err)
		return models.Payload{}, err
	}
	return res, nil
}

// GetPayloadByAlias returns the select payload based on its alias
func (s *Sqlite) GetPayloadByAlias(short string) (models.Payload, error) {

	row := s.db.QueryRow(SELECT_PAYLOAD_BY_ALIAS, short)

	var res models.Payload
	var hashes string
	var contentType sql.NullString
	err := row.Scan(&res.ID, &res.Name, &hashes, &res.Content, &contentType, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Payload{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Payload{}, err
	}
	if contentType.Valid {
		res.ContentType = contentType.String
	}

	err = json.Unmarshal([]byte(hashes), &res.Hashes)
	if err != nil {
		log.Warnln(err)
		return models.Payload{}, err
	}
	return res, nil
}

// GetInjection returns the select Injection based on its ID
func (s *Sqlite) GetInjection(id string) (models.Injection, error) {
	row := s.db.QueryRow(SELECT_INJECTION, id)

	var res models.Injection
	err := row.Scan(&res.ID, &res.Name, &res.Content, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Injection{}, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return models.Injection{}, err
	}
	return res, nil
}

// GetInjectionByName returns the select Injection based on its name
func (s *Sqlite) GetInjectionByName(name string) (models.Injection, error) {
	row := s.db.QueryRow(SELECT_INJECTION_BY_NAME, name)

	var res models.Injection
	err := row.Scan(&res.ID, &res.Name, &res.Content, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Injection{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Injection{}, err
	}
	return res, nil
}

// GetInjections returns all injection from the database
func (s *Sqlite) GetInjections() ([]models.Injection, error) {

	res := []models.Injection{}

	rows, err := s.db.Query(SELECT_ALL_INJECTION)
	if err != nil {
		log.Errorln("Error querying the db (Injection):", err)
		return nil, err
	}

	var tmpRes models.Injection
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.Name, &tmpRes.Content, &tmpRes.CreatedAt, &tmpRes.ModifiedAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return nil, err
	}

	return res, nil
}

//GetAlias returns the selected alias by its short name
func (s *Sqlite) GetAlias(alias string) (models.Alias, error) {

	row := s.db.QueryRow(SELECT_ALIAS_BY_SHORTNAME, alias)

	var res models.Alias
	err := row.Scan(&res.ID, &res.PayloadID, &res.Short, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Alias{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Alias{}, err
	}
	return res, nil
}

//GetAliasByID returns the selected alias by its ID
func (s *Sqlite) GetAliasByID(id string) (models.Alias, error) {

	row := s.db.QueryRow(SELECT_ALIAS_BY_ID, id)

	var res models.Alias
	err := row.Scan(&res.ID, &res.PayloadID, &res.Short, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Alias{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Alias{}, err
	}
	return res, nil
}

//GetAliasByPayloadID returns the selected alias by its payloadID
func (s *Sqlite) GetAliasByPayloadID(id string) (models.Alias, error) {

	row := s.db.QueryRow(SELECT_ALIAS_BY_PAYLOAD_ID, id)

	var res models.Alias
	err := row.Scan(&res.ID, &res.PayloadID, &res.Short, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Alias{}, models.NoSuchItem
	}

	if err != nil {
		log.Warnln(err)
		return models.Alias{}, err
	}
	return res, nil
}

//GetExecution returns the execution selected by its ID
func (s *Sqlite) GetExecution(id string) (models.Execution, error) {

	log.Debugln("GetExecution(", id, ")")
	row := s.db.QueryRow(SELECT_EXECUTION, id)

	var res models.Execution
	err := row.Scan(&res.ID, &res.PayloadID, &res.AliasID, &res.TriggeredAt)
	if err == sql.ErrNoRows {
		return models.Execution{}, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return models.Execution{}, err
	}
	return res, nil
}

//GetExecutions returns all the executions from the database
func (s *Sqlite) GetExecutions() ([]models.Execution, error) {

	res := []models.Execution{}

	rows, err := s.db.Query(SELECT_ALL_EXECUTIONS)
	if err != nil {
		log.Println("Error querying the db (Execution):", err)
		return nil, err
	}

	var tmpRes models.Execution
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.PayloadID, &tmpRes.AliasID, &tmpRes.TriggeredAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

//GetLoot returns the selected loot based on its ID
func (s *Sqlite) GetLoot(id string) (models.Loot, error) {

	log.Debugln("GetLoot(", id, ")")
	row := s.db.QueryRow(SELECT_LOOT, id)

	var res models.Loot
	err := row.Scan(&res.ID, &res.Data, &res.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Loot{}, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return models.Loot{}, err
	}
	return res, nil
}

//GetLoots returns all the loots from the database
func (s *Sqlite) GetLoots() ([]models.Loot, error) {

	res := []models.Loot{}

	rows, err := s.db.Query(SELECT_ALL_LOOT)
	if err != nil {
		log.Println("Error querying the db (Loot):", err)
		return nil, err
	}

	var tmpRes models.Loot
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.Data, &tmpRes.CreatedAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

//GetAliases returns all the Aliases from the database
func (s *Sqlite) GetAliases() ([]models.Alias, error) {

	res := []models.Alias{}

	rows, err := s.db.Query(SELECT_ALL_ALIASES)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var tmpRes models.Alias
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.PayloadID, &tmpRes.Short, &tmpRes.CreatedAt, &tmpRes.ModifiedAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

//GetUser return a user based on its ID
func (s *Sqlite) GetUser(id string) (models.User, error) {
	var user models.User
	var TOTPSecret sql.NullString
	var TFEnabled int

	row := s.db.QueryRow(SELECT_USER, id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &TFEnabled, &TOTPSecret, &user.CreatedAt, &user.ModifiedAt)
	if err == sql.ErrNoRows {
		return user, models.NoSuchItem
	}

	if TFEnabled == 1 {
		user.TwoFactorEnabled = true
	}
	if TOTPSecret.Valid {
		user.TOTPSecret = TOTPSecret.String
	}

	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

//GetUserByName return an user by its name
func (s *Sqlite) GetUserByName(name string) (models.User, error) {
	var user models.User
	var TOTPSecret sql.NullString
	var TFEnabled int

	row := s.db.QueryRow(SELECT_USER_BY_NAME, name)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &TFEnabled, &TOTPSecret, &user.CreatedAt, &user.ModifiedAt)
	if err == sql.ErrNoRows {
		return user, models.NoSuchItem
	}

	if TFEnabled == 1 {
		user.TwoFactorEnabled = true
	}
	if TOTPSecret.Valid {
		user.TOTPSecret = TOTPSecret.String
	}

	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

//UpdatePayload Not Implemented Yet
//Update the payload based on the provided one.
//Only the ID field must be correct. All other field will be changed
func (s *Sqlite) UpdatePayload(models.Payload) error {
	return nil
}

//UpdateUser Not Implemented Yet
//Update the User based on the provided one.
//Only the ID field must be correct. All other field will be changed
func (s *Sqlite) UpdateUser(models.User) error {
	return nil
}

//Delete

//DeletePayload delete the provided payload from the database
func (s *Sqlite) DeletePayload(p models.Payload) error {
	_, err := s.db.Exec(DELETE_PAYLOAD, p.ID)
	return err
}

//DeleteUser delete the provided User from the database
func (s *Sqlite) DeleteUser(u models.User) error {
	_, err := s.db.Exec(DELETE_USER, u.ID)
	return err
}

//DeleteAlias delete the provided Alias from the database
func (s *Sqlite) DeleteAlias(a models.Alias) error {
	_, err := s.db.Exec(DELETE_ALIAS, a.ID)
	return err
}

//DeleteExecution delete the provided Execution from the database
func (s *Sqlite) DeleteExecution(e models.Execution) error {
	_, err := s.db.Exec(DELETE_EXECUTION, e.ID)
	return err
}

//DeleteLoot delete the provided Loot from the database
func (s *Sqlite) DeleteLoot(c models.Loot) error {
	_, err := s.db.Exec(DELETE_LOOT, c.ID)
	return err
}

//DeleteInjection delete the provided Injection from the database
func (s *Sqlite) DeleteInjection(i models.Injection) error {
	_, err := s.db.Exec(DELETE_INJECTION, i.ID)
	return err
}
