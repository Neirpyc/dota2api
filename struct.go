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
	NumResults       int     `bson:"num_results" json:"num_results"`
	ResultsRemaining int     `bson:"results_remaining" json:"results_remaining"`
	Status           int     `json:"status"`
	TotalResults     int     `bson:"total_results" json:"total_results"`
}

type MatchDetails struct {
	Result Match `json:"result"`
}

type Match struct {
	DireTeamID            int64    `bson:"dire_team_id" json:"dire_team_id"`
	LobbyType             int      `bson:"lobby_type" json:"lobby_type"`
	MatchID               int64    `bson:"match_id" json:"match_id"`
	MatchSeqNum           int64    `bson:"match_seq_num" json:"match_seq_num"`
	Players               []Player `bson:"players" json:"players"`
	RadiantTeamID         int64    `bson:"radiant_team_id" json:"radiant_team_id"`
	StartTime             int64    `bson:"start_time" json:"start_time"`
	BarracksStatusDire    int      `bson:"barracks_status_dire" json:"barracks_status_dire"`
	BarracksStatusRadiant int      `bson:"barracks_status_radiant" json:"barracks_status_radiant"`
	Cluster               int      `bson:"cluster" json:"cluster"`
	DireCaptain           int64    `bson:"dire_captain" json:"dire_captain"`
	Duration              int      `bson:"duration" json:"duration"`
	FirstBloodTime        int      `bson:"first_blood_time" json:"first_blood_time"`
	GameMode              int      `bson:"game_mode" json:"game_mode"`
	HumanPlayers          int      `bson:"human_players" json:"human_players"`
	Leagueid              int      `bson:"leagueid" json:"leagueid"`
	NegativeVotes         int      `bson:"negative_votes" json:"negative_votes"`
	PositiveVotes         int      `bson:"positive_votes" json:"positive_votes"`
	RadiantCaptain        int64    `bson:"radiant_captain" json:"radiant_captain"`
	RadiantWin            bool     `bson:"radiant_win" json:"radiant_win"`
	TowerStatusDire       int      `bson:"tower_status_dire" json:"tower_status_dire"`
	Tower_StatusRadiant   int      `bson:"tower_status_radiant" json:"tower_status_radiant"`
	Error                 string   `bson:"error" json:"error"`
}

type Player struct {
	AccountID       int64            `bson:"account_id" json:"account_id"`
	HeroID          int              `bson:"hero_id" json:"hero_id"`
	PlayerSlot      int              `bson:"player_slot" json:"player_slot"`
	AbilityUpgrades []AbilityUpgrade `bson:"ability_upgrades" json:"ability_upgrades"`
	Assists         int              `bson:"assists" json:"assists"`
	Deaths          int              `bson:"deaths" json:"deaths"`
	Denies          int              `bson:"denies" json:"denies"`
	Gold            int              `bson:"gold" json:"gold"`
	GoldPerMin      int              `bson:"gold_per_min" json:"gold_per_min"`
	GoldSpent       int              `bson:"gold_spent" json:"gold_spent"`
	HeroDamage      int              `bson:"hero_damage" json:"hero_damage"`
	HeroHealing     int              `bson:"hero_healing" json:"hero_healing"`
	Item0           int              `bson:"item_0" json:"item_0"`
	Item1           int              `bson:"item_1" json:"item_1"`
	Item2           int              `bson:"item_2" json:"item_2"`
	Item3           int              `bson:"item_3" json:"item_3"`
	Item4           int              `bson:"item_4" json:"item_4"`
	Item5           int              `bson:"item_5" json:"item_5"`
	Kills           int              `bson:"kills" json:"kills"`
	LastHits        int              `bson:"last_hits" json:"last_hits"`
	LeaverStatus    int              `bson:"leaver_status" json:"leaver_status"`
	Level           int              `bson:"level" json:"level"`
	TowerDamage     int              `bson:"tower_damage" json:"tower_damage"`
	XpPerMin        int              `bson:"xp_per_min" json:"xp_per_min"`

	Avatar                   string `bson:"avatar" json:"avatar"`
	Avatarfull               string `bson:"avatarfull" json:"avatarfull"`
	Avatarmedium             string `bson:"avatarmedium" json:"avatarmedium"`
	Communityvisibilitystate int    `bson:"communityvisibilitystate" json:"communityvisibilitystate"`
	Lastlogoff               int64  `bson:"lastlogoff" json:"lastlogoff"`
	Personaname              string `bson:"personaname" json:"personaname"`
	Personastate             int    `bson:"personastate" json:"personastate"`
	Personastateflags        int    `bson:"personastateflags" json:"personastateflags"`
	Primaryclanid            string `bson:"primaryclanid" json:"primaryclanid"`
	Profilestate             int    `bson:"profilestate" json:"profilestate"`
	Profileurl               string `bson:"profileurl" json:"profileurl"`
	Realname                 string `bson:"realname" json:"realname"`
	Steamid                  string `bson:"steamid" json:"steamid"`
	Timecreated              int64  `bson:"timecreated" json:"timecreated"`
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
	LocalizedName string `bson:"localized_name" json:"localized_name"`
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
	FriendSince  int64  `bson:"friend_since" json:"friend_since"`
}
type TeamInfo struct {
	Result struct {
		Status float64 `json:"status"`
		Teams  []Team  `json:"teams"`
	} `json:"result"`
}

type Team struct {
	AdminAccountID               int64  `bson:"admin_account_id" json:"admin_account_id"`
	CountryCode                  string `bson:"country_code" json:"country_code"`
	GamesPlayedWithCurrentRoster int    `bson:"games_played_with_current_roster" json:"games_played_with_current_roster"`
	Logo                         int64  `json:"logo"`
	LogoSponsor                  int    `bson:"logo_sponsor" json:"logo_sponsor"`
	Name                         string `json:"name"`
	Player0AccountID             int64  `bson:"player_0_account_id" json:"player_0_account_id"`
	Player1AccountID             int64  `bson:"player_1_account_id" json:"player_1_account_id"`
	Player2AccountID             int64  `bson:"player_2_account_id" json:"player_2_account_id"`
	Player3AccountID             int64  `bson:"player_3_account_id" json:"player_3_account_id"`
	Player4AccountID             int64  `bson:"player_4_account_id" json:"player_4_account_id"`
	Player5AccountID             int64  `bson:"player_5_account_id" json:"player_5_account_id"`
	Rating                       string `json:"rating"`
	Tag                          string `json:"tag"`
	TeamID                       int64  `bson:"team_id" json:"team_id"`
	TimeCreated                  int64  `bson:"time_created" json:"time_created"`
	URL                          string `json:"url"`
}
