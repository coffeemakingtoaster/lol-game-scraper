package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/types"
	_ "github.com/mattn/go-sqlite3"
)

const SQLITE_PATH = "./data/data.db"
const SQLITE_DIR = "./data"

var sqliteConn *sql.DB

func getSqliteConn() *sql.DB {
	if sqliteConn != nil {
		return sqliteConn
	}
	if _, err := os.Stat(SQLITE_PATH); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(SQLITE_DIR, 0755)
		sqliteCreateCmd := exec.Command("sqlite3", SQLITE_PATH, "")
		err := sqliteCreateCmd.Run()
		if err != nil {
			fmt.Println("Could not create sqlite")
			panic(err)
		}
	}
	db, _ := sql.Open("sqlite3", SQLITE_PATH)
	// the init statement is writte to be idempotent
	_, err := db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Connection could not be used")
		panic(err)
	}
	return db
}

func SaveMatchToSqlite(match types.MatchData) bool {
	return saveMatchToSqlite(getSqliteConn(), match)
}

func saveMatchToSqlite(db *sql.DB, match types.MatchData) bool {
	matchID := match.Metadata.MatchID

	jsonData, err := json.Marshal(match)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Insert JSON data with a custom string ID into the database
	insertSQL := `INSERT INTO matches (id, data) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, matchID, string(jsonData))
	if err != nil {
		// Check if the error is a unique constraint violation
		if !strings.Contains(fmt.Sprint(err), "UNIQUE") {
			fmt.Printf("Other error: %v\n", err)
		}
		return false
	}
	return true
}
