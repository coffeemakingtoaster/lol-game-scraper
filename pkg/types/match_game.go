package types

type MatchData struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult    string `json:"endOfGameResult"`
		GameCreation       int64  `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int64  `json:"gameEndTimestamp"`
		GameID             int64  `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			ChampLevel                     int    `json:"champLevel"`
			ChampionID                     int    `json:"championId"`
			ChampionName                   string `json:"championName"`
			ChampionTransform              int    `json:"championTransform"`
			CommandPings                   int    `json:"commandPings"`
			ConsumablesPurchased           int    `json:"consumablesPurchased"`
			Deaths                         int    `json:"deaths"`
			IndividualPosition             string `json:"individualPosition"`
			InhibitorKills                 int    `json:"inhibitorKills"`
			InhibitorTakedowns             int    `json:"inhibitorTakedowns"`
			InhibitorsLost                 int    `json:"inhibitorsLost"`
			Kills                          int    `json:"kills"`
			Lane                           string `json:"lane"`
			ParticipantID                  int    `json:"participantId"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			Placement                      int    `json:"placement"`
			PlayerSubteamID                int    `json:"playerSubteamId"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RetreatPings                   int    `json:"retreatPings"`
			RiotIDGameName                 string `json:"riotIdGameName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SubteamPlacement               int    `json:"subteamPlacement"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TimePlayed                     int    `json:"timePlayed"`
			Win                            bool   `json:"win"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
	} `json:"info"`
}
