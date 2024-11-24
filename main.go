package main

import (
	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/queue"
)

// Define your entry user here
// TODO: make this an env var or something
const user_name = ""
const user_tagline = ""

func main() {
	sq := queue.New()
	// Start queue handler
	go sq.Run()
	sq.AddRiotAccToQueue(user_name, user_tagline)
}
