package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Dota2 struct {
}

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

func (d *Dota2) GetMatchHistory(accountId int64) {

	fmt.Println()
}

func (d *Dota2) GetMatchDetails(matchId int64) {

}

func (d *Dota2) GetPlayerSummaries(steamIds []int64) {

}

func (d *Dota2) GetLeagueListing() {}

func (d *Dota2) GetLiveLeagueGames() {}

func (d *Dota2) GetTeamInfoByTeamID() {}

func (d *Dota2) GetHeroes() {}

func (d *Dota2) GetTournamentPrizePool() {}

func (d *Dota2) GetGameItems() {}
