package dota2api

import (
	"encoding/json"
	"fmt"
	"time"
)

func getLiveGamesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetLiveLeagueGames", dota2.steamApiVersion)
}

type liveGamesJSON struct {
	Result struct {
		Games  []liveGameJSON `json:"games" bson:"games"`
		Status int            `json:"status" bson:"status"`
	} `json:"result" bson:"result"`
}

type liveGameJSON struct {
	Players []struct {
		AccountID int64  `json:"account_id" bson:"account_id"`
		Name      string `json:"name" bson:"name"`
		HeroID    int    `json:"hero_id" bson:"hero_id"`
		Team      int    `json:"team" bson:"team"`
	} `json:"players" bson:"players"`
	LobbyID           int64           `json:"lobby_id" bson:"lobby_id"`
	MatchID           int64           `json:"match_id" bson:"match_id"`
	Spectators        int             `json:"spectators" bson:"spectators"`
	LeagueID          int             `json:"league_id" bson:"league_id"`
	LeagueNodeId      int             `json:"league_node_id" bson:"league_node_id"`
	StreamDelayS      int             `json:"stream_delay_s" bson:"steam_delay_s"`
	RadiantSeriesWins int             `json:"radiant_series_wins" bson:"radiant_series_wins"`
	DireSeriesWins    int             `json:"dire_series_wins" bson:"dire_series_win"`
	SeriesType        int             `json:"series_type" bson:"series_type"`
	Scoreboard        scoreboardJSON  `json:"scoreboard" bson:"scoreboard"`
	DireTeam          playersTeamJSON `json:"dire_team,omitempty" bson:"dire_team"`
	RadiantTeam       playersTeamJSON `json:"radiant_team,omitempty" bson:"radiant_team"`
}

type playersTeamJSON struct {
	TeamName string `json:"team_name" bson:"team_name"`
	TeamID   int64  `json:"team_id" bson:"team_id"`
	TeamLogo int64  `json:"team_logo" bson:"team_logo"`
	Complete bool   `json:"complete" bson:"complete"`
}

type scoreboardJSON struct {
	Duration           float64  `json:"duration" bson:"duration"`
	RoshanRespawnTimer int      `json:"roshan_respawn_timer" bson:"roshan_respawn"`
	Radiant            sideJSON `json:"radiant" bson:"radiant"`
	Dire               sideJSON `json:"dire" bson:"dire"`
}

type sideJSON struct {
	Score         int `json:"score" bson:"score"`
	TowerState    int `json:"tower_state" bson:"tower_state"`
	BarracksState int `json:"barracks_state" bson:"barracks_state"`
	Picks         []struct {
		HeroID int `json:"hero_id" bson:"hero_id"`
	} `json:"picks" bson:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id" bson:"hero_id"`
	} `json:"bans" bson:"bans"`
	Players   []livePlayerJSON `json:"players" bson:"players"`
	Abilities LiveAbilities    `json:"abilities" bson:"abilities"`
}

type livePlayerJSON struct {
	PlayerSlot       int     `json:"player_slot" bson:"player_slot"`
	AccountID        int64   `json:"account_id" bson:"account_id"`
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

type LiveGame struct {
	Players     LiveGamePlayers
	LobbyId     int64
	MatchId     int64
	Spectators  int
	League      League
	StreamDelay time.Duration
	Series      Series
	Teams       LiveTeam
	ScoreBoard  ScoreBoard
}

type LiveGames []LiveGame

type LiveGamePlayer struct {
	AccountId int64
	Name      string
	Hero      Hero
	Team      int
}

type LiveGamePlayers []LiveGamePlayer

type League struct {
	LeagueId     int
	LeagueNodeId int
}

type Series struct {
	Radiant int
	Dire    int
}

type LiveTeam struct {
	TeamName string
	TeamId   int64
	TeamLogo int64
	Complete bool
}

type ScoreBoard struct {
	Duration           time.Duration
	RoshanRespawnTimer time.Duration
	Sides              Sides
}

type Sides struct {
	Radiant SideLive
	Dire    SideLive
}

type SideLive struct {
	Score          int
	BuildingsState BuildingsState
	Picks          []Hero
	Bans           []Hero
	Players        PlayersLive
	Abilities      LiveAbilities
}

type LiveTeams struct {
	Radiant LiveTeam
	Dire    LiveTeam
}

type PlayerLive struct {
	PlayerSlot   int
	AccountID    int64
	Hero         Hero
	KDA          KDA
	Stats        PlayerStatsLive
	Items        LivePlayerItems
	RespawnTimer time.Duration
	Position     Position
	NetWorth     int
}

type PlayersLive []PlayerLive

type PlayerStatsLive struct {
	LastHits      int
	Denies        int
	GoldPerMinute int
	XpPerMinute   int
	Level         int
	HeroDamage    Damage
	TowerDamage   Damage
	HeroHealing   Damage
}

type UltimateState struct {
	UltimateState    int
	UltimateCooldown time.Duration
}

type LivePlayerItems struct {
	Item0 Item
	Item1 Item
	Item2 Item
	Item3 Item
	Item4 Item
	Item5 Item
}

type Position struct {
	X, Y float64
}

type LiveAbility struct {
	AbilityID    int `json:"ability_id" bson:"ability_id"`
	AbilityLevel int `json:"ability_level" bson:"ability_level"`
}

type LiveAbilities []LiveAbility

func (d *Dota2) GetLiveLeagueGames() (liveGamesJSON, error) {
	var liveGames liveGamesJSON
	param := map[string]interface{}{
		"key": d.steamApiKey,
	}

	url, err := parseUrl(getLiveGamesUrl(d), param)
	fmt.Println(url)
	if err != nil {
		return liveGames, err
	}
	resp, err := d.Get(url)
	if err != nil {
		return liveGames, err
	}

	err = json.Unmarshal(resp, &liveGames)
	if err != nil {
		return liveGames, err
	}
	return liveGames, nil
}
