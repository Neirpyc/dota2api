package dota2api

import (
	"encoding/json"
	"fmt"
)

func getLeagueListUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2MatchUrl, "GetLeagueListing", dota2.steamApiVersion)
}

type LeagueList struct {
	Result struct {
		Leagues []struct {
			Name          string `json:"name" bson:"name"`
			Leagueid      int    `json:"leagueid" bson:"league_id"`
			Description   string `json:"description" bson:"description"`
			TournamentURL string `json:"tournament_url" bson:"tournament_url"`
			Itemdef       int    `json:"itemdef" bson:"item_def"`
		} `json:"leagues" bson:"leagues"`
	} `json:"result" bson:"result"`
}

func (d *Dota2) GetLeagueListing() (LeagueList, error) {
	var leagueList LeagueList
	param := map[string]interface{}{
		"key": d.steamApiKey,
	}

	url, err := parseUrl(getLeagueListUrl(d), param)
	fmt.Println(url)
	if err != nil {
		return leagueList, err
	}
	resp, err := d.Get(url)
	if err != nil {
		return leagueList, err
	}

	err = json.Unmarshal(resp, &leagueList)
	if err != nil {
		return leagueList, err
	}
	return leagueList, nil
}
