package dota2api

import (
	. "github.com/franela/goblin"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGetItems(t *testing.T) {
	g := Goblin(t)
	g.Describe("GetItems", func() {
		g.It("Should return no error", func() {
			api, _ := LoadConfig("config.ini")
			_, err := api.GetItems()
			g.Assert(err).Equal(nil)
		})
		g.It("Should return at least one hero", func() {
			api, _ := LoadConfig("config.ini")
			items, _ := api.GetItems()
			g.Assert(len(items.items) > 0).IsTrue()
		})
		g.It("Should work with concurrent usage", func() {
			api, _ := LoadConfig("config.ini")
			var wg sync.WaitGroup
			wg.Add(10)
			for i := 0; i < 10; i++ {
				go func() {
					items, err := api.GetItems()
					g.Assert(err).Equal(nil)
					g.Assert(len(items.items) > 0).IsTrue()
					wg.Done()
				}()
			}
			wg.Wait()
		})
		g.It("Should fill cache", func() {
			api, _ := LoadConfig("config.ini")
			_, _ = api.GetItems()
			items, err := api.getItemsFromCache()
			g.Assert(err).Equal(nil)
			g.Assert(len(items.items) > 0).IsTrue()
		})
	})
}

func TestItems_GetById(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.ini")
	items, _ := api.GetItems()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when ID is found", func() {
			for _, hero := range items.items {
				h, found := items.GetById(hero.ID)
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when ID is not found", func() {
			missingId := 1
			for _, hero := range items.items {
				missingId += hero.ID
			}
			_, found := items.GetById(missingId)
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkItems_GetById(b *testing.B) {
	api, _ := LoadConfig("config.ini")
	items, _ := api.GetItems()
	ids := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ids[i] = items.items[rand.Int()%len(items.items)].ID
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		items.GetById(ids[i])
	}
}

func TestItems_GetByName(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.ini")
	items, _ := api.GetItems()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when Name is found", func() {
			for _, hero := range items.items {
				h, found := items.GetByName(hero.Name)
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when Name is not found", func() {
			_, found := items.GetByName("irNe9GNzJm")
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkItems_GetByName(b *testing.B) {
	api, _ := LoadConfig("config.ini")
	items, _ := api.GetItems()
	names := make([]string, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		names[i] = items.items[rand.Int()%len(items.items)].Name
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		items.GetByName(names[i])
	}
}
