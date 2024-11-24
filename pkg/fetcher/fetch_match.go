package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/types"
)

func FetchMatchById(id string) (types.MatchData, error) {
	// Get user puuid
	requestURL := fmt.Sprintf("%slol/match/v5/matches/%s?api_key=%s", API_BASE, id, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return types.MatchData{}, err
	}

	if res.StatusCode != 200 {
		fmt.Printf("Match fetch status error: %d", res.StatusCode)
		return types.MatchData{}, errors.New("Unexpected response code")
	}
	target := new(types.MatchData)
	json.NewDecoder(res.Body).Decode(target)
	return *target, nil
}

func FetchMatchesByUserPUUID(puuid string) ([]string, error) {
	requestURL := fmt.Sprintf("%slol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s", API_BASE, puuid, api_key)
	res, err := http.Get(requestURL)
	if err != nil {
		return []string{}, err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Match fetch status error: %d", res.StatusCode)
		return []string{}, errors.New("Unexpected response code")
	}
	matchIds := new([]string)
	json.NewDecoder(res.Body).Decode(matchIds)
	fmt.Printf("ids: %v\n", matchIds)
	return *matchIds, nil
}

func GetParticipantPUUIDFromMatch(match types.MatchData, excludePUUID string) []string {
	res := []string{}
	for _, participant := range match.Info.Participants {
		if participant.Puuid != excludePUUID {
			res = append(res, participant.Puuid)
		}
	}
	return res
}
