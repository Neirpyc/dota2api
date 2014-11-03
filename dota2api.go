package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Dota2 struct {
	// steam api url
	SteamApi string
	// steam api key: http://steamcommunity.com/dev/apikey
	SteamApiKey string
	//Steam User
	SteamUser string
	// api version
	SteamApiVersion string

	// dota2 name in api
	Dota2Match string
	Dota2Econ  string

	// api version
	Dota2ApiVersion string

	// convert 64-bit steamID to 32-bit steamID
	// STEAMID64 - 76561197960265728 = STEAMID32
	ConvertInt int64

	// http request timeout
	Timeout int

	Dota2MatchUrl string
	Dota2EconUrl  string
	SteamUserUrl  string
}

//Get steamId by username
func (d *Dota2) ResolveVanityUrl(vanityurl string) (int64, error) {
	var steamId int64

	param := map[string]interface{}{
		"key":       d.SteamApiKey,
		"vanityurl": vanityurl,
	}
	url, err := parseUrl(getResolveVanityUrl(d), param)
	if err != nil {
		return steamId, err
	}
	resp, err := Get(url)
	if err != nil {
		return steamId, err
	}

	vanity := Vanity{}
	err = json.Unmarshal(resp, &vanity)
	if err != nil {
		return steamId, err
	}

	if vanity.Response.Success != 1 {
		return steamId, errors.New(string(resp))
	}

	steamId, err = strconv.ParseInt(vanity.Response.SteamId, 10, 64)
	if err != nil {
		return steamId, err
	}

	return steamId, nil
}

//Get match history
func (d *Dota2) GetMatchHistory(param map[string]interface{}) (MatchHistory, error) {
	var matchHistory MatchHistory

	param["key"] = d.SteamApiKey

	url, err := parseUrl(getMatchHistoryUrl(d), param)
	if err != nil {
		return matchHistory, err
	}
	resp, err := Get(url)
	if err != nil {
		return matchHistory, err
	}

	err = json.Unmarshal(resp, &matchHistory)
	if err != nil {
		return matchHistory, err
	}
	if matchHistory.Result.Status != 1 {
		return matchHistory, errors.New(string(resp))
	}

	return matchHistory, nil
}

//Get match history by sequence num
func (d *Dota2) GetMatchHistoryBySequenceNum(param map[string]interface{}) (MatchHistory, error) {
	var matchHistory MatchHistory

	param["key"] = d.SteamApiKey

	url, err := parseUrl(getMatchHistoryBySequenceNumUrl(d), param)
	if err != nil {
		return matchHistory, err
	}
	resp, err := Get(url)
	if err != nil {
		return matchHistory, err
	}

	err = json.Unmarshal(resp, &matchHistory)
	if err != nil {
		return matchHistory, err
	}
	if matchHistory.Result.Status != 1 {
		return matchHistory, errors.New(string(resp))
	}

	return matchHistory, nil
}

//Get match details
func (d *Dota2) GetMatchDetails(matchId int64) (MatchDetails, error) {

	var matchDetails MatchDetails

	param := map[string]interface{}{
		"key":      d.SteamApiKey,
		"match_id": matchId,
	}
	url, err := parseUrl(getMatchDetailsUrl(d), param)

	if err != nil {
		return matchDetails, err
	}
	resp, err := Get(url)
	if err != nil {
		return matchDetails, err
	}

	err = json.Unmarshal(resp, &matchDetails)
	if err != nil {
		return matchDetails, err
	}

	if matchDetails.Result.Error != "" {
		return matchDetails, errors.New(string(resp))
	}

	return matchDetails, nil
}

//Get player summaries
func (d *Dota2) GetPlayerSummaries(steamIds []int64) ([]Player, error) {
	var playerSummaries PlayerSummaries
	var players []Player

	param := map[string]interface{}{
		"key":      d.SteamApiKey,
		"steamids": strings.Join(ArrayIntToStr(steamIds), ","),
	}
	url, err := parseUrl(getPlayerSummariesUrl(d), param)

	if err != nil {
		return players, err
	}
	resp, err := Get(url)
	if err != nil {
		return players, err
	}

	err = json.Unmarshal(resp, &playerSummaries)
	if err != nil {
		return players, err
	}

	players = playerSummaries.Response.Players.Player
	return players, nil
}

//Get all heroes
func (d *Dota2) GetHeroes() ([]Hero, error) {
	var heroList Heroes
	var heroes []Hero

	param := map[string]interface{}{
		"key": d.SteamApiKey,
	}
	url, err := parseUrl(getHeroesUrl(d), param)

	if err != nil {
		return heroes, err
	}
	resp, err := Get(url)
	if err != nil {
		return heroes, err
	}

	err = json.Unmarshal(resp, &heroList)
	if err != nil {
		return heroes, err
	}

	heroes = heroList.Result.Heroes

	return heroes, nil
}

//Get friend list
func (d *Dota2) GetFriendList(steamid int64) ([]Friend, error) {
	var friendList FriendList
	var friends []Friend

	param := map[string]interface{}{
		"key":     d.SteamApiKey,
		"steamid": steamid,
	}
	url, err := parseUrl(getFriendListUrl(d), param)

	if err != nil {
		return friends, err
	}
	resp, err := Get(url)
	if err != nil {
		return friends, err
	}

	err = json.Unmarshal(resp, &friendList)
	if err != nil {
		return friends, err
	}

	friends = friendList.Friendslist.Friends

	return friends, nil
}

func (d *Dota2) GetLeagueListing() {
	fmt.Println()
}

func (d *Dota2) GetLiveLeagueGames() {}

//Get team info by teamId
func (d *Dota2) GetTeamInfoByTeamID(teamId int64) (Team, error) {
	var teamInfo TeamInfo
	var team Team

	param := map[string]interface{}{
		"key": d.SteamApiKey,
		"start_at_team_id=1753464": teamId,
		"teams_requested":          1,
	}
	url, err := parseUrl(getTeamInfoByTeamID(d), param)

	if err != nil {
		return team, err
	}
	resp, err := Get(url)
	if err != nil {
		return team, err
	}

	err = json.Unmarshal(resp, &teamInfo)
	if err != nil {
		return team, err
	}

	if teamInfo.Result.Status != 1 {
		return team, errors.New(string(resp))
	}

	if len(teamInfo.Result.Teams) > 0 {
		team = teamInfo.Result.Teams[0]
	} else {
		return team, errors.New("Teams > 1")
	}

	return team, nil
}

//Get all team info
func (d *Dota2) GetAllTeamInfo() ([]Team, error) {
	var teamInfo TeamInfo
	var team []Team

	param := map[string]interface{}{
		"key": d.SteamApiKey,
	}
	url, err := parseUrl(getTeamInfoByTeamID(d), param)

	if err != nil {
		return team, err
	}
	resp, err := Get(url)
	if err != nil {
		return team, err
	}

	err = json.Unmarshal(resp, &teamInfo)
	if err != nil {
		return team, err
	}

	team = teamInfo.Result.Teams

	return team, nil
}

func (d *Dota2) GetTournamentPrizePool() {}

func (d *Dota2) GetGameItems() {}

//Convert 64-bit steamId to 32-bit steamId
func (d *Dota2) GetAccountId(steamId int64) int64 {
	return steamId - ConvertInt
}
