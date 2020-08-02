package dota2api

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("LoadConfig", func() {
		g.It("Should load without error", func() {
			_, err := LoadConfig("config.ini")
			g.Assert(err).Equal(nil)
		})
	})
}
