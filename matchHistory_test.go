package dota2api

import (
	"errors"
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	matchHistoryResponse0 = "{\n\"result\":{\n\"status\":1,\n\"num_results\":2,\n\"total_results\":500,\n\"results_remaining\":498,\n\"matches\":[\n{\n\"match_id\":5569057130,\n\"match_seq_num\":4674249362,\n\"start_time\":1597496593,\n\"lobby_type\":12,\n\"radiant_team_id\":0,\n\"dire_team_id\":0,\n\"players\":[\n{\n\"account_id\":241341785,\n\"player_slot\":0,\n\"hero_id\":128\n},\n{\n\"account_id\":198212358,\n\"player_slot\":1,\n\"hero_id\":47\n},\n{\n\"account_id\":196978137,\n\"player_slot\":2,\n\"hero_id\":63\n},\n{\n\"account_id\":209793761,\n\"player_slot\":3,\n\"hero_id\":35\n}\n]\n\n},\n{\n\"match_id\":5569057115,\n\"match_seq_num\":4674248663,\n\"start_time\":1597496594,\n\"lobby_type\":12,\n\"radiant_team_id\":0,\n\"dire_team_id\":0,\n\"players\":[\n{\n\"account_id\":142181463,\n\"player_slot\":0,\n\"hero_id\":70\n},\n{\n\"account_id\":153072049,\n\"player_slot\":1,\n\"hero_id\":35\n},\n{\n\"account_id\":183925783,\n\"player_slot\":2,\n\"hero_id\":129\n},\n{\n\"account_id\":103439564,\n\"player_slot\":3,\n\"hero_id\":63\n}\n]\n\n}]}}"
	matchHistoryResponse1 = "{\n\"result\":{\n\"status\":1,\n\"num_results\":2,\n\"total_results\":500,\n\"results_remaining\":496,\n\"matches\":[{\"match_id\":43, \"match_seq_num\":45},{\"match_id\":42, \"match_seq_num\":44}]}}"
)

func TestDota2_GetMatchHistory(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	var matches MatchHistory
	var err error
	g.Describe("api.GetMatchHistory", func() {
		g.Describe("Basic test", func() {
			g.It("Should call the correct URL", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					switch req.URL.String() {
					case api.getMatchHistoryUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse0))}, nil
					case api.getHeroesUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
					default:
						g.Fail("Unnecessary API call " + req.URL.String())
					}
					return nil, nil
				}
				matches, err = api.GetMatchHistory()
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 2 matches", func() {
				g.Assert(matches.Count()).Equal(2)
			})
			g.It("Should return a valid content", func() {
				if matches.Count() != 2 {
					g.Fail("Wrong match count")
				}
				g.Assert(matches[0].MatchId).Equal(int64(5569057130))
				g.Assert(matches[0].MatchSeqNum).Equal(int64(4674249362))
				g.Assert(matches[0].StartTime.Equal(time.Unix(1597496593, 0))).IsTrue()
				g.Assert(matches[0].LobbyType).Equal(LobbyType(12))
				g.Assert(matches[0].Radiant.Id).Equal(0)
				g.Assert(matches[0].Dire.Id).Equal(0)
				g.Assert(matches[0].Radiant.players[2].AccountId).Equal(196978137)
				g.Assert(matches[1].MatchId).Equal(int64(5569057115))
			})
		})
		g.Describe("Test with cursor", func() {
			g.Before(func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					switch req.URL.String() {
					case api.getMatchHistoryUrl() + "?key=keyTEST&start_at_match_id=-1":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse0))}, nil
					case api.getMatchHistoryUrl() + "?key=keyTEST&start_at_match_id=5569057114":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse1))}, nil
					case api.getHeroesUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
					default:
						g.Fail("Unnecessary API call " + req.URL.String())
					}
					return nil, errors.New("unnecessary API call " + req.URL.String())
				}
			})
			c := NewCursor()
			g.It("Should accept an empty cursor", func() {
				_, err := api.GetMatchHistory(c)
				g.Assert(err).IsNil()
			})
			g.It("Should return an initialized cursor", func() {
				g.Assert(c.GetLastReceivedMatch()).Equal(int64(5569057115))
			})
			g.It("Should accept an existing cursor", func() {
				_, err = api.GetMatchHistory(c)
				g.Assert(err).IsNil()
			})
			g.It("Should update an initialized cursor", func() {
				g.Assert(c.GetLastReceivedMatch()).Equal(int64(42))
			})
		})
	})
}

func TestDota2_GetMatchHistoryBySequenceNum(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	var matches MatchHistory
	var err error
	g.Describe("api.GetMatchHistoryBySequenceNum", func() {
		g.Describe("Basic test", func() {
			g.It("Should call the correct URL", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					switch req.URL.String() {
					case api.getMatchHistoryBySequenceNumUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse0))}, nil
					case api.getHeroesUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
					default:
						g.Fail("Unnecessary API call " + req.URL.String())
					}
					return nil, nil
				}
				matches, err = api.GetMatchHistoryBySequenceNum()
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 2 matches", func() {
				g.Assert(matches.Count()).Equal(2)
			})
			g.It("Should return a valid content", func() {
				g.Assert(matches[0].MatchId).Equal(int64(5569057130))
				g.Assert(matches[0].MatchSeqNum).Equal(int64(4674249362))
				g.Assert(matches[0].StartTime.Equal(time.Unix(1597496593, 0))).IsTrue()
				g.Assert(matches[0].LobbyType).Equal(LobbyType(12))
				g.Assert(matches[0].Radiant.Id).Equal(0)
				g.Assert(matches[0].Dire.Id).Equal(0)
				g.Assert(matches[0].Radiant.players[2].AccountId).Equal(196978137)
				g.Assert(matches[1].MatchId).Equal(int64(5569057115))
			})
		})
		g.Describe("Test with cursor", func() {
			g.Before(func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					switch req.URL.String() {
					case api.getMatchHistoryBySequenceNumUrl() + "?key=keyTEST&start_at_match_seq_num=0":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse0))}, nil
					case api.getMatchHistoryBySequenceNumUrl() + "?key=keyTEST&start_at_match_seq_num=4674248664":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(matchHistoryResponse1))}, nil
					case api.getHeroesUrl() + "?key=keyTEST":
						return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
					default:
						g.Fail("Unnecessary API call " + req.URL.String())
					}
					return nil, errors.New("unnecessary API call " + req.URL.String())
				}
			})
			c := NewCursor()
			g.It("Should accept an empty cursor", func() {
				_, err := api.GetMatchHistoryBySequenceNum(c)
				g.Assert(err).IsNil()
			})
			g.It("Should return an initialized cursor", func() {
				g.Assert(c.GetLastReceivedMatch()).Equal(int64(4674248663))
			})
			g.It("Should accept an existing cursor", func() {
				_, err = api.GetMatchHistoryBySequenceNum(c)
				g.Assert(err).IsNil()
			})
			g.It("Should update an initialized cursor", func() {
				g.Assert(c.GetLastReceivedMatch()).Equal(int64(44))
			})
		})
	})
}
