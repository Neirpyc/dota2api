package dota2api

import (
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	vanityUrlResponse = "{\"response\":{\"steamid\":\"42\",\"success\":1}}"
)

func TestDota2_ResolveVanityUrl(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("api.GetHeroes", func() {
		var id SteamId
		var err error
		g.It("Should call the correct URL", func() {
			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				g.Assert(req.URL.String()).Equal(getResolveVanityUrl(&api) + "?key=keyTEST&vanityurl=vanityurl")
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(vanityUrlResponse))}, nil
			}
			id, err = api.ResolveVanityUrl(VanityUrl("vanityurl"))
		})
		g.It("Should return no error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return the correct SteamId", func() {
			g.Assert(id).Equal(NewSteamIdFrom64(42))
		})
	})
}
