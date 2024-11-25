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
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			ChampionID           int    `json:"championId"`
			ChampionName         string `json:"championName"`
			Deaths               int    `json:"deaths"`
			IndividualPosition   string `json:"individualPosition"`
			Kills                int    `json:"kills"`
			Lane                 string `json:"lane"`
			ParticipantID        int    `json:"participantId"`
			Placement            int    `json:"placement"`
			PlayerSubteamID      int    `json:"playerSubteamId"`
			Puuid                string `json:"puuid"`
			RiotIDGameName       string `json:"riotIdGameName"`
			RiotIDTagline        string `json:"riotIdTagline"`
			Role                 string `json:"role"`
			SubteamPlacement     int    `json:"subteamPlacement"`
			SummonerID           string `json:"summonerId"`
			SummonerLevel        int    `json:"summonerLevel"`
			SummonerName         string `json:"summonerName"`
			TeamEarlySurrendered bool   `json:"teamEarlySurrendered"`
			TeamID               int    `json:"teamId"`
			TimePlayed           int    `json:"timePlayed"`
			Win                  bool   `json:"win"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
	} `json:"info"`
}
