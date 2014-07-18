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




### Dota2 Api
----------------------------

#### GetMatchHistory 

**Url**

https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v001/

**Arguments**

key: Your Steam API Key

steamid (Optional): The 64 bit Steam ID the vanity URL resolves to. Not returned on resolution failures.

**Result Data**

status: 

num_result:

total_results:

results_remaining:

matches

> match_idi:
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

**Result Data**

#### GetPlayerSummaries

**Url**

https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/ 

**Arguments**

**Result Data**

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

#### GetHeroes

**Url**

https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001/

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
