package dota2api

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"time"
)

//go:generate ./genIterators -p dota2api -i genIterators.yaml -o iterators.go

func statusCodeError(got int, expected int) error {
	return errors.New(fmt.Sprintf("api returned status code %d where %d was expected", got, expected))
}

func invalidParameterTypeError(got reflect.Type) error {
	return errors.New(fmt.Sprintf("unaccepted parameter type %s", got.String()))
}

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

func applyDefaultValue(value string, def string) string {
	if value == "" {
		return def
	}
	return value
}

const timeout = time.Duration(20) * time.Second

func LoadConfig(conf Config) Dota2 {
	dota2 := Dota2{
		steamApi:        applyDefaultValue(conf.SteamApi, "https://api.steampowered.com"),
		steamApiKey:     conf.SteamApiKey,
		steamUser:       applyDefaultValue(conf.SteamUser, "ISteamUser"),
		steamApiVersion: applyDefaultValue(conf.SteamApiVersion, "V001"),
		dota2Match:      applyDefaultValue(conf.Dota2Match, "IDOTA2Match_570"),
		dota2Econ:       applyDefaultValue(conf.Dota2Econ, "IEconDOTA2_570"),
		dota2CDN:        applyDefaultValue(conf.Dota2CDN, "http://cdn.dota2.com/apps/dota2/images"),
		dota2ApiVersion: applyDefaultValue(conf.Dota2ApiVersion, "V001"),
		timeout: func() int {
			if conf.Timeout != 0 {
				return conf.Timeout
			}
			return 10
		}(),
		heroesCache: &getHeroesCache{},
		itemsCache:  &getItemsCache{},
		client: &http.Client{Transport: &http.Transport{
			ResponseHeaderTimeout: timeout,
			DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, timeout)
			},
			DisableKeepAlives: true,
		},
		},
	}
	dota2.dota2MatchUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.dota2Match)
	dota2.dota2EconUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.dota2Econ)
	dota2.steamUserUrl = fmt.Sprintf("%s/%s", dota2.steamApi, dota2.steamUser)

	return dota2
}

func LoadConfigFromFile(file string) (Dota2, error) {
	settingFile, err := ioutil.ReadFile(file)
	if err != nil {
		return Dota2{}, err
	}
	conf := Config{}
	err = yaml.Unmarshal(settingFile, &conf)
	if err != nil {
		return Dota2{}, err
	}

	if conf.SteamApiKey == "" {
		return Dota2{}, errors.New("field SteamApiKey is required. Find one at http://steamcommunity.com/dev/apikey")
	}

	return LoadConfig(conf), nil
}
