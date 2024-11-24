package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/types"
)

func FetchMatchById(id string) {
	// Get user puuid
	requestURL := fmt.Sprintf("%slol/match/v5/matches/%s?api_key=%s", API_BASE, id, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		panic("No")
	}
	target := new(types.MatchData)
	json.NewDecoder(res.Body).Decode(target)
	fmt.Printf("%v\n", target)
}

func GetMatchesByUserPUUID(puuid string) []string {
	requestURL := fmt.Sprintf("%slol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s", API_BASE, puuid, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		panic("No")
	}
	matchIds := new([]string)
	json.NewDecoder(res.Body).Decode(matchIds)
	fmt.Printf("ids: %v\n", matchIds)
	return *matchIds
}
