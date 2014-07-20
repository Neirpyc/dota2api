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

}
