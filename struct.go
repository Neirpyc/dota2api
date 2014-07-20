package dota2api

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
}

type MatchHistory struct {
	Result MatchResult `json:"result"`
}

type MatchResult struct {
	Matches          []Match `json:"matches"`
	NumResults       int     `json:"num_results"`
	ResultsRemaining int     `json:"results_remaining"`
	Status           int     `json:"status"`
	TotalResults     int     `json:"total_results"`
}

type Match struct {
	DireTeamID    int64    `json:"dire_team_id"`
	LobbyType     int      `json:"lobby_type"`
	MatchID       int64    `json:"match_id"`
	MatchSeqNum   int64    `json:"match_seq_num"`
	Players       []Player `json:"players"`
	RadiantTeamID int64    `json:"radiant_team_id"`
	StartTime     int64    `json:"start_time"`
}

type Player struct {
	AccountID  int64 `json:"account_id"`
	HeroID     int   `json:"hero_id"`
	PlayerSlot int   `json:"player_slot"`
}
