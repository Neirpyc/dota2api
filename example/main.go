package main

import (
	"encoding/json"
	"fmt"
	"github.com/Neirpyc/dota2api"
)

func main() {
	dota2, err := dota2api.LoadConfigFromFile("../config.yaml")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("#### Steam id  ####")
	fmt.Println()
	steamId, err := dota2.ResolveVanityUrl("Dendi")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(steamId)

	fmt.Println()
	fmt.Println("#### Account id ####")
	fmt.Println()

	matchHistory, err := dota2.GetMatchHistory(dota2api.HeroId(42), dota2api.MatchesRequested(1))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("#### Match History ####")
	fmt.Println()

	matchHistoryObject, _ := json.Marshal(matchHistory)
	fmt.Println(string(matchHistoryObject))

	matchDetails, err := dota2.GetMatchDetails(dota2api.MatchId(2606807053))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("#### Match Details ####")
	fmt.Println()
	res21, _ := json.Marshal(matchDetails)
	fmt.Println(string(res21))
	players, err := dota2.GetPlayerSummaries(dota2api.ParameterSteamIds(dota2api.NewSteamIdFrom64(76561198058479208)))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("#### Player Summaries ####")
	fmt.Println()

	playerSymmaryObject, _ := json.Marshal(players)
	fmt.Println(string(playerSymmaryObject))
	//
	heroes, err := dota2.GetHeroes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("#### Items ####")
	fmt.Println()
	fmt.Println(heroes)

	//
	friends, err := dota2.GetFriendList(dota2api.ParameterSteamId(dota2api.NewSteamIdFrom64(76561198090402335)))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("#### Friend List ####")
	fmt.Println()

	friendObject, _ := json.Marshal(friends)
	fmt.Println(string(friendObject))

	leagueList, err := dota2.GetLeagueListing()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("#### League Listing ####")
	fmt.Println()
	leagueListObject, _ := json.Marshal(leagueList)
	fmt.Println(string(leagueListObject))

	liveGames, err := dota2.GetLiveLeagueGames()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("#### Live Games ####")
	fmt.Println()
	liveGameObject, _ := json.Marshal(liveGames.Result.Games)
	fmt.Println(string(liveGameObject))

}
