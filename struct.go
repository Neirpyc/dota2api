package dota2api

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
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
	Duration           float64     `json:"duration" bson:"duration"`
	RoshanRespawnTimer int         `json:"roshan_respawn_timer" bson:"roshan_respawn"`
	Radiant            RadiantJSON `json:"radiant" bson:"radiant"`
	Dire               DireJSON    `json:"dire" bson:"dire"`
}

type RadiantJSON struct {
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

type DireJSON struct {
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
	Level            int     `json:"Level" bson:"Level"`
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
