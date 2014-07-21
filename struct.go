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

type MatchDetails struct {
	Result Match `json:"result"`
}

type Match struct {
	DireTeamID            int64    `json:"dire_team_id"`
	LobbyType             int      `json:"lobby_type"`
	MatchID               int64    `json:"match_id"`
	MatchSeqNum           int64    `json:"match_seq_num"`
	Players               []Player `json:"players"`
	RadiantTeamID         int64    `json:"radiant_team_id"`
	StartTime             int64    `json:"start_time"`
	BarracksStatusDire    int      `json:"barracks_status_dire"`
	BarracksStatusRadiant int      `json:"barracks_status_radiant"`
	Cluster               int      `json:"cluster"`
	DireCaptain           int64    `json:"dire_captain"`
	Duration              int      `json:"duration"`
	FirstBloodTime        int      `json:"first_blood_time"`
	GameMode              int      `json:"game_mode"`
	HumanPlayers          int      `json:"human_players"`
	Leagueid              int      `json:"leagueid"`
	NegativeVotes         int      `json:"negative_votes"`
	PositiveVotes         int      `json:"positive_votes"`
	RadiantCaptain        int64    `json:"radiant_captain"`
	RadiantWin            bool     `json:"radiant_win"`
	TowerStatusDire       int      `json:"tower_status_dire"`
	TowerStatusRadiant    int      `json:"tower_status_radiant"`
	Error                 string   `json:"error"`
}

type Player struct {
	AccountID       int64            `json:"account_id"`
	HeroID          int              `json:"hero_id"`
	PlayerSlot      int              `json:"player_slot"`
	AbilityUpgrades []AbilityUpgrade `json:"ability_upgrades"`
	Assists         int              `json:"assists"`
	Deaths          int              `json:"deaths"`
	Denies          int              `json:"denies"`
	Gold            int              `json:"gold"`
	GoldPerMin      int              `json:"gold_per_min"`
	GoldSpent       int              `json:"gold_spent"`
	HeroDamage      int              `json:"hero_damage"`
	HeroHealing     int              `json:"hero_healing"`
	Item0           int              `json:"item_0"`
	Item1           int              `json:"item_1"`
	Item2           int              `json:"item_2"`
	Item3           int              `json:"item_3"`
	Item4           int              `json:"item_4"`
	Item5           int              `json:"item_5"`
	Kills           int              `json:"kills"`
	LastHits        int              `json:"last_hits"`
	LeaverStatus    int              `json:"leaver_status"`
	Level           int              `json:"level"`
	TowerDamage     int              `json:"tower_damage"`
	XpPerMin        int              `json:"xp_per_min"`

	Avatar                   string `json:"avatar"`
	Avatarfull               string `json:"avatarfull"`
	Avatarmedium             string `json:"avatarmedium"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Lastlogoff               int64  `json:"lastlogoff"`
	Personaname              string `json:"personaname"`
	Personastate             int    `json:"personastate"`
	Personastateflags        int    `json:"personastateflags"`
	Primaryclanid            string `json:"primaryclanid"`
	Profilestate             int    `json:"profilestate"`
	Profileurl               string `json:"profileurl"`
	Realname                 string `json:"realname"`
	Steamid                  string `json:"steamid"`
	Timecreated              int64  `json:"timecreated"`
}

type AbilityUpgrade struct {
	Ability int `json:"ability"`
	Level   int `json:"level"`
	Time    int `json:"time"`
}

type PlayerSummaries struct {
	Response struct {
		Players struct {
			Player []Player `json:"player"`
		} `json:"players"`
	} `json:"response"`
}

type Heroes struct {
	Result struct {
		Count  int    `json:"count"`
		Heroes []Hero `json:"heroes"`
		Status int    `json:"status"`
	} `json:"result"`
}

type Hero struct {
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name"`
	Name          string `json:"name"`
}

type FriendList struct {
	Friendslist struct {
		Friends []Friend `json:"friends"`
	} `json:"friendslist"`
}

type Friend struct {
	SteamId      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int64  `json:"friend_since"`
}
