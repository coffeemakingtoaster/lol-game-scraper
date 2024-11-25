package queue

import (
	"fmt"
	"time"

	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/db"
	"github.com/coffeemakingtoaster/lol-game-scraper/pkg/fetcher"
)

type SummonerQueue struct {
	MatchQueue       chan string
	PUUIDQueue       chan string
	SavedMatches     int
	QueriedSummoners int
	IsReady          bool
}

func New() *SummonerQueue {
	sq := new(SummonerQueue)
	// 500 match slots in the queue
	sq.MatchQueue = make(chan string, 500)
	// 500 summoner slots in the queue
	sq.PUUIDQueue = make(chan string, 500)
	sq.SavedMatches = 0
	sq.QueriedSummoners = 0
	sq.IsReady = false
	return sq
}

// This should only be used once as all fetched participants of games already contain the PUUID
func (s *SummonerQueue) AddRiotAccToQueue(user_name string, user_tag string) {
	user := fetcher.FetchSummoner(user_name, user_tag)
	success := s.AddPuuidToQueue(user.PUUID)
	if !success {
		panic("Entryuser could not be added! This likely means that it has been fetched by another instance")
	}
	fmt.Printf("Initial user %s#%s to queue\n", user_name, user_tag)
	s.IsReady = true
}

func (s *SummonerQueue) AddPuuidToQueue(puuid string) bool {
	// Buffer of 5 just to make sure
	if len(s.PUUIDQueue) > (cap(s.PUUIDQueue) - 5) {
		return false
	}

	success := db.MarkPUUIDDone(puuid)
	if success {
		s.PUUIDQueue <- puuid
	}
	return success
}

func (s *SummonerQueue) Run() {
	for {
		// Check if queries are empty even though the queue should be ready for processing
		if (len(s.PUUIDQueue)+len(s.MatchQueue) == 0) && s.IsReady {
			panic("All Queues Empty! Scraper has run dry :c")
		}
		select {
		case matchId := <-s.MatchQueue:
			matchData, err := fetcher.FetchMatchById(matchId)
			// This means something went wrong. This is most likely due to the rate limit
			// therefore we will wait for a while and then try again later
			if err != nil {
				time.Sleep(15 * time.Second)
				s.MatchQueue <- matchId
				break
			}
			saved := db.SaveMatchToSqlite(matchData)
			if saved {
				s.SavedMatches++

				if s.SavedMatches%10 == 0 {
					fmt.Printf("Saved %d from %d summoners!\n", s.SavedMatches, s.QueriedSummoners)
				}
			}
			participants := fetcher.GetParticipantPUUIDFromMatch(matchData)
			for _, puuid := range participants {
				s.AddPuuidToQueue(puuid)
			}
			break
		case puuid := <-s.PUUIDQueue:
			if len(s.MatchQueue) > (cap(s.MatchQueue) - 25) {
				break
			}

			matchIds, err := fetcher.FetchMatchesByUserPUUID(puuid)

			// This means something went wrong. This is most likely due to the rate limit
			// therefore we will wait for a while and then try again later
			if err != nil {
				time.Sleep(15 * time.Second)
				s.AddPuuidToQueue(puuid)
				break
			}

			s.QueriedSummoners++

			for _, id := range matchIds {
				s.MatchQueue <- id
			}

			break
		}
	}
}
