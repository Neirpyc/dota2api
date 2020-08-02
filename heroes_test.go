package dota2api

import (
	. "github.com/franela/goblin"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGetHeroes(t *testing.T) {
	g := Goblin(t)
	g.Describe("GetHeroes", func() {
		g.It("Should return no error", func() {
			api, _ := LoadConfig("config.yaml")
			_, err := api.GetHeroes()
			g.Assert(err).Equal(nil)
		})
		g.It("Should return at least one hero", func() {
			api, _ := LoadConfig("config.yaml")
			heroes, _ := api.GetHeroes()
			g.Assert(len(heroes.heroes) > 0).IsTrue()
		})
		g.It("Should work with concurrent usage", func() {
			api, _ := LoadConfig("config.yaml")
			var wg sync.WaitGroup
			wg.Add(10)
			for i := 0; i < 10; i++ {
				go func() {
					heroes, err := api.GetHeroes()
					g.Assert(err).Equal(nil)
					g.Assert(len(heroes.heroes) > 0).IsTrue()
					wg.Done()
				}()
			}
			wg.Wait()
		})
		g.It("Should fill cache", func() {
			api, _ := LoadConfig("config.yaml")
			_, _ = api.GetHeroes()
			heroes, err := api.getHeroesFromCache()
			g.Assert(err).Equal(nil)
			g.Assert(len(heroes.heroes) > 0).IsTrue()
		})
	})
}

func TestHeroes_GetById(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when ID is found", func() {
			for _, hero := range heroes.heroes {
				h, found := heroes.GetById(hero.ID)
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when ID is not found", func() {
			missingId := 1
			for _, hero := range heroes.heroes {
				missingId += hero.ID
			}
			_, found := heroes.GetById(missingId)
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkHeroes_GetById(b *testing.B) {
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	ids := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ids[i] = heroes.heroes[rand.Int()%len(heroes.heroes)].ID
	}
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		heroes.GetById(ids[i])
	}
}

func TestHeroes_GetByName(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when Name is found", func() {
			for _, hero := range heroes.heroes {
				h, found := heroes.GetByName(hero.Name)
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when Name is not found", func() {
			_, found := heroes.GetByName("irNe9GNzJm")
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkHeroes_GetByName(b *testing.B) {
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	names := make([]string, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		names[i] = heroes.heroes[rand.Int()%len(heroes.heroes)].Name
	}
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		heroes.GetByName(names[i])
	}
}
