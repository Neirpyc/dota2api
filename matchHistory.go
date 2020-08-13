package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

const (
	LobbyInvalid = iota - 1
	LobbyPublicMatchmaking
	LobbyPractice
	LobbyTournament
	LobbyTutorial
	LobbyCoopWithAI
	LobbyTeamMatch
	LobbySoloQueue
	LobbyRankedMatchmaking
	LobbySoloMid1vs1
)

func getMatchHistoryUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetMatchHistory", dota2.dota2ApiVersion)
}

func getMatchHistoryBySequenceNumUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetMatchHistoryBySequenceNum", dota2.dota2ApiVersion)
}

type MatchHistoryJSON struct {
	Result matchHistoryResultJSON `json:"result"`
}

type matchHistoryResultJSON struct {
	Status           int                `json:"status" bson:"status"`
	NumResults       int                `json:"num_results" bson:"num_results"`
	TotalResults     int                `json:"total_results" bson:"total_results"`
	ResultsRemaining int                `json:"results_remaining" bson:"results_remaining"`
	Matches          []matchSummaryJSON `json:"matches" bson:"matches"`
}

type matchSummaryJSON struct {
	MatchID       int64               `json:"match_id" bson:"match_id"`
	MatchSeqNum   int64               `json:"match_seq_num" bson:"match_seq_num"`
	StartTime     int64               `json:"start_time" bson:"start_time"`
	LobbyType     int                 `json:"lobby_type" bson:"lobby_type"`
	RadiantTeamID int                 `json:"radiant_team_id" bson:"radiant_team_id"`
	DireTeamID    int                 `json:"dire_team_id" bson:"dire_team_id"`
	Players       []playerSummaryJSON `json:"players" bson:"players"`
}

func (m MatchHistoryJSON) toMatchSummary(d *Dota2) (MatchHistory, error) {
	var res MatchHistory

	res.Matches = make([]MatchSummary, len(m.Result.Matches))
	for i, src := range m.Result.Matches {
		res.Matches[i].LobbyType = LobbyType(src.LobbyType)
		res.Matches[i].StartTime = time.Unix(src.StartTime, 0)
		res.Matches[i].MatchId = src.MatchID
		res.Matches[i].MatchSeqNum = src.MatchSeqNum
		res.Matches[i].Radiant = Team{
			Id: src.RadiantTeamID,
		}
		res.Matches[i].Dire = Team{
			Id: src.DireTeamID,
		}
		heroes, err := d.GetHeroes()
		if err != nil {
			return res, err
		}
		for _, p := range src.Players {
			var h Hero
			var found bool
			if p.HeroID == 0 {
				h.ID = 0
			} else if h, found = heroes.GetById(p.HeroID); !found {
				return res, errors.New("hero ID not found")
			}
			player := Player{
				AccountId: p.AccountID,
				Hero:      h,
			}
			if p.PlayerSlot&128 > 0 {
				res.Matches[i].Dire.players = append(res.Matches[i].Dire.players, player)
			} else {
				res.Matches[i].Radiant.players = append(res.Matches[i].Radiant.players, player)
			}
		}
	}
	return res, nil
}

type playerSummaryJSON struct {
	AccountID  int `json:"account_id" bson:"account_id"`
	PlayerSlot int `json:"player_slot" bson:"player_slot"`
	HeroID     int `json:"hero_id" bson:"hero_id"`
}

type MatchHistory struct {
	Matches []MatchSummary
}

func (m MatchHistory) Count() int {
	return len(m.Matches)
}

type MatchSummary struct {
	MatchId     int64
	MatchSeqNum int64
	StartTime   time.Time
	LobbyType   LobbyType
	Radiant     Team
	Dire        Team
}

func (m MatchSummary) GetPlayer(pos int) (Player, bool) {
	if p, found := m.Radiant.GetPlayer(pos); found {
		return p, true
	}
	if p, found := m.Dire.GetPlayer(pos - len(m.Radiant.players)); found {
		return p, true
	}
	return Player{}, false
}

func (m MatchSummary) GetByHero(hero Hero) (Player, bool) {
	return m.GetByHeroId(hero.ID)
}

func (m MatchSummary) PlayerCount() int {
	return m.Radiant.Count() + m.Dire.Count()
}

func (m MatchSummary) GetByHeroId(id int) (Player, bool) {
	if p, found := m.Radiant.GetByHeroId(id); found {
		return p, true
	}
	if p, found := m.Dire.GetByHeroId(id); found {
		return p, true
	}
	return Player{}, false
}

type Player struct {
	AccountId int
	Hero      Hero
}

type Team struct {
	players []Player
	Id      int
}

func (t Team) Count() int {
	return len(t.players)
}

func (t Team) GetPlayer(pos int) (Player, bool) {
	if pos < len(t.players) {
		return t.players[pos], true
	}
	return Player{}, false
}

func (t Team) GetByHero(hero Hero) (Player, bool) {
	return t.GetByHeroId(hero.ID)
}

func (t Team) GetByHeroId(id int) (Player, bool) {
	for _, p := range t.players {
		if p.Hero.ID == id {
			return p, true
		}
	}
	return Player{}, false
}

type LobbyType int

func (l LobbyType) GetId() int {
	return int(l)
}

//todo update
func (l LobbyType) GetName() string {
	switch int(l) {
	case LobbyInvalid:
		return "Invalid"
	case LobbyPublicMatchmaking:
		return "Public Matchmaking"
	case LobbyPractice:
		return "Practice"
	case LobbyTournament:
		return "Tournament"
	case LobbyTutorial:
		return "Tutorial"
	case LobbyCoopWithAI:
		return "Co-op With AI"
	case LobbyTeamMatch:
		return "Team Match"
	case LobbySoloQueue:
		return "Solo Queue"
	case LobbyRankedMatchmaking:
		return "Ranked Matchmaking"
	case LobbySoloMid1vs1:
		return "Solo Mid 1 vs 1"
	default:
		return "Unknown"
	}
}

type Cursor struct {
	c *struct {
		begin     int64
		remaining int
	}
}

func NewCursor() Cursor {
	return Cursor{c: &struct {
		begin     int64
		remaining int
	}{
		begin:     -1,
		remaining: -1,
	}}
}

func (c Cursor) GetLastReceivedMatch() int64 {
	return c.c.begin
}

func (c Cursor) SetBegin(begin int64) {
	c.c.begin = begin
}

func (c Cursor) GetRemaining() int {
	return c.c.remaining
}

func (d *Dota2) GetMatchHistory(params ...interface{}) (MatchHistory, error) {
	var matchHistory MatchHistoryJSON
	var res MatchHistory
	var c Cursor

	parameters := make([]Parameter, 0)

	for _, p := range params {
		if reflect.TypeOf(p) == reflect.TypeOf(Cursor{}) {
			c = p.(Cursor)
		} else if v, ok := p.(Parameter); ok {
			parameters = append(parameters, v)
		} else {
			return res, errors.New("invalid parameter")
		}
	}
	param, err := getParameterMap(nil, []int{
		parameterKindHeroId,
		parameterKindMatchesRequested,
		parameterKindAccountId,
		parameterKindStartAtMatchId,
		parameterKindMinPlayers}, parameters)
	if err != nil {
		return res, err
	}
	if c.c != nil {
		c.c.begin--
		if _, f := param["start_at_match_id"]; !f {
			param["start_at_match_id"] = c.c.begin
		}
	}
	param["key"] = d.steamApiKey
	url, err := parseUrl(getMatchHistoryUrl(d), param)
	if err != nil {
		return res, err
	}
	resp, err := d.Get(url)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &matchHistory)
	if err != nil {
		return res, err
	}
	if matchHistory.Result.Status != 1 {
		return res, errors.New(string(resp))
	}
	if c.c != nil {
		if matchHistory.Result.NumResults > 0 {
			c.c.begin = matchHistory.Result.Matches[matchHistory.Result.NumResults-1].MatchID
		}
		c.c.remaining = matchHistory.Result.ResultsRemaining
	}

	res, err = matchHistory.toMatchSummary(d)

	return res, err
}

func (d *Dota2) GetMatchHistoryBySequenceNum(params ...interface{}) (MatchHistory, error) {
	var matchHistory MatchHistoryJSON
	var res MatchHistory
	var c Cursor

	parameters := make([]Parameter, 0)

	for _, p := range params {
		if reflect.TypeOf(p) == reflect.TypeOf(Cursor{}) {
			c = p.(Cursor)
		} else if v, ok := p.(Parameter); ok {
			parameters = append(parameters, v)
		} else {
			return res, errors.New("invalid parameter")
		}
	}
	param, err := getParameterMap(nil, []int{
		parameterStartMatchAtSeqNum,
		parameterKindMatchesRequested}, parameters)
	if err != nil {
		return res, err
	}
	if c.c != nil {
		c.c.begin++
		if _, f := param["start_at_match_seq_num"]; !f {
			param["start_at_match_seq_num"] = c.c.begin
		}
	}
	param["key"] = d.steamApiKey
	url, err := parseUrl(getMatchHistoryBySequenceNumUrl(d), param)
	if err != nil {
		return res, err
	}
	resp, err := d.Get(url)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &matchHistory)
	if err != nil {
		return res, err
	}
	if matchHistory.Result.Status != 1 {
		return res, errors.New(string(resp))
	}
	if c.c != nil {
		if len(matchHistory.Result.Matches) > 0 {
			c.c.begin = matchHistory.Result.Matches[len(matchHistory.Result.Matches)-1].MatchSeqNum
		}
	}

	res, err = matchHistory.toMatchSummary(d)

	return res, err
}

func HeroId(id int) ParameterInt {
	return ParameterInt{
		k:       "hero_id",
		v:       id,
		kindInt: parameterKindHeroId,
	}
}

func MatchesRequested(num int) ParameterInt {
	return ParameterInt{
		k:       "matches_requested",
		v:       num,
		kindInt: parameterKindMatchesRequested,
	}
}

func AccountId(id int64) ParameterInt {
	return ParameterInt{
		k:       "account_id",
		v:       int(int32(id)),
		kindInt: parameterKindAccountId,
	}
}

func StartAtMatchId(id int64) ParameterInt64 {
	return ParameterInt64{
		k:       "start_at_match_id",
		v:       id,
		kindInt: parameterKindStartAtMatchId,
	}
}

func MinPlayers(id int) ParameterInt {
	return ParameterInt{
		k:       "min_players",
		v:       id,
		kindInt: parameterKindMinPlayers,
	}
}

func StartAtMatchSeqNum(id int64) ParameterInt64 {
	return ParameterInt64{
		k:       "start_at_match_seq_num",
		v:       id,
		kindInt: parameterStartMatchAtSeqNum,
	}
}
