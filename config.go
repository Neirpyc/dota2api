package dota2api

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Timeout         int    `yaml:"Timeout"`
	SteamApiKey     string `yaml:"SteamApiKey"`
	SteamApi        string `yaml:"SteamApi"`
	SteamUser       string `yaml:"SteamUser"`
	SteamApiVersion string `yaml:"SteamApiVersion"`
	Dota2Match      string `yaml:"Dota2match"`
	Dota2Econ       string `yaml:"Dota2Econ"`
	Dota2CDN        string `yaml:"Dota2CDN"`
	Dota2ApiVersion string `yaml:"Dota2ApiVersion"`
}

func applyDefaultValue(value string, def string) string {
	if value == "" {
		return def
	}
	return value
}

func LoadConfig(file string) (Dota2, error) {
	dota2 := Dota2{}

	settingFile, err := ioutil.ReadFile(file)
	if err != nil {
		return dota2, err
	}
	conf := Config{}
	err = yaml.Unmarshal(settingFile, &conf)
	if err != nil {
		return dota2, err
	}

	dota2.steamApiKey = conf.SteamApiKey
	if dota2.steamApiKey == "" {
		return dota2, errors.New("SteamApiKey is empty.[http://steamcommunity.com/dev/apikey]")
	}

	dota2.steamApi = applyDefaultValue(conf.SteamApi, "https://api.steampowered.com")
	dota2.steamApiVersion = applyDefaultValue(conf.SteamApiVersion, "V001")
	dota2.steamUser = applyDefaultValue(conf.SteamUser, "SteamUser")
	dota2.dota2Match = applyDefaultValue(conf.Dota2Match, "IDOTA2Match_570")
	dota2.dota2Econ = applyDefaultValue(conf.Dota2Econ, "IEconDOTA2_570")
	dota2.dota2ApiVersion = applyDefaultValue(conf.Dota2ApiVersion, "V001")
	dota2.dota2CDN = applyDefaultValue(conf.Dota2CDN, "http://cdn.dota2.com/apps/dota2/images")
	dota2.timeout = conf.Timeout
	if dota2.timeout == 0 {
		dota2.timeout = 10
	}
	dota2.dota2MatchUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.dota2Match)
	dota2.dota2EconUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.dota2Econ)
	dota2.steamUserUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.steamUser)

	dota2.heroesCache = &getHeroesCache{}
	dota2.itemsCache = &getItemsCache{}

	return dota2, nil
}
