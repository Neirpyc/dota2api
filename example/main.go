package main

import (
	"fmt"
	"github.com/l2x/dota2api"
)

func main() {
	dota2, err := dota2api.LoadConfig("../config.ini")

	if err != nil {
		fmt.Println(err)
		return
	}

	steamId, err := dota2.ResolveVanityUrl("Dendi")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(steamId)

	accountId := dota2.GetAccountId(steamId)
	fmt.Println(accountId)

	param := map[string]interface{}{
		"account_id": accountId,
	}
	matchHistory, err := dota2.GetMatchHistory(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matchHistory)

	matchDetails, err := dota2.GetMatchDetails(786250116)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(matchDetails)

	steamIds := []int64{
		76561198058479208,
	}
	players, err := dota2.GetPlayerSummaries(steamIds)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(players)

	heroes, err := dota2.GetHeroes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(heroes)

	friends, err := dota2.GetFriendList(76561198090402335)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(friends)

	team, err := dota2.GetTeamInfoByTeamID(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(team)

}
