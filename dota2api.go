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
	steamApi string
	// steam api key: http://steamcommunity.com/dev/apikey
	steamApiKey string
	//Steam User
	steamUser string
	// api version
	steamApiVersion string

	// dota2 name in api
	dota2Match string
	dota2Econ  string

	// api version
	dota2ApiVersion string

	// convert 64-bit steamID to 32-bit steamID
	// STEAMID64 - 76561197960265728 = STEAMID32
	convertInt int64

	// http request timeout
	timeout int

	dota2MatchUrl string
	dota2EconUrl  string
	steamUserUrl  string

	//Caching
	heroesCache *getHeroesCache
	itemsCache  *getItemsCache
}

//Get steamId by username
func (d *Dota2) ResolveVanityUrl(vanityurl string) (int64, error) {
	var steamId int64

	param := map[string]interface{}{
		"key":       d.steamApiKey,
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

	param["key"] = d.steamApiKey

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

	param["key"] = d.steamApiKey

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
		"key":      d.steamApiKey,
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
func (d *Dota2) GetPlayerSummaries(steamIds []int64) (PlayerSummaries, error) {
	var playerSummaries PlayerSummaries
	var players PlayerSummaries

	param := map[string]interface{}{
		"key":      d.steamApiKey,
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

	players = playerSummaries
	return players, nil
}

//Get friend list
func (d *Dota2) GetFriendList(steamid int64) ([]Friend, error) {
	var friendList FriendList
	var friends []Friend

	param := map[string]interface{}{
		"key":     d.steamApiKey,
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

func (d *Dota2) GetLeagueListing() (LeagueList, error) {
	var leagueList LeagueList
	param := map[string]interface{}{
		"key": d.steamApiKey,
	}

	url, err := parseUrl(getLeagueListUrl(d), param)
	fmt.Println(url)
	if err != nil {
		return leagueList, err
	}
	resp, err := Get(url)
	if err != nil {
		return leagueList, err
	}

	err = json.Unmarshal(resp, &leagueList)
	if err != nil {
		return leagueList, err
	}
	return leagueList, nil
}

func (d *Dota2) GetLiveLeagueGames() (LiveGames, error) {
	var liveGames LiveGames
	param := map[string]interface{}{
		"key": d.steamApiKey,
	}

	url, err := parseUrl(getLiveGamesUrl(d), param)
	fmt.Println(url)
	if err != nil {
		return liveGames, err
	}
	resp, err := Get(url)
	if err != nil {
		return liveGames, err
	}

	err = json.Unmarshal(resp, &liveGames)
	if err != nil {
		return liveGames, err
	}
	return liveGames, nil
}
