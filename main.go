package main

import (
	"fmt"
	"os"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/queue"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting scraper...")
	// loads .env
	godotenv.Load()

	user_name := os.Getenv("ENTRY_USER_NAME")
	user_tagline := os.Getenv("ENTRY_USER_TAG")

	sq := queue.New()
	// Start queue handler
	sq.AddRiotAccToQueue(user_name, user_tagline)
	sq.Run()
}
