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
	response0       = `{"response":{"players":[{"steamid":"42","communityvisibilitystate":3,"profilestate":1,"personaname":"userNAME","profileurl":"profileURL","avatar":"avatar32URL","avatarmedium":"avatar64URL","avatarfull":"avatar184URL","avatarhash":"avatarHASH","lastlogoff":42,"personastate":3,"personastateflags":64}]}}`
	response1       = `{"response":{"players":[{"steamid":"42","communityvisibilitystate":3,"profilestate":1,"personaname":"userNAME","profileurl":"profileURL","avatar":"avatar32URL","avatarmedium":"avatar64URL","avatarfull":"avatar184URL","avatarhash":"avatarHASH","lastlogoff":42,"personastate":3,"personastateflags":64},{"steamid":"43","communityvisibilitystate":2,"profilestate":0,"personaname":"userNAME0","profileurl":"profileURL0","avatar":"avatar32URL0","avatarmedium":"avatar64URL0","avatarfull":"avatar184URL0","avatarhash":"avatarHASH0","lastlogoff":43,"personastate":2,"personastateflags":65}]}}`
	responseWithOpt = `{"response":{"players":[{"steamid":"42","commentpermission":"15","realname":"realNAME","primaryclanid":42,"timecreated":43,"loccountrycode":"locCOUNTRYcode","locstatecode":"locSTATEcode","loccityid":44,"gameid":"gameID","gameextrainfo":"gameEXTRAinfo","gameserverip":"gameSERVERip"}]}}`
)

func TestDota2_GetPlayerSummaries(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("api.GetPlayerSummaries", func() {
		g.Describe("Basic response", func() {
			var sum PlayerAccounts
			var err error
			g.It("Should call the correct request URI", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					g.Assert(req.URL.String()).Equal(getPlayerSummariesUrl(&api) + "?key=keyTEST&steamids=%5B42%5D")
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response0))}, nil
				}
				sum, err = api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 1 summary", func() {
				g.Assert(len(sum)).Equal(1)
			})
			g.It("Should parse non Optional flags", func() {
				g.Assert(len(sum)).IsNotZero()
				g.Assert(sum[0].SteamId).Equal(NewSteamIdFrom64(42))
				g.Assert(sum[0].CommunityVisibilityState).Equal(VisibilityFriendsOfFriends)
				g.Assert(sum[0].ProfileState).Equal(ProfileStateConfigured)
				g.Assert(sum[0].DisplayName).Equal("userNAME")
				g.Assert(sum[0].ProfileUrl).Equal("profileURL")
				g.Assert(sum[0].Avatar.Avatar32Url).Equal("avatar32URL")
				g.Assert(sum[0].Avatar.Avatar64Url).Equal("avatar64URL")
				g.Assert(sum[0].Avatar.Avatar184Url).Equal("avatar184URL")
				g.Assert(sum[0].Avatar.Hash).Equal("avatarHASH")
				g.Assert(sum[0].LastLogOff.Equal(time.Unix(42, 0)))
				g.Assert(sum[0].UserStatus).Equal(UserStatusAway)
				g.Assert(sum[0].PersonaStateFlag).Equal(64)
			})
		})
		g.Describe("Double response", func() {
			var sum PlayerAccounts
			var err error
			g.It("Should call the correct request URI", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					g.Assert(req.URL.String()).Equal(getPlayerSummariesUrl(&api) + "?key=keyTEST&steamids=%5B42%2C43%5D")
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response1))}, nil
				}
				sum, err = api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42), NewSteamIdFrom64(43)))
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 2 summaries", func() {
				g.Assert(len(sum)).Equal(2)
			})
			g.It("Should parse non Optional flags", func() {
				g.Assert(len(sum) > 1).IsTrue()
				//user
				g.Assert(sum[0].SteamId).Equal(NewSteamIdFrom64(42))
				g.Assert(sum[0].CommunityVisibilityState).Equal(VisibilityFriendsOfFriends)
				g.Assert(sum[0].ProfileState).Equal(ProfileStateConfigured)
				g.Assert(sum[0].DisplayName).Equal("userNAME")
				g.Assert(sum[0].ProfileUrl).Equal("profileURL")
				g.Assert(sum[0].Avatar.Avatar32Url).Equal("avatar32URL")
				g.Assert(sum[0].Avatar.Avatar64Url).Equal("avatar64URL")
				g.Assert(sum[0].Avatar.Avatar184Url).Equal("avatar184URL")
				g.Assert(sum[0].Avatar.Hash).Equal("avatarHASH")
				g.Assert(sum[0].LastLogOff.Equal(time.Unix(42, 0)))
				g.Assert(sum[0].UserStatus).Equal(UserStatusAway)
				g.Assert(sum[0].PersonaStateFlag).Equal(64)
				//user0
				g.Assert(sum[1].SteamId).Equal(NewSteamIdFrom64(43))
				g.Assert(sum[1].CommunityVisibilityState).Equal(VisibilityFriendsOnly)
				g.Assert(sum[1].ProfileState).Equal(ProfileStateEmpty)
				g.Assert(sum[1].DisplayName).Equal("userNAME0")
				g.Assert(sum[1].ProfileUrl).Equal("profileURL0")
				g.Assert(sum[1].Avatar.Avatar32Url).Equal("avatar32URL0")
				g.Assert(sum[1].Avatar.Avatar64Url).Equal("avatar64URL0")
				g.Assert(sum[1].Avatar.Avatar184Url).Equal("avatar184URL0")
				g.Assert(sum[1].Avatar.Hash).Equal("avatarHASH0")
				g.Assert(sum[1].LastLogOff.Equal(time.Unix(43, 0)))
				g.Assert(sum[1].UserStatus).Equal(UserStatusBusy)
				g.Assert(sum[1].PersonaStateFlag).Equal(65)
			})
		})
		g.Describe("Optional fields", func() {
			g.It("Should report missing fields", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response0))}, nil
				}
				sum, _ := api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
				_, f := sum[0].Optional.GameName()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.GameId()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.RealName()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.PrimaryClanId()
				g.Assert(f).IsFalse()
				f = sum[0].Optional.CommentPermission()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.TimeCreated()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.LocCityId()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.LocStateCode()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.LocCountryCode()
				g.Assert(f).IsFalse()
				_, f = sum[0].Optional.GameServerIp()
				g.Assert(f).IsFalse()
			})
			g.It("Should return correct values when found", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(responseWithOpt))}, nil
				}
				sum, _ := api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
				str, f := sum[0].Optional.GameName()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("gameEXTRAinfo")
				str, f = sum[0].Optional.GameId()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("gameID")
				str, f = sum[0].Optional.RealName()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("realNAME")
				i, f := sum[0].Optional.PrimaryClanId()
				g.Assert(f).IsTrue()
				g.Assert(i == 42).IsTrue()
				f = sum[0].Optional.CommentPermission()
				g.Assert(f).IsTrue()
				t, f := sum[0].Optional.TimeCreated()
				g.Assert(f).IsTrue()
				g.Assert(t.Equal(time.Unix(43, 0))).IsTrue()
				i, f = sum[0].Optional.LocCityId()
				g.Assert(f).IsTrue()
				g.Assert(i == 44).IsTrue()
				str, f = sum[0].Optional.LocStateCode()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("locSTATEcode")
				str, f = sum[0].Optional.LocCountryCode()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("locCOUNTRYcode")
				str, f = sum[0].Optional.GameServerIp()
				g.Assert(f).IsTrue()
				g.Assert(str).Equal("gameSERVERip")
			})
		})
	})
}
