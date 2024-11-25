package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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
	time.Sleep(1 * time.Second)
	return *target
}
