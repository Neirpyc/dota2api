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
	getFriendListResponse = "{\"friendslist\":{\"friends\":[{\"steamid\":\"42\",\"relationship\":\"friend\",\"friend_since\":1551817569},{\"steamid\":\"43\",\"relationship\":\"friend\",\"friend_since\":1548161655}]}}"
)

func TestDota2_GetFriendList(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("api.GetFriendList", func() {
		var friends Friends
		var err error
		g.It("Should call the correct URL", func() {
			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				g.Assert(req.URL.String()).Equal(api.getFriendListUrl() + "?key=keyTEST&steamid=42")
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(getFriendListResponse))}, nil
			}
			friends, err = api.GetFriendList(ParameterSteamId(NewSteamIdFrom64(42)))
		})
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return the correct values", func() {
			g.Assert(friends).Equal(Friends{Friend{
				SteamId:      NewSteamIdFrom64(42),
				RelationShip: "friend",
				FriendsSince: time.Unix(1551817569, 0),
			}, Friend{
				SteamId:      NewSteamIdFrom64(43),
				RelationShip: "friend",
				FriendsSince: time.Unix(1548161655, 0),
			}})
		})
	})
}
