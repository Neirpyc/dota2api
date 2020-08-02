package dota2api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"sort"
	"sync"
	"sync/atomic"
)

var sizes = []string{"lg", "sb", "full", "vert"}

const (
	heroPrefix = "npc_dota_hero_"
)

const (
	SizeLg = iota
	SizeSb
	SizeFull
	SizeVert
)

func getHeroesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2EconUrl, "GetHeroes", dota2.dota2ApiVersion)
}

type heroesJSON struct {
	Result struct {
		Count  int        `json:"count" bson:"count"`
		Heroes []heroJSON `json:"heroes" bson:"heroes"`
		Status int        `json:"status" bson:"status"`
	} `json:"result" bson:"result"`
}

type Heroes struct {
	heroes []Hero
}

// Returns the hero which has the given id
// If no matching hero is found, found = false, otherwise, found = true
//
// First tries with the index [id-1] which sometimes works, and is very fast to test
// If it doesn't work, it then run a dichotomy search.
func (h Heroes) GetById(id int) (hero Hero, found bool) {
	if id < len(h.heroes)-1 {
		if h.heroes[id-1].ID == id {
			return h.heroes[id-1], true
		}
	}
	beg, end := 0, len(h.heroes)-1
	for beg <= end {
		curr := (beg + end) / 2
		if h.heroes[curr].ID == id {
			return h.heroes[curr], true
		}
		if id > h.heroes[curr].ID {
			beg = curr + 1
		} else {
			end = curr - 1
		}
	}
	return Hero{}, false
}

// Returns the hero which has the given name
// If no matching hero is found, found = false, otherwise, found = true
//
// Runs a linear search
func (h Heroes) GetByName(name string) (hero Hero, found bool) {
	for _, currentHero := range h.heroes {
		if currentHero.Name.full == name {
			return currentHero, true
		}
	}
	return Hero{}, false
}

type heroName struct {
	name   string
	prefix string
	full   string
}

func (hN heroName) GetName() string {
	return hN.name
}

func (hN heroName) GetFullName() string {
	return hN.full
}

func (hN heroName) GetPrefix() string {
	return hN.prefix
}

type Hero struct {
	ID   int
	Name heroName
}

type heroJSON struct {
	ID   int
	Name string
}

type getHeroesCache struct {
	heroes    Heroes
	fromCache uint32
	mutex     sync.Mutex
}

// This function calls the API to get the list of the heroes
// Once a call has succeeded, the result is stored, and no further API call is made
// Instead, it returns a copy of the cached result
func (d *Dota2) GetHeroes() (Heroes, error) {
	var err error

	if atomic.LoadUint32(&d.heroesCache.fromCache) == 0 {
		if d.heroesCache.heroes, err = d.getHeroesFromAPI(); err == nil {
			atomic.StoreUint32(&d.heroesCache.fromCache, 1)
		}
	}
	return d.heroesCache.heroes, err
}

func (d *Dota2) getHeroesFromAPI() (Heroes, error) {
	d.heroesCache.mutex.Lock()
	defer d.heroesCache.mutex.Unlock()
	if d.heroesCache.fromCache == 0 {
		var heroesListJson heroesJSON
		var heroes Heroes

		param := map[string]interface{}{
			"key": d.steamApiKey,
		}
		url, err := parseUrl(getHeroesUrl(d), param)

		if err != nil {
			return heroes, err
		}
		resp, err := Get(url)
		if err != nil {
			return heroes, err
		}

		err = json.Unmarshal(resp, &heroesListJson)
		if err != nil {
			return heroes, err
		}

		heroes.heroes = make([]Hero, len(heroesListJson.Result.Heroes))
		for i, src := range heroesListJson.Result.Heroes {
			heroes.heroes[i] = Hero{
				ID: src.ID,
				Name: heroName{
					name:   src.Name[len(heroPrefix):],
					prefix: heroPrefix,
					full:   src.Name,
				},
			}
		}

		sort.Slice(heroes.heroes, func(i, j int) bool {
			return heroes.heroes[i].ID < heroes.heroes[j].ID
		})

		return heroes, nil
	}
	return d.getHeroesFromCache()
}

func (d *Dota2) getHeroesFromCache() (Heroes, error) {
	return d.heroesCache.heroes, nil
}

func (h Heroes) Count() int {
	return len(h.heroes)
}

func (h Heroes) ForEach(f func(hero Hero)) {
	for _, hero := range h.heroes {
		f(hero)
	}
}

func (h Heroes) GoForEach(f func(hero Hero, wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	wg.Add(len(h.heroes))
	for _, hero := range h.heroes {
		go f(hero, &wg)
	}
	wg.Wait()
}

func (d Dota2) GetHeroImage(hero Hero, size int) (image.Image, error) {
	ext := "png"
	if size == SizeVert {
		ext = "jpg"
	}
	url := fmt.Sprintf("%s/heroes/%s_%s.%s", d.dota2CDN, hero.Name.name, sizes[size], ext)
	res, err := Get(url)
	if err != nil {
		return nil, err
	}
	var img image.Image
	if size == SizeVert {
		img, err = jpeg.Decode(bytes.NewReader(res))
	} else {
		img, err = png.Decode(bytes.NewReader(res))
	}
	return img, nil
}
