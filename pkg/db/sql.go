package db

const createTableSQL = `CREATE TABLE IF NOT EXISTS matches (
		id TEXT PRIMARY KEY,
		data TEXT NOT NULL,
		is_relevant INTEGER NOT NULL
	);`
