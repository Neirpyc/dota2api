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
				g.Assert(req.URL.String()).Equal(getFriendListUrl(&api) + "?key=keyTEST&steamid=42")
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(getFriendListResponse))}, nil
			}
			friends, err = api.GetFriendList(ParameterSteamId(NewSteamIdFrom64(42)))
		})
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return the correct values", func() {
			g.Assert(friends.Count()).Equal(2)
			g.Assert(friends[0].SteamId).Equal(NewSteamIdFrom64(42))
			g.Assert(friends[0].FriendsSince).Equal(time.Unix(1551817569, 0))
			g.Assert(friends[0].RelationShip).Equal("friend")
			g.Assert(friends[1].SteamId).Equal(NewSteamIdFrom64(43))
			g.Assert(friends[1].FriendsSince).Equal(time.Unix(1548161655, 0))
			g.Assert(friends[1].RelationShip).Equal("friend")
		})
	})
}
