package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"
)

type Sqlite struct {
	file string
	db   *sql.DB
}

func New(cfg config.Config) (*Sqlite, error) {
	file := cfg.Storages["sqlite"].File
	fmt.Printf("Setup SQLite, using file: %+v\n", file)
	s := Sqlite{file: file}
	s.Init(cfg)

	fmt.Println(s)
	// Check if tables are created so we don't override
	needSetup := false
	_, err := s.db.Query(SELECT_ALL_PAYLOADS)
	if err != nil {
		needSetup = true
	}

	if needSetup {
		fmt.Println("Need setup")
		s.Setup()
	}

	fmt.Println("Set up done")
	_, err = s.db.Query(SELECT_ALL_PAYLOADS)
	if err != nil {
		fmt.Println(err)
	}
	return &s, nil
}

func (s *Sqlite) Init(config config.Config) error {
	var err error
	s.db, err = sql.Open("sqlite3", s.file)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) Setup() error {
	fmt.Println("Creating users' table")
	_, err := s.db.Exec(CREATE_TABLE_USERS)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Creating payloads' table")
	_, err = s.db.Exec(CREATE_TABLE_PAYLOADS)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Creating aliases' table")
	_, err = s.db.Exec(CREATE_TABLE_ALIASES)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Creating Execution' table")
	_, err = s.db.Exec(CREATE_TABLE_EXECUTIONS)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	//return last error, but keep executing all instruction
	return err
}

//Create
func (s *Sqlite) CreatePayload(payload models.Payload) (models.Payload, error) {
	_, err := s.db.Exec(INSERT_PAYLOAD, payload.ID, payload.Name, payload.Hash, payload.Content)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return models.Payload{}, err
	}

	return s.GetPayload(payload.ID)
}

func (s *Sqlite) CreateUser(user models.User) (models.User, error) {
	_, err := s.db.Exec(INSERT_USER, user.ID, user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	return s.GetUser(user.ID)
}

func (s *Sqlite) CreateAlias(alias models.Alias) (models.Alias, error) {
	_, err := s.db.Exec(INSERT_ALIAS, alias.ID, alias.PayloadID, alias.Short)
	if err != nil {
		fmt.Println(err)
		return models.Alias{}, err
	}

	return s.GetAlias(alias.ID)
}

func (s *Sqlite) CreateCollector(collector models.Collector) (models.Collector, error) {
	_, err := s.db.Exec(INSERT_COLLECTOR, collector.ID, collector.PayloadID, collector.Data)
	if err != nil {
		fmt.Println(err)
		return models.Collector{}, err
	}

	return s.GetCollector(collector.ID)
}

func (s *Sqlite) CreateExecution(execution models.Execution, payloadIDOrAlias string) (models.Execution, error) {
	// id, payload_id, alias_id
	// TODO : store the alias_ID and not the alias directly
	_, err := s.db.Exec(INSERT_EXECUTION, execution.ID, execution.PayloadID, payloadIDOrAlias)
	if err != nil {
		fmt.Println(err)
		return models.Execution{}, err
	}

	return s.GetExecution(execution.ID)
}

// Read
func (s *Sqlite) GetPayloads() ([]models.Payload, error) {

	fmt.Println("sqlite.GetPayloads")
	res := []models.Payload{}

	rows, err := s.db.Query(SELECT_ALL_PAYLOADS)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tmpRes models.Payload
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.Name, &tmpRes.Hash, &tmpRes.Content, &tmpRes.CreatedAt, &tmpRes.ModifiedAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	return res, nil
}

func (s *Sqlite) GetPayload(id string) (models.Payload, error) {

	row := s.db.QueryRow(SELECT_PAYLOAD_BY_ID, id)

	var res models.Payload
	err := row.Scan(&res.ID, &res.Name, &res.Hash, &res.Content, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Payload{}, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return models.Payload{}, err
	}
	return res, nil
}

func (s *Sqlite) GetPayloadByAlias(short string) (models.Payload, error) {

	row := s.db.QueryRow(SELECT_PAYLOAD_BY_ALIAS, short)

	var res models.Payload
	err := row.Scan(&res.ID, &res.Name, &res.Hash, &res.Content, &res.CreatedAt, &res.ModifiedAt)
	if err == sql.ErrNoRows {
		return models.Payload{}, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return models.Payload{}, err
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
		return models.Alias{}, err
	}
	return res, nil
}

func (s *Sqlite) GetExecution(id string) (models.Execution, error) {

	fmt.Println("GetExecution(", id, ")")
	row := s.db.QueryRow(SELECT_EXECUTION, id)

	var res models.Execution
	err := row.Scan(&res.ID, &res.PayloadID, &res.AliasID, &res.TriggeredAt)
	if err == sql.ErrNoRows {
		return models.Execution{}, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return models.Execution{}, err
	}
	return res, nil
}

func (s *Sqlite) GetExecutions() ([]models.Execution, error) {

	fmt.Println("GetExecutions")

	res := []models.Execution{}

	rows, err := s.db.Query(SELECT_ALL_EXECUTIONS)
	if err != nil {
		fmt.Println("Error querying the db (Execution):", err)
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
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	return res, nil
}

func (s *Sqlite) GetCollector(id string) (models.Collector, error) {

	fmt.Println("GetCollector(", id, ")")
	row := s.db.QueryRow(SELECT_COLLECTOR, id)

	var res models.Collector
	err := row.Scan(&res.ID, &res.PayloadID, &res.Data, &res.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Collector{}, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return models.Collector{}, err
	}
	return res, nil
}

func (s *Sqlite) GetCollectors() ([]models.Collector, error) {

	fmt.Println("GetCollectors")

	res := []models.Collector{}

	rows, err := s.db.Query(SELECT_ALL_COLLECTOR)
	if err != nil {
		fmt.Println("Error querying the db (Collector):", err)
		return nil, err
	}

	var tmpRes models.Collector
	for rows.Next() {
		rows.Scan(&tmpRes.ID, &tmpRes.PayloadID, &tmpRes.Data, &tmpRes.CreatedAt)
		res = append(res, tmpRes)
	}

	if err == sql.ErrNoRows {
		return nil, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	return res, nil
}

func (s *Sqlite) GetAliases() ([]models.Alias, error) {

	fmt.Println("GetAliases")

	res := []models.Alias{}

	rows, err := s.db.Query(SELECT_ALL_ALIASES)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	return res, nil
}

func (s *Sqlite) GetUser(id string) (models.User, error) {
	var user models.User

	row := s.db.QueryRow(SELECT_USER, id)

	err := row.Scan(&user)
	if err == sql.ErrNoRows {
		return user, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (s *Sqlite) GetUserByName(name string) (models.User, error) {
	var user models.User

	row := s.db.QueryRow(SELECT_USER_BY_NAME, name)

	err := row.Scan(&user)
	if err == sql.ErrNoRows {
		return user, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
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
func (s *Sqlite) DeletePayload(models.Payload) error {
	return nil
}
func (s *Sqlite) DeleteUser(models.User) error {
	return nil
}
