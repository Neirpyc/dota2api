package dota2api

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

func getHeroesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.Dota2EconUrl, "GetHeroes", dota2.Dota2ApiVersion)
}

type heroesJSON struct {
	Result struct {
		Count  int    `json:"count" bson:"count"`
		Heroes []Hero `json:"heroes" bson:"heroes"`
		Status int    `json:"status" bson:"status"`
	} `json:"result" bson:"result"`
}

type Heroes struct {
	heroes []Hero
}

func (h *Heroes) GetById(id int) (hero Hero, found bool) {
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

func (h *Heroes) GetByName(name string) (hero Hero, found bool) {
	for _, currentHero := range h.heroes {
		if currentHero.Name == name {
			return currentHero, true
		}
	}
	return Hero{}, false
}

type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type getHeroesCache struct {
	heroes          *Heroes
	heroesFromCache uint32
	getHeroesMutex  sync.Mutex
}

func newGetHeroesCache() *getHeroesCache {
	ret := getHeroesCache{
		heroesFromCache: 0,
	}
	return &ret
}

//Get all heroes
func (d *Dota2) GetHeroes() (*Heroes, error) {
	var err error
	if atomic.LoadUint32(&d.heroesCache.heroesFromCache) == 0 {
		if d.heroesCache.heroes, err = d.getHeroesFromAPI(); err == nil {
			atomic.StoreUint32(&d.heroesCache.heroesFromCache, 1)
		}
	}
	return d.heroesCache.heroes, err
}

func (d *Dota2) getHeroesFromAPI() (*Heroes, error) {
	d.heroesCache.getHeroesMutex.Lock()
	defer d.heroesCache.getHeroesMutex.Unlock()
	if d.heroesCache.heroesFromCache == 0 {
		var heroesListJson heroesJSON
		var heroes Heroes

		param := map[string]interface{}{
			"key": d.SteamApiKey,
		}
		url, err := parseUrl(getHeroesUrl(d), param)

		if err != nil {
			return nil, err
		}
		resp, err := Get(url)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resp, &heroesListJson)
		if err != nil {
			return nil, err
		}

		heroes.heroes = heroesListJson.Result.Heroes

		sort.Slice(heroes.heroes, func(i, j int) bool {
			return heroes.heroes[i].ID < heroes.heroes[j].ID
		})

		return &heroes, nil
	}
	return d.getHeroesFromCache()
}

func (d *Dota2) getHeroesFromCache() (*Heroes, error) {
	return d.heroesCache.heroes, nil
}