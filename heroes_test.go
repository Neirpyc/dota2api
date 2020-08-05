package dota2api

import (
	. "github.com/franela/goblin"
	"math/rand"
	"regexp"
	"sync"
	"testing"
	"time"
)

func TestGetHeroes(t *testing.T) {
	g := Goblin(t)
	g.Describe("GetHeroes", func() {
		g.It("Should return no error", func(done Done) {
			go func() {
				api, _ := LoadConfig("config.yaml")
				_, err := api.GetHeroes()
				g.Assert(err).Equal(nil)
				done()
			}()
		})
		g.It("Should return at least one hero", func(done Done) {
			go func() {
				api, _ := LoadConfig("config.yaml")
				heroes, _ := api.GetHeroes()
				g.Assert(len(heroes.heroes) > 0).IsTrue()
				done()
			}()
		})
		g.It("Should work with concurrent usage", func(done Done) {
			go func() {
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
				done()
			}()
		})
		g.It("Should have the correct count", func(done Done) {
			go func() {
				api, _ := LoadConfig("config.yaml")
				heroes, _ := api.GetHeroes()
				g.Assert(len(heroes.heroes)).Equal(heroes.Count())
				done()
			}()
		})
		g.It("Should fill cache", func(done Done) {
			go func() {
				api, _ := LoadConfig("config.yaml")
				_, _ = api.GetHeroes()
				heroes, err := api.getHeroesFromCache()
				g.Assert(err).Equal(nil)
				g.Assert(len(heroes.heroes) > 0).IsTrue()
				done()
			}()
		})
	})
}

func TestHeroes_Name(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Heroes Names", func() {
		g.It("Should return the correct full name", func() {
			match, _ := regexp.Match("^"+heroPrefix, []byte(heroes.heroes[0].Name.GetFullName()))
			g.Assert(match).IsTrue()
		})
		g.It("Should return the correct prefix", func() {
			match, _ := regexp.Match(heroPrefix+"$", []byte(heroes.heroes[0].Name.GetPrefix()))
			g.Assert(match).IsTrue()
		})
		g.It("Should return the correct name", func() {
			match, _ := regexp.Match("^"+heroPrefix, []byte(heroes.heroes[0].Name.GetName()))
			g.Assert(match).IsFalse()
		})
	})
}

func TestHeroes_ForEach(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Hero.ForEach", func() {
		g.It("Should work on synchronous request", func() {
			c := 0
			heroes.ForEach(func(hero Hero) {
				if hero.ID == 0 && hero.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c++
			})
			if c != heroes.Count() {
				g.Fail("Skipped element in for each")
			}
		})
		g.It("Should work on asynchronous request", func() {
			c := make(chan int, heroes.Count())
			heroes.ForEachAsync(func(hero Hero) {
				if hero.ID == 0 && hero.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c <- 1
			})
			for i := 0; i < heroes.Count(); i++ {
				select {
				case <-c:
					continue
				default:
					g.Fail("Skipped element in for each")
				}
			}
			heroes.GoForEach(func(hero Hero) {
				if hero.ID == 0 && hero.Name.GetName() == "" {
					g.Fail("Empty element in for each")
				}
				c <- 1
			})()
			for i := 0; i < heroes.Count(); i++ {
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

func TestDota2_GetHeroImage(t *testing.T) {
	g := Goblin(t)
	var wg sync.WaitGroup
	api, _ := LoadConfig("config.yaml")
	heroes, _ := api.GetHeroes()
	doTest := func(size int) {
		for i := 0; i < heroes.Count(); i += heroes.Count() / 10 {
			if i > heroes.Count() {
				continue
			}
			wg.Add(1)
			i := i
			go func() {
				img, err := api.GetHeroImage(heroes.heroes[i], size)
				g.Assert(err).Equal(nil)
				g.Assert(img == nil).IsFalse()
				g.Assert(img.Bounds().Dx() > 0).IsTrue()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	g.Describe("api.GetHeroImage", func() {
		g.It("Should return lg Images", func(done Done) {
			doTest(SizeLg)
			done()
		})
		g.It("Should return sb Images", func(done Done) {
			doTest(SizeSb)
			done()
		})
		g.It("Should return full Images", func(done Done) {
			doTest(SizeFull)
			done()
		})
		g.It("Should return vert Images", func(done Done) {
			doTest(SizeVert)
			done()
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
				h, found := heroes.GetByName(hero.Name.GetFullName())
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
		names[i] = heroes.heroes[rand.Int()%len(heroes.heroes)].Name.GetFullName()
	}
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		heroes.GetByName(names[i])
	}
}
