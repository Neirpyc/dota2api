package dota2api

import (
	. "github.com/franela/goblin"
	"net/http"
	"testing"
)

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func GetTestConfig() Config {
	return Config{
		Timeout:     1,
		SteamApiKey: "keyTEST",
	}
}

func TestLoadConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("LoadConfigFromFile", func() {
		g.It("Should load without error", func() {
			_, err := LoadConfigFromFile("config.yaml")
			g.Assert(err).Equal(nil)
		})
	})
}
