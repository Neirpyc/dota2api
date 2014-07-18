package dota2api

var (
	// steam api url
	SteamApi = "https://api.steampowered.com"

	// steam api key: http://steamcommunity.com/dev/apikey
	SteamApiKey = ""

	// dota2 name in api
	DotaName = "IDOTA2Match_570"

	// api version
	ApiVersion = "V001"

	// convert 64-bit steamID to 32-bit steamID
	// STEAMID64 - 76561197960265728 = STEAMID32
	ConvertInt = 76561197960265728

	// http request timeout
	Timeout = 10
)

func LoadConfig(file string) error {

	return nil
}
