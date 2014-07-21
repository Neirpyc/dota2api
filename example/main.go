package main

import (
	"fmt"
	"github.com/l2x/dota2api"
)

func main() {
	dota2, err := dota2api.LoadConfig("../config.ini")

	if err != nil {
		fmt.Println(err)
	}

	steamId, err := dota2.ResolveVanityUrl("specode")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(steamId)

	accountId := dota2.GetAccountId(steamId)
	param := map[string]interface{}{
		"account_id": accountId,
	}
	matchHistory, err := dota2.GetMatchHistory(param)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matchHistory)

	matchDetails, err := dota2.GetMatchDetails(786250116)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matchDetails)

	steamIds := []int64{
		76561198058479208,
	}
	players, err := dota2.GetPlayerSummaries(steamIds)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(players)

	heroes, err := dota2.GetHeroes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(heroes)

	friends, err := dota2.GetFriendList(76561198090402335)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(friends)

}
