package dota2api

import (
	"fmt"
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	response = "{\n\"result\":{\n\"players\":[\n{\n\"account_id\":62065,\n\"player_slot\":0,\n\"hero_id\":75,\n\"item_0\":63,\n\"item_1\":77,\n\"item_2\":236,\n\"item_3\":77,\n\"item_4\":485,\n\"item_5\":7,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":357,\n\"kills\":3,\n\"deaths\":10,\n\"assists\":9,\n\"leaver_status\":0,\n\"last_hits\":119,\n\"denies\":8,\n\"gold_per_min\":281,\n\"xp_per_min\":490,\n\"level\":21,\n\"hero_damage\":23281,\n\"tower_damage\":270,\n\"hero_healing\":0,\n\"gold\":1250,\n\"gold_spent\":9490,\n\"scaled_hero_damage\":15776,\n\"scaled_tower_damage\":79,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5378,\n\"time\":150,\n\"level\":1\n},\n{\n\"ability\":5377,\n\"time\":315,\n\"level\":2\n},\n{\n\"ability\":5378,\n\"time\":431,\n\"level\":3\n},\n{\n\"ability\":5377,\n\"time\":579,\n\"level\":4\n},\n{\n\"ability\":5378,\n\"time\":714,\n\"level\":5\n},\n{\n\"ability\":5377,\n\"time\":753,\n\"level\":6\n},\n{\n\"ability\":5380,\n\"time\":915,\n\"level\":7\n},\n{\n\"ability\":5378,\n\"time\":1100,\n\"level\":8\n},\n{\n\"ability\":5377,\n\"time\":1269,\n\"level\":9\n},\n{\n\"ability\":5906,\n\"time\":1343,\n\"level\":10\n},\n{\n\"ability\":5379,\n\"time\":1458,\n\"level\":11\n},\n{\n\"ability\":5380,\n\"time\":1602,\n\"level\":12\n},\n{\n\"ability\":5379,\n\"time\":1699,\n\"level\":13\n},\n{\n\"ability\":5379,\n\"time\":1802,\n\"level\":14\n},\n{\n\"ability\":6878,\n\"time\":1966,\n\"level\":15\n},\n{\n\"ability\":5379,\n\"time\":2061,\n\"level\":16\n},\n{\n\"ability\":5380,\n\"time\":2238,\n\"level\":17\n},\n{\n\"ability\":5943,\n\"time\":2442,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":4294967295,\n\"player_slot\":1,\n\"hero_id\":26,\n\"item_0\":254,\n\"item_1\":1,\n\"item_2\":0,\n\"item_3\":43,\n\"item_4\":214,\n\"item_5\":37,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":336,\n\"kills\":4,\n\"deaths\":4,\n\"assists\":5,\n\"leaver_status\":0,\n\"last_hits\":21,\n\"denies\":3,\n\"gold_per_min\":197,\n\"xp_per_min\":323,\n\"level\":17,\n\"hero_damage\":6989,\n\"tower_damage\":0,\n\"hero_healing\":0,\n\"gold\":331,\n\"gold_spent\":8200,\n\"scaled_hero_damage\":4937,\n\"scaled_tower_damage\":0,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5044,\n\"time\":222,\n\"level\":1\n},\n{\n\"ability\":5046,\n\"time\":373,\n\"level\":2\n},\n{\n\"ability\":5045,\n\"time\":475,\n\"level\":3\n},\n{\n\"ability\":5044,\n\"time\":553,\n\"level\":4\n},\n{\n\"ability\":5046,\n\"time\":714,\n\"level\":5\n},\n{\n\"ability\":5047,\n\"time\":838,\n\"level\":6\n},\n{\n\"ability\":5044,\n\"time\":879,\n\"level\":7\n},\n{\n\"ability\":5044,\n\"time\":942,\n\"level\":8\n},\n{\n\"ability\":5045,\n\"time\":1183,\n\"level\":9\n},\n{\n\"ability\":5046,\n\"time\":1325,\n\"level\":10\n},\n{\n\"ability\":465,\n\"time\":1535,\n\"level\":11\n},\n{\n\"ability\":5047,\n\"time\":1690,\n\"level\":12\n},\n{\n\"ability\":5046,\n\"time\":1874,\n\"level\":13\n},\n{\n\"ability\":5045,\n\"time\":1990,\n\"level\":14\n},\n{\n\"ability\":6056,\n\"time\":2178,\n\"level\":15\n},\n{\n\"ability\":5045,\n\"time\":2355,\n\"level\":16\n}\n]\n\n},\n{\n\"account_id\":250761061,\n\"player_slot\":2,\n\"hero_id\":11,\n\"item_0\":154,\n\"item_1\":75,\n\"item_2\":1,\n\"item_3\":168,\n\"item_4\":149,\n\"item_5\":63,\n\"backpack_0\":41,\n\"backpack_1\":75,\n\"backpack_2\":0,\n\"item_neutral\":381,\n\"kills\":5,\n\"deaths\":9,\n\"assists\":7,\n\"leaver_status\":0,\n\"last_hits\":290,\n\"denies\":6,\n\"gold_per_min\":464,\n\"xp_per_min\":790,\n\"level\":25,\n\"hero_damage\":17710,\n\"tower_damage\":1073,\n\"hero_healing\":0,\n\"gold\":803,\n\"gold_spent\":19175,\n\"scaled_hero_damage\":9895,\n\"scaled_tower_damage\":746,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5062,\n\"time\":221,\n\"level\":1\n},\n{\n\"ability\":5059,\n\"time\":286,\n\"level\":2\n},\n{\n\"ability\":5061,\n\"time\":338,\n\"level\":3\n},\n{\n\"ability\":5062,\n\"time\":395,\n\"level\":4\n},\n{\n\"ability\":5060,\n\"time\":513,\n\"level\":5\n},\n{\n\"ability\":5062,\n\"time\":581,\n\"level\":6\n},\n{\n\"ability\":5060,\n\"time\":683,\n\"level\":7\n},\n{\n\"ability\":5062,\n\"time\":736,\n\"level\":8\n},\n{\n\"ability\":5064,\n\"time\":805,\n\"level\":9\n},\n{\n\"ability\":6119,\n\"time\":874,\n\"level\":10\n},\n{\n\"ability\":5063,\n\"time\":948,\n\"level\":11\n},\n{\n\"ability\":5064,\n\"time\":1044,\n\"level\":12\n},\n{\n\"ability\":5063,\n\"time\":1143,\n\"level\":13\n},\n{\n\"ability\":5063,\n\"time\":1283,\n\"level\":14\n},\n{\n\"ability\":5919,\n\"time\":1391,\n\"level\":15\n},\n{\n\"ability\":5063,\n\"time\":1490,\n\"level\":16\n},\n{\n\"ability\":5064,\n\"time\":1547,\n\"level\":17\n},\n{\n\"ability\":6670,\n\"time\":1884,\n\"level\":18\n},\n{\n\"ability\":6912,\n\"time\":2421,\n\"level\":19\n}\n]\n\n},\n{\n\"account_id\":4294967295,\n\"player_slot\":3,\n\"hero_id\":42,\n\"item_0\":50,\n\"item_1\":127,\n\"item_2\":116,\n\"item_3\":252,\n\"item_4\":13,\n\"item_5\":137,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":363,\n\"kills\":3,\n\"deaths\":2,\n\"assists\":4,\n\"leaver_status\":0,\n\"last_hits\":301,\n\"denies\":12,\n\"gold_per_min\":433,\n\"xp_per_min\":591,\n\"level\":23,\n\"hero_damage\":24186,\n\"tower_damage\":764,\n\"hero_healing\":890,\n\"gold\":2499,\n\"gold_spent\":15920,\n\"scaled_hero_damage\":9896,\n\"scaled_tower_damage\":175,\n\"scaled_hero_healing\":323,\n\"ability_upgrades\":[\n{\n\"ability\":5086,\n\"time\":190,\n\"level\":1\n},\n{\n\"ability\":5087,\n\"time\":331,\n\"level\":2\n},\n{\n\"ability\":5087,\n\"time\":394,\n\"level\":3\n},\n{\n\"ability\":5088,\n\"time\":501,\n\"level\":4\n},\n{\n\"ability\":5087,\n\"time\":629,\n\"level\":5\n},\n{\n\"ability\":5089,\n\"time\":766,\n\"level\":6\n},\n{\n\"ability\":5087,\n\"time\":883,\n\"level\":7\n},\n{\n\"ability\":5088,\n\"time\":985,\n\"level\":8\n},\n{\n\"ability\":5088,\n\"time\":1127,\n\"level\":9\n},\n{\n\"ability\":6119,\n\"time\":1231,\n\"level\":10\n},\n{\n\"ability\":5088,\n\"time\":1283,\n\"level\":11\n},\n{\n\"ability\":5089,\n\"time\":1406,\n\"level\":12\n},\n{\n\"ability\":5086,\n\"time\":1489,\n\"level\":13\n},\n{\n\"ability\":5086,\n\"time\":1630,\n\"level\":14\n},\n{\n\"ability\":5928,\n\"time\":1719,\n\"level\":15\n},\n{\n\"ability\":5086,\n\"time\":1817,\n\"level\":16\n},\n{\n\"ability\":5089,\n\"time\":2119,\n\"level\":17\n},\n{\n\"ability\":6201,\n\"time\":2254,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":4294967295,\n\"player_slot\":4,\n\"hero_id\":99,\n\"item_0\":8,\n\"item_1\":21,\n\"item_2\":50,\n\"item_3\":36,\n\"item_4\":90,\n\"item_5\":242,\n\"backpack_0\":178,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":311,\n\"kills\":2,\n\"deaths\":11,\n\"assists\":11,\n\"leaver_status\":0,\n\"last_hits\":216,\n\"denies\":8,\n\"gold_per_min\":367,\n\"xp_per_min\":521,\n\"level\":22,\n\"hero_damage\":21095,\n\"tower_damage\":1279,\n\"hero_healing\":0,\n\"gold\":943,\n\"gold_spent\":14560,\n\"scaled_hero_damage\":15130,\n\"scaled_tower_damage\":745,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5549,\n\"time\":153,\n\"level\":1\n},\n{\n\"ability\":5550,\n\"time\":363,\n\"level\":2\n},\n{\n\"ability\":5549,\n\"time\":411,\n\"level\":3\n},\n{\n\"ability\":5548,\n\"time\":522,\n\"level\":4\n},\n{\n\"ability\":5549,\n\"time\":689,\n\"level\":5\n},\n{\n\"ability\":5551,\n\"time\":838,\n\"level\":6\n},\n{\n\"ability\":5549,\n\"time\":905,\n\"level\":7\n},\n{\n\"ability\":5550,\n\"time\":1069,\n\"level\":8\n},\n{\n\"ability\":5550,\n\"time\":1111,\n\"level\":9\n},\n{\n\"ability\":5917,\n\"time\":1207,\n\"level\":10\n},\n{\n\"ability\":5550,\n\"time\":1290,\n\"level\":11\n},\n{\n\"ability\":5551,\n\"time\":1377,\n\"level\":12\n},\n{\n\"ability\":5548,\n\"time\":1458,\n\"level\":13\n},\n{\n\"ability\":5548,\n\"time\":1644,\n\"level\":14\n},\n{\n\"ability\":5959,\n\"time\":1757,\n\"level\":15\n},\n{\n\"ability\":5548,\n\"time\":1868,\n\"level\":16\n},\n{\n\"ability\":5551,\n\"time\":2208,\n\"level\":17\n},\n{\n\"ability\":6360,\n\"time\":2262,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":55807544,\n\"player_slot\":128,\n\"hero_id\":44,\n\"item_0\":50,\n\"item_1\":116,\n\"item_2\":247,\n\"item_3\":208,\n\"item_4\":135,\n\"item_5\":168,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":376,\n\"kills\":15,\n\"deaths\":5,\n\"assists\":9,\n\"leaver_status\":0,\n\"last_hits\":322,\n\"denies\":31,\n\"gold_per_min\":633,\n\"xp_per_min\":905,\n\"level\":29,\n\"hero_damage\":41905,\n\"tower_damage\":2222,\n\"hero_healing\":0,\n\"gold\":1121,\n\"gold_spent\":26175,\n\"scaled_hero_damage\":20272,\n\"scaled_tower_damage\":1136,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5190,\n\"time\":159,\n\"level\":1\n},\n{\n\"ability\":5191,\n\"time\":329,\n\"level\":2\n},\n{\n\"ability\":5190,\n\"time\":453,\n\"level\":3\n},\n{\n\"ability\":5192,\n\"time\":549,\n\"level\":4\n},\n{\n\"ability\":5190,\n\"time\":623,\n\"level\":5\n},\n{\n\"ability\":5193,\n\"time\":728,\n\"level\":6\n},\n{\n\"ability\":5190,\n\"time\":862,\n\"level\":7\n},\n{\n\"ability\":5191,\n\"time\":933,\n\"level\":8\n},\n{\n\"ability\":5191,\n\"time\":1067,\n\"level\":9\n},\n{\n\"ability\":6034,\n\"time\":1182,\n\"level\":10\n},\n{\n\"ability\":5191,\n\"time\":1280,\n\"level\":11\n},\n{\n\"ability\":5193,\n\"time\":1426,\n\"level\":12\n},\n{\n\"ability\":5192,\n\"time\":1490,\n\"level\":13\n},\n{\n\"ability\":5192,\n\"time\":1614,\n\"level\":14\n},\n{\n\"ability\":433,\n\"time\":1682,\n\"level\":15\n},\n{\n\"ability\":5192,\n\"time\":1759,\n\"level\":16\n},\n{\n\"ability\":5193,\n\"time\":1855,\n\"level\":17\n},\n{\n\"ability\":7383,\n\"time\":1948,\n\"level\":18\n},\n{\n\"ability\":483,\n\"time\":2134,\n\"level\":19\n}\n]\n\n},\n{\n\"account_id\":216085151,\n\"player_slot\":129,\n\"hero_id\":7,\n\"item_0\":102,\n\"item_1\":1,\n\"item_2\":180,\n\"item_3\":34,\n\"item_4\":38,\n\"item_5\":69,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":287,\n\"kills\":1,\n\"deaths\":4,\n\"assists\":15,\n\"leaver_status\":0,\n\"last_hits\":67,\n\"denies\":2,\n\"gold_per_min\":274,\n\"xp_per_min\":482,\n\"level\":21,\n\"hero_damage\":8742,\n\"tower_damage\":1099,\n\"hero_healing\":400,\n\"gold\":2591,\n\"gold_spent\":9270,\n\"scaled_hero_damage\":4212,\n\"scaled_tower_damage\":540,\n\"scaled_hero_healing\":204,\n\"ability_upgrades\":[\n{\n\"ability\":5023,\n\"time\":207,\n\"level\":1\n},\n{\n\"ability\":5024,\n\"time\":346,\n\"level\":2\n},\n{\n\"ability\":5025,\n\"time\":541,\n\"level\":3\n},\n{\n\"ability\":5024,\n\"time\":641,\n\"level\":4\n},\n{\n\"ability\":5025,\n\"time\":809,\n\"level\":5\n},\n{\n\"ability\":5026,\n\"time\":919,\n\"level\":6\n},\n{\n\"ability\":5024,\n\"time\":1140,\n\"level\":7\n},\n{\n\"ability\":5025,\n\"time\":1234,\n\"level\":8\n},\n{\n\"ability\":5025,\n\"time\":1309,\n\"level\":9\n},\n{\n\"ability\":496,\n\"time\":1441,\n\"level\":10\n},\n{\n\"ability\":5024,\n\"time\":1497,\n\"level\":11\n},\n{\n\"ability\":5026,\n\"time\":1634,\n\"level\":12\n},\n{\n\"ability\":5023,\n\"time\":1830,\n\"level\":13\n},\n{\n\"ability\":5023,\n\"time\":2039,\n\"level\":14\n},\n{\n\"ability\":5919,\n\"time\":2083,\n\"level\":15\n},\n{\n\"ability\":5023,\n\"time\":2214,\n\"level\":16\n},\n{\n\"ability\":5026,\n\"time\":2356,\n\"level\":17\n},\n{\n\"ability\":6425,\n\"time\":2689,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":107353159,\n\"player_slot\":130,\n\"hero_id\":47,\n\"item_0\":206,\n\"item_1\":158,\n\"item_2\":263,\n\"item_3\":116,\n\"item_4\":63,\n\"item_5\":117,\n\"backpack_0\":0,\n\"backpack_1\":0,\n\"backpack_2\":75,\n\"item_neutral\":379,\n\"kills\":10,\n\"deaths\":2,\n\"assists\":11,\n\"leaver_status\":0,\n\"last_hits\":334,\n\"denies\":16,\n\"gold_per_min\":548,\n\"xp_per_min\":641,\n\"level\":24,\n\"hero_damage\":31367,\n\"tower_damage\":10991,\n\"hero_healing\":0,\n\"gold\":3675,\n\"gold_spent\":19385,\n\"scaled_hero_damage\":17146,\n\"scaled_tower_damage\":5286,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5218,\n\"time\":155,\n\"level\":1\n},\n{\n\"ability\":5220,\n\"time\":319,\n\"level\":2\n},\n{\n\"ability\":5218,\n\"time\":330,\n\"level\":3\n},\n{\n\"ability\":5220,\n\"time\":420,\n\"level\":4\n},\n{\n\"ability\":5218,\n\"time\":479,\n\"level\":5\n},\n{\n\"ability\":5221,\n\"time\":564,\n\"level\":6\n},\n{\n\"ability\":5218,\n\"time\":642,\n\"level\":7\n},\n{\n\"ability\":5220,\n\"time\":714,\n\"level\":8\n},\n{\n\"ability\":5220,\n\"time\":836,\n\"level\":9\n},\n{\n\"ability\":5906,\n\"time\":941,\n\"level\":10\n},\n{\n\"ability\":5219,\n\"time\":1001,\n\"level\":11\n},\n{\n\"ability\":5221,\n\"time\":1167,\n\"level\":12\n},\n{\n\"ability\":5219,\n\"time\":1265,\n\"level\":13\n},\n{\n\"ability\":5219,\n\"time\":1308,\n\"level\":14\n},\n{\n\"ability\":5943,\n\"time\":1439,\n\"level\":15\n},\n{\n\"ability\":5219,\n\"time\":1528,\n\"level\":16\n},\n{\n\"ability\":5221,\n\"time\":1798,\n\"level\":17\n},\n{\n\"ability\":5942,\n\"time\":2017,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":4294967295,\n\"player_slot\":131,\n\"hero_id\":50,\n\"item_0\":231,\n\"item_1\":102,\n\"item_2\":254,\n\"item_3\":0,\n\"item_4\":0,\n\"item_5\":0,\n\"backpack_0\":16,\n\"backpack_1\":0,\n\"backpack_2\":0,\n\"item_neutral\":375,\n\"kills\":4,\n\"deaths\":5,\n\"assists\":18,\n\"leaver_status\":0,\n\"last_hits\":91,\n\"denies\":5,\n\"gold_per_min\":315,\n\"xp_per_min\":540,\n\"level\":22,\n\"hero_damage\":8017,\n\"tower_damage\":61,\n\"hero_healing\":17460,\n\"gold\":1303,\n\"gold_spent\":12470,\n\"scaled_hero_damage\":4903,\n\"scaled_tower_damage\":49,\n\"scaled_hero_healing\":7527,\n\"ability_upgrades\":[\n{\n\"ability\":5233,\n\"time\":202,\n\"level\":1\n},\n{\n\"ability\":5235,\n\"time\":329,\n\"level\":2\n},\n{\n\"ability\":5233,\n\"time\":455,\n\"level\":3\n},\n{\n\"ability\":5234,\n\"time\":639,\n\"level\":4\n},\n{\n\"ability\":5233,\n\"time\":774,\n\"level\":5\n},\n{\n\"ability\":7304,\n\"time\":872,\n\"level\":6\n},\n{\n\"ability\":5233,\n\"time\":997,\n\"level\":7\n},\n{\n\"ability\":5235,\n\"time\":1193,\n\"level\":8\n},\n{\n\"ability\":5235,\n\"time\":1308,\n\"level\":9\n},\n{\n\"ability\":5941,\n\"time\":1404,\n\"level\":10\n},\n{\n\"ability\":5235,\n\"time\":1438,\n\"level\":11\n},\n{\n\"ability\":7304,\n\"time\":1666,\n\"level\":12\n},\n{\n\"ability\":5234,\n\"time\":1824,\n\"level\":13\n},\n{\n\"ability\":5234,\n\"time\":1971,\n\"level\":14\n},\n{\n\"ability\":6056,\n\"time\":2049,\n\"level\":15\n},\n{\n\"ability\":5234,\n\"time\":2083,\n\"level\":16\n},\n{\n\"ability\":7304,\n\"time\":2419,\n\"level\":17\n},\n{\n\"ability\":6528,\n\"time\":2532,\n\"level\":18\n}\n]\n\n},\n{\n\"account_id\":4294967295,\n\"player_slot\":132,\n\"hero_id\":35,\n\"item_0\":75,\n\"item_1\":259,\n\"item_2\":29,\n\"item_3\":152,\n\"item_4\":108,\n\"item_5\":139,\n\"backpack_0\":257,\n\"backpack_1\":75,\n\"backpack_2\":0,\n\"item_neutral\":311,\n\"kills\":5,\n\"deaths\":3,\n\"assists\":10,\n\"leaver_status\":0,\n\"last_hits\":224,\n\"denies\":28,\n\"gold_per_min\":464,\n\"xp_per_min\":561,\n\"level\":23,\n\"hero_damage\":29209,\n\"tower_damage\":3343,\n\"hero_healing\":0,\n\"gold\":2868,\n\"gold_spent\":18085,\n\"scaled_hero_damage\":12548,\n\"scaled_tower_damage\":1464,\n\"scaled_hero_healing\":0,\n\"ability_upgrades\":[\n{\n\"ability\":5155,\n\"time\":154,\n\"level\":1\n},\n{\n\"ability\":5154,\n\"time\":352,\n\"level\":2\n},\n{\n\"ability\":5156,\n\"time\":415,\n\"level\":3\n},\n{\n\"ability\":5155,\n\"time\":477,\n\"level\":4\n},\n{\n\"ability\":5155,\n\"time\":618,\n\"level\":5\n},\n{\n\"ability\":5157,\n\"time\":687,\n\"level\":6\n},\n{\n\"ability\":5155,\n\"time\":824,\n\"level\":7\n},\n{\n\"ability\":5154,\n\"time\":875,\n\"level\":8\n},\n{\n\"ability\":5156,\n\"time\":1000,\n\"level\":9\n},\n{\n\"ability\":5951,\n\"time\":1114,\n\"level\":10\n},\n{\n\"ability\":5154,\n\"time\":1220,\n\"level\":11\n},\n{\n\"ability\":5157,\n\"time\":1399,\n\"level\":12\n},\n{\n\"ability\":5156,\n\"time\":1472,\n\"level\":13\n},\n{\n\"ability\":5156,\n\"time\":1590,\n\"level\":14\n},\n{\n\"ability\":5907,\n\"time\":1827,\n\"level\":15\n},\n{\n\"ability\":5154,\n\"time\":1943,\n\"level\":16\n},\n{\n\"ability\":5157,\n\"time\":2093,\n\"level\":17\n},\n{\n\"ability\":6305,\n\"time\":2296,\n\"level\":18\n}\n]\n\n}\n]\n,\n\"radiant_win\":false,\n\"duration\":2552,\n\"pre_game_duration\":90,\n\"start_time\":1596363511,\n\"match_id\":5548608983,\n\"match_seq_num\":4655597189,\n\"tower_status_radiant\":262,\n\"tower_status_dire\":1974,\n\"barracks_status_radiant\":51,\n\"barracks_status_dire\":63,\n\"cluster\":136,\n\"first_blood_time\":18,\n\"lobby_type\":7,\n\"human_players\":10,\n\"leagueid\":0,\n\"positive_votes\":0,\n\"negative_votes\":0,\n\"game_mode\":22,\n\"flags\":1,\n\"engine\":1,\n\"radiant_score\":19,\n\"dire_score\":36,\n\"picks_bans\":[\n{\n\"is_pick\":true,\n\"hero_id\":35,\n\"team\":1,\n\"order\":0\n},\n{\n\"is_pick\":true,\n\"hero_id\":26,\n\"team\":0,\n\"order\":1\n},\n{\n\"is_pick\":true,\n\"hero_id\":50,\n\"team\":1,\n\"order\":2\n},\n{\n\"is_pick\":true,\n\"hero_id\":99,\n\"team\":0,\n\"order\":3\n},\n{\n\"is_pick\":true,\n\"hero_id\":75,\n\"team\":0,\n\"order\":4\n},\n{\n\"is_pick\":true,\n\"hero_id\":7,\n\"team\":1,\n\"order\":5\n},\n{\n\"is_pick\":true,\n\"hero_id\":47,\n\"team\":1,\n\"order\":6\n},\n{\n\"is_pick\":true,\n\"hero_id\":42,\n\"team\":0,\n\"order\":7\n},\n{\n\"is_pick\":true,\n\"hero_id\":11,\n\"team\":0,\n\"order\":8\n},\n{\n\"is_pick\":true,\n\"hero_id\":44,\n\"team\":1,\n\"order\":9\n},\n{\n\"is_pick\":false,\n\"hero_id\":46,\n\"team\":0,\n\"order\":10\n},\n{\n\"is_pick\":false,\n\"hero_id\":105,\n\"team\":0,\n\"order\":11\n},\n{\n\"is_pick\":false,\n\"hero_id\":104,\n\"team\":0,\n\"order\":12\n},\n{\n\"is_pick\":false,\n\"hero_id\":41,\n\"team\":0,\n\"order\":13\n},\n{\n\"is_pick\":false,\n\"hero_id\":23,\n\"team\":0,\n\"order\":14\n},\n{\n\"is_pick\":false,\n\"hero_id\":82,\n\"team\":1,\n\"order\":15\n},\n{\n\"is_pick\":false,\n\"hero_id\":32,\n\"team\":1,\n\"order\":16\n}\n]\n\n}\n}"
)

func TestDota2_GetMatchDetails(t *testing.T) {
	g := Goblin(t)
	mockClient := mockClient{}
	api := LoadConfig(GetTestConfig())
	api.client = &mockClient
	var details MatchDetails
	var err error
	g.Describe("api.GetMatchDetails", func() {
		g.It("Should call the correct URL", func() {
			mockClient.DoFunc = func(req *http.Request) (*http.Response, error) {
				switch req.URL.String() {
				case api.getMatchDetailsUrl() + "?key=keyTEST&match_id=42":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(response))}, nil
				case api.getHeroesUrl() + "?key=keyTEST":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(heroesResponse))}, nil
				case api.getItemsUrl() + "?key=keyTEST":
					return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(itemsResponse))}, nil
				default:
					g.Fail("Unnecessary API call")
				}
				return nil, nil
			}
			details, err = api.GetMatchDetails(MatchId(42))
		})
		g.It("Should not error", func() {
			g.Assert(err).IsNil()
		})
		g.It("Should return the correct match ID and SeqNum", func() {
			g.Assert(details.MatchID).Equal(int64(5548608983))
			g.Assert(details.MatchSeqNum == 4655597189).IsTrue()
		})
		g.It("Should return a HumanPlayers, 5 players in each team", func() {
			g.Assert(details.HumanPlayers).Equal(10)
			g.Assert(details.Radiant.Count()).Equal(5)
			g.Assert(details.Dire.Count()).Equal(5)
			details.ForEachPlayer(func(p PlayerDetails) {
				g.Assert(p.LeaverStatus).Equal(LeaverStatusNone)
			})
		})
		g.It("Should return Source2 as engine", func() {
			g.Assert(details.Engine).Equal(Source2)
		})
		g.Xit("Should return Ranked Matchmaking as GameMode", func() {
			g.Assert(details.GameMode).Equal(GameModeRankedMatchmaking)
			g.Assert(details.GameMode.GetString()).Equal("Ranked Matchmaking")
		})
		g.Xit("Should return ranked Matchmaking as LobbyType", func() {
			g.Assert(details.LobbyType).Equal(LobbyRankedMatchmaking)
			g.Assert(details.LobbyType.GetName()).Equal("Ranked Matchmaking")
		})
		g.It("Should return 19-36 as score", func() {
			g.Assert(details.Score.RadiantScore).Equal(19)
			g.Assert(details.Score.DireScore).Equal(36)
		})
		g.It("Should return Dire as winner", func() {
			g.Assert(details.Victory.DireWon()).IsTrue()
			g.Assert(details.Victory.RadiantWon()).IsFalse()
			g.Assert(details.Victory.GetWinningTeam()).Equal(Dire)
		})
		g.It("Should return correct time stamps", func() {
			g.Assert(details.Duration).Equal(2552 * time.Second)
			g.Assert(details.FirstBloodTime).Equal(18 * time.Second)
			g.Assert(details.StartTime.Equal(time.Unix(1596363511, 0))).IsTrue()
			g.Assert(details.PreGameDuration).Equal(90 * time.Second)
		})
		g.It("Should return correct BuildingState", func() {
			//Radiant
			g.Assert(details.BuildingsState.Radiant.Top.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Top.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Top.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Mid.RangedBarrack).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.MeleeBarrack).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T2Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Mid.T3Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Bot.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Radiant.Bot.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.T2Tower).IsFalse()
			g.Assert(details.BuildingsState.Radiant.Bot.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Radiant.T4TowerBot).IsFalse()
			g.Assert(details.BuildingsState.Radiant.T4TowerTop).IsFalse()
			//Dire
			g.Assert(details.BuildingsState.Dire.Top.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Top.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Top.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Mid.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Mid.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.RangedBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.MeleeBarrack).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.T1Tower).IsFalse()
			g.Assert(details.BuildingsState.Dire.Bot.T2Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.Bot.T3Tower).IsTrue()
			g.Assert(details.BuildingsState.Dire.T4TowerBot).IsTrue()
			g.Assert(details.BuildingsState.Dire.T4TowerTop).IsTrue()
		})
		g.It("Should return correct time stamps", func() {
			g.Assert(details.Duration).Equal(2552 * time.Second)
			g.Assert(details.FirstBloodTime).Equal(18 * time.Second)
			g.Assert(details.StartTime.Equal(time.Unix(1596363511, 0))).IsTrue()
			g.Assert(details.PreGameDuration).Equal(90 * time.Second)
		})
		g.It("Should return correct Picks and Bans", func() {
			bansHeroesIds := []int{46, 105, 104, 41, 23, 82, 32}
			bansTeams := []Victory{RadiantVictory, RadiantVictory, RadiantVictory, RadiantVictory, RadiantVictory, DireVictory, DireVictory}
			p, f := details.PicksBans.GetPick(0)
			g.Assert(f).IsTrue()
			g.Assert(p.IsPick()).IsTrue()
			p, f = details.PicksBans.GetPick(7)
			g.Assert(f).IsTrue()
			g.Assert(p.IsPick()).IsTrue()
			g.Assert(p.IsRadiant()).IsTrue()
			_, f = details.PicksBans.GetPick(17)
			g.Assert(f).IsFalse()
			for i, ban := range details.PicksBans.GetByPickType(Ban) {
				g.Assert(ban.IsBan()).IsTrue()
				g.Assert(ban.Hero.ID).Equal(bansHeroesIds[i])
				g.Assert(ban.Order).Equal(10 + i)
				g.Assert(ban.GetTeam() == int(bansTeams[i])).IsTrue()
			}
			picksHeroesIds := []int{35, 26, 50, 99, 75, 7, 47, 42, 11, 44}
			picksIsDire := []bool{true, false, true, false, false, true, true, false, false, true}
			for i, pick := range details.PicksBans.GetByPickType(Pick) {
				g.Assert(pick.GetType()).Equal(Pick)
				g.Assert(pick.Hero.ID).Equal(picksHeroesIds[i])
				g.Assert(pick.Order).Equal(i)
				g.Assert(pick.IsDire()).Equal(picksIsDire[i])
			}
			for _, pick := range details.PicksBans {
				p, f := details.PicksBans.GetPickByHero(pick.Hero)
				g.Assert(f).IsTrue()
				g.Assert(p.Hero).Equal(pick.Hero)
			}
			for _, pick := range details.PicksBans.GetByTeam(Dire) {
				g.Assert(pick.IsDire()).IsTrue()
			}
		})
		g.It("Should return the correct stats for player 0", func() {
			g.Assert(details.Radiant[0].Stats.Gold.Spent() == 9490).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.Spent().Raw() == 9490).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.Spent().ToString()).Equal("9.5k")
			g.Assert(details.Radiant[0].Stats.Gold.Current() == 1250).IsTrue()
			g.Assert(details.Radiant[0].Stats.Gold.NetWorth() == 10740).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroDamage.Raw() == 23281).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroDamage.Scaled() == 15776).IsTrue()
			g.Assert(details.Radiant[0].Stats.TowerDamage.Raw() == 270).IsTrue()
			g.Assert(details.Radiant[0].Stats.TowerDamage.Scaled() == 79).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroHealing.Raw() == 0).IsTrue()
			g.Assert(details.Radiant[0].Stats.HeroHealing.Scaled() == 0).IsTrue()

			g.Assert(details.Radiant[0].Stats.Gold.Current().ToString()).Equal("1.3k")
			g.Assert(details.Radiant[1].Stats.Gold.Current().ToString()).Equal("331")
		})
		g.It("Should return working stats", func() {
			details.ForEachPlayer(func(p PlayerDetails) {
				g.Assert(p.Stats.Gold.NetWorth()).Equal(p.Stats.Gold.Current() + p.Stats.Gold.Spent())
				g.Assert(p.Stats.HeroDamage.ScalingFactor()).Equal(float64(p.Stats.HeroDamage.Scaled()) / float64(p.Stats.HeroDamage.Raw()))
				g.Assert(p.Stats.TowerDamage.ScalingFactor()).Equal(float64(p.Stats.TowerDamage.Scaled()) / float64(p.Stats.TowerDamage.Raw()))
				g.Assert(p.Stats.HeroHealing.ScalingFactor()).Equal(float64(p.Stats.HeroHealing.Scaled()) / float64(p.Stats.HeroHealing.Raw()))
			})
		})
		g.It("Should return the correct items for player 0", func() {
			g.Assert(details.Radiant[0].Items.Item0.ID == 63).IsTrue()
			g.Assert(details.Radiant[0].Items.Item1.ID == 77).IsTrue()
			g.Assert(details.Radiant[0].Items.Item2.ID == 236).IsTrue()
			g.Assert(details.Radiant[0].Items.Item3.ID == 77).IsTrue()
			g.Assert(details.Radiant[0].Items.Item4.ID == 485).IsTrue()
			g.Assert(details.Radiant[0].Items.Item5.ID == 7).IsTrue()
			g.Assert(details.Radiant[0].Items.BackpackItem0.ID == 0).IsTrue()
			g.Assert(details.Radiant[0].Items.BackpackItem1.ID == 0).IsTrue()
			g.Assert(details.Radiant[0].Items.BackpackItem2.ID == 0).IsTrue()
			g.Assert(details.Radiant[0].Items.ItemNeutral.ID == 357).IsTrue()
		})
		g.It("Should return the correct AbilityBuild for player 0", func() {
			g.Assert(details.Radiant[0].AbilityUpgrades.Count()).Equal(18)
			aU, f := details.Radiant[0].AbilityUpgrades.GetByOrder(4)
			g.Assert(f).IsTrue()
			g.Assert(aU.Ability).Equal(5378)
			g.Assert(aU.Level).Equal(5)
			g.Assert(aU.Time).Equal(714 * time.Second)
			_, f = details.Radiant[0].AbilityUpgrades.GetByOrder(18)
			g.Assert(f).IsFalse()
			aU, f = details.Radiant[0].AbilityUpgrades.GetByLevel(10)
			g.Assert(f).IsTrue()
			g.Assert(aU.Ability).Equal(5906)
			g.Assert(aU.Level).Equal(10)
			g.Assert(aU.Time).Equal(1343 * time.Second)
			_, f = details.Radiant[0].AbilityUpgrades.GetByLevel(19)
			g.Assert(f).IsFalse()
			aUs := details.Radiant[0].AbilityUpgrades.GetByAbility(5378)
			g.Assert(aUs.Count()).Equal(4)
			for _, aU := range aUs {
				g.Assert(aU.Ability).Equal(5378)
			}
		})
		g.It("Should return working Abilities", func() {
			details.Dire.ForEach(func(p PlayerDetails) {
				for i, a := range p.AbilityUpgrades {
					for _, aU := range p.AbilityUpgrades.GetByAbility(a.Ability) {
						if aU.Ability == a.Ability {
							goto success
						}
					}
					g.Fail(fmt.Sprintf("missing required ability %d", a.Ability))
				success:
					aU, f := p.AbilityUpgrades.GetByOrder(i)
					g.Assert(f).IsTrue()
					g.Assert(aU).Equal(a)
					aU, f = p.AbilityUpgrades.GetByLevel(a.Level)
					g.Assert(f).IsTrue()
					g.Assert(aU).Equal(a)
				}
			})
		})
	})
}
