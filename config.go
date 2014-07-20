package dota2api

import (
	"errors"
	"fmt"
	"github.com/revel/config"
)

var (
	// steam api url
	SteamApi = "https://api.steampowered.com"
	// steam api key: http://steamcommunity.com/dev/apikey
	SteamApiKey = ""
	//Steam User
	SteamUser = "ISteamUser"
	// api version
	SteamApiVersion = "V001"

	// dota2 name in api
	Dota2Match = "IDOTA2Match_570"
	Dota2Econ  = "IEconDOTA2_570"

	// api version
	Dota2ApiVersion = "V001"

	// convert 64-bit steamID to 32-bit steamID
	// STEAMID64 - 76561197960265728 = STEAMID32
	ConvertInt int64 = 76561197960265728

	// http request timeout
	Timeout = 10

	Dota2MatchUrl = ""
	Dota2EconUrl  = ""
	SteamUserUrl  = ""
)

func LoadConfig(file string) (Dota2, error) {
	dota2 := Dota2{}

	config, err := config.ReadDefault(file)
	if err != nil {
		return dota2, err
	}

	SteamApiKey, err = config.String("steam", "steamApiKey")
	if err != nil {
		return dota2, err
	}
	if SteamApiKey == "" {
		return dota2, errors.New("SteamApiKey is empty.[http://steamcommunity.com/dev/apikey]")
	}

	SteamApi, err = config.String("steam", "steamApi")
	if err != nil {
		return dota2, err
	}

	SteamUser, err = config.String("steam", "steamUser")
	if err != nil {
		return dota2, err
	}

	SteamApiVersion, err = config.String("steam", "steamApiVersion")
	if err != nil {
		return dota2, err
	}

	Dota2Match, err = config.String("dota2", "dota2Match")
	if err != nil {
		return dota2, err
	}

	Dota2Econ, err = config.String("dota2", "dota2Econ")
	if err != nil {
		return dota2, err
	}

	Dota2ApiVersion, err = config.String("dota2", "dota2ApiVersion")
	if err != nil {
		return dota2, err
	}

	timeout, err := config.Int("", "timeout")
	if err == nil {
		Timeout = timeout
	}

	Dota2MatchUrl = fmt.Sprintf("%s/%s", SteamApi, Dota2Match)
	Dota2EconUrl = fmt.Sprintf("%s/%s", SteamApi, Dota2Econ)
	SteamUserUrl = fmt.Sprintf("%s/%s", SteamApi, SteamUser)

	return dota2, nil
}
