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

func IsMatchIDInSqlite(matchId string) bool {
	return isMatchIDInSqlite(getSqliteConn(), matchId)
}

func isMatchIDInSqlite(db *sql.DB, matchId string) bool {
	// Define the query to check if the match ID exists
	query := "SELECT COUNT(1) FROM matches WHERE id = ?"

	var count int

	// Execute the query with the provided match ID
	err := db.QueryRow(query, matchId).Scan(&count)
	if err != nil {
		log.Printf("Error querying the database: %v\n", err)
		return false
	}

	// Return true if the count is greater than 0
	return count > 0
}

func saveMatchToSqlite(db *sql.DB, match types.MatchData) bool {

	isRelevant := 1

	// Check if this is a relevant game
	if match.Info.GameMode == "ARAM" || match.Info.GameType != "MATCHED_GAME" {
		isRelevant = 0
	}

	matchID := match.Metadata.MatchID

	jsonData, err := json.Marshal(match)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Insert JSON data with a custom string ID into the database
	insertSQL := `INSERT INTO matches (id, data, is_relevant) VALUES (?, ?, ?)`
	_, err = db.Exec(insertSQL, matchID, string(jsonData), isRelevant)
	if err != nil {
		// Check if the error is a unique constraint violation
		if !strings.Contains(fmt.Sprint(err), "UNIQUE") {
			fmt.Printf("Other error: %v\n", err)
		}
		return false
	}
	return isRelevant == 1
}
