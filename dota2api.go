package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

//go:generate ./genIterators -p dota2api -i genIterators.yaml -o iterators.go

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

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

	// dota2 cdn
	dota2CDN string `yaml:"Dota2CDN"`

	// api version
	dota2ApiVersion string

	// http request timeout
	timeout int

	dota2MatchUrl string
	dota2EconUrl  string
	steamUserUrl  string

	//Caching
	heroesCache *getHeroesCache
	itemsCache  *getItemsCache

	//http client
	client HTTPClient
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
	resp, err := d.Get(url)
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
	resp, err := d.Get(url)
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
