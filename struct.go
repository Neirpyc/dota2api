package dota2api

type Vanity struct {
	Response VanityResp `json:"response"`
}

type VanityResp struct {
	SteamId string `json:"steamid"`
	Success int    `json:"success"`
}
