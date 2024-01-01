package gen

// HypixelPlayerResponse / HypixelPlayerResponse this code is generated with json 2 go
type HypixelPlayerResponse struct {
	Success bool `json:"success,omitempty"`
	Player  struct {
		ID           string `json:"_id,omitempty"`
		Achievements struct {
			ArenaBossed                      int `json:"arena_bossed,omitempty"`
			ArenaClimbTheRanks               int `json:"arena_climb_the_ranks,omitempty"`
			ArenaGladiator                   int `json:"arena_gladiator,omitempty"`
			ArenaGottaWearEmAll              int `json:"arena_gotta_wear_em_all,omitempty"`
			BlitzCoins                       int `json:"blitz_coins,omitempty"`
			BlitzKills                       int `json:"blitz_kills,omitempty"`
			BlitzWins                        int `json:"blitz_wins,omitempty"`
			GeneralCoins                     int `json:"general_coins,omitempty"`
			GeneralWins                      int `json:"general_wins,omitempty"`
			PaintballCoins                   int `json:"paintball_coins,omitempty"`
			PaintballKills                   int `json:"paintball_kills,omitempty"`
			PaintballWins                    int `json:"paintball_wins,omitempty"`
			QuakeKillingSprees               int `json:"quake_killing_sprees,omitempty"`
			QuakeKills                       int `json:"quake_kills,omitempty"`
			QuakeWins                        int `json:"quake_wins,omitempty"`
			TntgamesBowSpleefWins            int `json:"tntgames_bow_spleef_wins,omitempty"`
			TntgamesTntRunWins               int `json:"tntgames_tnt_run_wins,omitempty"`
			TntgamesWizardsWins              int `json:"tntgames_wizards_wins,omitempty"`
			VampirezCoins                    int `json:"vampirez_coins,omitempty"`
			VampirezKillSurvivors            int `json:"vampirez_kill_survivors,omitempty"`
			VampirezKillVampires             int `json:"vampirez_kill_vampires,omitempty"`
			VampirezSurvivorWins             int `json:"vampirez_survivor_wins,omitempty"`
			Walls3Coins                      int `json:"walls3_coins,omitempty"`
			Walls3Kills                      int `json:"walls3_kills,omitempty"`
			Walls3Wins                       int `json:"walls3_wins,omitempty"`
			WallsKills                       int `json:"walls_kills,omitempty"`
			WallsWins                        int `json:"walls_wins,omitempty"`
			WarlordsKills                    int `json:"warlords_kills,omitempty"`
			WarlordsPaladinLevel             int `json:"warlords_paladin_level,omitempty"`
			WarlordsWarriorLevel             int `json:"warlords_warrior_level,omitempty"`
			WarlordsShamanLevel              int `json:"warlords_shaman_level,omitempty"`
			WarlordsMageLevel                int `json:"warlords_mage_level,omitempty"`
			WarlordsAssist                   int `json:"warlords_assist,omitempty"`
			SkywarsWinsSolo                  int `json:"skywars_wins_solo,omitempty"`
			SkywarsKitsTeam                  int `json:"skywars_kits_team,omitempty"`
			SkywarsWinsTeam                  int `json:"skywars_wins_team,omitempty"`
			SkywarsKillsTeam                 int `json:"skywars_kills_team,omitempty"`
			SkywarsKitsSolo                  int `json:"skywars_kits_solo,omitempty"`
			SkywarsKillsSolo                 int `json:"skywars_kills_solo,omitempty"`
			SkywarsCages                     int `json:"skywars_cages,omitempty"`
			WallsCoins                       int `json:"walls_coins,omitempty"`
			TruecombatKitHoarderTeam         int `json:"truecombat_kit_hoarder_team,omitempty"`
			TruecombatKitHoarderSolo         int `json:"truecombat_kit_hoarder_solo,omitempty"`
			TruecombatSoloKiller             int `json:"truecombat_solo_killer,omitempty"`
			UhcChampion                      int `json:"uhc_champion,omitempty"`
			UhcMovingUp                      int `json:"uhc_moving_up,omitempty"`
			UhcHunter                        int `json:"uhc_hunter,omitempty"`
			SupersmashSmashChampion          int `json:"supersmash_smash_champion,omitempty"`
			SupersmashHeroSlayer             int `json:"supersmash_hero_slayer,omitempty"`
			CopsandcrimsSerialKiller         int `json:"copsandcrims_serial_killer,omitempty"`
			CopsandcrimsBombSpecialist       int `json:"copsandcrims_bomb_specialist,omitempty"`
			CopsandcrimsHeroTerrorist        int `json:"copsandcrims_hero_terrorist,omitempty"`
			CopsandcrimsCacBanker            int `json:"copsandcrims_cac_banker,omitempty"`
			GeneralChallenger                int `json:"general_challenger,omitempty"`
			GeneralQuestMaster               int `json:"general_quest_master,omitempty"`
			SkywarsKillsMega                 int `json:"skywars_kills_mega,omitempty"`
			SkywarsWinsMega                  int `json:"skywars_wins_mega,omitempty"`
			SkywarsKitsMega                  int `json:"skywars_kits_mega,omitempty"`
			SkyclashCardsUnlocked            int `json:"skyclash_cards_unlocked,omitempty"`
			SkyclashPacksOpened              int `json:"skyclash_packs_opened,omitempty"`
			SkyclashKills                    int `json:"skyclash_kills,omitempty"`
			WarlordsCtfWins                  int `json:"warlords_ctf_wins,omitempty"`
			WarlordsDomWins                  int `json:"warlords_dom_wins,omitempty"`
			ArcadeMiniwallsWinner            int `json:"arcade_miniwalls_winner,omitempty"`
			ArcadeArcadeWinner               int `json:"arcade_arcade_winner,omitempty"`
			ArcadeZombieKiller               int `json:"arcade_zombie_killer,omitempty"`
			ArcadeTeamWork                   int `json:"arcade_team_work,omitempty"`
			ArcadeFootballPro                int `json:"arcade_football_pro,omitempty"`
			ArcadeBountyHunter               int `json:"arcade_bounty_hunter,omitempty"`
			GingerbreadBanker                int `json:"gingerbread_banker,omitempty"`
			QuakeCoins                       int `json:"quake_coins,omitempty"`
			VampirezZombieKiller             int `json:"vampirez_zombie_killer,omitempty"`
			QuakeHeadshots                   int `json:"quake_headshots,omitempty"`
			BlitzWarVeteran                  int `json:"blitz_war_veteran,omitempty"`
			BlitzWinsTeams                   int `json:"blitz_wins_teams,omitempty"`
			BlitzLooter                      int `json:"blitz_looter,omitempty"`
			BedwarsBeds                      int `json:"bedwars_beds,omitempty"`
			BedwarsWins                      int `json:"bedwars_wins,omitempty"`
			BedwarsLevel                     int `json:"bedwars_level,omitempty"`
			BedwarsLootBox                   int `json:"bedwars_loot_box,omitempty"`
			ArcadeArcadeBanker               int `json:"arcade_arcade_banker,omitempty"`
			MurdermysteryWinsAsSurvivor      int `json:"murdermystery_wins_as_survivor,omitempty"`
			MurdermysteryKillsAsMurderer     int `json:"murdermystery_kills_as_murderer,omitempty"`
			Halloween2017Pumpkinator         int `json:"halloween2017_pumpkinator,omitempty"`
			Christmas2017Advent              int `json:"christmas2017_advent,omitempty"`
			TntgamesTntTagWins               int `json:"tntgames_tnt_tag_wins,omitempty"`
			TntgamesPvpRunKiller             int `json:"tntgames_pvp_run_killer,omitempty"`
			TntgamesPvpRunWins               int `json:"tntgames_pvp_run_wins,omitempty"`
			TntgamesTntWizardsKills          int `json:"tntgames_tnt_wizards_kills,omitempty"`
			Christmas2017NoChristmas         int `json:"christmas2017_no_christmas,omitempty"`
			Christmas2017SantaSaysRounds     int `json:"christmas2017_santa_says_rounds,omitempty"`
			Walls3Guardian                   int `json:"walls3_guardian,omitempty"`
			Walls3Rusher                     int `json:"walls3_rusher,omitempty"`
			ArcadeFarmhuntDominator          int `json:"arcade_farmhunt_dominator,omitempty"`
			TntgamesTntWizardsCaps           int `json:"tntgames_tnt_wizards_caps,omitempty"`
			TntgamesTntTriathlon             int `json:"tntgames_tnt_triathlon,omitempty"`
			TntgamesTntBanker                int `json:"tntgames_tnt_banker,omitempty"`
			DuelsDuelsWinner                 int `json:"duels_duels_winner,omitempty"`
			DuelsDuelsWinStreak              int `json:"duels_duels_win_streak,omitempty"`
			ArcadeZombiesRoundProgression    int `json:"arcade_zombies_round_progression,omitempty"`
			BridgeUniqueMapWins              int `json:"bridge_unique_map_wins,omitempty"`
			BridgeWins                       int `json:"bridge_wins,omitempty"`
			BridgeWinStreak                  int `json:"bridge_win_streak,omitempty"`
			BridgeOneVOneWins                int `json:"bridge_one_v_one_wins,omitempty"`
			BridgeFourVFourWins              int `json:"bridge_four_v_four_wins,omitempty"`
			BedwarsCollectorsEdition         int `json:"bedwars_collectors_edition,omitempty"`
			BedwarsBedwarsKiller             int `json:"bedwars_bedwars_killer,omitempty"`
			BridgeGoals                      int `json:"bridge_goals,omitempty"`
			BridgeTwoVTwoWins                int `json:"bridge_two_v_two_wins,omitempty"`
			Halloween2017PumpkinSmasher      int `json:"halloween2017_pumpkin_smasher,omitempty"`
			DuelsBridgeTeamsWins             int `json:"duels_bridge_teams_wins,omitempty"`
			DuelsUniqueMapWins               int `json:"duels_unique_map_wins,omitempty"`
			DuelsBridgeWinStreak             int `json:"duels_bridge_win_streak,omitempty"`
			DuelsGoals                       int `json:"duels_goals,omitempty"`
			DuelsBridgeWins                  int `json:"duels_bridge_wins,omitempty"`
			DuelsBridgeDoublesWins           int `json:"duels_bridge_doubles_wins,omitempty"`
			DuelsBridgeDuelsWins             int `json:"duels_bridge_duels_wins,omitempty"`
			PaintballKillStreaks             int `json:"paintball_kill_streaks,omitempty"`
			PaintballHatCollector            int `json:"paintball_hat_collector,omitempty"`
			GingerbreadMystery               int `json:"gingerbread_mystery,omitempty"`
			GingerbreadRacer                 int `json:"gingerbread_racer,omitempty"`
			CopsandcrimsHeadshotKills        int `json:"copsandcrims_headshot_kills,omitempty"`
			Christmas2017Advent2018          int `json:"christmas2017_advent_2018,omitempty"`
			ArcadeHideAndSeekHiderKills      int `json:"arcade_hide_and_seek_hider_kills,omitempty"`
			DuelsDuelsTraveller              int `json:"duels_duels_traveller,omitempty"`
			MurdermysteryHoarder             int `json:"murdermystery_hoarder,omitempty"`
			SkyblockTreasury                 int `json:"skyblock_treasury,omitempty"`
			SkyblockExcavator                int `json:"skyblock_excavator,omitempty"`
			SkyblockCombat                   int `json:"skyblock_combat,omitempty"`
			SkyblockGatherer                 int `json:"skyblock_gatherer,omitempty"`
			SkyblockHarvester                int `json:"skyblock_harvester,omitempty"`
			SkyblockAngler                   int `json:"skyblock_angler,omitempty"`
			SkyblockAugmentation             int `json:"skyblock_augmentation,omitempty"`
			SkyblockConcoctor                int `json:"skyblock_concoctor,omitempty"`
			SkyblockMinionLover              int `json:"skyblock_minion_lover,omitempty"`
			ArcadeZombiesNiceShot            int `json:"arcade_zombies_nice_shot,omitempty"`
			CopsandcrimsGrenadeKills         int `json:"copsandcrims_grenade_kills,omitempty"`
			Christmas2017Advent2019          int `json:"christmas2017_advent_2019,omitempty"`
			SkyblockUniqueGifts              int `json:"skyblock_unique_gifts,omitempty"`
			Christmas2017SecretSanta         int `json:"christmas2017_secret_santa,omitempty"`
			EasterThrowEggs                  int `json:"easter_throw_eggs,omitempty"`
			EasterEggFinder                  int `json:"easter_egg_finder,omitempty"`
			EasterMasterTracker              int `json:"easter_master_tracker,omitempty"`
			SkyblockDomesticator             int `json:"skyblock_domesticator,omitempty"`
			SkyblockSlayer                   int `json:"skyblock_slayer,omitempty"`
			SummerTreasureHoarder            int `json:"summer_treasure_hoarder,omitempty"`
			SummerShopaholic                 int `json:"summer_shopaholic,omitempty"`
			CopsandcrimsPracticeMakesPerfect int `json:"copsandcrims_practice_makes_perfect,omitempty"`
			BuildbattleGuessTheBuildGuesses  int `json:"buildbattle_guess_the_build_guesses,omitempty"`
			BuildbattleBuildBattleScore      int `json:"buildbattle_build_battle_score,omitempty"`
			ArcadeCtwSlayer                  int `json:"arcade_ctw_slayer,omitempty"`
			ArcadeCtwOhSheep                 int `json:"arcade_ctw_oh_sheep,omitempty"`
			Halloween2017CandyHoarder        int `json:"halloween2017_candy_hoarder,omitempty"`
			Halloween2017GetThemApples       int `json:"halloween2017_get_them_apples,omitempty"`
			SkywarsYouReAStar                int `json:"skywars_you_re_a_star,omitempty"`
			TntgamesBlockRunner              int `json:"tntgames_block_runner,omitempty"`
			Christmas2017Advent2020          int `json:"christmas2017_advent_2020,omitempty"`
			Christmas2017BestPresents        int `json:"christmas2017_best_presents,omitempty"`
			QuakeGodlikes                    int `json:"quake_godlikes,omitempty"`
			TntgamesClinic                   int `json:"tntgames_clinic,omitempty"`
			SkyblockDungeoneer               int `json:"skyblock_dungeoneer,omitempty"`
			SkyblockHardWorkingMiner         int `json:"skyblock_hard_working_miner,omitempty"`
			SkyblockGoblinKiller             int `json:"skyblock_goblin_killer,omitempty"`
			ArcadeDwSlayer                   int `json:"arcade_dw_slayer,omitempty"`
			ArcadeGalaxyWarsKills            int `json:"arcade_galaxy_wars_kills,omitempty"`
			ArcadeThrowOutKills              int `json:"arcade_throw_out_kills,omitempty"`
			MurdermysteryCountermeasures     int `json:"murdermystery_countermeasures,omitempty"`
			MurdermysteryHitman              int `json:"murdermystery_hitman,omitempty"`
			MurdermysterySurvivalSkills      int `json:"murdermystery_survival_skills,omitempty"`
			BuildbattleBuildBattleVoter      int `json:"buildbattle_build_battle_voter,omitempty"`
			BuildbattleBuildBattlePoints     int `json:"buildbattle_build_battle_points,omitempty"`
			Christmas2017Advent2021          int `json:"christmas2017_advent_2021,omitempty"`
			GeneralMasterLure                int `json:"general_master_lure,omitempty"`
			GeneralLuckiestOfTheSea          int `json:"general_luckiest_of_the_sea,omitempty"`
			GeneralTrashiestDiver            int `json:"general_trashiest_diver,omitempty"`
			SkyblockTreasureHunter           int `json:"skyblock_treasure_hunter,omitempty"`
			SummerGoneFishing                int `json:"summer_gone_fishing,omitempty"`
			ArcadePartySuperStar             int `json:"arcade_party_super_star,omitempty"`
			ArcadeHitwPracticeMakesPerfect   int `json:"arcade_hitw_practice_makes_perfect,omitempty"`
			ArcadeEnderSpleefBlockStealer    int `json:"arcade_ender_spleef_block_stealer,omitempty"`
			SkywarsHeads                     int `json:"skywars_heads,omitempty"`
			ArcadePixelPartyColorCoordinated int `json:"arcade_pixel_party_color_coordinated,omitempty"`
			Christmas2017Advent2022          int `json:"christmas2017_advent_2022,omitempty"`
			SkyblockCurator                  int `json:"skyblock_curator,omitempty"`
			SkyblockSbLevels                 int `json:"skyblock_sb_levels,omitempty"`
			PitGold                          int `json:"pit_gold,omitempty"`
			PitKills                         int `json:"pit_kills,omitempty"`
			WoolgamesWoolKills               int `json:"woolgames_wool_kills,omitempty"`
			WoolgamesWoolContest             int `json:"woolgames_wool_contest,omitempty"`
			ArenaPowerup                     int `json:"arena_powerup,omitempty"`
			WarlordsCoins                    int `json:"warlords_coins,omitempty"`
			SkyblockPeoplePleaser            int `json:"skyblock_people_pleaser,omitempty"`
			WoolgamesWoolWins                int `json:"woolgames_wool_wins,omitempty"`
			Christmas2017Advent2023          int `json:"christmas2017_advent_2023,omitempty"`
		} `json:"achievements,omitempty"`
		AchievementsOneTime    []string `json:"achievementsOneTime,omitempty"`
		Channel                string   `json:"channel,omitempty"`
		Displayname            string   `json:"displayname,omitempty"`
		FirstLogin             int64    `json:"firstLogin,omitempty"`
		FriendRequests         []any    `json:"friendRequests,omitempty"`
		Karma                  int      `json:"karma,omitempty"`
		LastLogin              int64    `json:"lastLogin,omitempty"`
		MostRecentlyThanked    string   `json:"mostRecentlyThanked,omitempty"`
		MostRecentlyTipped     string   `json:"mostRecentlyTipped,omitempty"`
		MostRecentlyTippedUUID string   `json:"mostRecentlyTippedUuid,omitempty"`
		NetworkExp             int      `json:"networkExp,omitempty"`
		ParkourCompletions     struct {
			Main []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"main,omitempty"`
			CopsnCrims []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"CopsnCrims,omitempty"`
			TheWallsLobby []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"TheWallsLobby,omitempty"`
			TruePVPParkour []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"TruePVPParkour,omitempty"`
			QuakeCraft []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"QuakeCraft,omitempty"`
			NewMainLobby []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"NewMainLobby,omitempty"`
			Skywars []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Skywars,omitempty"`
			SuperSmash []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"SuperSmash,omitempty"`
			Tnt []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"TNT,omitempty"`
			Arena []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Arena,omitempty"`
			Turbo []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Turbo,omitempty"`
			Prototype []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Prototype,omitempty"`
			Bedwars []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Bedwars,omitempty"`
			SkywarsAug2017 []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"SkywarsAug2017,omitempty"`
			MurderMystery []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"MurderMystery,omitempty"`
			Legacy []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Legacy,omitempty"`
			Tourney []struct {
				TimeStart int64 `json:"timeStart,omitempty"`
				TimeTook  int   `json:"timeTook,omitempty"`
			} `json:"Tourney,omitempty"`
		} `json:"parkourCompletions,omitempty"`
		Playername string `json:"playername,omitempty"`
		Quests     struct {
			Halloween2014 struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"halloween2014,omitempty"`
			PaintballExpert struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						Kill int `json:"kill,omitempty"`
						Play int `json:"play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"paintball_expert,omitempty"`
			SpaceMission struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"space_mission,omitempty"`
			WarlordsWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_win,omitempty"`
			TntAddict struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_addict,omitempty"`
			Waller struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Kill int `json:"kill,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"waller,omitempty"`
			SkywarsWeeklyKills struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SkywarsWeeklyKills int `json:"skywars_weekly_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_weekly_kills,omitempty"`
			ExplosiveGames struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"explosive_games,omitempty"`
			SerialKiller struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Quake     int `json:"quake,omitempty"`
						Tnt       int `json:"tnt,omitempty"`
						Blitz     int `json:"blitz,omitempty"`
						Paintball int `json:"paintball,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"serial_killer,omitempty"`
			NuggetWarriors struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Quake     int `json:"quake,omitempty"`
						Paintball int `json:"paintball,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"nugget_warriors,omitempty"`
			WelcomeToHell struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Megawalls int `json:"megawalls,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"welcome_to_hell,omitempty"`
			WarriorsJourney struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Blitzkill         int `json:"blitzkill,omitempty"`
						Quake25Kill       int `json:"quake25kill,omitempty"`
						Tntwin            int `json:"tntwin,omitempty"`
						Paintballwin      int `json:"paintballwin,omitempty"`
						Vampirezkillvamps int `json:"vampirezkillvamps,omitempty"`
						Megawallswin      int `json:"megawallswin,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warriors_journey,omitempty"`
			Blitzerk struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Killblitz10 int `json:"killblitz10,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitzerk,omitempty"`
			Megawaller struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						Win int `json:"win,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"megawaller,omitempty"`
			Gladiator struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"gladiator,omitempty"`
			QuakeDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						QuakeDailyPlay int `json:"quake_daily_play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"quake_daily_play,omitempty"`
			QuakeDailyKill struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						QuakeDailyKill int `json:"quake_daily_kill,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"quake_daily_kill,omitempty"`
			QuakeWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"quake_weekly_play,omitempty"`
			SkywarsTeamKills struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SkywarsTeamKills int `json:"skywars_team_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_team_kills,omitempty"`
			SkywarsTeamWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_team_win,omitempty"`
			SkywarsSoloKills struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SkywarsSoloKills int `json:"skywars_solo_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_solo_kills,omitempty"`
			SkywarsSoloWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_solo_win,omitempty"`
			CvcWinDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"cvc_win_daily_normal,omitempty"`
			CvcKillDailyNormal struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						CvcKillDailyNormal int `json:"cvc_kill_daily_normal,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"cvc_kill_daily_normal,omitempty"`
			CvcKill struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"cvc_kill,omitempty"`
			CvcWinDailyDeathmatch struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"cvc_win_daily_deathmatch,omitempty"`
			CvcKillWeekly struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						CvcPlayWeekly2 int `json:"cvc_play_weekly_2,omitempty"`
						CvcPlayWeekly  int `json:"cvc_play_weekly,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"cvc_kill_weekly,omitempty"`
			ArcadeGamer struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arcade_gamer,omitempty"`
			ArcadeWinner struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arcade_winner,omitempty"`
			ArcadeSpecialist struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						Play int `json:"play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arcade_specialist,omitempty"`
			SupersmashSoloWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"supersmash_solo_win,omitempty"`
			SupersmashSoloKills struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SupersmashSoloKills int `json:"supersmash_solo_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"supersmash_solo_kills,omitempty"`
			SupersmashTeamWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"supersmash_team_win,omitempty"`
			SupersmashTeamKills struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SupersmashTeamKills int `json:"supersmash_team_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"supersmash_team_kills,omitempty"`
			SupersmashWeeklyKills struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						SupersmashWeeklyKills int `json:"supersmash_weekly_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"supersmash_weekly_kills,omitempty"`
			QuakeDailyWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"quake_daily_win,omitempty"`
			BedwarsDailyWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_win,omitempty"`
			BedwarsDailyOneMore struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_one_more,omitempty"`
			BedwarsWeeklyBedElims struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						BedwarsBedElims int `json:"bedwars_bed_elims,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_bed_elims,omitempty"`
			Paintballer struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"paintballer,omitempty"`
			PaintballKiller struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"paintball_killer,omitempty"`
			WallsDailyPlay struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"walls_daily_play,omitempty"`
			WallsDailyKill struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						WallsDailyKill int `json:"walls_daily_kill,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"walls_daily_kill,omitempty"`
			WallsDailyWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"walls_daily_win,omitempty"`
			WallsWeekly struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						WallsWeeklyKills int `json:"walls_weekly_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"walls_weekly,omitempty"`
			MmDailyWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_daily_win,omitempty"`
			MmDailyPowerPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_daily_power_play,omitempty"`
			MmDailyTargetKill struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						MmTargetKills int `json:"mm_target_kills,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_daily_target_kill,omitempty"`
			MmWeeklyMurdererKills struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						MmWeeklyKillsAsMurderer int `json:"mm_weekly_kills_as_murderer,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_weekly_murderer_kills,omitempty"`
			MmWeeklyWins struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						MmWeeklyWin int `json:"mm_weekly_win,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_weekly_wins,omitempty"`
			BedwarsWeeklyPumpkinator struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						BedwarsSpecialWeeklyPumpkinator int `json:"bedwars_special_weekly_pumpkinator,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_pumpkinator,omitempty"`
			BedwarsWeeklySanta struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"bedwars_weekly_santa,omitempty"`
			MmSpecialWeeklySanta struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						MmSpecialWeeklySanta int `json:"mm_special_weekly_santa,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_special_weekly_santa,omitempty"`
			MegaWallsWin struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mega_walls_win,omitempty"`
			MegaWallsPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mega_walls_play,omitempty"`
			MegaWallsKill struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						MegaWallsKillDaily int `json:"mega_walls_kill_daily,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mega_walls_kill,omitempty"`
			MegaWallsWeekly struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						MegaWallsKillWeekly int `json:"mega_walls_kill_weekly,omitempty"`
						MegaWallsPlayWeekly int `json:"mega_walls_play_weekly,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mega_walls_weekly,omitempty"`
			TntTnttagDaily struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_tnttag_daily,omitempty"`
			TntTnttagWeekly struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						TntTnttagWeekly int `json:"tnt_tnttag_weekly,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_tnttag_weekly,omitempty"`
			TntBowspleefDaily struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"tnt_bowspleef_daily,omitempty"`
			TntBowspleefWeekly struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						TntBowspleefWeekly int `json:"tnt_bowspleef_weekly,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_bowspleef_weekly,omitempty"`
			TntPvprunDaily struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						TntPvprunDaily int `json:"tnt_pvprun_daily,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_pvprun_daily,omitempty"`
			TntPvprunWeekly struct {
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						TntPvprunWeekly int `json:"tnt_pvprun_weekly,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_pvprun_weekly,omitempty"`
			TntTntrunDaily struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"tnt_tntrun_daily,omitempty"`
			TntTntrunWeekly struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
			} `json:"tnt_tntrun_weekly,omitempty"`
			TntWizardsDaily struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						TntWizardsDailyKills int `json:"tnt_wizards_daily_kills,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_wizards_daily,omitempty"`
			TntWizardsWeekly struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						TntWizardsWeeklyKills int `json:"tnt_wizards_weekly_kills,omitempty"`
						TntWizardsWeeklyWin   int `json:"tnt_wizards_weekly_win,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_wizards_weekly,omitempty"`
			BedwarsWeeklyDreamWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						BedwarsDreamWins int `json:"bedwars_dream_wins,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_dream_win,omitempty"`
			BedwarsDailyNightmares struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
						BedwarsDailyNightmareBeds int `json:"bedwars_daily_nightmare_beds,omitempty"`
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_nightmares,omitempty"`
			TntDailyWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_daily_win,omitempty"`
			TntWeeklyPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						TntWeeklyPlay int `json:"tnt_weekly_play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_weekly_play,omitempty"`
			BedwarsDailyGifts struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Started    int64 `json:"started,omitempty"`
					Objectives struct {
					} `json:"objectives,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_gifts,omitempty"`
			GingerbreadMastery struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"gingerbread_mastery,omitempty"`
			GingerbreadMaps struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"gingerbread_maps,omitempty"`
			GingerbreadBlingBling struct {
				Active struct {
					Objectives struct {
						GingerbreadGoldPickedup int `json:"gingerbread_gold_pickedup,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"gingerbread_bling_bling,omitempty"`
			CrazyWallsDailyPlay struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"crazy_walls_daily_play,omitempty"`
			CrazyWallsDailyKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"crazy_walls_daily_kill,omitempty"`
			CrazyWallsDailyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"crazy_walls_daily_win,omitempty"`
			CrazyWallsWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"crazy_walls_weekly,omitempty"`
			TntWeeklySpecial struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						TntWeeklyCandyCanes int `json:"tnt_weekly_candy_canes,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"tnt_weekly_special,omitempty"`
			VampirezDailyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_daily_win,omitempty"`
			VampirezDailyKill struct {
				Active struct {
					Objectives struct {
						VampirezDailyKillVampire int `json:"vampirez_daily_kill_vampire,omitempty"`
						VampirezDailyKillZombie  int `json:"vampirez_daily_kill_zombie,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_daily_kill,omitempty"`
			VampirezDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_daily_play,omitempty"`
			ArenaWeeklyPlay struct {
				Active struct {
					Objectives struct {
						ArenaWeeklyPlay int `json:"arena_weekly_play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arena_weekly_play,omitempty"`
			VampirezWeeklyKill struct {
				Active struct {
					Objectives struct {
						VampirezWeeklyKillVampire int `json:"vampirez_weekly_kill_vampire,omitempty"`
						VampirezWeeklyKillZombie  int `json:"vampirez_weekly_kill_zombie,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_weekly_kill,omitempty"`
			ArenaDailyKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arena_daily_kills,omitempty"`
			ArenaDailyPlay struct {
				Active struct {
					Objectives struct {
						ArenaDailyPlay int `json:"arena_daily_play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arena_daily_play,omitempty"`
			GingerbreadRacer struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"gingerbread_racer,omitempty"`
			VampirezWeeklyWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_weekly_win,omitempty"`
			ArenaDailyWins struct {
				Active struct {
					Objectives struct {
						ArenaDailyWins int `json:"arena_daily_wins,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"arena_daily_wins,omitempty"`
			VampirezDailyHumanKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_daily_human_kill,omitempty"`
			VampirezWeeklyHumanKill struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"vampirez_weekly_human_kill,omitempty"`
			SkywarsArcadeWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_arcade_win,omitempty"`
			SkywarsSpecialNorthPole struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_special_north_pole,omitempty"`
			SkywarsWeeklyArcadeWinAll struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_weekly_arcade_win_all,omitempty"`
			SkywarsWeeklyFreeLootChest struct {
				Active struct {
					Objectives struct {
						SkywarsFreeLootChestWin int `json:"skywars_free_loot_chest_win,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_weekly_free_loot_chest,omitempty"`
			SkywarsCorruptWin struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_corrupt_win,omitempty"`
			SkywarsDailyTokens struct {
				Active struct {
					Objectives struct {
						SkywarsDailyTokensWins int `json:"skywars_daily_tokens_wins,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"skywars_daily_tokens,omitempty"`
			BlitzGameOfTheDay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_game_of_the_day,omitempty"`
			BlitzSpecialDailyNorthPole struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_special_daily_north_pole,omitempty"`
			BlitzKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_kills,omitempty"`
			BlitzWeeklyMaster struct {
				Active struct {
					Objectives struct {
						BlitzGamesPlayed int `json:"blitz_games_played,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_weekly_master,omitempty"`
			BlitzLootChestWeekly struct {
				Active struct {
					Objectives struct {
						Lootchestblitz  int `json:"lootchestblitz,omitempty"`
						Dealdamageblitz int `json:"dealdamageblitz,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_loot_chest_weekly,omitempty"`
			BlitzLootChestDaily struct {
				Active struct {
					Objectives struct {
						Lootchestblitz int `json:"lootchestblitz,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_loot_chest_daily,omitempty"`
			BlitzWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"blitz_win,omitempty"`
			DuelsWinner struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"duels_winner,omitempty"`
			DuelsWeeklyKills struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"duels_weekly_kills,omitempty"`
			DuelsPlayer struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"duels_player,omitempty"`
			DuelsWeeklyWins struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"duels_weekly_wins,omitempty"`
			DuelsKiller struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"duels_killer,omitempty"`
			MegaWallsFaithful struct {
				Active struct {
					Objectives struct {
						MegaWallsFaithfulPlay int `json:"mega_walls_faithful_play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mega_walls_faithful,omitempty"`
			WarlordsTdm struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_tdm,omitempty"`
			WarlordsCtf struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_ctf,omitempty"`
			WarlordsDomination struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_domination,omitempty"`
			WarlordsVictorious struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_victorious,omitempty"`
			WarlordsDedication struct {
				Active struct {
					Objectives struct {
						WarlordsWeeklyDedi int `json:"warlords_weekly_dedi,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_dedication,omitempty"`
			WarlordsAllStar struct {
				Active struct {
					Objectives struct {
						WarlordsWeeklyHeal   int `json:"warlords_weekly_heal,omitempty"`
						WarlordsWeeklyDamage int `json:"warlords_weekly_damage,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_all_star,omitempty"`
			WarlordsObjectives struct {
				Active struct {
					Objectives struct {
						WarlordsDailyObjectives int `json:"warlords_daily_objectives,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"warlords_objectives,omitempty"`
			BuildBattlePlayer struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"build_battle_player,omitempty"`
			BuildBattleWinner struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"build_battle_winner,omitempty"`
			BuildBattleWeekly struct {
				Active struct {
					Objectives struct {
						Play int `json:"play,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"build_battle_weekly,omitempty"`
			BuildBattleChristmasWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"build_battle_christmas_weekly,omitempty"`
			BuildBattleChristmas struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"build_battle_christmas,omitempty"`
			UhcSolo struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_solo,omitempty"`
			UhcTeam struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_team,omitempty"`
			UhcWeekly struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_weekly,omitempty"`
			UhcMadness struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_madness,omitempty"`
			TeamBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"team_brawler,omitempty"`
			UhcDm struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_dm,omitempty"`
			SoloBrawler struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"solo_brawler,omitempty"`
			UhcWeeklySpecialCookie struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"uhc_weekly_special_cookie,omitempty"`
			MmDailyInfector struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"mm_daily_infector,omitempty"`
			BedwarsWeeklyChallenges struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_challenges,omitempty"`
			WoolWarsDailyPlay struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"wool_wars_daily_play,omitempty"`
			WoolWarsDailyWins struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						Win int `json:"win,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"wool_wars_daily_wins,omitempty"`
			WoolWarsDailyKills struct {
				Completions []struct {
					Time int64 `json:"time,omitempty"`
				} `json:"completions,omitempty"`
				Active struct {
					Objectives struct {
						Kill int `json:"kill,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"wool_wars_daily_kills,omitempty"`
			WoolWeeklyPlay struct {
				Active struct {
					Objectives struct {
						Kill int `json:"kill,omitempty"`
						Win  int `json:"win,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"wool_weekly_play,omitempty"`
			WoolWarsWeeklyShears struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"wool_wars_weekly_shears,omitempty"`
			PitDailyKills struct {
				Active struct {
					Objectives struct {
						Kill int `json:"kill,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"pit_daily_kills,omitempty"`
			PitDailyContract struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"pit_daily_contract,omitempty"`
			PitWeeklyGold struct {
				Active struct {
					Objectives struct {
						PitWeeklyGold int `json:"pit_weekly_gold,omitempty"`
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"pit_weekly_gold,omitempty"`
			BedwarsWeeklyChallengesWin struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_challenges_win,omitempty"`
			BedwarsDailyBedBreaker struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_bed_breaker,omitempty"`
			BedwarsDailyFinalKiller struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_daily_final_killer,omitempty"`
			BedwarsWeeklyFinalKiller struct {
				Active struct {
					Objectives struct {
					} `json:"objectives,omitempty"`
					Started int64 `json:"started,omitempty"`
				} `json:"active,omitempty"`
			} `json:"bedwars_weekly_final_killer,omitempty"`
		} `json:"quests,omitempty"`
		Stats struct {
			Arcade struct {
				Blood                            bool     `json:"blood,omitempty"`
				BountyKillsOneinthequiver        int      `json:"bounty_kills_oneinthequiver,omitempty"`
				Coins                            int      `json:"coins,omitempty"`
				DeathsOneinthequiver             int      `json:"deaths_oneinthequiver,omitempty"`
				HeadshotsDayone                  int      `json:"headshots_dayone,omitempty"`
				KillsDayone                      int      `json:"kills_dayone,omitempty"`
				KillsDragonwars2                 int      `json:"kills_dragonwars2,omitempty"`
				KillsOneinthequiver              int      `json:"kills_oneinthequiver,omitempty"`
				MaxWave                          int      `json:"max_wave,omitempty"`
				PoopCollected                    int      `json:"poop_collected,omitempty"`
				StampLevel                       int      `json:"stamp_level,omitempty"`
				TimeStamp                        int64    `json:"time_stamp,omitempty"`
				SwKills                          int      `json:"sw_kills,omitempty"`
				SwShotsFired                     int      `json:"sw_shots_fired,omitempty"`
				SwEmpireKills                    int      `json:"sw_empire_kills,omitempty"`
				SwWeeklyKillsA                   int      `json:"sw_weekly_kills_a,omitempty"`
				SwDeaths                         int      `json:"sw_deaths,omitempty"`
				SwMonthlyKillsA                  int      `json:"sw_monthly_kills_a,omitempty"`
				SwRebelKills                     int      `json:"sw_rebel_kills,omitempty"`
				SwWeeklyKillsB                   int      `json:"sw_weekly_kills_b,omitempty"`
				Music                            bool     `json:"music,omitempty"`
				SwMonthlyKillsB                  int      `json:"sw_monthly_kills_b,omitempty"`
				SwGameWins                       int      `json:"sw_game_wins,omitempty"`
				RoundsHoleInTheWall              int      `json:"rounds_hole_in_the_wall,omitempty"`
				DttDropdown                      bool     `json:"dtt_dropdown,omitempty"`
				WinsDayone                       int      `json:"wins_dayone,omitempty"`
				WinsParty3                       int      `json:"wins_party_3,omitempty"`
				RoundsSantaSays                  int      `json:"rounds_santa_says,omitempty"`
				WinsSantaSays                    int      `json:"wins_santa_says,omitempty"`
				PowerkicksSoccer                 int      `json:"powerkicks_soccer,omitempty"`
				GoalsSoccer                      int      `json:"goals_soccer,omitempty"`
				MonthlyCoinsB                    int      `json:"monthly_coins_b,omitempty"`
				WeeklyCoinsA                     int      `json:"weekly_coins_a,omitempty"`
				WeeklyCoinsB                     int      `json:"weekly_coins_b,omitempty"`
				MonthlyCoinsA                    int      `json:"monthly_coins_a,omitempty"`
				ArrowsHitMiniWalls               int      `json:"arrows_hit_mini_walls,omitempty"`
				FinalKillsMiniWalls              int      `json:"final_kills_mini_walls,omitempty"`
				KillsMiniWalls                   int      `json:"kills_mini_walls,omitempty"`
				DeathsMiniWalls                  int      `json:"deaths_mini_walls,omitempty"`
				ArrowsShotMiniWalls              int      `json:"arrows_shot_mini_walls,omitempty"`
				WitherDamageMiniWalls            int      `json:"wither_damage_mini_walls,omitempty"`
				KillsThrowOut                    int      `json:"kills_throw_out,omitempty"`
				DeathsThrowOut                   int      `json:"deaths_throw_out,omitempty"`
				Dec2016Achievements2             bool     `json:"dec2016_achievements2,omitempty"`
				Dec2016Achievements              bool     `json:"dec2016_achievements,omitempty"`
				HitwRecordQ                      int      `json:"hitw_record_q,omitempty"`
				HitwRecordF                      int      `json:"hitw_record_f,omitempty"`
				WinsHoleInTheWall                int      `json:"wins_hole_in_the_wall,omitempty"`
				RoundsSimonSays                  int      `json:"rounds_simon_says,omitempty"`
				WinsSimonSays                    int      `json:"wins_simon_says,omitempty"`
				WinsFarmHunt                     int      `json:"wins_farm_hunt,omitempty"`
				WinsEnder                        int      `json:"wins_ender,omitempty"`
				BestRoundZombies                 int      `json:"best_round_zombies,omitempty"`
				DeathsZombies                    int      `json:"deaths_zombies,omitempty"`
				TimesKnockedDownZombies          int      `json:"times_knocked_down_zombies,omitempty"`
				ZombieKillsZombies               int      `json:"zombie_kills_zombies,omitempty"`
				PartyPooperHiderWinsHideAndSeek  int      `json:"party_pooper_hider_wins_hide_and_seek,omitempty"`
				HiderWinsHideAndSeek             int      `json:"hider_wins_hide_and_seek,omitempty"`
				SeekerWinsHideAndSeek            int      `json:"seeker_wins_hide_and_seek,omitempty"`
				PartyPooperSeekerWinsHideAndSeek int      `json:"party_pooper_seeker_wins_hide_and_seek,omitempty"`
				Showinfobook                     bool     `json:"showinfobook,omitempty"`
				DeliveredSantaSimulator          int      `json:"delivered_santa_simulator,omitempty"`
				SpottedSantaSimulator            int      `json:"spotted_santa_simulator,omitempty"`
				EggsFoundEasterSimulator         int      `json:"eggs_found_easter_simulator,omitempty"`
				WinsEasterSimulator              int      `json:"wins_easter_simulator,omitempty"`
				ItemsFoundScubaSimulator         int      `json:"items_found_scuba_simulator,omitempty"`
				TotalPointsScubaSimulator        int      `json:"total_points_scuba_simulator,omitempty"`
				CandyFoundHalloweenSimulator     int      `json:"candy_found_halloween_simulator,omitempty"`
				WinsGrinchSimulatorV2            int      `json:"wins_grinch_simulator_v2,omitempty"`
				GiftsGrinchSimulatorV2           int      `json:"gifts_grinch_simulator_v2,omitempty"`
				LastTourneyAd                    int64    `json:"lastTourneyAd,omitempty"`
				GiftsGrinchSimulatorV2Tourney    int      `json:"gifts_grinch_simulator_v2_tourney,omitempty"`
				WinsGrinchSimulatorV2Tourney     int      `json:"wins_grinch_simulator_v2_tourney,omitempty"`
				LossesGrinchSimulatorV2Tourney   int      `json:"losses_grinch_simulator_v2_tourney,omitempty"`
				WinsParty                        int      `json:"wins_party,omitempty"`
				AnimalWinsFarmHunt               int      `json:"animal_wins_farm_hunt,omitempty"`
				DangerousTauntsUsedFarmHunt      int      `json:"dangerous_taunts_used_farm_hunt,omitempty"`
				FireworkTauntsUsedFarmHunt       int      `json:"firework_taunts_used_farm_hunt,omitempty"`
				RiskyTauntsUsedFarmHunt          int      `json:"risky_taunts_used_farm_hunt,omitempty"`
				SafeTauntsUsedFarmHunt           int      `json:"safe_taunts_used_farm_hunt,omitempty"`
				TauntsUsedFarmHunt               int      `json:"taunts_used_farm_hunt,omitempty"`
				BowKillsFarmHunt                 int      `json:"bow_kills_farm_hunt,omitempty"`
				HunterBowKillsFarmHunt           int      `json:"hunter_bow_kills_farm_hunt,omitempty"`
				HunterKillsFarmHunt              int      `json:"hunter_kills_farm_hunt,omitempty"`
				HuntersBowKillsFarmHunt          int      `json:"hunters_bow_kills_farm_hunt,omitempty"`
				KillsFarmHunt                    int      `json:"kills_farm_hunt,omitempty"`
				PoopCollectedFarmHunt            int      `json:"poop_collected_farm_hunt,omitempty"`
				RoundWinsSantaSays               int      `json:"round_wins_santa_says,omitempty"`
				TopScoreSantaSays                int      `json:"top_score_santa_says,omitempty"`
				HunterWinsFarmHunt               int      `json:"hunter_wins_farm_hunt,omitempty"`
				AnimalBowKillsFarmHunt           int      `json:"animal_bow_kills_farm_hunt,omitempty"`
				AnimalKillsFarmHunt              int      `json:"animal_kills_farm_hunt,omitempty"`
				AnimalsBowKillsFarmHunt          int      `json:"animals_bow_kills_farm_hunt,omitempty"`
				ActiveVictoryDance               string   `json:"active_victory_dance,omitempty"`
				Packages                         []string `json:"packages,omitempty"`
				HoeHoeHoeBestScoreParty          int      `json:"hoe_hoe_hoe_best_score_party,omitempty"`
				HoeHoeHoeTotalScoreParty         int      `json:"hoe_hoe_hoe_total_score_party,omitempty"`
				JungleJumpBestTimeParty          int      `json:"jungle_jump_best_time_party,omitempty"`
				LawnMoowerMowedBestScoreParty    int      `json:"lawn_moower_mowed_best_score_party,omitempty"`
				LawnMoowerMowedTotalScoreParty   int      `json:"lawn_moower_mowed_total_score_party,omitempty"`
				Rpg16KillsBestScoreParty         int      `json:"rpg_16_kills_best_score_party,omitempty"`
				Rpg16KillsParty                  int      `json:"rpg_16_kills_party,omitempty"`
				TotalStarsParty                  int      `json:"total_stars_party,omitempty"`
				BlocksDestroyedEnder             int      `json:"blocks_destroyed_ender,omitempty"`
				PixelParty                       struct {
					GamesPlayed           int `json:"games_played,omitempty"`
					GamesPlayedNormal     int `json:"games_played_normal,omitempty"`
					HighestRound          int `json:"highest_round,omitempty"`
					RoundsCompleted       int `json:"rounds_completed,omitempty"`
					RoundsCompletedNormal int `json:"rounds_completed_normal,omitempty"`
				} `json:"pixel_party,omitempty"`
				BasicZombieKillsZombies                 int `json:"basic_zombie_kills_zombies,omitempty"`
				BestRoundZombiesDeadend                 int `json:"best_round_zombies_deadend,omitempty"`
				BestRoundZombiesDeadendNormal           int `json:"best_round_zombies_deadend_normal,omitempty"`
				BulletsHitZombies                       int `json:"bullets_hit_zombies,omitempty"`
				BulletsShotZombies                      int `json:"bullets_shot_zombies,omitempty"`
				DeathsZombiesDeadend                    int `json:"deaths_zombies_deadend,omitempty"`
				DeathsZombiesDeadendNormal              int `json:"deaths_zombies_deadend_normal,omitempty"`
				DoorsOpenedZombies                      int `json:"doors_opened_zombies,omitempty"`
				DoorsOpenedZombiesDeadend               int `json:"doors_opened_zombies_deadend,omitempty"`
				DoorsOpenedZombiesDeadendNormal         int `json:"doors_opened_zombies_deadend_normal,omitempty"`
				EmpoweredZombieKillsZombies             int `json:"empowered_zombie_kills_zombies,omitempty"`
				FastestTime10Zombies                    int `json:"fastest_time_10_zombies,omitempty"`
				FastestTime10ZombiesDeadendNormal       int `json:"fastest_time_10_zombies_deadend_normal,omitempty"`
				FireZombieKillsZombies                  int `json:"fire_zombie_kills_zombies,omitempty"`
				HeadshotsZombies                        int `json:"headshots_zombies,omitempty"`
				MagmaCubeZombieKillsZombies             int `json:"magma_cube_zombie_kills_zombies,omitempty"`
				MagmaZombieKillsZombies                 int `json:"magma_zombie_kills_zombies,omitempty"`
				PigZombieZombieKillsZombies             int `json:"pig_zombie_zombie_kills_zombies,omitempty"`
				TimesKnockedDownZombiesDeadend          int `json:"times_knocked_down_zombies_deadend,omitempty"`
				TimesKnockedDownZombiesDeadendNormal    int `json:"times_knocked_down_zombies_deadend_normal,omitempty"`
				TntBabyZombieKillsZombies               int `json:"tnt_baby_zombie_kills_zombies,omitempty"`
				TotalRoundsSurvivedZombies              int `json:"total_rounds_survived_zombies,omitempty"`
				TotalRoundsSurvivedZombiesDeadend       int `json:"total_rounds_survived_zombies_deadend,omitempty"`
				TotalRoundsSurvivedZombiesDeadendNormal int `json:"total_rounds_survived_zombies_deadend_normal,omitempty"`
				WindowsRepairedZombies                  int `json:"windows_repaired_zombies,omitempty"`
				WindowsRepairedZombiesDeadend           int `json:"windows_repaired_zombies_deadend,omitempty"`
				WindowsRepairedZombiesDeadendNormal     int `json:"windows_repaired_zombies_deadend_normal,omitempty"`
				WolfZombieKillsZombies                  int `json:"wolf_zombie_kills_zombies,omitempty"`
				ZombieKillsZombiesDeadend               int `json:"zombie_kills_zombies_deadend,omitempty"`
				ZombieKillsZombiesDeadendNormal         int `json:"zombie_kills_zombies_deadend_normal,omitempty"`
				LeaderboardSettings                     struct {
					ResetType string `json:"resetType,omitempty"`
				} `json:"leaderboardSettings,omitempty"`
				DiveBestScoreParty          int    `json:"dive_best_score_party,omitempty"`
				DiveTotalScoreParty         int    `json:"dive_total_score_party,omitempty"`
				LabEscapeBestTimeParty      int    `json:"lab_escape_best_time_party,omitempty"`
				OptionShowTutorialBook      string `json:"option_show_tutorial_book,omitempty"`
				WoolhuntAssists             int    `json:"woolhunt_assists,omitempty"`
				WoolhuntDeaths              int    `json:"woolhunt_deaths,omitempty"`
				WoolhuntDeathsToWoolholder  int    `json:"woolhunt_deaths_to_woolholder,omitempty"`
				WoolhuntDeathsWithWool      int    `json:"woolhunt_deaths_with_wool,omitempty"`
				WoolhuntExperiencedLosses   int    `json:"woolhunt_experienced_losses,omitempty"`
				WoolhuntExperiencedWins     int    `json:"woolhunt_experienced_wins,omitempty"`
				WoolhuntFastestWoolCapture  int    `json:"woolhunt_fastest_wool_capture,omitempty"`
				WoolhuntGoldEarned          int    `json:"woolhunt_gold_earned,omitempty"`
				WoolhuntGoldSpent           int    `json:"woolhunt_gold_spent,omitempty"`
				WoolhuntKills               int    `json:"woolhunt_kills,omitempty"`
				WoolhuntKillsOnWoolholder   int    `json:"woolhunt_kills_on_woolholder,omitempty"`
				WoolhuntKillsWithWool       int    `json:"woolhunt_kills_with_wool,omitempty"`
				WoolhuntLongestGame         int    `json:"woolhunt_longest_game,omitempty"`
				WoolhuntMostGoldEarned      int    `json:"woolhunt_most_gold_earned,omitempty"`
				WoolhuntMostKillsAndAssists int    `json:"woolhunt_most_kills_and_assists,omitempty"`
				WoolhuntParticipatedLosses  int    `json:"woolhunt_participated_losses,omitempty"`
				WoolhuntParticipatedWins    int    `json:"woolhunt_participated_wins,omitempty"`
				WoolhuntWoolsCaptured       int    `json:"woolhunt_wools_captured,omitempty"`
				WoolhuntWoolsStolen         int    `json:"woolhunt_wools_stolen,omitempty"`
				Dropper                     struct {
					GamesPlayed   int `json:"games_played,omitempty"`
					MapsCompleted int `json:"maps_completed,omitempty"`
					Fails         int `json:"fails,omitempty"`
				} `json:"dropper,omitempty"`
				BombardmentBestTimeParty  int `json:"bombardment_best_time_party,omitempty"`
				HighGroundBestScoreParty  int `json:"high_ground_best_score_party,omitempty"`
				HighGroundTotalScoreParty int `json:"high_ground_total_score_party,omitempty"`
				SpiderMazeBestTimeParty   int `json:"spider_maze_best_time_party,omitempty"`
			} `json:"Arcade,omitempty"`
			Arena struct {
				Coins         int      `json:"coins,omitempty"`
				Damage2V2     int      `json:"damage_2v2,omitempty"`
				Damage4V4     int      `json:"damage_4v4,omitempty"`
				Deaths2V2     int      `json:"deaths_2v2,omitempty"`
				Deaths4V4     int      `json:"deaths_4v4,omitempty"`
				Games2V2      int      `json:"games_2v2,omitempty"`
				Games4V4      int      `json:"games_4v4,omitempty"`
				Healed2V2     int      `json:"healed_2v2,omitempty"`
				Healed4V4     int      `json:"healed_4v4,omitempty"`
				Keys          int      `json:"keys,omitempty"`
				Kills4V4      int      `json:"kills_4v4,omitempty"`
				Losses2V2     int      `json:"losses_2v2,omitempty"`
				Losses4V4     int      `json:"losses_4v4,omitempty"`
				MagicalChest  int      `json:"magical_chest,omitempty"`
				Packages      []string `json:"packages,omitempty"`
				Penalty       int      `json:"penalty,omitempty"`
				Rating        float64  `json:"rating,omitempty"`
				WinStreaks2V2 int      `json:"win_streaks_2v2,omitempty"`
				WinStreaks4V4 int      `json:"win_streaks_4v4,omitempty"`
				Wins4V4       int      `json:"wins_4v4,omitempty"`
				Wins2V2       int      `json:"wins_2v2,omitempty"`
				WinStreaks1V1 int      `json:"win_streaks_1v1,omitempty"`
				Deaths1V1     int      `json:"deaths_1v1,omitempty"`
				Damage1V1     int      `json:"damage_1v1,omitempty"`
				Games1V1      int      `json:"games_1v1,omitempty"`
				Losses1V1     int      `json:"losses_1v1,omitempty"`
				Healed1V1     int      `json:"healed_1v1,omitempty"`
				Wins          int      `json:"wins,omitempty"`
				Ultimate      string   `json:"ultimate,omitempty"`
				Wins1V1       int      `json:"wins_1v1,omitempty"`
			} `json:"Arena,omitempty"`
			Battleground struct {
				Assists                   int      `json:"assists,omitempty"`
				AvengerPlays              int      `json:"avenger_plays,omitempty"`
				BerserkerPlays            int      `json:"berserker_plays,omitempty"`
				BrokenInventory           int      `json:"broken_inventory,omitempty"`
				ChosenClass               string   `json:"chosen_class,omitempty"`
				Coins                     int      `json:"coins,omitempty"`
				Damage                    int      `json:"damage,omitempty"`
				DamageAvenger             int      `json:"damage_avenger,omitempty"`
				DamageBerserker           int      `json:"damage_berserker,omitempty"`
				DamageMage                int      `json:"damage_mage,omitempty"`
				DamagePaladin             int      `json:"damage_paladin,omitempty"`
				DamagePrevented           int      `json:"damage_prevented,omitempty"`
				DamagePreventedAvenger    int      `json:"damage_prevented_avenger,omitempty"`
				DamagePreventedBerserker  int      `json:"damage_prevented_berserker,omitempty"`
				DamagePreventedMage       int      `json:"damage_prevented_mage,omitempty"`
				DamagePreventedPaladin    int      `json:"damage_prevented_paladin,omitempty"`
				DamagePreventedPyromancer int      `json:"damage_prevented_pyromancer,omitempty"`
				DamagePreventedWarrior    int      `json:"damage_prevented_warrior,omitempty"`
				DamagePyromancer          int      `json:"damage_pyromancer,omitempty"`
				DamageTaken               int      `json:"damage_taken,omitempty"`
				DamageWarrior             int      `json:"damage_warrior,omitempty"`
				Deaths                    int      `json:"deaths,omitempty"`
				FlagConquerTeam           int      `json:"flag_conquer_team,omitempty"`
				Heal                      int      `json:"heal,omitempty"`
				HealAvenger               int      `json:"heal_avenger,omitempty"`
				HealBerserker             int      `json:"heal_berserker,omitempty"`
				HealPaladin               int      `json:"heal_paladin,omitempty"`
				HealWarrior               int      `json:"heal_warrior,omitempty"`
				Kills                     int      `json:"kills,omitempty"`
				LifeLeeched               int      `json:"life_leeched,omitempty"`
				LifeLeechedBerserker      int      `json:"life_leeched_berserker,omitempty"`
				LifeLeechedWarrior        int      `json:"life_leeched_warrior,omitempty"`
				Losses                    int      `json:"losses,omitempty"`
				LossesMage                int      `json:"losses_mage,omitempty"`
				LossesPyromancer          int      `json:"losses_pyromancer,omitempty"`
				MagePlays                 int      `json:"mage_plays,omitempty"`
				MageSpec                  string   `json:"mage_spec,omitempty"`
				PaladinPlays              int      `json:"paladin_plays,omitempty"`
				PaladinSpec               string   `json:"paladin_spec,omitempty"`
				PyromancerPlays           int      `json:"pyromancer_plays,omitempty"`
				SelectedMount             string   `json:"selected_mount,omitempty"`
				WarriorPlays              int      `json:"warrior_plays,omitempty"`
				WarriorSpec               string   `json:"warrior_spec,omitempty"`
				WinStreak                 int      `json:"win_streak,omitempty"`
				Wins                      int      `json:"wins,omitempty"`
				WinsAvenger               int      `json:"wins_avenger,omitempty"`
				WinsCapturetheflag        int      `json:"wins_capturetheflag,omitempty"`
				WinsCapturetheflagRed     int      `json:"wins_capturetheflag_red,omitempty"`
				WinsDomination            int      `json:"wins_domination,omitempty"`
				WinsDominationRed         int      `json:"wins_domination_red,omitempty"`
				WinsPaladin               int      `json:"wins_paladin,omitempty"`
				WinsRed                   int      `json:"wins_red,omitempty"`
				HealMage                  int      `json:"heal_mage,omitempty"`
				HealPyromancer            int      `json:"heal_pyromancer,omitempty"`
				Packages                  []string `json:"packages,omitempty"`
				ShamanSpec                string   `json:"shaman_spec,omitempty"`
				KillsCapturetheflag       int      `json:"kills_capturetheflag,omitempty"`
				PlayStreak                int      `json:"play_streak,omitempty"`
				WinsCapturetheflagB       int      `json:"wins_capturetheflag_b,omitempty"`
				WinsMage                  int      `json:"wins_mage,omitempty"`
				WinsPyromancer            int      `json:"wins_pyromancer,omitempty"`
			} `json:"Battleground,omitempty"`
			HungerGames struct {
				Coins               int      `json:"coins,omitempty"`
				Deaths              int      `json:"deaths,omitempty"`
				Kills               int      `json:"kills,omitempty"`
				Wins                int      `json:"wins,omitempty"`
				MonthlyKillsA       int      `json:"monthly_kills_a,omitempty"`
				WeeklyKillsA        int      `json:"weekly_kills_a,omitempty"`
				MonthlyKillsB       int      `json:"monthly_kills_b,omitempty"`
				WeeklyKillsB        int      `json:"weekly_kills_b,omitempty"`
				Packages            []string `json:"packages,omitempty"`
				LastTourneyAd       int64    `json:"lastTourneyAd,omitempty"`
				WinsSoloNormal      int      `json:"wins_solo_normal,omitempty"`
				Autoarmor           bool     `json:"autoarmor,omitempty"`
				WinsBackup          int      `json:"wins_backup,omitempty"`
				WinsTeamsNormal     int      `json:"wins_teams_normal,omitempty"`
				ChestsOpened        int      `json:"chests_opened,omitempty"`
				ChestsOpenedKnight  int      `json:"chests_opened_knight,omitempty"`
				Damage              int      `json:"damage,omitempty"`
				DamageKnight        int      `json:"damage_knight,omitempty"`
				DamageTaken         int      `json:"damage_taken,omitempty"`
				DamageTakenKnight   int      `json:"damage_taken_knight,omitempty"`
				GamesPlayed         int      `json:"games_played,omitempty"`
				GamesPlayedKnight   int      `json:"games_played_knight,omitempty"`
				TimePlayed          int      `json:"time_played,omitempty"`
				TimePlayedKnight    int      `json:"time_played_knight,omitempty"`
				PotionsThrown       int      `json:"potions_thrown,omitempty"`
				PotionsThrownKnight int      `json:"potions_thrown_knight,omitempty"`
			} `json:"HungerGames,omitempty"`
			Mcgo struct {
				BombsDefused                int `json:"bombs_defused,omitempty"`
				BombsPlanted                int `json:"bombs_planted,omitempty"`
				BountyHunter                int `json:"bounty_hunter,omitempty"`
				Coins                       int `json:"coins,omitempty"`
				CopKills                    int `json:"cop_kills,omitempty"`
				CriminalKills               int `json:"criminal_kills,omitempty"`
				Deaths                      int `json:"deaths,omitempty"`
				GameWins                    int `json:"game_wins,omitempty"`
				GamesWins1010               int `json:"games_wins_10_10,omitempty"`
				GamesWins102014             int `json:"games_wins_10_2014,omitempty"`
				GamesWins21010              int `json:"games_wins_2_10_10,omitempty"`
				GamesWins3102014            int `json:"games_wins_3_10_2014,omitempty"`
				HeadshotKills               int `json:"headshot_kills,omitempty"`
				Kills                       int `json:"kills,omitempty"`
				KillsNew                    int `json:"killsNew,omitempty"`
				KillsNew102014              int `json:"killsNew_10_2014,omitempty"`
				KillsNew112014              int `json:"killsNew_11_2014,omitempty"`
				KillsNew1112014             int `json:"killsNew_1_11_2014,omitempty"`
				KillsNew2112014             int `json:"killsNew_2_11_2014,omitempty"`
				KillsNew3112014             int `json:"killsNew_3_11_2014,omitempty"`
				KillsNew4112014             int `json:"killsNew_4_11_2014,omitempty"`
				KillsNew5102014             int `json:"killsNew_5_10_2014,omitempty"`
				KillsNew5112014             int `json:"killsNew_5_11_2014,omitempty"`
				KillsNew6112014             int `json:"killsNew_6_11_2014,omitempty"`
				Kills102014                 int `json:"kills_10_2014,omitempty"`
				Kills112014                 int `json:"kills_11_2014,omitempty"`
				Kills1112014                int `json:"kills_1_11_2014,omitempty"`
				Kills2112014                int `json:"kills_2_11_2014,omitempty"`
				Kills3102014                int `json:"kills_3_10_2014,omitempty"`
				Kills3112014                int `json:"kills_3_11_2014,omitempty"`
				Kills4102014                int `json:"kills_4_10_2014,omitempty"`
				Kills4112014                int `json:"kills_4_11_2014,omitempty"`
				Kills5102014                int `json:"kills_5_10_2014,omitempty"`
				Kills5112014                int `json:"kills_5_11_2014,omitempty"`
				Kills6112014                int `json:"kills_6_11_2014,omitempty"`
				RoundWins                   int `json:"round_wins,omitempty"`
				ShotsFired                  int `json:"shots_fired,omitempty"`
				SmgDamageIncrease           int `json:"smg_damage_increase,omitempty"`
				StrengthTraining            int `json:"strength_training,omitempty"`
				BodyArmorCost               int `json:"body_armor_cost,omitempty"`
				RifleCostReduction          int `json:"rifle_cost_reduction,omitempty"`
				CarbineCostReduction        int `json:"carbine_cost_reduction,omitempty"`
				CarbineReloadSpeedReduction int `json:"carbine_reload_speed_reduction,omitempty"`
				RifleReloadSpeedReduction   int `json:"rifle_reload_speed_reduction,omitempty"`
				PocketChange                int `json:"pocket_change,omitempty"`
				MonthlyKillsB               int `json:"monthly_kills_b,omitempty"`
				Mcgo                        struct {
					Points int `json:"points,omitempty"`
				} `json:"mcgo,omitempty"`
				Packages                         []string `json:"packages,omitempty"`
				DeathsDeathmatch                 int      `json:"deaths_deathmatch,omitempty"`
				GameWinsSandstorm                int      `json:"game_wins_sandstorm,omitempty"`
				GameWinsDeathmatch               int      `json:"game_wins_deathmatch,omitempty"`
				CriminalKillsDeathmatch          int      `json:"criminal_kills_deathmatch,omitempty"`
				WeeklyKillsA                     int      `json:"weekly_kills_a,omitempty"`
				KillsDeathmatch                  int      `json:"kills_deathmatch,omitempty"`
				GameWinsCarrier                  int      `json:"game_wins_carrier,omitempty"`
				GameWinsOvergrown                int      `json:"game_wins_overgrown,omitempty"`
				GameWinsAlleyway                 int      `json:"game_wins_alleyway,omitempty"`
				GameWinsTemple                   int      `json:"game_wins_temple,omitempty"`
				CopKillsDeathmatch               int      `json:"cop_kills_deathmatch,omitempty"`
				CarbineDamageIncrease            int      `json:"carbine_damage_increase,omitempty"`
				CarbineRecoilReduction           int      `json:"carbine_recoil_reduction,omitempty"`
				RifleRecoilReduction             int      `json:"rifle_recoil_reduction,omitempty"`
				RifleDamageIncrease              int      `json:"rifle_damage_increase,omitempty"`
				GameWinsAtomic                   int      `json:"game_wins_atomic,omitempty"`
				WeeklyKillsB                     int      `json:"weekly_kills_b,omitempty"`
				SelectedCreeperChestplateDev     string   `json:"selectedCreeperChestplateDev,omitempty"`
				SelectedCreeperHelmetDev         string   `json:"selectedCreeperHelmetDev,omitempty"`
				SelectedOcelotChestplateDev      string   `json:"selectedOcelotChestplateDev,omitempty"`
				SelectedOcelotHelmetDev          string   `json:"selectedOcelotHelmetDev,omitempty"`
				MonthlyKillsA                    int      `json:"monthly_kills_a,omitempty"`
				SelectedKnifeDev                 string   `json:"selectedKnifeDev,omitempty"`
				SelectedCarbineDev               string   `json:"selectedCarbineDev,omitempty"`
				ShowLobbyPrefix                  bool     `json:"show_lobby_prefix,omitempty"`
				GameWinsReserve                  int      `json:"game_wins_reserve,omitempty"`
				SelectedLobbyPrefix              string   `json:"selected_lobby_prefix,omitempty"`
				GrenadeKills                     int      `json:"grenadeKills,omitempty"`
				GrenadeKills0                    int      `json:"grenade_kills,omitempty"`
				GameWinsMelonFactory             int      `json:"game_wins_melon factory,omitempty"`
				LastTourneyAd                    int64    `json:"lastTourneyAd,omitempty"`
				BombsPlantedTourneyMcgoDefusal0  int      `json:"bombs_planted_tourney_mcgo_defusal_0,omitempty"`
				CopKillsTourneyMcgoDefusal0      int      `json:"cop_kills_tourney_mcgo_defusal_0,omitempty"`
				CriminalKillsTourneyMcgoDefusal0 int      `json:"criminal_kills_tourney_mcgo_defusal_0,omitempty"`
				GamePlaysTourneyMcgoDefusal0     int      `json:"game_plays_tourney_mcgo_defusal_0,omitempty"`
				GameWinsTourneyMcgoDefusal0      int      `json:"game_wins_tourney_mcgo_defusal_0,omitempty"`
				HeadshotKillsTourneyMcgoDefusal0 int      `json:"headshot_kills_tourney_mcgo_defusal_0,omitempty"`
				KillsTourneyMcgoDefusal0         int      `json:"kills_tourney_mcgo_defusal_0,omitempty"`
				RoundWinsTourneyMcgoDefusal0     int      `json:"round_wins_tourney_mcgo_defusal_0,omitempty"`
				ShotsFiredTourneyMcgoDefusal0    int      `json:"shots_fired_tourney_mcgo_defusal_0,omitempty"`
				DeathsTourneyMcgoDefusal0        int      `json:"deaths_tourney_mcgo_defusal_0,omitempty"`
				GrenadeKillsTourneyMcgoDefusal0  int      `json:"grenade_kills_tourney_mcgo_defusal_0,omitempty"`
				BombsDefusedTourneyMcgoDefusal0  int      `json:"bombs_defused_tourney_mcgo_defusal_0,omitempty"`
				GamePlaysDeathmatch              int      `json:"game_plays_deathmatch,omitempty"`
				ScopedRifleHeadshots             int      `json:"scopedRifleHeadshots,omitempty"`
				ScopedRifleKills                 int      `json:"scopedRifleKills,omitempty"`
				RifleKills                       int      `json:"rifleKills,omitempty"`
				RifleHeadshots                   int      `json:"rifleHeadshots,omitempty"`
				ShotgunKills                     int      `json:"shotgunKills,omitempty"`
				ShotgunHeadshots                 int      `json:"shotgunHeadshots,omitempty"`
				CarbineKills                     int      `json:"carbineKills,omitempty"`
				CarbineHeadshots                 int      `json:"carbineHeadshots,omitempty"`
				SmgKills                         int      `json:"smgKills,omitempty"`
				SmgHeadshots                     int      `json:"smgHeadshots,omitempty"`
				Assists                          int      `json:"assists,omitempty"`
				GamePlays                        int      `json:"game_plays,omitempty"`
				GameWinsJunction                 int      `json:"game_wins_junction,omitempty"`
				PistolKills                      int      `json:"pistolKills,omitempty"`
				PistolHeadshots                  int      `json:"pistolHeadshots,omitempty"`
				AssistsDeathmatch                int      `json:"assists_deathmatch,omitempty"`
				AutoShotgunKills                 int      `json:"autoShotgunKills,omitempty"`
				AutoShotgunHeadshots             int      `json:"autoShotgunHeadshots,omitempty"`
				GameWinsDerailed                 int      `json:"game_wins_derailed,omitempty"`
				GameWinsBazaar                   int      `json:"game_wins_bazaar,omitempty"`
				ShoutTotal                       int      `json:"shoutTotal,omitempty"`
				BullpupHeadshots                 int      `json:"bullpupHeadshots,omitempty"`
				BullpupKills                     int      `json:"bullpupKills,omitempty"`
				HandgunKills                     int      `json:"handgunKills,omitempty"`
				HandgunHeadshots                 int      `json:"handgunHeadshots,omitempty"`
				SniperHeadshots                  int      `json:"sniperHeadshots,omitempty"`
				SniperKills                      int      `json:"sniperKills,omitempty"`
				MagnumKills                      int      `json:"magnumKills,omitempty"`
				MagnumHeadshots                  int      `json:"magnumHeadshots,omitempty"`
				LobbyPrefixColor                 string   `json:"lobbyPrefixColor,omitempty"`
				BombsPlantedTourneyMcgoDefusal1  int      `json:"bombs_planted_tourney_mcgo_defusal_1,omitempty"`
				CopKillsTourneyMcgoDefusal1      int      `json:"cop_kills_tourney_mcgo_defusal_1,omitempty"`
				CriminalKillsTourneyMcgoDefusal1 int      `json:"criminal_kills_tourney_mcgo_defusal_1,omitempty"`
				DeathsTourneyMcgoDefusal1        int      `json:"deaths_tourney_mcgo_defusal_1,omitempty"`
				GamePlaysTourneyMcgoDefusal1     int      `json:"game_plays_tourney_mcgo_defusal_1,omitempty"`
				GrenadeKillsTourneyMcgoDefusal1  int      `json:"grenade_kills_tourney_mcgo_defusal_1,omitempty"`
				HeadshotKillsTourneyMcgoDefusal1 int      `json:"headshot_kills_tourney_mcgo_defusal_1,omitempty"`
				KillsTourneyMcgoDefusal1         int      `json:"kills_tourney_mcgo_defusal_1,omitempty"`
				RoundWinsTourneyMcgoDefusal1     int      `json:"round_wins_tourney_mcgo_defusal_1,omitempty"`
				ShotsFiredTourneyMcgoDefusal1    int      `json:"shots_fired_tourney_mcgo_defusal_1,omitempty"`
				BombsDefusedTourneyMcgoDefusal1  int      `json:"bombs_defused_tourney_mcgo_defusal_1,omitempty"`
				GameWinsTourneyMcgoDefusal1      int      `json:"game_wins_tourney_mcgo_defusal_1,omitempty"`
				KillsGungame                     int      `json:"kills_gungame,omitempty"`
				GameWinsGungame                  int      `json:"game_wins_gungame,omitempty"`
				GameWinsAtomicV2                 int      `json:"game_wins_atomic v2,omitempty"`
				GameWinsRuins                    int      `json:"game_wins_ruins,omitempty"`
				GameWinsCastle                   int      `json:"game_wins_castle,omitempty"`
				DeathsGungame                    int      `json:"deaths_gungame,omitempty"`
				CarePackagesCollectedGungame     int      `json:"care_packages_collected_gungame,omitempty"`
				CopKillsGungame                  int      `json:"cop_kills_gungame,omitempty"`
				GamePlaysGungame                 int      `json:"game_plays_gungame,omitempty"`
			} `json:"MCGO,omitempty"`
			Paintball struct {
				Coins          int      `json:"coins,omitempty"`
				Deaths         int      `json:"deaths,omitempty"`
				Hat            string   `json:"hat,omitempty"`
				Headstart      int      `json:"headstart,omitempty"`
				Kills          int      `json:"kills,omitempty"`
				Killstreaks    int      `json:"killstreaks,omitempty"`
				Packages       []string `json:"packages,omitempty"`
				ShotsFired     int      `json:"shots_fired,omitempty"`
				Superluck      int      `json:"superluck,omitempty"`
				Wins           int      `json:"wins,omitempty"`
				Transfusion    int      `json:"transfusion,omitempty"`
				VotesLaLaLand  int      `json:"votes_LaLaLand,omitempty"`
				VotesGladiator int      `json:"votes_Gladiator,omitempty"`
				Fortune        int      `json:"fortune,omitempty"`
				FavoriteSlots  string   `json:"favorite_slots,omitempty"`
				ShowKillPrefix bool     `json:"showKillPrefix,omitempty"`
				Adrenaline     int      `json:"adrenaline,omitempty"`
				Godfather      int      `json:"godfather,omitempty"`
				Endurance      int      `json:"endurance,omitempty"`
			} `json:"Paintball,omitempty"`
			Quake struct {
				Coins                              float64  `json:"coins,omitempty"`
				Deaths                             int      `json:"deaths,omitempty"`
				Kills                              int      `json:"kills,omitempty"`
				Killstreaks                        int      `json:"killstreaks,omitempty"`
				Wins                               int      `json:"wins,omitempty"`
				Packages                           []string `json:"packages,omitempty"`
				Case                               string   `json:"case,omitempty"`
				Trigger                            string   `json:"trigger,omitempty"`
				Barrel                             string   `json:"barrel,omitempty"`
				Sight                              string   `json:"sight,omitempty"`
				Null                               string   `json:"null,omitempty"`
				Killsound                          string   `json:"killsound,omitempty"`
				KillsTeams                         int      `json:"kills_teams,omitempty"`
				DeathsTeams                        int      `json:"deaths_teams,omitempty"`
				KillstreaksTeams                   int      `json:"killstreaks_teams,omitempty"`
				WinsTeams                          int      `json:"wins_teams,omitempty"`
				MonthlyKillsB                      int      `json:"monthly_kills_b,omitempty"`
				WeeklyKillsB                       int      `json:"weekly_kills_b,omitempty"`
				WeeklyKillsA                       int      `json:"weekly_kills_a,omitempty"`
				MonthlyKillsA                      int      `json:"monthly_kills_a,omitempty"`
				CompassSelected                    bool     `json:"compass_selected,omitempty"`
				AlternativeGunCooldownIndicator    bool     `json:"alternative_gun_cooldown_indicator,omitempty"`
				EnableSound                        bool     `json:"enable_sound,omitempty"`
				ShowDashCooldown                   bool     `json:"showDashCooldown,omitempty"`
				HighestKillstreak                  int      `json:"highest_killstreak,omitempty"`
				InstantRespawn                     bool     `json:"instantRespawn,omitempty"`
				DistanceTravelledTeams             int      `json:"distance_travelled_teams,omitempty"`
				KillsSinceUpdateFeb2017Teams       int      `json:"kills_since_update_feb_2017_teams,omitempty"`
				ShotsFiredTeams                    int      `json:"shots_fired_teams,omitempty"`
				HeadshotsTeams                     int      `json:"headshots_teams,omitempty"`
				Headshots                          int      `json:"headshots,omitempty"`
				DistanceTravelled                  int      `json:"distance_travelled,omitempty"`
				ShotsFired                         int      `json:"shots_fired,omitempty"`
				KillsSinceUpdateFeb2017            int      `json:"kills_since_update_feb_2017,omitempty"`
				ShowKillPrefix                     bool     `json:"showKillPrefix,omitempty"`
				MessageOthersKillsDeaths           bool     `json:"messageOthers' Kills/deaths,omitempty"`
				MessageYourDeaths                  bool     `json:"messageYour Deaths,omitempty"`
				MessageKillstreaks                 bool     `json:"messageKillstreaks,omitempty"`
				MessageCoinMessages                bool     `json:"messageCoin Messages,omitempty"`
				MessageYourKills                   bool     `json:"messageYour Kills,omitempty"`
				MessagePowerupCollections          bool     `json:"messagePowerup Collections,omitempty"`
				MessageMultiKills                  bool     `json:"messageMulti-kills,omitempty"`
				KillsSoloTourney                   int      `json:"kills_solo_tourney,omitempty"`
				ShotsFiredSoloTourney              int      `json:"shots_fired_solo_tourney,omitempty"`
				KillsSinceUpdateFeb2017SoloTourney int      `json:"kills_since_update_feb_2017_solo_tourney,omitempty"`
				DistanceTravelledSoloTourney       int      `json:"distance_travelled_solo_tourney,omitempty"`
				DeathsSoloTourney                  int      `json:"deaths_solo_tourney,omitempty"`
				HeadshotsSoloTourney               int      `json:"headshots_solo_tourney,omitempty"`
				WinsSoloTourney                    int      `json:"wins_solo_tourney,omitempty"`
				KillstreaksSoloTourney             int      `json:"killstreaks_solo_tourney,omitempty"`
				MessageCoin                        bool     `json:"messageCoin,omitempty"`
				KillsDmTeams                       int      `json:"kills_dm_teams,omitempty"`
				KillsDm                            int      `json:"kills_dm,omitempty"`
				KillsTimeattack                    int      `json:"kills_timeattack,omitempty"`
				DashCooldown                       string   `json:"dash_cooldown,omitempty"`
				Muzzle                             string   `json:"muzzle,omitempty"`
				Beam                               string   `json:"beam,omitempty"`
				LastTourneyAd                      int64    `json:"lastTourneyAd,omitempty"`
				KillsTourneyUnknown                int      `json:"kills_tourney_unknown,omitempty"`
			} `json:"Quake,omitempty"`
			TNTGames struct {
				Coins                       int      `json:"coins,omitempty"`
				DeathsBowspleef             int      `json:"deaths_bowspleef,omitempty"`
				DeathsCapture               int      `json:"deaths_capture,omitempty"`
				KillsCapture                int      `json:"kills_capture,omitempty"`
				TagsBowspleef               int      `json:"tags_bowspleef,omitempty"`
				WinsCapture                 int      `json:"wins_capture,omitempty"`
				RecordTntrun                int      `json:"record_tntrun,omitempty"`
				CaptureClass                string   `json:"capture_class,omitempty"`
				AssistsCapture              int      `json:"assists_capture,omitempty"`
				WinsBowspleef               int      `json:"wins_bowspleef,omitempty"`
				Packages                    []string `json:"packages,omitempty"`
				NewIcewizardRegen           int      `json:"new_icewizard_regen,omitempty"`
				Wins                        int      `json:"wins,omitempty"`
				NewSpleefTripleshot         int      `json:"new_spleef_tripleshot,omitempty"`
				NewTntrunDoubleJumps        int      `json:"new_tntrun_double_jumps,omitempty"`
				NewWitherwizardRegen        int      `json:"new_witherwizard_regen,omitempty"`
				NewKineticwizardRegen       int      `json:"new_kineticwizard_regen,omitempty"`
				NewTntagSpeedy              int      `json:"new_tntag_speedy,omitempty"`
				NewWitherwizardExplode      int      `json:"new_witherwizard_explode,omitempty"`
				NewFirewizardRegen          int      `json:"new_firewizard_regen,omitempty"`
				NewFirewizardExplode        int      `json:"new_firewizard_explode,omitempty"`
				NewSpleefDoubleJumps        int      `json:"new_spleef_double_jumps,omitempty"`
				NewBloodwizardRegen         int      `json:"new_bloodwizard_regen,omitempty"`
				NewIcewizardExplode         int      `json:"new_icewizard_explode,omitempty"`
				NewPvprunDoubleJumps        int      `json:"new_pvprun_double_jumps,omitempty"`
				NewSpleefRepulsor           int      `json:"new_spleef_repulsor,omitempty"`
				NewBloodwizardExplode       int      `json:"new_bloodwizard_explode,omitempty"`
				NewKineticwizardExplode     int      `json:"new_kineticwizard_explode,omitempty"`
				Winstreak                   int      `json:"winstreak,omitempty"`
				RunPotionsSplashedOnPlayers int      `json:"run_potions_splashed_on_players,omitempty"`
				RecordPvprun                int      `json:"record_pvprun,omitempty"`
				KillsPvprun                 int      `json:"kills_pvprun,omitempty"`
				DeathsPvprun                int      `json:"deaths_pvprun,omitempty"`
				WizardsSelectedClass        string   `json:"wizards_selected_class,omitempty"`
				NewIcewizardDeaths          int      `json:"new_icewizard_deaths,omitempty"`
				NewIcewizardAssists         int      `json:"new_icewizard_assists,omitempty"`
				NewIcewizardKills           int      `json:"new_icewizard_kills,omitempty"`
				DeathsTntrun                int      `json:"deaths_tntrun,omitempty"`
				NewWitherwizardAssists      int      `json:"new_witherwizard_assists,omitempty"`
				NewWitherwizardKills        int      `json:"new_witherwizard_kills,omitempty"`
				NewWitherwizardDeaths       int      `json:"new_witherwizard_deaths,omitempty"`
				NewToxicwizardExplode       int      `json:"new_toxicwizard_explode,omitempty"`
				KillsTntag                  int      `json:"kills_tntag,omitempty"`
				VotesDreadPit               int      `json:"votes_Dread Pit,omitempty"`
				PointsCapture               int      `json:"points_capture,omitempty"`
				Flags                       struct {
					ShowTipHolograms                 bool `json:"show_tip_holograms,omitempty"`
					ShowTntrunActionbarInfo          bool `json:"show_tntrun_actionbar_info,omitempty"`
					ShowTnttagActionbarInfo          bool `json:"show_tnttag_actionbar_info,omitempty"`
					ShowWizardsActionbarInfo         bool `json:"show_wizards_actionbar_info,omitempty"`
					ShowWizardsCooldownNotifications bool `json:"show_wizards_cooldown_notifications,omitempty"`
					EnableExplosiveDash              bool `json:"enable_explosive_dash,omitempty"`
				} `json:"flags,omitempty"`
				NewAncientwizardRegen       int    `json:"new_ancientwizard_regen,omitempty"`
				NewAncientwizardExplode     int    `json:"new_ancientwizard_explode,omitempty"`
				NewAncientwizardDeaths      int    `json:"new_ancientwizard_deaths,omitempty"`
				NewAncientwizardHealing     int    `json:"new_ancientwizard_healing,omitempty"`
				NewWitherwizardHealing      int    `json:"new_witherwizard_healing,omitempty"`
				NewAncientwizardDamageTaken int    `json:"new_ancientwizard_damage_taken,omitempty"`
				KineticHealingCapture       int    `json:"kinetic_healing_capture,omitempty"`
				NewWitherwizardDamageTaken  int    `json:"new_witherwizard_damage_taken,omitempty"`
				AirTimeCapture              int    `json:"air_time_capture,omitempty"`
				NewAncientwizardAssists     int    `json:"new_ancientwizard_assists,omitempty"`
				NewAncientwizardKills       int    `json:"new_ancientwizard_kills,omitempty"`
				NewSelectedHat              string `json:"new_selected_hat,omitempty"`
				ActiveVoidMessage           string `json:"active_void_message,omitempty"`
				NewSpleefExlosiveDash       int    `json:"new_spleef_exlosive_dash,omitempty"`
				NewSpleefArrowrain          int    `json:"new_spleef_arrowrain,omitempty"`
				NewTntrunSpeedPotions       int    `json:"new_tntrun_speed_potions,omitempty"`
				WinsTntrun                  int    `json:"wins_tntrun,omitempty"`
				LastTourneyAd               int64  `json:"lastTourneyAd,omitempty"`
				DeathsTntag                 int    `json:"deaths_tntag,omitempty"`
			} `json:"TNTGames,omitempty"`
			Uhc struct {
				Coins              int  `json:"coins,omitempty"`
				Cache3             bool `json:"cache3,omitempty"`
				ClearupAchievement bool `json:"clearup_achievement,omitempty"`
				SavedStats         bool `json:"saved_stats,omitempty"`
			} `json:"UHC,omitempty"`
			VampireZ struct {
				Coins               int  `json:"coins,omitempty"`
				HumanDeaths         int  `json:"human_deaths,omitempty"`
				VampireDeaths       int  `json:"vampire_deaths,omitempty"`
				VampireKills        int  `json:"vampire_kills,omitempty"`
				ZombieKills         int  `json:"zombie_kills,omitempty"`
				UpdatedStats        bool `json:"updated_stats,omitempty"`
				MostVampireKillsNew int  `json:"most_vampire_kills_new,omitempty"`
			} `json:"VampireZ,omitempty"`
			Walls struct {
				Deaths int `json:"deaths,omitempty"`
				Losses int `json:"losses,omitempty"`
				Coins  int `json:"coins,omitempty"`
				Kills  int `json:"kills,omitempty"`
			} `json:"Walls,omitempty"`
			Walls3 struct {
				ChosenClass                             string   `json:"chosen_class,omitempty"`
				Coins                                   int      `json:"coins,omitempty"`
				Deaths                                  int      `json:"deaths,omitempty"`
				DeathsCreeper                           int      `json:"deaths_Creeper,omitempty"`
				DeathsSkeleton                          int      `json:"deaths_Skeleton,omitempty"`
				FinalDeaths                             int      `json:"finalDeaths,omitempty"`
				FinalKills                              int      `json:"finalKills,omitempty"`
				FinalKillsSkeleton                      int      `json:"finalKills_Skeleton,omitempty"`
				Kills                                   int      `json:"kills,omitempty"`
				KillsCreeper                            int      `json:"kills_Creeper,omitempty"`
				KillsSkeleton                           int      `json:"kills_Skeleton,omitempty"`
				Losses                                  int      `json:"losses,omitempty"`
				LossesCreeper                           int      `json:"losses_Creeper,omitempty"`
				LossesSkeleton                          int      `json:"losses_Skeleton,omitempty"`
				LossesZombie                            int      `json:"losses_Zombie,omitempty"`
				WeeklyDeaths                            int      `json:"weeklyDeaths,omitempty"`
				WeeklyDeathsSkeleton                    int      `json:"weeklyDeaths_Skeleton,omitempty"`
				WeeklyLosses                            int      `json:"weeklyLosses,omitempty"`
				WeeklyLossesSkeleton                    int      `json:"weeklyLosses_Skeleton,omitempty"`
				Wins                                    int      `json:"wins,omitempty"`
				WinsSkeleton                            int      `json:"wins_Skeleton,omitempty"`
				Packages                                []string `json:"packages,omitempty"`
				PlaysStandard                           int      `json:"plays_standard,omitempty"`
				DeathsNewSkeleton                       int      `json:"deaths_new_Skeleton,omitempty"`
				Assists                                 int      `json:"assists,omitempty"`
				DeathsNew                               int      `json:"deaths_new,omitempty"`
				AssistsSkeleton                         int      `json:"assists_Skeleton,omitempty"`
				SkeletonFinalKills                      int      `json:"skeleton_final_kills,omitempty"`
				TotalFinalKillsStandard                 int      `json:"total_final_kills_standard,omitempty"`
				FinalKillsStandard                      int      `json:"final_kills_standard,omitempty"`
				SkeletonKills                           int      `json:"skeleton_kills,omitempty"`
				FinalKills0                             int      `json:"final_kills,omitempty"`
				CreeperLosses                           int      `json:"creeper_losses,omitempty"`
				CreeperDeaths                           int      `json:"creeper_deaths,omitempty"`
				TotalFinalKills                         int      `json:"total_final_kills,omitempty"`
				CreeperKills                            int      `json:"creeper_kills,omitempty"`
				SkeletonTotalFinalKills                 int      `json:"skeleton_total_final_kills,omitempty"`
				SkeletonLosses                          int      `json:"skeleton_losses,omitempty"`
				SkeletonFinalKillsStandard              int      `json:"skeleton_final_kills_standard,omitempty"`
				SkeletonWins                            int      `json:"skeleton_wins,omitempty"`
				SkeletonWinsStandard                    int      `json:"skeleton_wins_standard,omitempty"`
				SkeletonDeaths                          int      `json:"skeleton_deaths,omitempty"`
				SkeletonTotalFinalKillsStandard         int      `json:"skeleton_total_final_kills_standard,omitempty"`
				ZombieLosses                            int      `json:"zombie_losses,omitempty"`
				IronArmorGiftedStandard                 int      `json:"iron_armor_gifted_standard,omitempty"`
				SkeletonAssists                         int      `json:"skeleton_assists,omitempty"`
				SkeletonDefenderAssists                 int      `json:"skeleton_defender_assists,omitempty"`
				SkeletonTotalDeaths                     int      `json:"skeleton_total_deaths,omitempty"`
				SkeletonArrowsFiredStandard             int      `json:"skeleton_arrows_fired_standard,omitempty"`
				SkeletonDefenderKillsStandard           int      `json:"skeleton_defender_kills_standard,omitempty"`
				DefenderAssistsStandard                 int      `json:"defender_assists_standard,omitempty"`
				SkeletonFinalKillsMelee                 int      `json:"skeleton_final_kills_melee,omitempty"`
				SkeletonFinalKillsMeleeBehindStandard   int      `json:"skeleton_final_kills_melee_behind_standard,omitempty"`
				SkeletonBlocksPlacedPreparation         int      `json:"skeleton_blocks_placed_preparation,omitempty"`
				TotalDeaths                             int      `json:"total_deaths,omitempty"`
				SkeletonBlocksPlaced                    int      `json:"skeleton_blocks_placed,omitempty"`
				LossesStandard                          int      `json:"losses_standard,omitempty"`
				TotalKills                              int      `json:"total_kills,omitempty"`
				ActivationsStandard                     int      `json:"activations_standard,omitempty"`
				TreasuresFound                          int      `json:"treasures_found,omitempty"`
				WoodChopped                             int      `json:"wood_chopped,omitempty"`
				ArrowsFiredStandard                     int      `json:"arrows_fired_standard,omitempty"`
				FinalKillsMelee                         int      `json:"final_kills_melee,omitempty"`
				SkeletonKillsMeleeStandard              int      `json:"skeleton_kills_melee_standard,omitempty"`
				SkeletonMetersWalked                    int      `json:"skeleton_meters_walked,omitempty"`
				SkeletonBlocksPlacedPreparationStandard int      `json:"skeleton_blocks_placed_preparation_standard,omitempty"`
				SkeletonTimePlayed                      int      `json:"skeleton_time_played,omitempty"`
				SkeletonFinalKillsMeleeBehind           int      `json:"skeleton_final_kills_melee_behind,omitempty"`
				SkeletonFinalDeathsStandard             int      `json:"skeleton_final_deaths_standard,omitempty"`
				TimePlayed                              int      `json:"time_played,omitempty"`
				DefenderKillsStandard                   int      `json:"defender_kills_standard,omitempty"`
				SkeletonAAssistsStandard                int      `json:"skeleton_a_assists_standard,omitempty"`
				WoodChoppedStandard                     int      `json:"wood_chopped_standard,omitempty"`
				SkeletonAActivationsStandard            int      `json:"skeleton_a_activations_standard,omitempty"`
				SkeletonBlocksBrokenStandard            int      `json:"skeleton_blocks_broken_standard,omitempty"`
				FinalDeaths0                            int      `json:"final_deaths,omitempty"`
				SkeletonMetersFallenStandard            int      `json:"skeleton_meters_fallen_standard,omitempty"`
				GamesPlayed                             int      `json:"games_played,omitempty"`
				BlocksPlacedStandard                    int      `json:"blocks_placed_standard,omitempty"`
				BlocksPlacedPreparation                 int      `json:"blocks_placed_preparation,omitempty"`
				SkeletonArrowsHit                       int      `json:"skeleton_arrows_hit,omitempty"`
				SkeletonArrowsHitStandard               int      `json:"skeleton_arrows_hit_standard,omitempty"`
				SkeletonIronOreBroken                   int      `json:"skeleton_iron_ore_broken,omitempty"`
				SkeletonActivations                     int      `json:"skeleton_activations,omitempty"`
				BlocksPlacedPreparationStandard         int      `json:"blocks_placed_preparation_standard,omitempty"`
				SkeletonKillsMelee                      int      `json:"skeleton_kills_melee,omitempty"`
				SkeletonBlocksPlacedStandard            int      `json:"skeleton_blocks_placed_standard,omitempty"`
				TotalKillsStandard                      int      `json:"total_kills_standard,omitempty"`
				SkeletonATotalKills                     int      `json:"skeleton_a_total_kills,omitempty"`
				SkeletonGamesPlayedStandard             int      `json:"skeleton_games_played_standard,omitempty"`
				TimePlayedStandard                      int      `json:"time_played_standard,omitempty"`
				SkeletonMetersWalkedSpeedStandard       int      `json:"skeleton_meters_walked_speed_standard,omitempty"`
				SkeletonTreasuresFoundStandard          int      `json:"skeleton_treasures_found_standard,omitempty"`
				FinalKillsMeleeBehind                   int      `json:"final_kills_melee_behind,omitempty"`
				KillsMeleeStandard                      int      `json:"kills_melee_standard,omitempty"`
				FinalKillsMeleeBehindStandard           int      `json:"final_kills_melee_behind_standard,omitempty"`
				KillsStandard                           int      `json:"kills_standard,omitempty"`
				IronOreBrokenStandard                   int      `json:"iron_ore_broken_standard,omitempty"`
				TreasuresFoundStandard                  int      `json:"treasures_found_standard,omitempty"`
				SkeletonTreasuresFound                  int      `json:"skeleton_treasures_found,omitempty"`
				BlocksBrokenStandard                    int      `json:"blocks_broken_standard,omitempty"`
				MetersWalkedSpeedStandard               int      `json:"meters_walked_speed_standard,omitempty"`
				SkeletonAAssists                        int      `json:"skeleton_a_assists,omitempty"`
				SkeletonAssistsStandard                 int      `json:"skeleton_assists_standard,omitempty"`
				SkeletonLossesStandard                  int      `json:"skeleton_losses_standard,omitempty"`
				SkeletonTotalDeathsStandard             int      `json:"skeleton_total_deaths_standard,omitempty"`
				SkeletonMetersWalkedSpeed               int      `json:"skeleton_meters_walked_speed,omitempty"`
				FinalDeathsStandard                     int      `json:"final_deaths_standard,omitempty"`
				SkeletonTotalKills                      int      `json:"skeleton_total_kills,omitempty"`
				KillsMelee                              int      `json:"kills_melee,omitempty"`
				SkeletonDefenderKills                   int      `json:"skeleton_defender_kills,omitempty"`
				DefenderKills                           int      `json:"defender_kills,omitempty"`
				SkeletonFinalKillsMeleeStandard         int      `json:"skeleton_final_kills_melee_standard,omitempty"`
				IronArmorGifted                         int      `json:"iron_armor_gifted,omitempty"`
				MetersWalked                            int      `json:"meters_walked,omitempty"`
				MetersFallenStandard                    int      `json:"meters_fallen_standard,omitempty"`
				BlocksBroken                            int      `json:"blocks_broken,omitempty"`
				Activations                             int      `json:"activations,omitempty"`
				AssistsStandard                         int      `json:"assists_standard,omitempty"`
				BlocksPlaced                            int      `json:"blocks_placed,omitempty"`
				SkeletonTimePlayedStandard              int      `json:"skeleton_time_played_standard,omitempty"`
				MetersWalkedSpeed                       int      `json:"meters_walked_speed,omitempty"`
				SkeletonKillsStandard                   int      `json:"skeleton_kills_standard,omitempty"`
				SkeletonTotalKillsStandard              int      `json:"skeleton_total_kills_standard,omitempty"`
				ArrowsFired                             int      `json:"arrows_fired,omitempty"`
				SkeletonMetersFallen                    int      `json:"skeleton_meters_fallen,omitempty"`
				IronOreBroken                           int      `json:"iron_ore_broken,omitempty"`
				DeathsStandard                          int      `json:"deaths_standard,omitempty"`
				TotalDeathsStandard                     int      `json:"total_deaths_standard,omitempty"`
				SkeletonAActivations                    int      `json:"skeleton_a_activations,omitempty"`
				SkeletonATotalKillsStandard             int      `json:"skeleton_a_total_kills_standard,omitempty"`
				SkeletonArrowsFired                     int      `json:"skeleton_arrows_fired,omitempty"`
				SkeletonBlocksBroken                    int      `json:"skeleton_blocks_broken,omitempty"`
				SkeletonDeathsStandard                  int      `json:"skeleton_deaths_standard,omitempty"`
				SkeletonIronArmorGifted                 int      `json:"skeleton_iron_armor_gifted,omitempty"`
				SkeletonDefenderAssistsStandard         int      `json:"skeleton_defender_assists_standard,omitempty"`
				FinalKillsMeleeStandard                 int      `json:"final_kills_melee_standard,omitempty"`
				DefenderAssists                         int      `json:"defender_assists,omitempty"`
				ArrowsHitStandard                       int      `json:"arrows_hit_standard,omitempty"`
				SkeletonMetersWalkedStandard            int      `json:"skeleton_meters_walked_standard,omitempty"`
				SkeletonADefenderAssists                int      `json:"skeleton_a_defender_assists,omitempty"`
				MetersFallen                            int      `json:"meters_fallen,omitempty"`
				SkeletonGamesPlayed                     int      `json:"skeleton_games_played,omitempty"`
				SkeletonIronArmorGiftedStandard         int      `json:"skeleton_iron_armor_gifted_standard,omitempty"`
				SkeletonFinalDeaths                     int      `json:"skeleton_final_deaths,omitempty"`
				SkeletonWoodChopped                     int      `json:"skeleton_wood_chopped,omitempty"`
				MetersWalkedStandard                    int      `json:"meters_walked_standard,omitempty"`
				ArrowsHit                               int      `json:"arrows_hit,omitempty"`
				SkeletonADefenderAssistsStandard        int      `json:"skeleton_a_defender_assists_standard,omitempty"`
				GamesPlayedStandard                     int      `json:"games_played_standard,omitempty"`
				SkeletonWoodChoppedStandard             int      `json:"skeleton_wood_chopped_standard,omitempty"`
				SkeletonActivationsStandard             int      `json:"skeleton_activations_standard,omitempty"`
				SkeletonIronOreBrokenStandard           int      `json:"skeleton_iron_ore_broken_standard,omitempty"`
				SkeletonFinalAssists                    int      `json:"skeleton_final_assists,omitempty"`
				SkeletonFinalAssistsStandard            int      `json:"skeleton_final_assists_standard,omitempty"`
				SkeletonFoodEatenStandard               int      `json:"skeleton_food_eaten_standard,omitempty"`
				DefenderFinalAssists                    int      `json:"defender_final_assists,omitempty"`
				SkeletonFoodEaten                       int      `json:"skeleton_food_eaten,omitempty"`
				SkeletonSteaksEaten                     int      `json:"skeleton_steaks_eaten,omitempty"`
				WitherDamageStandard                    int      `json:"wither_damage_standard,omitempty"`
				FoodEatenStandard                       int      `json:"food_eaten_standard,omitempty"`
				SkeletonDefenderFinalAssistsStandard    int      `json:"skeleton_defender_final_assists_standard,omitempty"`
				SkeletonSteaksEatenStandard             int      `json:"skeleton_steaks_eaten_standard,omitempty"`
				FinalAssists                            int      `json:"final_assists,omitempty"`
				SkeletonWitherDamageStandard            int      `json:"skeleton_wither_damage_standard,omitempty"`
				SkeletonDefenderFinalAssists            int      `json:"skeleton_defender_final_assists,omitempty"`
				FoodEaten                               int      `json:"food_eaten,omitempty"`
				FinalAssistsStandard                    int      `json:"final_assists_standard,omitempty"`
				SteaksEaten                             int      `json:"steaks_eaten,omitempty"`
				DefenderFinalAssistsStandard            int      `json:"defender_final_assists_standard,omitempty"`
				SkeletonWitherDamage                    int      `json:"skeleton_wither_damage,omitempty"`
				SteaksEatenStandard                     int      `json:"steaks_eaten_standard,omitempty"`
				WinsStandard                            int      `json:"wins_standard,omitempty"`
				WitherDamage                            int      `json:"wither_damage,omitempty"`
				Classes                                 struct {
					Herobrine struct {
						SkillLevelDChecked5 bool `json:"skill_level_dChecked5,omitempty"`
						SkillLevelD         int  `json:"skill_level_d,omitempty"`
						Unlocked            bool `json:"unlocked,omitempty"`
						Checked4            bool `json:"checked4,omitempty"`
					} `json:"herobrine,omitempty"`
					Enderman struct {
						SkillLevelDChecked5 bool `json:"skill_level_dChecked5,omitempty"`
						SkillLevelD         int  `json:"skill_level_d,omitempty"`
						Checked4            bool `json:"checked4,omitempty"`
						Unlocked            bool `json:"unlocked,omitempty"`
					} `json:"enderman,omitempty"`
					Skeleton struct {
						SkillLevelD         int  `json:"skill_level_d,omitempty"`
						SkillLevelBChecked5 bool `json:"skill_level_bChecked5,omitempty"`
						SkillLevelA         int  `json:"skill_level_a,omitempty"`
						SkillLevelDChecked5 bool `json:"skill_level_dChecked5,omitempty"`
						SkillLevelAChecked5 bool `json:"skill_level_aChecked5,omitempty"`
						Checked4            bool `json:"checked4,omitempty"`
						Unlocked            bool `json:"unlocked,omitempty"`
						SkillLevelGChecked5 bool `json:"skill_level_gChecked5,omitempty"`
						SkillLevelCChecked5 bool `json:"skill_level_cChecked5,omitempty"`
					} `json:"skeleton,omitempty"`
					Zombie struct {
						SkillLevelD         int  `json:"skill_level_d,omitempty"`
						SkillLevelDChecked5 bool `json:"skill_level_dChecked5,omitempty"`
						Checked4            bool `json:"checked4,omitempty"`
						Unlocked            bool `json:"unlocked,omitempty"`
					} `json:"zombie,omitempty"`
				} `json:"classes,omitempty"`
				KillsMeleeBehind                 int `json:"kills_melee_behind,omitempty"`
				KillsMeleeBehindStandard         int `json:"kills_melee_behind_standard,omitempty"`
				PotionsDrunk                     int `json:"potions_drunk,omitempty"`
				PotionsDrunkStandard             int `json:"potions_drunk_standard,omitempty"`
				SkeletonBActivations             int `json:"skeleton_b_activations,omitempty"`
				SkeletonBActivationsStandard     int `json:"skeleton_b_activations_standard,omitempty"`
				SkeletonCActivations             int `json:"skeleton_c_activations,omitempty"`
				SkeletonCActivationsStandard     int `json:"skeleton_c_activations_standard,omitempty"`
				SkeletonGActivations             int `json:"skeleton_g_activations,omitempty"`
				SkeletonGActivationsStandard     int `json:"skeleton_g_activations_standard,omitempty"`
				SkeletonKillsMeleeBehind         int `json:"skeleton_kills_melee_behind,omitempty"`
				SkeletonKillsMeleeBehindStandard int `json:"skeleton_kills_melee_behind_standard,omitempty"`
				SkeletonPotionsDrunk             int `json:"skeleton_potions_drunk,omitempty"`
				SkeletonPotionsDrunkStandard     int `json:"skeleton_potions_drunk_standard,omitempty"`
			} `json:"Walls3,omitempty"`
			GingerBread struct {
				FrameActive          string   `json:"frame_active,omitempty"`
				EngineActive         string   `json:"engine_active,omitempty"`
				BoosterActive        string   `json:"booster_active,omitempty"`
				Packages             []string `json:"packages,omitempty"`
				HelmetActive         string   `json:"helmet_active,omitempty"`
				JacketActive         string   `json:"jacket_active,omitempty"`
				PantsActive          string   `json:"pants_active,omitempty"`
				ShoesActive          string   `json:"shoes_active,omitempty"`
				SkinActive           string   `json:"skin_active,omitempty"`
				Horn                 string   `json:"horn,omitempty"`
				BoxPickups           int      `json:"box_pickups,omitempty"`
				LapsCompleted        int      `json:"laps_completed,omitempty"`
				BananaHitsReceived   int      `json:"banana_hits_received,omitempty"`
				CoinsPickedUp        int      `json:"coins_picked_up,omitempty"`
				Parts                string   `json:"parts,omitempty"`
				Coins                int      `json:"coins,omitempty"`
				BoxPickupsWeeklyB    float64  `json:"box_pickups_weekly_b,omitempty"`
				BoxPickupsMonthlyB   float64  `json:"box_pickups_monthly_b,omitempty"`
				BoxPickupsHypixelgp  float64  `json:"box_pickups_hypixelgp,omitempty"`
				HypixelgpPlays       float64  `json:"hypixelgp_plays,omitempty"`
				CanyonPlays          int      `json:"canyon_plays,omitempty"`
				BoxPickupsWeeklyA    int      `json:"box_pickups_weekly_a,omitempty"`
				BoxPickupsCanyon     int      `json:"box_pickups_canyon,omitempty"`
				BoxPickupsJunglerush int      `json:"box_pickups_junglerush,omitempty"`
				JunglerushPlays      int      `json:"junglerush_plays,omitempty"`
				BananaHitsSent       int      `json:"banana_hits_sent,omitempty"`
				RetroPlays           int      `json:"retro_plays,omitempty"`
				BoxPickupsRetro      int      `json:"box_pickups_retro,omitempty"`
				LastTourneyAd        int64    `json:"lastTourneyAd,omitempty"`
				BoxPickupsMonthlyA   int      `json:"box_pickups_monthly_a,omitempty"`
			} `json:"GingerBread,omitempty"`
			SkyWars struct {
				WinStreak                                    int      `json:"win_streak,omitempty"`
				SurvivedPlayers                              int      `json:"survived_players,omitempty"`
				Losses                                       int      `json:"losses,omitempty"`
				SurvivedPlayersKitBasicSoloDefault           int      `json:"survived_players_kit_basic_solo_default,omitempty"`
				LossesSolo                                   int      `json:"losses_solo,omitempty"`
				BlocksBroken                                 int      `json:"blocks_broken,omitempty"`
				DeathsKitBasicSoloDefault                    int      `json:"deaths_kit_basic_solo_default,omitempty"`
				LossesSoloNormal                             int      `json:"losses_solo_normal,omitempty"`
				BlocksPlaced                                 int      `json:"blocks_placed,omitempty"`
				DeathsSolo                                   int      `json:"deaths_solo,omitempty"`
				Deaths                                       int      `json:"deaths,omitempty"`
				Quits                                        int      `json:"quits,omitempty"`
				DeathsSoloNormal                             int      `json:"deaths_solo_normal,omitempty"`
				SurvivedPlayersSolo                          int      `json:"survived_players_solo,omitempty"`
				LossesKitBasicSoloDefault                    int      `json:"losses_kit_basic_solo_default,omitempty"`
				Packages                                     []string `json:"packages,omitempty"`
				Coins                                        int      `json:"coins,omitempty"`
				DeathsTeamNormal                             int      `json:"deaths_team_normal,omitempty"`
				LossesKitMiningTeamDefault                   int      `json:"losses_kit_mining_team_default,omitempty"`
				LossesTeamNormal                             int      `json:"losses_team_normal,omitempty"`
				DeathsTeam                                   int      `json:"deaths_team,omitempty"`
				DeathsKitMiningTeamDefault                   int      `json:"deaths_kit_mining_team_default,omitempty"`
				LossesTeam                                   int      `json:"losses_team,omitempty"`
				SurvivedPlayersKitMiningTeamDefault          int      `json:"survived_players_kit_mining_team_default,omitempty"`
				SurvivedPlayersTeam                          int      `json:"survived_players_team,omitempty"`
				Kills                                        int      `json:"kills,omitempty"`
				KillsTeamNormal                              int      `json:"kills_team_normal,omitempty"`
				GamesKitMiningTeamDefault                    int      `json:"games_kit_mining_team_default,omitempty"`
				KillsTeam                                    int      `json:"kills_team,omitempty"`
				KillsKitMiningTeamDefault                    int      `json:"kills_kit_mining_team_default,omitempty"`
				AssistsKitMiningTeamDefault                  int      `json:"assists_kit_mining_team_default,omitempty"`
				GamesTeam                                    int      `json:"games_team,omitempty"`
				AssistsTeam                                  int      `json:"assists_team,omitempty"`
				SoulsGathered                                int      `json:"souls_gathered,omitempty"`
				Assists                                      int      `json:"assists,omitempty"`
				Games                                        int      `json:"games,omitempty"`
				Souls                                        int      `json:"souls,omitempty"`
				LossesMegaNormal                             int      `json:"losses_mega_normal,omitempty"`
				DeathsMega                                   int      `json:"deaths_mega,omitempty"`
				SurvivedPlayersKitMegaMegaDefault            int      `json:"survived_players_kit_mega_mega_default,omitempty"`
				DeathsKitMegaMegaDefault                     int      `json:"deaths_kit_mega_mega_default,omitempty"`
				DeathsMegaNormal                             int      `json:"deaths_mega_normal,omitempty"`
				LossesKitMegaMegaDefault                     int      `json:"losses_kit_mega_mega_default,omitempty"`
				SurvivedPlayersMega                          int      `json:"survived_players_mega,omitempty"`
				LossesMega                                   int      `json:"losses_mega,omitempty"`
				AssistsMega                                  int      `json:"assists_mega,omitempty"`
				AssistsKitMegaMegaDefault                    int      `json:"assists_kit_mega_mega_default,omitempty"`
				KillsKitMegaMegaDefault                      int      `json:"kills_kit_mega_mega_default,omitempty"`
				KillsMega                                    int      `json:"kills_mega,omitempty"`
				KillsMegaNormal                              int      `json:"kills_mega_normal,omitempty"`
				EggThrown                                    int      `json:"egg_thrown,omitempty"`
				ArrowsHit                                    int      `json:"arrows_hit,omitempty"`
				ArrowsShot                                   int      `json:"arrows_shot,omitempty"`
				KillsKitBasicSoloDefault                     int      `json:"kills_kit_basic_solo_default,omitempty"`
				LossesSoloInsane                             int      `json:"losses_solo_insane,omitempty"`
				DeathsSoloInsane                             int      `json:"deaths_solo_insane,omitempty"`
				KillsSolo                                    int      `json:"kills_solo,omitempty"`
				KillsSoloInsane                              int      `json:"kills_solo_insane,omitempty"`
				SoulWell                                     int      `json:"soul_well,omitempty"`
				UsedSoulWell                                 bool     `json:"usedSoulWell,omitempty"`
				GamesSolo                                    int      `json:"games_solo,omitempty"`
				KillsSoloNormal                              int      `json:"kills_solo_normal,omitempty"`
				WinsSoloNormal                               int      `json:"wins_solo_normal,omitempty"`
				Wins                                         int      `json:"wins,omitempty"`
				WinsKitBasicSoloDefault                      int      `json:"wins_kit_basic_solo_default,omitempty"`
				WinsSolo                                     int      `json:"wins_solo,omitempty"`
				GamesKitBasicSoloDefault                     int      `json:"games_kit_basic_solo_default,omitempty"`
				VotesDragonice                               int      `json:"votes_Dragonice,omitempty"`
				ItemsEnchanted                               int      `json:"items_enchanted,omitempty"`
				WinsKitMiningTeamDefault                     int      `json:"wins_kit_mining_team_default,omitempty"`
				WinsTeamNormal                               int      `json:"wins_team_normal,omitempty"`
				WinsTeam                                     int      `json:"wins_team,omitempty"`
				DeathsTeamInsane                             int      `json:"deaths_team_insane,omitempty"`
				LossesTeamInsane                             int      `json:"losses_team_insane,omitempty"`
				KillsTeamInsane                              int      `json:"kills_team_insane,omitempty"`
				SoloMiningExpertise                          int      `json:"solo_mining_expertise,omitempty"`
				EnderpearlsThrown                            int      `json:"enderpearls_thrown,omitempty"`
				TeamEnderMastery                             int      `json:"team_ender_mastery,omitempty"`
				TeamInstantSmelting                          int      `json:"team_instant_smelting,omitempty"`
				AssistsKitBasicSoloDefault                   int      `json:"assists_kit_basic_solo_default,omitempty"`
				AssistsSolo                                  int      `json:"assists_solo,omitempty"`
				VotesCongo                                   int      `json:"votes_Congo,omitempty"`
				SoulWellRares                                int      `json:"soul_well_rares,omitempty"`
				MegaJuggernaut                               int      `json:"mega_juggernaut,omitempty"`
				VotesShire                                   int      `json:"votes_Shire,omitempty"`
				SoulWellLegendaries                          int      `json:"soul_well_legendaries,omitempty"`
				SoloMarksmanship                             int      `json:"solo_marksmanship,omitempty"`
				ActiveKitSOLO                                string   `json:"activeKit_SOLO,omitempty"`
				DeathsKitBasicSoloArmorsmith                 int      `json:"deaths_kit_basic_solo_armorsmith,omitempty"`
				GamesKitBasicSoloArmorsmith                  int      `json:"games_kit_basic_solo_armorsmith,omitempty"`
				LossesKitBasicSoloArmorsmith                 int      `json:"losses_kit_basic_solo_armorsmith,omitempty"`
				SurvivedPlayersKitBasicSoloArmorsmith        int      `json:"survived_players_kit_basic_solo_armorsmith,omitempty"`
				KillsKitBasicSoloArmorsmith                  int      `json:"kills_kit_basic_solo_armorsmith,omitempty"`
				WinsKitBasicSoloArmorsmith                   int      `json:"wins_kit_basic_solo_armorsmith,omitempty"`
				AssistsKitBasicSoloArmorsmith                int      `json:"assists_kit_basic_solo_armorsmith,omitempty"`
				VotesEntangled                               int      `json:"votes_Entangled,omitempty"`
				TeamResistanceBoost                          int      `json:"team_resistance_boost,omitempty"`
				ActiveKitTEAM                                string   `json:"activeKit_TEAM,omitempty"`
				SurvivedPlayersKitAttackingTeamKnight        int      `json:"survived_players_kit_attacking_team_knight,omitempty"`
				DeathsKitAttackingTeamKnight                 int      `json:"deaths_kit_attacking_team_knight,omitempty"`
				LossesKitAttackingTeamKnight                 int      `json:"losses_kit_attacking_team_knight,omitempty"`
				KillsKitAttackingTeamKnight                  int      `json:"kills_kit_attacking_team_knight,omitempty"`
				GamesKitAttackingTeamKnight                  int      `json:"games_kit_attacking_team_knight,omitempty"`
				ActiveCage                                   string   `json:"activeCage,omitempty"`
				AssistsKitAttackingTeamKnight                int      `json:"assists_kit_attacking_team_knight,omitempty"`
				WinsKitAttackingTeamKnight                   int      `json:"wins_kit_attacking_team_knight,omitempty"`
				MegaRusher                                   int      `json:"mega_rusher,omitempty"`
				PaidSouls                                    int      `json:"paid_souls,omitempty"`
				XezbethLuck                                  int      `json:"xezbeth_luck,omitempty"`
				WinsTeamInsane                               int      `json:"wins_team_insane,omitempty"`
				VotesJinzhou                                 int      `json:"votes_Jinzhou,omitempty"`
				VotesLongIsland                              int      `json:"votes_LongIsland,omitempty"`
				TeamMiningExpertise                          int      `json:"team_mining_expertise,omitempty"`
				MegaMiningExpertise                          int      `json:"mega_mining_expertise,omitempty"`
				DeathsKitAdvancedSoloEnchanter               int      `json:"deaths_kit_advanced_solo_enchanter,omitempty"`
				GamesKitAdvancedSoloEnchanter                int      `json:"games_kit_advanced_solo_enchanter,omitempty"`
				KillsKitAdvancedSoloEnchanter                int      `json:"kills_kit_advanced_solo_enchanter,omitempty"`
				LossesKitAdvancedSoloEnchanter               int      `json:"losses_kit_advanced_solo_enchanter,omitempty"`
				SurvivedPlayersKitAdvancedSoloEnchanter      int      `json:"survived_players_kit_advanced_solo_enchanter,omitempty"`
				WinsKitAdvancedSoloEnchanter                 int      `json:"wins_kit_advanced_solo_enchanter,omitempty"`
				VotesSiege                                   int      `json:"votes_Siege,omitempty"`
				SurvivedPlayersKitDefendingTeamArmorer       int      `json:"survived_players_kit_defending_team_armorer,omitempty"`
				DeathsKitDefendingTeamArmorer                int      `json:"deaths_kit_defending_team_armorer,omitempty"`
				LossesKitDefendingTeamArmorer                int      `json:"losses_kit_defending_team_armorer,omitempty"`
				KillsKitDefendingTeamArmorer                 int      `json:"kills_kit_defending_team_armorer,omitempty"`
				GamesKitDefendingTeamArmorer                 int      `json:"games_kit_defending_team_armorer,omitempty"`
				ExtraWheels                                  int      `json:"extra_wheels,omitempty"`
				HighestKillstreak                            int      `json:"highestKillstreak,omitempty"`
				Killstreak                                   int      `json:"killstreak,omitempty"`
				KillsWeeklyB                                 int      `json:"kills_weekly_b,omitempty"`
				KillsMonthlyB                                int      `json:"kills_monthly_b,omitempty"`
				MegaNourishment                              int      `json:"mega_nourishment,omitempty"`
				GamesPlayedSkywars                           int      `json:"games_played_skywars,omitempty"`
				LastMode                                     string   `json:"lastMode,omitempty"`
				TimePlayedKitAttackingTeamKnight             int      `json:"time_played_kit_attacking_team_knight,omitempty"`
				KillsWeeklyA                                 int      `json:"kills_weekly_a,omitempty"`
				MeleeKillsTeam                               int      `json:"melee_kills_team,omitempty"`
				VoidKillsKitAttackingTeamKnight              int      `json:"void_kills_kit_attacking_team_knight,omitempty"`
				VoidKills                                    int      `json:"void_kills,omitempty"`
				ChestsOpened                                 int      `json:"chests_opened,omitempty"`
				MeleeKills                                   int      `json:"melee_kills,omitempty"`
				MostKillsGameTeam                            int      `json:"most_kills_game_team,omitempty"`
				ChestsOpenedTeam                             int      `json:"chests_opened_team,omitempty"`
				TimePlayedTeam                               int      `json:"time_played_team,omitempty"`
				MeleeKillsKitAttackingTeamKnight             int      `json:"melee_kills_kit_attacking_team_knight,omitempty"`
				TimePlayed                                   int      `json:"time_played,omitempty"`
				MostKillsGame                                int      `json:"most_kills_game,omitempty"`
				MostKillsGameKitAttackingTeamKnight          int      `json:"most_kills_game_kit_attacking_team_knight,omitempty"`
				VoidKillsTeam                                int      `json:"void_kills_team,omitempty"`
				ChestsOpenedKitAttackingTeamKnight           int      `json:"chests_opened_kit_attacking_team_knight,omitempty"`
				KillsMonthlyA                                int      `json:"kills_monthly_a,omitempty"`
				RefillChestDestroy                           int      `json:"refill_chest_destroy,omitempty"`
				TimePlayedKitMegaMegaDefault                 int      `json:"time_played_kit_mega_mega_default,omitempty"`
				ChestsOpenedKitMegaMegaDefault               int      `json:"chests_opened_kit_mega_mega_default,omitempty"`
				ChestsOpenedMega                             int      `json:"chests_opened_mega,omitempty"`
				TimePlayedMega                               int      `json:"time_played_mega,omitempty"`
				SlimeExplainedLast                           int64    `json:"slime_explained_last,omitempty"`
				SlimeExplained                               int      `json:"slime_explained,omitempty"`
				WinStreakLab                                 int      `json:"win_streak_lab,omitempty"`
				SurvivedPlayersLabSolo                       int      `json:"survived_players_lab_solo,omitempty"`
				ChestsOpenedLabSolo                          int      `json:"chests_opened_lab_solo,omitempty"`
				DeathsLab                                    int      `json:"deaths_lab,omitempty"`
				DeathsLabSolo                                int      `json:"deaths_lab_solo,omitempty"`
				BlocksBrokenLab                              int      `json:"blocks_broken_lab,omitempty"`
				TimePlayedLabSolo                            int      `json:"time_played_lab_solo,omitempty"`
				TimePlayedLab                                int      `json:"time_played_lab,omitempty"`
				GamesLabKitAdvancedSoloEnchanter             int      `json:"games_lab_kit_advanced_solo_enchanter,omitempty"`
				GamesLab                                     int      `json:"games_lab,omitempty"`
				DeathsLabKitAdvancedSoloEnchanter            int      `json:"deaths_lab_kit_advanced_solo_enchanter,omitempty"`
				SurvivedPlayersLabKitAdvancedSoloEnchanter   int      `json:"survived_players_lab_kit_advanced_solo_enchanter,omitempty"`
				SurvivedPlayersLab                           int      `json:"survived_players_lab,omitempty"`
				ChestsOpenedLabKitAdvancedSoloEnchanter      int      `json:"chests_opened_lab_kit_advanced_solo_enchanter,omitempty"`
				CoinsGainedLab                               int      `json:"coins_gained_lab,omitempty"`
				ChestsOpenedLab                              int      `json:"chests_opened_lab,omitempty"`
				LossesLabSolo                                int      `json:"losses_lab_solo,omitempty"`
				GamesLabSolo                                 int      `json:"games_lab_solo,omitempty"`
				TimePlayedLabKitAdvancedSoloEnchanter        int      `json:"time_played_lab_kit_advanced_solo_enchanter,omitempty"`
				LossesLab                                    int      `json:"losses_lab,omitempty"`
				BlocksPlacedLab                              int      `json:"blocks_placed_lab,omitempty"`
				LossesLabKitAdvancedSoloEnchanter            int      `json:"losses_lab_kit_advanced_solo_enchanter,omitempty"`
				LongestBowShotLab                            int      `json:"longest_bow_shot_lab,omitempty"`
				LongestBowShotLabKitAdvancedSoloEnchanter    int      `json:"longest_bow_shot_lab_kit_advanced_solo_enchanter,omitempty"`
				LongestBowShotLabSolo                        int      `json:"longest_bow_shot_lab_solo,omitempty"`
				ArrowsShotLabSolo                            int      `json:"arrows_shot_lab_solo,omitempty"`
				ArrowsShotLabKitAdvancedSoloEnchanter        int      `json:"arrows_shot_lab_kit_advanced_solo_enchanter,omitempty"`
				ArrowsHitLabKitAdvancedSoloEnchanter         int      `json:"arrows_hit_lab_kit_advanced_solo_enchanter,omitempty"`
				ArrowsHitLab                                 int      `json:"arrows_hit_lab,omitempty"`
				QuitsLab                                     int      `json:"quits_lab,omitempty"`
				ArrowsShotLab                                int      `json:"arrows_shot_lab,omitempty"`
				ArrowsHitLabSolo                             int      `json:"arrows_hit_lab_solo,omitempty"`
				SurvivedPlayersLabTeam                       int      `json:"survived_players_lab_team,omitempty"`
				ChestsOpenedLabTeam                          int      `json:"chests_opened_lab_team,omitempty"`
				TimePlayedLabTeam                            int      `json:"time_played_lab_team,omitempty"`
				TimePlayedLabKitAttackingTeamDefault         int      `json:"time_played_lab_kit_attacking_team_default,omitempty"`
				SurvivedPlayersLabKitAttackingTeamDefault    int      `json:"survived_players_lab_kit_attacking_team_default,omitempty"`
				AssistsLabTeam                               int      `json:"assists_lab_team,omitempty"`
				AssistsLabKitAttackingTeamDefault            int      `json:"assists_lab_kit_attacking_team_default,omitempty"`
				DeathsLabKitAttackingTeamDefault             int      `json:"deaths_lab_kit_attacking_team_default,omitempty"`
				AssistsLab                                   int      `json:"assists_lab,omitempty"`
				LossesLabTeam                                int      `json:"losses_lab_team,omitempty"`
				LossesLabKitAttackingTeamDefault             int      `json:"losses_lab_kit_attacking_team_default,omitempty"`
				DeathsLabTeam                                int      `json:"deaths_lab_team,omitempty"`
				ChestsOpenedLabKitAttackingTeamDefault       int      `json:"chests_opened_lab_kit_attacking_team_default,omitempty"`
				LongestBowKillLabTeam                        int      `json:"longest_bow_kill_lab_team,omitempty"`
				MostKillsGameLabTeam                         int      `json:"most_kills_game_lab_team,omitempty"`
				LongestBowKillLab                            int      `json:"longest_bow_kill_lab,omitempty"`
				MostKillsGameLabKitAttackingTeamDefault      int      `json:"most_kills_game_lab_kit_attacking_team_default,omitempty"`
				LongestBowKillLabKitAttackingTeamDefault     int      `json:"longest_bow_kill_lab_kit_attacking_team_default,omitempty"`
				MostKillsGameLab                             int      `json:"most_kills_game_lab,omitempty"`
				KillsLabKitAttackingTeamDefault              int      `json:"kills_lab_kit_attacking_team_default,omitempty"`
				KillsLab                                     int      `json:"kills_lab,omitempty"`
				MeleeKillsLabTeam                            int      `json:"melee_kills_lab_team,omitempty"`
				KillsLabTeam                                 int      `json:"kills_lab_team,omitempty"`
				MeleeKillsLab                                int      `json:"melee_kills_lab,omitempty"`
				MeleeKillsLabKitAttackingTeamDefault         int      `json:"melee_kills_lab_kit_attacking_team_default,omitempty"`
				SkywarsChests                                int      `json:"skywars_chests,omitempty"`
				LevelFormatted                               string   `json:"levelFormatted,omitempty"`
				SkywarsExperience                            float64  `json:"skywars_experience,omitempty"`
				SoloInstantSmelting                          int      `json:"solo_instant_smelting,omitempty"`
				SoloEnderMastery                             int      `json:"solo_ender_mastery,omitempty"`
				TeamMarksmanship                             int      `json:"team_marksmanship,omitempty"`
				SoloResistanceBoost                          int      `json:"solo_resistance_boost,omitempty"`
				ArrowsShotKitMiningTeamDefault               int      `json:"arrows_shot_kit_mining_team_default,omitempty"`
				ArrowsShotSolo                               int      `json:"arrows_shot_solo,omitempty"`
				ChestsOpenedKitMiningTeamDefault             int      `json:"chests_opened_kit_mining_team_default,omitempty"`
				ChestsOpenedSolo                             int      `json:"chests_opened_solo,omitempty"`
				TimePlayedKitMiningTeamDefault               int      `json:"time_played_kit_mining_team_default,omitempty"`
				TimePlayedSolo                               int      `json:"time_played_solo,omitempty"`
				SkywarsChristmasBoxes                        int      `json:"skywars_christmas_boxes,omitempty"`
				ActiveKitTEAMSRandom                         bool     `json:"activeKit_TEAMS_random,omitempty"`
				ActiveKitTEAMS                               string   `json:"activeKit_TEAMS,omitempty"`
				ArrowsShotTourney                            int      `json:"arrows_shot_tourney,omitempty"`
				ArrowsShotTourneyKitMiningTeamDefault        int      `json:"arrows_shot_tourney_kit_mining_team_default,omitempty"`
				ArrowsShotTourneyTeamsTourney                int      `json:"arrows_shot_tourney_teams_tourney,omitempty"`
				ChestsOpenedTourney                          int      `json:"chests_opened_tourney,omitempty"`
				ChestsOpenedTourneyKitMiningTeamDefault      int      `json:"chests_opened_tourney_kit_mining_team_default,omitempty"`
				ChestsOpenedTourneyTeamsTourney              int      `json:"chests_opened_tourney_teams_tourney,omitempty"`
				CoinsGainedTourney                           int      `json:"coins_gained_tourney,omitempty"`
				DeathsTourney                                int      `json:"deaths_tourney,omitempty"`
				DeathsTourneyKitMiningTeamDefault            int      `json:"deaths_tourney_kit_mining_team_default,omitempty"`
				DeathsTourneyTeamsTourney                    int      `json:"deaths_tourney_teams_tourney,omitempty"`
				LossesTourney                                int      `json:"losses_tourney,omitempty"`
				LossesTourneyKitMiningTeamDefault            int      `json:"losses_tourney_kit_mining_team_default,omitempty"`
				LossesTourneyTeamsTourney                    int      `json:"losses_tourney_teams_tourney,omitempty"`
				QuitsTourney                                 int      `json:"quits_tourney,omitempty"`
				SurvivedPlayersTourney                       int      `json:"survived_players_tourney,omitempty"`
				SurvivedPlayersTourneyKitMiningTeamDefault   int      `json:"survived_players_tourney_kit_mining_team_default,omitempty"`
				SurvivedPlayersTourneyTeamsTourney           int      `json:"survived_players_tourney_teams_tourney,omitempty"`
				TimePlayedTourney                            int      `json:"time_played_tourney,omitempty"`
				TimePlayedTourneyKitMiningTeamDefault        int      `json:"time_played_tourney_kit_mining_team_default,omitempty"`
				TimePlayedTourneyTeamsTourney                int      `json:"time_played_tourney_teams_tourney,omitempty"`
				TourneySwInsaneDoubles0ArrowsShot            int      `json:"tourney_sw_insane_doubles_0_arrows_shot,omitempty"`
				TourneySwInsaneDoubles0ChestsOpened          int      `json:"tourney_sw_insane_doubles_0_chests_opened,omitempty"`
				TourneySwInsaneDoubles0Coins                 int      `json:"tourney_sw_insane_doubles_0_coins,omitempty"`
				TourneySwInsaneDoubles0CoinsGained           int      `json:"tourney_sw_insane_doubles_0_coins_gained,omitempty"`
				TourneySwInsaneDoubles0Deaths                int      `json:"tourney_sw_insane_doubles_0_deaths,omitempty"`
				TourneySwInsaneDoubles0Losses                int      `json:"tourney_sw_insane_doubles_0_losses,omitempty"`
				TourneySwInsaneDoubles0Quits                 int      `json:"tourney_sw_insane_doubles_0_quits,omitempty"`
				TourneySwInsaneDoubles0SurvivedPlayers       int      `json:"tourney_sw_insane_doubles_0_survived_players,omitempty"`
				TourneySwInsaneDoubles0TimePlayed            int      `json:"tourney_sw_insane_doubles_0_time_played,omitempty"`
				TourneySwInsaneDoubles0WinStreak             int      `json:"tourney_sw_insane_doubles_0_win_streak,omitempty"`
				WinStreakTourney                             int      `json:"win_streak_tourney,omitempty"`
				ArrowsHitTourney                             int      `json:"arrows_hit_tourney,omitempty"`
				ArrowsHitTourneyKitDefendingTeamBatguy       int      `json:"arrows_hit_tourney_kit_defending_team_batguy,omitempty"`
				ArrowsHitTourneyTeamsTourney                 int      `json:"arrows_hit_tourney_teams_tourney,omitempty"`
				ArrowsShotTourneyKitDefendingTeamBatguy      int      `json:"arrows_shot_tourney_kit_defending_team_batguy,omitempty"`
				BlocksBrokenTourney                          int      `json:"blocks_broken_tourney,omitempty"`
				BlocksPlacedTourney                          int      `json:"blocks_placed_tourney,omitempty"`
				ChestsOpenedTourneyKitDefendingTeamBatguy    int      `json:"chests_opened_tourney_kit_defending_team_batguy,omitempty"`
				DeathsTourneyKitDefendingTeamBatguy          int      `json:"deaths_tourney_kit_defending_team_batguy,omitempty"`
				LongestBowShotTourney                        int      `json:"longest_bow_shot_tourney,omitempty"`
				LongestBowShotTourneyKitDefendingTeamBatguy  int      `json:"longest_bow_shot_tourney_kit_defending_team_batguy,omitempty"`
				LongestBowShotTourneyTeamsTourney            int      `json:"longest_bow_shot_tourney_teams_tourney,omitempty"`
				LossesTourneyKitDefendingTeamBatguy          int      `json:"losses_tourney_kit_defending_team_batguy,omitempty"`
				SurvivedPlayersTourneyKitDefendingTeamBatguy int      `json:"survived_players_tourney_kit_defending_team_batguy,omitempty"`
				TimePlayedTourneyKitDefendingTeamBatguy      int      `json:"time_played_tourney_kit_defending_team_batguy,omitempty"`
				TourneySwInsaneDoubles0ArrowsHit             int      `json:"tourney_sw_insane_doubles_0_arrows_hit,omitempty"`
				TourneySwInsaneDoubles0BlocksBroken          int      `json:"tourney_sw_insane_doubles_0_blocks_broken,omitempty"`
				TourneySwInsaneDoubles0BlocksPlaced          int      `json:"tourney_sw_insane_doubles_0_blocks_placed,omitempty"`
				TourneySwInsaneDoubles0LongestBowShot        int      `json:"tourney_sw_insane_doubles_0_longest_bow_shot,omitempty"`
				EggThrownTourney                             int      `json:"egg_thrown_tourney,omitempty"`
				GamesTourney                                 int      `json:"games_tourney,omitempty"`
				GamesTourneyKitDefendingTeamBatguy           int      `json:"games_tourney_kit_defending_team_batguy,omitempty"`
				GamesTourneyTeamsTourney                     int      `json:"games_tourney_teams_tourney,omitempty"`
				KillsTourney                                 int      `json:"kills_tourney,omitempty"`
				KillsTourneyKitDefendingTeamBatguy           int      `json:"kills_tourney_kit_defending_team_batguy,omitempty"`
				KillsTourneyTeamsTourney                     int      `json:"kills_tourney_teams_tourney,omitempty"`
				MostKillsGameTourney                         int      `json:"most_kills_game_tourney,omitempty"`
				MostKillsGameTourneyKitDefendingTeamBatguy   int      `json:"most_kills_game_tourney_kit_defending_team_batguy,omitempty"`
				MostKillsGameTourneyTeamsTourney             int      `json:"most_kills_game_tourney_teams_tourney,omitempty"`
				TourneySwInsaneDoubles0EggThrown             int      `json:"tourney_sw_insane_doubles_0_egg_thrown,omitempty"`
				TourneySwInsaneDoubles0Games                 int      `json:"tourney_sw_insane_doubles_0_games,omitempty"`
				TourneySwInsaneDoubles0Kills                 int      `json:"tourney_sw_insane_doubles_0_kills,omitempty"`
				TourneySwInsaneDoubles0MostKillsGame         int      `json:"tourney_sw_insane_doubles_0_most_kills_game,omitempty"`
				TourneySwInsaneDoubles0VoidKills             int      `json:"tourney_sw_insane_doubles_0_void_kills,omitempty"`
				VoidKillsTourney                             int      `json:"void_kills_tourney,omitempty"`
				VoidKillsTourneyKitDefendingTeamBatguy       int      `json:"void_kills_tourney_kit_defending_team_batguy,omitempty"`
				VoidKillsTourneyTeamsTourney                 int      `json:"void_kills_tourney_teams_tourney,omitempty"`
				AssistsTourney                               int      `json:"assists_tourney,omitempty"`
				AssistsTourneyKitDefendingTeamBatguy         int      `json:"assists_tourney_kit_defending_team_batguy,omitempty"`
				AssistsTourneyTeamsTourney                   int      `json:"assists_tourney_teams_tourney,omitempty"`
				TourneySwInsaneDoubles0Assists               int      `json:"tourney_sw_insane_doubles_0_assists,omitempty"`
				ChestsOpenedTourneyKitMythicalEndLord        int      `json:"chests_opened_tourney_kit_mythical_end-lord,omitempty"`
				DeathsTourneyKitMythicalEndLord              int      `json:"deaths_tourney_kit_mythical_end-lord,omitempty"`
				EnderpearlsThrownTourney                     int      `json:"enderpearls_thrown_tourney,omitempty"`
				KillsTourneyKitMythicalEndLord               int      `json:"kills_tourney_kit_mythical_end-lord,omitempty"`
				LongestBowKillTourney                        int      `json:"longest_bow_kill_tourney,omitempty"`
				LongestBowKillTourneyKitMythicalEndLord      int      `json:"longest_bow_kill_tourney_kit_mythical_end-lord,omitempty"`
				LongestBowKillTourneyTeamsTourney            int      `json:"longest_bow_kill_tourney_teams_tourney,omitempty"`
				LossesTourneyKitMythicalEndLord              int      `json:"losses_tourney_kit_mythical_end-lord,omitempty"`
				MeleeKillsTourney                            int      `json:"melee_kills_tourney,omitempty"`
				MeleeKillsTourneyKitMythicalEndLord          int      `json:"melee_kills_tourney_kit_mythical_end-lord,omitempty"`
				MeleeKillsTourneyTeamsTourney                int      `json:"melee_kills_tourney_teams_tourney,omitempty"`
				MostKillsGameTourneyKitMythicalEndLord       int      `json:"most_kills_game_tourney_kit_mythical_end-lord,omitempty"`
				SurvivedPlayersTourneyKitMythicalEndLord     int      `json:"survived_players_tourney_kit_mythical_end-lord,omitempty"`
				TimePlayedTourneyKitMythicalEndLord          int      `json:"time_played_tourney_kit_mythical_end-lord,omitempty"`
				TourneySwInsaneDoubles0EnderpearlsThrown     int      `json:"tourney_sw_insane_doubles_0_enderpearls_thrown,omitempty"`
				TourneySwInsaneDoubles0LongestBowKill        int      `json:"tourney_sw_insane_doubles_0_longest_bow_kill,omitempty"`
				TourneySwInsaneDoubles0MeleeKills            int      `json:"tourney_sw_insane_doubles_0_melee_kills,omitempty"`
				ArrowsShotTourneyKitMythicalEndLord          int      `json:"arrows_shot_tourney_kit_mythical_end-lord,omitempty"`
				ArrowsHitTourneyKitMythicalEndLord           int      `json:"arrows_hit_tourney_kit_mythical_end-lord,omitempty"`
				LongestBowShotTourneyKitMythicalEndLord      int      `json:"longest_bow_shot_tourney_kit_mythical_end-lord,omitempty"`
				AssistsTourneyKitMythicalEndLord             int      `json:"assists_tourney_kit_mythical_end-lord,omitempty"`
				LastTourneyAd                                int64    `json:"lastTourneyAd,omitempty"`
				CosmeticTokens                               int      `json:"cosmetic_tokens,omitempty"`
				FastestWin                                   int      `json:"fastest_win,omitempty"`
				FastestWinKitMiningTeamDefault               int      `json:"fastest_win_kit_mining_team_default,omitempty"`
				FastestWinSolo                               int      `json:"fastest_win_solo,omitempty"`
				Heads                                        int      `json:"heads,omitempty"`
				HeadsKitMiningTeamDefault                    int      `json:"heads_kit_mining_team_default,omitempty"`
				HeadsSalty                                   int      `json:"heads_salty,omitempty"`
				HeadsSaltyKitMiningTeamDefault               int      `json:"heads_salty_kit_mining_team_default,omitempty"`
				HeadsSaltySolo                               int      `json:"heads_salty_solo,omitempty"`
				HeadsSolo                                    int      `json:"heads_solo,omitempty"`
				KillstreakKitMiningTeamDefault               int      `json:"killstreak_kit_mining_team_default,omitempty"`
				KillstreakSolo                               int      `json:"killstreak_solo,omitempty"`
				MostKillsGameKitMiningTeamDefault            int      `json:"most_kills_game_kit_mining_team_default,omitempty"`
				MostKillsGameSolo                            int      `json:"most_kills_game_solo,omitempty"`
				VoidKillsKitMiningTeamDefault                int      `json:"void_kills_kit_mining_team_default,omitempty"`
				VoidKillsSolo                                int      `json:"void_kills_solo,omitempty"`
				WinsSoloInsane                               int      `json:"wins_solo_insane,omitempty"`
				HeadCollection                               struct {
					Recent []struct {
						UUID      string `json:"uuid,omitempty"`
						Timestamp int64  `json:"timestamp,omitempty"`
						Mode      string `json:"mode,omitempty"`
						Sacrifice string `json:"sacrifice,omitempty"`
					} `json:"recent,omitempty"`
					Prestigious []struct {
						UUID      string `json:"uuid,omitempty"`
						Timestamp int64  `json:"timestamp,omitempty"`
						Mode      string `json:"mode,omitempty"`
						Sacrifice string `json:"sacrifice,omitempty"`
					} `json:"prestigious,omitempty"`
				} `json:"head_collection,omitempty"`
				LongestBowKill                       int `json:"longest_bow_kill,omitempty"`
				LongestBowKillKitMiningTeamDefault   int `json:"longest_bow_kill_kit_mining_team_default,omitempty"`
				LongestBowKillSolo                   int `json:"longest_bow_kill_solo,omitempty"`
				MeleeKillsKitMiningTeamDefault       int `json:"melee_kills_kit_mining_team_default,omitempty"`
				MeleeKillsSolo                       int `json:"melee_kills_solo,omitempty"`
				ArrowsHitKitMiningTeamDefault        int `json:"arrows_hit_kit_mining_team_default,omitempty"`
				ArrowsHitSolo                        int `json:"arrows_hit_solo,omitempty"`
				LongestBowShot                       int `json:"longest_bow_shot,omitempty"`
				LongestBowShotKitMiningTeamDefault   int `json:"longest_bow_shot_kit_mining_team_default,omitempty"`
				LongestBowShotSolo                   int `json:"longest_bow_shot_solo,omitempty"`
				ArrowsShotKitAdvancedSoloEnchanter   int `json:"arrows_shot_kit_advanced_solo_enchanter,omitempty"`
				AssistsKitAdvancedSoloEnchanter      int `json:"assists_kit_advanced_solo_enchanter,omitempty"`
				ChestsOpenedKitAdvancedSoloEnchanter int `json:"chests_opened_kit_advanced_solo_enchanter,omitempty"`
				TimePlayedKitAdvancedSoloEnchanter   int `json:"time_played_kit_advanced_solo_enchanter,omitempty"`
				Perkslot                             struct {
					Normal struct {
						Num3 string `json:"3,omitempty"`
						Num4 string `json:"4,omitempty"`
						Num5 string `json:"5,omitempty"`
					} `json:"normal,omitempty"`
					Insane struct {
						Num3 string `json:"3,omitempty"`
						Num4 string `json:"4,omitempty"`
						Num5 string `json:"5,omitempty"`
					} `json:"insane,omitempty"`
				} `json:"perkslot,omitempty"`
				ToggleSoloBulldozer             bool   `json:"toggle_solo_bulldozer,omitempty"`
				SoloBulldozer                   int    `json:"solo_bulldozer,omitempty"`
				ToggleSoloArrowRecovery         bool   `json:"toggle_solo_arrow_recovery,omitempty"`
				ToggleSoloMarksmanship          bool   `json:"toggle_solo_marksmanship,omitempty"`
				ToggleSoloSpeedBoost            bool   `json:"toggle_solo_speed_boost,omitempty"`
				ToggleSoloAnnoyOMite            bool   `json:"toggle_solo_annoy-o-mite,omitempty"`
				ToggleSoloKnowledge             bool   `json:"toggle_solo_knowledge,omitempty"`
				ToggleSoloRevenge               bool   `json:"toggle_solo_revenge,omitempty"`
				ToggleSoloMiningExpertise       bool   `json:"toggle_solo_mining_expertise,omitempty"`
				SoloFat                         int    `json:"solo_fat,omitempty"`
				SoloSavior                      int    `json:"solo_savior,omitempty"`
				ToggleSoloFat                   bool   `json:"toggle_solo_fat,omitempty"`
				ToggleSoloEnvironmentalExpert   bool   `json:"toggle_solo_environmental_expert,omitempty"`
				ToggleSoloNourishment           bool   `json:"toggle_solo_nourishment,omitempty"`
				ToggleSoloBlackMagic            bool   `json:"toggle_solo_black_magic,omitempty"`
				ToggleSoloBridger               bool   `json:"toggle_solo_bridger,omitempty"`
				ToggleSoloRobbery               bool   `json:"toggle_solo_robbery,omitempty"`
				ToggleSoloBarbarian             bool   `json:"toggle_solo_barbarian,omitempty"`
				ToggleSoloFrost                 bool   `json:"toggle_solo_frost,omitempty"`
				ToggleMegaArrowRecovery         bool   `json:"toggle_mega_arrow_recovery,omitempty"`
				ToggleMegaMiningExpertise       bool   `json:"toggle_mega_mining_expertise,omitempty"`
				ToggleSoloLuckyCharm            bool   `json:"toggle_solo_lucky_charm,omitempty"`
				ToggleMegaBlazingArrows         bool   `json:"toggle_mega_blazing_arrows,omitempty"`
				ToggleSoloSavior                bool   `json:"toggle_solo_savior,omitempty"`
				ToggleMegaTank                  bool   `json:"toggle_mega_tank,omitempty"`
				ToggleMegaMarksmanship          bool   `json:"toggle_mega_marksmanship,omitempty"`
				ToggleMegaJuggernaut            bool   `json:"toggle_mega_juggernaut,omitempty"`
				ToggleSoloNecromancer           bool   `json:"toggle_solo_necromancer,omitempty"`
				ToggleMegaBridger               bool   `json:"toggle_mega_bridger,omitempty"`
				ToggleMegaNecromancer           bool   `json:"toggle_mega_necromancer,omitempty"`
				ToggleTeamArrowRecovery         bool   `json:"toggle_team_arrow_recovery,omitempty"`
				ToggleMegaEnvironmentalExpert   bool   `json:"toggle_mega_environmental_expert,omitempty"`
				ToggleMegaNourishment           bool   `json:"toggle_mega_nourishment,omitempty"`
				ToggleMegaLuckyCharm            bool   `json:"toggle_mega_lucky_charm,omitempty"`
				ToggleMegaBlackMagic            bool   `json:"toggle_mega_black_magic,omitempty"`
				ToggleTeamBulldozer             bool   `json:"toggle_team_bulldozer,omitempty"`
				TeamBulldozer                   int    `json:"team_bulldozer,omitempty"`
				ToggleMegaNotoriety             bool   `json:"toggle_mega_notoriety,omitempty"`
				ToggleTeamSpeedBoost            bool   `json:"toggle_team_speed_boost,omitempty"`
				ToggleTeamMiningExpertise       bool   `json:"toggle_team_mining_expertise,omitempty"`
				TeamJuggernaut                  int    `json:"team_juggernaut,omitempty"`
				ToggleTeamKnowledge             bool   `json:"toggle_team_knowledge,omitempty"`
				ToggleSoloJuggernaut            bool   `json:"toggle_solo_juggernaut,omitempty"`
				ToggleTeamNourishment           bool   `json:"toggle_team_nourishment,omitempty"`
				TeamFat                         int    `json:"team_fat,omitempty"`
				ToggleTeamFat                   bool   `json:"toggle_team_fat,omitempty"`
				ToggleTeamBridger               bool   `json:"toggle_team_bridger,omitempty"`
				ToggleTeamJuggernaut            bool   `json:"toggle_team_juggernaut,omitempty"`
				TeamSavior                      int    `json:"team_savior,omitempty"`
				ToggleTeamEnvironmentalExpert   bool   `json:"toggle_team_environmental_expert,omitempty"`
				ToggleTeamLuckyCharm            bool   `json:"toggle_team_lucky_charm,omitempty"`
				ToggleTeamNecromancer           bool   `json:"toggle_team_necromancer,omitempty"`
				ToggleTeamRobbery               bool   `json:"toggle_team_robbery,omitempty"`
				ToggleTeamBlackMagic            bool   `json:"toggle_team_black_magic,omitempty"`
				ToggleTeamBarbarian             bool   `json:"toggle_team_barbarian,omitempty"`
				ToggleRankedArmorerPerk         bool   `json:"toggle_ranked_armorer_perk,omitempty"`
				ToggleMegaRusher                bool   `json:"toggle_mega_rusher,omitempty"`
				ToggleRankedBowmanPerk          bool   `json:"toggle_ranked_bowman_perk,omitempty"`
				ToggleRankedChampionPerk        bool   `json:"toggle_ranked_champion_perk,omitempty"`
				ToggleTeamMarksmanship          bool   `json:"toggle_team_marksmanship,omitempty"`
				ToggleRankedAthletePerk         bool   `json:"toggle_ranked_athlete_perk,omitempty"`
				ToggleRankedMagicianPerk        bool   `json:"toggle_ranked_magician_perk,omitempty"`
				ToggleTeamBlazingArrows         bool   `json:"toggle_team_blazing_arrows,omitempty"`
				ToggleRankedHealerPerk          bool   `json:"toggle_ranked_healer_perk,omitempty"`
				ToggleTeamDiamondpiercer        bool   `json:"toggle_team_diamondpiercer,omitempty"`
				ToggleRankedPaladinPerk         bool   `json:"toggle_ranked_paladin_perk,omitempty"`
				ToggleRankedArrowRecovery       bool   `json:"toggle_ranked_arrow_recovery,omitempty"`
				ToggleRankedHoundPerk           bool   `json:"toggle_ranked_hound_perk,omitempty"`
				ToggleRankedBlazingArrows       bool   `json:"toggle_ranked_blazing_arrows,omitempty"`
				ToggleRankedLastStand           bool   `json:"toggle_ranked_last_stand,omitempty"`
				ToggleRankedJuggernaut          bool   `json:"toggle_ranked_juggernaut,omitempty"`
				ToggleTeamResistanceBoost       bool   `json:"toggle_team_resistance_boost,omitempty"`
				ToggleRankedBlacksmithPerk      bool   `json:"toggle_ranked_blacksmith_perk,omitempty"`
				ToggleRankedMiningExpertise     bool   `json:"toggle_ranked_mining_expertise,omitempty"`
				ToggleRankedToughSkin           bool   `json:"toggle_ranked_tough_skin,omitempty"`
				ToggleRankedEnvironmentalExpert bool   `json:"toggle_ranked_environmental_expert,omitempty"`
				ToggleRankedBridger             bool   `json:"toggle_ranked_bridger,omitempty"`
				ToggleRankedScoutPerk           bool   `json:"toggle_ranked_scout_perk,omitempty"`
				ToggleTeamFrost                 bool   `json:"toggle_team_frost,omitempty"`
				ToggleRankedPyromancerPerk      bool   `json:"toggle_ranked_pyromancer_perk,omitempty"`
				ToggleRankedRusher              bool   `json:"toggle_ranked_rusher,omitempty"`
				ToggleTeamSavior                bool   `json:"toggle_team_savior,omitempty"`
				ToggleTeamAnnoyOMite            bool   `json:"toggle_team_annoy-o-mite,omitempty"`
				SoloJuggernaut                  int    `json:"solo_juggernaut,omitempty"`
				ToggleSoloBlazingArrows         bool   `json:"toggle_solo_blazing_arrows,omitempty"`
				ToggleSoloResistanceBoost       bool   `json:"toggle_solo_resistance_boost,omitempty"`
				LevelFormattedWithBrackets      string `json:"levelFormattedWithBrackets,omitempty"`
			} `json:"SkyWars,omitempty"`
			TrueCombat struct {
				Packages                    []string `json:"packages,omitempty"`
				WinStreak                   int      `json:"win_streak,omitempty"`
				Games                       int      `json:"games,omitempty"`
				CrazywallsLossesSoloChaos   int      `json:"crazywalls_losses_solo_chaos,omitempty"`
				CrazywallsGamesSoloChaos    int      `json:"crazywalls_games_solo_chaos,omitempty"`
				Deaths                      int      `json:"deaths,omitempty"`
				Coins                       int      `json:"coins,omitempty"`
				Losses                      int      `json:"losses,omitempty"`
				CrazywallsDeathsSoloChaos   int      `json:"crazywalls_deaths_solo_chaos,omitempty"`
				SurvivedPlayers             int      `json:"survived_players,omitempty"`
				ItemsEnchanted              int      `json:"items_enchanted,omitempty"`
				ArrowsShot                  int      `json:"arrows_shot,omitempty"`
				CrazywallsDeathsSolo        int      `json:"crazywalls_deaths_solo,omitempty"`
				Kills                       int      `json:"kills,omitempty"`
				CrazywallsGamesSolo         int      `json:"crazywalls_games_solo,omitempty"`
				CrazywallsKillsMonthlyBSolo int      `json:"crazywalls_kills_monthly_b_solo,omitempty"`
				CrazywallsLossesSolo        int      `json:"crazywalls_losses_solo,omitempty"`
				CrazywallsKillsSolo         int      `json:"crazywalls_kills_solo,omitempty"`
				CrazywallsKillsWeeklyASolo  int      `json:"crazywalls_kills_weekly_a_solo,omitempty"`
			} `json:"TrueCombat,omitempty"`
			SuperSmash struct {
				LastLevelTHEBULK int `json:"lastLevel_THE_BULK,omitempty"`
				SmashLevel       int `json:"smashLevel,omitempty"`
				WinStreak        int `json:"win_streak,omitempty"`
				ClassStats       struct {
					TheBulk struct {
						SeismicSlam struct {
							DamageDealt       int `json:"damage_dealt,omitempty"`
							DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
							Kills             int `json:"kills,omitempty"`
							KillsNormal       int `json:"kills_normal,omitempty"`
							KillsTeams        int `json:"kills_teams,omitempty"`
							DamageDealtTeams  int `json:"damage_dealt_teams,omitempty"`
						} `json:"seismic_slam,omitempty"`
						MonsterMash struct {
							DamageDealt       int `json:"damage_dealt,omitempty"`
							DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
							KillsNormal       int `json:"kills_normal,omitempty"`
							Kills             int `json:"kills,omitempty"`
							DamageDealtTeams  int `json:"damage_dealt_teams,omitempty"`
							Smasher           int `json:"smasher,omitempty"`
							SmasherNormal     int `json:"smasher_normal,omitempty"`
						} `json:"monster_mash,omitempty"`
						Smashed       int `json:"smashed,omitempty"`
						MonsterCharge struct {
							DamageDealt       int `json:"damage_dealt,omitempty"`
							DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
							DamageDealtTeams  int `json:"damage_dealt_teams,omitempty"`
							Smashed           int `json:"smashed,omitempty"`
							SmashedNormal     int `json:"smashed_normal,omitempty"`
							Kills             int `json:"kills,omitempty"`
							KillsNormal       int `json:"kills_normal,omitempty"`
						} `json:"monster_charge,omitempty"`
						Melee struct {
							DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
							Smashed           int `json:"smashed,omitempty"`
							SmashedNormal     int `json:"smashed_normal,omitempty"`
							DamageDealt       int `json:"damage_dealt,omitempty"`
							Kills             int `json:"kills,omitempty"`
							KillsNormal       int `json:"kills_normal,omitempty"`
						} `json:"melee,omitempty"`
						Games             int `json:"games,omitempty"`
						SmashedNormal     int `json:"smashed_normal,omitempty"`
						DeathsNormal      int `json:"deaths_normal,omitempty"`
						LossesNormal      int `json:"losses_normal,omitempty"`
						DamageDealt       int `json:"damage_dealt,omitempty"`
						Losses            int `json:"losses,omitempty"`
						Deaths            int `json:"deaths,omitempty"`
						GamesNormal       int `json:"games_normal,omitempty"`
						DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
						Boulder           struct {
							Kills             int `json:"kills,omitempty"`
							DamageDealt       int `json:"damage_dealt,omitempty"`
							DamageDealtNormal int `json:"damage_dealt_normal,omitempty"`
							KillsNormal       int `json:"kills_normal,omitempty"`
							DamageDealtTeams  int `json:"damage_dealt_teams,omitempty"`
						} `json:"boulder,omitempty"`
						Batarang struct {
							SmashedNormal int `json:"smashed_normal,omitempty"`
							Smashed       int `json:"smashed,omitempty"`
						} `json:"batarang,omitempty"`
						Kills            int `json:"kills,omitempty"`
						KillsNormal      int `json:"kills_normal,omitempty"`
						DamageDealtTeams int `json:"damage_dealt_teams,omitempty"`
						DeathsTeams      int `json:"deaths_teams,omitempty"`
						GamesTeams       int `json:"games_teams,omitempty"`
						LossesTeams      int `json:"losses_teams,omitempty"`
						KillsTeams       int `json:"kills_teams,omitempty"`
						Losses2V2        int `json:"losses_2v2,omitempty"`
						Games2V2         int `json:"games_2v2,omitempty"`
						ForceLightning   struct {
							Smashed       int `json:"smashed,omitempty"`
							SmashedNormal int `json:"smashed_normal,omitempty"`
						} `json:"force_lightning,omitempty"`
						Smasher       int `json:"smasher,omitempty"`
						SmasherNormal int `json:"smasher_normal,omitempty"`
					} `json:"THE_BULK,omitempty"`
				} `json:"class_stats,omitempty"`
				DamageDealt       int     `json:"damage_dealt,omitempty"`
				LossesWeeklyA     int     `json:"losses_weekly_a,omitempty"`
				GamesWeeklyA      int     `json:"games_weekly_a,omitempty"`
				LossesMonthlyB    int     `json:"losses_monthly_b,omitempty"`
				Losses            int     `json:"losses,omitempty"`
				SmashedNormal     int     `json:"smashed_normal,omitempty"`
				GamesMonthlyB     int     `json:"games_monthly_b,omitempty"`
				Quits             int     `json:"quits,omitempty"`
				DeathsNormal      int     `json:"deaths_normal,omitempty"`
				GamesNormal       int     `json:"games_normal,omitempty"`
				LossesNormal      int     `json:"losses_normal,omitempty"`
				Smashed           int     `json:"smashed,omitempty"`
				Deaths            int     `json:"deaths,omitempty"`
				Games             int     `json:"games,omitempty"`
				DamageDealtNormal int     `json:"damage_dealt_normal,omitempty"`
				Coins             int     `json:"coins,omitempty"`
				KillsMonthlyB     int     `json:"kills_monthly_b,omitempty"`
				Kills             int     `json:"kills,omitempty"`
				KillsWeeklyA      int     `json:"kills_weekly_a,omitempty"`
				KillsNormal       int     `json:"kills_normal,omitempty"`
				XpTHEBULK         float64 `json:"xp_THE_BULK,omitempty"`
				VotesLuxor        int     `json:"votes_Luxor,omitempty"`
				SmashLevelTotal   int     `json:"smash_level_total,omitempty"`
				DamageDealtTeams  int     `json:"damage_dealt_teams,omitempty"`
				LossesTeams       int     `json:"losses_teams,omitempty"`
				KillsTeams        int     `json:"kills_teams,omitempty"`
				DeathsTeams       int     `json:"deaths_teams,omitempty"`
				GamesTeams        int     `json:"games_teams,omitempty"`
				Losses2V2         int     `json:"losses_2v2,omitempty"`
				LossesWeeklyB     int     `json:"losses_weekly_b,omitempty"`
				Games2V2          int     `json:"games_2v2,omitempty"`
				GamesWeeklyB      int     `json:"games_weekly_b,omitempty"`
				SmashPlayStreak   int     `json:"smash_play_streak,omitempty"`
				Smasher           int     `json:"smasher,omitempty"`
				SmasherNormal     int     `json:"smasher_normal,omitempty"`
			} `json:"SuperSmash,omitempty"`
			SpeedUHC struct {
				Coins                                int  `json:"coins,omitempty"`
				FirstJoinLobbyInt                    int  `json:"firstJoinLobbyInt,omitempty"`
				Score                                int  `json:"score,omitempty"`
				MovedOver                            bool `json:"movedOver,omitempty"`
				Killstreak                           int  `json:"killstreak,omitempty"`
				BlocksBroken                         int  `json:"blocks_broken,omitempty"`
				BlocksPlaced                         int  `json:"blocks_placed,omitempty"`
				Deaths                               int  `json:"deaths,omitempty"`
				DeathsKitBasicNormalDefault          int  `json:"deaths_kit_basic_normal_default,omitempty"`
				DeathsMasteryWildSpecialist          int  `json:"deaths_mastery_wild_specialist,omitempty"`
				DeathsNormal                         int  `json:"deaths_normal,omitempty"`
				DeathsSolo                           int  `json:"deaths_solo,omitempty"`
				DeathsSoloNormal                     int  `json:"deaths_solo_normal,omitempty"`
				Losses                               int  `json:"losses,omitempty"`
				LossesKitBasicNormalDefault          int  `json:"losses_kit_basic_normal_default,omitempty"`
				LossesMasteryWildSpecialist          int  `json:"losses_mastery_wild_specialist,omitempty"`
				LossesNormal                         int  `json:"losses_normal,omitempty"`
				LossesSolo                           int  `json:"losses_solo,omitempty"`
				LossesSoloNormal                     int  `json:"losses_solo_normal,omitempty"`
				Quits                                int  `json:"quits,omitempty"`
				SurvivedPlayers                      int  `json:"survived_players,omitempty"`
				SurvivedPlayersKitBasicNormalDefault int  `json:"survived_players_kit_basic_normal_default,omitempty"`
				SurvivedPlayersNormal                int  `json:"survived_players_normal,omitempty"`
				SurvivedPlayersSolo                  int  `json:"survived_players_solo,omitempty"`
				WinStreak                            int  `json:"win_streak,omitempty"`
			} `json:"SpeedUHC,omitempty"`
			SkyClash struct {
				CardPacks                  int      `json:"card_packs,omitempty"`
				PerkHeartyStartDuplicates  int      `json:"perk_hearty_start_duplicates,omitempty"`
				PerkHitAndRunDuplicates    int      `json:"perk_hit_and_run_duplicates,omitempty"`
				Packages                   []string `json:"packages,omitempty"`
				PerkVoidMagnetNew          bool     `json:"perk_void_magnet_new,omitempty"`
				PerkVoidMagnet             int      `json:"perk_void_magnet,omitempty"`
				Playstreak                 int      `json:"playstreak,omitempty"`
				ActiveClass                int      `json:"active_class,omitempty"`
				WinStreak                  int      `json:"win_streak,omitempty"`
				LossesSolo                 int      `json:"losses_solo,omitempty"`
				DeathsKitGuardian          int      `json:"deaths_kit_guardian,omitempty"`
				Losses                     int      `json:"losses,omitempty"`
				GamesPlayedKitGuardian     int      `json:"games_played_kit_guardian,omitempty"`
				Deaths                     int      `json:"deaths,omitempty"`
				DeathsSolo                 int      `json:"deaths_solo,omitempty"`
				Coins                      int      `json:"coins,omitempty"`
				GamesPlayed                int      `json:"games_played,omitempty"`
				Quits                      int      `json:"quits,omitempty"`
				PlayStreak                 int      `json:"play_streak,omitempty"`
				LongestBowShotKitArcher    int      `json:"longest_bow_shot_kit_archer,omitempty"`
				LongestBowShot             int      `json:"longest_bow_shot,omitempty"`
				GamesPlayedKitArcher       int      `json:"games_played_kit_archer,omitempty"`
				BowHitsKitArcher           int      `json:"bow_hits_kit_archer,omitempty"`
				BowHits                    int      `json:"bow_hits,omitempty"`
				Games                      int      `json:"games,omitempty"`
				BowShots                   int      `json:"bow_shots,omitempty"`
				BowShotsKitArcher          int      `json:"bow_shots_kit_archer,omitempty"`
				DeathsKitArcher            int      `json:"deaths_kit_archer,omitempty"`
				Killstreak                 int      `json:"killstreak,omitempty"`
				MobsKilled                 int      `json:"mobs_killed,omitempty"`
				EnderchestsOpened          int      `json:"enderchests_opened,omitempty"`
				MobsKilledKitArcher        int      `json:"mobs_killed_kit_archer,omitempty"`
				EnderchestsOpenedKitArcher int      `json:"enderchests_opened_kit_archer,omitempty"`
				HighestKillstreak          int      `json:"highestKillstreak,omitempty"`
				Kills                      int      `json:"kills,omitempty"`
				KillsKitArcher             int      `json:"kills_kit_archer,omitempty"`
				MostKillsGame              int      `json:"most_kills_game,omitempty"`
				MeleeKillsKitArcher        int      `json:"melee_kills_kit_archer,omitempty"`
				MeleeKills                 int      `json:"melee_kills,omitempty"`
				KillsSolo                  int      `json:"kills_solo,omitempty"`
				KillsMonthlyB              int      `json:"kills_monthly_b,omitempty"`
				MostKillsGameKitArcher     int      `json:"most_kills_game_kit_archer,omitempty"`
				KillsWeeklyA               int      `json:"kills_weekly_a,omitempty"`
				VoidKillsKitArcher         int      `json:"void_kills_kit_archer,omitempty"`
				VoidKills                  int      `json:"void_kills,omitempty"`
				LongestBowKillKitArcher    int      `json:"longest_bow_kill_kit_archer,omitempty"`
				LongestBowKill             int      `json:"longest_bow_kill,omitempty"`
				BowKillsKitArcher          int      `json:"bow_kills_kit_archer,omitempty"`
				BowKills                   int      `json:"bow_kills,omitempty"`
			} `json:"SkyClash,omitempty"`
			Legacy struct {
				NextTokensSeconds   int    `json:"next_tokens_seconds,omitempty"`
				QuakecraftTokens    int    `json:"quakecraft_tokens,omitempty"`
				TotalTokens         int    `json:"total_tokens,omitempty"`
				Tokens              int    `json:"tokens,omitempty"`
				PaintballTokens     int    `json:"paintball_tokens,omitempty"`
				ArenaTokens         int    `json:"arena_tokens,omitempty"`
				PreferredChannel    string `json:"preferredChannel,omitempty"`
				Speed               string `json:"speed,omitempty"`
				LeaderboardSettings struct {
					ResetType      string `json:"resetType,omitempty"`
					QuakecraftMode string `json:"quakecraftMode,omitempty"`
				} `json:"leaderboardSettings,omitempty"`
				VampirezTokens    int `json:"vampirez_tokens,omitempty"`
				GingerbreadTokens int `json:"gingerbread_tokens,omitempty"`
				WallsTokens       int `json:"walls_tokens,omitempty"`
			} `json:"Legacy,omitempty"`
			Bedwars struct {
				FirstJoin7                                             bool     `json:"first_join_7,omitempty"`
				BedwarsBoxes                                           int      `json:"bedwars_boxes,omitempty"`
				ExperienceNew                                          int      `json:"Experience_new,omitempty"`
				GamesPlayedBedwars1                                    int      `json:"games_played_bedwars_1,omitempty"`
				Winstreak                                              int      `json:"winstreak,omitempty"`
				GoldResourcesCollectedBedwars                          int      `json:"gold_resources_collected_bedwars,omitempty"`
				VoidDeathsBedwars                                      int      `json:"void_deaths_bedwars,omitempty"`
				EightTwoVoidDeathsBedwars                              int      `json:"eight_two_void_deaths_bedwars,omitempty"`
				EightTwoBedsBrokenBedwars                              int      `json:"eight_two_beds_broken_bedwars,omitempty"`
				DiamondResourcesCollectedBedwars                       int      `json:"diamond_resources_collected_bedwars,omitempty"`
				DeathsBedwars                                          int      `json:"deaths_bedwars,omitempty"`
				EmeraldResourcesCollectedBedwars                       int      `json:"emerald_resources_collected_bedwars,omitempty"`
				ResourcesCollectedBedwars                              int      `json:"resources_collected_bedwars,omitempty"`
				EightTwoEmeraldResourcesCollectedBedwars               int      `json:"eight_two_emerald_resources_collected_bedwars,omitempty"`
				EightTwoPermanentItemsPurchasedBedwars                 int      `json:"eight_two_permanent _items_purchased_bedwars,omitempty"`
				EightTwoResourcesCollectedBedwars                      int      `json:"eight_two_resources_collected_bedwars,omitempty"`
				EightTwoDiamondResourcesCollectedBedwars               int      `json:"eight_two_diamond_resources_collected_bedwars,omitempty"`
				EightTwoWinsBedwars                                    int      `json:"eight_two_wins_bedwars,omitempty"`
				Coins                                                  int      `json:"coins,omitempty"`
				GamesPlayedBedwars                                     int      `json:"games_played_bedwars,omitempty"`
				PermanentItemsPurchasedBedwars                         int      `json:"permanent _items_purchased_bedwars,omitempty"`
				EightTwoIronResourcesCollectedBedwars                  int      `json:"eight_two_iron_resources_collected_bedwars,omitempty"`
				EightTwoDeathsBedwars                                  int      `json:"eight_two_deaths_bedwars,omitempty"`
				EightTwoItemsPurchasedBedwars                          int      `json:"eight_two__items_purchased_bedwars,omitempty"`
				EightTwoGamesPlayedBedwars                             int      `json:"eight_two_games_played_bedwars,omitempty"`
				EightTwoItemsPurchasedBedwars0                         int      `json:"eight_two_items_purchased_bedwars,omitempty"`
				EightTwoEntityAttackDeathsBedwars                      int      `json:"eight_two_entity_attack_deaths_bedwars,omitempty"`
				EightTwoGoldResourcesCollectedBedwars                  int      `json:"eight_two_gold_resources_collected_bedwars,omitempty"`
				EntityAttackDeathsBedwars                              int      `json:"entity_attack_deaths_bedwars,omitempty"`
				EightTwoVoidFinalKillsBedwars                          int      `json:"eight_two_void_final_kills_bedwars,omitempty"`
				ItemsPurchasedBedwars                                  int      `json:"items_purchased_bedwars,omitempty"`
				WinsBedwars                                            int      `json:"wins_bedwars,omitempty"`
				FinalKillsBedwars                                      int      `json:"final_kills_bedwars,omitempty"`
				VoidFinalKillsBedwars                                  int      `json:"void_final_kills_bedwars,omitempty"`
				IronResourcesCollectedBedwars                          int      `json:"iron_resources_collected_bedwars,omitempty"`
				BedsBrokenBedwars                                      int      `json:"beds_broken_bedwars,omitempty"`
				EightTwoFinalKillsBedwars                              int      `json:"eight_two_final_kills_bedwars,omitempty"`
				ItemsPurchasedBedwars0                                 int      `json:"_items_purchased_bedwars,omitempty"`
				FinalDeathsBedwars                                     int      `json:"final_deaths_bedwars,omitempty"`
				VoidKillsBedwars                                       int      `json:"void_kills_bedwars,omitempty"`
				EntityAttackFinalDeathsBedwars                         int      `json:"entity_attack_final_deaths_bedwars,omitempty"`
				EightTwoKillsBedwars                                   int      `json:"eight_two_kills_bedwars,omitempty"`
				EightTwoEntityAttackFinalKillsBedwars                  int      `json:"eight_two_entity_attack_final_kills_bedwars,omitempty"`
				EntityAttackFinalKillsBedwars                          int      `json:"entity_attack_final_kills_bedwars,omitempty"`
				EightTwoEntityAttackKillsBedwars                       int      `json:"eight_two_entity_attack_kills_bedwars,omitempty"`
				BedsLostBedwars                                        int      `json:"beds_lost_bedwars,omitempty"`
				KillsBedwars                                           int      `json:"kills_bedwars,omitempty"`
				EightTwoVoidKillsBedwars                               int      `json:"eight_two_void_kills_bedwars,omitempty"`
				EightTwoFinalDeathsBedwars                             int      `json:"eight_two_final_deaths_bedwars,omitempty"`
				EntityAttackKillsBedwars                               int      `json:"entity_attack_kills_bedwars,omitempty"`
				EightTwoEntityAttackFinalDeathsBedwars                 int      `json:"eight_two_entity_attack_final_deaths_bedwars,omitempty"`
				EightTwoBedsLostBedwars                                int      `json:"eight_two_beds_lost_bedwars,omitempty"`
				LossesBedwars                                          int      `json:"losses_bedwars,omitempty"`
				EightTwoLossesBedwars                                  int      `json:"eight_two_losses_bedwars,omitempty"`
				BedwarsBoxCommons                                      int      `json:"bedwars_box_commons,omitempty"`
				BedwarsBox                                             int      `json:"bedwars_box,omitempty"`
				ChestHistory                                           string   `json:"chest_history,omitempty"`
				BedwarsBoxRares                                        int      `json:"bedwars_box_rares,omitempty"`
				Packages                                               []string `json:"packages,omitempty"`
				ActiveNPCSkin                                          string   `json:"activeNPCSkin,omitempty"`
				ActiveIslandTopper                                     string   `json:"activeIslandTopper,omitempty"`
				ActiveDeathCry                                         string   `json:"activeDeathCry,omitempty"`
				Experience                                             float64  `json:"Experience,omitempty"`
				EightTwoVoidFinalDeathsBedwars                         int      `json:"eight_two_void_final_deaths_bedwars,omitempty"`
				VoidFinalDeathsBedwars                                 int      `json:"void_final_deaths_bedwars,omitempty"`
				FallDeathsBedwars                                      int      `json:"fall_deaths_bedwars,omitempty"`
				EightTwoFallDeathsBedwars                              int      `json:"eight_two_fall_deaths_bedwars,omitempty"`
				FallKillsBedwars                                       int      `json:"fall_kills_bedwars,omitempty"`
				EightTwoFallKillsBedwars                               int      `json:"eight_two_fall_kills_bedwars,omitempty"`
				EightOneEntityAttackFinalKillsBedwars                  int      `json:"eight_one_entity_attack_final_kills_bedwars,omitempty"`
				EightOneDeathsBedwars                                  int      `json:"eight_one_deaths_bedwars,omitempty"`
				EightOneVoidDeathsBedwars                              int      `json:"eight_one_void_deaths_bedwars,omitempty"`
				EightOneGamesPlayedBedwars                             int      `json:"eight_one_games_played_bedwars,omitempty"`
				EightOneEntityAttackDeathsBedwars                      int      `json:"eight_one_entity_attack_deaths_bedwars,omitempty"`
				EightOneFinalKillsBedwars                              int      `json:"eight_one_final_kills_bedwars,omitempty"`
				EightOneBedsBrokenBedwars                              int      `json:"eight_one_beds_broken_bedwars,omitempty"`
				EightOneIronResourcesCollectedBedwars                  int      `json:"eight_one_iron_resources_collected_bedwars,omitempty"`
				EightOneItemsPurchasedBedwars                          int      `json:"eight_one__items_purchased_bedwars,omitempty"`
				EightOneEmeraldResourcesCollectedBedwars               int      `json:"eight_one_emerald_resources_collected_bedwars,omitempty"`
				EightOneItemsPurchasedBedwars0                         int      `json:"eight_one_items_purchased_bedwars,omitempty"`
				EightOneDiamondResourcesCollectedBedwars               int      `json:"eight_one_diamond_resources_collected_bedwars,omitempty"`
				EightOnePermanentItemsPurchasedBedwars                 int      `json:"eight_one_permanent _items_purchased_bedwars,omitempty"`
				EightOneVoidKillsBedwars                               int      `json:"eight_one_void_kills_bedwars,omitempty"`
				EightOneWinsBedwars                                    int      `json:"eight_one_wins_bedwars,omitempty"`
				EightOneKillsBedwars                                   int      `json:"eight_one_kills_bedwars,omitempty"`
				EightOneResourcesCollectedBedwars                      int      `json:"eight_one_resources_collected_bedwars,omitempty"`
				EightOneGoldResourcesCollectedBedwars                  int      `json:"eight_one_gold_resources_collected_bedwars,omitempty"`
				ActiveKillEffect                                       string   `json:"activeKillEffect,omitempty"`
				SprayStorageNew                                        string   `json:"spray_storage_new,omitempty"`
				EightOneBedsLostBedwars                                int      `json:"eight_one_beds_lost_bedwars,omitempty"`
				EightOneFinalDeathsBedwars                             int      `json:"eight_one_final_deaths_bedwars,omitempty"`
				EightOneLossesBedwars                                  int      `json:"eight_one_losses_bedwars,omitempty"`
				EightOneVoidFinalDeathsBedwars                         int      `json:"eight_one_void_final_deaths_bedwars,omitempty"`
				FourFourItemsPurchasedBedwars                          int      `json:"four_four__items_purchased_bedwars,omitempty"`
				FourFourPermanentItemsPurchasedBedwars                 int      `json:"four_four_permanent _items_purchased_bedwars,omitempty"`
				FourFourVoidKillsBedwars                               int      `json:"four_four_void_kills_bedwars,omitempty"`
				FourFourKillsBedwars                                   int      `json:"four_four_kills_bedwars,omitempty"`
				FourFourFallDeathsBedwars                              int      `json:"four_four_fall_deaths_bedwars,omitempty"`
				FourFourFinalDeathsBedwars                             int      `json:"four_four_final_deaths_bedwars,omitempty"`
				FourFourDeathsBedwars                                  int      `json:"four_four_deaths_bedwars,omitempty"`
				FourFourGoldResourcesCollectedBedwars                  int      `json:"four_four_gold_resources_collected_bedwars,omitempty"`
				FourFourDiamondResourcesCollectedBedwars               int      `json:"four_four_diamond_resources_collected_bedwars,omitempty"`
				FourFourEntityAttackDeathsBedwars                      int      `json:"four_four_entity_attack_deaths_bedwars,omitempty"`
				FourFourVoidFinalDeathsBedwars                         int      `json:"four_four_void_final_deaths_bedwars,omitempty"`
				FourFourBedsLostBedwars                                int      `json:"four_four_beds_lost_bedwars,omitempty"`
				FourFourLossesBedwars                                  int      `json:"four_four_losses_bedwars,omitempty"`
				FourFourIronResourcesCollectedBedwars                  int      `json:"four_four_iron_resources_collected_bedwars,omitempty"`
				FourFourGamesPlayedBedwars                             int      `json:"four_four_games_played_bedwars,omitempty"`
				FourFourFallKillsBedwars                               int      `json:"four_four_fall_kills_bedwars,omitempty"`
				FourFourResourcesCollectedBedwars                      int      `json:"four_four_resources_collected_bedwars,omitempty"`
				FourFourVoidDeathsBedwars                              int      `json:"four_four_void_deaths_bedwars,omitempty"`
				FourFourItemsPurchasedBedwars0                         int      `json:"four_four_items_purchased_bedwars,omitempty"`
				FourFourEntityAttackKillsBedwars                       int      `json:"four_four_entity_attack_kills_bedwars,omitempty"`
				FourFourEmeraldResourcesCollectedBedwars               int      `json:"four_four_emerald_resources_collected_bedwars,omitempty"`
				FourFourBedsBrokenBedwars                              int      `json:"four_four_beds_broken_bedwars,omitempty"`
				FourFourWinsBedwars                                    int      `json:"four_four_wins_bedwars,omitempty"`
				FourFourVoidFinalKillsBedwars                          int      `json:"four_four_void_final_kills_bedwars,omitempty"`
				FourFourEntityAttackFinalKillsBedwars                  int      `json:"four_four_entity_attack_final_kills_bedwars,omitempty"`
				FourFourFinalKillsBedwars                              int      `json:"four_four_final_kills_bedwars,omitempty"`
				FourFourProjectileDeathsBedwars                        int      `json:"four_four_projectile_deaths_bedwars,omitempty"`
				ProjectileDeathsBedwars                                int      `json:"projectile_deaths_bedwars,omitempty"`
				FallFinalKillsBedwars                                  int      `json:"fall_final_kills_bedwars,omitempty"`
				FourFourFallFinalKillsBedwars                          int      `json:"four_four_fall_final_kills_bedwars,omitempty"`
				FourFourEntityAttackFinalDeathsBedwars                 int      `json:"four_four_entity_attack_final_deaths_bedwars,omitempty"`
				FallFinalDeathsBedwars                                 int      `json:"fall_final_deaths_bedwars,omitempty"`
				FourFourFallFinalDeathsBedwars                         int      `json:"four_four_fall_final_deaths_bedwars,omitempty"`
				EightOneVoidFinalKillsBedwars                          int      `json:"eight_one_void_final_kills_bedwars,omitempty"`
				FourThreeVoidKillsBedwars                              int      `json:"four_three_void_kills_bedwars,omitempty"`
				FourThreeWinsBedwars                                   int      `json:"four_three_wins_bedwars,omitempty"`
				FourThreeIronResourcesCollectedBedwars                 int      `json:"four_three_iron_resources_collected_bedwars,omitempty"`
				FourThreeGamesPlayedBedwars                            int      `json:"four_three_games_played_bedwars,omitempty"`
				FourThreeItemsPurchasedBedwars                         int      `json:"four_three__items_purchased_bedwars,omitempty"`
				FourThreeVoidDeathsBedwars                             int      `json:"four_three_void_deaths_bedwars,omitempty"`
				FourThreeEntityAttackDeathsBedwars                     int      `json:"four_three_entity_attack_deaths_bedwars,omitempty"`
				FourThreeGoldResourcesCollectedBedwars                 int      `json:"four_three_gold_resources_collected_bedwars,omitempty"`
				FourThreeDeathsBedwars                                 int      `json:"four_three_deaths_bedwars,omitempty"`
				FourThreeDiamondResourcesCollectedBedwars              int      `json:"four_three_diamond_resources_collected_bedwars,omitempty"`
				FourThreeItemsPurchasedBedwars0                        int      `json:"four_three_items_purchased_bedwars,omitempty"`
				FourThreeKillsBedwars                                  int      `json:"four_three_kills_bedwars,omitempty"`
				FourThreeResourcesCollectedBedwars                     int      `json:"four_three_resources_collected_bedwars,omitempty"`
				ActiveKillMessages                                     string   `json:"activeKillMessages,omitempty"`
				BedwarsBoxEpics                                        int      `json:"bedwars_box_epics,omitempty"`
				ActiveProjectileTrail                                  string   `json:"activeProjectileTrail,omitempty"`
				GlyphStorageNew                                        string   `json:"glyph_storage_new,omitempty"`
				FourFourProjectileKillsBedwars                         int      `json:"four_four_projectile_kills_bedwars,omitempty"`
				ProjectileKillsBedwars                                 int      `json:"projectile_kills_bedwars,omitempty"`
				FourThreeEmeraldResourcesCollectedBedwars              int      `json:"four_three_emerald_resources_collected_bedwars,omitempty"`
				FourThreeBedsBrokenBedwars                             int      `json:"four_three_beds_broken_bedwars,omitempty"`
				FourThreeEntityAttackFinalKillsBedwars                 int      `json:"four_three_entity_attack_final_kills_bedwars,omitempty"`
				FourThreeFinalKillsBedwars                             int      `json:"four_three_final_kills_bedwars,omitempty"`
				FourThreeVoidFinalKillsBedwars                         int      `json:"four_three_void_final_kills_bedwars,omitempty"`
				FourThreeEntityAttackKillsBedwars                      int      `json:"four_three_entity_attack_kills_bedwars,omitempty"`
				FourThreeEntityAttackFinalDeathsBedwars                int      `json:"four_three_entity_attack_final_deaths_bedwars,omitempty"`
				FourThreeFallDeathsBedwars                             int      `json:"four_three_fall_deaths_bedwars,omitempty"`
				FourThreeFallFinalKillsBedwars                         int      `json:"four_three_fall_final_kills_bedwars,omitempty"`
				FourThreeBedsLostBedwars                               int      `json:"four_three_beds_lost_bedwars,omitempty"`
				FourThreePermanentItemsPurchasedBedwars                int      `json:"four_three_permanent _items_purchased_bedwars,omitempty"`
				FourThreeLossesBedwars                                 int      `json:"four_three_losses_bedwars,omitempty"`
				FourThreeFinalDeathsBedwars                            int      `json:"four_three_final_deaths_bedwars,omitempty"`
				FourThreeVoidFinalDeathsBedwars                        int      `json:"four_three_void_final_deaths_bedwars,omitempty"`
				EightOneFallDeathsBedwars                              int      `json:"eight_one_fall_deaths_bedwars,omitempty"`
				EightOneEntityAttackKillsBedwars                       int      `json:"eight_one_entity_attack_kills_bedwars,omitempty"`
				EntityExplosionDeathsBedwars                           int      `json:"entity_explosion_deaths_bedwars,omitempty"`
				FourFourEntityExplosionDeathsBedwars                   int      `json:"four_four_entity_explosion_deaths_bedwars,omitempty"`
				SprayGlyphField                                        string   `json:"spray_glyph_field,omitempty"`
				FourFourEntityExplosionKillsBedwars                    int      `json:"four_four_entity_explosion_kills_bedwars,omitempty"`
				EntityExplosionKillsBedwars                            int      `json:"entity_explosion_kills_bedwars,omitempty"`
				EntityExplosionFinalKillsBedwars                       int      `json:"entity_explosion_final_kills_bedwars,omitempty"`
				FourFourEntityExplosionFinalKillsBedwars               int      `json:"four_four_entity_explosion_final_kills_bedwars,omitempty"`
				EightTwoFallFinalKillsBedwars                          int      `json:"eight_two_fall_final_kills_bedwars,omitempty"`
				ActiveVictoryDance                                     string   `json:"activeVictoryDance,omitempty"`
				FourFourFireTickDeathsBedwars                          int      `json:"four_four_fire_tick_deaths_bedwars,omitempty"`
				FireTickDeathsBedwars                                  int      `json:"fire_tick_deaths_bedwars,omitempty"`
				ChestHistoryNew                                        []string `json:"chest_history_new,omitempty"`
				FourFourProjectileFinalKillsBedwars                    int      `json:"four_four_projectile_final_kills_bedwars,omitempty"`
				ProjectileFinalKillsBedwars                            int      `json:"projectile_final_kills_bedwars,omitempty"`
				EightTwoEntityExplosionDeathsBedwars                   int      `json:"eight_two_entity_explosion_deaths_bedwars,omitempty"`
				FourThreeProjectileKillsBedwars                        int      `json:"four_three_projectile_kills_bedwars,omitempty"`
				FourFourProjectileFinalDeathsBedwars                   int      `json:"four_four_projectile_final_deaths_bedwars,omitempty"`
				ProjectileFinalDeathsBedwars                           int      `json:"projectile_final_deaths_bedwars,omitempty"`
				FourThreeFallKillsBedwars                              int      `json:"four_three_fall_kills_bedwars,omitempty"`
				EightTwoProjectileFinalKillsBedwars                    int      `json:"eight_two_projectile_final_kills_bedwars,omitempty"`
				FourThreeEntityExplosionDeathsBedwars                  int      `json:"four_three_entity_explosion_deaths_bedwars,omitempty"`
				Favourites1                                            string   `json:"favourites_1,omitempty"`
				BedwarsHalloweenBoxes                                  int      `json:"bedwars_halloween_boxes,omitempty"`
				FreeEventKeyBedwarsHalloweenBoxes2017                  bool     `json:"free_event_key_bedwars_halloween_boxes_2017,omitempty"`
				FreeEventKeyBedwarsChristmasBoxes2017                  bool     `json:"free_event_key_bedwars_christmas_boxes_2017,omitempty"`
				BedwarsChristmasBoxes                                  int      `json:"bedwars_christmas_boxes,omitempty"`
				FourFourWrappedPresentResourcesCollectedBedwars        int      `json:"four_four_wrapped_present_resources_collected_bedwars,omitempty"`
				WrappedPresentResourcesCollectedBedwars                int      `json:"wrapped_present_resources_collected_bedwars,omitempty"`
				FourThreeWrappedPresentResourcesCollectedBedwars       int      `json:"four_three_wrapped_present_resources_collected_bedwars,omitempty"`
				FourFourEntityExplosionFinalDeathsBedwars              int      `json:"four_four_entity_explosion_final_deaths_bedwars,omitempty"`
				EntityExplosionFinalDeathsBedwars                      int      `json:"entity_explosion_final_deaths_bedwars,omitempty"`
				EightOneEntityAttackFinalDeathsBedwars                 int      `json:"eight_one_entity_attack_final_deaths_bedwars,omitempty"`
				EightOneWrappedPresentResourcesCollectedBedwars        int      `json:"eight_one_wrapped_present_resources_collected_bedwars,omitempty"`
				FreeEventKeyBedwarsChristmasBoxes2018                  bool     `json:"free_event_key_bedwars_christmas_boxes_2018,omitempty"`
				SeenBetaMenu                                           int      `json:"seen_beta_menu,omitempty"`
				Favourites2                                            string   `json:"favourites_2,omitempty"`
				FreeEventKeyBedwarsLunarBoxes2018                      bool     `json:"free_event_key_bedwars_lunar_boxes_2018,omitempty"`
				BedwarsLunarBoxes                                      int      `json:"bedwars_lunar_boxes,omitempty"`
				BedwarsOpenedCommons                                   int      `json:"Bedwars_openedCommons,omitempty"`
				BedwarsOpenedChests                                    int      `json:"Bedwars_openedChests,omitempty"`
				BedwarsOpenedRares                                     int      `json:"Bedwars_openedRares,omitempty"`
				ActiveGlyph                                            string   `json:"activeGlyph,omitempty"`
				ActiveSprays                                           string   `json:"activeSprays,omitempty"`
				BedwarsEasterBoxes                                     int      `json:"bedwars_easter_boxes,omitempty"`
				SuffocationDeathsBedwars                               int      `json:"suffocation_deaths_bedwars,omitempty"`
				FourFourSuffocationDeathsBedwars                       int      `json:"four_four_suffocation_deaths_bedwars,omitempty"`
				FourFourWinstreak                                      int      `json:"four_four_winstreak,omitempty"`
				SelectedUltimate                                       string   `json:"selected_ultimate,omitempty"`
				EightTwoUltimateWinstreak                              int      `json:"eight_two_ultimate_winstreak,omitempty"`
				EightTwoUltimateItemsPurchasedBedwars                  int      `json:"eight_two_ultimate__items_purchased_bedwars,omitempty"`
				EightTwoUltimateResourcesCollectedBedwars              int      `json:"eight_two_ultimate_resources_collected_bedwars,omitempty"`
				EightTwoUltimateBedsLostBedwars                        int      `json:"eight_two_ultimate_beds_lost_bedwars,omitempty"`
				EightTwoUltimateGamesPlayedBedwars                     int      `json:"eight_two_ultimate_games_played_bedwars,omitempty"`
				EightTwoUltimateIronResourcesCollectedBedwars          int      `json:"eight_two_ultimate_iron_resources_collected_bedwars,omitempty"`
				EightTwoUltimateFinalDeathsBedwars                     int      `json:"eight_two_ultimate_final_deaths_bedwars,omitempty"`
				EightTwoUltimateDiamondResourcesCollectedBedwars       int      `json:"eight_two_ultimate_diamond_resources_collected_bedwars,omitempty"`
				EightTwoUltimatePermanentItemsPurchasedBedwars         int      `json:"eight_two_ultimate_permanent _items_purchased_bedwars,omitempty"`
				EightTwoUltimateLossesBedwars                          int      `json:"eight_two_ultimate_losses_bedwars,omitempty"`
				EightTwoUltimateGoldResourcesCollectedBedwars          int      `json:"eight_two_ultimate_gold_resources_collected_bedwars,omitempty"`
				EightTwoUltimateEntityAttackKillsBedwars               int      `json:"eight_two_ultimate_entity_attack_kills_bedwars,omitempty"`
				EightTwoUltimateKillsBedwars                           int      `json:"eight_two_ultimate_kills_bedwars,omitempty"`
				EightTwoUltimateItemsPurchasedBedwars0                 int      `json:"eight_two_ultimate_items_purchased_bedwars,omitempty"`
				EightTwoUltimateEntityAttackFinalDeathsBedwars         int      `json:"eight_two_ultimate_entity_attack_final_deaths_bedwars,omitempty"`
				UnderstandsResourceBank                                bool     `json:"understands_resource_bank,omitempty"`
				CastleBedsLostBedwars                                  int      `json:"castle_beds_lost_bedwars,omitempty"`
				CastlePermanentItemsPurchasedBedwars                   int      `json:"castle_permanent _items_purchased_bedwars,omitempty"`
				CastleDeathsBedwars                                    int      `json:"castle_deaths_bedwars,omitempty"`
				CastleItemsPurchasedBedwars                            int      `json:"castle__items_purchased_bedwars,omitempty"`
				CastleFinalDeathsBedwars                               int      `json:"castle_final_deaths_bedwars,omitempty"`
				CastleVoidFinalDeathsBedwars                           int      `json:"castle_void_final_deaths_bedwars,omitempty"`
				CastleEntityAttackKillsBedwars                         int      `json:"castle_entity_attack_kills_bedwars,omitempty"`
				CastleItemsPurchasedBedwars0                           int      `json:"castle_items_purchased_bedwars,omitempty"`
				CastleGoldResourcesCollectedBedwars                    int      `json:"castle_gold_resources_collected_bedwars,omitempty"`
				CastleResourcesCollectedBedwars                        int      `json:"castle_resources_collected_bedwars,omitempty"`
				CastleEntityAttackDeathsBedwars                        int      `json:"castle_entity_attack_deaths_bedwars,omitempty"`
				CastleEmeraldResourcesCollectedBedwars                 int      `json:"castle_emerald_resources_collected_bedwars,omitempty"`
				CastleLossesBedwars                                    int      `json:"castle_losses_bedwars,omitempty"`
				CastleVoidDeathsBedwars                                int      `json:"castle_void_deaths_bedwars,omitempty"`
				CastleIronResourcesCollectedBedwars                    int      `json:"castle_iron_resources_collected_bedwars,omitempty"`
				CastleKillsBedwars                                     int      `json:"castle_kills_bedwars,omitempty"`
				CastleWinstreak                                        int      `json:"castle_winstreak,omitempty"`
				UnderstandsStreaks                                     bool     `json:"understands_streaks,omitempty"`
				CastleBedsBrokenBedwars                                int      `json:"castle_beds_broken_bedwars,omitempty"`
				CastleEntityAttackFinalDeathsBedwars                   int      `json:"castle_entity_attack_final_deaths_bedwars,omitempty"`
				CastleGamesPlayedBedwars                               int      `json:"castle_games_played_bedwars,omitempty"`
				CastleFallKillsBedwars                                 int      `json:"castle_fall_kills_bedwars,omitempty"`
				CastleEntityExplosionFinalDeathsBedwars                int      `json:"castle_entity_explosion_final_deaths_bedwars,omitempty"`
				CastleVoidFinalKillsBedwars                            int      `json:"castle_void_final_kills_bedwars,omitempty"`
				CastleFinalKillsBedwars                                int      `json:"castle_final_kills_bedwars,omitempty"`
				CastleVoidKillsBedwars                                 int      `json:"castle_void_kills_bedwars,omitempty"`
				CastleWinsBedwars                                      int      `json:"castle_wins_bedwars,omitempty"`
				CastleFallDeathsBedwars                                int      `json:"castle_fall_deaths_bedwars,omitempty"`
				CastleEntityAttackFinalKillsBedwars                    int      `json:"castle_entity_attack_final_kills_bedwars,omitempty"`
				CastleFallFinalKillsBedwars                            int      `json:"castle_fall_final_kills_bedwars,omitempty"`
				CastleDiamondResourcesCollectedBedwars                 int      `json:"castle_diamond_resources_collected_bedwars,omitempty"`
				CastleFireTickKillsBedwars                             int      `json:"castle_fire_tick_kills_bedwars,omitempty"`
				CastleEntityExplosionKillsBedwars                      int      `json:"castle_entity_explosion_kills_bedwars,omitempty"`
				CastleProjectileFinalDeathsBedwars                     int      `json:"castle_projectile_final_deaths_bedwars,omitempty"`
				CastleFallFinalDeathsBedwars                           int      `json:"castle_fall_final_deaths_bedwars,omitempty"`
				CastleEntityExplosionDeathsBedwars                     int      `json:"castle_entity_explosion_deaths_bedwars,omitempty"`
				CastleProjectileFinalKillsBedwars                      int      `json:"castle_projectile_final_kills_bedwars,omitempty"`
				CastleEntityExplosionFinalKillsBedwars                 int      `json:"castle_entity_explosion_final_kills_bedwars,omitempty"`
				FourThreeWinstreak                                     int      `json:"four_three_winstreak,omitempty"`
				EightTwoRushWinstreak                                  int      `json:"eight_two_rush_winstreak,omitempty"`
				EightTwoRushItemsPurchasedBedwars                      int      `json:"eight_two_rush_items_purchased_bedwars,omitempty"`
				EightTwoRushLossesBedwars                              int      `json:"eight_two_rush_losses_bedwars,omitempty"`
				EightTwoRushItemsPurchasedBedwars0                     int      `json:"eight_two_rush__items_purchased_bedwars,omitempty"`
				EightTwoRushResourcesCollectedBedwars                  int      `json:"eight_two_rush_resources_collected_bedwars,omitempty"`
				EightTwoRushEntityAttackDeathsBedwars                  int      `json:"eight_two_rush_entity_attack_deaths_bedwars,omitempty"`
				EightTwoRushEmeraldResourcesCollectedBedwars           int      `json:"eight_two_rush_emerald_resources_collected_bedwars,omitempty"`
				EightTwoRushKillsBedwars                               int      `json:"eight_two_rush_kills_bedwars,omitempty"`
				EightTwoRushDeathsBedwars                              int      `json:"eight_two_rush_deaths_bedwars,omitempty"`
				EightTwoRushVoidKillsBedwars                           int      `json:"eight_two_rush_void_kills_bedwars,omitempty"`
				EightTwoRushEntityAttackFinalDeathsBedwars             int      `json:"eight_two_rush_entity_attack_final_deaths_bedwars,omitempty"`
				EightTwoRushFinalKillsBedwars                          int      `json:"eight_two_rush_final_kills_bedwars,omitempty"`
				EightTwoRushGoldResourcesCollectedBedwars              int      `json:"eight_two_rush_gold_resources_collected_bedwars,omitempty"`
				EightTwoRushGamesPlayedBedwars                         int      `json:"eight_two_rush_games_played_bedwars,omitempty"`
				EightTwoRushEntityAttackKillsBedwars                   int      `json:"eight_two_rush_entity_attack_kills_bedwars,omitempty"`
				EightTwoRushVoidFinalKillsBedwars                      int      `json:"eight_two_rush_void_final_kills_bedwars,omitempty"`
				EightTwoRushBedsLostBedwars                            int      `json:"eight_two_rush_beds_lost_bedwars,omitempty"`
				EightTwoRushFinalDeathsBedwars                         int      `json:"eight_two_rush_final_deaths_bedwars,omitempty"`
				EightTwoRushIronResourcesCollectedBedwars              int      `json:"eight_two_rush_iron_resources_collected_bedwars,omitempty"`
				FourFourUltimateWinstreak                              int      `json:"four_four_ultimate_winstreak,omitempty"`
				FourFourUltimateFinalKillsBedwars                      int      `json:"four_four_ultimate_final_kills_bedwars,omitempty"`
				FourFourUltimatePermanentItemsPurchasedBedwars         int      `json:"four_four_ultimate_permanent _items_purchased_bedwars,omitempty"`
				FourFourUltimateItemsPurchasedBedwars                  int      `json:"four_four_ultimate__items_purchased_bedwars,omitempty"`
				FourFourUltimateResourcesCollectedBedwars              int      `json:"four_four_ultimate_resources_collected_bedwars,omitempty"`
				FourFourUltimateBedsLostBedwars                        int      `json:"four_four_ultimate_beds_lost_bedwars,omitempty"`
				FourFourUltimateWinsBedwars                            int      `json:"four_four_ultimate_wins_bedwars,omitempty"`
				FourFourUltimateItemsPurchasedBedwars0                 int      `json:"four_four_ultimate_items_purchased_bedwars,omitempty"`
				FourFourUltimateDiamondResourcesCollectedBedwars       int      `json:"four_four_ultimate_diamond_resources_collected_bedwars,omitempty"`
				FourFourUltimateIronResourcesCollectedBedwars          int      `json:"four_four_ultimate_iron_resources_collected_bedwars,omitempty"`
				FourFourUltimateGoldResourcesCollectedBedwars          int      `json:"four_four_ultimate_gold_resources_collected_bedwars,omitempty"`
				FourFourUltimateBedsBrokenBedwars                      int      `json:"four_four_ultimate_beds_broken_bedwars,omitempty"`
				FourFourUltimateGamesPlayedBedwars                     int      `json:"four_four_ultimate_games_played_bedwars,omitempty"`
				FourFourUltimateVoidFinalKillsBedwars                  int      `json:"four_four_ultimate_void_final_kills_bedwars,omitempty"`
				FourFourUltimateDeathsBedwars                          int      `json:"four_four_ultimate_deaths_bedwars,omitempty"`
				FourFourUltimateVoidKillsBedwars                       int      `json:"four_four_ultimate_void_kills_bedwars,omitempty"`
				FourFourUltimateVoidFinalDeathsBedwars                 int      `json:"four_four_ultimate_void_final_deaths_bedwars,omitempty"`
				FourFourUltimateEntityAttackFinalKillsBedwars          int      `json:"four_four_ultimate_entity_attack_final_kills_bedwars,omitempty"`
				FourFourUltimateLossesBedwars                          int      `json:"four_four_ultimate_losses_bedwars,omitempty"`
				FourFourUltimateKillsBedwars                           int      `json:"four_four_ultimate_kills_bedwars,omitempty"`
				FourFourUltimateVoidDeathsBedwars                      int      `json:"four_four_ultimate_void_deaths_bedwars,omitempty"`
				FourFourUltimateFinalDeathsBedwars                     int      `json:"four_four_ultimate_final_deaths_bedwars,omitempty"`
				FourFourUltimateEntityAttackKillsBedwars               int      `json:"four_four_ultimate_entity_attack_kills_bedwars,omitempty"`
				FourFourUltimateEmeraldResourcesCollectedBedwars       int      `json:"four_four_ultimate_emerald_resources_collected_bedwars,omitempty"`
				FourFourUltimateEntityAttackFinalDeathsBedwars         int      `json:"four_four_ultimate_entity_attack_final_deaths_bedwars,omitempty"`
				FourFourUltimateEntityAttackDeathsBedwars              int      `json:"four_four_ultimate_entity_attack_deaths_bedwars,omitempty"`
				FourFourUltimateProjectileFinalKillsBedwars            int      `json:"four_four_ultimate_projectile_final_kills_bedwars,omitempty"`
				FourFourUltimateProjectileKillsBedwars                 int      `json:"four_four_ultimate_projectile_kills_bedwars,omitempty"`
				FourFourUltimateEntityExplosionDeathsBedwars           int      `json:"four_four_ultimate_entity_explosion_deaths_bedwars,omitempty"`
				CastleProjectileDeathsBedwars                          int      `json:"castle_projectile_deaths_bedwars,omitempty"`
				EightTwoWinstreak                                      int      `json:"eight_two_winstreak,omitempty"`
				CastleProjectileKillsBedwars                           int      `json:"castle_projectile_kills_bedwars,omitempty"`
				CastleFireTickFinalKillsBedwars                        int      `json:"castle_fire_tick_final_kills_bedwars,omitempty"`
				FourFourRushWinstreak                                  int      `json:"four_four_rush_winstreak,omitempty"`
				FourFourRushKillsBedwars                               int      `json:"four_four_rush_kills_bedwars,omitempty"`
				FourFourRushBedsBrokenBedwars                          int      `json:"four_four_rush_beds_broken_bedwars,omitempty"`
				FourFourRushLossesBedwars                              int      `json:"four_four_rush_losses_bedwars,omitempty"`
				FourFourRushItemsPurchasedBedwars                      int      `json:"four_four_rush_items_purchased_bedwars,omitempty"`
				FourFourRushGamesPlayedBedwars                         int      `json:"four_four_rush_games_played_bedwars,omitempty"`
				FourFourRushDeathsBedwars                              int      `json:"four_four_rush_deaths_bedwars,omitempty"`
				FourFourRushEntityAttackFinalKillsBedwars              int      `json:"four_four_rush_entity_attack_final_kills_bedwars,omitempty"`
				FourFourRushEmeraldResourcesCollectedBedwars           int      `json:"four_four_rush_emerald_resources_collected_bedwars,omitempty"`
				FourFourRushBedsLostBedwars                            int      `json:"four_four_rush_beds_lost_bedwars,omitempty"`
				FourFourRushEntityAttackDeathsBedwars                  int      `json:"four_four_rush_entity_attack_deaths_bedwars,omitempty"`
				FourFourRushFinalKillsBedwars                          int      `json:"four_four_rush_final_kills_bedwars,omitempty"`
				FourFourRushGoldResourcesCollectedBedwars              int      `json:"four_four_rush_gold_resources_collected_bedwars,omitempty"`
				FourFourRushResourcesCollectedBedwars                  int      `json:"four_four_rush_resources_collected_bedwars,omitempty"`
				FourFourRushVoidKillsBedwars                           int      `json:"four_four_rush_void_kills_bedwars,omitempty"`
				FourFourRushDiamondResourcesCollectedBedwars           int      `json:"four_four_rush_diamond_resources_collected_bedwars,omitempty"`
				FourFourRushFinalDeathsBedwars                         int      `json:"four_four_rush_final_deaths_bedwars,omitempty"`
				FourFourRushEntityAttackKillsBedwars                   int      `json:"four_four_rush_entity_attack_kills_bedwars,omitempty"`
				FourFourRushItemsPurchasedBedwars0                     int      `json:"four_four_rush__items_purchased_bedwars,omitempty"`
				FourFourRushVoidFinalDeathsBedwars                     int      `json:"four_four_rush_void_final_deaths_bedwars,omitempty"`
				FourFourRushPermanentItemsPurchasedBedwars             int      `json:"four_four_rush_permanent _items_purchased_bedwars,omitempty"`
				FourFourRushIronResourcesCollectedBedwars              int      `json:"four_four_rush_iron_resources_collected_bedwars,omitempty"`
				FourFourRushVoidDeathsBedwars                          int      `json:"four_four_rush_void_deaths_bedwars,omitempty"`
				FourFourRushEntityAttackFinalDeathsBedwars             int      `json:"four_four_rush_entity_attack_final_deaths_bedwars,omitempty"`
				FourFourRushFallDeathsBedwars                          int      `json:"four_four_rush_fall_deaths_bedwars,omitempty"`
				FourFourRushVoidFinalKillsBedwars                      int      `json:"four_four_rush_void_final_kills_bedwars,omitempty"`
				FourFourRushWinsBedwars                                int      `json:"four_four_rush_wins_bedwars,omitempty"`
				FourFourRushFallKillsBedwars                           int      `json:"four_four_rush_fall_kills_bedwars,omitempty"`
				EightTwoRushBedsBrokenBedwars                          int      `json:"eight_two_rush_beds_broken_bedwars,omitempty"`
				EightTwoRushEntityAttackFinalKillsBedwars              int      `json:"eight_two_rush_entity_attack_final_kills_bedwars,omitempty"`
				EightTwoRushDiamondResourcesCollectedBedwars           int      `json:"eight_two_rush_diamond_resources_collected_bedwars,omitempty"`
				EightTwoRushFallFinalDeathsBedwars                     int      `json:"eight_two_rush_fall_final_deaths_bedwars,omitempty"`
				EightTwoRushPermanentItemsPurchasedBedwars             int      `json:"eight_two_rush_permanent _items_purchased_bedwars,omitempty"`
				FourFourRushProjectileKillsBedwars                     int      `json:"four_four_rush_projectile_kills_bedwars,omitempty"`
				FourFourRushFallFinalKillsBedwars                      int      `json:"four_four_rush_fall_final_kills_bedwars,omitempty"`
				FourFourRushFallFinalDeathsBedwars                     int      `json:"four_four_rush_fall_final_deaths_bedwars,omitempty"`
				EightTwoRushVoidDeathsBedwars                          int      `json:"eight_two_rush_void_deaths_bedwars,omitempty"`
				EightTwoRushVoidFinalDeathsBedwars                     int      `json:"eight_two_rush_void_final_deaths_bedwars,omitempty"`
				FourFourRushEntityExplosionFinalKillsBedwars           int      `json:"four_four_rush_entity_explosion_final_kills_bedwars,omitempty"`
				EightTwoRushWinsBedwars                                int      `json:"eight_two_rush_wins_bedwars,omitempty"`
				EightTwoRushFallFinalKillsBedwars                      int      `json:"eight_two_rush_fall_final_kills_bedwars,omitempty"`
				EightTwoRushProjectileKillsBedwars                     int      `json:"eight_two_rush_projectile_kills_bedwars,omitempty"`
				FourFourUltimateFireTickDeathsBedwars                  int      `json:"four_four_ultimate_fire_tick_deaths_bedwars,omitempty"`
				FourFourUltimateFireTickKillsBedwars                   int      `json:"four_four_ultimate_fire_tick_kills_bedwars,omitempty"`
				FourFourUltimateProjectileFinalDeathsBedwars           int      `json:"four_four_ultimate_projectile_final_deaths_bedwars,omitempty"`
				FourFourUltimateEntityExplosionFinalKillsBedwars       int      `json:"four_four_ultimate_entity_explosion_final_kills_bedwars,omitempty"`
				FreeEventKeyBedwarsHalloweenBoxes2018                  bool     `json:"free_event_key_bedwars_halloween_boxes_2018,omitempty"`
				SpookyOpenAch                                          int      `json:"spooky_open_ach,omitempty"`
				ActiveBedDestroy                                       string   `json:"activeBedDestroy,omitempty"`
				FavoriteSlots                                          string   `json:"favorite_slots,omitempty"`
				FourFourRushEntityExplosionFinalDeathsBedwars          int      `json:"four_four_rush_entity_explosion_final_deaths_bedwars,omitempty"`
				LastTourneyAd                                          int64    `json:"lastTourneyAd,omitempty"`
				TourneyBedwars4S0Winstreak2                            int      `json:"tourney_bedwars4s_0_winstreak2,omitempty"`
				TourneyBedwars4S0DiamondResourcesCollectedBedwars      int      `json:"tourney_bedwars4s_0_diamond_resources_collected_bedwars,omitempty"`
				TourneyBedwars4S0FinalDeathsBedwars                    int      `json:"tourney_bedwars4s_0_final_deaths_bedwars,omitempty"`
				TourneyBedwars4S0VoidKillsBedwars                      int      `json:"tourney_bedwars4s_0_void_kills_bedwars,omitempty"`
				TourneyBedwars4S0ItemsPurchasedBedwars                 int      `json:"tourney_bedwars4s_0__items_purchased_bedwars,omitempty"`
				TourneyBedwars4S0ItemsPurchasedBedwars0                int      `json:"tourney_bedwars4s_0_items_purchased_bedwars,omitempty"`
				TourneyBedwars4S0BedsLostBedwars                       int      `json:"tourney_bedwars4s_0_beds_lost_bedwars,omitempty"`
				TourneyBedwars4S0LossesBedwars                         int      `json:"tourney_bedwars4s_0_losses_bedwars,omitempty"`
				TourneyBedwars4S0VoidDeathsBedwars                     int      `json:"tourney_bedwars4s_0_void_deaths_bedwars,omitempty"`
				TourneyBedwars4S0EntityAttackKillsBedwars              int      `json:"tourney_bedwars4s_0_entity_attack_kills_bedwars,omitempty"`
				TourneyBedwars4S0ResourcesCollectedBedwars             int      `json:"tourney_bedwars4s_0_resources_collected_bedwars,omitempty"`
				TourneyBedwars4S0IronResourcesCollectedBedwars         int      `json:"tourney_bedwars4s_0_iron_resources_collected_bedwars,omitempty"`
				TourneyBedwars4S0DeathsBedwars                         int      `json:"tourney_bedwars4s_0_deaths_bedwars,omitempty"`
				TourneyBedwars4S0GamesPlayedBedwars                    int      `json:"tourney_bedwars4s_0_games_played_bedwars,omitempty"`
				TourneyBedwars4S0PermanentItemsPurchasedBedwars        int      `json:"tourney_bedwars4s_0_permanent _items_purchased_bedwars,omitempty"`
				TourneyBedwars4S0EntityAttackFinalDeathsBedwars        int      `json:"tourney_bedwars4s_0_entity_attack_final_deaths_bedwars,omitempty"`
				TourneyBedwars4S0GoldResourcesCollectedBedwars         int      `json:"tourney_bedwars4s_0_gold_resources_collected_bedwars,omitempty"`
				TourneyBedwars4S0KillsBedwars                          int      `json:"tourney_bedwars4s_0_kills_bedwars,omitempty"`
				TourneyBedwars4S0EntityAttackDeathsBedwars             int      `json:"tourney_bedwars4s_0_entity_attack_deaths_bedwars,omitempty"`
				TourneyBedwars4S0BedsBrokenBedwars                     int      `json:"tourney_bedwars4s_0_beds_broken_bedwars,omitempty"`
				TourneyBedwars4S0EmeraldResourcesCollectedBedwars      int      `json:"tourney_bedwars4s_0_emerald_resources_collected_bedwars,omitempty"`
				TourneyBedwars4S0FinalKillsBedwars                     int      `json:"tourney_bedwars4s_0_final_kills_bedwars,omitempty"`
				TourneyBedwars4S0WinsBedwars                           int      `json:"tourney_bedwars4s_0_wins_bedwars,omitempty"`
				TourneyBedwars4S0EntityAttackFinalKillsBedwars         int      `json:"tourney_bedwars4s_0_entity_attack_final_kills_bedwars,omitempty"`
				TourneyBedwars4S0VoidFinalKillsBedwars                 int      `json:"tourney_bedwars4s_0_void_final_kills_bedwars,omitempty"`
				FourFourUltimateSuffocationKillsBedwars                int      `json:"four_four_ultimate_suffocation_kills_bedwars,omitempty"`
				LastHytaleAd                                           int64    `json:"lastHytaleAd,omitempty"`
				FreeEventKeyBedwarsChristmasBoxes2019                  bool     `json:"free_event_key_bedwars_christmas_boxes_2019,omitempty"`
				EightOneWinstreak                                      int      `json:"eight_one_winstreak,omitempty"`
				FourFourLuckyWinstreak                                 int      `json:"four_four_lucky_winstreak,omitempty"`
				FourFourLuckyPermanentItemsPurchasedBedwars            int      `json:"four_four_lucky_permanent _items_purchased_bedwars,omitempty"`
				FourFourLuckyEmeraldResourcesCollectedBedwars          int      `json:"four_four_lucky_emerald_resources_collected_bedwars,omitempty"`
				FourFourLuckyItemsPurchasedBedwars                     int      `json:"four_four_lucky_items_purchased_bedwars,omitempty"`
				FourFourLuckyFinalDeathsBedwars                        int      `json:"four_four_lucky_final_deaths_bedwars,omitempty"`
				FourFourLuckyLossesBedwars                             int      `json:"four_four_lucky_losses_bedwars,omitempty"`
				FourFourLuckyGamesPlayedBedwars                        int      `json:"four_four_lucky_games_played_bedwars,omitempty"`
				FourFourLuckyEntityAttackKillsBedwars                  int      `json:"four_four_lucky_entity_attack_kills_bedwars,omitempty"`
				FourFourLuckyItemsPurchasedBedwars0                    int      `json:"four_four_lucky__items_purchased_bedwars,omitempty"`
				FourFourLuckyKillsBedwars                              int      `json:"four_four_lucky_kills_bedwars,omitempty"`
				FourFourLuckyEntityAttackFinalDeathsBedwars            int      `json:"four_four_lucky_entity_attack_final_deaths_bedwars,omitempty"`
				FourFourLuckyResourcesCollectedBedwars                 int      `json:"four_four_lucky_resources_collected_bedwars,omitempty"`
				FourFourLuckyIronResourcesCollectedBedwars             int      `json:"four_four_lucky_iron_resources_collected_bedwars,omitempty"`
				FourFourLuckyBedsLostBedwars                           int      `json:"four_four_lucky_beds_lost_bedwars,omitempty"`
				FourFourLuckyGoldResourcesCollectedBedwars             int      `json:"four_four_lucky_gold_resources_collected_bedwars,omitempty"`
				FourFourLuckyDiamondResourcesCollectedBedwars          int      `json:"four_four_lucky_diamond_resources_collected_bedwars,omitempty"`
				FourFourLuckyVoidDeathsBedwars                         int      `json:"four_four_lucky_void_deaths_bedwars,omitempty"`
				FourFourLuckyBedsBrokenBedwars                         int      `json:"four_four_lucky_beds_broken_bedwars,omitempty"`
				FourFourLuckyFinalKillsBedwars                         int      `json:"four_four_lucky_final_kills_bedwars,omitempty"`
				FourFourLuckyVoidFinalKillsBedwars                     int      `json:"four_four_lucky_void_final_kills_bedwars,omitempty"`
				FourFourLuckyDeathsBedwars                             int      `json:"four_four_lucky_deaths_bedwars,omitempty"`
				FourFourLuckyVoidKillsBedwars                          int      `json:"four_four_lucky_void_kills_bedwars,omitempty"`
				FourFourLuckyEntityAttackDeathsBedwars                 int      `json:"four_four_lucky_entity_attack_deaths_bedwars,omitempty"`
				FourFourLuckyEntityAttackFinalKillsBedwars             int      `json:"four_four_lucky_entity_attack_final_kills_bedwars,omitempty"`
				FourFourLuckyFireTickKillsBedwars                      int      `json:"four_four_lucky_fire_tick_kills_bedwars,omitempty"`
				FourFourLuckyVoidFinalDeathsBedwars                    int      `json:"four_four_lucky_void_final_deaths_bedwars,omitempty"`
				FourFourLuckyWinsBedwars                               int      `json:"four_four_lucky_wins_bedwars,omitempty"`
				FourFourLuckyFireTickFinalDeathsBedwars                int      `json:"four_four_lucky_fire_tick_final_deaths_bedwars,omitempty"`
				FourFourLuckyFallFinalKillsBedwars                     int      `json:"four_four_lucky_fall_final_kills_bedwars,omitempty"`
				FourFourLuckyFireTickFinalKillsBedwars                 int      `json:"four_four_lucky_fire_tick_final_kills_bedwars,omitempty"`
				FourFourLuckyFallKillsBedwars                          int      `json:"four_four_lucky_fall_kills_bedwars,omitempty"`
				FreeEventKeyBedwarsLunarBoxes2019                      bool     `json:"free_event_key_bedwars_lunar_boxes_2019,omitempty"`
				FourFourLuckyFireTickDeathsBedwars                     int      `json:"four_four_lucky_fire_tick_deaths_bedwars,omitempty"`
				FourFourLuckyFallDeathsBedwars                         int      `json:"four_four_lucky_fall_deaths_bedwars,omitempty"`
				FourFourLuckyProjectileDeathsBedwars                   int      `json:"four_four_lucky_projectile_deaths_bedwars,omitempty"`
				FourFourLuckyProjectileKillsBedwars                    int      `json:"four_four_lucky_projectile_kills_bedwars,omitempty"`
				EightTwoUltimateVoidFinalDeathsBedwars                 int      `json:"eight_two_ultimate_void_final_deaths_bedwars,omitempty"`
				EightTwoUltimateDeathsBedwars                          int      `json:"eight_two_ultimate_deaths_bedwars,omitempty"`
				EightTwoUltimateEntityAttackDeathsBedwars              int      `json:"eight_two_ultimate_entity_attack_deaths_bedwars,omitempty"`
				EightTwoUltimateFinalKillsBedwars                      int      `json:"eight_two_ultimate_final_kills_bedwars,omitempty"`
				EightTwoUltimateBedsBrokenBedwars                      int      `json:"eight_two_ultimate_beds_broken_bedwars,omitempty"`
				EightTwoUltimateVoidFinalKillsBedwars                  int      `json:"eight_two_ultimate_void_final_kills_bedwars,omitempty"`
				EightTwoUltimateEntityAttackFinalKillsBedwars          int      `json:"eight_two_ultimate_entity_attack_final_kills_bedwars,omitempty"`
				EightTwoUltimateEmeraldResourcesCollectedBedwars       int      `json:"eight_two_ultimate_emerald_resources_collected_bedwars,omitempty"`
				EightTwoUltimateWinsBedwars                            int      `json:"eight_two_ultimate_wins_bedwars,omitempty"`
				EightTwoUltimateVoidDeathsBedwars                      int      `json:"eight_two_ultimate_void_deaths_bedwars,omitempty"`
				EightTwoUltimateVoidKillsBedwars                       int      `json:"eight_two_ultimate_void_kills_bedwars,omitempty"`
				EightTwoUltimateProjectileFinalKillsBedwars            int      `json:"eight_two_ultimate_projectile_final_kills_bedwars,omitempty"`
				FourFourRushEntityExplosionDeathsBedwars               int      `json:"four_four_rush_entity_explosion_deaths_bedwars,omitempty"`
				FourFourArmedWinstreak                                 int      `json:"four_four_armed_winstreak,omitempty"`
				FourFourArmedItemsPurchasedBedwars                     int      `json:"four_four_armed__items_purchased_bedwars,omitempty"`
				FourFourArmedBedsBrokenBedwars                         int      `json:"four_four_armed_beds_broken_bedwars,omitempty"`
				FourFourArmedDeathsBedwars                             int      `json:"four_four_armed_deaths_bedwars,omitempty"`
				FourFourArmedDiamondResourcesCollectedBedwars          int      `json:"four_four_armed_diamond_resources_collected_bedwars,omitempty"`
				FourFourArmedEmeraldResourcesCollectedBedwars          int      `json:"four_four_armed_emerald_resources_collected_bedwars,omitempty"`
				FourFourArmedEntityAttackDeathsBedwars                 int      `json:"four_four_armed_entity_attack_deaths_bedwars,omitempty"`
				FourFourArmedEntityAttackKillsBedwars                  int      `json:"four_four_armed_entity_attack_kills_bedwars,omitempty"`
				FourFourArmedFinalKillsBedwars                         int      `json:"four_four_armed_final_kills_bedwars,omitempty"`
				FourFourArmedGamesPlayedBedwars                        int      `json:"four_four_armed_games_played_bedwars,omitempty"`
				FourFourArmedGoldResourcesCollectedBedwars             int      `json:"four_four_armed_gold_resources_collected_bedwars,omitempty"`
				FourFourArmedIronResourcesCollectedBedwars             int      `json:"four_four_armed_iron_resources_collected_bedwars,omitempty"`
				FourFourArmedItemsPurchasedBedwars0                    int      `json:"four_four_armed_items_purchased_bedwars,omitempty"`
				FourFourArmedKillsBedwars                              int      `json:"four_four_armed_kills_bedwars,omitempty"`
				FourFourArmedPermanentItemsPurchasedBedwars            int      `json:"four_four_armed_permanent _items_purchased_bedwars,omitempty"`
				FourFourArmedProjectileDeathsBedwars                   int      `json:"four_four_armed_projectile_deaths_bedwars,omitempty"`
				FourFourArmedProjectileKillsBedwars                    int      `json:"four_four_armed_projectile_kills_bedwars,omitempty"`
				FourFourArmedResourcesCollectedBedwars                 int      `json:"four_four_armed_resources_collected_bedwars,omitempty"`
				FourFourArmedVoidDeathsBedwars                         int      `json:"four_four_armed_void_deaths_bedwars,omitempty"`
				FourFourArmedVoidFinalKillsBedwars                     int      `json:"four_four_armed_void_final_kills_bedwars,omitempty"`
				FourFourArmedVoidKillsBedwars                          int      `json:"four_four_armed_void_kills_bedwars,omitempty"`
				FourFourArmedWinsBedwars                               int      `json:"four_four_armed_wins_bedwars,omitempty"`
				FourFourArmedBedsLostBedwars                           int      `json:"four_four_armed_beds_lost_bedwars,omitempty"`
				FourFourArmedFinalDeathsBedwars                        int      `json:"four_four_armed_final_deaths_bedwars,omitempty"`
				FourFourArmedLossesBedwars                             int      `json:"four_four_armed_losses_bedwars,omitempty"`
				FourFourArmedProjectileFinalDeathsBedwars              int      `json:"four_four_armed_projectile_final_deaths_bedwars,omitempty"`
				FourFourArmedProjectileFinalKillsBedwars               int      `json:"four_four_armed_projectile_final_kills_bedwars,omitempty"`
				FourFourArmedFallKillsBedwars                          int      `json:"four_four_armed_fall_kills_bedwars,omitempty"`
				FourFourArmedFallDeathsBedwars                         int      `json:"four_four_armed_fall_deaths_bedwars,omitempty"`
				FourFourArmedEntityAttackFinalDeathsBedwars            int      `json:"four_four_armed_entity_attack_final_deaths_bedwars,omitempty"`
				FourFourArmedEntityAttackFinalKillsBedwars             int      `json:"four_four_armed_entity_attack_final_kills_bedwars,omitempty"`
				FourFourArmedFallFinalDeathsBedwars                    int      `json:"four_four_armed_fall_final_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0Winstreak2                       int      `json:"tourney_bedwars_two_four_0_winstreak2,omitempty"`
				TourneyBedwarsTwoFour0ItemsPurchasedBedwars            int      `json:"tourney_bedwars_two_four_0__items_purchased_bedwars,omitempty"`
				TourneyBedwarsTwoFour0BedsBrokenBedwars                int      `json:"tourney_bedwars_two_four_0_beds_broken_bedwars,omitempty"`
				TourneyBedwarsTwoFour0DeathsBedwars                    int      `json:"tourney_bedwars_two_four_0_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EntityAttackDeathsBedwars        int      `json:"tourney_bedwars_two_four_0_entity_attack_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EntityAttackFinalKillsBedwars    int      `json:"tourney_bedwars_two_four_0_entity_attack_final_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0FallDeathsBedwars                int      `json:"tourney_bedwars_two_four_0_fall_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0FinalKillsBedwars                int      `json:"tourney_bedwars_two_four_0_final_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0GamesPlayedBedwars               int      `json:"tourney_bedwars_two_four_0_games_played_bedwars,omitempty"`
				TourneyBedwarsTwoFour0GoldResourcesCollectedBedwars    int      `json:"tourney_bedwars_two_four_0_gold_resources_collected_bedwars,omitempty"`
				TourneyBedwarsTwoFour0IronResourcesCollectedBedwars    int      `json:"tourney_bedwars_two_four_0_iron_resources_collected_bedwars,omitempty"`
				TourneyBedwarsTwoFour0ItemsPurchasedBedwars0           int      `json:"tourney_bedwars_two_four_0_items_purchased_bedwars,omitempty"`
				TourneyBedwarsTwoFour0PermanentItemsPurchasedBedwars   int      `json:"tourney_bedwars_two_four_0_permanent _items_purchased_bedwars,omitempty"`
				TourneyBedwarsTwoFour0ResourcesCollectedBedwars        int      `json:"tourney_bedwars_two_four_0_resources_collected_bedwars,omitempty"`
				TourneyBedwarsTwoFour0VoidDeathsBedwars                int      `json:"tourney_bedwars_two_four_0_void_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0VoidFinalKillsBedwars            int      `json:"tourney_bedwars_two_four_0_void_final_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0WinsBedwars                      int      `json:"tourney_bedwars_two_four_0_wins_bedwars,omitempty"`
				TourneyBedwarsTwoFour0KillsBedwars                     int      `json:"tourney_bedwars_two_four_0_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0VoidKillsBedwars                 int      `json:"tourney_bedwars_two_four_0_void_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0BedsLostBedwars                  int      `json:"tourney_bedwars_two_four_0_beds_lost_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EmeraldResourcesCollectedBedwars int      `json:"tourney_bedwars_two_four_0_emerald_resources_collected_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EntityAttackKillsBedwars         int      `json:"tourney_bedwars_two_four_0_entity_attack_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EntityExplosionDeathsBedwars     int      `json:"tourney_bedwars_two_four_0_entity_explosion_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0FinalDeathsBedwars               int      `json:"tourney_bedwars_two_four_0_final_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0LossesBedwars                    int      `json:"tourney_bedwars_two_four_0_losses_bedwars,omitempty"`
				TourneyBedwarsTwoFour0VoidFinalDeathsBedwars           int      `json:"tourney_bedwars_two_four_0_void_final_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0EntityAttackFinalDeathsBedwars   int      `json:"tourney_bedwars_two_four_0_entity_attack_final_deaths_bedwars,omitempty"`
				TourneyBedwarsTwoFour0FallKillsBedwars                 int      `json:"tourney_bedwars_two_four_0_fall_kills_bedwars,omitempty"`
				TourneyBedwarsTwoFour0FallFinalKillsBedwars            int      `json:"tourney_bedwars_two_four_0_fall_final_kills_bedwars,omitempty"`
				FourFourMagicFinalDeathsBedwars                        int      `json:"four_four_magic_final_deaths_bedwars,omitempty"`
				MagicFinalDeathsBedwars                                int      `json:"magic_final_deaths_bedwars,omitempty"`
				FourFourMagicDeathsBedwars                             int      `json:"four_four_magic_deaths_bedwars,omitempty"`
				MagicDeathsBedwars                                     int      `json:"magic_deaths_bedwars,omitempty"`
				FreeEventKeyBedwarsEasterBoxes2020                     bool     `json:"free_event_key_bedwars_easter_boxes_2020,omitempty"`
				FourFourUltimatePermanentItemsPurchasedBedwars0        int      `json:"four_four_ultimate_permanent_items_purchased_bedwars,omitempty"`
				FourFourUltimateSuffocationFinalDeathsBedwars          int      `json:"four_four_ultimate_suffocation_final_deaths_bedwars,omitempty"`
				FourFourUltimateMagicFinalDeathsBedwars                int      `json:"four_four_ultimate_magic_final_deaths_bedwars,omitempty"`
				EightTwoUltimatePermanentItemsPurchasedBedwars0        int      `json:"eight_two_ultimate_permanent_items_purchased_bedwars,omitempty"`
				EightTwoVoidlessWinstreak                              int      `json:"eight_two_voidless_winstreak,omitempty"`
				EightTwoVoidlessItemsPurchasedBedwars                  int      `json:"eight_two_voidless__items_purchased_bedwars,omitempty"`
				EightTwoVoidlessBedsBrokenBedwars                      int      `json:"eight_two_voidless_beds_broken_bedwars,omitempty"`
				EightTwoVoidlessBedsLostBedwars                        int      `json:"eight_two_voidless_beds_lost_bedwars,omitempty"`
				EightTwoVoidlessDeathsBedwars                          int      `json:"eight_two_voidless_deaths_bedwars,omitempty"`
				EightTwoVoidlessDiamondResourcesCollectedBedwars       int      `json:"eight_two_voidless_diamond_resources_collected_bedwars,omitempty"`
				EightTwoVoidlessEntityAttackDeathsBedwars              int      `json:"eight_two_voidless_entity_attack_deaths_bedwars,omitempty"`
				EightTwoVoidlessEntityAttackFinalDeathsBedwars         int      `json:"eight_two_voidless_entity_attack_final_deaths_bedwars,omitempty"`
				EightTwoVoidlessEntityAttackKillsBedwars               int      `json:"eight_two_voidless_entity_attack_kills_bedwars,omitempty"`
				EightTwoVoidlessFinalDeathsBedwars                     int      `json:"eight_two_voidless_final_deaths_bedwars,omitempty"`
				EightTwoVoidlessFinalKillsBedwars                      int      `json:"eight_two_voidless_final_kills_bedwars,omitempty"`
				EightTwoVoidlessGamesPlayedBedwars                     int      `json:"eight_two_voidless_games_played_bedwars,omitempty"`
				EightTwoVoidlessGoldResourcesCollectedBedwars          int      `json:"eight_two_voidless_gold_resources_collected_bedwars,omitempty"`
				EightTwoVoidlessIronResourcesCollectedBedwars          int      `json:"eight_two_voidless_iron_resources_collected_bedwars,omitempty"`
				EightTwoVoidlessItemsPurchasedBedwars0                 int      `json:"eight_two_voidless_items_purchased_bedwars,omitempty"`
				EightTwoVoidlessKillsBedwars                           int      `json:"eight_two_voidless_kills_bedwars,omitempty"`
				EightTwoVoidlessLossesBedwars                          int      `json:"eight_two_voidless_losses_bedwars,omitempty"`
				EightTwoVoidlessMagicFinalKillsBedwars                 int      `json:"eight_two_voidless_magic_final_kills_bedwars,omitempty"`
				EightTwoVoidlessPermanentItemsPurchasedBedwars         int      `json:"eight_two_voidless_permanent_items_purchased_bedwars,omitempty"`
				EightTwoVoidlessResourcesCollectedBedwars              int      `json:"eight_two_voidless_resources_collected_bedwars,omitempty"`
				FourFourPermanentItemsPurchasedBedwars0                int      `json:"four_four_permanent_items_purchased_bedwars,omitempty"`
				PermanentItemsPurchasedBedwars0                        int      `json:"permanent_items_purchased_bedwars,omitempty"`
				FourThreePermanentItemsPurchasedBedwars0               int      `json:"four_three_permanent_items_purchased_bedwars,omitempty"`
				EightTwoPermanentItemsPurchasedBedwars0                int      `json:"eight_two_permanent_items_purchased_bedwars,omitempty"`
				FreeEventKeyBedwarsEasterBoxes2021                     bool     `json:"free_event_key_bedwars_easter_boxes_2021,omitempty"`
				Practice                                               struct {
					Selected string `json:"selected,omitempty"`
					Mlg      struct {
						FailedAttempts     int `json:"failed_attempts,omitempty"`
						SuccessfulAttempts int `json:"successful_attempts,omitempty"`
					} `json:"mlg,omitempty"`
					Records struct {
						BridgingDistance30ElevationNONEAngleSTRAIGHT int `json:"bridging_distance_30:elevation_NONE:angle_STRAIGHT:,omitempty"`
					} `json:"records,omitempty"`
					Bridging struct {
						BlocksPlaced       int `json:"blocks_placed,omitempty"`
						FailedAttempts     int `json:"failed_attempts,omitempty"`
						SuccessfulAttempts int `json:"successful_attempts,omitempty"`
					} `json:"bridging,omitempty"`
					FireballJumping struct {
						FailedAttempts     int `json:"failed_attempts,omitempty"`
						SuccessfulAttempts int `json:"successful_attempts,omitempty"`
					} `json:"fireball_jumping,omitempty"`
				} `json:"practice,omitempty"`
				CastleMagicDeathsBedwars                              int    `json:"castle_magic_deaths_bedwars,omitempty"`
				CastlePermanentItemsPurchasedBedwars0                 int    `json:"castle_permanent_items_purchased_bedwars,omitempty"`
				EightTwoVoidlessEmeraldResourcesCollectedBedwars      int    `json:"eight_two_voidless_emerald_resources_collected_bedwars,omitempty"`
				EightTwoVoidlessFallFinalKillsBedwars                 int    `json:"eight_two_voidless_fall_final_kills_bedwars,omitempty"`
				SelectedChallengeType                                 string `json:"selected_challenge_type,omitempty"`
				TourneyBedwarsEightTwo0Winstreak2                     int    `json:"tourney_bedwars_eight_two_0_winstreak2,omitempty"`
				TourneyBedwarsEightTwo0ItemsPurchasedBedwars          int    `json:"tourney_bedwars_eight_two_0__items_purchased_bedwars,omitempty"`
				TourneyBedwarsEightTwo0BedsLostBedwars                int    `json:"tourney_bedwars_eight_two_0_beds_lost_bedwars,omitempty"`
				TourneyBedwarsEightTwo0DeathsBedwars                  int    `json:"tourney_bedwars_eight_two_0_deaths_bedwars,omitempty"`
				TourneyBedwarsEightTwo0EntityAttackDeathsBedwars      int    `json:"tourney_bedwars_eight_two_0_entity_attack_deaths_bedwars,omitempty"`
				TourneyBedwarsEightTwo0EntityAttackFinalDeathsBedwars int    `json:"tourney_bedwars_eight_two_0_entity_attack_final_deaths_bedwars,omitempty"`
				TourneyBedwarsEightTwo0EntityAttackKillsBedwars       int    `json:"tourney_bedwars_eight_two_0_entity_attack_kills_bedwars,omitempty"`
				TourneyBedwarsEightTwo0FinalDeathsBedwars             int    `json:"tourney_bedwars_eight_two_0_final_deaths_bedwars,omitempty"`
				TourneyBedwarsEightTwo0GamesPlayedBedwars             int    `json:"tourney_bedwars_eight_two_0_games_played_bedwars,omitempty"`
				TourneyBedwarsEightTwo0GoldResourcesCollectedBedwars  int    `json:"tourney_bedwars_eight_two_0_gold_resources_collected_bedwars,omitempty"`
				TourneyBedwarsEightTwo0IronResourcesCollectedBedwars  int    `json:"tourney_bedwars_eight_two_0_iron_resources_collected_bedwars,omitempty"`
				TourneyBedwarsEightTwo0ItemsPurchasedBedwars0         int    `json:"tourney_bedwars_eight_two_0_items_purchased_bedwars,omitempty"`
				TourneyBedwarsEightTwo0KillsBedwars                   int    `json:"tourney_bedwars_eight_two_0_kills_bedwars,omitempty"`
				TourneyBedwarsEightTwo0LossesBedwars                  int    `json:"tourney_bedwars_eight_two_0_losses_bedwars,omitempty"`
				TourneyBedwarsEightTwo0MagicKillsBedwars              int    `json:"tourney_bedwars_eight_two_0_magic_kills_bedwars,omitempty"`
				TourneyBedwarsEightTwo0ResourcesCollectedBedwars      int    `json:"tourney_bedwars_eight_two_0_resources_collected_bedwars,omitempty"`
				FourFourUltimateMagicFinalKillsBedwars                int    `json:"four_four_ultimate_magic_final_kills_bedwars,omitempty"`
				FourFourLuckyMagicDeathsBedwars                       int    `json:"four_four_lucky_magic_deaths_bedwars,omitempty"`
				Slumber                                               struct {
					Quest struct {
						LastStarted struct {
							NpcReceptionStart int64 `json:"npc_reception_start,omitempty"`
						} `json:"lastStarted,omitempty"`
						Started struct {
							NpcReceptionStart bool `json:"npc_reception_start,omitempty"`
						} `json:"started,omitempty"`
					} `json:"quest,omitempty"`
				} `json:"slumber,omitempty"`
			} `json:"Bedwars,omitempty"`
			MurderMystery struct {
				MurdermysteryBooks                                       []string `json:"murdermystery_books,omitempty"`
				DetectiveChance                                          int      `json:"detective_chance,omitempty"`
				MurdererChance                                           int      `json:"murderer_chance,omitempty"`
				GamesArchivesMURDERCLASSIC                               int      `json:"games_archives_MURDER_CLASSIC,omitempty"`
				Wins                                                     int      `json:"wins,omitempty"`
				Games                                                    int      `json:"games,omitempty"`
				GamesMURDERCLASSIC                                       int      `json:"games_MURDER_CLASSIC,omitempty"`
				GamesArchives                                            int      `json:"games_archives,omitempty"`
				Coins                                                    int      `json:"coins,omitempty"`
				WinsArchivesMURDERCLASSIC                                int      `json:"wins_archives_MURDER_CLASSIC,omitempty"`
				WinsArchives                                             int      `json:"wins_archives,omitempty"`
				WinsMURDERCLASSIC                                        int      `json:"wins_MURDER_CLASSIC,omitempty"`
				CoinsPickedupMURDERCLASSIC                               int      `json:"coins_pickedup_MURDER_CLASSIC,omitempty"`
				WinsLibraryMURDERCLASSIC                                 int      `json:"wins_library_MURDER_CLASSIC,omitempty"`
				WinsLibrary                                              int      `json:"wins_library,omitempty"`
				GamesLibraryMURDERCLASSIC                                int      `json:"games_library_MURDER_CLASSIC,omitempty"`
				CoinsPickedupLibrary                                     int      `json:"coins_pickedup_library,omitempty"`
				CoinsPickedupLibraryMURDERCLASSIC                        int      `json:"coins_pickedup_library_MURDER_CLASSIC,omitempty"`
				GamesLibrary                                             int      `json:"games_library,omitempty"`
				CoinsPickedup                                            int      `json:"coins_pickedup,omitempty"`
				GamesGoldRushMURDERCLASSIC                               int      `json:"games_gold_rush_MURDER_CLASSIC,omitempty"`
				DeathsGoldRush                                           int      `json:"deaths_gold_rush,omitempty"`
				DeathsGoldRushMURDERCLASSIC                              int      `json:"deaths_gold_rush_MURDER_CLASSIC,omitempty"`
				Deaths                                                   int      `json:"deaths,omitempty"`
				GamesGoldRush                                            int      `json:"games_gold_rush,omitempty"`
				CoinsPickedupGoldRush                                    int      `json:"coins_pickedup_gold_rush,omitempty"`
				DeathsMURDERCLASSIC                                      int      `json:"deaths_MURDER_CLASSIC,omitempty"`
				CoinsPickedupGoldRushMURDERCLASSIC                       int      `json:"coins_pickedup_gold_rush_MURDER_CLASSIC,omitempty"`
				KnifeKillsMURDERCLASSIC                                  int      `json:"knife_kills_MURDER_CLASSIC,omitempty"`
				KnifeKillsHollywood                                      int      `json:"knife_kills_hollywood,omitempty"`
				ThrownKnifeKills                                         int      `json:"thrown_knife_kills,omitempty"`
				KillsHollywood                                           int      `json:"kills_hollywood,omitempty"`
				CoinsPickedupHollywoodMURDERCLASSIC                      int      `json:"coins_pickedup_hollywood_MURDER_CLASSIC,omitempty"`
				ThrownKnifeKillsHollywood                                int      `json:"thrown_knife_kills_hollywood,omitempty"`
				KillsHollywoodMURDERCLASSIC                              int      `json:"kills_hollywood_MURDER_CLASSIC,omitempty"`
				KnifeKillsHollywoodMURDERCLASSIC                         int      `json:"knife_kills_hollywood_MURDER_CLASSIC,omitempty"`
				GamesHollywood                                           int      `json:"games_hollywood,omitempty"`
				KillsMURDERCLASSIC                                       int      `json:"kills_MURDER_CLASSIC,omitempty"`
				ThrownKnifeKillsMURDERCLASSIC                            int      `json:"thrown_knife_kills_MURDER_CLASSIC,omitempty"`
				ThrownKnifeKillsHollywoodMURDERCLASSIC                   int      `json:"thrown_knife_kills_hollywood_MURDER_CLASSIC,omitempty"`
				GamesHollywoodMURDERCLASSIC                              int      `json:"games_hollywood_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererMURDERCLASSIC                             int      `json:"kills_as_murderer_MURDER_CLASSIC,omitempty"`
				KillsAsMurderer                                          int      `json:"kills_as_murderer,omitempty"`
				KnifeKills                                               int      `json:"knife_kills,omitempty"`
				KillsAsMurdererHollywoodMURDERCLASSIC                    int      `json:"kills_as_murderer_hollywood_MURDER_CLASSIC,omitempty"`
				CoinsPickedupHollywood                                   int      `json:"coins_pickedup_hollywood,omitempty"`
				DeathsHollywoodMURDERCLASSIC                             int      `json:"deaths_hollywood_MURDER_CLASSIC,omitempty"`
				Kills                                                    int      `json:"kills,omitempty"`
				KillsAsMurdererHollywood                                 int      `json:"kills_as_murderer_hollywood,omitempty"`
				DeathsHollywood                                          int      `json:"deaths_hollywood,omitempty"`
				DetectiveWinsTowerfall                                   int      `json:"detective_wins_towerfall,omitempty"`
				GamesTowerfallMURDERCLASSIC                              int      `json:"games_towerfall_MURDER_CLASSIC,omitempty"`
				CoinsPickedupTowerfall                                   int      `json:"coins_pickedup_towerfall,omitempty"`
				BowKillsTowerfall                                        int      `json:"bow_kills_towerfall,omitempty"`
				KillsTowerfall                                           int      `json:"kills_towerfall,omitempty"`
				WinsTowerfallMURDERCLASSIC                               int      `json:"wins_towerfall_MURDER_CLASSIC,omitempty"`
				WasHero                                                  int      `json:"was_hero,omitempty"`
				WasHeroMURDERCLASSIC                                     int      `json:"was_hero_MURDER_CLASSIC,omitempty"`
				BowKillsMURDERCLASSIC                                    int      `json:"bow_kills_MURDER_CLASSIC,omitempty"`
				DetectiveWinsTowerfallMURDERCLASSIC                      int      `json:"detective_wins_towerfall_MURDER_CLASSIC,omitempty"`
				GamesTowerfall                                           int      `json:"games_towerfall,omitempty"`
				BowKills                                                 int      `json:"bow_kills,omitempty"`
				WasHeroTowerfall                                         int      `json:"was_hero_towerfall,omitempty"`
				DetectiveWins                                            int      `json:"detective_wins,omitempty"`
				KillsTowerfallMURDERCLASSIC                              int      `json:"kills_towerfall_MURDER_CLASSIC,omitempty"`
				BowKillsTowerfallMURDERCLASSIC                           int      `json:"bow_kills_towerfall_MURDER_CLASSIC,omitempty"`
				DetectiveWinsMURDERCLASSIC                               int      `json:"detective_wins_MURDER_CLASSIC,omitempty"`
				WinsTowerfall                                            int      `json:"wins_towerfall,omitempty"`
				WasHeroTowerfallMURDERCLASSIC                            int      `json:"was_hero_towerfall_MURDER_CLASSIC,omitempty"`
				CoinsPickedupTowerfallMURDERCLASSIC                      int      `json:"coins_pickedup_towerfall_MURDER_CLASSIC,omitempty"`
				DeathsTowerfallMURDERCLASSIC                             int      `json:"deaths_towerfall_MURDER_CLASSIC,omitempty"`
				DeathsTowerfall                                          int      `json:"deaths_towerfall,omitempty"`
				GamesMURDERASSASSINS                                     int      `json:"games_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupArchives                                    int      `json:"coins_pickedup_archives,omitempty"`
				KnifeKillsArchives                                       int      `json:"knife_kills_archives,omitempty"`
				DeathsArchivesMURDERASSASSINS                            int      `json:"deaths_archives_MURDER_ASSASSINS,omitempty"`
				KnifeKillsArchivesMURDERASSASSINS                        int      `json:"knife_kills_archives_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupArchivesMURDERASSASSINS                     int      `json:"coins_pickedup_archives_MURDER_ASSASSINS,omitempty"`
				KillsArchives                                            int      `json:"kills_archives,omitempty"`
				KillsArchivesMURDERASSASSINS                             int      `json:"kills_archives_MURDER_ASSASSINS,omitempty"`
				DeathsMURDERASSASSINS                                    int      `json:"deaths_MURDER_ASSASSINS,omitempty"`
				GamesArchivesMURDERASSASSINS                             int      `json:"games_archives_MURDER_ASSASSINS,omitempty"`
				DeathsArchives                                           int      `json:"deaths_archives,omitempty"`
				KnifeKillsMURDERASSASSINS                                int      `json:"knife_kills_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupMURDERASSASSINS                             int      `json:"coins_pickedup_MURDER_ASSASSINS,omitempty"`
				KillsMURDERASSASSINS                                     int      `json:"kills_MURDER_ASSASSINS,omitempty"`
				DeathsAncientTombMURDERASSASSINS                         int      `json:"deaths_ancient_tomb_MURDER_ASSASSINS,omitempty"`
				GamesAncientTomb                                         int      `json:"games_ancient_tomb,omitempty"`
				GamesAncientTombMURDERASSASSINS                          int      `json:"games_ancient_tomb_MURDER_ASSASSINS,omitempty"`
				KnifeKillsAncientTomb                                    int      `json:"knife_kills_ancient_tomb,omitempty"`
				KnifeKillsAncientTombMURDERASSASSINS                     int      `json:"knife_kills_ancient_tomb_MURDER_ASSASSINS,omitempty"`
				DeathsAncientTomb                                        int      `json:"deaths_ancient_tomb,omitempty"`
				KillsAncientTomb                                         int      `json:"kills_ancient_tomb,omitempty"`
				KillsAncientTombMURDERASSASSINS                          int      `json:"kills_ancient_tomb_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupAncientTombMURDERASSASSINS                  int      `json:"coins_pickedup_ancient_tomb_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupAncientTomb                                 int      `json:"coins_pickedup_ancient_tomb,omitempty"`
				Showqueuebook                                            bool     `json:"showqueuebook,omitempty"`
				BowKillsCruiseShip                                       int      `json:"bow_kills_cruise_ship,omitempty"`
				ThrownKnifeKillsCruiseShip                               int      `json:"thrown_knife_kills_cruise_ship,omitempty"`
				ThrownKnifeKillsCruiseShipMURDERASSASSINS                int      `json:"thrown_knife_kills_cruise_ship_MURDER_ASSASSINS,omitempty"`
				KillsCruiseShip                                          int      `json:"kills_cruise_ship,omitempty"`
				ThrownKnifeKillsMURDERASSASSINS                          int      `json:"thrown_knife_kills_MURDER_ASSASSINS,omitempty"`
				GamesCruiseShipMURDERASSASSINS                           int      `json:"games_cruise_ship_MURDER_ASSASSINS,omitempty"`
				BowKillsCruiseShipMURDERASSASSINS                        int      `json:"bow_kills_cruise_ship_MURDER_ASSASSINS,omitempty"`
				KillsCruiseShipMURDERASSASSINS                           int      `json:"kills_cruise_ship_MURDER_ASSASSINS,omitempty"`
				WinsCruiseShipMURDERASSASSINS                            int      `json:"wins_cruise_ship_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupCruiseShip                                  int      `json:"coins_pickedup_cruise_ship,omitempty"`
				GamesCruiseShip                                          int      `json:"games_cruise_ship,omitempty"`
				KnifeKillsCruiseShip                                     int      `json:"knife_kills_cruise_ship,omitempty"`
				CoinsPickedupCruiseShipMURDERASSASSINS                   int      `json:"coins_pickedup_cruise_ship_MURDER_ASSASSINS,omitempty"`
				KnifeKillsCruiseShipMURDERASSASSINS                      int      `json:"knife_kills_cruise_ship_MURDER_ASSASSINS,omitempty"`
				WinsCruiseShip                                           int      `json:"wins_cruise_ship,omitempty"`
				WinsMURDERASSASSINS                                      int      `json:"wins_MURDER_ASSASSINS,omitempty"`
				BowKillsMURDERASSASSINS                                  int      `json:"bow_kills_MURDER_ASSASSINS,omitempty"`
				WinsMURDERHARDCORE                                       int      `json:"wins_MURDER_HARDCORE,omitempty"`
				GamesMURDERHARDCORE                                      int      `json:"games_MURDER_HARDCORE,omitempty"`
				GamesLibraryMURDERHARDCORE                               int      `json:"games_library_MURDER_HARDCORE,omitempty"`
				CoinsPickedupLibraryMURDERHARDCORE                       int      `json:"coins_pickedup_library_MURDER_HARDCORE,omitempty"`
				WinsLibraryMURDERHARDCORE                                int      `json:"wins_library_MURDER_HARDCORE,omitempty"`
				CoinsPickedupMURDERHARDCORE                              int      `json:"coins_pickedup_MURDER_HARDCORE,omitempty"`
				DeathsMURDERHARDCORE                                     int      `json:"deaths_MURDER_HARDCORE,omitempty"`
				CoinsPickedupCruiseShipMURDERHARDCORE                    int      `json:"coins_pickedup_cruise_ship_MURDER_HARDCORE,omitempty"`
				GamesCruiseShipMURDERHARDCORE                            int      `json:"games_cruise_ship_MURDER_HARDCORE,omitempty"`
				WinsCruiseShipMURDERHARDCORE                             int      `json:"wins_cruise_ship_MURDER_HARDCORE,omitempty"`
				DeathsCruiseShip                                         int      `json:"deaths_cruise_ship,omitempty"`
				DeathsCruiseShipMURDERHARDCORE                           int      `json:"deaths_cruise_ship_MURDER_HARDCORE,omitempty"`
				WinsHollywood                                            int      `json:"wins_hollywood,omitempty"`
				CoinsPickedupHollywoodMURDERHARDCORE                     int      `json:"coins_pickedup_hollywood_MURDER_HARDCORE,omitempty"`
				WinsHollywoodMURDERHARDCORE                              int      `json:"wins_hollywood_MURDER_HARDCORE,omitempty"`
				GamesHollywoodMURDERHARDCORE                             int      `json:"games_hollywood_MURDER_HARDCORE,omitempty"`
				DeathsLibrary                                            int      `json:"deaths_library,omitempty"`
				DeathsLibraryMURDERHARDCORE                              int      `json:"deaths_library_MURDER_HARDCORE,omitempty"`
				WinsGoldRush                                             int      `json:"wins_gold_rush,omitempty"`
				WinsGoldRushMURDERHARDCORE                               int      `json:"wins_gold_rush_MURDER_HARDCORE,omitempty"`
				CoinsPickedupGoldRushMURDERHARDCORE                      int      `json:"coins_pickedup_gold_rush_MURDER_HARDCORE,omitempty"`
				GamesGoldRushMURDERHARDCORE                              int      `json:"games_gold_rush_MURDER_HARDCORE,omitempty"`
				GrantedChests                                            int      `json:"granted_chests,omitempty"`
				MmChests                                                 int      `json:"mm_chests,omitempty"`
				DeathsHypixelWorld                                       int      `json:"deaths_hypixel_world,omitempty"`
				GamesHypixelWorldMURDERCLASSIC                           int      `json:"games_hypixel_world_MURDER_CLASSIC,omitempty"`
				DeathsHypixelWorldMURDERCLASSIC                          int      `json:"deaths_hypixel_world_MURDER_CLASSIC,omitempty"`
				GamesHypixelWorld                                        int      `json:"games_hypixel_world,omitempty"`
				CoinsPickedupHypixelWorldMURDERCLASSIC                   int      `json:"coins_pickedup_hypixel_world_MURDER_CLASSIC,omitempty"`
				CoinsPickedupHypixelWorld                                int      `json:"coins_pickedup_hypixel_world,omitempty"`
				KillsGoldRushMURDERCLASSIC                               int      `json:"kills_gold_rush_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererGoldRushMURDERCLASSIC                     int      `json:"kills_as_murderer_gold_rush_MURDER_CLASSIC,omitempty"`
				ThrownKnifeKillsGoldRush                                 int      `json:"thrown_knife_kills_gold_rush,omitempty"`
				KillsAsMurdererGoldRush                                  int      `json:"kills_as_murderer_gold_rush,omitempty"`
				ThrownKnifeKillsGoldRushMURDERCLASSIC                    int      `json:"thrown_knife_kills_gold_rush_MURDER_CLASSIC,omitempty"`
				KillsGoldRush                                            int      `json:"kills_gold_rush,omitempty"`
				KnifeKillsGoldRush                                       int      `json:"knife_kills_gold_rush,omitempty"`
				KnifeKillsGoldRushMURDERCLASSIC                          int      `json:"knife_kills_gold_rush_MURDER_CLASSIC,omitempty"`
				WinsGoldRushMURDERCLASSIC                                int      `json:"wins_gold_rush_MURDER_CLASSIC,omitempty"`
				WinsHollywoodMURDERCLASSIC                               int      `json:"wins_hollywood_MURDER_CLASSIC,omitempty"`
				MmChristmasChests                                        int      `json:"mm_christmas_chests,omitempty"`
				WinsHypixelWorldMURDERCLASSIC                            int      `json:"wins_hypixel_world_MURDER_CLASSIC,omitempty"`
				WinsHypixelWorld                                         int      `json:"wins_hypixel_world,omitempty"`
				GamesHeadquartersMURDERCLASSIC                           int      `json:"games_headquarters_MURDER_CLASSIC,omitempty"`
				WinsHeadquartersMURDERCLASSIC                            int      `json:"wins_headquarters_MURDER_CLASSIC,omitempty"`
				CoinsPickedupHeadquarters                                int      `json:"coins_pickedup_headquarters,omitempty"`
				WinsHeadquarters                                         int      `json:"wins_headquarters,omitempty"`
				GamesHeadquarters                                        int      `json:"games_headquarters,omitempty"`
				CoinsPickedupHeadquartersMURDERCLASSIC                   int      `json:"coins_pickedup_headquarters_MURDER_CLASSIC,omitempty"`
				DeathsHeadquartersMURDERCLASSIC                          int      `json:"deaths_headquarters_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererHeadquartersMURDERCLASSIC                 int      `json:"kills_as_murderer_headquarters_MURDER_CLASSIC,omitempty"`
				KillsHeadquarters                                        int      `json:"kills_headquarters,omitempty"`
				ThrownKnifeKillsHeadquartersMURDERCLASSIC                int      `json:"thrown_knife_kills_headquarters_MURDER_CLASSIC,omitempty"`
				DeathsHeadquarters                                       int      `json:"deaths_headquarters,omitempty"`
				KnifeKillsHeadquarters                                   int      `json:"knife_kills_headquarters,omitempty"`
				ThrownKnifeKillsHeadquarters                             int      `json:"thrown_knife_kills_headquarters,omitempty"`
				KillsAsMurdererHeadquarters                              int      `json:"kills_as_murderer_headquarters,omitempty"`
				KnifeKillsHeadquartersMURDERCLASSIC                      int      `json:"knife_kills_headquarters_MURDER_CLASSIC,omitempty"`
				KillsHeadquartersMURDERCLASSIC                           int      `json:"kills_headquarters_MURDER_CLASSIC,omitempty"`
				BowKillsLibraryMURDERCLASSIC                             int      `json:"bow_kills_library_MURDER_CLASSIC,omitempty"`
				KillsLibraryMURDERCLASSIC                                int      `json:"kills_library_MURDER_CLASSIC,omitempty"`
				DeathsLibraryMURDERCLASSIC                               int      `json:"deaths_library_MURDER_CLASSIC,omitempty"`
				BowKillsLibrary                                          int      `json:"bow_kills_library,omitempty"`
				KillsLibrary                                             int      `json:"kills_library,omitempty"`
				GamesAncientTombMURDERCLASSIC                            int      `json:"games_ancient_tomb_MURDER_CLASSIC,omitempty"`
				WinsAncientTombMURDERCLASSIC                             int      `json:"wins_ancient_tomb_MURDER_CLASSIC,omitempty"`
				CoinsPickedupAncientTombMURDERCLASSIC                    int      `json:"coins_pickedup_ancient_tomb_MURDER_CLASSIC,omitempty"`
				WinsAncientTomb                                          int      `json:"wins_ancient_tomb,omitempty"`
				KnifeKillsHypixelWorld                                   int      `json:"knife_kills_hypixel_world,omitempty"`
				ThrownKnifeKillsHypixelWorld                             int      `json:"thrown_knife_kills_hypixel_world,omitempty"`
				KillsAsMurdererHypixelWorld                              int      `json:"kills_as_murderer_hypixel_world,omitempty"`
				ThrownKnifeKillsHypixelWorldMURDERCLASSIC                int      `json:"thrown_knife_kills_hypixel_world_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererHypixelWorldMURDERCLASSIC                 int      `json:"kills_as_murderer_hypixel_world_MURDER_CLASSIC,omitempty"`
				KillsHypixelWorldMURDERCLASSIC                           int      `json:"kills_hypixel_world_MURDER_CLASSIC,omitempty"`
				KillsHypixelWorld                                        int      `json:"kills_hypixel_world,omitempty"`
				KnifeKillsHypixelWorldMURDERCLASSIC                      int      `json:"knife_kills_hypixel_world_MURDER_CLASSIC,omitempty"`
				WasHeroHeadquarters                                      int      `json:"was_hero_headquarters,omitempty"`
				WasHeroHeadquartersMURDERCLASSIC                         int      `json:"was_hero_headquarters_MURDER_CLASSIC,omitempty"`
				BowKillsHeadquartersMURDERCLASSIC                        int      `json:"bow_kills_headquarters_MURDER_CLASSIC,omitempty"`
				BowKillsHeadquarters                                     int      `json:"bow_kills_headquarters,omitempty"`
				CoinsPickedupCruiseShipMURDERCLASSIC                     int      `json:"coins_pickedup_cruise_ship_MURDER_CLASSIC,omitempty"`
				GamesCruiseShipMURDERCLASSIC                             int      `json:"games_cruise_ship_MURDER_CLASSIC,omitempty"`
				WinsCruiseShipMURDERCLASSIC                              int      `json:"wins_cruise_ship_MURDER_CLASSIC,omitempty"`
				CoinsPickedupArchivesMURDERCLASSIC                       int      `json:"coins_pickedup_archives_MURDER_CLASSIC,omitempty"`
				DetectiveWinsHeadquartersMURDERCLASSIC                   int      `json:"detective_wins_headquarters_MURDER_CLASSIC,omitempty"`
				DetectiveWinsHeadquarters                                int      `json:"detective_wins_headquarters,omitempty"`
				WasHeroTransport                                         int      `json:"was_hero_transport,omitempty"`
				KillsTransportMURDERCLASSIC                              int      `json:"kills_transport_MURDER_CLASSIC,omitempty"`
				WinsTransportMURDERCLASSIC                               int      `json:"wins_transport_MURDER_CLASSIC,omitempty"`
				GamesTransport                                           int      `json:"games_transport,omitempty"`
				KillsTransport                                           int      `json:"kills_transport,omitempty"`
				WinsTransport                                            int      `json:"wins_transport,omitempty"`
				DetectiveWinsTransport                                   int      `json:"detective_wins_transport,omitempty"`
				DetectiveWinsTransportMURDERCLASSIC                      int      `json:"detective_wins_transport_MURDER_CLASSIC,omitempty"`
				BowKillsTransportMURDERCLASSIC                           int      `json:"bow_kills_transport_MURDER_CLASSIC,omitempty"`
				CoinsPickedupTransport                                   int      `json:"coins_pickedup_transport,omitempty"`
				BowKillsTransport                                        int      `json:"bow_kills_transport,omitempty"`
				WasHeroTransportMURDERCLASSIC                            int      `json:"was_hero_transport_MURDER_CLASSIC,omitempty"`
				GamesTransportMURDERCLASSIC                              int      `json:"games_transport_MURDER_CLASSIC,omitempty"`
				CoinsPickedupTransportMURDERCLASSIC                      int      `json:"coins_pickedup_transport_MURDER_CLASSIC,omitempty"`
				DetectiveWinsHypixelWorld                                int      `json:"detective_wins_hypixel_world,omitempty"`
				DetectiveWinsHypixelWorldMURDERCLASSIC                   int      `json:"detective_wins_hypixel_world_MURDER_CLASSIC,omitempty"`
				DeathsCruiseShipMURDERCLASSIC                            int      `json:"deaths_cruise_ship_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererArchives                                  int      `json:"kills_as_murderer_archives,omitempty"`
				KillsArchivesMURDERCLASSIC                               int      `json:"kills_archives_MURDER_CLASSIC,omitempty"`
				KnifeKillsArchivesMURDERCLASSIC                          int      `json:"knife_kills_archives_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererArchivesMURDERCLASSIC                     int      `json:"kills_as_murderer_archives_MURDER_CLASSIC,omitempty"`
				WinsWidowSDenMURDERCLASSIC                               int      `json:"wins_widow's_den_MURDER_CLASSIC,omitempty"`
				WinsWidowSDen                                            int      `json:"wins_widow's_den,omitempty"`
				CoinsPickedupWidowSDen                                   int      `json:"coins_pickedup_widow's_den,omitempty"`
				GamesWidowSDenMURDERCLASSIC                              int      `json:"games_widow's_den_MURDER_CLASSIC,omitempty"`
				CoinsPickedupWidowSDenMURDERCLASSIC                      int      `json:"coins_pickedup_widow's_den_MURDER_CLASSIC,omitempty"`
				GamesWidowSDen                                           int      `json:"games_widow's_den,omitempty"`
				CoinsPickedupSanPeratico                                 int      `json:"coins_pickedup_san_peratico,omitempty"`
				WinsSanPeratico                                          int      `json:"wins_san_peratico,omitempty"`
				WinsSanPeraticoMURDERCLASSIC                             int      `json:"wins_san_peratico_MURDER_CLASSIC,omitempty"`
				GamesSanPeraticoMURDERCLASSIC                            int      `json:"games_san_peratico_MURDER_CLASSIC,omitempty"`
				CoinsPickedupSanPeraticoMURDERCLASSIC                    int      `json:"coins_pickedup_san_peratico_MURDER_CLASSIC,omitempty"`
				GamesSanPeratico                                         int      `json:"games_san_peratico,omitempty"`
				CoinsPickedupArchivesTopFloorMURDERCLASSIC               int      `json:"coins_pickedup_archives_top_floor_MURDER_CLASSIC,omitempty"`
				GamesArchivesTopFloor                                    int      `json:"games_archives_top_floor,omitempty"`
				CoinsPickedupArchivesTopFloor                            int      `json:"coins_pickedup_archives_top_floor,omitempty"`
				WinsArchivesTopFloorMURDERCLASSIC                        int      `json:"wins_archives_top_floor_MURDER_CLASSIC,omitempty"`
				GamesArchivesTopFloorMURDERCLASSIC                       int      `json:"games_archives_top_floor_MURDER_CLASSIC,omitempty"`
				WinsArchivesTopFloor                                     int      `json:"wins_archives_top_floor,omitempty"`
				CoinsPickedupHeadquartersMURDERASSASSINS                 int      `json:"coins_pickedup_headquarters_MURDER_ASSASSINS,omitempty"`
				GamesHeadquartersMURDERASSASSINS                         int      `json:"games_headquarters_MURDER_ASSASSINS,omitempty"`
				KillsHeadquartersMURDERASSASSINS                         int      `json:"kills_headquarters_MURDER_ASSASSINS,omitempty"`
				KnifeKillsHeadquartersMURDERASSASSINS                    int      `json:"knife_kills_headquarters_MURDER_ASSASSINS,omitempty"`
				WinsHeadquartersMURDERASSASSINS                          int      `json:"wins_headquarters_MURDER_ASSASSINS,omitempty"`
				BowKillsMURDERINFECTION                                  int      `json:"bow_kills_MURDER_INFECTION,omitempty"`
				BowKillsWidowSDen                                        int      `json:"bow_kills_widow's_den,omitempty"`
				BowKillsWidowSDenMURDERINFECTION                         int      `json:"bow_kills_widow's_den_MURDER_INFECTION,omitempty"`
				GamesMURDERINFECTION                                     int      `json:"games_MURDER_INFECTION,omitempty"`
				GamesWidowSDenMURDERINFECTION                            int      `json:"games_widow's_den_MURDER_INFECTION,omitempty"`
				KillsMURDERINFECTION                                     int      `json:"kills_MURDER_INFECTION,omitempty"`
				KillsAsSurvivor                                          int      `json:"kills_as_survivor,omitempty"`
				KillsAsSurvivorMURDERINFECTION                           int      `json:"kills_as_survivor_MURDER_INFECTION,omitempty"`
				KillsAsSurvivorWidowSDen                                 int      `json:"kills_as_survivor_widow's_den,omitempty"`
				KillsAsSurvivorWidowSDenMURDERINFECTION                  int      `json:"kills_as_survivor_widow's_den_MURDER_INFECTION,omitempty"`
				KillsWidowSDen                                           int      `json:"kills_widow's_den,omitempty"`
				KillsWidowSDenMURDERINFECTION                            int      `json:"kills_widow's_den_MURDER_INFECTION,omitempty"`
				LongestTimeAsSurvivorSeconds                             int      `json:"longest_time_as_survivor_seconds,omitempty"`
				LongestTimeAsSurvivorSecondsMURDERINFECTION              int      `json:"longest_time_as_survivor_seconds_MURDER_INFECTION,omitempty"`
				LongestTimeAsSurvivorSecondsWidowSDen                    int      `json:"longest_time_as_survivor_seconds_widow's_den,omitempty"`
				LongestTimeAsSurvivorSecondsWidowSDenMURDERINFECTION     int      `json:"longest_time_as_survivor_seconds_widow's_den_MURDER_INFECTION,omitempty"`
				SurvivorWins                                             int      `json:"survivor_wins,omitempty"`
				SurvivorWinsMURDERINFECTION                              int      `json:"survivor_wins_MURDER_INFECTION,omitempty"`
				SurvivorWinsWidowSDen                                    int      `json:"survivor_wins_widow's_den,omitempty"`
				SurvivorWinsWidowSDenMURDERINFECTION                     int      `json:"survivor_wins_widow's_den_MURDER_INFECTION,omitempty"`
				TotalTimeSurvivedSeconds                                 int      `json:"total_time_survived_seconds,omitempty"`
				TotalTimeSurvivedSecondsMURDERINFECTION                  int      `json:"total_time_survived_seconds_MURDER_INFECTION,omitempty"`
				TotalTimeSurvivedSecondsWidowSDen                        int      `json:"total_time_survived_seconds_widow's_den,omitempty"`
				TotalTimeSurvivedSecondsWidowSDenMURDERINFECTION         int      `json:"total_time_survived_seconds_widow's_den_MURDER_INFECTION,omitempty"`
				DeathsWidowSDen                                          int      `json:"deaths_widow's_den,omitempty"`
				DeathsWidowSDenMURDERCLASSIC                             int      `json:"deaths_widow's_den_MURDER_CLASSIC,omitempty"`
				Packages                                                 []string `json:"packages,omitempty"`
				QuickestDetectiveWinTimeSeconds                          int      `json:"quickest_detective_win_time_seconds,omitempty"`
				QuickestDetectiveWinTimeSecondsMURDERCLASSIC             int      `json:"quickest_detective_win_time_seconds_MURDER_CLASSIC,omitempty"`
				QuickestDetectiveWinTimeSecondsHeadquarters              int      `json:"quickest_detective_win_time_seconds_headquarters,omitempty"`
				QuickestDetectiveWinTimeSecondsHeadquartersMURDERCLASSIC int      `json:"quickest_detective_win_time_seconds_headquarters_MURDER_CLASSIC,omitempty"`
				CoinsPickedupVilla                                       int      `json:"coins_pickedup_villa,omitempty"`
				CoinsPickedupVillaMURDERCLASSIC                          int      `json:"coins_pickedup_villa_MURDER_CLASSIC,omitempty"`
				DeathsVilla                                              int      `json:"deaths_villa,omitempty"`
				DeathsVillaMURDERCLASSIC                                 int      `json:"deaths_villa_MURDER_CLASSIC,omitempty"`
				GamesVilla                                               int      `json:"games_villa,omitempty"`
				GamesVillaMURDERCLASSIC                                  int      `json:"games_villa_MURDER_CLASSIC,omitempty"`
				WinsVilla                                                int      `json:"wins_villa,omitempty"`
				WinsVillaMURDERCLASSIC                                   int      `json:"wins_villa_MURDER_CLASSIC,omitempty"`
				KillsAsMurdererTowerfall                                 int      `json:"kills_as_murderer_towerfall,omitempty"`
				KillsAsMurdererTowerfallMURDERCLASSIC                    int      `json:"kills_as_murderer_towerfall_MURDER_CLASSIC,omitempty"`
				KnifeKillsTowerfall                                      int      `json:"knife_kills_towerfall,omitempty"`
				KnifeKillsTowerfallMURDERCLASSIC                         int      `json:"knife_kills_towerfall_MURDER_CLASSIC,omitempty"`
				CoinsPickedupMountain                                    int      `json:"coins_pickedup_mountain,omitempty"`
				CoinsPickedupMountainMURDERCLASSIC                       int      `json:"coins_pickedup_mountain_MURDER_CLASSIC,omitempty"`
				DeathsMountain                                           int      `json:"deaths_mountain,omitempty"`
				DeathsMountainMURDERCLASSIC                              int      `json:"deaths_mountain_MURDER_CLASSIC,omitempty"`
				GamesMountain                                            int      `json:"games_mountain,omitempty"`
				GamesMountainMURDERCLASSIC                               int      `json:"games_mountain_MURDER_CLASSIC,omitempty"`
				DeathsDarkfall                                           int      `json:"deaths_darkfall,omitempty"`
				DeathsDarkfallMURDERCLASSIC                              int      `json:"deaths_darkfall_MURDER_CLASSIC,omitempty"`
				GamesDarkfall                                            int      `json:"games_darkfall,omitempty"`
				GamesDarkfallMURDERCLASSIC                               int      `json:"games_darkfall_MURDER_CLASSIC,omitempty"`
				CoinsPickedupAquarium                                    int      `json:"coins_pickedup_aquarium,omitempty"`
				CoinsPickedupAquariumMURDERCLASSIC                       int      `json:"coins_pickedup_aquarium_MURDER_CLASSIC,omitempty"`
				DeathsAquarium                                           int      `json:"deaths_aquarium,omitempty"`
				DeathsAquariumMURDERCLASSIC                              int      `json:"deaths_aquarium_MURDER_CLASSIC,omitempty"`
				GamesAquarium                                            int      `json:"games_aquarium,omitempty"`
				GamesAquariumMURDERCLASSIC                               int      `json:"games_aquarium_MURDER_CLASSIC,omitempty"`
				WinsMountain                                             int      `json:"wins_mountain,omitempty"`
				WinsMountainMURDERCLASSIC                                int      `json:"wins_mountain_MURDER_CLASSIC,omitempty"`
				CoinsPickedupMURDERDOUBLEUP                              int      `json:"coins_pickedup_MURDER_DOUBLE_UP,omitempty"`
				CoinsPickedupArchivesTopFloorMURDERDOUBLEUP              int      `json:"coins_pickedup_archives_top_floor_MURDER_DOUBLE_UP,omitempty"`
				DeathsMURDERDOUBLEUP                                     int      `json:"deaths_MURDER_DOUBLE_UP,omitempty"`
				DeathsArchivesTopFloor                                   int      `json:"deaths_archives_top_floor,omitempty"`
				DeathsArchivesTopFloorMURDERDOUBLEUP                     int      `json:"deaths_archives_top_floor_MURDER_DOUBLE_UP,omitempty"`
				GamesMURDERDOUBLEUP                                      int      `json:"games_MURDER_DOUBLE_UP,omitempty"`
				GamesArchivesTopFloorMURDERDOUBLEUP                      int      `json:"games_archives_top_floor_MURDER_DOUBLE_UP,omitempty"`
				DeathsTransport                                          int      `json:"deaths_transport,omitempty"`
				DeathsTransportMURDERASSASSINS                           int      `json:"deaths_transport_MURDER_ASSASSINS,omitempty"`
				GamesTransportMURDERASSASSINS                            int      `json:"games_transport_MURDER_ASSASSINS,omitempty"`
				DeathsHeadquartersMURDERASSASSINS                        int      `json:"deaths_headquarters_MURDER_ASSASSINS,omitempty"`
				CoinsPickedupTransportMURDERASSASSINS                    int      `json:"coins_pickedup_transport_MURDER_ASSASSINS,omitempty"`
				DeathsMountainMURDERASSASSINS                            int      `json:"deaths_mountain_MURDER_ASSASSINS,omitempty"`
				GamesMountainMURDERASSASSINS                             int      `json:"games_mountain_MURDER_ASSASSINS,omitempty"`
			} `json:"MurderMystery,omitempty"`
			BuildBattle struct {
				BuildbattleLoadout []string `json:"buildbattle_loadout,omitempty"`
				Coins              int      `json:"coins,omitempty"`
				CorrectGuesses     int      `json:"correct_guesses,omitempty"`
				GamesPlayed        int      `json:"games_played,omitempty"`
				MonthlyCoinsA      int      `json:"monthly_coins_a,omitempty"`
				Score              int      `json:"score,omitempty"`
				WeeklyCoinsA       int      `json:"weekly_coins_a,omitempty"`
				MonthlyCoinsB      int      `json:"monthly_coins_b,omitempty"`
				WeeklyCoinsB       int      `json:"weekly_coins_b,omitempty"`
				TeamsMostPoints    int      `json:"teams_most_points,omitempty"`
				TotalVotes         int      `json:"total_votes,omitempty"`
				Packages           []string `json:"packages,omitempty"`
			} `json:"BuildBattle,omitempty"`
			Duels struct {
				DuelsRecentlyPlayed             string   `json:"duels_recently_played,omitempty"`
				GamesPlayedDuels                int      `json:"games_played_duels,omitempty"`
				ShowLbOption                    string   `json:"show_lb_option,omitempty"`
				DuelsWinstreakBestBowDuel       int      `json:"duels_winstreak_best_bow_duel,omitempty"`
				RematchOption1                  string   `json:"rematch_option_1,omitempty"`
				ChatEnabled                     string   `json:"chat_enabled,omitempty"`
				BestTntGamesWinstreak           int      `json:"best_tnt_games_winstreak,omitempty"`
				CurrentWinstreak                int      `json:"current_winstreak,omitempty"`
				BestOverallWinstreak            int      `json:"best_overall_winstreak,omitempty"`
				CurrentTntGamesWinstreak        int      `json:"current_tnt_games_winstreak,omitempty"`
				DuelsWinstreakBestBowspleefDuel int      `json:"duels_winstreak_best_bowspleef_duel,omitempty"`
				BowspleefDuelRoundsPlayed       int      `json:"bowspleef_duel_rounds_played,omitempty"`
				BowspleefDuelWins               int      `json:"bowspleef_duel_wins,omitempty"`
				Wins                            int      `json:"wins,omitempty"`
				RoundsPlayed                    int      `json:"rounds_played,omitempty"`
				BowspleefDuelDeaths             int      `json:"bowspleef_duel_deaths,omitempty"`
				Losses                          int      `json:"losses,omitempty"`
				Deaths                          int      `json:"deaths,omitempty"`
				BowspleefDuelLosses             int      `json:"bowspleef_duel_losses,omitempty"`
				Selected1                       string   `json:"selected_1,omitempty"`
				Selected2                       string   `json:"selected_2,omitempty"`
				BridgeDoublesWins               int      `json:"bridge_doubles_wins,omitempty"`
				BridgeFourWins                  int      `json:"bridge_four_wins,omitempty"`
				BridgeDuelWins                  int      `json:"bridge_duel_wins,omitempty"`
				BridgeDuelDeaths                int      `json:"bridge_duel_deaths,omitempty"`
				Bridge3V3V3V3Losses             int      `json:"bridge_3v3v3v3_losses,omitempty"`
				Bridge3V3V3V3RoundsPlayed       int      `json:"bridge_3v3v3v3_rounds_played,omitempty"`
				BridgeDoublesDeaths             int      `json:"bridge_doubles_deaths,omitempty"`
				Goals                           int      `json:"goals,omitempty"`
				BridgeDuelGoals                 int      `json:"bridge_duel_goals,omitempty"`
				BridgeDoublesGoals              int      `json:"bridge_doubles_goals,omitempty"`
				BridgeFourRoundsPlayed          int      `json:"bridge_four_rounds_played,omitempty"`
				BridgeFourKills                 int      `json:"bridge_four_kills,omitempty"`
				BridgeDoublesRoundsPlayed       int      `json:"bridge_doubles_rounds_played,omitempty"`
				BridgeDoublesKills              int      `json:"bridge_doubles_kills,omitempty"`
				BridgeDoublesLosses             int      `json:"bridge_doubles_losses,omitempty"`
				BridgeDuelKills                 int      `json:"bridge_duel_kills,omitempty"`
				BridgeDuelLosses                int      `json:"bridge_duel_losses,omitempty"`
				BridgeDuelRoundsPlayed          int      `json:"bridge_duel_rounds_played,omitempty"`
				Packages                        []string `json:"packages,omitempty"`
				ClassicRookieTitlePrestige      int      `json:"classic_rookie_title_prestige,omitempty"`
				BlitzRookieTitlePrestige        int      `json:"blitz_rookie_title_prestige,omitempty"`
				BowRookieTitlePrestige          int      `json:"bow_rookie_title_prestige,omitempty"`
				SumoRookieTitlePrestige         int      `json:"sumo_rookie_title_prestige,omitempty"`
				TntGamesRookieTitlePrestige     int      `json:"tnt_games_rookie_title_prestige,omitempty"`
				BridgeRookieTitlePrestige       int      `json:"bridge_rookie_title_prestige,omitempty"`
				SkywarsRookieTitlePrestige      int      `json:"skywars_rookie_title_prestige,omitempty"`
				OpRookieTitlePrestige           int      `json:"op_rookie_title_prestige,omitempty"`
				UhcRookieTitlePrestige          int      `json:"uhc_rookie_title_prestige,omitempty"`
				ComboRookieTitlePrestige        int      `json:"combo_rookie_title_prestige,omitempty"`
				AllModesRookieTitlePrestige     int      `json:"all_modes_rookie_title_prestige,omitempty"`
				MegaWallsRookieTitlePrestige    int      `json:"mega_walls_rookie_title_prestige,omitempty"`
				NoDebuffRookieTitlePrestige     int      `json:"no_debuff_rookie_title_prestige,omitempty"`
				TournamentRookieTitlePrestige   int      `json:"tournament_rookie_title_prestige,omitempty"`
				Selected2New                    string   `json:"selected_2_new,omitempty"`
				Selected1New                    string   `json:"selected_1_new,omitempty"`
				SwDuelsKitNew3                  string   `json:"sw_duels_kit_new3,omitempty"`
				MeleeSwings                     int      `json:"melee_swings,omitempty"`
				SwDuelBlocksPlaced              int      `json:"sw_duel_blocks_placed,omitempty"`
				BlocksPlaced                    int      `json:"blocks_placed,omitempty"`
				MeleeHits                       int      `json:"melee_hits,omitempty"`
				SwDuelDamageDealt               int      `json:"sw_duel_damage_dealt,omitempty"`
				HealthRegenerated               int      `json:"health_regenerated,omitempty"`
				SwDuelRoundsPlayed              int      `json:"sw_duel_rounds_played,omitempty"`
				DamageDealt                     int      `json:"damage_dealt,omitempty"`
				SwDuelHealthRegenerated         int      `json:"sw_duel_health_regenerated,omitempty"`
				SwDuelMeleeHits                 int      `json:"sw_duel_melee_hits,omitempty"`
				SwDuelMeleeSwings               int      `json:"sw_duel_melee_swings,omitempty"`
				SwDuelBowHits                   int      `json:"sw_duel_bow_hits,omitempty"`
				BowShots                        int      `json:"bow_shots,omitempty"`
				BowHits                         int      `json:"bow_hits,omitempty"`
				SwDuelBowShots                  int      `json:"sw_duel_bow_shots,omitempty"`
				BestWinstreakModeSwDuel         int      `json:"best_winstreak_mode_sw_duel,omitempty"`
				MapsWonOn                       []string `json:"maps_won_on,omitempty"`
				CurrentWinstreakModeSwDuel      int      `json:"current_winstreak_mode_sw_duel,omitempty"`
				CurrentSkywarsWinstreak         int      `json:"current_skywars_winstreak,omitempty"`
				BestSkywarsWinstreak            int      `json:"best_skywars_winstreak,omitempty"`
				SwDuelArmorerKitWins            int      `json:"sw_duel_armorer_kit_wins,omitempty"`
				SwDuelKills                     int      `json:"sw_duel_kills,omitempty"`
				Kills                           int      `json:"kills,omitempty"`
				Coins                           int      `json:"coins,omitempty"`
				SwDuelKitWins                   int      `json:"sw_duel_kit_wins,omitempty"`
				KitWins                         int      `json:"kit_wins,omitempty"`
				ArmorerKitWins                  int      `json:"armorer_kit_wins,omitempty"`
				SwDuelWins                      int      `json:"sw_duel_wins,omitempty"`
				DuelsChests                     int      `json:"duels_chests,omitempty"`
				OpDuelDamageDealt               int      `json:"op_duel_damage_dealt,omitempty"`
				OpDuelHealthRegenerated         int      `json:"op_duel_health_regenerated,omitempty"`
				OpDuelMeleeHits                 int      `json:"op_duel_melee_hits,omitempty"`
				OpDuelMeleeSwings               int      `json:"op_duel_melee_swings,omitempty"`
				OpDuelRoundsPlayed              int      `json:"op_duel_rounds_played,omitempty"`
				BridgeMapWins                   []string `json:"bridgeMapWins,omitempty"`
				CurrentBridgeWinstreak          int      `json:"current_bridge_winstreak,omitempty"`
				CurrentWinstreakModeBridgeFour  int      `json:"current_winstreak_mode_bridge_four,omitempty"`
				BestWinstreakModeBridgeFour     int      `json:"best_winstreak_mode_bridge_four,omitempty"`
				BestBridgeWinstreak             int      `json:"best_bridge_winstreak,omitempty"`
				BridgeFourBridgeKills           int      `json:"bridge_four_bridge_kills,omitempty"`
				BridgeFourDamageDealt           int      `json:"bridge_four_damage_dealt,omitempty"`
				BridgeFourGoals                 int      `json:"bridge_four_goals,omitempty"`
				BridgeFourHealthRegenerated     int      `json:"bridge_four_health_regenerated,omitempty"`
				BridgeFourMeleeHits             int      `json:"bridge_four_melee_hits,omitempty"`
				BridgeFourMeleeSwings           int      `json:"bridge_four_melee_swings,omitempty"`
				BridgeKills                     int      `json:"bridge_kills,omitempty"`
				DuelsWinstreakBestBridgeFour    int      `json:"duels_winstreak_best_bridge_four,omitempty"`
				BridgeDeaths                    int      `json:"bridge_deaths,omitempty"`
				BridgeFourBlocksPlaced          int      `json:"bridge_four_blocks_placed,omitempty"`
				BridgeFourBowHits               int      `json:"bridge_four_bow_hits,omitempty"`
				BridgeFourBowShots              int      `json:"bridge_four_bow_shots,omitempty"`
				BridgeFourBridgeDeaths          int      `json:"bridge_four_bridge_deaths,omitempty"`
				UhcDuelDamageDealt              int      `json:"uhc_duel_damage_dealt,omitempty"`
				UhcDuelHealthRegenerated        int      `json:"uhc_duel_health_regenerated,omitempty"`
				UhcDuelMeleeHits                int      `json:"uhc_duel_melee_hits,omitempty"`
				UhcDuelMeleeSwings              int      `json:"uhc_duel_melee_swings,omitempty"`
				UhcDuelRoundsPlayed             int      `json:"uhc_duel_rounds_played,omitempty"`
				DuelsRecentlyPlayed2            string   `json:"duels_recently_played2,omitempty"`
				GoldenApplesEaten               int      `json:"golden_apples_eaten,omitempty"`
				UhcDuelBlocksPlaced             int      `json:"uhc_duel_blocks_placed,omitempty"`
				UhcDuelBowHits                  int      `json:"uhc_duel_bow_hits,omitempty"`
				UhcDuelBowShots                 int      `json:"uhc_duel_bow_shots,omitempty"`
				UhcDuelGoldenApplesEaten        int      `json:"uhc_duel_golden_apples_eaten,omitempty"`
				BowspleefDuelBowShots           int      `json:"bowspleef_duel_bow_shots,omitempty"`
				MovedToRedis2                   bool     `json:"moved_to_redis_2,omitempty"`
				BoxingRookieTitlePrestige       int      `json:"boxing_rookie_title_prestige,omitempty"`
				ParkourRookieTitlePrestige      int      `json:"parkour_rookie_title_prestige,omitempty"`
				MovedToRedis3                   bool     `json:"moved_to_redis_3,omitempty"`
			} `json:"Duels,omitempty"`
			SkyBlock struct {
				Profiles struct {
					Three84A029294Fc445E863F2C42Fe9709Cb struct {
						ProfileID string `json:"profile_id,omitempty"`
						CuteName  string `json:"cute_name,omitempty"`
					} `json:"384a029294fc445e863f2c42fe9709cb,omitempty"`
					FiveDf1948A4824440Ea1D54Fed3Cd04C47 struct {
						ProfileID string `json:"profile_id,omitempty"`
						CuteName  string `json:"cute_name,omitempty"`
					} `json:"5df1948a4824440ea1d54fed3cd04c47,omitempty"`
					Zero714B118B4424A8D9B1515A04636A496 struct {
						ProfileID string `json:"profile_id,omitempty"`
						CuteName  string `json:"cute_name,omitempty"`
					} `json:"0714b118b4424a8d9b1515a04636a496,omitempty"`
					Four38Ee92Df729438Aac22D4447B432745 struct {
						ProfileID string `json:"profile_id,omitempty"`
						CuteName  string `json:"cute_name,omitempty"`
					} `json:"438ee92df729438aac22d4447b432745,omitempty"`
					Ec1827Cfeae8436A947Cf0595C0345E3 struct {
						ProfileID string `json:"profile_id,omitempty"`
						CuteName  string `json:"cute_name,omitempty"`
					} `json:"ec1827cfeae8436a947cf0595c0345e3,omitempty"`
				} `json:"profiles,omitempty"`
			} `json:"SkyBlock,omitempty"`
			Pit struct {
				Profile struct {
					MovedAchievements1 bool  `json:"moved_achievements_1,omitempty"`
					OutgoingOffers     []any `json:"outgoing_offers,omitempty"`
					MovedAchievements2 bool  `json:"moved_achievements_2,omitempty"`
					LeaderboardStats   struct {
					} `json:"leaderboard_stats,omitempty"`
					LastSave  int64 `json:"last_save,omitempty"`
					KingQuest struct {
						Kills int `json:"kills,omitempty"`
					} `json:"king_quest,omitempty"`
					LastPassiveXp int64 `json:"last_passive_xp,omitempty"`
					InvArmor      struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"inv_armor,omitempty"`
					ItemStash struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"item_stash,omitempty"`
					LoginMessages   []any `json:"login_messages,omitempty"`
					HotbarFavorites []int `json:"hotbar_favorites,omitempty"`
					SpireStashInv   struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"spire_stash_inv,omitempty"`
					Xp          int `json:"xp,omitempty"`
					InvContents struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"inv_contents,omitempty"`
					ZeroPointThreeGoldTransfer bool `json:"zero_point_three_gold_transfer,omitempty"`
					InvEnderchest              struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"inv_enderchest,omitempty"`
					Bounties        []any `json:"bounties,omitempty"`
					SpireStashArmor struct {
						Type int   `json:"type,omitempty"`
						Data []int `json:"data,omitempty"`
					} `json:"spire_stash_armor,omitempty"`
					Cash                float64 `json:"cash,omitempty"`
					CashDuringPrestige0 float64 `json:"cash_during_prestige_0,omitempty"`
				} `json:"profile,omitempty"`
				PitStatsPtl struct {
					CashEarned          int `json:"cash_earned,omitempty"`
					Joins               int `json:"joins,omitempty"`
					JumpedIntoPit       int `json:"jumped_into_pit,omitempty"`
					DamageDealt         int `json:"damage_dealt,omitempty"`
					DamageReceived      int `json:"damage_received,omitempty"`
					Deaths              int `json:"deaths,omitempty"`
					LeftClicks          int `json:"left_clicks,omitempty"`
					MeleeDamageDealt    int `json:"melee_damage_dealt,omitempty"`
					MeleeDamageReceived int `json:"melee_damage_received,omitempty"`
					PlaytimeMinutes     int `json:"playtime_minutes,omitempty"`
					SwordHits           int `json:"sword_hits,omitempty"`
					Assists             int `json:"assists,omitempty"`
					BowDamageReceived   int `json:"bow_damage_received,omitempty"`
					GappleEaten         int `json:"gapple_eaten,omitempty"`
					Kills               int `json:"kills,omitempty"`
					LaunchedByLaunchers int `json:"launched_by_launchers,omitempty"`
					MaxStreak           int `json:"max_streak,omitempty"`
					ArrowHits           int `json:"arrow_hits,omitempty"`
					ArrowsFired         int `json:"arrows_fired,omitempty"`
					BowDamageDealt      int `json:"bow_damage_dealt,omitempty"`
				} `json:"pit_stats_ptl,omitempty"`
			} `json:"Pit,omitempty"`
			Housing struct {
				Packages []string `json:"packages,omitempty"`
			} `json:"Housing,omitempty"`
			MainLobby struct {
				QuestNPCTutorials struct {
					ZoneSpawn          bool `json:"zone_spawn,omitempty"`
					ZoneCastle         bool `json:"zone_castle,omitempty"`
					ZoneVillage        bool `json:"zone_village,omitempty"`
					Dockmaster         bool `json:"dockmaster,omitempty"`
					ZoneAncientStatues bool `json:"zone_ancient_statues,omitempty"`
					ZoneLabyrinth      bool `json:"zone_labyrinth,omitempty"`
					ZonePort           bool `json:"zone_port,omitempty"`
					SummerGuide2022    bool `json:"summer_guide_2022,omitempty"`
					ZoneFarm           bool `json:"zone_farm,omitempty"`
					ZoneMines          bool `json:"zone_mines,omitempty"`
					EasterGuide2023    bool `json:"easter_guide_2023,omitempty"`
					ZoneFishingHut     bool `json:"zone_fishing_hut,omitempty"`
					ZoneRuins          bool `json:"zone_ruins,omitempty"`
				} `json:"questNPCTutorials,omitempty"`
				DiscoveredZones struct {
					Castle         bool `json:"castle,omitempty"`
					Village        bool `json:"village,omitempty"`
					AncientStatues bool `json:"ancient_statues,omitempty"`
					Labyrinth      bool `json:"labyrinth,omitempty"`
					Port           bool `json:"port,omitempty"`
					Farm           bool `json:"farm,omitempty"`
					Mines          bool `json:"mines,omitempty"`
					FishingHut     bool `json:"fishing_hut,omitempty"`
					Ruins          bool `json:"ruins,omitempty"`
				} `json:"discoveredZones,omitempty"`
				HistoricalRecords struct {
					PtlRelease         bool `json:"ptl_release,omitempty"`
					BuildBattleRelease bool `json:"build_battle_release,omitempty"`
					SkyblockRelease    bool `json:"skyblock_release,omitempty"`
					AugustMegaUpdate   bool `json:"august_mega_update,omitempty"`
					GuildUpdate        bool `json:"guild_update,omitempty"`
					SmpBeta            bool `json:"smp_beta,omitempty"`
				} `json:"historicalRecords,omitempty"`
				Fishing struct {
					Stats struct {
						Num2023 struct {
							Easter struct {
								Water struct {
									Fish     int `json:"fish,omitempty"`
									Treasure int `json:"treasure,omitempty"`
									Junk     int `json:"junk,omitempty"`
								} `json:"water,omitempty"`
							} `json:"easter,omitempty"`
						} `json:"2023,omitempty"`
						Permanent struct {
							Water struct {
								Fish     int `json:"fish,omitempty"`
								Junk     int `json:"junk,omitempty"`
								Treasure int `json:"treasure,omitempty"`
							} `json:"water,omitempty"`
							Individual struct {
								Fish struct {
									Clownfish  int `json:"clownfish,omitempty"`
									Salmon     int `json:"salmon,omitempty"`
									Cod        int `json:"cod,omitempty"`
									Pufferfish int `json:"pufferfish,omitempty"`
								} `json:"fish,omitempty"`
								Junk struct {
									String           int `json:"string,omitempty"`
									RottenFlesh      int `json:"rotten_flesh,omitempty"`
									Bowl             int `json:"bowl,omitempty"`
									WaterBottle      int `json:"water_bottle,omitempty"`
									Leather          int `json:"leather,omitempty"`
									Bone             int `json:"bone,omitempty"`
									InkSac           int `json:"ink_sac,omitempty"`
									BrokenFishingRod int `json:"broken_fishing_rod,omitempty"`
									SoggyPaper       int `json:"soggy_paper,omitempty"`
									TripwireHook     int `json:"tripwire_hook,omitempty"`
									Stick            int `json:"stick,omitempty"`
									RabbitHide       int `json:"rabbit_hide,omitempty"`
								} `json:"junk,omitempty"`
								Treasure struct {
									Saddle        int `json:"saddle,omitempty"`
									EnchantedBow  int `json:"enchanted_bow,omitempty"`
									NameTag       int `json:"name_tag,omitempty"`
									DiamondSword  int `json:"diamond_sword,omitempty"`
									EnchantedBook int `json:"enchanted_book,omitempty"`
									Compass       int `json:"compass,omitempty"`
								} `json:"treasure,omitempty"`
							} `json:"individual,omitempty"`
						} `json:"permanent,omitempty"`
					} `json:"stats,omitempty"`
					Fireproofing struct {
						Scales int `json:"scales,omitempty"`
						Flame  int `json:"flame,omitempty"`
					} `json:"fireproofing,omitempty"`
					Orbs struct {
						Selene    int `json:"selene,omitempty"`
						Helios    int `json:"helios,omitempty"`
						Nyx       int `json:"nyx,omitempty"`
						Aphrodite int `json:"aphrodite,omitempty"`
					} `json:"orbs,omitempty"`
					SpecialFish struct {
						EggTheFish           bool `json:"egg_the_fish,omitempty"`
						FishMongerSuitHelmet bool `json:"fish_monger_suit_helmet,omitempty"`
					} `json:"special_fish,omitempty"`
					Enchants struct {
						DumpsterDiver struct {
							Level int `json:"level,omitempty"`
						} `json:"dumpster_diver,omitempty"`
					} `json:"enchants,omitempty"`
				} `json:"fishing,omitempty"`
				Packages []string `json:"packages,omitempty"`
			} `json:"MainLobby,omitempty"`
			WoolGames struct {
				Progression struct {
					AvailableLayers int     `json:"available_layers,omitempty"`
					Experience      float64 `json:"experience,omitempty"`
				} `json:"progression,omitempty"`
				Coins    int `json:"coins,omitempty"`
				WoolWars struct {
					SelectedClass string `json:"selected_class,omitempty"`
					Stats         struct {
						Classes struct {
							Archer struct {
								Deaths int `json:"deaths,omitempty"`
							} `json:"archer,omitempty"`
							Swordsman struct {
								Deaths         int `json:"deaths,omitempty"`
								Kills          int `json:"kills,omitempty"`
								Assists        int `json:"assists,omitempty"`
								PowerupsGotten int `json:"powerups_gotten,omitempty"`
								BlocksBroken   int `json:"blocks_broken,omitempty"`
								WoolPlaced     int `json:"wool_placed,omitempty"`
							} `json:"swordsman,omitempty"`
							Tank struct {
								Deaths int `json:"deaths,omitempty"`
							} `json:"tank,omitempty"`
							Golem struct {
								BlocksBroken   int `json:"blocks_broken,omitempty"`
								Deaths         int `json:"deaths,omitempty"`
								PowerupsGotten int `json:"powerups_gotten,omitempty"`
								WoolPlaced     int `json:"wool_placed,omitempty"`
							} `json:"golem,omitempty"`
						} `json:"classes,omitempty"`
						Deaths         int `json:"deaths,omitempty"`
						GamesPlayed    int `json:"games_played,omitempty"`
						Kills          int `json:"kills,omitempty"`
						Assists        int `json:"assists,omitempty"`
						BlocksBroken   int `json:"blocks_broken,omitempty"`
						PowerupsGotten int `json:"powerups_gotten,omitempty"`
						WoolPlaced     int `json:"wool_placed,omitempty"`
						Tourney        struct {
							WoolWars1 struct {
								Assists        int `json:"assists,omitempty"`
								Deaths         int `json:"deaths,omitempty"`
								GamesPlayed    int `json:"games_played,omitempty"`
								BlocksBroken   int `json:"blocks_broken,omitempty"`
								Kills          int `json:"kills,omitempty"`
								PowerupsGotten int `json:"powerups_gotten,omitempty"`
								WoolPlaced     int `json:"wool_placed,omitempty"`
								Wins           int `json:"wins,omitempty"`
							} `json:"wool_wars_1,omitempty"`
						} `json:"tourney,omitempty"`
						Wins int `json:"wins,omitempty"`
					} `json:"stats,omitempty"`
				} `json:"wool_wars,omitempty"`
				Packages      []string `json:"packages,omitempty"`
				LastTourneyAd int64    `json:"lastTourneyAd,omitempty"`
			} `json:"WoolGames,omitempty"`
		} `json:"stats,omitempty"`
		ThanksSent       int    `json:"thanksSent,omitempty"`
		TimePlaying      int    `json:"timePlaying,omitempty"`
		TournamentTokens int    `json:"tournamentTokens,omitempty"`
		UUID             string `json:"uuid,omitempty"`
		Eugene           struct {
			DailyTwoKExp int64 `json:"dailyTwoKExp,omitempty"`
		} `json:"eugene,omitempty"`
		Language string `json:"language,omitempty"`
		Voting   struct {
			Total               int   `json:"total,omitempty"`
			TotalMcsorg         int   `json:"total_mcsorg,omitempty"`
			SecondaryMcsorg     int   `json:"secondary_mcsorg,omitempty"`
			LastMcsorg          int64 `json:"last_mcsorg,omitempty"`
			VotesToday          int   `json:"votesToday,omitempty"`
			LastVote            int64 `json:"last_vote,omitempty"`
			SecondaryMcsl       int   `json:"secondary_mcsl,omitempty"`
			TotalMcsl           int   `json:"total_mcsl,omitempty"`
			LastMcsl            int64 `json:"last_mcsl,omitempty"`
			TotalMcmp           int   `json:"total_mcmp,omitempty"`
			SecondaryMcmp       int   `json:"secondary_mcmp,omitempty"`
			LastMcmp            int64 `json:"last_mcmp,omitempty"`
			TotalTopg           int   `json:"total_topg,omitempty"`
			SecondaryTopg       int   `json:"secondary_topg,omitempty"`
			LastTopg            int64 `json:"last_topg,omitempty"`
			TotalMinestatus     int   `json:"total_minestatus,omitempty"`
			SecondaryMinestatus int   `json:"secondary_minestatus,omitempty"`
			LastMinestatus      int64 `json:"last_minestatus,omitempty"`
			TotalMcipl          int   `json:"total_mcipl,omitempty"`
			SecondaryMcipl      int   `json:"secondary_mcipl,omitempty"`
			LastMcipl           int64 `json:"last_mcipl,omitempty"`
			TotalMcf            int   `json:"total_mcf,omitempty"`
			SecondaryMcf        int   `json:"secondary_mcf,omitempty"`
			LastMcf             int64 `json:"last_mcf,omitempty"`
		} `json:"voting,omitempty"`
		PetConsumables struct {
			Feather      int `json:"FEATHER,omitempty"`
			Bread        int `json:"BREAD,omitempty"`
			Cake         int `json:"CAKE,omitempty"`
			CookedBeef   int `json:"COOKED_BEEF,omitempty"`
			RedRose      int `json:"RED_ROSE,omitempty"`
			WaterBucket  int `json:"WATER_BUCKET,omitempty"`
			Melon        int `json:"MELON,omitempty"`
			Stick        int `json:"STICK,omitempty"`
			WoodSword    int `json:"WOOD_SWORD,omitempty"`
			MilkBucket   int `json:"MILK_BUCKET,omitempty"`
			GoldRecord   int `json:"GOLD_RECORD,omitempty"`
			Pork         int `json:"PORK,omitempty"`
			Leash        int `json:"LEASH,omitempty"`
			LavaBucket   int `json:"LAVA_BUCKET,omitempty"`
			CarrotItem   int `json:"CARROT_ITEM,omitempty"`
			RottenFlesh  int `json:"ROTTEN_FLESH,omitempty"`
			SlimeBall    int `json:"SLIME_BALL,omitempty"`
			RawFish      int `json:"RAW_FISH,omitempty"`
			Bone         int `json:"BONE,omitempty"`
			MagmaCream   int `json:"MAGMA_CREAM,omitempty"`
			Cookie       int `json:"COOKIE,omitempty"`
			Apple        int `json:"APPLE,omitempty"`
			Wheat        int `json:"WHEAT,omitempty"`
			MushroomSoup int `json:"MUSHROOM_SOUP,omitempty"`
			HayBlock     int `json:"HAY_BLOCK,omitempty"`
			PumpkinPie   int `json:"PUMPKIN_PIE,omitempty"`
			BakedPotato  int `json:"BAKED_POTATO,omitempty"`
		} `json:"petConsumables,omitempty"`
		HousingMeta struct {
			AllowedBlocks      []string `json:"allowedBlocks,omitempty"`
			TutorialStep       string   `json:"tutorialStep,omitempty"`
			Packages           []string `json:"packages,omitempty"`
			GivenCookies104981 []string `json:"given_cookies_104981,omitempty"`
			FirstHouseJoinMs   int64    `json:"firstHouseJoinMs,omitempty"`
			GivenCookies104995 []string `json:"given_cookies_104995,omitempty"`
			GivenCookies105002 []string `json:"given_cookies_105002,omitempty"`
			GivenCookies105013 []string `json:"given_cookies_105013,omitempty"`
			GivenCookies105019 []string `json:"given_cookies_105019,omitempty"`
			GivenCookies105021 []string `json:"given_cookies_105021,omitempty"`
			PlotSize           string   `json:"plotSize,omitempty"`
			PlayerSettings     struct {
				Border string `json:"BORDER,omitempty"`
			} `json:"playerSettings,omitempty"`
			GivenCookies105080 []string `json:"given_cookies_105080,omitempty"`
			GivenCookies105104 []string `json:"given_cookies_105104,omitempty"`
		} `json:"housingMeta,omitempty"`
		VanityMeta struct {
			Packages []string `json:"packages,omitempty"`
		} `json:"vanityMeta,omitempty"`
		RewardConsumed          bool   `json:"rewardConsumed,omitempty"`
		LastAdsenseGenerateTime int64  `json:"lastAdsenseGenerateTime,omitempty"`
		LastClaimedReward       int64  `json:"lastClaimedReward,omitempty"`
		TotalRewards            int    `json:"totalRewards,omitempty"`
		TotalDailyRewards       int    `json:"totalDailyRewards,omitempty"`
		RewardStreak            int    `json:"rewardStreak,omitempty"`
		RewardScore             int    `json:"rewardScore,omitempty"`
		RewardHighScore         int    `json:"rewardHighScore,omitempty"`
		AdsenseTokens           int    `json:"adsense_tokens,omitempty"`
		FlashingSalePopup       int64  `json:"flashingSalePopup,omitempty"`
		FlashingSaleOpens       int    `json:"flashingSaleOpens,omitempty"`
		FlashingSaleClicks      int    `json:"flashingSaleClicks,omitempty"`
		FlashingSalePoppedUp    int    `json:"flashingSalePoppedUp,omitempty"`
		NewPackageRank          string `json:"newPackageRank,omitempty"`
		LevelUpVIP              int64  `json:"levelUp_VIP,omitempty"`
		SpecialtyCooldowns      struct {
			Vip3     bool `json:"VIP3,omitempty"`
			Vip0     bool `json:"VIP0,omitempty"`
			Vip1     bool `json:"VIP1,omitempty"`
			Vip2     bool `json:"VIP2,omitempty"`
			Vip6     bool `json:"VIP6,omitempty"`
			Vip5     bool `json:"VIP5,omitempty"`
			Vip4     bool `json:"VIP4,omitempty"`
			VipPlus0 bool `json:"VIP_PLUS0,omitempty"`
			VipPlus1 bool `json:"VIP_PLUS1,omitempty"`
			VipPlus4 bool `json:"VIP_PLUS4,omitempty"`
			VipPlus5 bool `json:"VIP_PLUS5,omitempty"`
			VipPlus2 bool `json:"VIP_PLUS2,omitempty"`
			VipPlus3 bool `json:"VIP_PLUS3,omitempty"`
			VipPlus6 bool `json:"VIP_PLUS6,omitempty"`
		} `json:"specialtyCooldowns,omitempty"`
		LevelUpVIPPLUS int64 `json:"levelUp_VIP_PLUS,omitempty"`
		PetStats       struct {
			SheepYellow struct {
				Name     string `json:"name,omitempty"`
				Exercise struct {
					Value     int   `json:"value,omitempty"`
					Timestamp int64 `json:"timestamp,omitempty"`
				} `json:"EXERCISE,omitempty"`
				Thirst struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"THIRST,omitempty"`
				Hunger struct {
					Value     int   `json:"value,omitempty"`
					Timestamp int64 `json:"timestamp,omitempty"`
				} `json:"HUNGER,omitempty"`
				Experience int `json:"experience,omitempty"`
			} `json:"SHEEP_YELLOW,omitempty"`
			CowBaby struct {
				Thirst struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"THIRST,omitempty"`
				Exercise struct {
					Value     int   `json:"value,omitempty"`
					Timestamp int64 `json:"timestamp,omitempty"`
				} `json:"EXERCISE,omitempty"`
				Hunger struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"HUNGER,omitempty"`
				Name       string `json:"name,omitempty"`
				Experience int    `json:"experience,omitempty"`
			} `json:"COW_BABY,omitempty"`
			WildOcelotBaby struct {
				Hunger struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"HUNGER,omitempty"`
				Thirst struct {
					Value     int   `json:"value,omitempty"`
					Timestamp int64 `json:"timestamp,omitempty"`
				} `json:"THIRST,omitempty"`
				Exercise struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"EXERCISE,omitempty"`
				Experience int    `json:"experience,omitempty"`
				Name       string `json:"name,omitempty"`
			} `json:"WILD_OCELOT_BABY,omitempty"`
			Zombie struct {
				Name string `json:"name,omitempty"`
			} `json:"ZOMBIE,omitempty"`
			Totem struct {
				Name string `json:"name,omitempty"`
			} `json:"TOTEM,omitempty"`
			Pig struct {
				Name     string `json:"name,omitempty"`
				Exercise struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"EXERCISE,omitempty"`
				Hunger struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"HUNGER,omitempty"`
				Thirst struct {
					Timestamp int64 `json:"timestamp,omitempty"`
					Value     int   `json:"value,omitempty"`
				} `json:"THIRST,omitempty"`
				Experience int `json:"experience,omitempty"`
			} `json:"PIG,omitempty"`
		} `json:"petStats,omitempty"`
		PetJourneyTimestamp int64 `json:"petJourneyTimestamp,omitempty"`
		SocialMedia         struct {
			Twitter string `json:"TWITTER,omitempty"`
			Links   struct {
				Youtube string `json:"YOUTUBE,omitempty"`
				Discord string `json:"DISCORD,omitempty"`
			} `json:"links,omitempty"`
			Prompt bool `json:"prompt,omitempty"`
		} `json:"socialMedia,omitempty"`
		UserLanguage        string   `json:"userLanguage,omitempty"`
		SantaQuestStarted   bool     `json:"SANTA_QUEST_STARTED,omitempty"`
		LastLogout          int64    `json:"lastLogout,omitempty"`
		NetworkUpdateBook   string   `json:"network_update_book,omitempty"`
		AchievementTracking []string `json:"achievementTracking,omitempty"`
		GuildNotifications  bool     `json:"guildNotifications,omitempty"`
		GiftingMeta         struct {
			RealBundlesReceived int      `json:"realBundlesReceived,omitempty"`
			BundlesReceived     int      `json:"bundlesReceived,omitempty"`
			GiftsGiven          int      `json:"giftsGiven,omitempty"`
			BundlesGiven        int      `json:"bundlesGiven,omitempty"`
			RealBundlesGiven    int      `json:"realBundlesGiven,omitempty"`
			Milestones          []string `json:"milestones,omitempty"`
		} `json:"giftingMeta,omitempty"`
		FortuneBuff     int `json:"fortuneBuff,omitempty"`
		AchievementSync struct {
			QuakeTiered int `json:"quake_tiered,omitempty"`
		} `json:"achievementSync,omitempty"`
		Challenges struct {
			AllTime struct {
				BEDWARSSupport                    int `json:"BEDWARS__support,omitempty"`
				BEDWARSOffensive                  int `json:"BEDWARS__offensive,omitempty"`
				MCGOKillingSpreeChallenge         int `json:"MCGO__killing_spree_challenge,omitempty"`
				ARCADEGalaxyWarsChallenge         int `json:"ARCADE__galaxy_wars_challenge,omitempty"`
				BEDWARSDefensive                  int `json:"BEDWARS__defensive,omitempty"`
				MCGOKnifeChallenge                int `json:"MCGO__knife_challenge,omitempty"`
				TNTGAMESTntTagChallenge           int `json:"TNTGAMES__tnt_tag_challenge,omitempty"`
				TNTGAMESPvpRunChallenge           int `json:"TNTGAMES__pvp_run_challenge,omitempty"`
				TNTGAMESBowSpleefChallenge        int `json:"TNTGAMES__bow_spleef_challenge,omitempty"`
				TNTGAMESTntWizardChallenge        int `json:"TNTGAMES__tnt_wizard_challenge,omitempty"`
				MURDERMYSTERYHero                 int `json:"MURDER_MYSTERY__hero,omitempty"`
				ARCADEFarmHuntChallenge           int `json:"ARCADE__farm_hunt_challenge,omitempty"`
				ARCADEHideAndSeekChallenge        int `json:"ARCADE__hide_and_seek_challenge,omitempty"`
				DUELSFeedTheVoidChallenge         int `json:"DUELS__feed_the_void_challenge,omitempty"`
				PAINTBALLFinishChallenge          int `json:"PAINTBALL__finish_challenge,omitempty"`
				QUAKECRAFTDonTBlinkChallenge      int `json:"QUAKECRAFT__don't_blink_challenge,omitempty"`
				QUAKECRAFTComboChallenge          int `json:"QUAKECRAFT__combo_challenge,omitempty"`
				MURDERMYSTERYSerialKiller         int `json:"MURDER_MYSTERY__serial_killer,omitempty"`
				MCGOGrenadeChallenge              int `json:"MCGO__grenade_challenge,omitempty"`
				DUELSTeamsChallenge               int `json:"DUELS__teams_challenge,omitempty"`
				PAINTBALLKillingSpreeChallenge    int `json:"PAINTBALL__killing_spree_challenge,omitempty"`
				PAINTBALLKillStreakChallenge      int `json:"PAINTBALL__kill_streak_challenge,omitempty"`
				QUAKECRAFTKillingStreakChallenge  int `json:"QUAKECRAFT__killing_streak_challenge,omitempty"`
				QUAKECRAFTPowerupChallenge        int `json:"QUAKECRAFT__powerup_challenge,omitempty"`
				TNTGAMESTntWizardsChallenge       int `json:"TNTGAMES__tnt_wizards_challenge,omitempty"`
				TNTGAMESTntRunChallenge           int `json:"TNTGAMES__tnt_run_challenge,omitempty"`
				ARCADEPartyGamesChallenge         int `json:"ARCADE__party_games_challenge,omitempty"`
				MURDERMYSTERYSherlock             int `json:"MURDER_MYSTERY__sherlock,omitempty"`
				ARCADEZombiesChallenge            int `json:"ARCADE__zombies_challenge,omitempty"`
				WALLSLootingChallenge             int `json:"WALLS__looting_challenge,omitempty"`
				SUPERSMASHSmashChallenge          int `json:"SUPER_SMASH__smash_challenge,omitempty"`
				SUPERSMASHLeaderboardChallenge    int `json:"SUPER_SMASH__leaderboard_challenge,omitempty"`
				ARCADECreeperAttackChallenge      int `json:"ARCADE__creeper_attack_challenge,omitempty"`
				WOOLGAMESFlawlessChallenge        int `json:"WOOL_GAMES__flawless_challenge,omitempty"`
				WOOLGAMESMercilessKillerChallenge int `json:"WOOL_GAMES__merciless_killer_challenge,omitempty"`
			} `json:"all_time,omitempty"`
			DayA struct {
				TNTGAMESTntRunChallenge      int `json:"TNTGAMES__tnt_run_challenge,omitempty"`
				TNTGAMESBowSpleefChallenge   int `json:"TNTGAMES__bow_spleef_challenge,omitempty"`
				PAINTBALLKillStreakChallenge int `json:"PAINTBALL__kill_streak_challenge,omitempty"`
				PAINTBALLFinishChallenge     int `json:"PAINTBALL__finish_challenge,omitempty"`
			} `json:"day_a,omitempty"`
			DayD struct {
				PAINTBALLKillingSpreeChallenge int `json:"PAINTBALL__killing_spree_challenge,omitempty"`
				PAINTBALLKillStreakChallenge   int `json:"PAINTBALL__kill_streak_challenge,omitempty"`
				QUAKECRAFTComboChallenge       int `json:"QUAKECRAFT__combo_challenge,omitempty"`
			} `json:"day_d,omitempty"`
		} `json:"challenges,omitempty"`
		ParkourCheckpointBests struct {
			Duels struct {
				Num0 int `json:"0,omitempty"`
			} `json:"Duels,omitempty"`
			Legacy struct {
				Num0 int `json:"0,omitempty"`
				Num1 int `json:"1,omitempty"`
				Num2 int `json:"2,omitempty"`
			} `json:"Legacy,omitempty"`
			Prototype struct {
				Num0 int `json:"0,omitempty"`
				Num1 int `json:"1,omitempty"`
			} `json:"Prototype,omitempty"`
			Tourney struct {
				Num0 int `json:"0,omitempty"`
				Num1 int `json:"1,omitempty"`
				Num2 int `json:"2,omitempty"`
				Num3 int `json:"3,omitempty"`
			} `json:"Tourney,omitempty"`
			CopsnCrims struct {
				Num0 int `json:"0,omitempty"`
				Num1 int `json:"1,omitempty"`
			} `json:"CopsnCrims,omitempty"`
		} `json:"parkourCheckpointBests,omitempty"`
		AchievementPoints     int `json:"achievementPoints,omitempty"`
		AchievementRewardsNew struct {
			ForPoints2000 int64 `json:"for_points_2000,omitempty"`
			ForPoints3000 int64 `json:"for_points_3000,omitempty"`
		} `json:"achievementRewardsNew,omitempty"`
		AchievementTotem struct {
			CanCustomize     bool     `json:"canCustomize,omitempty"`
			AllowedMaxHeight int      `json:"allowed_max_height,omitempty"`
			UnlockedParts    []string `json:"unlockedParts,omitempty"`
			SelectedParts    struct {
				Slot0 string `json:"slot_0,omitempty"`
			} `json:"selectedParts,omitempty"`
			UnlockedColors []string `json:"unlockedColors,omitempty"`
			SelectedColors struct {
			} `json:"selectedColors,omitempty"`
		} `json:"achievementTotem,omitempty"`
		Tourney struct {
			FirstJoinLobby int64 `json:"first_join_lobby,omitempty"`
			Bedwars4S0     struct {
				Playtime       int   `json:"playtime,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
			} `json:"bedwars4s_0,omitempty"`
			TotalTributes int `json:"total_tributes,omitempty"`
			McgoDefusal0  struct {
				SeenRPbook     bool  `json:"seenRPbook,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
			} `json:"mcgo_defusal_0,omitempty"`
			BedwarsTwoFour0 struct {
				GamesPlayed    int   `json:"games_played,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
			} `json:"bedwars_two_four_0,omitempty"`
			QuakeSolo20 struct {
				GamesPlayed    int   `json:"games_played,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
				FirstGame      int64 `json:"first_game,omitempty"`
			} `json:"quake_solo2_0,omitempty"`
			SwInsaneDoubles0 struct {
				GamesPlayed    int `json:"games_played,omitempty"`
				Playtime       int `json:"playtime,omitempty"`
				TributesEarned int `json:"tributes_earned,omitempty"`
			} `json:"sw_insane_doubles_0,omitempty"`
			GrinchSimulator0 struct {
				GamesPlayed    int   `json:"games_played,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
			} `json:"grinch_simulator_0,omitempty"`
			McgoDefusal1 struct {
				SeenRPbook     bool  `json:"seenRPbook,omitempty"`
				GamesPlayed    int   `json:"games_played,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
			} `json:"mcgo_defusal_1,omitempty"`
			BedwarsEightTwo0 struct {
				GamesPlayed    int `json:"games_played,omitempty"`
				Playtime       int `json:"playtime,omitempty"`
				TributesEarned int `json:"tributes_earned,omitempty"`
			} `json:"bedwars_eight_two_0,omitempty"`
			WoolWars1 struct {
				GamesPlayed    int   `json:"games_played,omitempty"`
				Playtime       int   `json:"playtime,omitempty"`
				FirstWin       int64 `json:"first_win,omitempty"`
				TributesEarned int   `json:"tributes_earned,omitempty"`
			} `json:"wool_wars_1,omitempty"`
		} `json:"tourney,omitempty"`
		CurrentCloak       string `json:"currentCloak,omitempty"`
		ParticlePack       string `json:"particlePack,omitempty"`
		CurrentClickEffect string `json:"currentClickEffect,omitempty"`
		Monthlycrates      struct {
			One2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"1-2017,omitempty"`
			One2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"1-2018,omitempty"`
			One2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"1-2019,omitempty"`
			One02016 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"10-2016,omitempty"`
			One02017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"10-2017,omitempty"`
			One02018 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
			} `json:"10-2018,omitempty"`
			One12016 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"11-2016,omitempty"`
			One12017 struct {
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"11-2017,omitempty"`
			One12018 struct {
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"11-2018,omitempty"`
			One22017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"12-2017,omitempty"`
			One22018 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"12-2018,omitempty"`
			Two2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"2-2017,omitempty"`
			Two2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"2-2018,omitempty"`
			Two2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"2-2019,omitempty"`
			Three2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"3-2017,omitempty"`
			Four2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"4-2017,omitempty"`
			Four2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"4-2019,omitempty"`
			Five2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"5-2017,omitempty"`
			Six2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"6-2017,omitempty"`
			Six2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"6-2018,omitempty"`
			Seven2017 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"7-2017,omitempty"`
			Seven2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"7-2018,omitempty"`
			Seven2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"7-2019,omitempty"`
			Eight2016 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"8-2016,omitempty"`
			Eight2017 struct {
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"8-2017,omitempty"`
			Eight2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"8-2018,omitempty"`
			Eight2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"8-2019,omitempty"`
			Nine2016 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"9-2016,omitempty"`
			Nine2017 struct {
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"9-2017,omitempty"`
			Nine2018 struct {
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"9-2018,omitempty"`
			Nine2019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"9-2019,omitempty"`
			One12019 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"11-2019,omitempty"`
			One2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"1-2020,omitempty"`
			Two2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"2-2020,omitempty"`
			Three2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"3-2020,omitempty"`
			Four2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"4-2020,omitempty"`
			Five2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"5-2020,omitempty"`
			Six2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"6-2020,omitempty"`
			Seven2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"7-2020,omitempty"`
			Eight2020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"8-2020,omitempty"`
			Nine2020 struct {
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"9-2020,omitempty"`
			One12020 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"11-2020,omitempty"`
			One22020 struct {
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"12-2020,omitempty"`
			One2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"1-2021,omitempty"`
			Two2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"2-2021,omitempty"`
			Three2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"3-2021,omitempty"`
			Four2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
			} `json:"4-2021,omitempty"`
			Five2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"5-2021,omitempty"`
			Seven2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"7-2021,omitempty"`
			Eight2021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"8-2021,omitempty"`
			One02021 struct {
				Regular bool `json:"REGULAR,omitempty"`
			} `json:"10-2021,omitempty"`
			One22021 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"12-2021,omitempty"`
			Two2022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"2-2022,omitempty"`
			Three2022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"3-2022,omitempty"`
			Four2022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"4-2022,omitempty"`
			Five2022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"5-2022,omitempty"`
			Seven2022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"7-2022,omitempty"`
			One02022 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"10-2022,omitempty"`
			One2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
			} `json:"1-2023,omitempty"`
			Two2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
			} `json:"2-2023,omitempty"`
			Three2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"3-2023,omitempty"`
			Four2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
			} `json:"4-2023,omitempty"`
			Five2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"5-2023,omitempty"`
			Six2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
			} `json:"6-2023,omitempty"`
			Seven2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"7-2023,omitempty"`
			Eight2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"8-2023,omitempty"`
			Nine2023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"9-2023,omitempty"`
			One02023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"10-2023,omitempty"`
			One12023 struct {
				Regular bool `json:"REGULAR,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
			} `json:"11-2023,omitempty"`
			One22023 struct {
				Normal  bool `json:"NORMAL,omitempty"`
				Vip     bool `json:"VIP,omitempty"`
				VipPlus bool `json:"VIP_PLUS,omitempty"`
				Mvp     bool `json:"MVP,omitempty"`
				MvpPlus bool `json:"MVP_PLUS,omitempty"`
			} `json:"12-2023,omitempty"`
		} `json:"monthlycrates,omitempty"`
		Main2017Tutorial       bool   `json:"main2017Tutorial,omitempty"`
		CurrentGadget          string `json:"currentGadget,omitempty"`
		SnowballFightIntro2019 bool   `json:"snowball_fight_intro_2019,omitempty"`
		GiftsGrinch            int    `json:"gifts_grinch,omitempty"`
		ClaimedPotatoTalisman  int64  `json:"claimed_potato_talisman,omitempty"`
		Outfit                 struct {
			Chestplate string `json:"CHESTPLATE,omitempty"`
			Leggings   string `json:"LEGGINGS,omitempty"`
			Boots      string `json:"BOOTS,omitempty"`
			Helmet     string `json:"HELMET,omitempty"`
		} `json:"outfit,omitempty"`
		SkyblockFreeCookie int64  `json:"skyblock_free_cookie,omitempty"`
		ClaimedCenturyCake int64  `json:"claimed_century_cake,omitempty"`
		LevelUpMVPPLUS     int64  `json:"levelUp_MVP_PLUS,omitempty"`
		RankPlusColor      string `json:"rankPlusColor,omitempty"`
		QuestSettings      struct {
			AutoActivate bool `json:"autoActivate,omitempty"`
		} `json:"questSettings,omitempty"`
		ScorpiusBribe120           int64 `json:"scorpius_bribe_120,omitempty"`
		AnniversaryNPCVisited2021  []int `json:"anniversaryNPCVisited2021,omitempty"`
		AnniversaryNPCProgress2021 int   `json:"anniversaryNPCProgress2021,omitempty"`
		Easter2021Cooldowns2       struct {
			VipPlus3 bool `json:"VIP_PLUS3,omitempty"`
			VipPlus2 bool `json:"VIP_PLUS2,omitempty"`
			VipPlus1 bool `json:"VIP_PLUS1,omitempty"`
			VipPlus0 bool `json:"VIP_PLUS0,omitempty"`
			Normal0  bool `json:"NORMAL0,omitempty"`
			Normal1  bool `json:"NORMAL1,omitempty"`
			Normal2  bool `json:"NORMAL2,omitempty"`
			Normal3  bool `json:"NORMAL3,omitempty"`
			MvpPlus0 bool `json:"MVP_PLUS0,omitempty"`
			MvpPlus1 bool `json:"MVP_PLUS1,omitempty"`
			MvpPlus2 bool `json:"MVP_PLUS2,omitempty"`
			MvpPlus3 bool `json:"MVP_PLUS3,omitempty"`
			Mvp0     bool `json:"MVP0,omitempty"`
			Mvp1     bool `json:"MVP1,omitempty"`
			Mvp2     bool `json:"MVP2,omitempty"`
			Mvp3     bool `json:"MVP3,omitempty"`
			Vip0     bool `json:"VIP0,omitempty"`
			Vip1     bool `json:"VIP1,omitempty"`
			Vip2     bool `json:"VIP2,omitempty"`
			Vip3     bool `json:"VIP3,omitempty"`
		} `json:"easter2021Cooldowns2,omitempty"`
		AdventRewards2017 struct {
			Day2  int64 `json:"day2,omitempty"`
			Day3  int64 `json:"day3,omitempty"`
			Day4  int64 `json:"day4,omitempty"`
			Day5  int64 `json:"day5,omitempty"`
			Day6  int64 `json:"day6,omitempty"`
			Day7  int64 `json:"day7,omitempty"`
			Day8  int64 `json:"day8,omitempty"`
			Day9  int64 `json:"day9,omitempty"`
			Day10 int64 `json:"day10,omitempty"`
			Day11 int64 `json:"day11,omitempty"`
			Day12 int64 `json:"day12,omitempty"`
			Day13 int64 `json:"day13,omitempty"`
			Day14 int64 `json:"day14,omitempty"`
			Day15 int64 `json:"day15,omitempty"`
			Day16 int64 `json:"day16,omitempty"`
			Day17 int64 `json:"day17,omitempty"`
			Day18 int64 `json:"day18,omitempty"`
			Day19 int64 `json:"day19,omitempty"`
			Day20 int64 `json:"day20,omitempty"`
			Day21 int64 `json:"day21,omitempty"`
			Day23 int64 `json:"day23,omitempty"`
			Day24 int64 `json:"day24,omitempty"`
			Day25 int64 `json:"day25,omitempty"`
			Day22 int64 `json:"day22,omitempty"`
			Day1  int64 `json:"day1,omitempty"`
		} `json:"adventRewards2017,omitempty"`
		AdventRewards2019 struct {
			Day1  int64 `json:"day1,omitempty"`
			Day2  int64 `json:"day2,omitempty"`
			Day3  int64 `json:"day3,omitempty"`
			Day4  int64 `json:"day4,omitempty"`
			Day5  int64 `json:"day5,omitempty"`
			Day6  int64 `json:"day6,omitempty"`
			Day7  int64 `json:"day7,omitempty"`
			Day8  int64 `json:"day8,omitempty"`
			Day9  int64 `json:"day9,omitempty"`
			Day10 int64 `json:"day10,omitempty"`
			Day11 int64 `json:"day11,omitempty"`
			Day12 int64 `json:"day12,omitempty"`
			Day13 int64 `json:"day13,omitempty"`
			Day14 int64 `json:"day14,omitempty"`
			Day15 int64 `json:"day15,omitempty"`
			Day16 int64 `json:"day16,omitempty"`
			Day17 int64 `json:"day17,omitempty"`
			Day18 int64 `json:"day18,omitempty"`
			Day19 int64 `json:"day19,omitempty"`
			Day20 int64 `json:"day20,omitempty"`
			Day21 int64 `json:"day21,omitempty"`
			Day22 int64 `json:"day22,omitempty"`
			Day23 int64 `json:"day23,omitempty"`
			Day24 int64 `json:"day24,omitempty"`
			Day25 int64 `json:"day25,omitempty"`
		} `json:"adventRewards2019,omitempty"`
		AdventRewards2020 struct {
			Day1  int64 `json:"day1,omitempty"`
			Day2  int64 `json:"day2,omitempty"`
			Day5  int64 `json:"day5,omitempty"`
			Day6  int64 `json:"day6,omitempty"`
			Day7  int64 `json:"day7,omitempty"`
			Day8  int64 `json:"day8,omitempty"`
			Day9  int64 `json:"day9,omitempty"`
			Day12 int64 `json:"day12,omitempty"`
			Day13 int64 `json:"day13,omitempty"`
			Day15 int64 `json:"day15,omitempty"`
			Day16 int64 `json:"day16,omitempty"`
			Day17 int64 `json:"day17,omitempty"`
			Day19 int64 `json:"day19,omitempty"`
			Day20 int64 `json:"day20,omitempty"`
			Day21 int64 `json:"day21,omitempty"`
			Day22 int64 `json:"day22,omitempty"`
			Day23 int64 `json:"day23,omitempty"`
			Day24 int64 `json:"day24,omitempty"`
			Day25 int64 `json:"day25,omitempty"`
		} `json:"adventRewards2020,omitempty"`
		AdventRewardsV22018 struct {
			Day1  int64 `json:"day1,omitempty"`
			Day2  int64 `json:"day2,omitempty"`
			Day3  int64 `json:"day3,omitempty"`
			Day4  int64 `json:"day4,omitempty"`
			Day5  int64 `json:"day5,omitempty"`
			Day6  int64 `json:"day6,omitempty"`
			Day9  int64 `json:"day9,omitempty"`
			Day10 int64 `json:"day10,omitempty"`
			Day13 int64 `json:"day13,omitempty"`
			Day15 int64 `json:"day15,omitempty"`
			Day16 int64 `json:"day16,omitempty"`
			Day17 int64 `json:"day17,omitempty"`
			Day18 int64 `json:"day18,omitempty"`
			Day19 int64 `json:"day19,omitempty"`
			Day21 int64 `json:"day21,omitempty"`
			Day22 int64 `json:"day22,omitempty"`
			Day23 int64 `json:"day23,omitempty"`
			Day24 int64 `json:"day24,omitempty"`
			Day25 int64 `json:"day25,omitempty"`
		} `json:"adventRewards_v2_2018,omitempty"`
		ClaimedYear143Cake  int64  `json:"claimed_year143_cake,omitempty"`
		VanityFavorites     string `json:"vanityFavorites,omitempty"`
		Summer2021Cooldowns struct {
			MvpPlus3 bool `json:"MVP_PLUS3,omitempty"`
			MvpPlus2 bool `json:"MVP_PLUS2,omitempty"`
			MvpPlus1 bool `json:"MVP_PLUS1,omitempty"`
			MvpPlus0 bool `json:"MVP_PLUS0,omitempty"`
			Normal3  bool `json:"NORMAL3,omitempty"`
			Normal2  bool `json:"NORMAL2,omitempty"`
			Normal1  bool `json:"NORMAL1,omitempty"`
			Normal0  bool `json:"NORMAL0,omitempty"`
			VipPlus0 bool `json:"VIP_PLUS0,omitempty"`
			VipPlus1 bool `json:"VIP_PLUS1,omitempty"`
			VipPlus2 bool `json:"VIP_PLUS2,omitempty"`
			VipPlus3 bool `json:"VIP_PLUS3,omitempty"`
			Mvp0     bool `json:"MVP0,omitempty"`
			Mvp1     bool `json:"MVP1,omitempty"`
			Mvp2     bool `json:"MVP2,omitempty"`
			Mvp3     bool `json:"MVP3,omitempty"`
			Vip0     bool `json:"VIP0,omitempty"`
			Vip1     bool `json:"VIP1,omitempty"`
			Vip2     bool `json:"VIP2,omitempty"`
			Vip3     bool `json:"VIP3,omitempty"`
		} `json:"summer2021Cooldowns,omitempty"`
		Halloween2021Cooldowns struct {
			MvpPlus3 bool `json:"MVP_PLUS3,omitempty"`
			MvpPlus2 bool `json:"MVP_PLUS2,omitempty"`
			MvpPlus1 bool `json:"MVP_PLUS1,omitempty"`
			MvpPlus0 bool `json:"MVP_PLUS0,omitempty"`
			Normal3  bool `json:"NORMAL3,omitempty"`
			Normal2  bool `json:"NORMAL2,omitempty"`
			Normal1  bool `json:"NORMAL1,omitempty"`
			Normal0  bool `json:"NORMAL0,omitempty"`
			VipPlus0 bool `json:"VIP_PLUS0,omitempty"`
			VipPlus1 bool `json:"VIP_PLUS1,omitempty"`
			VipPlus2 bool `json:"VIP_PLUS2,omitempty"`
			VipPlus3 bool `json:"VIP_PLUS3,omitempty"`
			Mvp0     bool `json:"MVP0,omitempty"`
			Mvp1     bool `json:"MVP1,omitempty"`
			Mvp2     bool `json:"MVP2,omitempty"`
			Mvp3     bool `json:"MVP3,omitempty"`
			Vip0     bool `json:"VIP0,omitempty"`
			Vip1     bool `json:"VIP1,omitempty"`
			Vip2     bool `json:"VIP2,omitempty"`
			Vip3     bool `json:"VIP3,omitempty"`
		} `json:"halloween2021Cooldowns,omitempty"`
		Seasonal struct {
			Christmas struct {
				Num2021 struct {
					AdventRewards struct {
						Day1  int64 `json:"day1,omitempty"`
						Day2  int64 `json:"day2,omitempty"`
						Day4  int64 `json:"day4,omitempty"`
						Day5  int64 `json:"day5,omitempty"`
						Day6  int64 `json:"day6,omitempty"`
						Day7  int64 `json:"day7,omitempty"`
						Day8  int64 `json:"day8,omitempty"`
						Day9  int64 `json:"day9,omitempty"`
						Day11 int64 `json:"day11,omitempty"`
						Day12 int64 `json:"day12,omitempty"`
						Day13 int64 `json:"day13,omitempty"`
						Day15 int64 `json:"day15,omitempty"`
						Day16 int64 `json:"day16,omitempty"`
						Day17 int64 `json:"day17,omitempty"`
						Day18 int64 `json:"day18,omitempty"`
						Day19 int64 `json:"day19,omitempty"`
						Day20 int64 `json:"day20,omitempty"`
						Day22 int64 `json:"day22,omitempty"`
						Day23 int64 `json:"day23,omitempty"`
						Day24 int64 `json:"day24,omitempty"`
						Day25 int64 `json:"day25,omitempty"`
					} `json:"adventRewards,omitempty"`
					Presents struct {
						Classic1      bool `json:"CLASSIC_1,omitempty"`
						Classic2      bool `json:"CLASSIC_2,omitempty"`
						Arcade2       bool `json:"ARCADE_2,omitempty"`
						Arcade3       bool `json:"ARCADE_3,omitempty"`
						Arcade1       bool `json:"ARCADE_1,omitempty"`
						Prototype1    bool `json:"PROTOTYPE_1,omitempty"`
						MainLobby1    bool `json:"MAIN_LOBBY_1,omitempty"`
						MainLobby2    bool `json:"MAIN_LOBBY_2,omitempty"`
						MainLobby12   bool `json:"MAIN_LOBBY_12,omitempty"`
						MainLobby6    bool `json:"MAIN_LOBBY_6,omitempty"`
						MainLobby7    bool `json:"MAIN_LOBBY_7,omitempty"`
						MainLobby8    bool `json:"MAIN_LOBBY_8,omitempty"`
						MainLobby37   bool `json:"MAIN_LOBBY_37,omitempty"`
						MainLobby38   bool `json:"MAIN_LOBBY_38,omitempty"`
						MainLobby11   bool `json:"MAIN_LOBBY_11,omitempty"`
						MainLobby10   bool `json:"MAIN_LOBBY_10,omitempty"`
						MainLobby36   bool `json:"MAIN_LOBBY_36,omitempty"`
						MainLobby22   bool `json:"MAIN_LOBBY_22,omitempty"`
						MainLobby25   bool `json:"MAIN_LOBBY_25,omitempty"`
						MainLobby4    bool `json:"MAIN_LOBBY_4,omitempty"`
						MainLobby35   bool `json:"MAIN_LOBBY_35,omitempty"`
						MainLobby3    bool `json:"MAIN_LOBBY_3,omitempty"`
						MainLobby24   bool `json:"MAIN_LOBBY_24,omitempty"`
						MainLobby23   bool `json:"MAIN_LOBBY_23,omitempty"`
						MainLobby27   bool `json:"MAIN_LOBBY_27,omitempty"`
						MainLobby5    bool `json:"MAIN_LOBBY_5,omitempty"`
						MainLobby33   bool `json:"MAIN_LOBBY_33,omitempty"`
						MainLobby34   bool `json:"MAIN_LOBBY_34,omitempty"`
						MainLobby13   bool `json:"MAIN_LOBBY_13,omitempty"`
						MainLobby14   bool `json:"MAIN_LOBBY_14,omitempty"`
						MainLobby15   bool `json:"MAIN_LOBBY_15,omitempty"`
						MainLobby9    bool `json:"MAIN_LOBBY_9,omitempty"`
						MainLobby28   bool `json:"MAIN_LOBBY_28,omitempty"`
						MainLobby29   bool `json:"MAIN_LOBBY_29,omitempty"`
						MainLobby21   bool `json:"MAIN_LOBBY_21,omitempty"`
						MainLobby32   bool `json:"MAIN_LOBBY_32,omitempty"`
						MainLobby17   bool `json:"MAIN_LOBBY_17,omitempty"`
						MainLobby16   bool `json:"MAIN_LOBBY_16,omitempty"`
						MainLobby31   bool `json:"MAIN_LOBBY_31,omitempty"`
						MainLobby19   bool `json:"MAIN_LOBBY_19,omitempty"`
						MainLobby18   bool `json:"MAIN_LOBBY_18,omitempty"`
						MainLobby30   bool `json:"MAIN_LOBBY_30,omitempty"`
						MainLobby20   bool `json:"MAIN_LOBBY_20,omitempty"`
						MainLobby26   bool `json:"MAIN_LOBBY_26,omitempty"`
						MainLobby40   bool `json:"MAIN_LOBBY_40,omitempty"`
						MainLobby39   bool `json:"MAIN_LOBBY_39,omitempty"`
						Bedwars1      bool `json:"BEDWARS_1,omitempty"`
						Bedwars2      bool `json:"BEDWARS_2,omitempty"`
						Bedwars3      bool `json:"BEDWARS_3,omitempty"`
						Bedwars4      bool `json:"BEDWARS_4,omitempty"`
						Bedwars5      bool `json:"BEDWARS_5,omitempty"`
						Skywars1      bool `json:"SKYWARS_1,omitempty"`
						Skywars2      bool `json:"SKYWARS_2,omitempty"`
						Skywars5      bool `json:"SKYWARS_5,omitempty"`
						Skywars4      bool `json:"SKYWARS_4,omitempty"`
						Skywars3      bool `json:"SKYWARS_3,omitempty"`
						Murder1       bool `json:"MURDER_1,omitempty"`
						Murder3       bool `json:"MURDER_3,omitempty"`
						Murder2       bool `json:"MURDER_2,omitempty"`
						Housing1      bool `json:"HOUSING_1,omitempty"`
						Housing5      bool `json:"HOUSING_5,omitempty"`
						Housing3      bool `json:"HOUSING_3,omitempty"`
						Housing2      bool `json:"HOUSING_2,omitempty"`
						Housing4      bool `json:"HOUSING_4,omitempty"`
						Buildbattle1  bool `json:"BUILDBATTLE_1,omitempty"`
						Buildbattle2  bool `json:"BUILDBATTLE_2,omitempty"`
						Buildbattle3  bool `json:"BUILDBATTLE_3,omitempty"`
						Duels1        bool `json:"DUELS_1,omitempty"`
						Duels3        bool `json:"DUELS_3,omitempty"`
						Duels2        bool `json:"DUELS_2,omitempty"`
						Prototype3    bool `json:"PROTOTYPE_3,omitempty"`
						Prototype2    bool `json:"PROTOTYPE_2,omitempty"`
						Uhc1          bool `json:"UHC_1,omitempty"`
						Uhc2          bool `json:"UHC_2,omitempty"`
						Uhc3          bool `json:"UHC_3,omitempty"`
						Tnt1          bool `json:"TNT_1,omitempty"`
						Tnt2          bool `json:"TNT_2,omitempty"`
						Tnt3          bool `json:"TNT_3,omitempty"`
						Classic3      bool `json:"CLASSIC_3,omitempty"`
						CopsAndCrims3 bool `json:"COPS_AND_CRIMS_3,omitempty"`
						CopsAndCrims2 bool `json:"COPS_AND_CRIMS_2,omitempty"`
						CopsAndCrims1 bool `json:"COPS_AND_CRIMS_1,omitempty"`
						Blitz1        bool `json:"BLITZ_1,omitempty"`
						Blitz2        bool `json:"BLITZ_2,omitempty"`
						Blitz3        bool `json:"BLITZ_3,omitempty"`
						Megawalls1    bool `json:"MEGAWALLS_1,omitempty"`
						Megawalls2    bool `json:"MEGAWALLS_2,omitempty"`
						Megawalls3    bool `json:"MEGAWALLS_3,omitempty"`
						Smash1        bool `json:"SMASH_1,omitempty"`
						Smash2        bool `json:"SMASH_2,omitempty"`
						Smash3        bool `json:"SMASH_3,omitempty"`
						Warlords2     bool `json:"WARLORDS_2,omitempty"`
						Warlords3     bool `json:"WARLORDS_3,omitempty"`
						Warlords1     bool `json:"WARLORDS_1,omitempty"`
					} `json:"presents,omitempty"`
				} `json:"2021,omitempty"`
				Num2022 struct {
					AdventRewards struct {
						Day1  int64 `json:"day1,omitempty"`
						Day2  int64 `json:"day2,omitempty"`
						Day6  int64 `json:"day6,omitempty"`
						Day9  int64 `json:"day9,omitempty"`
						Day10 int64 `json:"day10,omitempty"`
						Day11 int64 `json:"day11,omitempty"`
						Day13 int64 `json:"day13,omitempty"`
						Day14 int64 `json:"day14,omitempty"`
						Day16 int64 `json:"day16,omitempty"`
						Day17 int64 `json:"day17,omitempty"`
						Day18 int64 `json:"day18,omitempty"`
						Day20 int64 `json:"day20,omitempty"`
						Day21 int64 `json:"day21,omitempty"`
						Day22 int64 `json:"day22,omitempty"`
						Day23 int64 `json:"day23,omitempty"`
						Day24 int64 `json:"day24,omitempty"`
					} `json:"adventRewards,omitempty"`
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
					Presents struct {
						Classic1      bool `json:"CLASSIC_1,omitempty"`
						Prototype1    bool `json:"PROTOTYPE_1,omitempty"`
						Prototype3    bool `json:"PROTOTYPE_3,omitempty"`
						Arcade1       bool `json:"ARCADE_1,omitempty"`
						MainLobby1    bool `json:"MAIN_LOBBY_1,omitempty"`
						MainLobby2    bool `json:"MAIN_LOBBY_2,omitempty"`
						MainLobby31   bool `json:"MAIN_LOBBY_31,omitempty"`
						MainLobby3    bool `json:"MAIN_LOBBY_3,omitempty"`
						MainLobby40   bool `json:"MAIN_LOBBY_40,omitempty"`
						MainLobby18   bool `json:"MAIN_LOBBY_18,omitempty"`
						MainLobby34   bool `json:"MAIN_LOBBY_34,omitempty"`
						MainLobby36   bool `json:"MAIN_LOBBY_36,omitempty"`
						MainLobby11   bool `json:"MAIN_LOBBY_11,omitempty"`
						MainLobby12   bool `json:"MAIN_LOBBY_12,omitempty"`
						MainLobby25   bool `json:"MAIN_LOBBY_25,omitempty"`
						MainLobby5    bool `json:"MAIN_LOBBY_5,omitempty"`
						MainLobby27   bool `json:"MAIN_LOBBY_27,omitempty"`
						MainLobby21   bool `json:"MAIN_LOBBY_21,omitempty"`
						MainLobby19   bool `json:"MAIN_LOBBY_19,omitempty"`
						MainLobby20   bool `json:"MAIN_LOBBY_20,omitempty"`
						MainLobby22   bool `json:"MAIN_LOBBY_22,omitempty"`
						MainLobby13   bool `json:"MAIN_LOBBY_13,omitempty"`
						MainLobby38   bool `json:"MAIN_LOBBY_38,omitempty"`
						MainLobby30   bool `json:"MAIN_LOBBY_30,omitempty"`
						MainLobby28   bool `json:"MAIN_LOBBY_28,omitempty"`
						MainLobby29   bool `json:"MAIN_LOBBY_29,omitempty"`
						MainLobby37   bool `json:"MAIN_LOBBY_37,omitempty"`
						MainLobby17   bool `json:"MAIN_LOBBY_17,omitempty"`
						MainLobby33   bool `json:"MAIN_LOBBY_33,omitempty"`
						MainLobby24   bool `json:"MAIN_LOBBY_24,omitempty"`
						MainLobby26   bool `json:"MAIN_LOBBY_26,omitempty"`
						MainLobby7    bool `json:"MAIN_LOBBY_7,omitempty"`
						MainLobby32   bool `json:"MAIN_LOBBY_32,omitempty"`
						MainLobby39   bool `json:"MAIN_LOBBY_39,omitempty"`
						MainLobby6    bool `json:"MAIN_LOBBY_6,omitempty"`
						MainLobby4    bool `json:"MAIN_LOBBY_4,omitempty"`
						MainLobby35   bool `json:"MAIN_LOBBY_35,omitempty"`
						MainLobby9    bool `json:"MAIN_LOBBY_9,omitempty"`
						MainLobby8    bool `json:"MAIN_LOBBY_8,omitempty"`
						MainLobby15   bool `json:"MAIN_LOBBY_15,omitempty"`
						MainLobby16   bool `json:"MAIN_LOBBY_16,omitempty"`
						MainLobby10   bool `json:"MAIN_LOBBY_10,omitempty"`
						MainLobby14   bool `json:"MAIN_LOBBY_14,omitempty"`
						MainLobby23   bool `json:"MAIN_LOBBY_23,omitempty"`
						Housing1      bool `json:"HOUSING_1,omitempty"`
						Bedwars1      bool `json:"BEDWARS_1,omitempty"`
						Bedwars2      bool `json:"BEDWARS_2,omitempty"`
						Bedwars4      bool `json:"BEDWARS_4,omitempty"`
						Bedwars5      bool `json:"BEDWARS_5,omitempty"`
						Bedwars3      bool `json:"BEDWARS_3,omitempty"`
						Skywars1      bool `json:"SKYWARS_1,omitempty"`
						Skywars4      bool `json:"SKYWARS_4,omitempty"`
						Skywars2      bool `json:"SKYWARS_2,omitempty"`
						Skywars5      bool `json:"SKYWARS_5,omitempty"`
						Skywars3      bool `json:"SKYWARS_3,omitempty"`
						Tnt1          bool `json:"TNT_1,omitempty"`
						Tnt2          bool `json:"TNT_2,omitempty"`
						Tnt3          bool `json:"TNT_3,omitempty"`
						Duels1        bool `json:"DUELS_1,omitempty"`
						Duels3        bool `json:"DUELS_3,omitempty"`
						Duels2        bool `json:"DUELS_2,omitempty"`
						Uhc1          bool `json:"UHC_1,omitempty"`
						Uhc2          bool `json:"UHC_2,omitempty"`
						Uhc3          bool `json:"UHC_3,omitempty"`
						Megawalls1    bool `json:"MEGAWALLS_1,omitempty"`
						Megawalls2    bool `json:"MEGAWALLS_2,omitempty"`
						Megawalls3    bool `json:"MEGAWALLS_3,omitempty"`
						Buildbattle1  bool `json:"BUILDBATTLE_1,omitempty"`
						Buildbattle2  bool `json:"BUILDBATTLE_2,omitempty"`
						Buildbattle3  bool `json:"BUILDBATTLE_3,omitempty"`
						Arcade2       bool `json:"ARCADE_2,omitempty"`
						Arcade3       bool `json:"ARCADE_3,omitempty"`
						Prototype2    bool `json:"PROTOTYPE_2,omitempty"`
						Blitz1        bool `json:"BLITZ_1,omitempty"`
						Blitz2        bool `json:"BLITZ_2,omitempty"`
						Blitz3        bool `json:"BLITZ_3,omitempty"`
						Murder1       bool `json:"MURDER_1,omitempty"`
						Murder2       bool `json:"MURDER_2,omitempty"`
						Murder3       bool `json:"MURDER_3,omitempty"`
						Classic2      bool `json:"CLASSIC_2,omitempty"`
						Classic3      bool `json:"CLASSIC_3,omitempty"`
						Wool1         bool `json:"WOOL_1,omitempty"`
						Wool2         bool `json:"WOOL_2,omitempty"`
						Wool3         bool `json:"WOOL_3,omitempty"`
						Housing3      bool `json:"HOUSING_3,omitempty"`
						Housing2      bool `json:"HOUSING_2,omitempty"`
						CopsAndCrims1 bool `json:"COPS_AND_CRIMS_1,omitempty"`
						CopsAndCrims2 bool `json:"COPS_AND_CRIMS_2,omitempty"`
						CopsAndCrims3 bool `json:"COPS_AND_CRIMS_3,omitempty"`
						Smash1        bool `json:"SMASH_1,omitempty"`
						Smash2        bool `json:"SMASH_2,omitempty"`
						Smash3        bool `json:"SMASH_3,omitempty"`
						Warlords1     bool `json:"WARLORDS_1,omitempty"`
						Warlords3     bool `json:"WARLORDS_3,omitempty"`
						Warlords2     bool `json:"WARLORDS_2,omitempty"`
					} `json:"presents,omitempty"`
				} `json:"2022,omitempty"`
				Num2023 struct {
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
					Presents struct {
						Prototype1    bool `json:"PROTOTYPE_1,omitempty"`
						Classic1      bool `json:"CLASSIC_1,omitempty"`
						MainLobby1    bool `json:"MAIN_LOBBY_1,omitempty"`
						MainLobby2    bool `json:"MAIN_LOBBY_2,omitempty"`
						MainLobby28   bool `json:"MAIN_LOBBY_28,omitempty"`
						MainLobby3    bool `json:"MAIN_LOBBY_3,omitempty"`
						MainLobby13   bool `json:"MAIN_LOBBY_13,omitempty"`
						MainLobby27   bool `json:"MAIN_LOBBY_27,omitempty"`
						MainLobby15   bool `json:"MAIN_LOBBY_15,omitempty"`
						Tnt1          bool `json:"TNT_1,omitempty"`
						Tnt2          bool `json:"TNT_2,omitempty"`
						Tnt3          bool `json:"TNT_3,omitempty"`
						Arcade1       bool `json:"ARCADE_1,omitempty"`
						MainLobby4    bool `json:"MAIN_LOBBY_4,omitempty"`
						MainLobby10   bool `json:"MAIN_LOBBY_10,omitempty"`
						MainLobby35   bool `json:"MAIN_LOBBY_35,omitempty"`
						MainLobby11   bool `json:"MAIN_LOBBY_11,omitempty"`
						MainLobby12   bool `json:"MAIN_LOBBY_12,omitempty"`
						MainLobby20   bool `json:"MAIN_LOBBY_20,omitempty"`
						MainLobby5    bool `json:"MAIN_LOBBY_5,omitempty"`
						MainLobby29   bool `json:"MAIN_LOBBY_29,omitempty"`
						MainLobby9    bool `json:"MAIN_LOBBY_9,omitempty"`
						MainLobby31   bool `json:"MAIN_LOBBY_31,omitempty"`
						MainLobby30   bool `json:"MAIN_LOBBY_30,omitempty"`
						MainLobby8    bool `json:"MAIN_LOBBY_8,omitempty"`
						MainLobby38   bool `json:"MAIN_LOBBY_38,omitempty"`
						MainLobby34   bool `json:"MAIN_LOBBY_34,omitempty"`
						MainLobby6    bool `json:"MAIN_LOBBY_6,omitempty"`
						MainLobby19   bool `json:"MAIN_LOBBY_19,omitempty"`
						MainLobby26   bool `json:"MAIN_LOBBY_26,omitempty"`
						MainLobby23   bool `json:"MAIN_LOBBY_23,omitempty"`
						MainLobby25   bool `json:"MAIN_LOBBY_25,omitempty"`
						MainLobby22   bool `json:"MAIN_LOBBY_22,omitempty"`
						MainLobby21   bool `json:"MAIN_LOBBY_21,omitempty"`
						MainLobby24   bool `json:"MAIN_LOBBY_24,omitempty"`
						MainLobby14   bool `json:"MAIN_LOBBY_14,omitempty"`
						MainLobby39   bool `json:"MAIN_LOBBY_39,omitempty"`
						MainLobby40   bool `json:"MAIN_LOBBY_40,omitempty"`
						MainLobby33   bool `json:"MAIN_LOBBY_33,omitempty"`
						MainLobby17   bool `json:"MAIN_LOBBY_17,omitempty"`
						MainLobby18   bool `json:"MAIN_LOBBY_18,omitempty"`
						MainLobby16   bool `json:"MAIN_LOBBY_16,omitempty"`
						MainLobby36   bool `json:"MAIN_LOBBY_36,omitempty"`
						MainLobby37   bool `json:"MAIN_LOBBY_37,omitempty"`
						MainLobby7    bool `json:"MAIN_LOBBY_7,omitempty"`
						MainLobby32   bool `json:"MAIN_LOBBY_32,omitempty"`
						Bedwars2      bool `json:"BEDWARS_2,omitempty"`
						Blitz3        bool `json:"BLITZ_3,omitempty"`
						Blitz2        bool `json:"BLITZ_2,omitempty"`
						Blitz1        bool `json:"BLITZ_1,omitempty"`
						Megawalls3    bool `json:"MEGAWALLS_3,omitempty"`
						Megawalls1    bool `json:"MEGAWALLS_1,omitempty"`
						Megawalls2    bool `json:"MEGAWALLS_2,omitempty"`
						Arcade2       bool `json:"ARCADE_2,omitempty"`
						Arcade3       bool `json:"ARCADE_3,omitempty"`
						CopsAndCrims1 bool `json:"COPS_AND_CRIMS_1,omitempty"`
						CopsAndCrims3 bool `json:"COPS_AND_CRIMS_3,omitempty"`
						CopsAndCrims2 bool `json:"COPS_AND_CRIMS_2,omitempty"`
						Uhc2          bool `json:"UHC_2,omitempty"`
						Uhc1          bool `json:"UHC_1,omitempty"`
						Uhc3          bool `json:"UHC_3,omitempty"`
						Warlords1     bool `json:"WARLORDS_1,omitempty"`
						Warlords3     bool `json:"WARLORDS_3,omitempty"`
						Warlords2     bool `json:"WARLORDS_2,omitempty"`
						Smash1        bool `json:"SMASH_1,omitempty"`
						Smash2        bool `json:"SMASH_2,omitempty"`
						Smash3        bool `json:"SMASH_3,omitempty"`
						Housing1      bool `json:"HOUSING_1,omitempty"`
						Housing3      bool `json:"HOUSING_3,omitempty"`
						Housing2      bool `json:"HOUSING_2,omitempty"`
						Skywars5      bool `json:"SKYWARS_5,omitempty"`
						Skywars4      bool `json:"SKYWARS_4,omitempty"`
						Skywars3      bool `json:"SKYWARS_3,omitempty"`
						Skywars1      bool `json:"SKYWARS_1,omitempty"`
						Skywars2      bool `json:"SKYWARS_2,omitempty"`
						Classic3      bool `json:"CLASSIC_3,omitempty"`
						Classic2      bool `json:"CLASSIC_2,omitempty"`
						Prototype3    bool `json:"PROTOTYPE_3,omitempty"`
						Prototype2    bool `json:"PROTOTYPE_2,omitempty"`
						Bedwars4      bool `json:"BEDWARS_4,omitempty"`
						Murder2       bool `json:"MURDER_2,omitempty"`
						Murder1       bool `json:"MURDER_1,omitempty"`
						Murder3       bool `json:"MURDER_3,omitempty"`
						Buildbattle1  bool `json:"BUILDBATTLE_1,omitempty"`
						Buildbattle2  bool `json:"BUILDBATTLE_2,omitempty"`
						Buildbattle3  bool `json:"BUILDBATTLE_3,omitempty"`
						Duels1        bool `json:"DUELS_1,omitempty"`
						Duels2        bool `json:"DUELS_2,omitempty"`
						Duels3        bool `json:"DUELS_3,omitempty"`
						Wool1         bool `json:"WOOL_1,omitempty"`
						Wool3         bool `json:"WOOL_3,omitempty"`
						Wool2         bool `json:"WOOL_2,omitempty"`
						Bedwars5      bool `json:"BEDWARS_5,omitempty"`
						Bedwars3      bool `json:"BEDWARS_3,omitempty"`
						Bedwars1      bool `json:"BEDWARS_1,omitempty"`
					} `json:"presents,omitempty"`
					Bingo struct {
						Easy struct {
							Objectives struct {
								Pbpowerup           int `json:"Pbpowerup,omitempty"`
								Tntrunsurviveminute int `json:"Tntrunsurviveminute,omitempty"`
								Quakedash           int `json:"Quakedash,omitempty"`
								Cvcthrowprojectile  int `json:"Cvcthrowprojectile,omitempty"`
							} `json:"objectives,omitempty"`
						} `json:"easy,omitempty"`
					} `json:"bingo,omitempty"`
					AdventRewards struct {
						Day1  int64 `json:"day1,omitempty"`
						Day2  int64 `json:"day2,omitempty"`
						Day6  int64 `json:"day6,omitempty"`
						Day10 int64 `json:"day10,omitempty"`
						Day11 int64 `json:"day11,omitempty"`
						Day12 int64 `json:"day12,omitempty"`
						Day13 int64 `json:"day13,omitempty"`
						Day16 int64 `json:"day16,omitempty"`
						Day17 int64 `json:"day17,omitempty"`
						Day18 int64 `json:"day18,omitempty"`
						Day20 int64 `json:"day20,omitempty"`
						Day22 int64 `json:"day22,omitempty"`
						Day23 int64 `json:"day23,omitempty"`
						Day24 int64 `json:"day24,omitempty"`
						Day25 int64 `json:"day25,omitempty"`
					} `json:"adventRewards,omitempty"`
				} `json:"2023,omitempty"`
			} `json:"christmas,omitempty"`
			Anniversary struct {
				Num2022 struct {
					AnniversaryNPCVisited  []int `json:"anniversaryNPCVisited,omitempty"`
					AnniversaryNPCProgress int   `json:"anniversaryNPCProgress,omitempty"`
				} `json:"2022,omitempty"`
				Num2023 struct {
					Bingo struct {
						Easy struct {
							Objectives struct {
								Pbpowerup               int `json:"Pbpowerup,omitempty"`
								Arcadehiderdamage       int `json:"Arcadehiderdamage,omitempty"`
								Tnttagplayer            int `json:"Tnttagplayer,omitempty"`
								Pixelpartysurvive       int `json:"Pixelpartysurvive,omitempty"`
								Murderbowgold           int `json:"Murderbowgold,omitempty"`
								Bedwarsdiamond          int `json:"Bedwarsdiamond,omitempty"`
								Quakedash               int `json:"Quakedash,omitempty"`
								Megawallsdefense        int `json:"Megawallsdefense,omitempty"`
								Skywarsvoidkill         int `json:"Skywarsvoidkill,omitempty"`
								Tntrunsurviveminute     int `json:"Tntrunsurviveminute,omitempty"`
								Cvcthrowprojectile      int `json:"Cvcthrowprojectile,omitempty"`
								Vampzvampirekill        int `json:"Vampzvampirekill,omitempty"`
								Pitkill                 int `json:"Pitkill,omitempty"`
								Wwplacewool             int `json:"Wwplacewool,omitempty"`
								Tkrcollectbox           int `json:"Tkrcollectbox,omitempty"`
								Arcadezombiesdoor       int `json:"Arcadezombiesdoor,omitempty"`
								Wallswoodpickaxe        int `json:"Wallswoodpickaxe,omitempty"`
								Blitzchests             int `json:"Blitzchests,omitempty"`
								Smashthrowoff           int `json:"Smashthrowoff,omitempty"`
								Arcadekillcreeper       int `json:"Arcadekillcreeper,omitempty"`
								Maincatchfish           int `json:"Maincatchfish,omitempty"`
								Wizardscapture          int `json:"Wizardscapture,omitempty"`
								Arenaultimate           int `json:"Arenaultimate,omitempty"`
								Bbguess                 int `json:"Bbguess,omitempty"`
								Arcadeblockingdeadkills int `json:"Arcadeblockingdeadkills,omitempty"`
							} `json:"objectives,omitempty"`
						} `json:"easy,omitempty"`
						Medium struct {
							Objectives struct {
								Maincatchtreasure   int `json:"Maincatchtreasure,omitempty"`
								Arenabuffs          int `json:"Arenabuffs,omitempty"`
								Arcadesupplychests  int `json:"Arcadesupplychests,omitempty"`
								Pbtntrain           int `json:"Pbtntrain,omitempty"`
								Quakeheadshot       int `json:"Quakeheadshot,omitempty"`
								Bowspleefsurvivetwo int `json:"Bowspleefsurvivetwo,omitempty"`
								Warlordsdamageflag  int `json:"Warlordsdamageflag,omitempty"`
								Arcadetop3Round     int `json:"Arcadetop3round,omitempty"`
							} `json:"objectives,omitempty"`
						} `json:"medium,omitempty"`
					} `json:"bingo,omitempty"`
				} `json:" 2023,omitempty"`
			} `json:"anniversary,omitempty"`
			Easter struct {
				Num2022 struct {
					MainlobbyEgghunt1284514  bool `json:"mainlobby_egghunt_128_45_-14,omitempty"`
					MainlobbyEgghunt15257184 bool `json:"mainlobby_egghunt_-152_57_184,omitempty"`
					MainlobbyEgghunt56957    bool `json:"mainlobby_egghunt_-56_95_7,omitempty"`
					MainlobbyEgghunt337344   bool `json:"mainlobby_egghunt_-33_73_-44,omitempty"`
					MainlobbyEgghunt153452   bool `json:"mainlobby_egghunt_-15_34_-52,omitempty"`
					MainlobbyEgghunt252266   bool `json:"mainlobby_egghunt_25_22_66,omitempty"`
					MainlobbyEgghunt286818   bool `json:"mainlobby_egghunt_28_68_-18,omitempty"`
					MainlobbyEgghunt357022   bool `json:"mainlobby_egghunt_35_70_22,omitempty"`
					MainlobbyEgghunt366455   bool `json:"mainlobby_egghunt_36_64_55,omitempty"`
					MainlobbyEgghunt27056    bool `json:"mainlobby_egghunt_2_70_56,omitempty"`
					MainlobbyEgghunt83865    bool `json:"mainlobby_egghunt_83_86_-5,omitempty"`
					MainlobbyEgghunt947043   bool `json:"mainlobby_egghunt_94_70_43,omitempty"`
					MainlobbyEgghunt140595   bool `json:"mainlobby_egghunt_140_59_5,omitempty"`
					MainlobbyEgghunt64104129 bool `json:"mainlobby_egghunt_64_104_-129,omitempty"`
					MainlobbyEgghunt886590   bool `json:"mainlobby_egghunt_-88_65_-90,omitempty"`
					MainlobbyEgghunt13664142 bool `json:"mainlobby_egghunt_-136_64_-142,omitempty"`
					MainlobbyEgghunt17520109 bool `json:"mainlobby_egghunt_-175_20_-109,omitempty"`
					MainlobbyEgghunt2027822  bool `json:"mainlobby_egghunt_-202_78_-22,omitempty"`
					MainlobbyEgghunt1526541  bool `json:"mainlobby_egghunt_-152_65_41,omitempty"`
					MainlobbyEgghunt9812920  bool `json:"mainlobby_egghunt_-98_129_-20,omitempty"`
					MainlobbyEgghunt7710918  bool `json:"mainlobby_egghunt_-77_109_18,omitempty"`
					MainlobbyEgghunt808959   bool `json:"mainlobby_egghunt_-80_89_59,omitempty"`
					MainlobbyEgghunt11762107 bool `json:"mainlobby_egghunt_-117_62_107,omitempty"`
					MainlobbyEgghunt8265160  bool `json:"mainlobby_egghunt_-82_65_160,omitempty"`
					MainlobbyEgghunt1157203  bool `json:"mainlobby_egghunt_-11_57_203,omitempty"`
					MainlobbyEgghunt4253253  bool `json:"mainlobby_egghunt_42_53_253,omitempty"`
					MainlobbyEgghunt12652179 bool `json:"mainlobby_egghunt_126_52_179,omitempty"`
					MainlobbyEgghunt6525121  bool `json:"mainlobby_egghunt_65_25_121,omitempty"`
					MainlobbyEgghunt12967118 bool `json:"mainlobby_egghunt_129_67_118,omitempty"`
					MainlobbyEgghunt1156186  bool `json:"mainlobby_egghunt_115_61_86,omitempty"`
					MainlobbyEgghuntReward   bool `json:"mainlobby_egghunt_reward,omitempty"`
				} `json:"2022,omitempty"`
				Num2023 struct {
					MainlobbyEgghunt59945    bool `json:"mainlobby_egghunt_-59_94_-5,omitempty"`
					MainlobbyEgghunt20652345 bool `json:"mainlobby_egghunt_-206_52_345,omitempty"`
					Levelling                struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
					MainlobbyEgghunt296727   bool `json:"mainlobby_egghunt_29_67_-27,omitempty"`
					MainlobbyEgghunt228559   bool `json:"mainlobby_egghunt_-22_85_-59,omitempty"`
					MainlobbyEgghunt558756   bool `json:"mainlobby_egghunt_-55_87_56,omitempty"`
					MainlobbyEgghunt1087886  bool `json:"mainlobby_egghunt_-108_78_86,omitempty"`
					MainlobbyEgghunt7756191  bool `json:"mainlobby_egghunt_-77_56_191,omitempty"`
					MainlobbyEgghunt3354254  bool `json:"mainlobby_egghunt_33_54_254,omitempty"`
					MainlobbyEgghunt100756   bool `json:"mainlobby_egghunt_100_75_6,omitempty"`
					MainlobbyEgghunt12872134 bool `json:"mainlobby_egghunt_128_72_-134,omitempty"`
					MainlobbyEgghunt59105137 bool `json:"mainlobby_egghunt_59_105_-137,omitempty"`
					MainlobbyEgghunt14573137 bool `json:"mainlobby_egghunt_-145_73_-137,omitempty"`
					MainlobbyEgghunt058500   bool `json:"mainlobby_egghunt_0_58_-500,omitempty"`
					MainlobbyEgghunt2051436  bool `json:"mainlobby_egghunt_-20_51_-436,omitempty"`
					MainlobbyEgghunt6854399  bool `json:"mainlobby_egghunt_68_54_-399,omitempty"`
					MainlobbyEgghunt1457518  bool `json:"mainlobby_egghunt_-14_57_-518,omitempty"`
					MainlobbyEgghunt7054565  bool `json:"mainlobby_egghunt_70_54_-565,omitempty"`
					MainlobbyEgghunt11526509 bool `json:"mainlobby_egghunt_115_26_-509,omitempty"`
					MainlobbyEgghunt9625579  bool `json:"mainlobby_egghunt_96_25_-579,omitempty"`
					MainlobbyEgghunt32743137 bool `json:"mainlobby_egghunt_327_43_-137,omitempty"`
					MainlobbyEgghunt374600   bool `json:"mainlobby_egghunt_374_60_0,omitempty"`
					MainlobbyEgghunt38456149 bool `json:"mainlobby_egghunt_384_56_149,omitempty"`
					MainlobbyEgghunt37592139 bool `json:"mainlobby_egghunt_375_92_139,omitempty"`
					MainlobbyEgghunt13353354 bool `json:"mainlobby_egghunt_133_53_354,omitempty"`
					MainlobbyEgghunt13749347 bool `json:"mainlobby_egghunt_137_49_347,omitempty"`
					MainlobbyEgghunt452364   bool `json:"mainlobby_egghunt_4_52_364,omitempty"`
					MainlobbyEgghunt1857402  bool `json:"mainlobby_egghunt_-18_57_402,omitempty"`
					MainlobbyEgghunt37065228 bool `json:"mainlobby_egghunt_-370_65_228,omitempty"`
					MainlobbyEgghunt3725010  bool `json:"mainlobby_egghunt_-372_50_10,omitempty"`
					MainlobbyEgghunt878162   bool `json:"mainlobby_egghunt_-87_81_-62,omitempty"`
					MainlobbyEgghuntReward   bool `json:"mainlobby_egghunt_reward,omitempty"`
				} `json:"2023,omitempty"`
			} `json:"easter,omitempty"`
			Summer struct {
				Num2022 struct {
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
				} `json:"2022,omitempty"`
				Num2023 struct {
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
				} `json:"2023,omitempty"`
			} `json:"summer,omitempty"`
			Silver    int `json:"silver,omitempty"`
			Halloween struct {
				Num2022 struct {
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
				} `json:"2022,omitempty"`
				Num2023 struct {
					Levelling struct {
						Experience int `json:"experience,omitempty"`
					} `json:"levelling,omitempty"`
					Bingo struct {
						Easy struct {
							Objectives struct {
								Cvcbrains              int `json:"Cvcbrains,omitempty"`
								Quakedash              int `json:"Quakedash,omitempty"`
								Paintballpoultryghiest int `json:"Paintballpoultryghiest,omitempty"`
							} `json:"objectives,omitempty"`
						} `json:"easy,omitempty"`
					} `json:"bingo,omitempty"`
					SkyBlockAlchemistIntro bool `json:"skyBlockAlchemistIntro,omitempty"`
				} `json:"2023,omitempty"`
			} `json:"halloween,omitempty"`
		} `json:"seasonal,omitempty"`
		CompletedChristmasQuests2022 int    `json:"completed_christmas_quests_2022,omitempty"`
		CurrentPet                   string `json:"currentPet,omitempty"`
		ClaimedCenturyCake200        int64  `json:"claimed_century_cake200,omitempty"`
		Easter2022Cooldowns2         struct {
			MvpPlus3 bool `json:"MVP_PLUS3,omitempty"`
			MvpPlus2 bool `json:"MVP_PLUS2,omitempty"`
			MvpPlus1 bool `json:"MVP_PLUS1,omitempty"`
			MvpPlus0 bool `json:"MVP_PLUS0,omitempty"`
			Normal3  bool `json:"NORMAL3,omitempty"`
			Normal2  bool `json:"NORMAL2,omitempty"`
			Normal1  bool `json:"NORMAL1,omitempty"`
			Normal0  bool `json:"NORMAL0,omitempty"`
			VipPlus0 bool `json:"VIP_PLUS0,omitempty"`
			VipPlus1 bool `json:"VIP_PLUS1,omitempty"`
			VipPlus2 bool `json:"VIP_PLUS2,omitempty"`
			VipPlus3 bool `json:"VIP_PLUS3,omitempty"`
			Mvp0     bool `json:"MVP0,omitempty"`
			Mvp1     bool `json:"MVP1,omitempty"`
			Mvp2     bool `json:"MVP2,omitempty"`
			Mvp3     bool `json:"MVP3,omitempty"`
			Vip0     bool `json:"VIP0,omitempty"`
			Vip1     bool `json:"VIP1,omitempty"`
			Vip2     bool `json:"VIP2,omitempty"`
			Vip3     bool `json:"VIP3,omitempty"`
		} `json:"easter2022Cooldowns2,omitempty"`
		Leveling struct {
			ClaimedRewards []int `json:"claimedRewards,omitempty"`
		} `json:"leveling,omitempty"`
		Rowfive            []string `json:"Rowfive,omitempty"`
		Rowtwo             []string `json:"Rowtwo,omitempty"`
		Rowone             []string `json:"Rowone,omitempty"`
		Rowthree           []string `json:"Rowthree,omitempty"`
		Rowfour            []string `json:"Rowfour,omitempty"`
		Diagonaltwo        []string `json:"Diagonaltwo,omitempty"`
		Diagonalone        []string `json:"Diagonalone,omitempty"`
		Columnone          []string `json:"Columnone,omitempty"`
		Columnthree        []string `json:"Columnthree,omitempty"`
		Columntwo          []string `json:"Columntwo,omitempty"`
		Columnfour         []string `json:"Columnfour,omitempty"`
		Columnfive         []string `json:"Columnfive,omitempty"`
		BlackOut           []string `json:"blackOut,omitempty"`
		MostRecentGameType string   `json:"mostRecentGameType,omitempty"`
	} `json:"player,omitempty"`
}
