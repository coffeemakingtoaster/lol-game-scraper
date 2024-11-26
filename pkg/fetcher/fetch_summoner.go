package fetcher

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/types"
)

func FetchSummoner(user_name string, user_tagline string) types.AccountResponse {
	// Get user puuid
	requestURL := fmt.Sprintf("%sriot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", API_BASE, user_name, user_tagline, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		panic("No")
	}
	target := new(types.AccountResponse)
	json.NewDecoder(res.Body).Decode(target)
	return *target
}

func GetRandomPUUID() string {
	id := getRandomSummonerID()
	return summonerIDToPUUID(id)
}

func getRandomSummonerID() string {

	ranks := []string{"IRON", "CHALLENGER", "GRANDMASTER", "MASTER", "DIAMOND", "EMERALD", "PLATINUM", "GOLD", "SILVER", "BRONZE"}
	divisions := []string{"I", "II", "III", "IV"}

	queue := "RANKED_SOLO_5x5"
	rank := ranks[rand.Intn(len(ranks))]
	division := divisions[rand.Intn(len(divisions))]
	if rank == "CHALLENGER" || strings.Contains(rank, "MASTER") {
		division = "I"
	}
	requestURL := fmt.Sprintf("https://euw1.api.riotgames.com/lol/league-exp/v4/entries/%s/%s/%s?page=1&api_key=%s", queue, rank, division, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		panic("No")
	}
	t := new(types.LeagueExp)
	json.NewDecoder(res.Body).Decode(t)
	target := *t
	chosenSummoner := target[rand.Intn(len(target))]
	return chosenSummoner.SummonerID
}

func summonerIDToPUUID(summonerID string) string {
	requestURL := fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/%s?api_key=%s", summonerID, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		panic("No")
	}
	target := new(types.SummonerResponse)
	json.NewDecoder(res.Body).Decode(target)
	fmt.Printf("Found user with PUUID: %s\n", target.Puuid)
	return target.Puuid
}
