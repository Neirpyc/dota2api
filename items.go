package dota2api

import (
	"bytes"
	"encoding/json"
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

func getItemsUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.dota2EconUrl, "GetGameItems", dota2.dota2ApiVersion)
}

func getItemImageUrl(d *Dota2, name itemName) string {
	return fmt.Sprintf("%s/items/%s_lg.png", d.dota2CDN, name.name)
}

type itemsJSON struct {
	Result struct {
		Items  []ItemJSON `json:"items" bson:"items"`
		Status int        `json:"status" bson:"status"`
	}
}

type ItemJSON struct {
	ID         int    `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Cost       int    `json:"cost" bson:"cost"`
	SecretShop int    `json:"secret_shop" bson:"secret_shop"`
	SideShop   int    `json:"side_shop" bson:"side_shop"`
	Recipe     int    `json:"recipe" bson:"recipe"`
}

type itemName struct {
	name   string
	prefix string
	full   string
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
func (i Items) GetById(id int) (item Item, found bool) {
	if id < len(i.items) && id > 0 {
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
	return
}

func (i Items) GetByPos(pos int) Item {
	return i.items[pos]
}

// Returns the item which has the given name
// If no matching item is found, found = false, otherwise, found = true
//
// Runs a linear search
func (i Items) GetByName(name string) (item Item, found bool) {
	for _, currentItem := range i.items {
		if currentItem.Name.full == name {
			return currentItem, true
		}
	}
	return Item{}, false
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
func (d *Dota2) GetItems() (Items, error) {
	var err error
	if atomic.LoadUint32(&d.itemsCache.fromCache) == 0 {
		err = d.fillItemsCache()
	}
	return d.itemsCache.items, err
}

func (d *Dota2) fillItemsCache() error {
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
			return err
		}
		resp, err := d.Get(url)
		if err != nil {
			return err
		}

		err = json.Unmarshal(resp, &itemsListJson)
		if err != nil {
			return err
		}

		items.items = make([]Item, len(itemsListJson.Result.Items))
		for i, src := range itemsListJson.Result.Items {
			items.items[i] = Item{
				ID:         src.ID,
				Name:       itemNameFromFullName(src.Name),
				Cost:       0,
				SecretShop: src.SecretShop == 1,
				SideShop:   src.SideShop == 1,
				Recipe:     src.Recipe == 1,
			}
		}

		sort.Slice(items.items, func(i, j int) bool {
			return items.items[i].ID < items.items[j].ID
		})

		d.itemsCache.items = items
		atomic.StoreUint32(&d.itemsCache.fromCache, 1)
		return nil
	}
	return nil
}

func (d *Dota2) getItemsFromCache() (Items, error) {
	return d.itemsCache.items, nil
}

func (i Items) Count() int {
	return len(i.items)
}

func (d Dota2) GetItemImage(item Item) (img image.Image, err error) {
	url := getItemImageUrl(&d, item.Name)
	res, err := d.Get(url)
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
