package dota2api

import (
	. "github.com/franela/goblin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	heroesResponse = "{\n\"result\":{\n\"heroes\":[\n{\n\"name\":\"npc_dota_hero_antimage\",\n\"id\":1\n},\n{\n\"name\":\"npc_dota_hero_axe\",\n\"id\":2\n},\n{\n\"name\":\"npc_dota_hero_bane\",\n\"id\":3\n},\n{\n\"name\":\"npc_dota_hero_bloodseeker\",\n\"id\":4\n},\n{\n\"name\":\"npc_dota_hero_crystal_maiden\",\n\"id\":5\n},\n{\n\"name\":\"npc_dota_hero_drow_ranger\",\n\"id\":6\n},\n{\n\"name\":\"npc_dota_hero_earthshaker\",\n\"id\":7\n},\n{\n\"name\":\"npc_dota_hero_juggernaut\",\n\"id\":8\n},\n{\n\"name\":\"npc_dota_hero_mirana\",\n\"id\":9\n},\n{\n\"name\":\"npc_dota_hero_nevermore\",\n\"id\":11\n},\n{\n\"name\":\"npc_dota_hero_morphling\",\n\"id\":10\n},\n{\n\"name\":\"npc_dota_hero_phantom_lancer\",\n\"id\":12\n},\n{\n\"name\":\"npc_dota_hero_puck\",\n\"id\":13\n},\n{\n\"name\":\"npc_dota_hero_pudge\",\n\"id\":14\n},\n{\n\"name\":\"npc_dota_hero_razor\",\n\"id\":15\n},\n{\n\"name\":\"npc_dota_hero_sand_king\",\n\"id\":16\n},\n{\n\"name\":\"npc_dota_hero_storm_spirit\",\n\"id\":17\n},\n{\n\"name\":\"npc_dota_hero_sven\",\n\"id\":18\n},\n{\n\"name\":\"npc_dota_hero_tiny\",\n\"id\":19\n},\n{\n\"name\":\"npc_dota_hero_vengefulspirit\",\n\"id\":20\n},\n{\n\"name\":\"npc_dota_hero_windrunner\",\n\"id\":21\n},\n{\n\"name\":\"npc_dota_hero_zuus\",\n\"id\":22\n},\n{\n\"name\":\"npc_dota_hero_kunkka\",\n\"id\":23\n},\n{\n\"name\":\"npc_dota_hero_lina\",\n\"id\":25\n},\n{\n\"name\":\"npc_dota_hero_lich\",\n\"id\":31\n},\n{\n\"name\":\"npc_dota_hero_lion\",\n\"id\":26\n},\n{\n\"name\":\"npc_dota_hero_shadow_shaman\",\n\"id\":27\n},\n{\n\"name\":\"npc_dota_hero_slardar\",\n\"id\":28\n},\n{\n\"name\":\"npc_dota_hero_tidehunter\",\n\"id\":29\n},\n{\n\"name\":\"npc_dota_hero_witch_doctor\",\n\"id\":30\n},\n{\n\"name\":\"npc_dota_hero_riki\",\n\"id\":32\n},\n{\n\"name\":\"npc_dota_hero_enigma\",\n\"id\":33\n},\n{\n\"name\":\"npc_dota_hero_tinker\",\n\"id\":34\n},\n{\n\"name\":\"npc_dota_hero_sniper\",\n\"id\":35\n},\n{\n\"name\":\"npc_dota_hero_necrolyte\",\n\"id\":36\n},\n{\n\"name\":\"npc_dota_hero_warlock\",\n\"id\":37\n},\n{\n\"name\":\"npc_dota_hero_beastmaster\",\n\"id\":38\n},\n{\n\"name\":\"npc_dota_hero_queenofpain\",\n\"id\":39\n},\n{\n\"name\":\"npc_dota_hero_venomancer\",\n\"id\":40\n},\n{\n\"name\":\"npc_dota_hero_faceless_void\",\n\"id\":41\n},\n{\n\"name\":\"npc_dota_hero_skeleton_king\",\n\"id\":42\n},\n{\n\"name\":\"npc_dota_hero_death_prophet\",\n\"id\":43\n},\n{\n\"name\":\"npc_dota_hero_phantom_assassin\",\n\"id\":44\n},\n{\n\"name\":\"npc_dota_hero_pugna\",\n\"id\":45\n},\n{\n\"name\":\"npc_dota_hero_templar_assassin\",\n\"id\":46\n},\n{\n\"name\":\"npc_dota_hero_viper\",\n\"id\":47\n},\n{\n\"name\":\"npc_dota_hero_luna\",\n\"id\":48\n},\n{\n\"name\":\"npc_dota_hero_dragon_knight\",\n\"id\":49\n},\n{\n\"name\":\"npc_dota_hero_dazzle\",\n\"id\":50\n},\n{\n\"name\":\"npc_dota_hero_rattletrap\",\n\"id\":51\n},\n{\n\"name\":\"npc_dota_hero_leshrac\",\n\"id\":52\n},\n{\n\"name\":\"npc_dota_hero_furion\",\n\"id\":53\n},\n{\n\"name\":\"npc_dota_hero_life_stealer\",\n\"id\":54\n},\n{\n\"name\":\"npc_dota_hero_dark_seer\",\n\"id\":55\n},\n{\n\"name\":\"npc_dota_hero_clinkz\",\n\"id\":56\n},\n{\n\"name\":\"npc_dota_hero_omniknight\",\n\"id\":57\n},\n{\n\"name\":\"npc_dota_hero_enchantress\",\n\"id\":58\n},\n{\n\"name\":\"npc_dota_hero_huskar\",\n\"id\":59\n},\n{\n\"name\":\"npc_dota_hero_night_stalker\",\n\"id\":60\n},\n{\n\"name\":\"npc_dota_hero_broodmother\",\n\"id\":61\n},\n{\n\"name\":\"npc_dota_hero_bounty_hunter\",\n\"id\":62\n},\n{\n\"name\":\"npc_dota_hero_weaver\",\n\"id\":63\n},\n{\n\"name\":\"npc_dota_hero_jakiro\",\n\"id\":64\n},\n{\n\"name\":\"npc_dota_hero_batrider\",\n\"id\":65\n},\n{\n\"name\":\"npc_dota_hero_chen\",\n\"id\":66\n},\n{\n\"name\":\"npc_dota_hero_spectre\",\n\"id\":67\n},\n{\n\"name\":\"npc_dota_hero_doom_bringer\",\n\"id\":69\n},\n{\n\"name\":\"npc_dota_hero_ancient_apparition\",\n\"id\":68\n},\n{\n\"name\":\"npc_dota_hero_ursa\",\n\"id\":70\n},\n{\n\"name\":\"npc_dota_hero_spirit_breaker\",\n\"id\":71\n},\n{\n\"name\":\"npc_dota_hero_gyrocopter\",\n\"id\":72\n},\n{\n\"name\":\"npc_dota_hero_alchemist\",\n\"id\":73\n},\n{\n\"name\":\"npc_dota_hero_invoker\",\n\"id\":74\n},\n{\n\"name\":\"npc_dota_hero_silencer\",\n\"id\":75\n},\n{\n\"name\":\"npc_dota_hero_obsidian_destroyer\",\n\"id\":76\n},\n{\n\"name\":\"npc_dota_hero_lycan\",\n\"id\":77\n},\n{\n\"name\":\"npc_dota_hero_brewmaster\",\n\"id\":78\n},\n{\n\"name\":\"npc_dota_hero_shadow_demon\",\n\"id\":79\n},\n{\n\"name\":\"npc_dota_hero_lone_druid\",\n\"id\":80\n},\n{\n\"name\":\"npc_dota_hero_chaos_knight\",\n\"id\":81\n},\n{\n\"name\":\"npc_dota_hero_meepo\",\n\"id\":82\n},\n{\n\"name\":\"npc_dota_hero_treant\",\n\"id\":83\n},\n{\n\"name\":\"npc_dota_hero_ogre_magi\",\n\"id\":84\n},\n{\n\"name\":\"npc_dota_hero_undying\",\n\"id\":85\n},\n{\n\"name\":\"npc_dota_hero_rubick\",\n\"id\":86\n},\n{\n\"name\":\"npc_dota_hero_disruptor\",\n\"id\":87\n},\n{\n\"name\":\"npc_dota_hero_nyx_assassin\",\n\"id\":88\n},\n{\n\"name\":\"npc_dota_hero_naga_siren\",\n\"id\":89\n},\n{\n\"name\":\"npc_dota_hero_keeper_of_the_light\",\n\"id\":90\n},\n{\n\"name\":\"npc_dota_hero_wisp\",\n\"id\":91\n},\n{\n\"name\":\"npc_dota_hero_visage\",\n\"id\":92\n},\n{\n\"name\":\"npc_dota_hero_slark\",\n\"id\":93\n},\n{\n\"name\":\"npc_dota_hero_medusa\",\n\"id\":94\n},\n{\n\"name\":\"npc_dota_hero_troll_warlord\",\n\"id\":95\n},\n{\n\"name\":\"npc_dota_hero_centaur\",\n\"id\":96\n},\n{\n\"name\":\"npc_dota_hero_magnataur\",\n\"id\":97\n},\n{\n\"name\":\"npc_dota_hero_shredder\",\n\"id\":98\n},\n{\n\"name\":\"npc_dota_hero_bristleback\",\n\"id\":99\n},\n{\n\"name\":\"npc_dota_hero_tusk\",\n\"id\":100\n},\n{\n\"name\":\"npc_dota_hero_skywrath_mage\",\n\"id\":101\n},\n{\n\"name\":\"npc_dota_hero_abaddon\",\n\"id\":102\n},\n{\n\"name\":\"npc_dota_hero_elder_titan\",\n\"id\":103\n},\n{\n\"name\":\"npc_dota_hero_legion_commander\",\n\"id\":104\n},\n{\n\"name\":\"npc_dota_hero_ember_spirit\",\n\"id\":106\n},\n{\n\"name\":\"npc_dota_hero_earth_spirit\",\n\"id\":107\n},\n{\n\"name\":\"npc_dota_hero_terrorblade\",\n\"id\":109\n},\n{\n\"name\":\"npc_dota_hero_phoenix\",\n\"id\":110\n},\n{\n\"name\":\"npc_dota_hero_oracle\",\n\"id\":111\n},\n{\n\"name\":\"npc_dota_hero_techies\",\n\"id\":105\n},\n{\n\"name\":\"npc_dota_hero_winter_wyvern\",\n\"id\":112\n},\n{\n\"name\":\"npc_dota_hero_arc_warden\",\n\"id\":113\n},\n{\n\"name\":\"npc_dota_hero_abyssal_underlord\",\n\"id\":108\n},\n{\n\"name\":\"npc_dota_hero_monkey_king\",\n\"id\":114\n},\n{\n\"name\":\"npc_dota_hero_pangolier\",\n\"id\":120\n},\n{\n\"name\":\"npc_dota_hero_dark_willow\",\n\"id\":119\n},\n{\n\"name\":\"npc_dota_hero_grimstroke\",\n\"id\":121\n},\n{\n\"name\":\"npc_dota_hero_mars\",\n\"id\":129\n},\n{\n\"name\":\"npc_dota_hero_void_spirit\",\n\"id\":126\n},\n{\n\"name\":\"npc_dota_hero_snapfire\",\n\"id\":128\n}\n]\n,\n\"status\":200,\n\"count\":119\n}\n}"
)

func TestGetHeroes(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	g.Describe("GetHeroes", func() {
		var heroes Heroes
		var err error
		g.It("Should call the correct URL", func() {
			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				g.Assert(req.URL.String()).Equal(getHeroesUrl(&api) + "?key=keyTEST")
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
			}
			heroes, err = api.GetHeroes()
		})
		g.It("Should return no error", func() {
			g.Assert(err).Equal(nil)
		})
		g.It("Should return 119 heroes", func() {
			g.Assert(heroes.Count()).Equal(119)
		})
		g.It("Should work with concurrent usage", func() {
			var calledMutex sync.Mutex
			var called bool
			var wg sync.WaitGroup
			api.heroesCache.fromCache = 0

			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				g.Assert(req.URL.String()).Equal(getHeroesUrl(&api) + "?key=keyTEST")
				calledMutex.Lock()
				defer calledMutex.Unlock()
				g.Assert(called).IsFalse()
				called = true
				return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
			}

			wg.Add(10)
			for i := 0; i < 10; i++ {
				go func() {
					h, err := api.GetHeroes()
					g.Assert(err).IsNil()
					g.Assert(h).Equal(heroes)
					defer wg.Done()
				}()
			}
			wg.Wait()
			g.Assert(called).IsTrue()
		})
		g.It("Should fill cache", func() {
			g.Assert(api.heroesCache.fromCache).Equal(uint32(1))
		})
	})
}

func TestHeroName(t *testing.T) {
	g := Goblin(t)
	g.Describe("Heroes Names", func() {
		g.It("Should return the correct names", func() {
			for _, name := range []string{"name0", "a", "anotherName"} {
				hN := heroNameFromFullName(heroPrefix + name)
				g.Assert(hN.GetName()).Equal(name)
				g.Assert(hN.GetFullName()).Equal(heroPrefix + name)
				g.Assert(hN.GetPrefix()).Equal(heroPrefix)
			}
		})
	})
}

func TestDota2_GetHeroImage(t *testing.T) {
	g := Goblin(t)
	var wg sync.WaitGroup
	api, _ := LoadConfigFromFile("config.yaml")
	heroes, _ := api.GetHeroes()
	doTest := func(size int) {
		for i := 0; i < heroes.Count(); i += heroes.Count() / 10 {
			if i > heroes.Count() {
				continue
			}
			wg.Add(1)
			i := i
			go func() {
				img, err := api.GetHeroImage(heroes.heroes[i], size)
				g.Assert(err).Equal(nil)
				g.Assert(img == nil).IsFalse()
				g.Assert(img.Bounds().Dx() > 0).IsTrue()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	g.Describe("api.GetHeroImage", func() {
		g.It("Should return lg Images", func(done Done) {
			doTest(SizeLg)
			done()
		})
		g.It("Should return sb Images", func(done Done) {
			doTest(SizeSb)
			done()
		})
		g.It("Should return full Images", func(done Done) {
			doTest(SizeFull)
			done()
		})
		g.It("Should return vert Images", func(done Done) {
			doTest(SizeVert)
			done()
		})
	})
}

func TestHeroes_GetById(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when ID is found", func() {
			for _, hero := range heroes.heroes {
				h, found := heroes.GetById(hero.ID)
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when ID is not found", func() {
			missingId := 1
			for _, hero := range heroes.heroes {
				missingId += hero.ID
			}
			_, found := heroes.GetById(missingId)
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkHeroes_GetById(b *testing.B) {
	api, _ := LoadConfigFromFile("config.yaml")
	heroes, _ := api.GetHeroes()
	ids := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		ids[i] = heroes.heroes[rand.Int()%len(heroes.heroes)].ID
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		heroes.GetById(ids[i])
	}
}

func TestHeroes_GetByName(t *testing.T) {
	g := Goblin(t)
	api, _ := LoadConfigFromFile("config.yaml")
	heroes, _ := api.GetHeroes()
	g.Describe("Items.GetById", func() {
		g.It("Should return the correct hero when Name is found", func() {
			for _, hero := range heroes.heroes {
				h, found := heroes.GetByName(hero.Name.GetFullName())
				g.Assert(found).IsTrue()
				g.Assert(h.ID).Equal(hero.ID)
				g.Assert(h.Name).Equal(hero.Name)
			}
		})
		g.It("Should return Error when Name is not found", func() {
			_, found := heroes.GetByName("irNe9GNzJm")
			g.Assert(found).IsFalse()
		})
	})
}

func BenchmarkHeroes_GetByName(b *testing.B) {
	api, _ := LoadConfigFromFile("config.yaml")
	heroes, _ := api.GetHeroes()
	names := make([]string, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		names[i] = heroes.heroes[rand.Int()%len(heroes.heroes)].Name.GetFullName()
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		heroes.GetByName(names[i])
	}
}
