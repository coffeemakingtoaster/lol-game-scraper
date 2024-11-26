package main

import (
	"fmt"
	"os"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/fetcher"
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
	if len(user_name) > 0 && len(user_tagline) > 0 {
		fmt.Println("Adding provided entry user...")
		sq.AddRiotAccToQueue(user_name, user_tagline)
	} else {
		fmt.Println("No entry user provided!\nAdding random ranked player...\n")
		puuid := fetcher.GetRandomPUUID()
		sq.AddPuuidToQueue(puuid)
	}
	sq.Run()
}
