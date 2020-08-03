package dota2api

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestDota2_GetMatchHistory(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	g.Describe("api.GetMatchHistory base", func() {
		hist, err := api.GetMatchHistory()
		g.It("Should return no error", func() {
			g.Assert(err == nil).IsTrue()
		})
		g.It("Should return at least one result", func() {
			g.Assert(hist.Count() > 0).IsTrue()
		})
		g.It("Should return a match seq num for each result", func() {
			for _, match := range hist.Matches {
				g.Assert(match.MatchSeqNum != 0).IsTrue()
			}
		})
		g.It("Should return a match ID for each result", func() {
			for _, match := range hist.Matches {
				g.Assert(match.MatchId != 0).IsTrue()
			}
		})
		g.It("Should return a start time for each result", func() {
			for _, match := range hist.Matches {
				g.Assert(match.StartTime.Unix() != 0).IsTrue()
			}
		})
		g.It("Should return a working LobbyType for each result", func() {
			for _, match := range hist.Matches {
				match.LobbyType.GetId()
				g.Assert(match.LobbyType.GetName() != "").IsTrue()
			}
		})
		g.It("Should return a team for each result", func() {
			for _, match := range hist.Matches {
				g.Assert(match.Radiant.players != nil || match.Dire.players != nil).IsTrue()
			}
		})
	})
	g.Describe("api.GetMatchHistory parameters", func() {
		g.It("Should refuse invalid parameters", func() {
			_, err := api.GetMatchHistory(42, 37)
			g.Assert(err != nil).IsTrue()
		})
		g.It("Should refuse even one invalid parameter", func() {
			_, err := api.GetMatchHistory(MatchesRequested(52), 37)
			g.Assert(err != nil).IsTrue()
		})
		matches, err := api.GetMatchHistory(MatchesRequested(42))
		g.It("Should accept a valid parameter", func() {
			g.Assert(err == nil).IsTrue()
		})
		g.It("Should work with matchesRequested parameter", func() {
			g.Assert(matches.Count() == 42).IsTrue()
		})
		g.It("Should work with accountId and heroId parameter", func() {
			id := int64(76561198067618887)
			matches, err := api.GetMatchHistory(AccountId(id), HeroId(42), MatchesRequested(100))
			g.Assert(err == nil).IsTrue()
			if err != nil {
				return
			}
			g.Assert(matches.Count() <= 100).IsTrue()
		topLoop:
			for _, match := range matches.Matches {
				flag := 0
				for i := 0; ; i++ {
					if p, found := match.GetPlayer(i); found {
						if p.AccountId == int(int32(id)) {
							flag++
						}
						if p.Hero.ID == 42 {
							flag++
						}
						if flag == 2 {
							continue topLoop
						}
					} else {
						break
					}
				}
				g.Fail("Could not find requested accountId and/or heroId")
			}
		})
		g.It("Should work with startAtMatchId parameter", func() {
			id := matches.Matches[matches.Count()/2].MatchId
			matches, err := api.GetMatchHistory(StartAtMatchId(id))
			g.Assert(err == nil).IsTrue()
			if err != nil {
				return
			}
			for _, match := range matches.Matches {
				if match.MatchId > id {
					g.Fail("Match with previous ID returned")
				}
			}
		})
	})
	g.Describe("api.GetMatchHistory cursor", func() {
		c := NewCursor()
		matches, err := api.GetMatchHistory(c, MatchesRequested(50))
		g.It("Should accept a cursor parameter", func() {
			g.Assert(err == nil).IsTrue()
		})
		g.It("Should modify the cursor parameter", func() {
			g.Assert(c.c != nil).IsTrue()
		})
		g.It("Should modify correctly the cursor", func() {
			g.Assert(matches.Matches[matches.Count()-1].MatchId == c.GetLastReceivedMatch()).IsTrue()
		})
		old := c.GetLastReceivedMatch()
		matches, _ = api.GetMatchHistory(c, MatchesRequested(50))
		g.It("Should modify correctly the cursor when reusing a cursor", func() {
			g.Assert(matches.Matches[matches.Count()-1].MatchId == c.GetLastReceivedMatch()).IsTrue()
		})
		g.It("Should send a new batch of matches when reusing a cursor", func() {
			for _, match := range matches.Matches {
				g.Assert(match.MatchId < old).IsTrue()
			}
		})
		g.It("Should error when no match is remaining", func() {
			for c.GetRemaining() > 0 {
				_, _ = api.GetMatchHistory(c, MatchesRequested(200))
			}
			matches, err = api.GetMatchHistory(c, MatchesRequested(200))
			g.Assert(err == nil).IsTrue()
			g.Assert(matches.Count() == 0).IsTrue()
		})
	})
}

func TestDota2_GetMatchHistory_Teams(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	matches, _ := api.GetMatchHistory(MatchesRequested(100), AccountId(76561198067618887))
	check := func(player Player, p Player, f bool, checkAccountId bool) {
		g.Assert(f).IsTrue()
		if checkAccountId {
			g.Assert(p.AccountId == player.AccountId).IsTrue()
		}
		g.Assert(p.Hero.ID == player.Hero.ID).IsTrue()
	}
	g.Describe("Teams", func() {
		g.It("Should return found=false on non found pos", func() {
			for _, match := range matches.Matches {
				_, f := match.GetPlayer(match.Radiant.Count() + match.Dire.Count() + 1)
				g.Assert(f).IsFalse()
			}
		})
		heroes, _ := api.GetHeroes()
		nonPlayedHero := func(match MatchSummary) Hero {
		heroLoop:
			for _, hero := range heroes.heroes {
				for _, player := range match.Radiant.players {
					if player.Hero.ID == hero.ID {
						continue heroLoop
					}
				}
				for _, player := range match.Dire.players {
					if player.Hero.ID == hero.ID {
						continue heroLoop
					}
				}
				return hero
			}
			return Hero{}
		}
		g.It("Should return found=false on non hero", func() {
			for _, match := range matches.Matches {
				_, f := match.GetByHero(nonPlayedHero(match))
				g.Assert(f).IsFalse()
			}
		})
		g.It("Should return found=false on non heroId", func() {
			for _, match := range matches.Matches {
				_, f := match.GetByHeroId(nonPlayedHero(match).ID)
				g.Assert(f).IsFalse()
			}
		})
		g.It("Should find player by pos", func() {
			for _, match := range matches.Matches {
				c := 0
				loopF := func(player Player) {
					p, f := match.GetPlayer(c)
					check(player, p, f, true)
					c++
				}
				for _, player := range match.Radiant.players {
					loopF(player)
				}
				for _, player := range match.Dire.players {
					loopF(player)
				}
			}
		})
		g.It("Should find player by Hero", func() {
			for _, match := range matches.Matches {
				loopF := func(player Player) {
					p, f := match.GetByHero(player.Hero)
					check(player, p, f, false)
				}
				for _, player := range match.Radiant.players {
					loopF(player)
				}
				for _, player := range match.Dire.players {
					loopF(player)
				}
			}
		})
		g.It("Should find player by HeroId", func() {
			for _, match := range matches.Matches {
				loopF := func(player Player) {
					p, f := match.GetByHeroId(player.Hero.ID)
					check(player, p, f, false)
				}
				for _, player := range match.Radiant.players {
					loopF(player)
				}
				for _, player := range match.Dire.players {
					loopF(player)
				}
			}
		})
		g.It("Should find player by Team.getByHero", func() {
			for _, match := range matches.Matches {
				for _, player := range match.Radiant.players {
					p, f := match.Radiant.GetByHero(player.Hero)
					check(player, p, f, false)
				}
				for _, player := range match.Dire.players {
					p, f := match.Dire.GetByHero(player.Hero)
					check(player, p, f, false)
				}
			}
		})
	})
}
