package dota2api

import (
	"bytes"
	. "github.com/franela/goblin"
	"image"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	response0       = `{"response":{"players":[{"steamid":"42","communityvisibilitystate":3,"profilestate":1,"personaname":"userNAME","profileurl":"profileURL","avatar":"avatar32URL","avatarmedium":"avatar64URL","avatarfull":"avatar184URL","avatarhash":"avatarHASH","lastlogoff":42,"personastate":3,"personastateflags":64}]}}`
	response1       = `{"response":{"players":[{"steamid":"42","communityvisibilitystate":3,"profilestate":1,"personaname":"userNAME","profileurl":"profileURL","avatar":"avatar32URL","avatarmedium":"avatar64URL","avatarfull":"avatar184URL","avatarhash":"avatarHASH","lastlogoff":42,"personastate":3,"personastateflags":64},{"steamid":"43","communityvisibilitystate":2,"profilestate":0,"personaname":"userNAME0","profileurl":"profileURL0","avatar":"avatar32URL0","avatarmedium":"avatar64URL0","avatarfull":"avatar184URL0","avatarhash":"avatarHASH0","lastlogoff":43,"personastate":2,"personastateflags":65}]}}`
	responseWithOpt = `{"response":{"players":[{"steamid":"42","commentpermission":"15","realname":"realNAME","primaryclanid":42,"timecreated":43,"loccountrycode":"locCOUNTRYcode","locstatecode":"locSTATEcode","loccityid":44,"gameid":"gameID","gameextrainfo":"gameEXTRAinfo","gameserverip":"gameSERVERip"}]}}`
)

func TestDota2_GetPlayerSummaries(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("api.GetPlayerSummaries", func() {
		g.Describe("Basic response", func() {
			var sum PlayerAccounts
			var err error
			g.It("Should call the correct request URI", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					g.Assert(req.URL.String()).Equal(api.getPlayerSummariesUrl() + "?key=keyTEST&steamids=%5B42%5D")
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response0))}, nil
				}
				sum, err = api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 1 summary", func() {
				g.Assert(len(sum)).Equal(1)
			})
			g.It("Should parse non Optional flags", func() {
				g.Assert(sum).Equal(PlayerAccounts{PlayerAccount{
					SteamId:                  NewSteamIdFrom64(42),
					CommunityVisibilityState: VisibilityFriendsOfFriends,
					ProfileState:             ProfileStateConfigured,
					DisplayName:              "userNAME",
					LastLogOff:               time.Unix(42, 0),
					ProfileUrl:               "profileURL",
					Avatar: Avatar{
						Avatar32Url:  "avatar32URL",
						Avatar64Url:  "avatar64URL",
						Avatar184Url: "avatar184URL",
						Hash:         "avatarHASH",
						api:          &api,
					},
					UserStatus:       UserStatusAway,
					PersonaStateFlag: 64,
					Optional:         Optional{},
				}})
			})
		})
		g.Describe("Double response", func() {
			var sum PlayerAccounts
			var err error
			g.It("Should call the correct request URI", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					g.Assert(req.URL.String()).Equal(api.getPlayerSummariesUrl() + "?key=keyTEST&steamids=%5B42%2C43%5D")
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response1))}, nil
				}
				sum, err = api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42), NewSteamIdFrom64(43)))
			})
			g.It("Should return no error", func() {
				g.Assert(err).IsNil()
			})
			g.It("Should return 2 summaries", func() {
				g.Assert(len(sum)).Equal(2)
			})
			g.It("Should parse non Optional flags", func() {
				g.Assert(sum).Equal(PlayerAccounts{PlayerAccount{
					SteamId:                  NewSteamIdFrom64(42),
					CommunityVisibilityState: VisibilityFriendsOfFriends,
					ProfileState:             ProfileStateConfigured,
					DisplayName:              "userNAME",
					LastLogOff:               time.Unix(42, 0),
					ProfileUrl:               "profileURL",
					Avatar: Avatar{
						Avatar32Url:  "avatar32URL",
						Avatar64Url:  "avatar64URL",
						Avatar184Url: "avatar184URL",
						Hash:         "avatarHASH",
						api:          &api,
					},
					UserStatus:       UserStatusAway,
					PersonaStateFlag: 64,
					Optional:         Optional{},
				}, PlayerAccount{
					SteamId:                  NewSteamIdFrom64(43),
					CommunityVisibilityState: VisibilityFriendsOnly,
					ProfileState:             ProfileStateEmpty,
					DisplayName:              "userNAME0",
					LastLogOff:               time.Unix(43, 0),
					ProfileUrl:               "profileURL0",
					Avatar: Avatar{
						Avatar32Url:  "avatar32URL0",
						Avatar64Url:  "avatar64URL0",
						Avatar184Url: "avatar184URL0",
						Hash:         "avatarHASH0",
						api:          &api,
					},
					UserStatus:       UserStatusBusy,
					PersonaStateFlag: 65,
					Optional:         Optional{},
				}})
			})
		})
		g.Describe("Optional fields", func() {
			g.It("Should report missing fields", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response0))}, nil
				}
				sum, _ := api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
				g.Assert(sum[0].Optional).Equal(Optional{})
			})
			g.It("Should return correct values when found", func() {
				mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(responseWithOpt))}, nil
				}
				sum, _ := api.GetPlayerSummaries(ParameterSteamIds(NewSteamIdFrom64(42)))
				g.Assert(sum[0].Optional).Equal(Optional{
					gameIdPresent:            true,
					gameId:                   "gameID",
					gameNamePresent:          true,
					gameName:                 "gameEXTRAinfo",
					primaryClanIdPresent:     true,
					primaryClanId:            42,
					commentPermissionPresent: true,
					timeCreatedPresent:       true,
					timeCreated:              time.Unix(43, 0),
					locCountryCodePresent:    true,
					locCountryCode:           "locCOUNTRYcode",
					locStateCodePresent:      true,
					locStateCode:             "locSTATEcode",
					locCityIdPresent:         true,
					locCityId:                44,
					gameServerIpPresent:      true,
					gameServerIp:             "gameSERVERip",
					realNamePresent:          true,
					realName:                 "realNAME",
				})
			})
		})
	})
}

func TestAvatar_Avatar(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	a := Avatar{
		Avatar32Url:  "avatar32",
		Avatar64Url:  "avatar64",
		Avatar184Url: "avatar184",
		api:          &api,
	}
	getFunc := func(jpg bool, name string) func(req *http.Request) (*http.Response, error) {
		return func(req *http.Request) (*http.Response, error) {
			g.Assert(req.URL.String()).Equal(name)
			var b []byte
			if jpg {
				b = getJpgTest()
			} else {
				b = getPngTest()
			}
			return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(b))}, nil
		}
	}
	check := func(img image.Image, err error) {
		g.Assert(err).IsNil()
		g.Assert(validateTestImage(img)).IsTrue()
	}
	var img image.Image
	var err error
	g.Describe("avatar.Avatar32", func() {
		g.It("Should query the correct URL", func() {
			mockClient.DoFunc = getFunc(true, "avatar32")
			img, err = a.Avatar32()
		})
		g.It("Should decode jpeg image", func() {
			check(img, err)
		})
		g.It("Should decode png image", func() {
			mockClient.DoFunc = getFunc(false, "avatar32")
			img, err = a.Avatar32()
			check(img, err)
		})
	})
	g.Describe("avatar.Avatar64", func() {
		g.It("Should query the correct URL", func() {
			mockClient.DoFunc = getFunc(true, "avatar64")
			img, err = a.Avatar64()
		})
		g.It("Should decode jpeg image", func() {
			check(img, err)
		})
		g.It("Should decode png image", func() {
			mockClient.DoFunc = getFunc(false, "avatar64")
			img, err = a.Avatar64()
			check(img, err)
		})
	})
	g.Describe("avatar.Avatar184", func() {
		g.It("Should query the correct URL", func() {
			mockClient.DoFunc = getFunc(true, "avatar184")
			img, err = a.Avatar184()
		})
		g.It("Should decode jpeg image", func() {
			check(img, err)
		})
		g.It("Should decode png image", func() {
			mockClient.DoFunc = getFunc(false, "avatar184")
			img, err = a.Avatar184()
			check(img, err)
		})
	})
}
