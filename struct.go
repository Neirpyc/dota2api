package dota2api

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
}

type MatchHistory struct {
	Result MatchHistoryResult `json:"result"`
}

type MatchHistoryResult struct {
	Status           int            `json:"status" bson:"status"`
	LeagueId         int64          `json:"league_id"  bson:"league_id"`
	NumResults       int            `json:"num_results" bson:"num_results"`
	TotalResults     int            `json:"total_results" bson:"total_results"`
	ResultsRemaining int            `json:"results_remaining" bson:"results_remaining"`
	Matches          []MatchSummary `json:"matches" bson:"matches"`
}

type MatchSummary struct {
	SeriesID      int             `json:"series_id" bson:"series_id"`
	SeriesType    int             `json:"series_type" bson:"series_type"`
	MatchID       int64           `json:"match_id" bson:"match_id"`
	MatchSeqNum   int64           `json:"match_seq_num" bson:"match_seq_num"`
	StartTime     int             `json:"start_time" bson:"start_time"`
	LobbyType     int             `json:"lobby_type" bson:"loby_type"`
	RadiantTeamID int             `json:"radiant_team_id" bson:"radiant_team_id"`
	DireTeamID    int             `json:"dire_team_id" bson:"dire_team_id"`
	Players       []PlayerSummary `json:"players" bson:"players"`
}

type PlayerSummary struct {
	AccountID  int `json:"account_id" bson:"account_id"`
	PlayerSlot int `json:"player_slot" bson:"player_slot"`
	HeroID     int `json:"hero_id" bson:"hero_id"`
}

type MatchDetails struct {
	Result Match `json:"result"`
}

type Match struct {
	Error                 string      `bson:"error" json:"error" bson:"error"`
	Players               []Player    `json:"players" bson:"players"`
	RadiantWin            bool        `json:"radiant_win" bson:"radiant_win"`
	Duration              int         `json:"duration" bson:"duration"`
	PreGameDuration       int         `json:"pre_game_duration" bson:"pre_game_duration"`
	StartTime             int         `json:"start_time" bson:"start_time"`
	MatchID               int64       `json:"match_id" bson:"match_id"`
	MatchSeqNum           int64       `json:"match_seq_num" bson:"match_seq_num"`
	TowerStatusRadiant    int         `json:"tower_status_radiant" bson:"tower_status_radiant"`
	TowerStatusDire       int         `json:"tower_status_dire" bson:"tower_status_dire"`
	BarracksStatusRadiant int         `json:"barracks_status_radiant" bson:"barracks_status_radiant"`
	BarracksStatusDire    int         `json:"barracks_status_dire" bson:"barracks_status_dire"`
	Cluster               int         `json:"cluster" bson:"cluster"`
	FirstBloodTime        int         `json:"first_blood_time" bson:"first_blood_time"`
	LobbyType             int         `json:"lobby_type" bson:"lobby_type"`
	HumanPlayers          int         `json:"human_players" bson:"human_players"`
	Leagueid              int         `json:"league_id" bson:"league_id"`
	PositiveVotes         int         `json:"positive_votes" bson:"positive_votes"`
	NegativeVotes         int         `json:"negative_votes" bson:"negative_votes"`
	GameMode              int         `json:"game_mode" bson:"game_mode"`
	Flags                 int         `json:"flags" bson:"flags"`
	Engine                int         `json:"engine" bson:"engine"`
	RadiantScore          int         `json:"radiant_score" bson:"radiant_score"`
	DireScore             int         `json:"dire_score" bson:"dire_score"`
	TournamentID          int         `json:"tournament_id" bson:"tournament_id"`
	TournamentRound       int         `json:"tournament_round" bson:"tournament_round"`
	RadiantTeamID         int         `json:"radiant_team_id" bson:"radiant_team_id"`
	RadiantName           string      `json:"radiant_name" bson:"radiant_name"`
	RadiantLogo           int         `json:"radiant_logo" bson:"radiant_logo"`
	RadiantTeamComplete   int         `json:"radiant_team_complete" bson:"radiant_team_complete"`
	DireTeamID            int         `json:"dire_team_id" bson:"dire_team_id"`
	DireName              string      `json:"dire_name" bson:"dire_name"`
	DireLogo              int         `json:"dire_logo" bson:"dire_logo"`
	DireTeamComplete      int         `json:"dire_team_complete" bson:"dire_team_complete"`
	RadiantCaptain        int         `json:"radiant_captain" bson:"radian_captain"`
	DireCaptain           int         `json:"dire_captain" bson:"dire_captain"`
	PicksBans             []PicksBans `json:"picks_bans" bson:"picks_bans"`
}

type PicksBans struct {
	IsPick bool `json:"is_pick" bson:"is_pick"`
	HeroID int  `json:"hero_id" bson:"hero_id"`
	Team   int  `json:"team" bson:"team"`
	Order  int  `json:"order" bson:"order"`
}

type Player struct {
	AccountID       int              `json:"account_id" bson:"account_id"`
	PlayerSlot      int              `json:"player_slot" bson:"player_slot"`
	HeroID          int              `json:"hero_id" bson:"hero_id"`
	Item0           int              `json:"item_0" bson:"item_0"`
	Item1           int              `json:"item_1" bson:"item_1"`
	Item2           int              `json:"item_2" bson:"item_2"`
	Item3           int              `json:"item_3" bson:"item_3"`
	Item4           int              `json:"item_4" bson:"item_4"`
	Item5           int              `json:"item_5" bson:"item_5"`
	Kills           int              `json:"kills" bson:"kills"`
	Deaths          int              `json:"deaths" bson:"deaths"`
	Assists         int              `json:"assists" bson:"assists"`
	LeaverStatus    int              `json:"leaver_status" bson:"lever_status"`
	LastHits        int              `json:"last_hits" bson:"last_hits"`
	Denies          int              `json:"denies" bson:"denies"`
	GoldPerMin      int              `json:"gold_per_min" bson:"gold_per_min"`
	XpPerMin        int              `json:"xp_per_min" bson:"xp_per_min"`
	Level           int              `json:"level" bson:"level"`
	Gold            int              `json:"gold" bson:"gold"`
	GoldSpent       int              `json:"gold_spent" bson:"gold_spent"`
	HeroDamage      int              `json:"hero_damage" bson:"hero_damage"`
	TowerDamage     int              `json:"tower_damage" bson:"tower_damage"`
	HeroHealing     int              `json:"hero_healing" bson:"hero_healing"`
	AbilityUpgrades []AbilityUpgrade `json:"ability_upgrades" bson:"ability_upgrades"`
}

type AbilityUpgrade struct {
	Ability int `json:"ability" bson:"ability"`
	Level   int `json:"level" bson:"level"`
	Time    int `json:"time" bson:"time"`
}

type PlayerSummaries struct {
	Response struct {
		Players []PlayerAccount `json:"players" bson:"players"`
	} `json:"response" bson:"response"`
}

type PlayerAccount struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Lastlogoff               int    `json:"lastlogoff"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Gameextrainfo            string `json:"gameextrainfo"`
	Gameid                   string `json:"gameid"`
	Loccountrycode           string `json:"loccountrycode"`
	Locstatecode             string `json:"locstatecode"`
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

type LeagueList struct {
	Result struct {
		Leagues []struct {
			Name          string `json:"name" bson:"name"`
			Leagueid      int    `json:"leagueid" bson:"league_id"`
			Description   string `json:"description" bson:"description"`
			TournamentURL string `json:"tournament_url" bson:"tournament_url"`
			Itemdef       int    `json:"itemdef" bson:"item_def"`
		} `json:"leagues" bson:"leagues"`
	} `json:"result" bson:"result"`
}

type LiveGames struct {
	Result struct {
		Games  []Game `json:"games" bson:"games"`
		Status int    `json:"status" bson:"status"`
	} `json:"result" bson:"result"`
}

type Game struct {
	Players []struct {
		AccountID int    `json:"account_id" bson:"account_id"`
		Name      string `json:"name" bson:"name"`
		HeroID    int    `json:"hero_id" bson:"hero_id"`
		Team      int    `json:"team" bson:"team"`
	} `json:"players" bson:"players"`
	LobbyID           int64       `json:"lobby_id" bson:"lobby_id"`
	MatchID           int64       `json:"match_id" bson:"match_id"`
	Spectators        int         `json:"spectators" bson:"spectators"`
	SeriesID          int         `json:"series_id" bson:"series_id"`
	GameNumber        int         `json:"game_number" bson:"game_number"`
	LeagueID          int         `json:"league_id" bson:"league_id"`
	StreamDelayS      int         `json:"stream_delay_s" bson:"steam_delay_s"`
	RadiantSeriesWins int         `json:"radiant_series_wins" bson:"radiant_series_wins"`
	DireSeriesWins    int         `json:"dire_series_wins" bson:"dire_series_win"`
	SeriesType        int         `json:"series_type" bson:"series_type"`
	LeagueSeriesID    int         `json:"league_series_id" bson:"league_series_id"`
	LeagueGameID      int         `json:"league_game_id" bson:"league_game_id"`
	StageName         string      `json:"stage_name" bson:"stage_name"`
	LeagueTier        int         `json:"league_tier" bson:"league_tier"`
	Scoreboard        Scoreboard  `json:"scoreboard" bson:"scoreboard"`
	DireTeam          PlayersTeam `json:"dire_team,omitempty" bson:"dire_team"`
	RadiantTeam       PlayersTeam `json:"radiant_team,omitempty" bson:"radiant_team"`
}

type PlayersTeam struct {
	TeamName string `json:"team_name" bson:"team_name"`
	TeamID   int    `json:"team_id" bson:"team_id"`
	TeamLogo int64  `json:"team_logo" bson:"team_logo"`
	Complete bool   `json:"complete" bson:"complete"`
}

type Scoreboard struct {
	Duration           float64 `json:"duration" bson:"duration"`
	RoshanRespawnTimer int     `json:"roshan_respawn_timer" bson:"roshan_respawn"`
	Radiant            Radiant `json:"radiant" bson:"radiant"`
	Dire               Dire    `json:"dire" bson:"dire"`
}

type Radiant struct {
	Score         int `json:"score" bson:"score"`
	TowerState    int `json:"tower_state" bson:"tower_state"`
	BarracksState int `json:"barracks_state" bson:"barracks_state"`
	Picks         []struct {
		HeroID int `json:"hero_id" bson:"hero_id"`
	} `json:"picks" bson:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id" bson:"hero_id"`
	} `json:"bans" bson:"bans"`
	Players   []LivePlayer `json:"players" bson:"players"`
	Abilities []struct {
		AbilityID    int `json:"ability_id" bson:"ability_id"`
		AbilityLevel int `json:"ability_level" bson:"ability_level"`
	} `json:"abilities" bson:"abilities"`
}

type Dire struct {
	Score         int `json:"score" bson:"score"`
	TowerState    int `json:"tower_state" bson:"tower_state"`
	BarracksState int `json:"barracks_state" bson:"barracks_state"`
	Picks         []struct {
		HeroID int `json:"hero_id" bson:"heor_id"`
	} `json:"picks" bson:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id" bson:"hero_id"`
	} `json:"bans" bson:"bans"`
	Players   []LivePlayer `json:"players" bson:"players"`
	Abilities []struct {
		AbilityID    int `json:"ability_id" bson:"ability_id"`
		AbilityLevel int `json:"ability_level" bson:"ability_level"`
	} `json:"abilities" bson:"abilities"`
}

type LivePlayer struct {
	PlayerSlot       int     `json:"player_slot" bson:"player_slot"`
	AccountID        int     `json:"account_id" bson:"account_id"`
	HeroID           int     `json:"hero_id" bson:"hero_id"`
	Kills            int     `json:"kills" bson:"kills"`
	Death            int     `json:"death" bson:"death"`
	Assists          int     `json:"assists" bson:"assists"`
	LastHits         int     `json:"last_hits" bson:"last_hits"`
	Denies           int     `json:"denies" bson:"denies"`
	Gold             int     `json:"gold" bson:"gold"`
	Level            int     `json:"level" bson:"level"`
	GoldPerMin       int     `json:"gold_per_min" bson:"gold_per_min"`
	XpPerMin         int     `json:"xp_per_min" bson:"xp_per_min"`
	UltimateState    int     `json:"ultimate_state" bson:"ultimate_state"`
	UltimateCooldown int     `json:"ultimate_cooldown" bson:"ultimate_cooldown"`
	Item0            int     `json:"item0" bson:"item0"`
	Item1            int     `json:"item1" bson:"item1"`
	Item2            int     `json:"item2" bson:"item2"`
	Item3            int     `json:"item3" bson:"item3"`
	Item4            int     `json:"item4" bson:"item4"`
	Item5            int     `json:"item5" bson:"item5"`
	RespawnTimer     int     `json:"respawn_timer" bson:"respawn"`
	PositionX        float64 `json:"position_x" bson:"position_x"`
	PositionY        float64 `json:"position_y" bson:"position_y"`
	NetWorth         int     `json:"net_worth" bson:"net_worth"`
}
