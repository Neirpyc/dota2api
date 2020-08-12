package dota2api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strconv"
	"strings"
	"time"
)

func getPlayerSummariesUrl(dota2 *Dota2) string {
	return fmt.Sprintf("%s/%s/%s/", dota2.steamUserUrl, "GetPlayerSummaries", "V002")
}

type playerSummariesJSON struct {
	Response struct {
		Players []playerAccountJSON `json:"players" bson:"players"`
	} `json:"response" bson:"response"`
}

type playerAccountJSON struct {
	Steamid                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	LastLogoff               int64  `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	AvatarHash               string `json:"avatarhash"`
	PersonaState             int    `json:"personastate"`
	PersonaStateFlags        int    `json:"personastateflags"`
}

func (p playerSummariesJSON) toPlayerAccounts(rawResponse []byte) PlayerAccounts {
	var ret = make(PlayerAccounts, len(p.Response.Players))
	for i, player := range p.Response.Players {
		ret[i] = PlayerAccount{
			SteamId: NewSteamIdFrom64(func() uint64 {
				i, err := strconv.ParseUint(player.Steamid, 10, 64)
				if err != nil {
					panic(err)
				}
				return i
			}()),
			CommunityVisibilityState: CommunityVisibilityState(player.CommunityVisibilityState),
			ProfileState:             ProfileState(player.ProfileState),
			DisplayName:              player.PersonaName,
			LastLogOff:               time.Unix(player.LastLogoff, 0),
			ProfileUrl:               player.ProfileUrl,
			Avatar: Avatar{
				Avatar32Url:  player.Avatar,
				Avatar64Url:  player.AvatarMedium,
				Avatar184Url: player.AvatarMedium,
				Hash:         player.AvatarHash,
			},
			UserStatus: UserStatus(player.PersonaState),
			/*Optional: func() Optional {
				var o Optional
				reader := bytes.NewReader(rawResponse)
				dec := json.NewDecoder(reader)
				dec.DisallowUnknownFields()
				gameId
				if v, err := dec.Decode()
			}(),*/
		}
	}
	return ret
}

type PlayerAccounts []PlayerAccount

type CommunityVisibilityState int

const (
	VisibilityPrivate CommunityVisibilityState = iota + 1
	VisibilityFriendOnly
	VisibilityFriendsOfFriends
	VisibilityUsersOnly
	VisibilityPublic
)

type ProfileState int

const (
	ProfileStateEmpty ProfileState = iota
	ProfileStateConfigured
)

type Avatar struct {
	Avatar32Url  string
	Avatar64Url  string
	Avatar184Url string
	Hash         string
}

func (a Avatar) getUrl(url string) (img image.Image, err error) {
	var ret []byte
	ret, err = Get(url)
	if err != nil {
		return
	}
	rd := bytes.NewReader(ret)
	img, err = jpeg.Decode(rd)
	if err != nil {
		img, err = png.Decode(rd)
	}
	return
}

func (a Avatar) Avatar32() (image.Image, error) {
	return a.getUrl(a.Avatar32Url)
}

func (a Avatar) Avatar64() (image.Image, error) {
	return a.getUrl(a.Avatar64Url)
}

func (a Avatar) Avatar184() (image.Image, error) {
	return a.getUrl(a.Avatar184Url)
}

type UserStatus int

const (
	UserStatusOffline UserStatus = iota
	UserStatusOnline
	UserStatusBusy
	UserStatusAway
	UserStatusSnooze
	UserStatusLookingToTrade
	UserStatusLookingToPlay
)

type Optional struct {
	gameIdPresent            bool
	gameId                   string
	gameNamePresent          bool
	gameName                 string
	primaryClanIdPresent     bool
	primaryClanId            uint64
	commentPermissionPresent bool
	timeCreatedPresent       bool
	timeCreated              time.Time
	locCountryCodePresent    bool
	locCountryCode           string
	locStateCodePresent      bool
	locStateCode             string
	locCityIdPresent         bool
	locCityId                uint64
	gameServerIpPresent      bool
	gameServerIp             string
}

func (o Optional) GameId() (gameId string, present bool) {
	return o.gameId, o.gameIdPresent
}

func (o Optional) GameName() (gameName string, present bool) {
	return o.gameName, o.gameNamePresent
}

func (o Optional) PrimaryClanId() (primaryClanId uint64, present bool) {
	return o.primaryClanId, o.primaryClanIdPresent
}

func (o Optional) CommentPermission() (commentPermission bool) {
	return o.commentPermissionPresent
}

func (o Optional) TimeCreated() (timeCreated time.Time, present bool) {
	return o.timeCreated, o.timeCreatedPresent
}

func (o Optional) LocCountryCode() (code string, present bool) {
	return o.locCountryCode, o.locCountryCodePresent
}

func (o Optional) LocStateCode() (code string, present bool) {
	return o.locStateCode, o.locStateCodePresent
}

func (o Optional) LocCityId() (id uint64, present bool) {
	return o.locCityId, o.locCityIdPresent
}

func (o Optional) GameServerIp() (ip string, present bool) {
	return o.gameServerIp, o.gameServerIpPresent
}

type PlayerAccount struct {
	SteamId                  SteamId
	CommunityVisibilityState CommunityVisibilityState
	ProfileState             ProfileState
	DisplayName              string
	LastLogOff               time.Time
	ProfileUrl               string
	Avatar                   Avatar
	UserStatus               UserStatus
	Optional                 Optional //todo get a working Optional
}

//Get player summaries
func (d *Dota2) GetPlayerSummaries(steamIds []int64) (PlayerAccounts, error) {
	var playerAccounts PlayerAccounts
	var playerSummaries playerSummariesJSON

	param := map[string]interface{}{
		"key":      d.steamApiKey,
		"steamids": strings.Join(ArrayIntToStr(steamIds), ","),
	}
	url, err := parseUrl(getPlayerSummariesUrl(d), param)

	if err != nil {
		return playerAccounts, err
	}
	resp, err := Get(url)
	if err != nil {
		return playerAccounts, err
	}

	err = json.Unmarshal(resp, &playerSummaries)
	if err != nil {
		return playerAccounts, err
	}

	return playerSummaries.toPlayerAccounts(resp), nil
}
