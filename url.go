package dota2api

import (
	"fmt"
)

func getResolveVanityUrl() string {

	return fmt.Sprintf("%s/%s/%s/", SteamUserUrl, "ResolveVanityURL", SteamApiVersion)
}

func getMatchHistoryUrl() string {

	return fmt.Sprintf("%s/%s/%s/", Dota2MatchUrl, "GetMatchHistory", Dota2ApiVersion)
}

func getMatchDetailsUrl() string {

	return fmt.Sprintf("%s/%s/%s/", Dota2MatchUrl, "GetMatchDetails", Dota2ApiVersion)
}

func getPlayerSummariesUrl() string {

	return fmt.Sprintf("%s/%s/%s/", SteamUserUrl, "GetPlayerSummaries", SteamApiVersion)
}

func getHeroesUrl() string {

	return fmt.Sprintf("%s/%s/%s/", Dota2EconUrl, "GetHeroes", Dota2ApiVersion)
}

func getFriendListUrl() string {

	return fmt.Sprintf("%s/%s/%s/", SteamUserUrl, "GetFriendList", SteamApiVersion)
}

func getTeamInfoByTeamID() string {

	return fmt.Sprintf("%s/%s/%s/", Dota2MatchUrl, "GetTeamInfoByTeamID", Dota2ApiVersion)
}
