package dota2api

import (
	"fmt"
)

func getResolveVanityUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "ResolveVanityURL", dota2.steamApiVersion)
}

func getPlayerSummariesUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "GetPlayerSummaries", "V002")
}

func getFriendListUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "GetFriendList", dota2.steamApiVersion)
}

func getLeagueListUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetLeagueListing", dota2.steamApiVersion)
}

func getLiveGamesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetLiveLeagueGames", dota2.steamApiVersion)
}
