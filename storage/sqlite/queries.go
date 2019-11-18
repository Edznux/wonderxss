package sqlite

var CREATE_TABLE_USERS = `
CREATE TABLE users (
	id          TEXT NOT NULL PRIMARY KEY,
	username    TEXT NOT NULL unique,
	password    TEXT NOT NULL,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	modified_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
`

var CREATE_TABLE_EXECUTIONS = `
CREATE TABLE executions (
	id            TEXT NOT NULL PRIMARY KEY,
	payload_id    TEXT NOT NULL,
	alias_id      TEXT NOT NULL,
	triggered_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	FOREIGN KEY(alias_id) REFERENCES aliases(id),
	FOREIGN KEY(payload_id) REFERENCES payloads(id)
);
`

var CREATE_TABLE_ALIASES = `
CREATE TABLE aliases (
	id          TEXT NOT NULL PRIMARY KEY,
	payload_id  TEXT NOT NULL,
	alias       TEXT NOT NULL unique,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	modified_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	FOREIGN KEY(payload_id) REFERENCES payloads(id)
);
`

var CREATE_TABLE_PAYLOADS = `
CREATE TABLE payloads (
	id          TEXT NOT NULL PRIMARY KEY,
	name        TEXT NOT NULL unique,
	hash        TEXT NOT NULL,
	content     TEXT NOT NULL,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	modified_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
`

var CREATE_TABLE_COLLECTORS = `
CREATE TABLE collectors (
	id          TEXT NOT NULL PRIMARY KEY,
	content     TEXT NOT NULL,
	created_at  DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
);
`

var SELECT_ALIAS_BY_SHORTNAME = `
SELECT *
FROM aliases
WHERE alias = ?;
`

var SELECT_ALIAS_BY_PAYLOAD_ID = `
SELECT aliases.id, aliases.payload_id, aliases.alias, aliases.created_at, aliases.modified_at
FROM aliases, payloads
WHERE aliases.payload_id = ?
AND aliases.payload_id = payloads.id;
`

var SELECT_ALIAS_BY_ID = `
SELECT *
FROM aliases
WHERE id = ?;
`

var SELECT_PAYLOAD_BY_ALIAS = `
SELECT payloads.id, payloads.name, payloads.hash, payloads.content, payloads.created_at, payloads.modified_at
FROM aliases, payloads
WHERE alias = ?
AND payloads.id = aliases.payload_id;
`

var SELECT_ALL_PAYLOADS = `
SELECT *
FROM payloads;
`

var SELECT_EXECUTION = `
SELECT id, payload_id, alias_id, triggered_at
FROM executions
WHERE id = ?;
`

var SELECT_ALL_EXECUTIONS = `
SELECT id, payload_id, alias_id, triggered_at
FROM executions;
`

var SELECT_ALL_ALIASES = `
SELECT id, payload_id, alias, created_at, modified_at
FROM aliases;
`

var SELECT_PAYLOAD_BY_ID = `
SELECT id, name, hash, content, created_at, modified_at
FROM payloads
WHERE id = ?;
`

var SELECT_PAYLOAD_BY_NAME = `
SELECT id, name, hash, content, created_at, modified_at
FROM payloads
WHERE name = ?;
`

var SELECT_USER = `
SELECT *
FROM users
WHERE id = ?;
`

var SELECT_COLLECTOR = `
SELECT *
FROM collectors
WHERE id = ?;
`

var SELECT_ALL_COLLECTOR = `
SELECT *
FROM collectors;
`

var INSERT_PAYLOAD = `INSERT INTO payloads (id, name, hash, content) VALUES (?, ?, ?, ?);`

var INSERT_USER = `INSERT INTO users (id, username, password) VALUES (?, ?, ?);`

var INSERT_ALIAS = `INSERT INTO aliases (id, payload_id, alias) VALUES (?, ?, ?);`

var INSERT_EXECUTION = `INSERT INTO executions (id, payload_id, alias_id) VALUES (?, ?, ?);`

var INSERT_COLLECTOR = `INSERT INTO collectors (id, payload_id, alias_id) VALUES (?, ?, ?);`
