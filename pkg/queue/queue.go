package queue

import "github.com/coffeemakingtoaster/lol-game-scraper/pkg/fetcher"

type SummonerQueue struct {
	status string
	Queue  chan string
}

func New() *SummonerQueue {
	sq := new(SummonerQueue)
	// 500 match slots in the queue
	sq.Queue = make(chan string, 500)
	return sq
}

func (s *SummonerQueue) AddRiotAccToQueue(user_name string, user_tag string) {
	user := fetcher.FetchSummoner(user_name, user_tag)
	s.AddSummonnerToQueue(user.PUUID)
}

func (s *SummonerQueue) AddSummonnerToQueue(puuid string) {
	// Is there room? 20 games are added so we want 25 empty
	if len(s.Queue) > (cap(s.Queue) - 25) {
		// dont do anything
		return
	}
	// TODO: Check if summoner has been fetched by other
	matchIds := fetcher.GetMatchesByUserPUUID(puuid)
	for _, matchId := range matchIds {
		s.Queue <- matchId
	}
}

func (s *SummonerQueue) Run() {
	for {
		select {
		case matchId := <-s.Queue:

		}
	}
}
