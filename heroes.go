package dota2api

import (
	"bytes"
	"encoding/json"
	"errors"
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

type HeroImageSize int

const (
	SizeLg HeroImageSize = iota
	SizeSb
	SizeFull
	SizeVert
)

func (api Dota2) getHeroesUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.dota2EconUrl, "GetHeroes", api.dota2ApiVersion)
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
	if id < len(h.heroes) && id > 0 {
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
	name string
	full string
}

func (hN heroName) GetName() string {
	return hN.name
}

func (hN heroName) GetFullName() string {
	return hN.full
}

func (hN heroName) GetPrefix() string {
	return heroPrefix
}

func heroNameFromFullName(name string) heroName {
	return heroName{
		name: name[len(heroPrefix):],
		full: name,
	}
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
func (api *Dota2) GetHeroes() (Heroes, error) {
	var err error

	if atomic.LoadUint32(&api.heroesCache.fromCache) == 0 {
		err = api.fillHeroesCache()
	}
	return api.heroesCache.heroes, err
}

func (h heroesJSON) toHeroes() Heroes {
	var heroes Heroes
	heroes.heroes = make([]Hero, len(h.Result.Heroes))
	for i, hero := range h.Result.Heroes {
		heroes.heroes[i] = hero.toHero()
	}
	sort.Slice(heroes.heroes, func(i, j int) bool {
		return heroes.heroes[i].ID < heroes.heroes[j].ID
	})
	return heroes
}

func (h heroJSON) toHero() Hero {
	return Hero{
		ID: h.ID,
		Name: heroName{
			name: h.Name[len(heroPrefix):],
			full: h.Name,
		},
	}
}

func (api Dota2) fillHeroesCache() error {
	api.heroesCache.mutex.Lock()
	defer api.heroesCache.mutex.Unlock()
	if api.heroesCache.fromCache == 0 {
		var heroesListJson heroesJSON

		param := map[string]interface{}{
			"key": api.steamApiKey,
		}
		url, err := parseUrl(api.getHeroesUrl(), param)

		if err != nil {
			return err
		}
		resp, err := api.Get(url)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(resp, &heroesListJson); err != nil {
			return err
		}
		if heroesListJson.Result.Status != 200 {
			return errors.New("non 200 status code")
		}

		api.heroesCache.heroes = heroesListJson.toHeroes()
		atomic.StoreUint32(&api.heroesCache.fromCache, 1)
		return nil
	}
	return nil
}

func (h Heroes) Count() int {
	return len(h.heroes)
}

func getHeroImageUrl(d Dota2, name heroName, size HeroImageSize) string {
	ext := "png"
	if size == SizeVert {
		ext = "jpg"
	}
	return fmt.Sprintf("%s/heroes/%s_%s.%s", d.dota2CDN, name.name, sizes[size], ext)
}

func (api Dota2) GetHeroImage(hero Hero, size HeroImageSize) (image.Image, error) {
	url := getHeroImageUrl(api, hero.Name, size)
	res, err := api.Get(url)
	if err != nil {
		return nil, err
	}
	var img image.Image
	if size == SizeVert {
		img, err = jpeg.Decode(bytes.NewReader(res))
	}
	if size != SizeVert || err != nil {
		img, err = png.Decode(bytes.NewReader(res))
	}
	if err != nil && size != SizeVert {
		img, err = jpeg.Decode(bytes.NewReader(res))
	}
	return img, err
}
