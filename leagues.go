package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (api Dota2) getLiveGamesUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.dota2MatchUrl, "GetLiveLeagueGames", api.steamApiVersion)
}

type liveGamesJSON struct {
	Result struct {
		Games  []liveGameJSON `json:"games" bson:"games"`
		Status int            `json:"status" bson:"status"`
	} `json:"result" bson:"result"`
}

type liveGameJSON struct {
	Players           liveGamePlayersJSON `json:"players" bson:"players"`
	LobbyID           int64               `json:"lobby_id" bson:"lobby_id"`
	MatchID           int64               `json:"match_id" bson:"match_id"`
	Spectators        int                 `json:"spectators" bson:"spectators"`
	LeagueID          int                 `json:"league_id" bson:"league_id"`
	LeagueNodeId      int                 `json:"league_node_id" bson:"league_node_id"`
	StreamDelayS      int                 `json:"stream_delay_s" bson:"steam_delay_s"`
	RadiantSeriesWins int                 `json:"radiant_series_wins" bson:"radiant_series_wins"`
	DireSeriesWins    int                 `json:"dire_series_wins" bson:"dire_series_win"`
	SeriesType        int                 `json:"series_type" bson:"series_type"`
	Scoreboard        scoreboardJSON      `json:"scoreboard" bson:"scoreboard"`
	DireTeam          playersTeamJSON     `json:"dire_team,omitempty" bson:"dire_team"`
	RadiantTeam       playersTeamJSON     `json:"radiant_team,omitempty" bson:"radiant_team"`
}

type liveGamePlayersJSON []liveGamePlayerJSON

type liveGamePlayerJSON struct {
	AccountID int64  `json:"account_id" bson:"account_id"`
	Name      string `json:"name" bson:"name"`
	HeroID    int    `json:"hero_id" bson:"hero_id"`
	Team      int    `json:"team" bson:"team"`
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
	Score         int             `json:"score" bson:"score"`
	TowerState    uint16          `json:"tower_state" bson:"tower_state"`
	BarracksState uint8           `json:"barracks_state" bson:"barracks_state"`
	Picks         picksOrBansJSON `json:"picks" bson:"picks"`
	Bans          picksOrBansJSON `json:"bans" bson:"bans"`
	Players       livePlayersJSON `json:"players" bson:"players"`
	Abilities     LiveAbilities   `json:"abilities" bson:"abilities"`
}

type picksOrBansJSON []struct {
	HeroID int `json:"hero_id" bson:"hero_id"`
}

type livePlayersJSON []livePlayerJSON

type livePlayerJSON struct {
	PlayerSlot       int     `json:"player_slot" bson:"player_slot"`
	AccountID        int64   `json:"account_id" bson:"account_id"`
	HeroID           int     `json:"hero_id" bson:"hero_id"`
	Kills            int     `json:"kills" bson:"kills"`
	Deaths           int     `json:"death" bson:"death"`
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
	Teams       LiveTeams
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
	BuildingsState TeamBuildingsState
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
	PlayerSlot    int
	AccountID     int64
	Hero          Hero
	KDA           KDA
	Stats         PlayerStatsLive
	Items         LivePlayerItems
	RespawnTimer  time.Duration
	UltimateState UltimateState
	Position      Position
	Gold          PlayerGold
}

type PlayersLive []PlayerLive

type PlayerStatsLive struct {
	LastHits      int
	Denies        int
	GoldPerMinute int
	XpPerMinute   int
	Level         int
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

func (l liveGamesJSON) toLiveGames(api Dota2) (LiveGames, error) {
	ret := make(LiveGames, len(l.Result.Games))
	for i, g := range l.Result.Games {
		var err error
		if ret[i], err = g.toLiveGame(api); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (g liveGameJSON) toLiveGame(api Dota2) (LiveGame, error) {
	ret := LiveGame{
		LobbyId:    g.LobbyID,
		MatchId:    g.MatchID,
		Spectators: g.Spectators,
		League: League{
			LeagueId:     g.LeagueID,
			LeagueNodeId: g.LeagueNodeId,
		},
		StreamDelay: time.Duration(g.StreamDelayS) * time.Second,
		Series: Series{
			Radiant: g.RadiantSeriesWins,
			Dire:    g.DireSeriesWins,
		},
		Teams: LiveTeams{
			Radiant: LiveTeam{
				TeamName: g.RadiantTeam.TeamName,
				TeamId:   g.RadiantTeam.TeamID,
				TeamLogo: g.RadiantTeam.TeamLogo,
				Complete: g.RadiantTeam.Complete,
			},
			Dire: LiveTeam{
				TeamName: g.DireTeam.TeamName,
				TeamId:   g.DireTeam.TeamID,
				TeamLogo: g.DireTeam.TeamLogo,
				Complete: g.DireTeam.Complete,
			},
		},
	}
	var err error
	if ret.Players, err = g.Players.toPlayers(api); err != nil {
		return ret, err
	}
	ret.ScoreBoard, err = g.Scoreboard.toScoreBoard(api)
	return ret, err
}

func (p liveGamePlayersJSON) toPlayers(api Dota2) (LiveGamePlayers, error) {
	l := make(LiveGamePlayers, len(p))
	for i, player := range p {
		var err error
		if l[i], err = player.toPlayer(api); err != nil {
			return nil, err
		}
	}
	return l, nil
}

func (p liveGamePlayerJSON) toPlayer(api Dota2) (LiveGamePlayer, error) {
	ret := LiveGamePlayer{
		AccountId: p.AccountID,
		Name:      p.Name,
		Team:      p.Team,
	}

	if p.HeroID != 0 {
		if h, err := api.GetHeroes(); err != nil {
			return ret, err
		} else {
			var f bool
			if ret.Hero, f = h.GetById(p.HeroID); !f {
				return ret, errors.New("hero not found")
			}
		}
	}
	return ret, nil
}

func (s scoreboardJSON) toScoreBoard(api Dota2) (ScoreBoard, error) {
	ret := ScoreBoard{
		Duration:           time.Duration(s.Duration * float64(time.Second)),
		RoshanRespawnTimer: time.Duration(s.RoshanRespawnTimer) * time.Second,
		Sides:              Sides{},
	}
	var err error
	if ret.Sides.Radiant, err = s.Radiant.toSideLive(api); err != nil {
		return ret, err
	}
	ret.Sides.Dire, err = s.Dire.toSideLive(api)
	return ret, err
}

func (s sideJSON) toSideLive(api Dota2) (SideLive, error) {
	ret := SideLive{
		Score:          s.Score,
		BuildingsState: TeamBuildingsState{}.from(s.TowerState, s.BarracksState),
		Abilities:      s.Abilities,
	}
	var err error
	if ret.Picks, err = s.Picks.toHeroSlice(api); err != nil {
		return ret, err
	}
	if ret.Bans, err = s.Bans.toHeroSlice(api); err != nil {
		return ret, err
	}
	ret.Players, err = s.Players.toLivePlayers(api)
	return ret, err
}

func (l livePlayersJSON) toLivePlayers(api Dota2) (PlayersLive, error) {
	ret := make(PlayersLive, len(l))
	for i, player := range l {
		var err error
		if ret[i], err = player.toLivePlayer(api); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (l livePlayerJSON) toLivePlayer(api Dota2) (PlayerLive, error) {
	p := PlayerLive{
		PlayerSlot: l.PlayerSlot,
		AccountID:  l.AccountID,
		KDA: KDA{
			Kills:   l.Kills,
			Deaths:  l.Deaths,
			Assists: l.Assists,
		},
		Stats: PlayerStatsLive{
			LastHits:      l.LastHits,
			Denies:        l.Denies,
			GoldPerMinute: l.GoldPerMin,
			XpPerMinute:   l.XpPerMin,
			Level:         l.Level,
		},
		RespawnTimer: time.Duration(l.RespawnTimer) * time.Second,
		UltimateState: UltimateState{
			UltimateState:    l.UltimateState,
			UltimateCooldown: time.Duration(l.UltimateCooldown) * time.Second,
		},
		Position: Position{
			X: l.PositionX,
			Y: l.PositionY,
		},
		Gold: PlayerGold{
			current: l.Gold,
			spent:   l.NetWorth - l.Gold,
		},
	}

	var f bool

	if l.HeroID != 0 {
		h, err := api.GetHeroes()
		if err != nil {
			return PlayerLive{}, err
		}
		if p.Hero, f = h.GetById(l.HeroID); !f {
			return p, errors.New("unknown hero id")
		}
	}

	items, err := api.GetItems()
	if err != nil {
		return PlayerLive{}, err
	}
	fields := []*Item{&p.Items.Item0, &p.Items.Item1, &p.Items.Item2, &p.Items.Item3, &p.Items.Item4, &p.Items.Item5}
	src := []int{l.Item0, l.Item1, l.Item2, l.Item3, l.Item4, l.Item5}
	for i, fi := range fields {
		if src[i] != 0 {
			if *fi, f = items.GetById(src[i]); !f {
				return p, errors.New("unknown item id")
			}
		}
	}

	return p, nil
}

func (p picksOrBansJSON) toHeroSlice(api Dota2) ([]Hero, error) {
	h, err := api.GetHeroes()
	if err != nil {
		return nil, err
	}
	ret := make([]Hero, len(p))
	for i, choice := range p {
		var f bool
		if ret[i], f = h.GetById(choice.HeroID); !f {
			return nil, err
		}
	}
	return ret, nil
}

func (api Dota2) GetLiveLeagueGames(params ...Parameter) (LiveGames, error) {
	var liveGames liveGamesJSON

	param, err := getParameterMap(nil, nil, params)
	if err != nil {
		return LiveGames{}, err
	}
	param["key"] = api.steamApiKey

	url, err := parseUrl(api.getLiveGamesUrl(), param)
	fmt.Println(url)
	if err != nil {
		return LiveGames{}, err
	}
	resp, err := api.Get(url)
	if err != nil {
		return LiveGames{}, err
	}

	err = json.Unmarshal(resp, &liveGames)
	if err != nil {
		return LiveGames{}, err
	}
	if liveGames.Result.Status != 200 {
		return LiveGames{}, errors.New("non 200 status")
	}
	return liveGames.toLiveGames(api)
}
