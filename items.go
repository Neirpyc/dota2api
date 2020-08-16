package dota2api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"sort"
	"sync"
	"sync/atomic"
)

const itemPrefix = "item_"

func (api Dota2) getItemsUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.dota2EconUrl, "GetGameItems", api.dota2ApiVersion)
}

func getItemImageUrl(d *Dota2, name itemName) string {
	return fmt.Sprintf("%s/items/%s_lg.png", d.dota2CDN, name.name)
}

type itemsJSON struct {
	Result struct {
		Items  []itemJSON `json:"items" bson:"items"`
		Status int        `json:"status" bson:"status"`
	}
}

type itemJSON struct {
	ID         int    `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Cost       int    `json:"cost" bson:"cost"`
	SecretShop int    `json:"secret_shop" bson:"secret_shop"`
	SideShop   int    `json:"side_shop" bson:"side_shop"`
	Recipe     int    `json:"recipe" bson:"recipe"`
}

type itemName struct {
	name string
	full string
}

func (iN itemName) GetName() string {
	return iN.name
}

func (iN itemName) GetFullName() string {
	return iN.full
}

func (iN itemName) GetPrefix() string {
	return itemPrefix
}

func itemNameFromFullName(name string) itemName {
	return itemName{
		name: name[len(itemPrefix):],
		full: name,
	}
}

type Item struct {
	ID         int
	Name       itemName
	Cost       int
	SecretShop bool
	SideShop   bool
	Recipe     bool
}

type Items struct {
	items []Item
}

// Returns the item which has the given id
// If no matching item is found, found = false, otherwise, found = true
//
// First tries with the index [id-1] which sometimes works, and is very fast to test
// If it doesn't work, it then run a dichotomy search.
func (i Items) GetById(id int) (Item, error) {
	if id < len(i.items) && id > 0 {
		if i.items[id-1].ID == id {
			return i.items[id-1], nil
		}
	}
	beg, end := 0, len(i.items)-1
	for beg <= end {
		curr := (beg + end) / 2
		if i.items[curr].ID == id {
			return i.items[curr], nil
		}
		if id > i.items[curr].ID {
			beg = curr + 1
		} else {
			end = curr - 1
		}
	}
	return Item{}, errors.New(fmt.Sprintf("item with id %d not found", id))
}

func (i Items) GetByPos(pos int) Item {
	return i.items[pos]
}

// Returns the item which has the given name
// If no matching item is found, found = false, otherwise, found = true
//
// Runs a linear search
func (i Items) GetByName(name string) (Item, error) {
	for _, currentItem := range i.items {
		if currentItem.Name.full == name {
			return currentItem, nil
		}
	}
	return Item{}, errors.New(fmt.Sprintf("item with name %s not found", name))
}

func (i Items) Item(pos int) Item {
	return i.items[pos]
}

type getItemsCache struct {
	items     Items
	fromCache uint32
	mutex     sync.Mutex
}

// This function calls the API to get the list of the items
// Once a call has succeeded, the result is stored, and no further API call is made
// Instead, it returns a copy of the cached result
func (api *Dota2) GetItems() (Items, error) {
	var err error
	if atomic.LoadUint32(&api.itemsCache.fromCache) == 0 {
		err = api.fillItemsCache()
	}
	return api.itemsCache.items, err
}

func (i itemsJSON) toItems() Items {
	var items Items
	items.items = make([]Item, len(i.Result.Items))
	for i, item := range i.Result.Items {
		items.items[i] = item.toItem()
	}
	sort.Slice(items.items, func(i, j int) bool {
		return items.items[i].ID < items.items[j].ID
	})
	return items
}

func (i itemJSON) toItem() Item {
	return Item{
		ID: i.ID,
		Name: itemName{
			name: i.Name[len(itemPrefix):],
			full: i.Name,
		},
		Cost:       i.Cost,
		SecretShop: i.SecretShop == 1,
		SideShop:   i.SideShop == 1,
		Recipe:     i.Recipe == 1,
	}
}

func (api *Dota2) fillItemsCache() error {
	api.itemsCache.mutex.Lock()
	defer api.itemsCache.mutex.Unlock()
	if api.itemsCache.fromCache == 0 {
		var itemsListJson itemsJSON

		param := map[string]interface{}{
			"key": api.steamApiKey,
		}
		url, err := parseUrl(api.getItemsUrl(), param)

		if err != nil {
			return err
		}
		resp, err := api.Get(url)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(resp, &itemsListJson); err != nil {
			return err
		}

		if itemsListJson.Result.Status != 200 {
			return errors.New("non 200 status code")
		}

		api.itemsCache.items = itemsListJson.toItems()
		atomic.StoreUint32(&api.itemsCache.fromCache, 1)
		return nil
	}
	return nil
}

func (api *Dota2) getItemsFromCache() (Items, error) {
	return api.itemsCache.items, nil
}

func (i Items) Count() int {
	return len(i.items)
}

func (api Dota2) GetItemImage(item Item) (img image.Image, err error) {
	url := getItemImageUrl(&api, item.Name)
	res, err := api.Get(url)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(res)
	img, err = png.Decode(r)
	if err != nil {
		_, err = r.Seek(0, io.SeekStart)
		if err != nil {
			return
		}
		img, err = jpeg.Decode(r)
	}
	return
}
