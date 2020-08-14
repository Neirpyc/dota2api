// +build integration

package dota2api

import (
	. "github.com/franela/goblin"
	"testing"
)

//These tests make sure that the API returns data in the format the code can parse

func TestDota2_GetPlayerSummaries_Integration(t *testing.T) {
	g := Goblin(t)
	api, err := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetPlayerSummaries Integration Test", func() {
		g.It("Should return summaries from the correct players", func() {
			g.Assert(err).IsNil()
			steamId0 := NewSteamIdFrom64(76561198054320440)
			steamId1 := NewSteamIdFrom64(76561198048536965)
			sum, err := api.GetPlayerSummaries(ParameterSteamIds(steamId0, steamId1))
			g.Assert(err).IsNil()
			g.Assert(sum[0].SteamId).Equal(steamId0)
			g.Assert(sum[1].SteamId).Equal(steamId1)
		})
	})
}
