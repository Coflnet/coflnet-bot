package hypixel

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/rs/zerolog/log"
)

func PlayerData(uuid string) (*PlayerDataResponse, error) {

  proxyUrl := fmt.Sprintf("http://sky-proxy.sky:8000/Proxy/hypixel?path=%s", url.QueryEscape("/player?uuid="+uuid))

  log.Info().Msgf("loading hypixel player data for uuid %s", uuid)
  log.Info().Msgf("url for proxy is: %s", proxyUrl)
	u, err := url.Parse(proxyUrl)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing url")
		return nil, err
	}

	request := http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{
			"API_KEY": {Key()},
		},
	}

  log.Info().Msgf("sending request: %s", request.URL.String())
	resp, err := http.DefaultClient.Do(&request)


	if err != nil {
		log.Error().Err(err).Msgf("error getting player data")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error reading response body")
		return nil, err
	}

	var data PlayerDataResponse
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		log.Error().Err(err).Msgf("error unmarshalling response")
		metrics.ErrorOccurred()
		return nil, err
	}

	return &data, nil
}

//goland:noinspection ALL
type PlayerDataResponse struct {
	Success bool `json:"success"`
	Player  struct {
		ID                     string        `json:"_id"`
		AchievementsOneTime    []string      `json:"achievementsOneTime"`
		Channel                string        `json:"channel"`
		Displayname            string        `json:"displayname"`
		FirstLogin             int64         `json:"firstLogin"`
		FriendRequests         []interface{} `json:"friendRequests"`
		FriendRequestsUUID     []interface{} `json:"friendRequestsUuid"`
		Karma                  int           `json:"karma"`
		KnownAliases           []string      `json:"knownAliases"`
		KnownAliasesLower      []string      `json:"knownAliasesLower"`
		LastLogin              int64         `json:"lastLogin"`
		MostRecentlyThanked    string        `json:"mostRecentlyThanked"`
		MostRecentlyTipped     string        `json:"mostRecentlyTipped"`
		MostRecentlyTippedUUID string        `json:"mostRecentlyTippedUuid"`
		ParkourCompletions     struct {
			Main []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"main"`
			CopsnCrims []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"CopsnCrims"`
			TheWallsLobby []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"TheWallsLobby"`
			TruePVPParkour []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"TruePVPParkour"`
			QuakeCraft []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"QuakeCraft"`
			NewMainLobby []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"NewMainLobby"`
			Skywars []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Skywars"`
			SuperSmash []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"SuperSmash"`
			Tnt []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"TNT"`
			Arena []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Arena"`
			Turbo []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Turbo"`
			Prototype []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Prototype"`
			Bedwars []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Bedwars"`
			SkywarsAug2017 []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"SkywarsAug2017"`
			MurderMystery []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"MurderMystery"`
			Legacy []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Legacy"`
			Tourney []struct {
				TimeStart int64 `json:"timeStart"`
				TimeTook  int   `json:"timeTook"`
			} `json:"Tourney"`
		} `json:"parkourCompletions"`
		Playername string `json:"playername"`
		Quests     struct {
			Halloween2014 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"halloween2014"`
			PaintballExpert struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Kill int `json:"kill"`
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"paintball_expert"`
			SpaceMission struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"space_mission"`
			WarlordsWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warlords_win"`
			TntAddict struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_addict"`
			Waller struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Kill int `json:"kill"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"waller"`
			SkywarsWeeklyKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsWeeklyKills int `json:"skywars_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_weekly_kills"`
			ExplosiveGames struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"explosive_games"`
			SerialKiller struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Quake     int `json:"quake"`
						Tnt       int `json:"tnt"`
						Blitz     int `json:"blitz"`
						Paintball int `json:"paintball"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"serial_killer"`
			NuggetWarriors struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Quake     int `json:"quake"`
						Paintball int `json:"paintball"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"nugget_warriors"`
			WelcomeToHell struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Megawalls int `json:"megawalls"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"welcome_to_hell"`
			WarriorsJourney struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Blitzkill         int `json:"blitzkill"`
						Quake25Kill       int `json:"quake25kill"`
						Tntwin            int `json:"tntwin"`
						Paintballwin      int `json:"paintballwin"`
						Vampirezkillvamps int `json:"vampirezkillvamps"`
						Megawallswin      int `json:"megawallswin"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"warriors_journey"`
			Blitzerk struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Killblitz10 int `json:"killblitz10"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"blitzerk"`
			Megawaller struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						Win int `json:"win"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"megawaller"`
			Gladiator struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"gladiator"`
			QuakeDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_play"`
			QuakeDailyKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_kill"`
			QuakeWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						QuakeWeeklyStreak int `json:"quake_weekly_streak"`
						QuakeWeeklyPlay   int `json:"quake_weekly_play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_weekly_play"`
			SkywarsTeamKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsTeamKills int `json:"skywars_team_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_team_kills"`
			SkywarsTeamWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_team_win"`
			SkywarsSoloKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SkywarsSoloKills int `json:"skywars_solo_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_solo_kills"`
			SkywarsSoloWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"skywars_solo_win"`
			CvcWinDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"cvc_win_daily_normal"`
			CvcKillDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"cvc_kill_daily_normal"`
			CvcKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						CvcKillDailyDeathmatch int `json:"cvc_kill_daily_deathmatch"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"cvc_kill"`
			CvcWinDailyDeathmatch struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"cvc_win_daily_deathmatch"`
			CvcKillWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						CvcPlayWeekly2 int `json:"cvc_play_weekly_2"`
						CvcPlayWeekly  int `json:"cvc_play_weekly"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"cvc_kill_weekly"`
			ArcadeGamer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_gamer"`
			ArcadeWinner struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_winner"`
			ArcadeSpecialist struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arcade_specialist"`
			SupersmashSoloWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_win"`
			SupersmashSoloKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_solo_kills"`
			SupersmashTeamWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_team_win"`
			SupersmashTeamKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SupersmashTeamKills int `json:"supersmash_team_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_team_kills"`
			SupersmashWeeklyKills struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						SupersmashWeeklyKills int `json:"supersmash_weekly_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"supersmash_weekly_kills"`
			QuakeDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"quake_daily_win"`
			BedwarsDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_win"`
			BedwarsDailyOneMore struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_daily_one_more"`
			BedwarsWeeklyBedElims struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsBedElims int `json:"bedwars_bed_elims"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_bed_elims"`
			Paintballer struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"paintballer"`
			PaintballKiller struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"paintball_killer"`
			WallsDailyPlay struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_play"`
			WallsDailyKill struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_kill"`
			WallsDailyWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_daily_win"`
			WallsWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"walls_weekly"`
			MmDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_daily_win"`
			MmDailyPowerPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_daily_power_play"`
			MmDailyTargetKill struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						MmTargetKills int `json:"mm_target_kills"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_daily_target_kill"`
			MmWeeklyMurdererKills struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						MmWeeklyKillsAsMurderer int `json:"mm_weekly_kills_as_murderer"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_weekly_murderer_kills"`
			MmWeeklyWins struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						MmWeeklyWin int `json:"mm_weekly_win"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_weekly_wins"`
			BedwarsWeeklyPumpkinator struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						BedwarsSpecialWeeklyPumpkinator int `json:"bedwars_special_weekly_pumpkinator"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_weekly_pumpkinator"`
			BedwarsWeeklySanta struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
			} `json:"bedwars_weekly_santa"`
			MmSpecialWeeklySanta struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						MmSpecialWeeklySanta int `json:"mm_special_weekly_santa"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mm_special_weekly_santa"`
			MegaWallsWin struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_win"`
			MegaWallsPlay struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_play"`
			MegaWallsKill struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_kill"`
			MegaWallsWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"mega_walls_weekly"`
			TntTnttagDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntTnttagDaily int `json:"tnt_tnttag_daily"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tnttag_daily"`
			TntTnttagWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						TntTnttagWeekly int `json:"tnt_tnttag_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_tnttag_weekly"`
			TntBowspleefDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_daily"`
			TntBowspleefWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntBowspleefWeekly int `json:"tnt_bowspleef_weekly"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_bowspleef_weekly"`
			TntPvprunDaily struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						TntPvprunDaily int `json:"tnt_pvprun_daily"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_pvprun_daily"`
			TntPvprunWeekly struct {
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						TntPvprunWeekly int `json:"tnt_pvprun_weekly"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_pvprun_weekly"`
			TntTntrunDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_daily"`
			TntTntrunWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_tntrun_weekly"`
			TntWizardsDaily struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWizardsDailyKills int `json:"tnt_wizards_daily_kills"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_wizards_daily"`
			TntWizardsWeekly struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						TntWizardsWeeklyKills int `json:"tnt_wizards_weekly_kills"`
						TntWizardsWeeklyWin   int `json:"tnt_wizards_weekly_win"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"tnt_wizards_weekly"`
			BedwarsWeeklyDreamWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						BedwarsDreamWins int `json:"bedwars_dream_wins"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_dream_win"`
			BedwarsDailyNightmares struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
						BedwarsDailyNightmareBeds int `json:"bedwars_daily_nightmare_beds"`
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_daily_nightmares"`
			TntDailyWin struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_daily_win"`
			TntWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_play"`
			BedwarsDailyGifts struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Started    int64 `json:"started"`
					Objectives struct {
					} `json:"objectives"`
				} `json:"active"`
			} `json:"bedwars_daily_gifts"`
			GingerbreadMastery struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_mastery"`
			GingerbreadMaps struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_maps"`
			GingerbreadBlingBling struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_bling_bling"`
			CrazyWallsDailyPlay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"crazy_walls_daily_play"`
			CrazyWallsDailyKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"crazy_walls_daily_kill"`
			CrazyWallsDailyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"crazy_walls_daily_win"`
			CrazyWallsWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"crazy_walls_weekly"`
			TntWeeklySpecial struct {
				Completions []struct {
					Time int64 `json:"time"`
				} `json:"completions"`
				Active struct {
					Objectives struct {
						TntWeeklyCandyCanes int `json:"tnt_weekly_candy_canes"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"tnt_weekly_special"`
			VampirezDailyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_win"`
			VampirezDailyKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_kill"`
			VampirezDailyPlay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_play"`
			ArenaWeeklyPlay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arena_weekly_play"`
			VampirezWeeklyKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_weekly_kill"`
			ArenaDailyKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arena_daily_kills"`
			ArenaDailyPlay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arena_daily_play"`
			GingerbreadRacer struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"gingerbread_racer"`
			VampirezWeeklyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_weekly_win"`
			ArenaDailyWins struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"arena_daily_wins"`
			VampirezDailyHumanKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_daily_human_kill"`
			VampirezWeeklyHumanKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"vampirez_weekly_human_kill"`
			SkywarsArcadeWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_arcade_win"`
			SkywarsSpecialNorthPole struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_special_north_pole"`
			SkywarsWeeklyArcadeWinAll struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_weekly_arcade_win_all"`
			SkywarsWeeklyFreeLootChest struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_weekly_free_loot_chest"`
			SkywarsCorruptWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_corrupt_win"`
			SkywarsDailyTokens struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"skywars_daily_tokens"`
			BlitzGameOfTheDay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_game_of_the_day"`
			BlitzSpecialDailyNorthPole struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_special_daily_north_pole"`
			BlitzKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_kills"`
			BlitzWeeklyMaster struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_weekly_master"`
			BlitzLootChestWeekly struct {
				Active struct {
					Objectives struct {
						Lootchestblitz  int `json:"lootchestblitz"`
						Dealdamageblitz int `json:"dealdamageblitz"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_loot_chest_weekly"`
			BlitzLootChestDaily struct {
				Active struct {
					Objectives struct {
						Lootchestblitz int `json:"lootchestblitz"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_loot_chest_daily"`
			BlitzWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"blitz_win"`
			DuelsWinner struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_winner"`
			DuelsWeeklyKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_weekly_kills"`
			DuelsPlayer struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_player"`
			DuelsWeeklyWins struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_weekly_wins"`
			DuelsKiller struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"duels_killer"`
			MegaWallsFaithful struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mega_walls_faithful"`
			WarlordsTdm struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_tdm"`
			WarlordsCtf struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_ctf"`
			WarlordsDomination struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_domination"`
			WarlordsVictorious struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_victorious"`
			WarlordsDedication struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_dedication"`
			WarlordsAllStar struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_all_star"`
			WarlordsObjectives struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"warlords_objectives"`
			BuildBattlePlayer struct {
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_player"`
			BuildBattleWinner struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_winner"`
			BuildBattleWeekly struct {
				Active struct {
					Objectives struct {
						Play int `json:"play"`
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_weekly"`
			BuildBattleChristmasWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_christmas_weekly"`
			BuildBattleChristmas struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"build_battle_christmas"`
			UhcSolo struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_solo"`
			UhcTeam struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_team"`
			UhcWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_weekly"`
			UhcMadness struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_madness"`
			TeamBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"team_brawler"`
			UhcDm struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_dm"`
			SoloBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"solo_brawler"`
			UhcWeeklySpecialCookie struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"uhc_weekly_special_cookie"`
			MmDailyInfector struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"mm_daily_infector"`
			BedwarsWeeklyChallenges struct {
				Active struct {
					Objectives struct {
					} `json:"objectives"`
					Started int64 `json:"started"`
				} `json:"active"`
			} `json:"bedwars_weekly_challenges"`
		} `json:"quests"`
		ThanksSent       int    `json:"thanksSent"`
		TimePlaying      int    `json:"timePlaying"`
		TournamentTokens int    `json:"tournamentTokens"`
		UUID             string `json:"uuid"`
		Eugene           struct {
			DailyTwoKExp int64 `json:"dailyTwoKExp"`
		} `json:"eugene"`
		McVersionRp string `json:"mcVersionRp"`
		Language    string `json:"language"`
		Voting      struct {
			Total               int   `json:"total"`
			TotalMcsorg         int   `json:"total_mcsorg"`
			SecondaryMcsorg     int   `json:"secondary_mcsorg"`
			LastMcsorg          int64 `json:"last_mcsorg"`
			VotesToday          int   `json:"votesToday"`
			LastVote            int64 `json:"last_vote"`
			SecondaryMcsl       int   `json:"secondary_mcsl"`
			TotalMcsl           int   `json:"total_mcsl"`
			LastMcsl            int64 `json:"last_mcsl"`
			TotalMcmp           int   `json:"total_mcmp"`
			SecondaryMcmp       int   `json:"secondary_mcmp"`
			LastMcmp            int64 `json:"last_mcmp"`
			TotalTopg           int   `json:"total_topg"`
			SecondaryTopg       int   `json:"secondary_topg"`
			LastTopg            int64 `json:"last_topg"`
			TotalMinestatus     int   `json:"total_minestatus"`
			SecondaryMinestatus int   `json:"secondary_minestatus"`
			LastMinestatus      int64 `json:"last_minestatus"`
			TotalMcipl          int   `json:"total_mcipl"`
			SecondaryMcipl      int   `json:"secondary_mcipl"`
			LastMcipl           int64 `json:"last_mcipl"`
			TotalMcf            int   `json:"total_mcf"`
			SecondaryMcf        int   `json:"secondary_mcf"`
			LastMcf             int64 `json:"last_mcf"`
		} `json:"voting"`
		PetConsumables struct {
			Feather      int `json:"FEATHER"`
			Bread        int `json:"BREAD"`
			Cake         int `json:"CAKE"`
			CookedBeef   int `json:"COOKED_BEEF"`
			RedRose      int `json:"RED_ROSE"`
			WaterBucket  int `json:"WATER_BUCKET"`
			Melon        int `json:"MELON"`
			Stick        int `json:"STICK"`
			WoodSword    int `json:"WOOD_SWORD"`
			MilkBucket   int `json:"MILK_BUCKET"`
			GoldRecord   int `json:"GOLD_RECORD"`
			Pork         int `json:"PORK"`
			Leash        int `json:"LEASH"`
			LavaBucket   int `json:"LAVA_BUCKET"`
			CarrotItem   int `json:"CARROT_ITEM"`
			RottenFlesh  int `json:"ROTTEN_FLESH"`
			SlimeBall    int `json:"SLIME_BALL"`
			RawFish      int `json:"RAW_FISH"`
			Bone         int `json:"BONE"`
			MagmaCream   int `json:"MAGMA_CREAM"`
			Cookie       int `json:"COOKIE"`
			Apple        int `json:"APPLE"`
			Wheat        int `json:"WHEAT"`
			MushroomSoup int `json:"MUSHROOM_SOUP"`
			HayBlock     int `json:"HAY_BLOCK"`
			PumpkinPie   int `json:"PUMPKIN_PIE"`
			BakedPotato  int `json:"BAKED_POTATO"`
		} `json:"petConsumables"`
		HousingMeta struct {
			AllowedBlocks      []string `json:"allowedBlocks"`
			TutorialStep       string   `json:"tutorialStep"`
			Packages           []string `json:"packages"`
			GivenCookies104981 []string `json:"given_cookies_104981"`
			FirstHouseJoinMs   int64    `json:"firstHouseJoinMs"`
			GivenCookies104995 []string `json:"given_cookies_104995"`
			GivenCookies105002 []string `json:"given_cookies_105002"`
			GivenCookies105013 []string `json:"given_cookies_105013"`
			GivenCookies105019 []string `json:"given_cookies_105019"`
			GivenCookies105021 []string `json:"given_cookies_105021"`
			PlotSize           string   `json:"plotSize"`
			PlayerSettings     struct {
				Border string `json:"BORDER"`
			} `json:"playerSettings"`
			GivenCookies105080 []string `json:"given_cookies_105080"`
			GivenCookies105104 []string `json:"given_cookies_105104"`
		} `json:"housingMeta"`
		VanityMeta struct {
			Packages []string `json:"packages"`
		} `json:"vanityMeta"`
		RewardConsumed          bool   `json:"rewardConsumed"`
		LevelingReward29        bool   `json:"levelingReward_29"`
		LastAdsenseGenerateTime int64  `json:"lastAdsenseGenerateTime"`
		TotalRewards            int    `json:"totalRewards"`
		TotalDailyRewards       int    `json:"totalDailyRewards"`
		RewardStreak            int    `json:"rewardStreak"`
		RewardScore             int    `json:"rewardScore"`
		RewardHighScore         int    `json:"rewardHighScore"`
		AdsenseTokens           int    `json:"adsense_tokens"`
		FlashingSalePopup       int64  `json:"flashingSalePopup"`
		FlashingSaleOpens       int    `json:"flashingSaleOpens"`
		FlashingSaleClicks      int    `json:"flashingSaleClicks"`
		FlashingSalePoppedUp    int    `json:"flashingSalePoppedUp"`
		NewPackageRank          string `json:"newPackageRank"`
		LevelUpVIP              int64  `json:"levelUp_VIP"`
		SpecialtyCooldowns      struct {
			Vip3     bool `json:"VIP3"`
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip6     bool `json:"VIP6"`
			Vip5     bool `json:"VIP5"`
			Vip4     bool `json:"VIP4"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus4 bool `json:"VIP_PLUS4"`
			VipPlus5 bool `json:"VIP_PLUS5"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			VipPlus6 bool `json:"VIP_PLUS6"`
		} `json:"specialtyCooldowns"`
		LevelUpVIPPLUS int64 `json:"levelUp_VIP_PLUS"`
		PetStats       struct {
			SheepYellow struct {
				Name     string `json:"name"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Hunger struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"HUNGER"`
				Experience int `json:"experience"`
			} `json:"SHEEP_YELLOW"`
			CowBaby struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"THIRST"`
				Exercise struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"EXERCISE"`
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Name       string `json:"name"`
				Experience int    `json:"experience"`
			} `json:"COW_BABY"`
			WildOcelotBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"HUNGER"`
				Thirst struct {
					Value     int   `json:"value"`
					Timestamp int64 `json:"timestamp"`
				} `json:"THIRST"`
				Exercise struct {
					Timestamp int64 `json:"timestamp"`
					Value     int   `json:"value"`
				} `json:"EXERCISE"`
				Experience int    `json:"experience"`
				Name       string `json:"name"`
			} `json:"WILD_OCELOT_BABY"`
			Zombie struct {
				Name string `json:"name"`
			} `json:"ZOMBIE"`
			Totem struct {
				Name string `json:"name"`
			} `json:"TOTEM"`
		} `json:"petStats"`
		PetJourneyTimestamp int64 `json:"petJourneyTimestamp"`
		SocialMedia         struct {
			Twitter string `json:"TWITTER"`
			Links   struct {
				Youtube string `json:"YOUTUBE"`
				Discord string `json:"DISCORD"`
			} `json:"links"`
			Prompt bool `json:"prompt"`
		} `json:"socialMedia"`
		UserLanguage        string   `json:"userLanguage"`
		SantaQuestStarted   bool     `json:"SANTA_QUEST_STARTED"`
		LastLogout          int64    `json:"lastLogout"`
		NetworkUpdateBook   string   `json:"network_update_book"`
		AchievementTracking []string `json:"achievementTracking"`
		GuildNotifications  bool     `json:"guildNotifications"`
		GiftingMeta         struct {
			RealBundlesReceived int      `json:"realBundlesReceived"`
			BundlesReceived     int      `json:"bundlesReceived"`
			GiftsGiven          int      `json:"giftsGiven"`
			BundlesGiven        int      `json:"bundlesGiven"`
			RealBundlesGiven    int      `json:"realBundlesGiven"`
			Milestones          []string `json:"milestones"`
		} `json:"giftingMeta"`
		FortuneBuff     int `json:"fortuneBuff"`
		AchievementSync struct {
			QuakeTiered int `json:"quake_tiered"`
		} `json:"achievementSync"`
		Challenges struct {
			AllTime struct {
				BEDWARSSupport                   int `json:"BEDWARS__support"`
				BEDWARSOffensive                 int `json:"BEDWARS__offensive"`
				MCGOKillingSpreeChallenge        int `json:"MCGO__killing_spree_challenge"`
				ARCADEGalaxyWarsChallenge        int `json:"ARCADE__galaxy_wars_challenge"`
				BEDWARSDefensive                 int `json:"BEDWARS__defensive"`
				MCGOKnifeChallenge               int `json:"MCGO__knife_challenge"`
				TNTGAMESTntTagChallenge          int `json:"TNTGAMES__tnt_tag_challenge"`
				TNTGAMESPvpRunChallenge          int `json:"TNTGAMES__pvp_run_challenge"`
				TNTGAMESBowSpleefChallenge       int `json:"TNTGAMES__bow_spleef_challenge"`
				TNTGAMESTntWizardChallenge       int `json:"TNTGAMES__tnt_wizard_challenge"`
				MURDERMYSTERYHero                int `json:"MURDER_MYSTERY__hero"`
				ARCADEFarmHuntChallenge          int `json:"ARCADE__farm_hunt_challenge"`
				ARCADEHideAndSeekChallenge       int `json:"ARCADE__hide_and_seek_challenge"`
				DUELSFeedTheVoidChallenge        int `json:"DUELS__feed_the_void_challenge"`
				PAINTBALLFinishChallenge         int `json:"PAINTBALL__finish_challenge"`
				QUAKECRAFTDonTBlinkChallenge     int `json:"QUAKECRAFT__don't_blink_challenge"`
				QUAKECRAFTComboChallenge         int `json:"QUAKECRAFT__combo_challenge"`
				MURDERMYSTERYSerialKiller        int `json:"MURDER_MYSTERY__serial_killer"`
				MCGOGrenadeChallenge             int `json:"MCGO__grenade_challenge"`
				DUELSTeamsChallenge              int `json:"DUELS__teams_challenge"`
				PAINTBALLKillingSpreeChallenge   int `json:"PAINTBALL__killing_spree_challenge"`
				PAINTBALLKillStreakChallenge     int `json:"PAINTBALL__kill_streak_challenge"`
				QUAKECRAFTKillingStreakChallenge int `json:"QUAKECRAFT__killing_streak_challenge"`
				QUAKECRAFTPowerupChallenge       int `json:"QUAKECRAFT__powerup_challenge"`
				TNTGAMESTntWizardsChallenge      int `json:"TNTGAMES__tnt_wizards_challenge"`
				TNTGAMESTntRunChallenge          int `json:"TNTGAMES__tnt_run_challenge"`
				ARCADEPartyGamesChallenge        int `json:"ARCADE__party_games_challenge"`
				MURDERMYSTERYSherlock            int `json:"MURDER_MYSTERY__sherlock"`
			} `json:"all_time"`
			DayL struct {
				MCGOGrenadeChallenge             int `json:"MCGO__grenade_challenge"`
				QUAKECRAFTComboChallenge         int `json:"QUAKECRAFT__combo_challenge"`
				QUAKECRAFTPowerupChallenge       int `json:"QUAKECRAFT__powerup_challenge"`
				QUAKECRAFTKillingStreakChallenge int `json:"QUAKECRAFT__killing_streak_challenge"`
				PAINTBALLKillStreakChallenge     int `json:"PAINTBALL__kill_streak_challenge"`
			} `json:"day_l"`
			DayA struct {
				PAINTBALLKillingSpreeChallenge int `json:"PAINTBALL__killing_spree_challenge"`
				PAINTBALLKillStreakChallenge   int `json:"PAINTBALL__kill_streak_challenge"`
			} `json:"day_a"`
			DayB struct {
				PAINTBALLKillingSpreeChallenge int `json:"PAINTBALL__killing_spree_challenge"`
				PAINTBALLKillStreakChallenge   int `json:"PAINTBALL__kill_streak_challenge"`
			} `json:"day_b"`
		} `json:"challenges"`
		ParkourCheckpointBests struct {
			Duels struct {
				Num0 int `json:"0"`
			} `json:"Duels"`
			Legacy struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
			} `json:"Legacy"`
			Prototype struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
			} `json:"Prototype"`
			Tourney struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"Tourney"`
			CopsnCrims struct {
				Num0 int `json:"0"`
				Num1 int `json:"1"`
			} `json:"CopsnCrims"`
		} `json:"parkourCheckpointBests"`
		AchievementPoints     int `json:"achievementPoints"`
		AchievementRewardsNew struct {
			ForPoints2000 int64 `json:"for_points_2000"`
			ForPoints3000 int64 `json:"for_points_3000"`
		} `json:"achievementRewardsNew"`
		AchievementTotem struct {
			CanCustomize     bool     `json:"canCustomize"`
			AllowedMaxHeight int      `json:"allowed_max_height"`
			UnlockedParts    []string `json:"unlockedParts"`
			SelectedParts    struct {
				Slot0 string `json:"slot_0"`
			} `json:"selectedParts"`
			UnlockedColors []string `json:"unlockedColors"`
			SelectedColors struct {
			} `json:"selectedColors"`
		} `json:"achievementTotem"`
		Tourney struct {
			FirstJoinLobby int64 `json:"first_join_lobby"`
			Bedwars4S0     struct {
				Playtime       int   `json:"playtime"`
				TributesEarned int   `json:"tributes_earned"`
				FirstWin       int64 `json:"first_win"`
			} `json:"bedwars4s_0"`
			TotalTributes int `json:"total_tributes"`
			McgoDefusal0  struct {
				SeenRPbook     bool  `json:"seenRPbook"`
				Playtime       int   `json:"playtime"`
				FirstWin       int64 `json:"first_win"`
				TributesEarned int   `json:"tributes_earned"`
			} `json:"mcgo_defusal_0"`
			BedwarsTwoFour0 struct {
				GamesPlayed    int   `json:"games_played"`
				Playtime       int   `json:"playtime"`
				TributesEarned int   `json:"tributes_earned"`
				FirstWin       int64 `json:"first_win"`
			} `json:"bedwars_two_four_0"`
			QuakeSolo20 struct {
				GamesPlayed    int   `json:"games_played"`
				Playtime       int   `json:"playtime"`
				TributesEarned int   `json:"tributes_earned"`
				FirstGame      int64 `json:"first_game"`
			} `json:"quake_solo2_0"`
			SwInsaneDoubles0 struct {
				GamesPlayed    int `json:"games_played"`
				Playtime       int `json:"playtime"`
				TributesEarned int `json:"tributes_earned"`
			} `json:"sw_insane_doubles_0"`
			GrinchSimulator0 struct {
				GamesPlayed    int   `json:"games_played"`
				Playtime       int   `json:"playtime"`
				FirstWin       int64 `json:"first_win"`
				TributesEarned int   `json:"tributes_earned"`
			} `json:"grinch_simulator_0"`
			McgoDefusal1 struct {
				SeenRPbook     bool  `json:"seenRPbook"`
				GamesPlayed    int   `json:"games_played"`
				Playtime       int   `json:"playtime"`
				FirstWin       int64 `json:"first_win"`
				TributesEarned int   `json:"tributes_earned"`
			} `json:"mcgo_defusal_1"`
			BedwarsEightTwo0 struct {
				GamesPlayed    int `json:"games_played"`
				Playtime       int `json:"playtime"`
				TributesEarned int `json:"tributes_earned"`
			} `json:"bedwars_eight_two_0"`
		} `json:"tourney"`
		CurrentCloak       string `json:"currentCloak"`
		ParticlePack       string `json:"particlePack"`
		CurrentClickEffect string `json:"currentClickEffect"`
		Monthlycrates      struct {
			One2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"1-2017"`
			One2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"1-2018"`
			One2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"1-2019"`
			One02016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"10-2016"`
			One02017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"10-2017"`
			One02018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
			} `json:"10-2018"`
			One12016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"11-2016"`
			One12017 struct {
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"11-2017"`
			One12018 struct {
				Regular bool `json:"REGULAR"`
			} `json:"11-2018"`
			One22017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"12-2017"`
			One22018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"12-2018"`
			Two2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"2-2017"`
			Two2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"2-2018"`
			Two2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"2-2019"`
			Three2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"3-2017"`
			Four2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"4-2017"`
			Four2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"4-2019"`
			Five2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"5-2017"`
			Six2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"6-2017"`
			Six2018 struct {
				Regular bool `json:"REGULAR"`
			} `json:"6-2018"`
			Seven2017 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"7-2017"`
			Seven2018 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"7-2018"`
			Seven2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"7-2019"`
			Eight2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"8-2016"`
			Eight2017 struct {
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"8-2017"`
			Eight2018 struct {
				Regular bool `json:"REGULAR"`
			} `json:"8-2018"`
			Eight2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"8-2019"`
			Nine2016 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"9-2016"`
			Nine2017 struct {
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
			} `json:"9-2017"`
			Nine2018 struct {
				Regular bool `json:"REGULAR"`
			} `json:"9-2018"`
			Nine2019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"9-2019"`
			One12019 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"11-2019"`
			One2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"1-2020"`
			Two2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"2-2020"`
			Three2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"3-2020"`
			Four2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"4-2020"`
			Five2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"5-2020"`
			Six2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"6-2020"`
			Seven2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"7-2020"`
			Eight2020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"8-2020"`
			Nine2020 struct {
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
			} `json:"9-2020"`
			One12020 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"11-2020"`
			One22020 struct {
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
				Vip     bool `json:"VIP"`
				Regular bool `json:"REGULAR"`
			} `json:"12-2020"`
			One2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"1-2021"`
			Two2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				MvpPlus bool `json:"MVP_PLUS"`
				Mvp     bool `json:"MVP"`
				VipPlus bool `json:"VIP_PLUS"`
			} `json:"2-2021"`
			Three2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2021"`
			Four2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
			} `json:"4-2021"`
			Five2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"5-2021"`
			Seven2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"7-2021"`
			Eight2021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"8-2021"`
			One02021 struct {
				Regular bool `json:"REGULAR"`
			} `json:"10-2021"`
			One22021 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"12-2021"`
			Two2022 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"2-2022"`
			Three2022 struct {
				Regular bool `json:"REGULAR"`
				Vip     bool `json:"VIP"`
				VipPlus bool `json:"VIP_PLUS"`
				Mvp     bool `json:"MVP"`
				MvpPlus bool `json:"MVP_PLUS"`
			} `json:"3-2022"`
			Four2022 struct {
				Regular bool `json:"REGULAR"`
			} `json:"4-2022"`
		} `json:"monthlycrates"`
		Main2017Tutorial       bool   `json:"main2017Tutorial"`
		CurrentGadget          string `json:"currentGadget"`
		SnowballFightIntro2019 bool   `json:"snowball_fight_intro_2019"`
		GiftsGrinch            int    `json:"gifts_grinch"`
		ClaimedPotatoTalisman  int64  `json:"claimed_potato_talisman"`
		Outfit                 struct {
			Chestplate string `json:"CHESTPLATE"`
			Leggings   string `json:"LEGGINGS"`
			Boots      string `json:"BOOTS"`
			Helmet     string `json:"HELMET"`
		} `json:"outfit"`
		SkyblockFreeCookie int64  `json:"skyblock_free_cookie"`
		ClaimedCenturyCake int64  `json:"claimed_century_cake"`
		LevelUpMVPPLUS     int64  `json:"levelUp_MVP_PLUS"`
		LevelingReward94   bool   `json:"levelingReward_94"`
		RankPlusColor      string `json:"rankPlusColor"`
		QuestSettings      struct {
			AutoActivate bool `json:"autoActivate"`
		} `json:"questSettings"`
		LevelingReward149          bool  `json:"levelingReward_149"`
		ScorpiusBribe120           int64 `json:"scorpius_bribe_120"`
		AnniversaryNPCVisited2021  []int `json:"anniversaryNPCVisited2021"`
		AnniversaryNPCProgress2021 int   `json:"anniversaryNPCProgress2021"`
		Easter2021Cooldowns2       struct {
			VipPlus3 bool `json:"VIP_PLUS3"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			Normal0  bool `json:"NORMAL0"`
			Normal1  bool `json:"NORMAL1"`
			Normal2  bool `json:"NORMAL2"`
			Normal3  bool `json:"NORMAL3"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			MvpPlus3 bool `json:"MVP_PLUS3"`
			Mvp0     bool `json:"MVP0"`
			Mvp1     bool `json:"MVP1"`
			Mvp2     bool `json:"MVP2"`
			Mvp3     bool `json:"MVP3"`
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip3     bool `json:"VIP3"`
		} `json:"easter2021Cooldowns2"`
		AdventRewards2017 struct {
			Day2  int64 `json:"day2"`
			Day3  int64 `json:"day3"`
			Day4  int64 `json:"day4"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day7  int64 `json:"day7"`
			Day8  int64 `json:"day8"`
			Day9  int64 `json:"day9"`
			Day10 int64 `json:"day10"`
			Day11 int64 `json:"day11"`
			Day12 int64 `json:"day12"`
			Day13 int64 `json:"day13"`
			Day14 int64 `json:"day14"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day17 int64 `json:"day17"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day21 int64 `json:"day21"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
			Day22 int64 `json:"day22"`
			Day1  int64 `json:"day1"`
		} `json:"adventRewards2017"`
		AdventRewards2019 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day3  int64 `json:"day3"`
			Day4  int64 `json:"day4"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day7  int64 `json:"day7"`
			Day8  int64 `json:"day8"`
			Day9  int64 `json:"day9"`
			Day10 int64 `json:"day10"`
			Day11 int64 `json:"day11"`
			Day12 int64 `json:"day12"`
			Day13 int64 `json:"day13"`
			Day14 int64 `json:"day14"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day17 int64 `json:"day17"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day21 int64 `json:"day21"`
			Day22 int64 `json:"day22"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards2019"`
		AdventRewards2020 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day7  int64 `json:"day7"`
			Day8  int64 `json:"day8"`
			Day9  int64 `json:"day9"`
			Day12 int64 `json:"day12"`
			Day13 int64 `json:"day13"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day17 int64 `json:"day17"`
			Day19 int64 `json:"day19"`
			Day20 int64 `json:"day20"`
			Day21 int64 `json:"day21"`
			Day22 int64 `json:"day22"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards2020"`
		AdventRewardsV22018 struct {
			Day1  int64 `json:"day1"`
			Day2  int64 `json:"day2"`
			Day3  int64 `json:"day3"`
			Day4  int64 `json:"day4"`
			Day5  int64 `json:"day5"`
			Day6  int64 `json:"day6"`
			Day9  int64 `json:"day9"`
			Day10 int64 `json:"day10"`
			Day13 int64 `json:"day13"`
			Day15 int64 `json:"day15"`
			Day16 int64 `json:"day16"`
			Day17 int64 `json:"day17"`
			Day18 int64 `json:"day18"`
			Day19 int64 `json:"day19"`
			Day21 int64 `json:"day21"`
			Day22 int64 `json:"day22"`
			Day23 int64 `json:"day23"`
			Day24 int64 `json:"day24"`
			Day25 int64 `json:"day25"`
		} `json:"adventRewards_v2_2018"`
		ClaimedYear143Cake  int64  `json:"claimed_year143_cake"`
		VanityFavorites     string `json:"vanityFavorites"`
		Summer2021Cooldowns struct {
			MvpPlus3 bool `json:"MVP_PLUS3"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Normal3  bool `json:"NORMAL3"`
			Normal2  bool `json:"NORMAL2"`
			Normal1  bool `json:"NORMAL1"`
			Normal0  bool `json:"NORMAL0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			Mvp0     bool `json:"MVP0"`
			Mvp1     bool `json:"MVP1"`
			Mvp2     bool `json:"MVP2"`
			Mvp3     bool `json:"MVP3"`
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip3     bool `json:"VIP3"`
		} `json:"summer2021Cooldowns"`
		Halloween2021Cooldowns struct {
			MvpPlus3 bool `json:"MVP_PLUS3"`
			MvpPlus2 bool `json:"MVP_PLUS2"`
			MvpPlus1 bool `json:"MVP_PLUS1"`
			MvpPlus0 bool `json:"MVP_PLUS0"`
			Normal3  bool `json:"NORMAL3"`
			Normal2  bool `json:"NORMAL2"`
			Normal1  bool `json:"NORMAL1"`
			Normal0  bool `json:"NORMAL0"`
			VipPlus0 bool `json:"VIP_PLUS0"`
			VipPlus1 bool `json:"VIP_PLUS1"`
			VipPlus2 bool `json:"VIP_PLUS2"`
			VipPlus3 bool `json:"VIP_PLUS3"`
			Mvp0     bool `json:"MVP0"`
			Mvp1     bool `json:"MVP1"`
			Mvp2     bool `json:"MVP2"`
			Mvp3     bool `json:"MVP3"`
			Vip0     bool `json:"VIP0"`
			Vip1     bool `json:"VIP1"`
			Vip2     bool `json:"VIP2"`
			Vip3     bool `json:"VIP3"`
		} `json:"halloween2021Cooldowns"`
		Seasonal struct {
			Christmas struct {
				Num2021 struct {
					AdventRewards struct {
						Day1  int64 `json:"day1"`
						Day2  int64 `json:"day2"`
						Day4  int64 `json:"day4"`
						Day5  int64 `json:"day5"`
						Day6  int64 `json:"day6"`
						Day7  int64 `json:"day7"`
						Day8  int64 `json:"day8"`
						Day9  int64 `json:"day9"`
						Day11 int64 `json:"day11"`
						Day12 int64 `json:"day12"`
						Day13 int64 `json:"day13"`
						Day15 int64 `json:"day15"`
						Day16 int64 `json:"day16"`
						Day17 int64 `json:"day17"`
						Day18 int64 `json:"day18"`
						Day19 int64 `json:"day19"`
						Day20 int64 `json:"day20"`
						Day22 int64 `json:"day22"`
						Day23 int64 `json:"day23"`
						Day24 int64 `json:"day24"`
						Day25 int64 `json:"day25"`
					} `json:"adventRewards"`
					Presents struct {
						Classic1      bool `json:"CLASSIC_1"`
						Classic2      bool `json:"CLASSIC_2"`
						Arcade2       bool `json:"ARCADE_2"`
						Arcade3       bool `json:"ARCADE_3"`
						Arcade1       bool `json:"ARCADE_1"`
						Prototype1    bool `json:"PROTOTYPE_1"`
						MainLobby1    bool `json:"MAIN_LOBBY_1"`
						MainLobby2    bool `json:"MAIN_LOBBY_2"`
						MainLobby12   bool `json:"MAIN_LOBBY_12"`
						MainLobby6    bool `json:"MAIN_LOBBY_6"`
						MainLobby7    bool `json:"MAIN_LOBBY_7"`
						MainLobby8    bool `json:"MAIN_LOBBY_8"`
						MainLobby37   bool `json:"MAIN_LOBBY_37"`
						MainLobby38   bool `json:"MAIN_LOBBY_38"`
						MainLobby11   bool `json:"MAIN_LOBBY_11"`
						MainLobby10   bool `json:"MAIN_LOBBY_10"`
						MainLobby36   bool `json:"MAIN_LOBBY_36"`
						MainLobby22   bool `json:"MAIN_LOBBY_22"`
						MainLobby25   bool `json:"MAIN_LOBBY_25"`
						MainLobby4    bool `json:"MAIN_LOBBY_4"`
						MainLobby35   bool `json:"MAIN_LOBBY_35"`
						MainLobby3    bool `json:"MAIN_LOBBY_3"`
						MainLobby24   bool `json:"MAIN_LOBBY_24"`
						MainLobby23   bool `json:"MAIN_LOBBY_23"`
						MainLobby27   bool `json:"MAIN_LOBBY_27"`
						MainLobby5    bool `json:"MAIN_LOBBY_5"`
						MainLobby33   bool `json:"MAIN_LOBBY_33"`
						MainLobby34   bool `json:"MAIN_LOBBY_34"`
						MainLobby13   bool `json:"MAIN_LOBBY_13"`
						MainLobby14   bool `json:"MAIN_LOBBY_14"`
						MainLobby15   bool `json:"MAIN_LOBBY_15"`
						MainLobby9    bool `json:"MAIN_LOBBY_9"`
						MainLobby28   bool `json:"MAIN_LOBBY_28"`
						MainLobby29   bool `json:"MAIN_LOBBY_29"`
						MainLobby21   bool `json:"MAIN_LOBBY_21"`
						MainLobby32   bool `json:"MAIN_LOBBY_32"`
						MainLobby17   bool `json:"MAIN_LOBBY_17"`
						MainLobby16   bool `json:"MAIN_LOBBY_16"`
						MainLobby31   bool `json:"MAIN_LOBBY_31"`
						MainLobby19   bool `json:"MAIN_LOBBY_19"`
						MainLobby18   bool `json:"MAIN_LOBBY_18"`
						MainLobby30   bool `json:"MAIN_LOBBY_30"`
						MainLobby20   bool `json:"MAIN_LOBBY_20"`
						MainLobby26   bool `json:"MAIN_LOBBY_26"`
						MainLobby40   bool `json:"MAIN_LOBBY_40"`
						MainLobby39   bool `json:"MAIN_LOBBY_39"`
						Bedwars1      bool `json:"BEDWARS_1"`
						Bedwars2      bool `json:"BEDWARS_2"`
						Bedwars3      bool `json:"BEDWARS_3"`
						Bedwars4      bool `json:"BEDWARS_4"`
						Bedwars5      bool `json:"BEDWARS_5"`
						Skywars1      bool `json:"SKYWARS_1"`
						Skywars2      bool `json:"SKYWARS_2"`
						Skywars5      bool `json:"SKYWARS_5"`
						Skywars4      bool `json:"SKYWARS_4"`
						Skywars3      bool `json:"SKYWARS_3"`
						Murder1       bool `json:"MURDER_1"`
						Murder3       bool `json:"MURDER_3"`
						Murder2       bool `json:"MURDER_2"`
						Housing1      bool `json:"HOUSING_1"`
						Housing5      bool `json:"HOUSING_5"`
						Housing3      bool `json:"HOUSING_3"`
						Housing2      bool `json:"HOUSING_2"`
						Housing4      bool `json:"HOUSING_4"`
						Buildbattle1  bool `json:"BUILDBATTLE_1"`
						Buildbattle2  bool `json:"BUILDBATTLE_2"`
						Buildbattle3  bool `json:"BUILDBATTLE_3"`
						Duels1        bool `json:"DUELS_1"`
						Duels3        bool `json:"DUELS_3"`
						Duels2        bool `json:"DUELS_2"`
						Prototype3    bool `json:"PROTOTYPE_3"`
						Prototype2    bool `json:"PROTOTYPE_2"`
						Uhc1          bool `json:"UHC_1"`
						Uhc2          bool `json:"UHC_2"`
						Uhc3          bool `json:"UHC_3"`
						Tnt1          bool `json:"TNT_1"`
						Tnt2          bool `json:"TNT_2"`
						Tnt3          bool `json:"TNT_3"`
						Classic3      bool `json:"CLASSIC_3"`
						CopsAndCrims3 bool `json:"COPS_AND_CRIMS_3"`
						CopsAndCrims2 bool `json:"COPS_AND_CRIMS_2"`
						CopsAndCrims1 bool `json:"COPS_AND_CRIMS_1"`
						Blitz1        bool `json:"BLITZ_1"`
						Blitz2        bool `json:"BLITZ_2"`
						Blitz3        bool `json:"BLITZ_3"`
						Megawalls1    bool `json:"MEGAWALLS_1"`
						Megawalls2    bool `json:"MEGAWALLS_2"`
						Megawalls3    bool `json:"MEGAWALLS_3"`
						Smash1        bool `json:"SMASH_1"`
						Smash2        bool `json:"SMASH_2"`
						Smash3        bool `json:"SMASH_3"`
						Warlords2     bool `json:"WARLORDS_2"`
						Warlords3     bool `json:"WARLORDS_3"`
						Warlords1     bool `json:"WARLORDS_1"`
					} `json:"presents"`
				} `json:"2021"`
			} `json:"christmas"`
			Anniversary struct {
				Num2022 struct {
					AnniversaryNPCVisited  []int `json:"anniversaryNPCVisited"`
					AnniversaryNPCProgress int   `json:"anniversaryNPCProgress"`
				} `json:"2022"`
			} `json:"anniversary"`
			Easter struct {
				Num2022 struct {
					MainlobbyEgghunt1284514  bool `json:"mainlobby_egghunt_128_45_-14"`
					MainlobbyEgghunt15257184 bool `json:"mainlobby_egghunt_-152_57_184"`
					MainlobbyEgghunt56957    bool `json:"mainlobby_egghunt_-56_95_7"`
					MainlobbyEgghunt337344   bool `json:"mainlobby_egghunt_-33_73_-44"`
					MainlobbyEgghunt153452   bool `json:"mainlobby_egghunt_-15_34_-52"`
					MainlobbyEgghunt252266   bool `json:"mainlobby_egghunt_25_22_66"`
					MainlobbyEgghunt286818   bool `json:"mainlobby_egghunt_28_68_-18"`
					MainlobbyEgghunt357022   bool `json:"mainlobby_egghunt_35_70_22"`
					MainlobbyEgghunt366455   bool `json:"mainlobby_egghunt_36_64_55"`
					MainlobbyEgghunt27056    bool `json:"mainlobby_egghunt_2_70_56"`
					MainlobbyEgghunt83865    bool `json:"mainlobby_egghunt_83_86_-5"`
					MainlobbyEgghunt947043   bool `json:"mainlobby_egghunt_94_70_43"`
					MainlobbyEgghunt140595   bool `json:"mainlobby_egghunt_140_59_5"`
					MainlobbyEgghunt64104129 bool `json:"mainlobby_egghunt_64_104_-129"`
					MainlobbyEgghunt886590   bool `json:"mainlobby_egghunt_-88_65_-90"`
					MainlobbyEgghunt13664142 bool `json:"mainlobby_egghunt_-136_64_-142"`
					MainlobbyEgghunt17520109 bool `json:"mainlobby_egghunt_-175_20_-109"`
					MainlobbyEgghunt2027822  bool `json:"mainlobby_egghunt_-202_78_-22"`
					MainlobbyEgghunt1526541  bool `json:"mainlobby_egghunt_-152_65_41"`
					MainlobbyEgghunt9812920  bool `json:"mainlobby_egghunt_-98_129_-20"`
					MainlobbyEgghunt7710918  bool `json:"mainlobby_egghunt_-77_109_18"`
					MainlobbyEgghunt808959   bool `json:"mainlobby_egghunt_-80_89_59"`
					MainlobbyEgghunt11762107 bool `json:"mainlobby_egghunt_-117_62_107"`
					MainlobbyEgghunt8265160  bool `json:"mainlobby_egghunt_-82_65_160"`
					MainlobbyEgghunt1157203  bool `json:"mainlobby_egghunt_-11_57_203"`
					MainlobbyEgghunt4253253  bool `json:"mainlobby_egghunt_42_53_253"`
					MainlobbyEgghunt12652179 bool `json:"mainlobby_egghunt_126_52_179"`
					MainlobbyEgghunt6525121  bool `json:"mainlobby_egghunt_65_25_121"`
					MainlobbyEgghunt12967118 bool `json:"mainlobby_egghunt_129_67_118"`
					MainlobbyEgghunt1156186  bool `json:"mainlobby_egghunt_115_61_86"`
					MainlobbyEgghuntReward   bool `json:"mainlobby_egghunt_reward"`
				} `json:"2022"`
			} `json:"easter"`
		} `json:"seasonal"`
		CompletedChristmasQuests2022 int    `json:"completed_christmas_quests_2022"`
		CurrentPet                   string `json:"currentPet"`
		ClaimedCenturyCake200        int64  `json:"claimed_century_cake200"`
		MostRecentGameType           string `json:"mostRecentGameType"`
	} `json:"player"`
}
