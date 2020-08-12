package dota2api

import (
	. "github.com/franela/goblin"
	"math/rand"
	"regexp"
	"sync"
	"testing"
	"time"
)

func TestGetItems(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	g.Describe("GetItems", func() {
		g.It("Should return no error", func(done Done) {
			go func() {
				_, err := api.GetItems()
				g.Assert(err).Equal(nil)
				done()
			}()
		})
		g.It("Should return at least one hero", func(done Done) {
			go func() {
				items, _ := api.GetItems()
				g.Assert(len(items.items) > 0).IsTrue()
				done()
			}()
		})
		g.It("Should work with concurrent usage", func(done Done) {
			go func() {
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
				done()
			}()
		})
		g.It("Should fill cache", func(done Done) {
			go func() {
				_, _ = api.GetItems()
				items, err := api.getItemsFromCache()
				g.Assert(err).Equal(nil)
				g.Assert(len(items.items) > 0).IsTrue()
				done()
			}()
		})
	})
}

func TestItems_Name(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	g.Describe("Items Names", func() {
		g.It("Should return the correct full name", func() {
			match, _ := regexp.Match("^"+itemPrefix, []byte(items.items[0].Name.GetFullName()))
			g.Assert(match).IsTrue()
		})
		g.It("Should return the correct prefix", func() {
			match, _ := regexp.Match(itemPrefix+"$", []byte(items.items[0].Name.GetPrefix()))
			g.Assert(match).IsTrue()
		})
		g.It("Should return the correct name", func() {
			match, _ := regexp.Match("^"+itemPrefix, []byte(items.items[0].Name.GetName()))
			g.Assert(match).IsFalse()
		})
	})
}

func TestDota2_GetItemImage(t *testing.T) {
	g := Goblin(t)
	var wg sync.WaitGroup
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	g.Describe("api.GetHeroImage", func() {
		g.It("Should return lg Images", func() {
			for i := 0; i < items.Count(); i += items.Count() / 10 {
				if i > items.Count() {
					continue
				}
				wg.Add(1)
				i := i
				go func() {
					img, err := api.GetItemImage(items.items[i])
					g.Assert(err).Equal(nil)
					g.Assert(img == nil).IsFalse()
					g.Assert(img.Bounds().Dx() > 0).IsTrue()
					wg.Done()
				}()
			}
		})
	})
}

func TestItems_ForEach(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	g.Describe("Item.ForEach", func() {
		g.It("Should work on synchronous request", func() {
			c := 0
			items.ForEach(func(item Item) {
				if item.ID == 0 && item.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c++
			})
			if c != items.Count() {
				g.Fail("Skipped element in for each")
			}
		})
		g.It("Should work on asynchronous request", func() {
			c := make(chan int, items.Count())
			items.GoForEach(func(item Item) {
				if item.ID == 0 && item.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c <- 1
			})()
			for i := 0; i < items.Count(); i++ {
				select {
				case <-c:
					continue
				default:
					g.Fail("Skipped element in for each")
				}
			}
			items.GoForEach(func(item Item) {
				if item.ID == 0 && item.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c <- 1
			})()
			for i := 0; i < items.Count(); i++ {
				select {
				case <-c:
					continue
				default:
					g.Fail("Skipped element in for each")
				}
			}
		})
	})
}

func TestItems_GetById(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when ID is found", func() {
			for _, item := range items.items {
				i, found := items.GetById(item.ID)
				g.Assert(found).IsTrue()
				g.Assert(i.ID).Equal(item.ID)
				g.Assert(i.Name).Equal(item.Name)
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
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	ids := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ids[i] = items.items[rand.Int()%len(items.items)].ID
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items.GetById(ids[i])
	}
}

func TestItems_GetByName(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when Name is found", func() {
			for _, item := range items.items {
				i, found := items.GetByName(item.Name.GetFullName())
				g.Assert(found).IsTrue()
				g.Assert(i.ID).Equal(item.ID)
				g.Assert(i.Name).Equal(item.Name)
			}
		})
		g.It("Should return Error when Name is not found", func() {
			_, found := items.GetByName("irNe9GNzJm")
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkItems_GetByName(b *testing.B) {
	api, _ := LoadConfig("config.yaml")
	items, _ := api.GetItems()
	names := make([]string, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		names[i] = items.items[rand.Int()%len(items.items)].Name.GetFullName()
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items.GetByName(names[i])
	}
}
