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

var CREATE_TABLE_LOOTS = `
CREATE TABLE loots (
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

var SELECT_ALIAS_BY_SHORTNAME = `
SELECT *
FROM aliases
WHERE alias = ?;
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

var SELECT_LOOT = `
SELECT id, payload_id, alias_id, triggered_at
FROM loots
WHERE id = ?;
`

var SELECT_ALL_LOOTS = `
SELECT id, payload_id, alias_id, triggered_at
FROM loots;
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

var INSERT_PAYLOAD = `INSERT INTO payloads (id, name, hash, content) VALUES (?, ?, ?, ?);`

var INSERT_USER = `INSERT INTO users (id, username, password) VALUES (?, ?, ?);`

var INSERT_ALIAS = `INSERT INTO aliases (id, payload_id, alias) VALUES (?, ?, ?);`

var INSERT_LOOT = `INSERT INTO loots (id, payload_id, alias_id) VALUES (?, ?, ?);`
