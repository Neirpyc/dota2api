package dota2api

import (
	"fmt"
	. "github.com/franela/goblin"
	"testing"
	"time"
)

func TestDota2_GetMatchDetails1(t *testing.T) {
	g := Goblin(t)
	matchId := int64(5548608983)
	api, _ := LoadConfig("config.yaml")
	g.Describe("GetMatchDetails parameters", func() {
		g.It("Should error with no parameter", func() {
			_, err := api.GetMatchDetails()
			g.Assert(err == nil).IsFalse()
		})
		g.It("Should error with wrong parameter", func() {
			_, err := api.GetMatchDetails(MatchesRequested(42))
			g.Assert(err == nil).IsFalse()
		})
		g.It("Should error with wrong and right parameter", func() {
			_, err := api.GetMatchDetails(MatchId(matchId), MatchesRequested(42))
			g.Assert(err == nil).IsFalse()
		})
		g.It("Should not error with correct parameter", func() {
			_, err := api.GetMatchDetails(MatchId(matchId))
			g.Assert(err == nil).IsTrue()
		})
	})
}

func TestDota2_GetMatchDetails2(t *testing.T) {
	g := Goblin(t)
	matchId := int64(5548608983)
	api, _ := LoadConfig("config.yaml")
	details, err := api.GetMatchDetails(MatchId(matchId))
	g.Describe(fmt.Sprintf("GetMatchDetails %d", matchId), func() {
		g.It("Should not error", func() {
			g.Assert(err == nil).IsTrue()
		})
		g.It("Should return the correct match ID and SeqNum", func() {
			g.Assert(details.MatchID).Equal(matchId)
			g.Assert(details.MatchSeqNum == 4655597189).IsTrue()
		})
		g.It("Should return a HumanPlayers, 5 players in each team", func() {
			g.Assert(details.HumanPlayers).Equal(10)
			g.Assert(details.Radiant.Count()).Equal(5)
			g.Assert(details.Dire.Count()).Equal(5)
			for _, p := range details.Dire {
				g.Assert(p.LeaverStatus == LeaverStatusNone).IsTrue()
			}
			for _, p := range details.Radiant {
				g.Assert(p.LeaverStatus == LeaverStatusNone).IsTrue()
			}
		})
		g.It("Should return Source2 as engine", func() {
			g.Assert(details.Engine == EngineSource2).IsTrue()
		})
		g.Xit("Should return Ranked Matchmaking as GameMode", func() {
			g.Assert(details.GameMode).Equal(GameModeRankedMatchmaking)
			g.Assert(details.GameMode.GetString()).Equal("Ranked Matchmaking")
		})
		g.Xit("Should return ranked Matchmaking as LobbyType", func() {
			g.Assert(details.LobbyType).Equal(LobbyRankedMatchmaking)
			g.Assert(details.LobbyType.GetName()).Equal("Ranked Matchmaking")
		})
		g.It("Should return 19-36 as score", func() {
			g.Assert(details.Score.RadiantScore).Equal(19)
			g.Assert(details.Score.DireScore).Equal(36)
		})
		g.It("Should return Dire as winner", func() {
			g.Assert(details.Victory.DireWon()).IsTrue()
			g.Assert(details.Victory.RadiantWon()).IsFalse()
			g.Assert(details.Victory.GetWinningTeam()).Equal(Dire)
		})
		g.It("Should return correct time stamps", func() {
			g.Assert(details.Duration).Equal(2552 * time.Second)
			g.Assert(details.FirstBloodTime).Equal(18 * time.Second)
			g.Assert(details.StartTime.Equal(time.Unix(1596363511, 0))).IsTrue()
			g.Assert(details.PreGameDuration).Equal(90 * time.Second)
		})
		g.It("Should return correct BuildingState", func() {
			//Radiant
			g.Assert(details.BuildingsState.Radiant.Top.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Top.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Mid.RangedBarrack).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.MeleeBarrack).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T2Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T3Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Bot.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Bot.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.T2Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.T4TowerBot).IsFalse()
			g.Assert(details.BuildingsState.Radiant.T4TowerTop).IsFalse()
			//Dire
			g.Assert(details.BuildingsState.Dire.Top.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Top.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Mid.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Bot.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.T4TowerBot).IsTrue()
			g.Assert(details.BuildingsState.Dire.T4TowerTop).IsTrue()
		})
		g.It("Should return correct time stamps", func() {
			g.Assert(details.Duration).Equal(2552 * time.Second)
			g.Assert(details.FirstBloodTime).Equal(18 * time.Second)
			g.Assert(details.StartTime.Equal(time.Unix(1596363511, 0))).IsTrue()
			g.Assert(details.PreGameDuration).Equal(90 * time.Second)
		})
		g.It("Should return correct Picks and Bans", func() {
			bansHeroesIds := []int{46, 105, 104, 41, 23, 82, 32}
			bansTeams := []int{Radiant, Radiant, Radiant, Radiant, Radiant, Dire, Dire}
			p, f := details.PicksBans.GetPick(0)
			g.Assert(f).IsTrue()
			g.Assert(p.IsPick()).IsTrue()
			p, f = details.PicksBans.GetPick(7)
			g.Assert(f).IsTrue()
			g.Assert(p.IsPick()).IsTrue()
			g.Assert(p.IsRadiant()).IsTrue()
			_, f = details.PicksBans.GetPick(17)
			g.Assert(f).IsFalse()
			for i, ban := range details.PicksBans.GetByPickType(Ban) {
				g.Assert(ban.IsBan()).IsTrue()
				g.Assert(ban.Hero.ID).Equal(bansHeroesIds[i])
				g.Assert(ban.Order).Equal(10 + i)
				g.Assert(ban.GetTeam()).Equal(bansTeams[i])
			}
			picksHeroesIds := []int{35, 26, 50, 99, 75, 7, 47, 42, 11, 44}
			picksIsDire := []bool{true, false, true, false, false, true, true, false, false, true}
			for i, pick := range details.PicksBans.GetByPickType(Pick) {
				g.Assert(pick.GetType()).Equal(Pick)
				g.Assert(pick.Hero.ID).Equal(picksHeroesIds[i])
				g.Assert(pick.Order).Equal(i)
				g.Assert(pick.IsDire()).Equal(picksIsDire[i])
			}
			for _, pick := range details.PicksBans {
				p, f := details.PicksBans.GetPickByHero(pick.Hero)
				g.Assert(f).IsTrue()
				g.Assert(p.Hero).Equal(pick.Hero)
			}
			for _, pick := range details.PicksBans.GetByTeam(Dire) {
				g.Assert(pick.IsDire()).IsTrue()
			}
		})
		g.It("Should return the correct stats for player 0", func() {
			g.Assert(details.Radiant[0].Stats.Gold.Spent() == 9490).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.Spent().Raw() == 9490).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.Spent().ToString()).Equal("9.5k")
			g.Assert(details.Radiant[0].Stats.Gold.Current() == 1250).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.NetWorth() == 10740).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroDamage.Raw() == 23281).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroDamage.Scaled() == 15776).IsTrue()
			g.Assert(details.Radiant[0].Stats.TowerDamage.Raw() == 270).IsTrue()
			g.Assert(details.Radiant[0].Stats.TowerDamage.Scaled() == 79).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroHealing.Raw() == 0).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroHealing.Scaled() == 0).IsTrue()

			g.Assert(details.Radiant[0].Stats.Gold.Current().ToString()).Equal("1.3k")
			g.Assert(details.Radiant[1].Stats.Gold.Current().ToString()).Equal("331")
		})
		g.It("Should return working stats", func() {
			for _, p := range details.Dire {
				g.Assert(p.Stats.Gold.NetWorth()).Equal(p.Stats.Gold.Current() + p.Stats.Gold.Spent())
				g.Assert(p.Stats.HeroDamage.ScalingFactor()).Equal(float64(p.Stats.HeroDamage.Scaled()) / float64(p.Stats.HeroDamage.Raw()))
				g.Assert(p.Stats.TowerDamage.ScalingFactor()).Equal(float64(p.Stats.TowerDamage.Scaled()) / float64(p.Stats.TowerDamage.Raw()))
				g.Assert(p.Stats.HeroHealing.ScalingFactor()).Equal(float64(p.Stats.HeroHealing.Scaled()) / float64(p.Stats.HeroHealing.Raw()))
			}
			for _, p := range details.Radiant {
				g.Assert(p.Stats.Gold.NetWorth()).Equal(p.Stats.Gold.Current() + p.Stats.Gold.Spent())
				g.Assert(p.Stats.HeroDamage.ScalingFactor()).Equal(float64(p.Stats.HeroDamage.Scaled()) / float64(p.Stats.HeroDamage.Raw()))
				g.Assert(p.Stats.TowerDamage.ScalingFactor()).Equal(float64(p.Stats.TowerDamage.Scaled()) / float64(p.Stats.TowerDamage.Raw()))
				g.Assert(p.Stats.HeroHealing.ScalingFactor()).Equal(float64(p.Stats.HeroHealing.Scaled()) / float64(p.Stats.HeroHealing.Raw()))
			}
		})
		g.It("Should return the correct AbilityBuild for player 0", func() {
			g.Assert(details.Radiant[0].AbilityUpgrades.Count()).Equal(18)
			aU, f := details.Radiant[0].AbilityUpgrades.GetByOrder(4)
			g.Assert(f).IsTrue()
			g.Assert(aU.Ability).Equal(5378)
			g.Assert(aU.Level).Equal(5)
			g.Assert(aU.Time).Equal(714 * time.Second)
			_, f = details.Radiant[0].AbilityUpgrades.GetByOrder(18)
			g.Assert(f).IsFalse()
			aU, f = details.Radiant[0].AbilityUpgrades.GetByLevel(10)
			g.Assert(f).IsTrue()
			g.Assert(aU.Ability).Equal(5906)
			g.Assert(aU.Level).Equal(10)
			g.Assert(aU.Time).Equal(1343 * time.Second)
			_, f = details.Radiant[0].AbilityUpgrades.GetByLevel(19)
			g.Assert(f).IsFalse()
			aUs := details.Radiant[0].AbilityUpgrades.GetByAbility(5378)
			g.Assert(aUs.Count()).Equal(4)
			for _, aU := range aUs {
				g.Assert(aU.Ability).Equal(5378)
			}
		})
		g.It("Should return working Abilities", func() {
			for _, p := range details.Dire {
				for i, a := range p.AbilityUpgrades {
					for _, aU := range p.AbilityUpgrades.GetByAbility(a.Ability) {
						if aU.Ability == a.Ability {
							goto success
						}
					}
					g.Fail(fmt.Sprintf("missing required ability %d", a.Ability))
				success:
					aU, f := p.AbilityUpgrades.GetByOrder(i)
					g.Assert(f).IsTrue()
					g.Assert(aU).Equal(a)
					aU, f = p.AbilityUpgrades.GetByLevel(a.Level)
					g.Assert(f).IsTrue()
					g.Assert(aU).Equal(a)
				}
			}
		})
	})
}
