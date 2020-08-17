package dota2api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (api Dota2) getResolveVanityUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.steamUserUrl, "ResolveVanityURL", api.steamApiVersion)
}

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
}

//Get steamId by username
func (api Dota2) ResolveVanityUrl(params ...Parameter) (SteamId, error) {
	var steamId SteamId

	param, err := getParameterMap([]parameterKind{parameterVanityUrl}, nil, params)
	if err != nil {
		return steamId, err
	}
	param["key"] = api.steamApiKey

	url, err := parseUrl(api.getResolveVanityUrl(), param)
	if err != nil {
		return steamId, err
	}
	resp, err := api.Get(url)
	if err != nil {
		return steamId, err
	}

	vanity := Vanity{}
	err = json.Unmarshal(resp, &vanity)
	if err != nil {
		return steamId, err
	}

	if vanity.Response.Success != 1 {
		return steamId, statusCodeError(vanity.Response.Success, 1)
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
