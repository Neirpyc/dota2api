package dota2api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strconv"
	"time"
)

func (api Dota2) getPlayerSummariesUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.steamUserUrl, "GetPlayerSummaries", "V002")
}

type playerSummariesJSON struct {
	Response struct {
		Players []playerAccountJSON `json:"players" bson:"players"`
	} `json:"response" bson:"response"`
}

type playerAccountJSON struct {
	Steamid                  string          `json:"steamid"`
	CommunityVisibilityState int             `json:"communityvisibilitystate"`
	ProfileState             int             `json:"profilestate"`
	PersonaName              string          `json:"personaname"`
	LastLogoff               int64           `json:"lastlogoff"`
	ProfileUrl               string          `json:"profileurl"`
	Avatar                   string          `json:"avatar"`
	AvatarMedium             string          `json:"avatarmedium"`
	AvatarFull               string          `json:"avatarfull"`
	AvatarHash               string          `json:"avatarhash"`
	PersonaState             int             `json:"personastate"`
	PersonaStateFlags        int             `json:"personastateflags"`
	CommentPermission        json.RawMessage `json:"commentpermission"`
	RealName                 json.RawMessage `json:"realname"`
	PrimaryClanId            json.RawMessage `json:"primaryclanid"`
	TimeCreated              json.RawMessage `json:"timecreated"`
	LocCountryCode           json.RawMessage `json:"loccountrycode"`
	LocStateCode             json.RawMessage `json:"locstatecode"`
	LocCityId                json.RawMessage `json:"loccityid"`
	GameId                   json.RawMessage `json:"gameid"`
	GameExtraInfo            json.RawMessage `json:"gameextrainfo"`
	GameServerIp             json.RawMessage `json:"gameserverip"`
}

func (p playerSummariesJSON) toPlayerAccounts(api *Dota2) PlayerAccounts {
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
			PersonaStateFlag:         player.PersonaStateFlags,
			Avatar: Avatar{
				Avatar32Url:  player.Avatar,
				Avatar64Url:  player.AvatarMedium,
				Avatar184Url: player.AvatarFull,
				Hash:         player.AvatarHash,
				api:          api,
			},
			UserStatus: UserStatus(player.PersonaState),
			Optional: func() (o Optional) {
				checkField := func(src json.RawMessage) (string, bool) {
					if len(src) > 0 {
						src = bytes.TrimRight(bytes.TrimLeft(src, "\""), "\"")
						return string(src), true
					}
					return "", false
				}
				var tmp string
				var err error
				_, o.commentPermissionPresent = checkField(player.CommentPermission)
				o.realName, o.realNamePresent = checkField(player.RealName)
				if tmp, o.primaryClanIdPresent = checkField(player.PrimaryClanId); o.primaryClanIdPresent {
					if o.primaryClanId, err = strconv.ParseUint(tmp, 10, 64); err != nil {
						panic(err)
					}
				}
				if tmp, o.timeCreatedPresent = checkField(player.TimeCreated); o.timeCreatedPresent {
					if unix, err := strconv.ParseInt(tmp, 10, 64); err != nil {
						panic(err)
					} else {
						o.timeCreated = time.Unix(unix, 0)
					}
				}
				o.locCountryCode, o.locCountryCodePresent = checkField(player.LocCountryCode)
				o.locStateCode, o.locStateCodePresent = checkField(player.LocStateCode)
				if tmp, o.locCityIdPresent = checkField(player.LocCityId); o.locCityIdPresent {
					if o.locCityId, err = strconv.ParseUint(tmp, 10, 64); err != nil {
						panic(err)
					}
				}
				o.gameId, o.gameIdPresent = checkField(player.GameId)
				o.gameName, o.gameNamePresent = checkField(player.GameExtraInfo)
				o.gameServerIp, o.gameServerIpPresent = checkField(player.GameServerIp)

				return
			}(),
		}
	}
	return ret
}

type PlayerAccounts []PlayerAccount

type CommunityVisibilityState int

const (
	VisibilityPrivate CommunityVisibilityState = iota + 1
	VisibilityFriendsOnly
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
	api          *Dota2
}

func (a Avatar) getUrl(url string) (img image.Image, err error) {
	var ret []byte
	ret, err = a.api.Get(url)
	if err != nil {
		return
	}
	rd := bytes.NewReader(ret)
	img, err = jpeg.Decode(rd)
	if err != nil {
		_, err = rd.Seek(0, io.SeekStart)
		if err != nil {
			return
		}
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
	realNamePresent          bool
	realName                 string
}

func (o Optional) GameId() (gameId string, present bool) {
	return o.gameId, o.gameIdPresent
}

func (o Optional) RealName() (realName string, present bool) {
	return o.realName, o.realNamePresent
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
	PersonaStateFlag         int
	Optional                 Optional
}

//Get player summaries
func (api *Dota2) GetPlayerSummaries(params ...Parameter) (PlayerAccounts, error) {
	var playerAccounts PlayerAccounts
	var playerSummaries playerSummariesJSON

	param, err := getParameterMap([]parameterKind{parameterSteamIds}, nil, params)
	if err != nil {
		return playerAccounts, err
	}
	param["key"] = api.steamApiKey

	url, err := parseUrl(api.getPlayerSummariesUrl(), param)

	if err != nil {
		return playerAccounts, err
	}
	resp, err := api.Get(url)
	if err != nil {
		return playerAccounts, err
	}

	err = json.Unmarshal(resp, &playerSummaries)
	if err != nil {
		return playerAccounts, err
	}

	return playerSummaries.toPlayerAccounts(api), nil
}

func ParameterSteamIds(ids ...SteamId) Parameter {
	idsUint64 := make([]uint64, len(ids))
	for i, id := range ids {
		current, err := id.SteamId64()
		if err != nil {
			panic(err)
		}
		idsUint64[i] = current
	}
	return ParameterString{
		k:       "steamids",
		v:       ArrayIntToStr(idsUint64),
		kindInt: parameterSteamIds,
	}
}
