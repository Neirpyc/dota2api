dota2-api
=========

golang dota2-api


### Usage


### Steam Api
---------------------------
#### ResolveVanityURL 

**Url**

http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/

**Arguments**

key: Your Steam API Key

vanityurl: The user's vanity URL that you would like to retrieve a steam ID for, e.g. http://steamcommunity.com/id/gabelogannewell would use "gabelogannewell"

**Result Data**

success: The status of the request. 1 if successful, 42 if there was no match.

steamid (Optional): The 64 bit Steam ID the vanity URL resolves to. Not returned on resolution failures.

message (Optional): The message associated with the request status. Currently only used on resolution failures.

#### GetPlayerSummaries

**Url**

http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/

key

steamids


**Result Data**

A list of profile objects. Contained information varies depending on whether or not the user has their profile set to Friends only or Private.

steamid: The user's 64 bit ID

communityvisibilitystate: An integer that describes the access setting of the profile

>1 - Private

>2 - Friends only

>3 - Friends of Friends

>4 - Users Only

>5 - Public

profilestate: If set to 1 the user has configured the profile.

personaname: User's display name.

lastlogoff: A unix timestamp of when the user was last online.

profileurl: The URL to the user's Steam Community profile.

avatar: A 32x32 image

avatarmedium: A 64x64 image

avatarfull: A 184x184 image

personastate: The user's status

>0 - Offline (Also set when the profile is Private)

>1 - Online

>2 - Busy

>3 - Away

>4 - Snooze

>5 - Looking to trade

>6 - Looking to play

commentpermission (Optional): If present the profile allows public comments.

realname (Optional): The user's real name.

primaryclanid (Optional): The 64 bit ID of the user's primary group.

timecreated (Optional): A unix timestamp of the date the profile was created.

loccountrycode (Optional): ISO 3166 code of where the user is located.

locstatecode (Optional): Variable length code representing the state the user is located in.

loccityid (Optional): An integer ID internal to Steam representing the user's city.

gameid (Optional): If the user is in game this will be set to it's app ID as a string.

gameextrainfo (Optional): The title of the game.

gameserverip (Optional): The server URL given as an IP address and port number separated by a colon, this will not be present or set to "0.0.0.0:0" if none is available.


### Dota2 Api
----------------------------

#### GetMatchHistory 

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v001/

**Arguments**

key: Your Steam API Key

hero_id (Optional) (uint32): A list of hero IDs can be found via the GetHeroes method.

game_mode (Optional) (uint32):

>0 - None

>1 - All Pick

>2 - Captain's Mode

>3 - Random Draft

>4 - Single Draft

>5 - All Random

>6 - Intro

>7 - Diretide

>8 - Reverse Captain's Mode

>9 - The Greeviling

>10 - Tutorial

>11 - Mid Only

>12 - Least Played

>13 - New Player Pool

>14 - Compendium Matchmaking

>16 - Captain's Draft

skill (Optional) (uint32): Skill bracket for the matches (Ignored if an account ID is specified).

>0 - Any

>1 - Normal

>2 - High

>3 - Very High

date_min (Optional) (uint32): Minimum date range for returned matches (unix timestamp, rounded to the nearest day).

date_max (Optional) (uint32): Maximum date range for returned matches (unix timestamp, rounded to the nearest day).

min_players (Optional) (string): Minimum amount of players in a match for the match to be returned.

account_id (Optional) (string): 32-bit account ID.

league_id (Optional) (string): Only return matches from this league. A list of league IDs can be found via the GetLeagueListing method.

start_at_match_id (Optional) (string): Start searching for matches equal to or older than this match ID.

matches_requested (Optional) (string): Amount of matches to include in results (default: 25).

tournament_games_only (Optional) (string): Whether to limit results to tournament matches

**Result Data**

status: 

num_result:

total_results:

results_remaining:

matches

> match_id:

> match_seq_num:

> start_time:

> lobby_type:

> radiant_team_id:

> dire_team_id:

> players:

>> account_id:

>> player_slot:

>> hero_id:


#### GetMatchDetails

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v001/ 

**Arguments**

match_id: Match id

key: Your Steam API Key

**Result Data**

players: List of players in the match.

account_id: 32-bit account ID

player_slot: 

hero_id: The hero's unique ID. A list of hero IDs can be found via the GetHeroes method.

item_0: ID of the top-left inventory item.

item_1: ID of the top-center inventory item.

item_2: ID of the top-right inventory item.

item_3: ID of the bottom-left inventory item.

item_4: ID of the bottom-center inventory item.

item_5: ID of the bottom-right inventory item.

kills: The amount of kills attributed to this player.

deaths: The amount of times this player died during the match.

assists: The amount of assists attributed to this player.

leaver_status: What the values here represent are not yet known.

gold: The amount of gold the player had remaining at the end of the match.

last_hits: The amount of last-hits the player got during the match.

denies: The amount of denies the player got during the match.

gold_per_min: The player's overall gold/minute.

xp_per_min: The player's overall experience/minute.

gold_spent: The amount of gold the player spent during the match.

hero_damage: The amount of damage the player dealt to heroes.

tower_damage: The amount of damage the player dealt to towers.

hero_healing: The amount of health the player had healed on heroes.

level: The player's level at match end.

ability_upgrades: A list detailing a player's ability upgrades.

ability: ID of the ability upgraded.

time: Time since match start that the ability was upgraded.

level: The level of the player at time of upgrading.

additional_units: Additional playable units owned by the player.

unitname: The name of the unit

item_0: ID of the top-left inventory item.

item_1: ID of the top-center inventory item.

item_2: ID of the top-right inventory item.

item_3: ID of the bottom-left inventory item.

item_4: ID of the bottom-center inventory item.

item_5: ID of the bottom-right inventory item.

season: The season the game was played in.

radiant_win: Dictates the winner of the match, true for radiant; false for dire.

duration: The length of the match, in seconds since the match began.

start_time: Unix timestamp of when the match began.

match_id: The matches unique ID.

match_seq_num: A 'sequence number', representing the order in which matches were recorded.

tower_status_radiant: 

tower_status_dire: 

barracks_status_radiant:

barracks_status_dire:

cluster: The server cluster the match was played upon. Used for downloading replays of matches.

first_blood_time: The time in seconds since the match began when first-blood occured.

lobby_type:

>-1 - Invalid

>0 - Public matchmaking

>1 - Practise

>2 - Tournament

>3 - Tutorial

>4 - Co-op with bots.

>5 - Team match

>6 - Solo Queue

human_players: The amount of human players within the match.

leagueid: The league that this match was a part of. A list of league IDs can be found via the GetLeagueListing method.

positive_votes: The number of thumbs-up the game has received by users.

negative_votes: The number of thumbs-down the game has received by users.

game_mode:

>0 - None

>1 - All Pick

>2 - Captain's Mode

>3 - Random Draft

>4 - Single Draft

>5 - All Random

>6 - Intro

>7 - Diretide

>8 - Reverse Captain's Mode

>9 - The Greeviling

>10 - Tutorial

>11 - Mid Only

>12 - Least Played

>13 - New Player Pool

>14 - Compendium Matchmaking

picks_bans: A list of the picks and bans in the match, if the game mode is Captains Mode.

is_pick: Whether this entry is a pick (true) or a ban (false).

hero_id: The hero's unique ID. A list of hero IDs can be found via the GetHeroes method.

team: The team who chose the pick or ban; 0 for Radiant, 1 for Dire.

order: The order of which the picks and bans were selected; 0-19.

#### GetHeroes

**Url**

https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/

**Arguments**

language (Optional) (string): The language to provide hero names in.

itemizedonly (Optional) (bool): Return a list of itemized heroes only.

**Result Data**

heroes: List of heroes.

name: The tokenized string for the name of the hero.

id: ID of the hero.

localized_name: The localized name of the hero for use in name display.

count: Number of results.

#### GetLeagueListing

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetLeagueListing/v0001/ 

**Arguments**

**Result Data**

#### GetLiveLeagueGames 

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetLiveLeagueGames/v0001/

**Arguments**

**Result Data**

#### GetTeamInfoByTeamID 

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetTeamInfoByTeamID/v001/ 

**Arguments**

**Result Data**


#### GetTournamentPrizePool 

**Url**

https://api.steampowered.com/IEconDOTA2_570/GetTournamentPrizePool/v1/

**Arguments**

**Result Data**

#### GetGameItems

**Url**

https://api.steampowered.com/IEconDOTA2_570/GetGameItems/v0001/

**Arguments**

**Result Data**
