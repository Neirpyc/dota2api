// +build integration

package dota2api

import (
	"fmt"
	. "github.com/franela/goblin"
	"testing"
	"time"
)

//These tests make sure that the API returns data in the format the code can parse

func TestDota2_GetPlayerSummaries_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetPlayerSummaries Integration Test", func() {
		steamId0 := NewSteamIdFrom64(76561198054320440)
		steamId1 := NewSteamIdFrom64(76561198048536965)
		sum, err := api.GetPlayerSummaries(ParameterSteamIds(steamId0, steamId1))
		g.It("Should return summaries from the correct players", func() {
			g.Assert(err).IsNil()
			if sum[0].SteamId == steamId0 {
				g.Assert(sum[0].SteamId).Equal(steamId0)
				g.Assert(sum[1].SteamId).Equal(steamId1)
			} else {
				g.Assert(sum[0].SteamId).Equal(steamId1)
				g.Assert(sum[1].SteamId).Equal(steamId0)
			}
		})
	})
}

func TestDota2_GetFriendList_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetFriendList Integration Test", func() {
		friendList, err := api.GetFriendList(ParameterSteamId(NewSteamIdFrom64(76561198392783738)))
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return non empty friends", func() {
			g.Assert(friendList.Count() > 0).IsTrue()
			friendList.ForEach(func(friend Friend) {
				id, err := friend.SteamId.SteamId64()
				g.Assert(err).IsNil()
				g.Assert(id != 0).IsTrue()
				g.Assert(friend.RelationShip != "").IsTrue()
				g.Assert(friend.FriendsSince.Equal(time.Time{})).IsFalse()
			})
		})
	})
}

func TestDota2_GetHeroes_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetHeroes Integration Test", func() {
		heroes, err := api.GetHeroes()
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return non empty heroes", func() {
			g.Assert(heroes.Count() > 0).IsTrue()
			heroes.ForEach(func(hero Hero) {
				g.Assert(hero.Name.GetName() != "").IsTrue()
				g.Assert(hero.Name.GetFullName() != "").IsTrue()
				g.Assert(hero.Id != 0).IsTrue()
			})
		})
	})
}

func TestDota2_GetItems_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetItems Integration Test", func() {
		items, err := api.GetItems()
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return non empty items", func() {
			g.Assert(items.Count() > 0).IsTrue()
			items.ForEach(func(item Item) {
				g.Assert(item.Name.GetName() != "").IsTrue()
				g.Assert(item.Name.GetFullName() != "").IsTrue()
				g.Assert(item.Id != 0).IsTrue()
			})
		})
	})
}

func TestDota2_GetLiveLeagueGames_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetLiveLeagueGames Integration Test", func() {
		games, err := api.GetLiveLeagueGames()
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return non empty games", func() {
			g.Assert(games.Count() > 0).IsTrue()
			games.ForEachGame(func(game LiveGame) {
				g.Assert(game.ScoreBoard).IsNotZero()
			})
		})
	})
}

func TestDota2_GetMatchDetails_Integration(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	g.Describe("api.GetHeroes Integration Test", func() {
		details, err := api.GetMatchDetails(MatchId(5548608983))
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.Describe("Should return the correct game", func() {
			g.It("Should return the correct match ID and SeqNum", func() {
				g.Assert(details.MatchID).Equal(int64(5548608983))
				g.Assert(details.MatchSeqNum == 4655597189).IsTrue()
			})
			g.It("Should return a HumanPlayers, 5 players in each team", func() {
				g.Assert(details.HumanPlayers).Equal(10)
				g.Assert(details.Radiant.Count()).Equal(5)
				g.Assert(details.Dire.Count()).Equal(5)
				details.ForEachPlayer(func(p PlayerDetails) {
					g.Assert(p.LeaverStatus).Equal(LeaverStatusNone)
				})
			})
			g.It("Should return Source2 as engine", func() {
				g.Assert(details.Engine).Equal(Source2)
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
				g.Assert(details.BuildingsState).Equal(BuildingsState{
					Dire: TeamBuildingsState{
						Top: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       true,
							T3Tower:       true,
							MeleeBarrack:  true,
							RangedBarrack: true,
						},
						Mid: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       true,
							T3Tower:       true,
							MeleeBarrack:  true,
							RangedBarrack: true,
						},
						Bot: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       true,
							T3Tower:       true,
							MeleeBarrack:  true,
							RangedBarrack: true,
						},
						T4TowerBot: true,
						T4TowerTop: true,
					},
					Radiant: TeamBuildingsState{
						Top: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       true,
							T3Tower:       true,
							MeleeBarrack:  true,
							RangedBarrack: true,
						},
						Mid: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       false,
							T3Tower:       false,
							MeleeBarrack:  false,
							RangedBarrack: false,
						},
						Bot: LaneBuildingsState{
							T1Tower:       false,
							T2Tower:       false,
							T3Tower:       true,
							MeleeBarrack:  true,
							RangedBarrack: true,
						},
						T4TowerBot: false,
						T4TowerTop: false,
					},
				})
			})
			g.It("Should return correct time stamps", func() {
				g.Assert(details.Duration).Equal(2552 * time.Second)
				g.Assert(details.FirstBloodTime).Equal(18 * time.Second)
				g.Assert(details.StartTime.Equal(time.Unix(1596363511, 0))).IsTrue()
				g.Assert(details.PreGameDuration).Equal(90 * time.Second)
			})
			g.It("Should return correct Picks and Bans", func() {
				bansHeroesIds := []int{46, 105, 104, 41, 23, 82, 32}
				bansTeams := []Victory{RadiantVictory, RadiantVictory, RadiantVictory, RadiantVictory, RadiantVictory, DireVictory, DireVictory}
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
					g.Assert(ban.Hero.Id).Equal(bansHeroesIds[i])
					g.Assert(ban.Order).Equal(10 + i)
					g.Assert(ban.GetTeam() == int(bansTeams[i])).IsTrue()
				}
				picksHeroesIds := []int{35, 26, 50, 99, 75, 7, 47, 42, 11, 44}
				picksIsDire := []bool{true, false, true, false, false, true, true, false, false, true}
				for i, pick := range details.PicksBans.GetByPickType(Pick) {
					g.Assert(pick.GetType()).Equal(Pick)
					g.Assert(pick.Hero.Id).Equal(picksHeroesIds[i])
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
				details.ForEachPlayer(func(p PlayerDetails) {
					g.Assert(p.Stats.Gold.NetWorth()).Equal(p.Stats.Gold.Current() + p.Stats.Gold.Spent())
					g.Assert(p.Stats.HeroDamage.ScalingFactor()).Equal(float64(p.Stats.HeroDamage.Scaled()) / float64(p.Stats.HeroDamage.Raw()))
					g.Assert(p.Stats.TowerDamage.ScalingFactor()).Equal(float64(p.Stats.TowerDamage.Scaled()) / float64(p.Stats.TowerDamage.Raw()))
					g.Assert(p.Stats.HeroHealing.ScalingFactor()).Equal(float64(p.Stats.HeroHealing.Scaled()) / float64(p.Stats.HeroHealing.Raw()))
				})
			})
			g.It("Should return the correct items for player 0", func() {
				g.Assert(details.Radiant[0].Items.Item0.Id == 63).IsTrue()
				g.Assert(details.Radiant[0].Items.Item1.Id == 77).IsTrue()
				g.Assert(details.Radiant[0].Items.Item2.Id == 236).IsTrue()
				g.Assert(details.Radiant[0].Items.Item3.Id == 77).IsTrue()
				g.Assert(details.Radiant[0].Items.Item4.Id == 485).IsTrue()
				g.Assert(details.Radiant[0].Items.Item5.Id == 7).IsTrue()
				g.Assert(details.Radiant[0].Items.BackpackItem0.Id == 0).IsTrue()
				g.Assert(details.Radiant[0].Items.BackpackItem1.Id == 0).IsTrue()
				g.Assert(details.Radiant[0].Items.BackpackItem2.Id == 0).IsTrue()
				g.Assert(details.Radiant[0].Items.ItemNeutral.Id == 357).IsTrue()
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
				details.Dire.ForEach(func(p PlayerDetails) {
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
				})
			})
		})
	})
}
