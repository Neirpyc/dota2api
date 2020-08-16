package dota2api

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func (api Dota2) getMatchDetailsUrl() string {
	return fmt.Sprintf("%s/%s/%s/", api.dota2MatchUrl, "GetMatchDetails", api.dota2ApiVersion)
}

type matchDetailsJSON struct {
	Result matchJSON `json:"result"`
}

type matchJSON struct {
	Error                 string          `json:"error" json:"error" bson:"error"`
	Players               []playerJSON    `json:"players" bson:"players"`
	RadiantWin            bool            `json:"radiant_win" bson:"radiant_win"`
	Duration              int             `json:"duration" bson:"duration"`
	PreGameDuration       int             `json:"pre_game_duration" bson:"pre_game_duration"`
	StartTime             int64           `json:"start_time" bson:"start_time"`
	MatchID               int64           `json:"match_id" bson:"match_id"`
	MatchSeqNum           int64           `json:"match_seq_num" bson:"match_seq_num"`
	TowerStatusRadiant    uint16          `json:"tower_status_radiant" bson:"tower_status_radiant"`
	TowerStatusDire       uint16          `json:"tower_status_dire" bson:"tower_status_dire"`
	BarracksStatusRadiant uint8           `json:"barracks_status_radiant" bson:"barracks_status_radiant"`
	BarracksStatusDire    uint8           `json:"barracks_status_dire" bson:"barracks_status_dire"`
	Cluster               int             `json:"cluster" bson:"cluster"`
	FirstBloodTime        int             `json:"first_blood_time" bson:"first_blood_time"`
	LobbyType             int             `json:"lobby_type" bson:"lobby_type"`
	HumanPlayers          int             `json:"human_players" bson:"human_players"`
	Leagueid              int             `json:"league_id" bson:"league_id"`
	PositiveVotes         int             `json:"positive_votes" bson:"positive_votes"`
	NegativeVotes         int             `json:"negative_votes" bson:"negative_votes"`
	GameMode              int             `json:"game_mode" bson:"game_mode"`
	Flags                 int             `json:"flags" bson:"flags"`
	Engine                int             `json:"engine" bson:"engine"`
	RadiantScore          int             `json:"radiant_score" bson:"radiant_score"`
	DireScore             int             `json:"dire_score" bson:"dire_score"`
	TournamentID          int             `json:"tournament_id" bson:"tournament_id"`
	TournamentRound       int             `json:"tournament_round" bson:"tournament_round"`
	RadiantTeamID         int             `json:"radiant_team_id" bson:"radiant_team_id"`
	RadiantName           string          `json:"radiant_name" bson:"radiant_name"`
	RadiantLogo           int             `json:"radiant_logo" bson:"radiant_logo"`
	RadiantTeamComplete   int             `json:"radiant_team_complete" bson:"radiant_team_complete"`
	DireTeamID            int             `json:"dire_team_id" bson:"dire_team_id"`
	DireName              string          `json:"dire_name" bson:"dire_name"`
	DireLogo              int             `json:"dire_logo" bson:"dire_logo"`
	DireTeamComplete      int             `json:"dire_team_complete" bson:"dire_team_complete"`
	RadiantCaptain        int             `json:"radiant_captain" bson:"radian_captain"`
	DireCaptain           int             `json:"dire_captain" bson:"dire_captain"`
	PicksBans             []picksBansJSON `json:"picks_bans" bson:"picks_bans"`
}

type playerJSON struct {
	AccountID         int                  `json:"account_id" bson:"account_id"`
	PlayerSlot        int                  `json:"player_slot" bson:"player_slot"`
	HeroID            int                  `json:"hero_id" bson:"hero_id"`
	Item0             int                  `json:"item_0" bson:"item_0"`
	Item1             int                  `json:"item_1" bson:"item_1"`
	Item2             int                  `json:"item_2" bson:"item_2"`
	Item3             int                  `json:"item_3" bson:"item_3"`
	Item4             int                  `json:"item_4" bson:"item_4"`
	Item5             int                  `json:"item_5" bson:"item_5"`
	ItemNeutral       int                  `json:"item_neutral" bson:"item_neutral"`
	Backpack0         int                  `json:"backpack_0" bson:"backpack_0"`
	Backpack1         int                  `json:"backpack_1" bson:"backpack_1"`
	Backpack2         int                  `json:"backpack_2" bson:"backpack_2"`
	Kills             int                  `json:"kills" bson:"kills"`
	Deaths            int                  `json:"deaths" bson:"deaths"`
	Assists           int                  `json:"assists" bson:"assists"`
	LeaverStatus      int                  `json:"leaver_status" bson:"lever_status"`
	LastHits          int                  `json:"last_hits" bson:"last_hits"`
	Denies            int                  `json:"denies" bson:"denies"`
	GoldPerMin        int                  `json:"gold_per_min" bson:"gold_per_min"`
	XpPerMin          int                  `json:"xp_per_min" bson:"xp_per_min"`
	Level             int                  `json:"Level" bson:"Level"`
	Gold              int                  `json:"gold" bson:"gold"`
	GoldSpent         int                  `json:"gold_spent" bson:"gold_spent"`
	HeroDamage        int                  `json:"hero_damage" bson:"hero_damage"`
	ScaledHeroDamage  int                  `json:"scaled_hero_damage" bson:"scaled_hero_damage"`
	TowerDamage       int                  `json:"tower_damage" bson:"tower_damage"`
	ScaledTowerDamage int                  `json:"scaled_tower_damage" bson:"scaled_tower_damage"`
	HeroHealing       int                  `json:"hero_healing" bson:"hero_healing"`
	ScaledHeroHealing int                  `json:"scaled_hero_healing" bson:"scaled_hero_healing"`
	AbilityUpgrades   []abilityUpgradeJSON `json:"ability_upgrades" bson:"ability_upgrades"`
}

type picksBansJSON struct {
	IsPick bool `json:"is_pick" bson:"is_pick"`
	HeroId int  `json:"hero_id" bson:"hero_id"`
	Team   int  `json:"team" bson:"team"`
	Order  int  `json:"order" bson:"order"`
}

type abilityUpgradeJSON struct {
	Ability int `json:"ability" bson:"ability"`
	Level   int `json:"Level" bson:"Level"`
	Time    int `json:"time" bson:"time"`
}

type abilityUpgradesJSON []abilityUpgradeJSON

type MatchDetails struct {
	Radiant         TeamDetails
	Dire            TeamDetails
	Victory         Victory
	Duration        time.Duration
	PreGameDuration time.Duration
	StartTime       time.Time
	MatchID         int64
	MatchSeqNum     int64
	BuildingsState  BuildingsState
	Cluster         int
	FirstBloodTime  time.Duration
	LobbyType       LobbyType
	HumanPlayers    int
	Votes           Votes
	GameMode        GameMode
	Flags           int
	Engine          Engine
	Score           Score
	PicksBans       PicksBans
}

type PicksBans []PickBan

func (p PicksBans) GetPick(order int) (PickBan, bool) {
	beg, end := 0, len(p)-1
	for beg <= end {
		curr := (beg + end) / 2
		if p[curr].Order == order {
			return p[curr], true
		}
		if order > p[curr].Order {
			beg = curr + 1
		} else {
			end = curr - 1
		}
	}
	return PickBan{}, false
}

func (p PicksBans) GetPickByHero(hero Hero) (PickBan, bool) {
	for _, pickBan := range p {
		if pickBan.Hero.ID == hero.ID {
			return pickBan, true
		}
	}
	return PickBan{}, false
}

func (p PicksBans) GetByTeam(team Side) PicksBans {
	var ret PicksBans
	for _, pickBan := range p {
		if pickBan.team == int(team) {
			ret = append(ret, pickBan)
		}
	}
	return ret
}

func (p PicksBans) GetByPickType(pickType PickType) PicksBans {
	var ret PicksBans
	for _, pickBan := range p {
		if pickBan.GetType() == pickType {
			ret = append(ret, pickBan)
		}
	}
	return ret
}

type PickBan struct {
	isPick bool
	Hero   Hero
	team   int
	Order  int
}

type PickType int

const (
	Pick PickType = iota
	Ban
)

func (p PickBan) IsPick() bool {
	return p.isPick
}

func (p PickBan) IsBan() bool {
	return !p.isPick
}

func (p PickBan) GetType() PickType {
	if p.isPick {
		return Pick
	}
	return Ban
}

func (p PickBan) IsRadiant() bool {
	return p.team == int(RadiantVictory)
}

func (p PickBan) IsDire() bool {
	return p.team == int(DireVictory)
}

func (p PickBan) GetTeam() int {
	return p.team
}

type AbilityUpgrades []AbilityUpgrade

func (a AbilityUpgrades) Count() int {
	return len(a)
}

func (a AbilityUpgrades) GetByAbility(ability int) AbilityUpgrades {
	var ret []AbilityUpgrade
	for _, aU := range a {
		if aU.Ability == ability {
			ret = append(ret, aU)
		}
	}
	return ret
}

func (a AbilityUpgrades) GetByOrder(order int) (AbilityUpgrade, bool) {
	if order >= 0 && order < len(a) {
		return a[order], true
	}
	return AbilityUpgrade{}, false
}

func (a AbilityUpgrades) GetByLevel(level int) (AbilityUpgrade, bool) {
	beg, end := 0, len(a)-1
	for beg <= end {
		curr := (beg + end) / 2
		if a[curr].Level == level {
			return a[curr], true
		}
		if level > a[curr].Level {
			beg = curr + 1
		} else {
			end = curr - 1
		}
	}
	return AbilityUpgrade{}, false
}

type AbilityUpgrade struct {
	Level   int
	Time    time.Duration
	Ability int
}

type Engine int

const (
	Source1 Engine = iota
	Source2
)

type TeamDetails []PlayerDetails

func (t TeamDetails) Count() int {
	return len(t)
}

type PlayerDetails struct {
	AccountId       int
	Hero            Hero
	Items           PlayersItems
	KDA             KDA
	LeaverStatus    LeaverStatus
	Stats           PlayerStats
	AbilityUpgrades AbilityUpgrades
}

type PlayerStats struct {
	LastHits      int
	Denies        int
	GoldPerMinute int
	XpPerMinute   int
	Level         int
	HeroDamage    Damage
	TowerDamage   Damage
	HeroHealing   Damage
	Gold          PlayerGold
}

type PlayerGold struct {
	current int
	spent   int
}

func (p PlayerGold) Current() Gold {
	return Gold(p.current)
}

func (p PlayerGold) Spent() Gold {
	return Gold(p.spent)
}

func (p PlayerGold) NetWorth() Gold {
	return Gold(p.current + p.spent)
}

type Gold int

func (g Gold) Raw() int {
	return int(g)
}

type Damage struct {
	raw    int
	scaled int
}

func (d Damage) Raw() int {
	return d.raw
}

func (d Damage) Scaled() int {
	return d.scaled
}

func (d Damage) ScalingFactor() float64 {
	return float64(d.scaled) / float64(d.raw)
}

func (g Gold) ToString() string {
	g += 50
	if g < 1000 {
		return strconv.Itoa(int(g - 50))
	}
	if (g-(g/1000*1000))/100 > 0 {
		return fmt.Sprintf("%d.%dk", g/1000, g%1000/100)
	}
	return fmt.Sprintf("%dk", g%1000)
}

type LeaverStatus int

const (
	LeaverStatusNone LeaverStatus = iota
	LeaverStatusDisconnected
	LeaverStatusDisconnectedTooLong
	LeaverStatusAbandoned
	LeaverStatusAFK
	LeaverStatusNeverConnected
	LeaverStatusNeverConnectedTooLong
)

type PlayersItems struct {
	Item0         Item
	Item1         Item
	Item2         Item
	Item3         Item
	Item4         Item
	Item5         Item
	ItemNeutral   Item
	BackpackItem0 Item
	BackpackItem1 Item
	BackpackItem2 Item
}

type KDA struct {
	Kills   int
	Deaths  int
	Assists int
}

type Side int

const (
	Radiant Side = iota
	Dire
)

type Victory int

const (
	RadiantVictory Victory = iota
	DireVictory
)

func (v Victory) RadiantWon() bool {
	return v == RadiantVictory
}

func (v Victory) DireWon() bool {
	return v == DireVictory
}

func (v Victory) GetWinningTeam() Side {
	return Side(v)
}

type BuildingsState struct {
	Dire    TeamBuildingsState
	Radiant TeamBuildingsState
}

type TeamBuildingsState struct {
	Top        LaneBuildingsState
	Mid        LaneBuildingsState
	Bot        LaneBuildingsState
	T4TowerBot bool
	T4TowerTop bool
}

type LaneBuildingsState struct {
	T1Tower       bool
	T2Tower       bool
	T3Tower       bool
	MeleeBarrack  bool
	RangedBarrack bool
}

func (t TeamBuildingsState) from(towersState uint16, barracksState uint8) TeamBuildingsState {
	setAndShiftTower := func(b *bool) {
		*b = towersState&1 == 1
		towersState = towersState >> 1
	}
	setAndShiftTower(&t.Top.T1Tower)
	setAndShiftTower(&t.Top.T2Tower)
	setAndShiftTower(&t.Top.T3Tower)
	setAndShiftTower(&t.Mid.T1Tower)
	setAndShiftTower(&t.Mid.T2Tower)
	setAndShiftTower(&t.Mid.T3Tower)
	setAndShiftTower(&t.Bot.T1Tower)
	setAndShiftTower(&t.Bot.T2Tower)
	setAndShiftTower(&t.Bot.T3Tower)
	setAndShiftTower(&t.T4TowerTop)
	setAndShiftTower(&t.T4TowerBot)
	setAndShiftBarracks := func(b *bool) {
		*b = barracksState&1 == 1
		barracksState = barracksState >> 1
	}
	setAndShiftBarracks(&t.Top.MeleeBarrack)
	setAndShiftBarracks(&t.Top.RangedBarrack)
	setAndShiftBarracks(&t.Mid.MeleeBarrack)
	setAndShiftBarracks(&t.Mid.RangedBarrack)
	setAndShiftBarracks(&t.Bot.MeleeBarrack)
	setAndShiftBarracks(&t.Bot.RangedBarrack)
	return t
}

type Votes struct {
	PositiveVotes int
	NegativeVotes int
}

type GameMode int

//todo update GameModes strings
func (g GameMode) GetString() string {
	switch g {
	case GameModeNone:
		return "None"
	case GameModeAllPick:
		return "All Pick"
	case GameModeCaptainsDraft:
		return "Captain's Draft"
	case GameModeRandomDraft:
		return "Random Draft"
	case GameModeSingleDraft:
		return "Single Draft"
	case GameModeAllRandom:
		return "All Random"
	case GameModeIntro:
		return "Into"
	case GameModeDireTide:
		return "Dire Tide"
	case GameModeReverseCaptainsMode:
		return "Reverse Captain's Mode"
	case GameModeTheGreeviling:
		return "The Greeviling"
	case GameModeTutorial:
		return "Tutorial"
	case GameModeMidOnly:
		return "Mid Only"
	case GameModeLeastPlayed:
		return "Least Played"
	case GameModeNewPlayerPool:
		return "New Player Pool"
	case GameModeCompendiumMatchmaking:
		return "Compendium Matchmaking"
	case GameModeCoopVsBots:
		return "Co-op vs Bots"
	case GameModeAllRandomDeathmatch:
		return "All Random Deathmatch"
	case GameMode1v1MidOnly:
		return "1v1 Mid Only"
	case GameModeRankedMatchmaking:
		return "Ranked Matchmaking"
	case GameModeTurboMode:
		return "Turbo Mode"
	default:
		return "Unknown"
	}
}

//todo update GameModes
const (
	GameModeNone GameMode = iota
	GameModeAllPick
	GameModeCaptainsDraft
	GameModeRandomDraft
	GameModeSingleDraft
	GameModeAllRandom
	GameModeIntro
	GameModeDireTide
	GameModeReverseCaptainsMode
	GameModeTheGreeviling
	GameModeTutorial
	GameModeMidOnly
	GameModeLeastPlayed
	GameModeNewPlayerPool
	GameModeCompendiumMatchmaking
	GameModeCoopVsBots
	GameModeAllRandomDeathmatch
	GameMode1v1MidOnly
	GameModeRankedMatchmaking
	GameModeTurboMode
)

type Score struct {
	RadiantScore int
	DireScore    int
}

func (p playerJSON) toPlayerDetails(heroes Heroes, items Items) PlayerDetails {
	h, _ := heroes.GetById(p.HeroID)

	i0, _ := items.GetById(p.Item0)
	i1, _ := items.GetById(p.Item1)
	i2, _ := items.GetById(p.Item2)
	i3, _ := items.GetById(p.Item3)
	i4, _ := items.GetById(p.Item4)
	i5, _ := items.GetById(p.Item5)
	iN, _ := items.GetById(p.ItemNeutral)
	iB0, _ := items.GetById(p.Backpack0)
	iB1, _ := items.GetById(p.Backpack1)
	iB2, _ := items.GetById(p.Backpack2)
	player := PlayerDetails{
		AccountId: p.AccountID,
		Hero:      h,
		Items: PlayersItems{
			Item0:         i0,
			Item1:         i1,
			Item2:         i2,
			Item3:         i3,
			Item4:         i4,
			Item5:         i5,
			ItemNeutral:   iN,
			BackpackItem0: iB0,
			BackpackItem1: iB1,
			BackpackItem2: iB2,
		},
		KDA: KDA{
			Kills:   p.Kills,
			Deaths:  p.Deaths,
			Assists: p.Assists,
		},
		LeaverStatus: LeaverStatus(p.LeaverStatus),
		Stats: PlayerStats{
			LastHits:      p.LastHits,
			Denies:        p.Denies,
			GoldPerMinute: p.GoldPerMin,
			XpPerMinute:   p.XpPerMin,
			Level:         p.Level,
			HeroDamage: Damage{
				raw:    p.HeroDamage,
				scaled: p.ScaledHeroDamage,
			},
			TowerDamage: Damage{
				raw:    p.TowerDamage,
				scaled: p.ScaledTowerDamage,
			},
			HeroHealing: Damage{
				raw:    p.HeroHealing,
				scaled: p.ScaledHeroHealing,
			},
			Gold: PlayerGold{
				current: p.Gold,
				spent:   p.GoldSpent,
			},
		},
		AbilityUpgrades: make(AbilityUpgrades, len(p.AbilityUpgrades)),
	}
	return player
}

func (a abilityUpgradesJSON) toAbilityUpgrade() AbilityUpgrades {
	ret := make([]AbilityUpgrade, len(a))
	for i, aU := range a {
		ret[i] = AbilityUpgrade{
			Time:    time.Duration(int64(aU.Time) * int64(time.Second)),
			Ability: aU.Ability,
			Level:   aU.Level,
		}
	}
	return ret
}

//Get match details
func (api Dota2) GetMatchDetails(params ...Parameter) (MatchDetails, error) {

	var matchDetails matchDetailsJSON
	var match MatchDetails

	param, err := getParameterMap([]int{parameterKindMatchId}, nil, params)
	if err != nil {
		return match, err
	}
	param["key"] = api.steamApiKey
	url, err := parseUrl(api.getMatchDetailsUrl(), param)

	if err != nil {
		return match, err
	}
	resp, err := api.Get(url)
	if err != nil {
		return match, err
	}

	err = json.Unmarshal(resp, &matchDetails)
	if err != nil {
		return match, err
	}

	if matchDetails.Result.Error != "" {
		return match, errors.New(matchDetails.Result.Error)
	}

	match = MatchDetails{
		Radiant: make([]PlayerDetails, 0),
		Dire:    make([]PlayerDetails, 0),
		Victory: func() Victory {
			if matchDetails.Result.RadiantWin {
				return RadiantVictory
			}
			return DireVictory
		}(),
		Duration:        time.Duration(int64(matchDetails.Result.Duration) * int64(time.Second)),
		PreGameDuration: time.Duration(int64(matchDetails.Result.PreGameDuration) * int64(time.Second)),
		StartTime:       time.Unix(matchDetails.Result.StartTime, 0),
		MatchID:         matchDetails.Result.MatchID,
		MatchSeqNum:     matchDetails.Result.MatchSeqNum,
		BuildingsState: BuildingsState{
			Dire: TeamBuildingsState{}.from(matchDetails.Result.TowerStatusDire,
				matchDetails.Result.BarracksStatusDire),
			Radiant: TeamBuildingsState{}.from(matchDetails.Result.TowerStatusRadiant,
				matchDetails.Result.BarracksStatusRadiant),
		},
		Cluster:        matchDetails.Result.Cluster,
		FirstBloodTime: time.Duration(int64(matchDetails.Result.FirstBloodTime) * int64(time.Second)),
		LobbyType:      LobbyType(matchDetails.Result.LobbyType),
		HumanPlayers:   matchDetails.Result.HumanPlayers,
		Votes: Votes{
			PositiveVotes: matchDetails.Result.PositiveVotes,
			NegativeVotes: matchDetails.Result.NegativeVotes,
		},
		GameMode: GameMode(matchDetails.Result.GameMode),
		Flags:    matchDetails.Result.Flags,
		Engine:   Engine(matchDetails.Result.Engine),
		Score: Score{
			RadiantScore: matchDetails.Result.RadiantScore,
			DireScore:    matchDetails.Result.DireScore,
		},
	}

	heroes, err := api.GetHeroes()
	if err != nil {
		return match, err
	}

	items, err := api.GetItems()
	if err != nil {
		return match, err
	}

	for _, player := range matchDetails.Result.Players {
		p := player.toPlayerDetails(heroes, items)

		p.AbilityUpgrades = abilityUpgradesJSON(player.AbilityUpgrades).toAbilityUpgrade()

		sort.Slice(p.AbilityUpgrades, func(i, j int) bool {
			return p.AbilityUpgrades[i].Level < p.AbilityUpgrades[j].Level
		})

		if player.PlayerSlot&128 == 0 {
			match.Radiant = append(match.Radiant, p)
		} else {
			match.Dire = append(match.Dire, p)
		}
	}

	for _, pickBan := range matchDetails.Result.PicksBans {
		h, _ := heroes.GetById(pickBan.HeroId)
		match.PicksBans = append(match.PicksBans, PickBan{
			isPick: pickBan.IsPick,
			Hero:   h,
			team:   pickBan.Team,
			Order:  pickBan.Order,
		})
	}

	return match, nil
}

func MatchId(matchId int64) ParameterInt64 {
	return ParameterInt64{
		k:       "match_id",
		v:       matchId,
		kindInt: parameterKindMatchId,
	}
}
