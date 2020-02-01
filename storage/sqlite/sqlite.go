package sqlite

import (
	"database/sql"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"

	sqlite3 "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	file string
	db   *sql.DB
}

func New() (*Sqlite, error) {
	cfg := config.Current
	file := cfg.Storages["sqlite"].File
	log.Debugf("Setup SQLite, using file: %+v\n", file)
	s := Sqlite{file: file}

	return &s, nil
}

func (s *Sqlite) Init() error {
	var err error
	s.db, err = sql.Open("sqlite3", s.file)
	if err != nil {
		return err
	}
	return nil
}

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

	log.Debugln("Creating Collectors' table")
	_, err = s.db.Exec(CREATE_TABLE_COLLECTORS)
	if err != nil {
		log.Errorln(err)
		lastErr = err
	}
	return lastErr
}

//Create
func (s *Sqlite) CreatePayload(payload models.Payload) (models.Payload, error) {
	_, err := s.db.Exec(INSERT_PAYLOAD, payload.ID, payload.Name, payload.Hashes.String(), payload.Content, payload.ContentType)
	if err != nil {
		log.Errorln(err)
		return models.Payload{}, err
	}

	return s.GetPayload(payload.ID)
}

func (s *Sqlite) CreateUser(user models.User) (models.User, error) {
	_, err := s.db.Exec(INSERT_USER, user.ID, user.Username, user.Password)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

func (s *Sqlite) CreateOTP(user models.User, TOTPSecret string) (models.User, error) {
	_, err := s.db.Exec(UPDATE_ADD_TOTP, 1, TOTPSecret, user.ID)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

func (s *Sqlite) RemoveOTP(user models.User) (models.User, error) {
	_, err := s.db.Exec(UPDATE_ADD_TOTP, 0, "", user.ID)
	if err != nil {
		log.Errorln(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

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

func (s *Sqlite) CreateCollector(collector models.Collector) (models.Collector, error) {
	_, err := s.db.Exec(INSERT_COLLECTOR, collector.ID, collector.Data)
	if err != nil {
		log.Errorln(err)
		return models.Collector{}, err
	}

	return s.GetCollector(collector.ID)
}

func (s *Sqlite) CreateInjection(injection models.Injection) (models.Injection, error) {
	_, err := s.db.Exec(INSERT_INJECTION, injection.ID, injection.Name, injection.Content)
	if err != nil {
		log.Errorln("CreateInjection failed:", err)
		return models.Injection{}, err
	}

	return s.GetInjection(injection.ID)
}

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

// Read
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

func (s *Sqlite) GetCollector(id string) (models.Collector, error) {

	log.Debugln("GetCollector(", id, ")")
	row := s.db.QueryRow(SELECT_COLLECTOR, id)

	var res models.Collector
	err := row.Scan(&res.ID, &res.Data, &res.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Collector{}, models.NoSuchItem
	}

	if err != nil {
		log.Println(err)
		return models.Collector{}, err
	}
	return res, nil
}

func (s *Sqlite) GetCollectors() ([]models.Collector, error) {

	res := []models.Collector{}

	rows, err := s.db.Query(SELECT_ALL_COLLECTOR)
	if err != nil {
		log.Println("Error querying the db (Collector):", err)
		return nil, err
	}

	var tmpRes models.Collector
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

//Update
func (s *Sqlite) UpdatePayload(models.Payload) error {
	return nil
}

func (s *Sqlite) UpdateUser(models.User) error {
	return nil
}

//Delete
func (s *Sqlite) DeletePayload(p models.Payload) error {
	_, err := s.db.Exec(DELETE_PAYLOAD, p.ID)
	return err
}
func (s *Sqlite) DeleteUser(u models.User) error {
	_, err := s.db.Exec(DELETE_USER, u.ID)
	return err
}
func (s *Sqlite) DeleteAlias(a models.Alias) error {
	_, err := s.db.Exec(DELETE_ALIAS, a.ID)
	return err
}
func (s *Sqlite) DeleteExecution(e models.Execution) error {
	_, err := s.db.Exec(DELETE_EXECUTION, e.ID)
	return err
}
func (s *Sqlite) DeleteCollector(c models.Collector) error {
	_, err := s.db.Exec(DELETE_COLLECTOR, c.ID)
	return err
}
func (s *Sqlite) DeleteInjection(i models.Injection) error {
	_, err := s.db.Exec(DELETE_INJECTION, i.ID)
	return err
}
