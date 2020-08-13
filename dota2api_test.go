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

func TestLoadConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("LoadConfig", func() {
		g.It("Should load without error", func() {
			_, err := LoadConfig("config.yaml")
			g.Assert(err).Equal(nil)
		})
	})
}
