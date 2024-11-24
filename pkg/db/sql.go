package db

const createTableSQL = `CREATE TABLE IF NOT EXISTS matches (
		id TEXT PRIMARY KEY AUTOINCREMENT,
		data TEXT NOT NULL
	);`
