package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Dota2 struct {
}

//Get steamId by username
func (d *Dota2) ResolveVanityUrl(vanityurl string) (int64, error) {
	var steamId int64

	param := map[string]interface{}{
		"key":       SteamApiKey,
		"vanityurl": vanityurl,
	}
	url, err := parseUrl(getResolveVanityUrl(), param)
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

	param["key"] = SteamApiKey

	url, err := parseUrl(getMatchHistoryUrl(), param)
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
		"key":      SteamApiKey,
		"match_id": matchId,
	}
	url, err := parseUrl(getMatchDetailsUrl(), param)

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
		"key":      SteamApiKey,
		"steamids": strings.Join(ArrayIntToStr(steamIds), ","),
	}
	url, err := parseUrl(getPlayerSummariesUrl(), param)

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
		"key": SteamApiKey,
	}
	url, err := parseUrl(getHeroesUrl(), param)

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
		"key":     SteamApiKey,
		"steamid": steamid,
	}
	url, err := parseUrl(getFriendListUrl(), param)

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

func (d *Dota2) GetTeamInfoByTeamID() {}

func (d *Dota2) GetTournamentPrizePool() {}

func (d *Dota2) GetGameItems() {}

//Convert 64-bit steamId to 32-bit steamId
func (d *Dota2) GetAccountId(steamId int64) int64 {
	return steamId - ConvertInt
}
