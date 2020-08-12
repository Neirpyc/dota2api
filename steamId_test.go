package dota2api

import (
	"fmt"
	. "github.com/franela/goblin"
	"testing"
)

func TestNewSteamIdFrom32(t *testing.T) {
	g := Goblin(t)
	g.Describe("NewSteamIdFrom32", func() {
		steamId := NewSteamIdFrom32(4039455774)
		g.It("Should set the Id", func() {
			g.Assert(steamId.id == 4039455774).IsTrue()
		})
		g.It("Should set the Id type", func() {
			g.Assert(steamId.isId64).IsFalse()
		})
	})
}

func TestNewSteamIdFrom64(t *testing.T) {
	g := Goblin(t)
	g.Describe("NewSteamIdFrom64", func() {
		steamId := NewSteamIdFrom64(8674665223082153551)
		g.It("Should set the Id", func() {
			g.Assert(steamId.id == 8674665223082153551).IsTrue()
		})
		g.It("Should set the Id type", func() {
			g.Assert(steamId.isId64).IsTrue()
		})
	})
}

func TestSteamId_SetSteamId32(t *testing.T) {
	g := Goblin(t)
	g.Describe("SteamId.SetSteamId32", func() {
		g.Describe("With a previously empty SteamId", func() {
			steamId := SteamId{}
			steamId.SetSteamId32(2854263694)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 2854263694).IsTrue(fmt.Sprintf("expected %d, got %d", 2854263694, steamId.id))
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsFalse()
			})
		})
		g.Describe("With a previously SteamId32 SteamId", func() {
			steamId := NewSteamIdFrom32(4039455774)
			steamId.SetSteamId32(2854263694)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 2854263694).IsTrue(fmt.Sprintf("expected %d, got %d", 2854263694, steamId.id))
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsFalse()
			})
		})
		g.Describe("With a previously SteamId64 SteamId", func() {
			steamId := NewSteamIdFrom64(8674665223082153551)
			steamId.SetSteamId32(2854263694)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 2854263694).IsTrue(fmt.Sprintf("expected %d, got %d", 2854263694, steamId.id))
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsFalse()
			})
		})
	})
}

func TestSteamId_SetSteamId64(t *testing.T) {
	g := Goblin(t)
	g.Describe("SteamId.SetSteamId64", func() {
		g.Describe("With a previously empty SteamId", func() {
			steamId := SteamId{}
			steamId.SetSteamId64(15352856648520921629)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 15352856648520921629).IsTrue()
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsTrue()
			})
		})
		g.Describe("With a previously SteamId32 SteamId", func() {
			steamId := NewSteamIdFrom32(4039455774)
			steamId.SetSteamId64(15352856648520921629)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 15352856648520921629).IsTrue()
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsTrue()
			})
		})
		g.Describe("With a previously SteamId64 SteamId", func() {
			steamId := NewSteamIdFrom64(8674665223082153551)
			steamId.SetSteamId64(15352856648520921629)
			g.It("Should set the Id", func() {
				g.Assert(steamId.id == 15352856648520921629).IsTrue()
			})
			g.It("Should set the Id type", func() {
				g.Assert(steamId.isId64).IsTrue()
			})
		})
	})
}

func TestSteamId_SteamId32(t *testing.T) {
	g := Goblin(t)
	g.Describe("SteamId.SteamId32", func() {
		g.Describe("With a previously SteamId32 SteamId", func() {
			steamId := NewSteamIdFrom32(4039455774)
			id := steamId.SteamId32()
			g.It("Should set the correct Id", func() {
				g.Assert(id == 4039455774).IsTrue()
			})
		})
		g.Describe("With a previously SteamId64 SteamId", func() {
			steamId := NewSteamIdFrom64(8674665223082153551)
			id := steamId.SteamId32()
			g.It("Should set the correct Id", func() {
				g.Assert(id == 1597969999).IsTrue()
			})
		})
	})
}

func TestSteamId_SteamId64(t *testing.T) {
	g := Goblin(t)
	g.Describe("SteamId.SteamId64", func() {
		g.Describe("With a previously SteamId32 SteamId", func() {
			steamId := NewSteamIdFrom32(4039455774)
			id, found := steamId.SteamId64()
			g.It("Should return found == false", func() {
				g.Assert(found).IsFalse()
			})
			g.It("Should return no Id", func() {
				g.Assert(id == 0).IsTrue()
			})
		})
		g.Describe("With a previously SteamId64 SteamId", func() {
			steamId := NewSteamIdFrom64(8674665223082153551)
			id, found := steamId.SteamId64()
			g.It("Should return found == true", func() {
				g.Assert(found).IsTrue()
			})
			g.It("Should return the correct Id", func() {
				g.Assert(id == 8674665223082153551).IsTrue()
			})
		})
	})
}
