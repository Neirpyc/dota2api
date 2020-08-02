# Dota2-API

## Introduction
This project is a fork from [l2x](https://github.com/l2x) /[dota2api](https://github.com/l2x/dota2api)

It should be used to fetch information from the Dota webapi, 
but provides higher level wrappers, so it is easier to use.

## Installation
```shell script
go get github.com/Neirpyc/dota2api
```
## Setup
Reads settings from the file `config.ini`.

Here is an example:
```ini
---
Timeout: 10
SteamApiKey:  "<Your Steam API Key>"
SteamApi: "https://api.steampowered.com"
SteamUser: "ISteamUser"
SteamApiVersion: "V001"
Dota2Match: "IDOTA2Match_570"
Dota2Econ: "IEconDOTA2_570"
Dota2ApiVersion: "V001"

```

## Usage

This project is in active development, therefore, new functions will be added, existing one 
(mostly legacy function from l2x's code) will be removed, and new ones could be modified.

Yet, here are examples for the currently working non-legacy function:

#### Heroes
These functions should be used to get the name or the ID of a hero
```go
heroes, err := api.GetHeroes()
if err != nil{
    panic(err)
}
h, found := heroes.GetById(10)
if !found{
    panic("ID not found")
}
fmt.Printf("Name: %s, ID: %d\n", h.Name.GetFullName(), h.ID)
h, found = heroes.GetByName("npc_dota_hero_antimage")
if !found{
    panic("Name not found")
}
fmt.Printf("Name: %s, ID: %d\n", h.Name.GetFullName(), h.ID)
```

#### Items
These functions should be used to get the name or the ID of a hero
```go
items, err := api.GetItems
    panic(err)
}
i, found := items.GetById(10)
if !found{
    panic("ID not found")
}
fmt.Printf("Name: %s, ID: %d, Cost:%d\n", i.Name.GetFullName(), i.ID, i.Cost)
i, found = items.GetByName("item_blink")
if !found{
    panic("Name not found")
}
fmt.Printf("Name: %s, ID: %d, Cost:%d\n", i.Name.GetFullName(), i.ID, i.Cost)
```
