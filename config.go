package dota2api

var (
	//steam api url
	SteamApi = "https://api.steampowered.com"

	//steam api key: http://steamcommunity.com/dev/apikey
	SteamApiKey = ""

	//dota2 name in api
	DotaName = "IDOTA2Match_570"

	//api version
	ApiVersion = "V001"

	// http request timeout
	Timeout = 10
)

func LoadConfig(file string) error {

	return nil
}
