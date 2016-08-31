package dota2api

import (
	"fmt"
)

func getResolveVanityUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.SteamUserUrl, "ResolveVanityURL", dota2.SteamApiVersion)
}

func getMatchHistoryUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2MatchUrl, "GetMatchHistory", dota2.Dota2ApiVersion)
}

func getMatchHistoryBySequenceNumUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2MatchUrl, "GetMatchHistoryBySequenceNum", dota2.Dota2ApiVersion)
}

func getMatchDetailsUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2MatchUrl, "GetMatchDetails", dota2.Dota2ApiVersion)
}

func getPlayerSummariesUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.SteamUserUrl, "GetPlayerSummaries", "V002")
}

func getHeroesUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2EconUrl, "GetHeroes", dota2.Dota2ApiVersion)
}

func getFriendListUrl(dota2 *Dota2) string {

	return fmt.Sprintf("%s/%s/%s/", dota2.SteamUserUrl, "GetFriendList", dota2.SteamApiVersion)
}

func getLeagueListUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2MatchUrl, "GetLeagueListing", dota2.SteamApiVersion)
}

func getLiveGamesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2MatchUrl, "GetLiveLeagueGames", dota2.SteamApiVersion)
}


