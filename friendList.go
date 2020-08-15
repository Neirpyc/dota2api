package dota2api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func getFriendListUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "GetFriendList", dota2.steamApiVersion)
}

type friendListJSON struct {
	Friendslist struct {
		Friends []friendJSON `json:"friends"`
	} `json:"friendslist"`
}

func (f friendListJSON) ToFriends() Friends {
	friends := make(Friends, len(f.Friendslist.Friends))
	for i, friend := range f.Friendslist.Friends {
		friends[i] = Friend{
			SteamId: func() SteamId {
				if id, err := strconv.ParseUint(friend.SteamId, 10, 64); err != nil {
					panic(err)
				} else {
					return NewSteamIdFrom64(id)
				}
			}(),
			RelationShip: friend.Relationship,
			FriendsSince: time.Unix(friend.FriendSince, 0),
		}
	}
	return friends
}

type friendJSON struct {
	SteamId      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int64  `bson:"friend_since" json:"friend_since"`
}

type Friend struct {
	SteamId      SteamId
	RelationShip string
	FriendsSince time.Time
}

type Friends []Friend

func (f Friends) Count() int {
	return len(f)
}

//Get friend list
func (d *Dota2) GetFriendList(params ...Parameter) (Friends, error) {
	var friendList friendListJSON

	param, err := getParameterMap([]int{parameterSteamId}, nil, params)
	if err != nil {
		return Friends{}, err
	}

	param["key"] = d.steamApiKey
	url, err := parseUrl(getFriendListUrl(d), param)
	if err != nil {
		return Friends{}, err
	}

	resp, err := d.Get(url)
	if err != nil {
		return Friends{}, err
	}

	err = json.Unmarshal(resp, &friendList)
	if err != nil {
		return Friends{}, err
	}

	return friendList.ToFriends(), nil
}

func ParameterSteamId(s SteamId) ParameterInt64 {
	return ParameterInt64{
		k:       "steamid",
		v:       int64(s.id),
		kindInt: parameterSteamId,
	}
}