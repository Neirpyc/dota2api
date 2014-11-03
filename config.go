package dota2api

import (
	"errors"
	"fmt"

	"github.com/revel/config"
)

var (
	// convert 64-bit steamID to 32-bit steamID
	// STEAMID64 - 76561197960265728 = STEAMID32
	ConvertInt int64 = 76561197960265728
)

type Conf struct {
	config *config.Config
}

func (c *Conf) String(section string, option string) (string, error) {
	return c.config.String(section, option)
}

func (c *Conf) StringDefault(section string, option string, defaultValue string) string {
	v, err := c.config.String(section, option)
	if err == nil {
		return v
	}
	return defaultValue
}

func (c *Conf) Int(section string, option string) (int, error) {
	return c.config.Int(section, option)
}

func (c *Conf) IntDefault(section string, option string, defaultValue int) int {
	v, err := c.config.Int(section, option)
	if err == nil {
		return v
	}
	return defaultValue
}

func LoadConfig(file string) (Dota2, error) {
	dota2 := Dota2{}
	config, err := config.ReadDefault(file)
	if err != nil {
		return dota2, err
	}
	conf := &Conf{config}

	dota2.SteamApiKey, err = conf.String("steam", "steamApiKey")
	if err != nil {
		return dota2, err
	}
	if dota2.SteamApiKey == "" {
		return dota2, errors.New("SteamApiKey is empty.[http://steamcommunity.com/dev/apikey]")
	}

	dota2.SteamApi = conf.StringDefault("steam", "steamApi", "https://api.steampowered.com")
	dota2.SteamUser = conf.StringDefault("steam", "steamUser", "SteamUser")
	dota2.SteamApiVersion = conf.StringDefault("steam", "steamApiVersion", "V001")
	dota2.Dota2Match = conf.StringDefault("dota2", "dota2Match", "IDOTA2Match_570")
	dota2.Dota2Econ = conf.StringDefault("dota2", "dota2Econ", "IEconDOTA2_570")
	dota2.Dota2ApiVersion = conf.StringDefault("dota2", "dota2ApiVersion", "V001")
	dota2.Timeout = conf.IntDefault("", "timeout", 10)

	dota2.Dota2MatchUrl = fmt.Sprintf("%s/%s", dota2.SteamApi, dota2.Dota2Match)
	dota2.Dota2EconUrl = fmt.Sprintf("%s/%s", dota2.SteamApi, dota2.Dota2Econ)
	dota2.SteamUserUrl = fmt.Sprintf("%s/%s", dota2.SteamApi, dota2.SteamUser)

	return dota2, nil
}
