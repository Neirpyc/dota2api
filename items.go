package dota2api

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
)

func getItemsUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2EconUrl, "GetGameItems", dota2.dota2ApiVersion)
}

type itemsJSON struct {
	Result struct {
		Items  []Item `json:"items" bson:"items"`
		Status int    `json:"status" bson:"status"`
	}
}

type Item struct {
	ID         int    `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Cost       int    `json:"cost" bson:"cost"`
	SecretShop int    `json:"secret_shop" bson:"secret_shop"`
	SideShop   int    `json:"side_shop" bson:"side_shop"`
	Recipe     int    `json:"recipe" bson:"recipe"`
}

type Items struct {
	items []Item
}

// Returns the item which has the given id
// If no matching item is found, found = false, otherwise, found = true
//
// First tries with the index [id-1] which sometimes works, and is very fast to test
// If it doesn't work, it then run a dichotomy search.
func (i Items) GetById(id int) (item Item, found bool) {
	if id < len(i.items)-1 {
		if i.items[id-1].ID == id {
			return i.items[id-1], true
		}
	}
	beg, end := 0, len(i.items)-1
	for beg <= end {
		curr := (beg + end) / 2
		if i.items[curr].ID == id {
			return i.items[curr], true
		}
		if id > i.items[curr].ID {
			beg = curr + 1
		} else {
			end = curr - 1
		}
	}
	return Item{}, false
}

// Returns the item which has the given name
// If no matching item is found, found = false, otherwise, found = true
//
// Runs a linear search
func (i Items) GetByName(name string) (item Item, found bool) {
	for _, currentItem := range i.items {
		if currentItem.Name == name {
			return currentItem, true
		}
	}
	return Item{}, false
}

type getItemsCache struct {
	items     Items
	fromCache uint32
	mutex     sync.Mutex
}

// This function calls the API to get the list of the items
// Once a call has succeeded, the result is stored, and no further API call is made
// Instead, it returns a copy of the cached result
func (d *Dota2) GetItems() (Items, error) {
	var err error
	if atomic.LoadUint32(&d.itemsCache.fromCache) == 0 {
		if d.itemsCache.items, err = d.getItemsFromAPI(); err == nil {
			atomic.StoreUint32(&d.itemsCache.fromCache, 1)
		}
	}
	return d.itemsCache.items, err
}

func (d *Dota2) getItemsFromAPI() (Items, error) {
	d.itemsCache.mutex.Lock()
	defer d.itemsCache.mutex.Unlock()
	if d.itemsCache.fromCache == 0 {
		var itemsListJson itemsJSON
		var items Items

		param := map[string]interface{}{
			"key": d.steamApiKey,
		}
		url, err := parseUrl(getItemsUrl(d), param)

		if err != nil {
			return items, err
		}
		resp, err := Get(url)
		if err != nil {
			return items, err
		}

		err = json.Unmarshal(resp, &itemsListJson)
		if err != nil {
			return items, err
		}

		items.items = itemsListJson.Result.Items

		sort.Slice(items.items, func(i, j int) bool {
			return items.items[i].ID < items.items[j].ID
		})

		return items, nil
	}
	return d.getItemsFromCache()
}

func (d *Dota2) getItemsFromCache() (Items, error) {
	return d.itemsCache.items, nil
}

func (i Items) Count() int {
	return len(i.items)
}

func (i Items) ForEach(f func(item Item)) {
	for _, item := range i.items {
		f(item)
	}
}
