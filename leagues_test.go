package dota2api

import (
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	leaguesGamesResponse = "{\n\"result\":{\n\"games\":[{\n\"players\":[\n{\n\"account_id\":42,\n\"name\":\"name\",\n\"hero_id\":120,\n\"team\":1\n}\n]\n,\n\"radiant_team\":{\n\"team_name\":\"TeamRad\",\n\"team_id\":42,\n\"team_logo\":42,\n\"complete\":true\n},\n\"lobby_id\":42,\n\"match_id\":42,\n\"spectators\":42,\n\"league_id\":42,\n\"league_node_id\":42,\n\"stream_delay_s\":300,\n\"radiant_series_wins\":1,\n\"series_type\":1,\n\"scoreboard\":{\n\"duration\":173.82421875,\n\"roshan_respawn_timer\":42,\n\"radiant\":{\n\"score\":1,\n\"tower_state\":2047,\n\"barracks_state\":63,\n\"picks\":[\n{\n\"hero_id\":19\n}\n]\n,\n\"bans\":[\n{\n\"hero_id\":110\n}\n]\n,\n\"players\":[\n{\n\"player_slot\":1,\n\"account_id\":42,\n\"hero_id\":120,\n\"kills\":42,\n\"death\":42,\n\"assists\":42,\n\"last_hits\":9,\n\"denies\":2,\n\"gold\":235,\n\"level\":3,\n\"gold_per_min\":243,\n\"xp_per_min\":367,\n\"ultimate_state\":1,\n\"ultimate_cooldown\":42,\n\"item0\":42,\n\"respawn_timer\":42,\n\"position_x\":-1272.0074462890625,\n\"position_y\":-1174.009521484375,\n\"net_worth\":1235\n}\n]\n,\n\"abilities\":[\n{\n\"ability_id\":5106,\n\"ability_level\":1\n}]\n}}}], \"status\": 200}}"
)

func TestDota2_GetLiveLeagueGames(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("api.GetHeroes", func() {
		var lgs LiveGames
		var err error
		g.It("Should call the correct URL", func() {
			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				switch req.URL.String() {
				case api.getLiveGamesUrl() + "?key=keyTEST":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(leaguesGamesResponse))}, nil
				case api.getHeroesUrl() + "?key=keyTEST":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
				case api.getItemsUrl() + "?key=keyTEST":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(itemsResponse))}, nil
				default:
					g.Fail("Unnecessary API call")
				}
				return nil, nil
			}
			lgs, err = api.GetLiveLeagueGames()
		})
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return the correct value", func() {
			g.Assert(lgs).Equal(LiveGames{LiveGame{
				Players: LiveGamePlayers{LiveGamePlayer{
					AccountId: 42,
					Name:      "name",
					Hero: Hero{
						ID: 120,
						Name: heroName{
							name: "pangolier",
							full: "npc_dota_hero_pangolier",
						},
					},
					Team: 1,
				}},
				LobbyId:    42,
				MatchId:    42,
				Spectators: 42,
				League: League{
					LeagueId:     42,
					LeagueNodeId: 42,
				},
				StreamDelay: 300 * time.Second,
				Series: Series{
					Radiant: 1,
					Dire:    0,
				},
				Teams: LiveTeams{Radiant: LiveTeam{
					TeamName: "TeamRad",
					TeamId:   42,
					TeamLogo: 42,
					Complete: true,
				}},
				ScoreBoard: ScoreBoard{
					Duration:           time.Duration(173.82421875 * float64(time.Second)),
					RoshanRespawnTimer: 42 * time.Second,
					Sides: Sides{
						Radiant: SideLive{
							Score: 1,
							BuildingsState: TeamBuildingsState{
								Top: LaneBuildingsState{
									T1Tower:       true,
									T2Tower:       true,
									T3Tower:       true,
									MeleeBarrack:  true,
									RangedBarrack: true,
								},
								Mid: LaneBuildingsState{
									T1Tower:       true,
									T2Tower:       true,
									T3Tower:       true,
									MeleeBarrack:  true,
									RangedBarrack: true,
								},
								Bot: LaneBuildingsState{
									T1Tower:       true,
									T2Tower:       true,
									T3Tower:       true,
									MeleeBarrack:  true,
									RangedBarrack: true,
								},
								T4TowerBot: true,
								T4TowerTop: true,
							},
							Picks: []Hero{{
								ID: 19,
								Name: heroName{
									name: "tiny",
									full: "npc_dota_hero_tiny",
								},
							}},
							Bans: []Hero{{
								ID: 110,
								Name: heroName{
									name: "phoenix",
									full: "npc_dota_hero_phoenix",
								},
							}},
							Players: PlayersLive{PlayerLive{
								PlayerSlot: 1,
								AccountID:  42,
								Hero: Hero{
									ID: 120,
									Name: heroName{
										name: "pangolier",
										full: "npc_dota_hero_pangolier",
									},
								},
								KDA: KDA{
									Kills:   42,
									Deaths:  42,
									Assists: 42,
								},
								Stats: PlayerStatsLive{
									LastHits:      9,
									Denies:        2,
									GoldPerMinute: 243,
									XpPerMinute:   367,
									Level:         3,
								},
								Items: LivePlayerItems{
									Item0: Item{
										ID: 42,
										Name: itemName{
											name: "ward_observer",
											full: "item_ward_observer",
										},
										Cost:       0,
										SecretShop: false,
										SideShop:   false,
										Recipe:     false,
									},
								},
								RespawnTimer: 42 * time.Second,
								UltimateState: UltimateState{
									UltimateState:    1,
									UltimateCooldown: 42 * time.Second,
								},
								Position: Position{
									X: -1272.0074462890625,
									Y: -1174.009521484375,
								},
								Gold: PlayerGold{
									current: 235,
									spent:   1000,
								},
							}},
							Abilities: LiveAbilities{LiveAbility{
								AbilityID:    5106,
								AbilityLevel: 1,
							}},
						},
						Dire: SideLive{
							Picks:   []Hero{},
							Bans:    []Hero{},
							Players: PlayersLive{},
						},
					},
				},
			},
			})
		})
	})
}
