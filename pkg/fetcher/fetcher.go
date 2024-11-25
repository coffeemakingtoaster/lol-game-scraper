package fetcher

import (
	"os"

	"github.com/joho/godotenv"
)

const API_BASE = "https://europe.api.riotgames.com/"

var api_key string

func init() {
	godotenv.Load()
	api_key = os.Getenv("RIOT_API_KEY")
}
