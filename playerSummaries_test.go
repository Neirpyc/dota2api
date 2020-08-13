package dota2api

import (
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	response0 = `{"response":{"players":[{"steamid":"42","communityvisibilitystate":3,"profilestate":1,"personaname":"userNAME","profileurl":"profileURL","avatar":"avatar32URL","avatarmedium":"avatar64URL","avatarfull":"avatar184URL","avatarhash":"avatarHASH","lastlogoff":42,"personastate":3,"personastateflags":64}]}}`
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
				g.Assert(sum[0].ProfileState).Equal(ProfileStateConfigured)
				g.Assert(sum[0].UserStatus).Equal(UserStatusAway)
				g.Assert(sum[0].PersonaStateFlag).Equal(64)
			})
		})
	})
}
