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
	Status int `json:"status"`
	NumResults int `json:"num_results"`
	TotalResults int `json:"total_results"`
	ResultsRemaining int `json:"results_remaining"`
	Matches [] MatchSummary `json:"matches"`
}

type MatchSummary struct {
	SeriesID int `json:"series_id"`
	SeriesType int `json:"series_type"`
	MatchID int64 `json:"match_id"`
	MatchSeqNum int64 `json:"match_seq_num"`
	StartTime int `json:"start_time"`
	LobbyType int `json:"lobby_type"`
	RadiantTeamID int `json:"radiant_team_id"`
	DireTeamID int `json:"dire_team_id"`
	Players [] PlayerSummary `json:"players"`
}

type PlayerSummary struct {
	AccountID int `json:"account_id"`
	PlayerSlot int `json:"player_slot"`
	HeroID int `json:"hero_id"`
}

type MatchDetails struct {
	Result Match `json:"result"`
}

type Match struct {
	Error string `bson:"error" json:"error"`
	Players [] Player `json:"players"`
	RadiantWin bool `json:"radiant_win"`
	Duration int `json:"duration"`
	PreGameDuration int `json:"pre_game_duration"`
	StartTime int `json:"start_time"`
	MatchID int64 `json:"match_id"`
	MatchSeqNum int64 `json:"match_seq_num"`
	TowerStatusRadiant int `json:"tower_status_radiant"`
	TowerStatusDire int `json:"tower_status_dire"`
	BarracksStatusRadiant int `json:"barracks_status_radiant"`
	BarracksStatusDire int `json:"barracks_status_dire"`
	Cluster int `json:"cluster"`
	FirstBloodTime int `json:"first_blood_time"`
	LobbyType int `json:"lobby_type"`
	HumanPlayers int `json:"human_players"`
	Leagueid int `json:"leagueid"`
	PositiveVotes int `json:"positive_votes"`
	NegativeVotes int `json:"negative_votes"`
	GameMode int `json:"game_mode"`
	Flags int `json:"flags"`
	Engine int `json:"engine"`
	RadiantScore int `json:"radiant_score"`
	DireScore int `json:"dire_score"`
	TournamentID int `json:"tournament_id"`
	TournamentRound int `json:"tournament_round"`
	RadiantTeamID int `json:"radiant_team_id"`
	RadiantName string `json:"radiant_name"`
	RadiantLogo int `json:"radiant_logo"`
	RadiantTeamComplete int `json:"radiant_team_complete"`
	DireTeamID int `json:"dire_team_id"`
	DireName string `json:"dire_name"`
	DireLogo int `json:"dire_logo"`
	DireTeamComplete int `json:"dire_team_complete"`
	RadiantCaptain int `json:"radiant_captain"`
	DireCaptain int `json:"dire_captain"`
    PicksBans [] PicksBans `json:"picks_bans"`
}

type PicksBans struct {
	IsPick bool `json:"is_pick"`
	HeroID int `json:"hero_id"`
	Team int `json:"team"`
	Order int `json:"order"`
}

type Player struct {
	AccountID int `json:"account_id"`
	PlayerSlot int `json:"player_slot"`
	HeroID int `json:"hero_id"`
	Item0 int `json:"item_0"`
	Item1 int `json:"item_1"`
	Item2 int `json:"item_2"`
	Item3 int `json:"item_3"`
	Item4 int `json:"item_4"`
	Item5 int `json:"item_5"`
	Kills int `json:"kills"`
	Deaths int `json:"deaths"`
	Assists int `json:"assists"`
	LeaverStatus int `json:"leaver_status"`
	LastHits int `json:"last_hits"`
	Denies int `json:"denies"`
	GoldPerMin int `json:"gold_per_min"`
	XpPerMin int `json:"xp_per_min"`
	Level int `json:"level"`
	Gold int `json:"gold"`
	GoldSpent int `json:"gold_spent"`
	HeroDamage int `json:"hero_damage"`
	TowerDamage int `json:"tower_damage"`
	HeroHealing int `json:"hero_healing"`
	AbilityUpgrades [] AbilityUpgrade `json:"ability_upgrades"`
}

type AbilityUpgrade struct {
	Ability int `json:"ability"`
	Level   int `json:"level"`
	Time    int `json:"time"`
}

type PlayerSummaries struct {
	Response struct {
		 Players [] PlayerAccount `json:"players"`
	 } `json:"response"`
}


type PlayerAccount struct {
	Steamid string `json:"steamid"`
	Communityvisibilitystate int `json:"communityvisibilitystate"`
	Profilestate int `json:"profilestate"`
	Personaname string `json:"personaname"`
	Lastlogoff int `json:"lastlogoff"`
	Profileurl string `json:"profileurl"`
	Avatar string `json:"avatar"`
	Avatarmedium string `json:"avatarmedium"`
	Avatarfull string `json:"avatarfull"`
	Personastate int `json:"personastate"`
	Realname string `json:"realname"`
	Primaryclanid string `json:"primaryclanid"`
	Timecreated int `json:"timecreated"`
	Personastateflags int `json:"personastateflags"`
	Gameextrainfo string `json:"gameextrainfo"`
	Gameid string `json:"gameid"`
	Loccountrycode string `json:"loccountrycode"`
	Locstatecode string `json:"locstatecode"`
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

type LeagueList struct {
	Result struct {
	   Leagues []struct {
		   Name string `json:"name"`
		   Leagueid int `json:"leagueid"`
		   Description string `json:"description"`
		   TournamentURL string `json:"tournament_url"`
		   Itemdef int `json:"itemdef"`
	   } `json:"leagues"`
   	} `json:"result"`
}

type LiveGames struct {
	Result struct {
	   Games []Game `json:"games"`
	   Status int `json:"status"`
	} `json:"result"`
}


type Game struct {
	Players []struct {
		AccountID int `json:"account_id"`
		Name string `json:"name"`
		HeroID int `json:"hero_id"`
		Team int `json:"team"`
	} `json:"players"`
	LobbyID int64 `json:"lobby_id"`
	MatchID int64 `json:"match_id"`
	Spectators int `json:"spectators"`
	SeriesID int `json:"series_id"`
	GameNumber int `json:"game_number"`
	LeagueID int `json:"league_id"`
	StreamDelayS int `json:"stream_delay_s"`
	RadiantSeriesWins int `json:"radiant_series_wins"`
	DireSeriesWins int `json:"dire_series_wins"`
	SeriesType int `json:"series_type"`
	LeagueSeriesID int `json:"league_series_id"`
	LeagueGameID int `json:"league_game_id"`
	StageName string `json:"stage_name"`
	LeagueTier int `json:"league_tier"`
	Scoreboard Scoreboard `json:"scoreboard"`
	DireTeam PlayersTeam `json:"dire_team,omitempty"`
	RadiantTeam PlayersTeam `json:"radiant_team,omitempty"`
}

type RadiantTeam struct {
	TeamName string `json:"team_name"`
	TeamID int `json:"team_id"`
	TeamLogo int64 `json:"team_logo"`
	Complete bool `json:"complete"`
}

type PlayersTeam struct {
	TeamName string `json:"team_name"`
	TeamID int `json:"team_id"`
	TeamLogo int64 `json:"team_logo"`
	Complete bool `json:"complete"`
}

type Scoreboard struct {
	Duration float64 `json:"duration"`
	RoshanRespawnTimer int `json:"roshan_respawn_timer"`
	Radiant Radiant `json:"radiant"`
	Dire Dire `json:"dire"`
}


type Radiant struct {
	Score int `json:"score"`
	TowerState int `json:"tower_state"`
	BarracksState int `json:"barracks_state"`
	Picks []struct {
		HeroID int `json:"hero_id"`
	} `json:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id"`
	} `json:"bans"`
	Players []LivePlayer `json:"players"`
	Abilities []struct {
		AbilityID int `json:"ability_id"`
		AbilityLevel int `json:"ability_level"`
	} `json:"abilities"`
}


type Dire struct {
	Score int `json:"score"`
	TowerState int `json:"tower_state"`
	BarracksState int `json:"barracks_state"`
	Picks []struct {
		HeroID int `json:"hero_id"`
	} `json:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id"`
	} `json:"bans"`
	Players []LivePlayer `json:"players"`
	Abilities []struct {
		AbilityID int `json:"ability_id"`
		AbilityLevel int `json:"ability_level"`
	} `json:"abilities"`
}

type LivePlayer struct {
	PlayerSlot int `json:"player_slot"`
	AccountID int `json:"account_id"`
	HeroID int `json:"hero_id"`
	Kills int `json:"kills"`
	Death int `json:"death"`
	Assists int `json:"assists"`
	LastHits int `json:"last_hits"`
	Denies int `json:"denies"`
	Gold int `json:"gold"`
	Level int `json:"level"`
	GoldPerMin int `json:"gold_per_min"`
	XpPerMin int `json:"xp_per_min"`
	UltimateState int `json:"ultimate_state"`
	UltimateCooldown int `json:"ultimate_cooldown"`
	Item0 int `json:"item0"`
	Item1 int `json:"item1"`
	Item2 int `json:"item2"`
	Item3 int `json:"item3"`
	Item4 int `json:"item4"`
	Item5 int `json:"item5"`
	RespawnTimer int `json:"respawn_timer"`
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
	NetWorth int `json:"net_worth"`
}