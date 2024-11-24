package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/types"
)

const SQLITE_PATH = "./data/data.db"

var sqliteConn *sql.DB

func getSqliteConn() *sql.DB {
	if sqliteConn != nil {
		return sqliteConn
	}
	if _, err := os.Stat(SQLITE_PATH); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(SQLITE_PATH, 0755)
		sqliteCreateCmd := exec.Command("sqlite3", SQLITE_PATH, createTableSQL)
		err := sqliteCreateCmd.Run()
		if err != nil {
			panic(err)
		}
	}
	db, _ := sql.Open("sqlite3", SQLITE_PATH)
	return db
}

func SaveMatchToSqlite(match types.MatchData) {
	saveMatchToSqlite(getSqliteConn(), match)
}

func saveMatchToSqlite(db *sql.DB, match types.MatchData) {
	matchID := match.Metadata.MatchID

	jsonData, err := json.Marshal(match)
	if err != nil {
		log.Fatal(err)
	}

	// Insert JSON data with a custom string ID into the database
	insertSQL := `INSERT INTO matches (id, data) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, matchID, string(jsonData))
	if err != nil {
		log.Fatal(err)
	}
}
