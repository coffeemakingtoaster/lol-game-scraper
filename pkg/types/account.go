package types

type AccountResponse struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type SummonerResponse struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}
