package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/edznux/wonder-xss/config"
	"github.com/edznux/wonder-xss/storage/models"
)

type Sqlite struct {
	file string
	db   *sql.DB
}

func New(config config.Config) (*Sqlite, error) {
	fmt.Println("Setup SQLite")
	fmt.Printf("%+v", config)
	s := Sqlite{file: "db.sqlite"}
	s.Init(config)

	fmt.Println(s)
	// Check if tables are created to we don't override
	needSetup := false
	rows, err := s.db.Query(SELECT_ALL_PAYLOADS)
	fmt.Println(rows, err)
	if err != nil {
		needSetup = true
	}

	if needSetup {
		fmt.Println("Need setup")
		s.Setup()
	}

	fmt.Println("Set up done")
	rows, err = s.db.Query(SELECT_ALL_PAYLOADS)
	fmt.Println(rows, err)
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

	fmt.Println("Creating loots' table")
	_, err = s.db.Exec(CREATE_TABLE_LOOTS)
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

func (s *Sqlite) CreateLoot(loot models.Loot, payloadIDOrAlias string) (models.Loot, error) {
	// id, payload_id, alias_id
	// TODO : store the alias_ID and not the alias directly
	_, err := s.db.Exec(INSERT_LOOT, loot.ID, loot.PayloadID, payloadIDOrAlias)
	if err != nil {
		fmt.Println(err)
		return models.Loot{}, err
	}

	return s.GetLoot(loot.ID)
}

// Read
func (s *Sqlite) GetPayloads() ([]models.Payload, error) {

	fmt.Println("GetPayloads")
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

func (s *Sqlite) GetLoot(id string) (models.Loot, error) {

	fmt.Println("GetLoot(", id, ")")
	row := s.db.QueryRow(SELECT_LOOT, id)

	var res models.Loot
	err := row.Scan(&res.ID, &res.PayloadID, &res.AliasID, &res.TriggeredAt)
	if err == sql.ErrNoRows {
		return models.Loot{}, models.NoSuchItem
	}

	if err != nil {
		fmt.Println(err)
		return models.Loot{}, err
	}
	return res, nil
}

func (s *Sqlite) GetLoots() ([]models.Loot, error) {

	fmt.Println("GetLoots")

	res := []models.Loot{}

	rows, err := s.db.Query(SELECT_ALL_LOOTS)
	if err != nil {
		fmt.Println("Error querying the db (loots):", err)
		return nil, err
	}

	var tmpRes models.Loot
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
