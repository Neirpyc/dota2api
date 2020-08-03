package dota2api

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestDota2_GetMatchHistory(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	_, err := api.GetMatchHistory()
	g.Describe("api.GetMatchHistory", func() {
		g.It("Should return no error", func() {
			g.Assert(err == nil).IsTrue()
		})
	})
}
