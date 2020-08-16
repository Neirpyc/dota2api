package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func getResolveVanityUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "ResolveVanityURL", dota2.steamApiVersion)
}

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
}

//Get steamId by username
func (d *Dota2) ResolveVanityUrl(params ...Parameter) (SteamId, error) {
	var steamId SteamId

	param, err := getParameterMap([]int{parameterVanityUrl}, nil, params)
	if err != nil {
		return steamId, err
	}
	param["key"] = d.steamApiKey

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

	steamId.id, err = strconv.ParseUint(vanity.Response.SteamId, 10, 64)
	if err != nil {
		return steamId, err
	}
	steamId.isId64 = true
	return steamId, nil
}

func VanityUrl(url string) ParameterString {
	return ParameterString{
		k:       "vanityurl",
		v:       url,
		kindInt: parameterVanityUrl,
	}
}
