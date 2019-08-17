package model

type move int
type damageCategory byte
type target byte

const (
	physical damageCategory = iota
	special
	statusEffect
)

const (
	singleNotUser target = iota
	none
	singleRandomOpposing
	allOpposing
	allButUser
	user
	usersSide
	bothSides
	opposingSide
	partner
	singleUsersSide
	singleOpposingSide
	singleDirectlyOpposite
)

const (
	ABSORB move = iota
	ACIDARMOR
	ACIDSPRAY
	ACID
	ACROBATICS
	ACUPRESSURE
	AERIALACE
	AEROBLAST
	AFTERYOU
	AGILITY
	AIRCUTTER
	AIRSLASH
	ALLYSWITCH
	AMNESIA
	ANCIENTPOWER
	AQUAJET
	AQUARING
	AQUATAIL
	ARMTHRUST
	AROMATHERAPY
	ASSIST
	ASSURANCE
	ASTONISH
	ATTACKORDER
	ATTRACT
	AURASPHERE
	AURORABEAM
	AUTOTOMIZE
	AVALANCHE
	BARRAGE
	BARRIER
	BATONPASS
	BEATUP
	BELLYDRUM
	BESTOW
	BIDE
	BIND
	BITE
	BLASTBURN
	BLAZEKICK
	BLIZZARD
	BLOCK
	BLUEFLARE
	BODYSLAM
	BOLTSTRIKE
	BONECLUB
	BONERUSH
	BONEMERANG
	BOUNCE
	BRAVEBIRD
	BRICKBREAK
	BRINE
	BUBBLE
	BUBBLEBEAM
	BUGBITE
	BUGBUZZ
	BULKUP
	BULLDOZE
	BULLETPUNCH
	BULLETSEED
	CALMMIND
	CAMOUFLAGE
	CAPTIVATE
	CHARGEBEAM
	CHARGE
	CHARM
	CHATTER
	CHIPAWAY
	CIRCLETHROW
	CLAMP
	CLEARSMOG
	CLOSECOMBAT
	COIL
	COMETPUNCH
	CONFUSERAY
	CONFUSION
	CONSTRICT
	CONVERSION2
	CONVERSION
	COPYCAT
	COSMICPOWER
	COTTONGUARD
	COTTONSPORE
	COUNTER
	COVET
	CRABHAMMER
	CROSSCHOP
	CROSSPOISON
	CRUNCH
	CRUSHCLAW
	CRUSHGRIP
	CURSE
	CUT
	DARKPULSE
	DARKVOID
	DEFENDORDER
	DEFENSECURL
	DEFOG
	DESTINYBOND
	DETECT
	DIG
	DISABLE
	DISCHARGE
	DIVE
	DIZZYPUNCH
	DOOMDESIRE
	DOUBLEHIT
	DOUBLEKICK
	DOUBLETEAM
	DOUBLEEDGE
	DOUBLESLAP
	DRACOMETEOR
	DRAGONCLAW
	DRAGONDANCE
	DRAGONPULSE
	DRAGONRAGE
	DRAGONRUSH
	DRAGONTAIL
	DRAGONBREATH
	DRAINPUNCH
	DREAMEATER
	DRILLPECK
	DRILLRUN
	DUALCHOP
	DYNAMICPUNCH
	EARTHPOWER
	EARTHQUAKE
	ECHOEDVOICE
	EGGBOMB
	ELECTROBALL
	ELECTROWEB
	EMBARGO
	EMBER
	ENCORE
	ENDEAVOR
	ENDURE
	ENERGYBALL
	ENTRAINMENT
	ERUPTION
	EXPLOSION
	EXTRASENSORY
	EXTREMESPEED
	FACADE
	FEINTATTACK
	FAKEOUT
	FAKETEARS
	FALSESWIPE
	FEATHERDANCE
	FEINT
	FIERYDANCE
	FINALGAMBIT
	FIREBLAST
	FIREFANG
	FIREPLEDGE
	FIREPUNCH
	FIRESPIN
	FISSURE
	FLAIL
	FLAMEBURST
	FLAMECHARGE
	FLAMEWHEEL
	FLAMETHROWER
	FLAREBLITZ
	FLASHCANNON
	FLASH
	FLATTER
	FLING
	FLY
	FOCUSBLAST
	FOCUSENERGY
	FOCUSPUNCH
	FOLLOWME
	FORCEPALM
	FORESIGHT
	FOULPLAY
	FREEZESHOCK
	FRENZYPLANT
	FROSTBREATH
	FRUSTRATION
	FURYATTACK
	FURYCUTTER
	FURYSWIPES
	FUSIONBOLT
	FUSIONFLARE
	FUTURESIGHT
	GASTROACID
	GEARGRIND
	GIGADRAIN
	GIGAIMPACT
	GLACIATE
	GLARE
	GRASSKNOT
	GRASSPLEDGE
	GRASSWHISTLE
	GRAVITY
	GROWL
	GROWTH
	GRUDGE
	GUARDSPLIT
	GUARDSWAP
	GUILLOTINE
	GUNKSHOT
	GUST
	GYROBALL
	HAIL
	HAMMERARM
	HARDEN
	HAZE
	HEADCHARGE
	HEADSMASH
	HEADBUTT
	HEALBELL
	HEALBLOCK
	HEALORDER
	HEALPULSE
	HEALINGWISH
	HEARTSTAMP
	HEARTSWAP
	HEATCRASH
	HEATWAVE
	HEAVYSLAM
	HELPINGHAND
	HEX
	HIGHJUMPKICK
	HIDDENPOWER
	HONECLAWS
	HORNATTACK
	HORNDRILL
	HORNLEECH
	HOWL
	HURRICANE
	HYDROCANNON
	HYDROPUMP
	HYPERBEAM
	HYPERFANG
	HYPERVOICE
	HYPNOSIS
	ICEBALL
	ICEBEAM
	ICEBURN
	ICEFANG
	ICEPUNCH
	ICESHARD
	ICICLECRASH
	ICICLESPEAR
	ICYWIND
	IMPRISON
	INCINERATE
	INFERNO
	INGRAIN
	IRONDEFENSE
	IRONHEAD
	IRONTAIL
	JUDGMENT
	JUMPKICK
	KARATECHOP
	KINESIS
	KNOCKOFF
	LASTRESORT
	LAVAPLUME
	LEAFBLADE
	LEAFSTORM
	LEAFTORNADO
	LEECHLIFE
	LEECHSEED
	LEER
	LICK
	LIGHTSCREEN
	LOCKON
	LOVELYKISS
	LOWKICK
	LOWSWEEP
	LUCKYCHANT
	LUNARDANCE
	LUSTERPURGE
	MACHPUNCH
	MAGICCOAT
	MAGICROOM
	MAGICALLEAF
	MAGMASTORM
	MAGNETBOMB
	MAGNETRISE
	MAGNITUDE
	MEFIRST
	MEANLOOK
	MEDITATE
	MEGADRAIN
	MEGAKICK
	MEGAPUNCH
	MEGAHORN
	MEMENTO
	METALBURST
	METALCLAW
	METALSOUND
	METEORMASH
	METRONOME
	MILKDRINK
	MIMIC
	MINDREADER
	MINIMIZE
	MIRACLEEYE
	MIRRORCOAT
	MIRRORMOVE
	MIRRORSHOT
	MISTBALL
	MIST
	MOONLIGHT
	MORNINGSUN
	MUDBOMB
	MUDSHOT
	MUDSPORT
	MUDSLAP
	MUDDYWATER
	NASTYPLOT
	NATURALGIFT
	NATUREPOWER
	NEEDLEARM
	NIGHTDAZE
	NIGHTSHADE
	NIGHTSLASH
	NIGHTMARE
	OCTAZOOKA
	ODORSLEUTH
	OMINOUSWIND
	OUTRAGE
	OVERHEAT
	PAINSPLIT
	PAYDAY
	PAYBACK
	PECK
	PERISHSONG
	PETALDANCE
	PINMISSILE
	PLUCK
	POISONFANG
	POISONGAS
	POISONJAB
	POISONSTING
	POISONTAIL
	POISONPOWDER
	POUND
	POWDERSNOW
	POWERGEM
	POWERSPLIT
	POWERSWAP
	POWERTRICK
	POWERWHIP
	PRESENT
	PROTECT
	PSYBEAM
	PSYCHUP
	PSYCHIC_move
	PSYCHOBOOST
	PSYCHOCUT
	PSYCHOSHIFT
	PSYSHOCK
	PSYSTRIKE
	PSYWAVE
	PUNISHMENT
	PURSUIT
	QUASH
	QUICKATTACK
	QUICKGUARD
	QUIVERDANCE
	RAGEPOWDER
	RAGE
	RAINDANCE
	RAPIDSPIN
	RAZORLEAF
	RAZORSHELL
	RAZORWIND
	RECOVER
	RECYCLE
	REFLECTTYPE
	REFLECT
	REFRESH
	RELICSONG
	REST
	RETALIATE
	RETURN
	REVENGE
	REVERSAL
	ROAROFTIME
	ROAR
	ROCKBLAST
	ROCKCLIMB
	ROCKPOLISH
	ROCKSLIDE
	ROCKSMASH
	ROCKTHROW
	ROCKTOMB
	ROCKWRECKER
	ROLEPLAY
	ROLLINGKICK
	ROLLOUT
	ROOST
	ROUND
	SACREDFIRE
	SACREDSWORD
	SAFEGUARD
	SANDTOMB
	SANDATTACK
	SANDSTORM
	SCALD
	SCARYFACE
	SCRATCH
	SCREECH
	SEARINGSHOT
	SECRETPOWER
	SECRETSWORD
	SEEDBOMB
	SEEDFLARE
	SEISMICTOSS
	SELFDESTRUCT
	SHADOWBALL
	SHADOWCLAW
	SHADOWFORCE
	SHADOWPUNCH
	SHADOWSNEAK
	SHARPEN
	SHEERCOLD
	SHELLSMASH
	SHIFTGEAR
	SHOCKWAVE
	SIGNALBEAM
	SILVERWIND
	SIMPLEBEAM
	SING
	SKETCH
	SKILLSWAP
	SKULLBASH
	SKYATTACK
	SKYDROP
	SKYUPPERCUT
	SLACKOFF
	SLAM
	SLASH
	SLEEPPOWDER
	SLEEPTALK
	SLUDGEBOMB
	SLUDGEWAVE
	SLUDGE
	SMACKDOWN
	SMELLINGSALTS
	SMOG
	SMOKESCREEN
	SNARL
	SNATCH
	SNORE
	SOAK
	SOFTBOILED
	SOLARBEAM
	SONICBOOM
	SPACIALREND
	SPARK
	SPIDERWEB
	SPIKECANNON
	SPIKES
	SPITUP
	SPITE
	SPLASH
	SPORE
	STEALTHROCK
	STEAMROLLER
	STEELWING
	STOCKPILE
	STOMP
	STONEEDGE
	STOREDPOWER
	STORMTHROW
	STRENGTH
	STRINGSHOT
	STRUGGLEBUG
	STRUGGLE
	STUNSPORE
	SUBMISSION
	SUBSTITUTE
	SUCKERPUNCH
	SUNNYDAY
	SUPERFANG
	SUPERPOWER
	SUPERSONIC
	SURF
	SWAGGER
	SWALLOW
	SWEETKISS
	SWEETSCENT
	SWIFT
	SWITCHEROO
	SWORDSDANCE
	SYNCHRONOISE
	SYNTHESIS
	TACKLE
	TAILGLOW
	TAILSLAP
	TAILWHIP
	TAILWIND
	TAKEDOWN
	TAUNT
	TECHNOBLAST
	TEETERDANCE
	TELEKINESIS
	TELEPORT
	THIEF
	THRASH
	THUNDERFANG
	THUNDERWAVE
	THUNDER
	THUNDERBOLT
	THUNDERPUNCH
	THUNDERSHOCK
	TICKLE
	TORMENT
	TOXICSPIKES
	TOXIC
	TRANSFORM
	TRIATTACK
	TRICKROOM
	TRICK
	TRIPLEKICK
	TRUMPCARD
	TWINEEDLE
	TWISTER
	UTURN
	UPROAR
	VCREATE
	VACUUMWAVE
	VENOSHOCK
	VICEGRIP
	VINEWHIP
	VITALTHROW
	VOLTSWITCH
	VOLTTACKLE
	WAKEUPSLAP
	WATERGUN
	WATERPLEDGE
	WATERPULSE
	WATERSPORT
	WATERSPOUT
	WATERFALL
	WEATHERBALL
	WHIRLPOOL
	WHIRLWIND
	WIDEGUARD
	WILDCHARGE
	WILLOWISP
	WINGATTACK
	WISH
	WITHDRAW
	WONDERROOM
	WOODHAMMER
	WORKUP
	WORRYSEED
	WRAP
	WRINGOUT
	XSCISSOR
	YAWN
	ZAPCANNON
	ZENHEADBUTT
)

// TODOs:
// * replace functionCodes with functions
// * replace flags with array of constants

var moveData = []MoveData{
	{
		Name: "Absorb",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       20,
		Type:        GRASS,
		Category:    special,
		Accuracy:    100,
		PP:          25,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "A nutrient-draining attack. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name: "Acid Armor",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +2})
		},
		Type:        POISON,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user alters its cellular structure to liquefy itself, sharply raising its Defense stat.",
	},
	{
		Name: "Acid Spray",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -2})
		},
		Power:           40,
		Type:            POISON,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user spits fluid that works to melt the target. This harshly reduces the target's Sp. Def stat.",
	},
	{
		Name: "Acid",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           40,
		Type:            POISON,
		Category:        special,
		Accuracy:        100,
		PP:              30,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "The opposing team is attacked with a spray of harsh acid. The acid may also lower the targets' Sp.Def stats.",
	},
	{
		Name:         "Acrobatics",
		functionCode: "086",
		Power:        55,
		Type:         FLYING,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user nimbly strikes the target. If the user is not holding an item, this attack inflicts massive damage.",
	},
	{
		Name:         "Acupressure",
		functionCode: "037",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           30,
		Target:       singleUsersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user applies pressure to stress points, sharply boosting one of its stats.",
	},
	{
		Name:        "Aerial Ace",
		Power:       60,
		Type:        FLYING,
		Category:    physical,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user confounds the foe with speed, then slashes. The attack lands without fail.",
	},
	{
		Name:        "Aeroblast",
		Power:       100,
		Type:        FLYING,
		Category:    special,
		Accuracy:    95,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "befh",
		Description: "A vortex of air is shot at the target to inflict damage. Critical hits land more easily.",
	},
	{
		Name:         "After You",
		functionCode: "11D",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "The user helps the target and makes it use its move right after the user.",
	},
	{
		Name: "Agility",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: +2})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user relaxes and lightens its body to move faster. It sharply boosts the Speed stat.",
	},
	{
		Name:        "Air Cutter",
		Power:       55,
		Type:        FLYING,
		Category:    special,
		Accuracy:    95,
		PP:          25,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "befh",
		Description: "The user launches razor-like wind to slash the opposing team. Critical hits land more easily.",
	},
	{
		Name:            "Air Slash",
		functionCode:    "00F",
		Power:           75,
		Type:            FLYING,
		Category:        special,
		Accuracy:        95,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks with a blade of air that slices even the sky. It may also make the target flinch.",
	},
	{
		Name:         "Ally Switch",
		functionCode: "120",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           15,
		Target:       user,
		Priority:     1,
		Flags:        "",
		Description:  "The user teleports using a strange power and switches its place with one of its allies.",
	},
	{
		Name: "Amnesia",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: +2})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user temporarily empties its mind to forget its concerns. It sharply raises the user's Sp. Def stat.",
	},
	{
		Name: "Ancient Power",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: +1, defense: +1, speed: +1, spattack: +1, spdefense: +1})
		},
		Power:           60,
		Type:            ROCK,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks with a prehistoric power. It may also raise all the user's stats at once.",
	},
	{
		Name:        "Aqua Jet",
		Power:       40,
		Type:        WATER,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "abef",
		Description: "The user lunges at the target at a speed that makes it almost invisible. It is sure to strike first.",
	},
	{
		Name:         "Aqua Ring",
		functionCode: "0DA",
		Type:         WATER,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user envelops itself in a veil made of water. It regains some HP on every turn.",
	},
	{
		Name:        "Aqua Tail",
		Power:       90,
		Type:        WATER,
		Category:    physical,
		Accuracy:    90,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user attacks by swinging its tail as if it were a vicious wave in a raging storm.",
	},
	{
		Name:         "Arm Thrust",
		functionCode: "0C0",
		Power:        15,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user looses a flurry of open-palmed arm thrusts that hit two to five times in a row.",
	},
	{
		Name:         "Aromatherapy",
		functionCode: "019",
		Type:         GRASS,
		Category:     statusEffect,
		PP:           5,
		Target:       bothSides,
		Priority:     0,
		Flags:        "dk",
		Description:  "The user releases a soothing scent that heals all status problems affecting the user's party.",
	},
	{
		Name:         "Assist",
		functionCode: "0B5",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "The user hurriedly and randomly uses a move among those known by other Pokémon in the party.",
	},
	{
		Name:         "Assurance",
		functionCode: "082",
		Power:        50,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "If the target has already taken some damage in the same turn, this attack's power is doubled.",
	},
	{
		Name:            "Astonish",
		functionCode:    "00F",
		Power:           30,
		Type:            GHOST,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user attacks the target while shouting in a startling fashion. It may also make the target flinch.",
	},
	{
		Name:        "Attack Order",
		Power:       90,
		Type:        BUG,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "befh",
		Description: "The user calls out its underlings to pummel the target. Critical hits land more easily.",
	},
	{
		Name:         "Attract",
		functionCode: "016",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "If it is the opposite gender of the user, the target becomes infatuated and less likely to attack.",
	},
	{
		Name:        "Aura Sphere",
		Power:       90,
		Type:        FIGHTING,
		Category:    special,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user looses a blast of aura power from deep within its body at the target. This move is certain to hit.",
	},
	{
		Name: "Aurora Beam",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: -1})
		},
		Power:           65,
		Type:            ICE,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is hit with a rainbow-colored beam. This may also lower the target's Attack stat.",
	},
	{
		Name:         "Autotomize",
		functionCode: "031",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: +2})
		},
		Type:        STEEL,
		Category:    statusEffect,
		PP:          15,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user sheds part of its body to make itself lighter and sharply raise its Speed stat.",
	},
	{
		Name:         "Avalanche",
		functionCode: "081",
		Power:        60,
		Type:         ICE,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     -4,
		Flags:        "abef",
		Description:  "An attack move that inflicts double the damage if the user has been hurt by the target in the same turn.",
	},
	{
		Name:         "Barrage",
		functionCode: "0C0",
		Power:        15,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Round objects are hurled at the target to strike two to five times in a row.",
	},
	{
		Name: "Barrier",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +2})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user throws up a sturdy wall that sharply raises its Defense stat.",
	},
	{
		Name:         "Baton Pass",
		functionCode: "0ED",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           40,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "The user switches places with a party Pokémon in waiting, passing along any stat changes.",
	},
	{
		Name:         "Beat Up",
		functionCode: "0C1",
		Power:        1,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user gets all party Pokémon to attack the target. The more party Pokémon, the greater the number of attacks.",
	},
	{
		Name:         "Belly Drum",
		functionCode: "03A",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user maximizes its Attack stat in exchange for HP equal to half its max HP.",
	},
	{
		Name:         "Bestow",
		functionCode: "0F3",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user passes its held item to the target when the target isn't holding an item.",
	},
	{
		Name:         "Bide",
		functionCode: "0D4",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		PP:           10,
		Target:       user,
		Priority:     1,
		Flags:        "abf",
		Description:  "The user endures attacks for two turns, then strikes back to cause double the damage taken.",
	},
	{
		Name:         "Bind",
		functionCode: "0CF",
		Power:        15,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "Things such as long bodies or tentacles are used to bind and squeeze the foe for four to five turns.",
	},
	{
		Name:            "Bite",
		functionCode:    "00F",
		Power:           60,
		Type:            DARK,
		Category:        physical,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is bitten with viciously sharp fangs. It may make the target flinch.",
	},
	{
		Name:         "Blast Burn",
		functionCode: "0C2",
		Power:        150,
		Type:         FIRE,
		Category:     special,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target is razed by a fiery explosion. The user must rest on the next turn, however.",
	},
	{
		Name: "Blaze Kick",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           85,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abeh",
		Description:     "The user launches a kick that lands a critical hit more easily. It may also leave the target with a burn.",
	},
	{
		Name:         "Blizzard",
		functionCode: "00D",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Freeze{})
		},
		Power:           120,
		Type:            ICE,
		Category:        special,
		Accuracy:        70,
		PP:              5,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "A howling blizzard is summoned to strike the opposing team. It may also freeze them solid.",
	},
	{
		Name:         "Block",
		functionCode: "0EF",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user blocks the target's way with arms spread wide to prevent escape.",
	},
	{
		Name:         "Blue Flare",
		functionCode: "00A",
		// TODO: For Blue Flare only, will double the power of the next Fusion Bolt used this round.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           130,
		Type:            FIRE,
		Category:        special,
		Accuracy:        85,
		PP:              5,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks by engulfing the target in an intense, yet beautiful, blue flame. It may leave the target with a burn.",
	},
	{
		Name: "Body Slam",
		// TODO: Gen 6: For Body Slam only, power is doubled and accuracy is perfect if the target is Minimized.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           85,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user drops onto the target with its full body weight. It may leave the target with paralysis.",
	},
	{
		Name:         "Bolt Strike",
		functionCode: "007",
		// TODO: For Bolt Strike only, will double the power of the next Fusion Flare used this round.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           130,
		Type:            ELECTRIC,
		Category:        physical,
		Accuracy:        85,
		PP:              5,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user charges its target, surrounding itself with a great amount of electricity. It may leave the target with paralysis.",
	},
	{
		Name:            "Bone Club",
		functionCode:    "00F",
		Power:           65,
		Type:            GROUND,
		Category:        physical,
		Accuracy:        85,
		PP:              20,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user clubs the target with a bone. It may also make the target flinch.",
	},
	{
		Name:         "Bone Rush",
		functionCode: "0C0",
		Power:        25,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user strikes the target with a hard bone two to five times in a row.",
	},
	{
		Name:         "Bonemerang",
		functionCode: "0BD",
		Power:        50,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user throws the bone it holds. The bone loops to hit the target twice, coming and going.",
	},
	{
		Name:            "Bounce",
		functionCode:    "0CC",
		Power:           85,
		Type:            FLYING,
		Category:        physical,
		Accuracy:        85,
		PP:              5,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abefl",
		Description:     "The user bounces up high, then drops on the target on the second turn. It may also leave the target with paralysis.",
	},
	{
		Name: "Brave Bird",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 3)
		},
		Power:       120,
		Type:        FLYING,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user tucks in its wings and charges from a low altitude. The user also takes serious damage.",
	},
	{
		Name:         "Brick Break",
		functionCode: "10A",
		Power:        75,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks with a swift chop. It can also break any barrier such as Light Screen and Reflect.",
	},
	{
		Name:         "Brine",
		functionCode: "080",
		Power:        65,
		Type:         WATER,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "If the target's HP is down to about half, this attack will hit with double the power.",
	},
	{
		Name: "Bubble",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           20,
		Type:            WATER,
		Category:        special,
		Accuracy:        100,
		PP:              30,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "A spray of countless bubbles is jetted at the opposing team. It may also lower the targets' Speed stats.",
	},
	{
		Name: "Bubble Beam",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           65,
		Type:            WATER,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "A spray of bubbles is forcefully ejected at the target. It may also lower its Speed stat.",
	},
	{
		Name:         "Bug Bite",
		functionCode: "0F4",
		Power:        60,
		Type:         BUG,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user bites the target. If the target is holding a Berry, the user eats it and gains its effect.",
	},
	{
		Name: "Bug Buzz",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           90,
		Type:            BUG,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bek",
		Description:     "The user vibrates its wings to generate a damaging sound wave. It may also lower the target's Sp. Def stat.",
	},
	{
		Name: "Bulk Up",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, defense: +1})
		},
		Type:        FIGHTING,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user tenses its muscles to bulk up its body, boosting both its Attack and Defense stats.",
	},
	{
		Name: "Bulldoze",
		// TODO: For Bulldoze only, power is halved if Grassy Terrain is in effect.
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           60,
		Type:            GROUND,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 100,
		Target:          allButUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user stomps down on the ground and attacks everything in the area. Hit Pokémon's Speed stat is reduced.",
	},
	{
		Name:        "Bullet Punch",
		Power:       40,
		Type:        STEEL,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "abefj",
		Description: "The user strikes the target with tough punches as fast as bullets. This move always goes first.",
	},
	{
		Name:         "Bullet Seed",
		functionCode: "0C0",
		Power:        25,
		Type:         GRASS,
		Category:     physical,
		Accuracy:     100,
		PP:           30,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user forcefully shoots seeds at the target. Two to five seeds are shot in rapid succession.",
	},
	{
		Name: "Calm Mind",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: +1, spdefense: +1})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user quietly focuses its mind and calms its spirit to raise its Sp. Atk and Sp. Def stats.",
	},
	{
		Name:         "Camouflage",
		functionCode: "060",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user's type is changed depending on its environment, such as at water's edge, in grass, or in a cave.",
	},
	{
		Name:         "Captivate",
		functionCode: "04E",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           20,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "bce",
		Description:  "If it is the opposite gender of the user, the target is charmed into harshly lowering its Sp. Atk stat.",
	},
	{
		Name: "Charge Beam",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: +1})
		},
		Power:           50,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 70,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks with an electric charge. The user may use any remaining electricity to raise its Sp. Atk stat.",
	},
	{
		Name:         "Charge",
		functionCode: "021",
		// TODO: Until the end of the next round, the power of the user's damaging Electric attacks is doubled.
		// The effect ends if the user is switched out.
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: +1})
		},
		Type:        ELECTRIC,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user boosts the power of the Electric move it uses on the next turn. It also raises the user's Sp. Def stat.",
	},
	{
		Name: "Charm",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: -2})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user gazes at the target rather charmingly, making it less wary. The target's Attack is harshly lowered.",
	},
	{
		Name:         "Chatter",
		functionCode: "014",
		Power:        60,
		Type:         FLYING,
		Category:     special,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bk",
		Description:  "The user attacks using a sound wave based on words it has learned. It may also confuse the target.",
	},
	{
		Name:         "Chip Away",
		functionCode: "0A9",
		Power:        70,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "Looking for an opening, the user strikes continually. The target's stat changes don't affect this attack's damage.",
	},
	{
		Name:         "Circle Throw",
		functionCode: "0EC",
		Power:        60,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     -6,
		Flags:        "abe",
		Description:  "The user throws the target and drags out another Pokémon in its party. In the wild, the battle ends.",
	},
	{
		Name:         "Clamp",
		functionCode: "0CF",
		Power:        35,
		Type:         WATER,
		Category:     physical,
		Accuracy:     85,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is clamped and squeezed by the user's very thick and sturdy shell for four to five turns.",
	},
	{
		Name: "Clear Smog",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			t.statStages = emptyStages
			log.add(GenericUpdateLog{Index: ti, StatStages: emptyStages})
			log.f("%s's stat changes were removed!", t.Name)
		},
		Power:       50,
		Type:        POISON,
		Category:    special,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user attacks by throwing a clump of special mud. All status changes are returned to normal.",
	},
	{
		Name: "Close Combat",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{defense: -1, spdefense: -1})
		},
		Power:       120,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user fights the target up close without guarding itself. It also cuts the user's Defense and Sp. Def.",
	},
	{
		Name: "Coil",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, defense: +1})
			changeAccuracy(l, t, ti, +1)
		},
		Type:        POISON,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user coils up and concentrates. This raises its Attack and Defense stats as well as its accuracy.",
	},
	{
		Name:         "Comet Punch",
		functionCode: "0C0",
		Power:        18,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefj",
		Description:  "The target is hit with a flurry of punches that strike two to five times in a row.",
	},
	{
		Name: "Confuse Ray",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        GHOST,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The target is exposed to a sinister ray that triggers confusion.",
	},
	{
		Name: "Confusion",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           50,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is hit by a weak telekinetic force. It may also leave the target confused.",
	},
	{
		Name: "Constrict",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           10,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              35,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is attacked with long, creeping tentacles or vines. It may also lower the target's Speed stat.",
	},
	{
		Name:         "Conversion 2",
		functionCode: "05F",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           30,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "The user changes its type to make itself resistant to the type of the attack the opponent used last.",
	},
	{
		Name:         "Conversion",
		functionCode: "05E",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           30,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user changes its type to become the same type as one of its moves.",
	},
	{
		Name:         "Copycat",
		functionCode: "0AF",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       none,
		Priority:     0,
		Flags:        "",
		Description:  "The user mimics the move used immediately before it. The move fails if no other move has been used yet.",
	},
	{
		Name: "Cosmic Power",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +1, spdefense: +1})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user absorbs a mystical power from space to raise its Defense and Sp. Def stats.",
	},
	{
		Name: "Cotton Guard",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +3})
		},
		Type:        GRASS,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user protects itself by wrapping its body in soft cotton, drastically raising the user's Defense stat.",
	},
	{
		Name: "Cotton Spore",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -2})
		},
		Type:        GRASS,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          40,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user releases cotton-like spores that cling to the target, harshly reducing its Speed stat.",
	},
	{
		Name:         "Counter",
		functionCode: "071",
		Power:        1,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       none,
		Priority:     -5,
		Flags:        "ab",
		Description:  "A retaliation move that counters any physical attack, inflicting double the damage taken.",
	},
	{
		Name:         "Covet",
		functionCode: "0F1",
		Power:        60,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "The user endearingly approaches the target, then steals the target's held item.",
	},
	{
		Name:        "Crabhammer",
		Power:       90,
		Type:        WATER,
		Category:    physical,
		Accuracy:    90,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The target is hammered with a large pincer. Critical hits land more easily.",
	},
	{
		Name:        "Cross Chop",
		Power:       100,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    80,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The user delivers a double chop with its forearms crossed. Critical hits land more easily.",
	},
	{
		Name: "Cross Poison",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           70,
		Type:            POISON,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abefh",
		Description:     "A slashing attack with a poisonous blade that may also leave the target poisoned. Critical hits land more easily.",
	},
	{
		Name: "Crunch",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Power:           80,
		Type:            DARK,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user crunches up the target with sharp fangs. It may also lower the target's Defense stat.",
	},
	{
		Name: "Crush Claw",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Power:           75,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        95,
		PP:              10,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user slashes the target with hard and sharp claws. It may also lower the target's Defense.",
	},
	{
		Name:         "Crush Grip",
		functionCode: "08C",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is crushed with great force. The attack is more powerful the more HP the target has left.",
	},
	{
		Name:         "Curse",
		functionCode: "10D",
		Type:         GHOST,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "A move that works differently for the Ghost type than for all the other types.",
	},
	{
		Name:        "Cut",
		Power:       50,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    95,
		PP:          30,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is cut with a scythe or a claw. It can also be used to cut down thin trees.",
	},
	{
		Name:            "Dark Pulse",
		functionCode:    "00F",
		Power:           80,
		Type:            DARK,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user releases a horrible aura imbued with dark thoughts. It may also make the target flinch.",
	},
	{
		Name: "Dark Void",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        DARK,
		Category:    statusEffect,
		Accuracy:    80,
		PP:          10,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "Opposing Pokémon are dragged into a world of total darkness that makes them sleep.",
	},
	{
		Name: "Defend Order",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +1, spdefense: +1})
		},
		Type:        BUG,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user calls out its underlings to shield its body, raising its Defense and Sp. Def stats.",
	},
	{
		Name:         "Defense Curl",
		functionCode: "01E",
		// TODO: The user curls up (making the user's Ice Ball/Rollout do double damage),
		// even if Defense is not boosted. Curling up is not cumulative.
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user curls up to conceal weak spots and raise its Defense stat.",
	},
	{
		Name:         "Defog",
		functionCode: "049",
		Type:         FLYING,
		Category:     statusEffect,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "A strong wind blows away the target's obstacles such as Reflect or Light Screen. It also lowers the target's evasiveness.",
	},
	{
		Name:         "Destiny Bond",
		functionCode: "0E7",
		Type:         GHOST,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "When this move is used, if the user faints, the Pokémon that landed the knockout hit also faints.",
	},
	{
		Name:         "Detect",
		functionCode: "0AA",
		Type:         FIGHTING,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     3,
		Flags:        "",
		Description:  "It enables the user to evade all attacks. Its chance of failing rises if it is used in succession.",
	},
	{
		Name:         "Dig",
		functionCode: "0CA",
		Power:        80,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user burrows, then attacks on the second turn. It can also be used to exit dungeons.",
	},
	{
		Name:         "Disable",
		functionCode: "0B9",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "For four turns, this move prevents the target from using the move it last used.",
	},
	{
		Name: "Discharge",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           80,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          allButUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "A flare of electricity is loosed to strike the area around the user. It may also cause paralysis.",
	},
	{
		Name:         "Dive",
		functionCode: "0CB",
		Power:        80,
		Type:         WATER,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "Diving on the first turn, the user floats up and attacks on the second turn. It can be used to dive deep in the ocean.",
	},
	{
		Name: "Dizzy Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           70,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abej",
		Description:     "The target is hit with rhythmically launched punches that may also leave it confused.",
	},
	{
		Name:         "Doom Desire",
		functionCode: "111",
		Power:        140,
		Type:         STEEL,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "Two turns after this move is used, the user blasts the target with a concentrated bundle of light.",
	},
	{
		Name:         "Double Hit",
		functionCode: "0BD",
		Power:        35,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user slams the target with a long tail, vines, or tentacle. The target is hit twice in a row.",
	},
	{
		Name:         "Double Kick",
		functionCode: "0BD",
		Power:        30,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           30,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is quickly kicked twice in succession using both feet.",
	},
	{
		Name: "Double Team",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeEvasion(log, s, si, +1)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          15,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "By moving rapidly, the user makes illusory copies of itself to raise its evasiveness.",
	},
	{
		Name: "Double-Edge",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 3)
		},
		Power:       120,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "A reckless, life-risking tackle. It also damages the user by a fairly large amount, however.",
	},
	{
		Name:         "Double Slap",
		functionCode: "0C0",
		Power:        15,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is slapped repeatedly, back and forth, two to five times in a row.",
	},
	{
		Name: "Draco Meteor",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: -2})
		},
		Power:           140,
		Type:            DRAGON,
		Category:        special,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "Comets are summoned down from the sky onto the target. The attack's recoil harshly reduces the user's Sp. Atk stat.",
	},
	{
		Name:        "Dragon Claw",
		Power:       80,
		Type:        DRAGON,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user slashes the target with huge, sharp claws.",
	},
	{
		Name: "Dragon Dance",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, speed: +1})
		},
		Type:        DRAGON,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user vigorously performs a mystic, Powerful dance that boosts its Attack and Speed stats.",
	},
	{
		Name:        "Dragon Pulse",
		Power:       90,
		Type:        DRAGON,
		Category:    special,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The target is attacked with a shock wave generated by the user's gaping mouth.",
	},
	{
		Name: "Dragon Rage",
		damageFunction: func(source, target *Pokemon) (int, float64, float64) {
			return fixedDamage(40, DRAGON, target)
		},
		Power:       1,
		Type:        DRAGON,
		Category:    special,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "This attack hits the target with a shock wave of pure rage. This attack always inflicts 40,HP damage.",
	},
	{
		Name:            "Dragon Rush",
		functionCode:    "00F",
		Power:           100,
		Type:            DRAGON,
		Category:        physical,
		Accuracy:        75,
		PP:              10,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user tackles the target while exhibiting overwhelming menace. It may also make the target flinch.",
	},
	{
		Name:         "Dragon Tail",
		functionCode: "0EC",
		Power:        60,
		Type:         DRAGON,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     -6,
		Flags:        "abe",
		Description:  "The user knocks away the target and drags out another Pokémon in its party. In the wild, the battle ends.",
	},
	{
		Name: "Dragon Breath",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           60,
		Type:            DRAGON,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user exhales a mighty gust that inflicts damage. It may also leave the target with paralysis.",
	},
	{
		Name: "Drain Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       75,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefj",
		Description: "An energy-draining punch. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name: "Dream Eater",
		applicable: func(s, t *Pokemon) bool {
			switch t.NonVolatileCondition.(type) {
			case *Sleep:
				return true
			}
			return false
		},
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       100,
		Type:        PSYCHIC,
		Category:    special,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user eats the dreams of a sleeping target. It absorbs half the damage caused to heal the user's HP.",
	},
	{
		Name:        "Drill Peck",
		Power:       80,
		Type:        FLYING,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "A corkscrewing attack with the sharp beak acting as a drill.",
	},
	{
		Name:        "Drill Run",
		Power:       80,
		Type:        GROUND,
		Category:    physical,
		Accuracy:    95,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The user crashes into its target while rotating its body like a drill. Critical hits land more easily.",
	},
	{
		Name:         "Dual Chop",
		functionCode: "0BD",
		Power:        40,
		Type:         DRAGON,
		Category:     physical,
		Accuracy:     90,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks its target by hitting it with brutal strikes. The target is hit twice in a row.",
	},
	{
		Name: "Dynamic Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           100,
		Type:            FIGHTING,
		Category:        physical,
		Accuracy:        50,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abej",
		Description:     "The user punches the target with full, concentrated power. It confuses the target if it hits.",
	},
	{
		Name: "Earth Power",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           90,
		Type:            GROUND,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user makes the ground under the target erupt with power. It may also lower the target's Sp. Def.",
	},
	{
		Name:         "Earthquake",
		functionCode: "076",
		Power:        100,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user sets off an earthquake that strikes every Pokémon around it.",
	},
	{
		Name:         "Echoed Voice",
		functionCode: "092",
		Power:        40,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "befk",
		Description:  "The user attacks the target with an echoing voice. If this move is used every turn, it does greater damage.",
	},
	{
		Name:        "Egg Bomb",
		Power:       100,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    75,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "A large egg is hurled at the target with maximum force to inflict damage.",
	},
	{
		Name:         "Electro Ball",
		functionCode: "099",
		Power:        1,
		Type:         ELECTRIC,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user hurls an electric orb at the target. The faster the user is than the target, the greater the damage.",
	},
	{
		Name: "Electroweb",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:       55,
		Type:        ELECTRIC,
		Category:    special,
		Accuracy:    95,
		PP:          15,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bef",
		Description: "The user captures and attacks opposing Pokémon by using an electric net. It reduces the targets' Speed stat.",
	},
	{
		Name:         "Embargo",
		functionCode: "0F8",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "It prevents the target from using its held item. Its Trainer is also prevented from using items on it.",
	},
	{
		Name: "Ember",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           40,
		Type:            FIRE,
		Category:        special,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is attacked with small flames. It may also leave the target with a burn.",
	},
	{
		Name:         "Encore",
		functionCode: "0BC",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user compels the target to keep using only the move it last used for three turns.",
	},
	{
		Name:         "Endeavor",
		functionCode: "06E",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "An attack move that cuts down the target's HP to equal the user's HP.",
	},
	{
		Name:         "Endure",
		functionCode: "0E8",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     3,
		Flags:        "",
		Description:  "The user endures any attack with at least 1,HP. Its chance of failing rises if it is used in succession.",
	},
	{
		Name: "Energy Ball",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           80,
		Type:            GRASS,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user draws power from nature and fires it at the target. It may also lower the target's Sp. Def.",
	},
	{
		Name:         "Entrainment",
		functionCode: "066",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user dances with an odd rhythm that compels the target to mimic it, making the target's Ability the same as the user's.",
	},
	{
		Name:         "Eruption",
		functionCode: "08B",
		Power:        150,
		Type:         FIRE,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user attacks the opposing team with explosive fury. The lower the user's HP, the less powerful this attack becomes.",
	},
	{
		Name:         "Explosion",
		functionCode: "0E0",
		Power:        250,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user explodes to inflict damage on those around it. The user faints upon using this move.",
	},
	{
		Name:            "Extrasensory",
		functionCode:    "00F",
		Power:           80,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              30,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks with an odd, unseeable power. It may also make the target flinch.",
	},
	{
		Name:        "Extreme Speed",
		Power:       80,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          5,
		Target:      singleNotUser,
		Priority:    2,
		Flags:       "abef",
		Description: "The user charges the target at blinding speed. This attack always goes before any other move.",
	},
	{
		Name:         "Facade",
		functionCode: "07E",
		Power:        70,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "An attack move that doubles its power if the user is poisoned, burned, or has paralysis.",
	},
	{
		Name:        "Feint Attack",
		Power:       60,
		Type:        DARK,
		Category:    physical,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user approaches the target disarmingly, then throws a sucker punch. It hits without fail.",
	},
	{
		Name:            "Fake Out",
		functionCode:    "012",
		Power:           40,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        3,
		Flags:           "abe",
		Description:     "An attack that hits first and makes the target flinch. It only works the first turn the user is in battle.",
	},
	{
		Name: "Fake Tears",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -2})
		},
		Type:        DARK,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user feigns crying to fluster the target, harshly lowering its Sp. Def stat.",
	},
	{
		Name:         "False Swipe",
		functionCode: "0E9",
		Power:        40,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A restrained attack that prevents the target from fainting. The target is left with at least 1,HP.",
	},
	{
		Name: "Feather Dance",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: -2})
		},
		Type:        FLYING,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user covers the target's body with a mass of down that harshly lowers its Attack stat.",
	},
	{
		Name:         "Feint",
		functionCode: "0AD",
		Power:        30,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     2,
		Flags:        "",
		Description:  "An attack that hits a target using Protect or Detect. It also lifts the effects of those moves.",
	},
	{
		Name: "Fiery Dance",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: +1})
		},
		Power:           80,
		Type:            FIRE,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "Cloaked in flames, the user dances and flaps its wings. It may also raise the user's Sp. Atk stat.",
	},
	{
		Name:         "Final Gambit",
		functionCode: "0E1",
		Power:        1,
		Type:         FIGHTING,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "ab",
		Description:  "The user risks everything to attack its target. The user faints but does damage equal to the user's HP.",
	},
	{
		Name: "Fire Blast",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           120,
		Type:            FIRE,
		Category:        special,
		Accuracy:        85,
		PP:              5,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is attacked with an intense blast of all-consuming fire. It may also leave the target with a burn.",
	},
	{
		Name:            "Fire Fang",
		functionCode:    "00B",
		Power:           65,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        95,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user bites with flame-cloaked fangs. It may also make the target flinch or leave it burned.",
	},
	{
		Name:         "Fire Pledge",
		functionCode: "107",
		Power:        50,
		Type:         FIRE,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A column of fire hits opposing Pokémon. When used with its Grass equivalent, its damage increases into a vast sea of fire.",
	},
	{
		Name: "Fire Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           75,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abej",
		Description:     "The target is punched with a fiery fist. It may leave the target with a burn.",
	},
	{
		Name:         "Fire Spin",
		functionCode: "0CF",
		Power:        35,
		Type:         FIRE,
		Category:     special,
		Accuracy:     85,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target becomes trapped within a fierce vortex of fire that rages for four to five turns.",
	},
	{
		Name:         "Fissure",
		functionCode: "070",
		Power:        1,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     30,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user opens up a fissure in the ground and drops the target in. The target instantly faints if it hits.",
	},
	{
		Name:         "Flail",
		functionCode: "098",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user flails about aimlessly to attack. It becomes more powerful the less HP the user has.",
	},
	{
		Name:         "Flame Burst",
		functionCode: "074",
		Power:        70,
		Type:         FIRE,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user attacks the target with a bursting flame. The bursting flame damages Pokémon next to the target as well.",
	},
	{
		Name: "Flame Charge",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{speed: +1})
		},
		Power:           50,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user cloaks itself with flame and attacks. Building up more power, it raises the user's Speed stat.",
	},
	{
		Name: "Flame Wheel",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           60,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abeg",
		Description:     "The user cloaks itself in fire and charges at the target. It may also leave the target with a burn.",
	},
	{
		Name: "Flamethrower",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           95,
		Type:            FIRE,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is scorched with an intense blast of fire. It may also leave the target with a burn.",
	},
	{
		Name: "Flare Blitz",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 3)
			// 10% chance to paralyze. not using AddEffectChance because recoil always happens
			if random.Float64() < 0.1 {
				inflictNonVolatileCondition(l, t, ti, Burn{})
			}
		},
		Power:           120,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abefg",
		Description:     "The user cloaks itself in fire and charges at the target. The user sustains serious damage and may leave the target burned.",
	},
	{
		Name: "Flash Cannon",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           80,
		Type:            STEEL,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user gathers all its light energy and releases it at once. It may also lower the target's Sp. Def stat.",
	},
	{
		Name: "Flash",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user flashes a light that cuts the target's accuracy. It can also be used to illuminate caves.",
	},
	{
		Name: "Flatter",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(log, t, ti, Stats{spattack: +1})
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        DARK,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "Flattery is used to confuse the target. However, it also raises the target's Sp. Atk stat.",
	},
	{
		Name:         "Fling",
		functionCode: "0F7",
		Power:        1,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user flings its held item at the target to attack. Its power and effects depend on the item.",
	},
	{
		Name:         "Fly",
		functionCode: "0C9",
		Power:        90,
		Type:         FLYING,
		Category:     physical,
		Accuracy:     95,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefl",
		Description:  "The user soars, then strikes its target on the second turn. It can also be used for flying to any familiar town.",
	},
	{
		Name: "Focus Blast",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           120,
		Type:            FIGHTING,
		Category:        special,
		Accuracy:        70,
		PP:              5,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user heightens its mental focus and unleashes its power. It may also lower the target's Sp. Def.",
	},
	{
		Name:         "Focus Energy",
		functionCode: "023",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           30,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user takes a deep breath and focuses so that critical hits land more easily.",
	},
	{
		Name:         "Focus Punch",
		functionCode: "115",
		Power:        150,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     -3,
		Flags:        "abj",
		Description:  "The user focuses its mind before launching a punch. It will fail if the user is hit before it is used.",
	},
	{
		Name:         "Follow Me",
		functionCode: "117",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     3,
		Flags:        "",
		Description:  "The user draws attention to itself, making all targets take aim only at the user.",
	},
	{
		Name: "Force Palm",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           60,
		Type:            FIGHTING,
		Category:        physical,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The target is attacked with a shock wave. It may also leave the target with paralysis.",
	},
	{
		Name:         "Foresight",
		functionCode: "0A7",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "Enables a Ghost-type target to be hit by Normal- and Fighting-type attacks. It also enables an evasive target to be hit.",
	},
	{
		Name:         "Foul Play",
		functionCode: "121",
		Power:        95,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user turns the target's power against it. The higher the target's Attack stat, the greater the damage.",
	},
	{
		Name:            "Freeze Shock",
		functionCode:    "0C5",
		Power:           140,
		Type:            ICE,
		Category:        physical,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "On the second turn, the user hits the target with electrically charged ice. It may leave the target with paralysis.",
	},
	{
		Name:         "Frenzy Plant",
		functionCode: "0C2",
		Power:        150,
		Type:         GRASS,
		Category:     special,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user slams the target with an enormous tree. The user can't move on the next turn.",
	},
	{
		Name:         "Frost Breath",
		functionCode: "0A0",
		Power:        40,
		Type:         ICE,
		Category:     special,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user blows a cold breath on the target. This attack always results in a critical hit.",
	},
	{
		Name:         "Frustration",
		functionCode: "08A",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A full-power attack that grows more powerful the less the user likes its Trainer.",
	},
	{
		Name:         "Fury Attack",
		functionCode: "0C0",
		Power:        15,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is jabbed repeatedly with a horn or beak two to five times in a row.",
	},
	{
		Name:         "Fury Cutter",
		functionCode: "091",
		Power:        20,
		Type:         BUG,
		Category:     physical,
		Accuracy:     95,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is slashed with scythes or claws. Its power increases if it hits in succession.",
	},
	{
		Name:         "Fury Swipes",
		functionCode: "0C0",
		Power:        18,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     80,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The target is raked with sharp claws or scythes for two to five times in quick succession.",
	},
	{
		Name:         "Fusion Bolt",
		functionCode: "079",
		Power:        100,
		Type:         ELECTRIC,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user throws down a giant thunderbolt. This attack does greater damage when influenced by an enormous flame.",
	},
	{
		Name:         "Fusion Flare",
		functionCode: "07A",
		Power:        100,
		Type:         FIRE,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "befg",
		Description:  "The user brings down a giant flame. This attack does greater damage when influenced by an enormous thunderbolt.",
	},
	{
		Name:         "Future Sight",
		functionCode: "111",
		Power:        100,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "Two turns after this move is used, a hunk of psychic energy attacks the target.",
	},
	{
		Name:         "Gastro Acid",
		functionCode: "068",
		Type:         POISON,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user hurls up its stomach acids on the target. The fluid eliminates the effect of the target's Ability.",
	},
	{
		Name:         "Gear Grind",
		functionCode: "0BD",
		Power:        50,
		Type:         STEEL,
		Category:     physical,
		Accuracy:     85,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks by throwing two steel gears at its target.",
	},
	{
		Name: "Giga Drain",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       75,
		Type:        GRASS,
		Category:    special,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "A nutrient-draining attack. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name:         "Giga Impact",
		functionCode: "0C2",
		Power:        150,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user charges at the target using every bit of its power. The user must rest on the next turn.",
	},
	{
		Name: "Glaciate",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           65,
		Type:            ICE,
		Category:        special,
		Accuracy:        95,
		PP:              10,
		AddEffectChance: 100,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks by blowing freezing cold air at opposing Pokémon. This attack reduces the targets' Speed stat.",
	},
	{
		Name: "Glare",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    90,
		PP:          30,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user intimidates the target with the pattern on its belly to cause paralysis.",
	},
	{
		Name:         "Grass Knot",
		functionCode: "09A",
		Power:        1,
		Type:         GRASS,
		Category:     special,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user snares the target with grass and trips it. The heavier the target, the greater the damage.",
	},
	{
		Name:         "Grass Pledge",
		functionCode: "106",
		Power:        50,
		Type:         GRASS,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A column of grass hits opposing Pokémon. When used with its water equivalent, its damage increases into a vast swamp.",
	},
	{
		Name: "Grass Whistle",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        GRASS,
		Category:    statusEffect,
		Accuracy:    55,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bcek",
		Description: "The user plays a pleasant melody that lulls the target into a deep sleep.",
	},
	{
		Name:         "Gravity",
		functionCode: "118",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           5,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "Gravity is intensified for five turns, making moves involving flying unusable and negating Levitation.",
	},
	{
		Name: "Growl",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: -1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          40,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bcek",
		Description: "The user growls in an endearing way, making the opposing team less wary. The foes' Attack stats are lowered.",
	},
	{
		Name:         "Growth",
		functionCode: "028",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, spattack: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user's body grows all at once, raising the Atk and Sp. Atk stats.",
	},
	{
		Name:         "Grudge",
		functionCode: "0E6",
		Type:         GHOST,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "If the user faints, the user's grudge fully depletes the PP of the opponent's move that knocked it out.",
	},
	{
		Name:         "Guard Split",
		functionCode: "059",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "b",
		Description:  "The user employs its psychic power to average its Defense and Sp. Def stats with those of its target's.",
	},
	{
		Name: "Guard Swap",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			s.statStages.defense, t.statStages.defense = t.statStages.defense, s.statStages.defense
			s.statStages.spdefense, t.statStages.spdefense = t.statStages.spdefense, s.statStages.spdefense
			log.add(GenericUpdateLog{Index: si, StatStages: s.statStages})
			log.add(GenericUpdateLog{Index: ti, StatStages: t.statStages})
			log.f("%s switched all changes to its Defense and Sp. Def with the target!", s.Name)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user employs its psychic power to switch changes to its Defense and Sp. Def with the target.",
	},
	{
		Name:         "Guillotine",
		functionCode: "070",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     30,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "A vicious, tearing attack with big pincers. The target will faint instantly if this attack hits.",
	},
	{
		Name: "Gunk Shot",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           120,
		Type:            POISON,
		Category:        physical,
		Accuracy:        70,
		PP:              5,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user shoots filthy garbage at the target to attack. It may also poison the target.",
	},
	{
		Name:         "Gust",
		functionCode: "077",
		Power:        40,
		Type:         FLYING,
		Category:     special,
		Accuracy:     100,
		PP:           35,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A gust of wind is whipped up by wings and launched at the target to inflict damage.",
	},
	{
		Name:         "Gyro Ball",
		functionCode: "08D",
		Power:        1,
		Type:         STEEL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user tackles the target with a high-speed spin. The slower the user, the greater the damage.",
	},
	{
		Name:         "Hail",
		functionCode: "102",
		Type:         ICE,
		Category:     statusEffect,
		PP:           10,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user summons a hail storm lasting five turns. It damages all Pokémon except the Ice type.",
	},
	{
		Name: "Hammer Arm",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{speed: -1})
		},
		Power:       100,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    90,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefj",
		Description: "The user swings and hits with its strong and heavy fist. It lowers the user's Speed, however.",
	},
	{
		Name: "Harden",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user stiffens all the muscles in its body to raise its Defense stat.",
	},
	{
		Name:         "Haze",
		functionCode: "051",
		Type:         ICE,
		Category:     statusEffect,
		PP:           30,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user creates a haze that eliminates every stat change among all the Pokémon engaged in battle.",
	},
	{
		Name: "Head Charge",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 4)
		},
		Power:       120,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user charges its head into its target, using its powerful guard hair. It also damages the user a little.",
	},
	{
		Name: "Head Smash",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 2)
		},
		Power:       150,
		Type:        ROCK,
		Category:    physical,
		Accuracy:    80,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user attacks the target with a hazardous, full-power headbutt. The user also takes terrible damage.",
	},
	{
		Name:            "Headbutt",
		functionCode:    "00F",
		Power:           70,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user sticks out its head and attacks by charging straight into the target. It may also make the target flinch.",
	},
	{
		Name:         "Heal Bell",
		functionCode: "019",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "The user makes a soothing bell chime to heal the status problems of all the party Pokémon.",
	},
	{
		Name:         "Heal Block",
		functionCode: "0BB",
		Type:         PSYCHIC,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "bce",
		Description:  "For five turns, the user prevents the opposing team from using any moves, Abilities, or held items that recover HP.",
	},
	{
		Name:         "Heal Order",
		functionCode: "0D5",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp/2)
		},
		Type:        BUG,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user calls out its underlings to heal it. The user regains up to half of its max HP.",
	},
	{
		Name: "Heal Pulse",
		// TODO: Fail if target has substitute
		// cannot be used if USER is affected by Heal Block, NOT target!
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, t, ti, t.Stats.hp/2)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bci",
		Description: "The user emits a healing pulse which restores the target's HP by up to half of its max HP.",
	},
	{
		Name:         "Healing Wish",
		functionCode: "0E3",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user faints. In return, the Pokémon taking its place will have its HP restored and status cured.",
	},
	{
		Name:            "Heart Stamp",
		functionCode:    "00F",
		Power:           60,
		Type:            PSYCHIC,
		Category:        physical,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user unleashes a vicious blow after its cute act makes the target less wary. It may also make the target flinch.",
	},
	{
		Name: "Heart Swap",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			s.statStages, t.statStages = t.statStages, s.statStages
			log.add(GenericUpdateLog{Index: si, StatStages: s.statStages})
			log.add(GenericUpdateLog{Index: ti, StatStages: t.statStages})
			log.f("%s switched stat changes with the target!", s.Name)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user employs its psychic power to switch stat changes with the target.",
	},
	{
		Name:         "Heat Crash",
		functionCode: "09B",
		Power:        1,
		Type:         FIRE,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user slams its target with its flame-covered body. The more the user outweighs the target, the greater the damage.",
	},
	{
		Name: "Heat Wave",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           100,
		Type:            FIRE,
		Category:        special,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks by exhaling hot breath on the opposing team. It may also leave targets with a burn.",
	},
	{
		Name:         "Heavy Slam",
		functionCode: "09B",
		Power:        1,
		Type:         STEEL,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user slams into the target with its heavy body. The more the user outweighs the target, the greater its damage.",
	},
	{
		Name:         "Helping Hand",
		functionCode: "09C",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       partner,
		Priority:     5,
		Flags:        "",
		Description:  "The user assists an ally by boosting the power of its attack.",
	},
	{
		Name:         "Hex",
		functionCode: "07F",
		Power:        50,
		Type:         GHOST,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "This relentless attack does massive damage to a target affected by status problems.",
	},
	{
		Name:         "High Jump Kick",
		functionCode: "10B",
		Power:        130,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefl",
		Description:  "The target is attacked with a knee kick from a jump. If it misses, the user is hurt instead.",
	},
	{
		Name:         "Hidden Power",
		functionCode: "090",
		Power:        1,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A unique attack that varies in type and intensity depending on the Pokémon using it.",
	},
	{
		Name: "Hone Claws",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1})
			changeAccuracy(l, t, ti, +1)
		},
		Type:        DARK,
		Category:    statusEffect,
		PP:          15,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user sharpens its claws to boost its Attack stat and accuracy.",
	},
	{
		Name:        "Horn Attack",
		Power:       65,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          25,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is jabbed with a sharply pointed horn to inflict damage.",
	},
	{
		Name:         "Horn Drill",
		functionCode: "070",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     30,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "The user stabs the target with a horn that rotates like a drill. If it hits, the target faints instantly.",
	},
	{
		Name: "Horn Leech",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       75,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abe",
		Description: "The user drains the target's energy with its horns. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name: "Howl",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user howls loudly to raise its spirit, boosting its Attack stat.",
	},
	{
		Name:            "Hurricane",
		functionCode:    "015",
		Power:           120,
		Type:            FLYING,
		Category:        special,
		Accuracy:        70,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks by wrapping its opponent in a fierce wind that flies up into the sky. It may also confuse the target.",
	},
	{
		Name:         "Hydro Cannon",
		functionCode: "0C2",
		Power:        150,
		Type:         WATER,
		Category:     special,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target is hit with a watery blast. The user must rest on the next turn, however.",
	},
	{
		Name:        "Hydro Pump",
		Power:       120,
		Type:        WATER,
		Category:    special,
		Accuracy:    80,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The target is blasted by a huge volume of water launched under great pressure.",
	},
	{
		Name:         "Hyper Beam",
		functionCode: "0C2",
		Power:        150,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target is attacked with a powerful beam. The user must rest on the next turn to regain its energy.",
	},
	{
		Name:            "Hyper Fang",
		functionCode:    "00F",
		Power:           80,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        90,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user bites hard on the target with its sharp front fangs. It may also make the target flinch.",
	},
	{
		Name:        "Hyper Voice",
		Power:       90,
		Type:        NORMAL,
		Category:    special,
		Accuracy:    100,
		PP:          10,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bek",
		Description: "The user lets loose a horribly echoing shout with the power to inflict damage.",
	},
	{
		Name: "Hypnosis",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		Accuracy:    60,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user employs hypnotic suggestion to make the target fall into a deep sleep.",
	},
	{
		Name:         "Ice Ball",
		functionCode: "0D3",
		Power:        30,
		Type:         ICE,
		Category:     physical,
		Accuracy:     90,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user continually rolls into the target over five turns. It becomes stronger each time it hits.",
	},
	{
		Name: "Ice Beam",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Freeze{})
		},
		Power:           95,
		Type:            ICE,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is struck with an icy-cold beam of energy. It may also freeze the target solid.",
	},
	{
		Name:            "Ice Burn",
		functionCode:    "0C6",
		Power:           140,
		Type:            ICE,
		Category:        special,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "On the second turn, an ultracold, freezing wind surrounds the target. This may leave the target with a burn.",
	},
	{
		Name:            "Ice Fang",
		functionCode:    "00E",
		Power:           65,
		Type:            ICE,
		Category:        physical,
		Accuracy:        95,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user bites with cold-infused fangs. It may also make the target flinch or leave it frozen.",
	},
	{
		Name: "Ice Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Freeze{})
		},
		Power:           75,
		Type:            ICE,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abej",
		Description:     "The target is punched with an icy fist. It may also leave the target frozen.",
	},
	{
		Name:        "Ice Shard",
		Power:       40,
		Type:        ICE,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "bef",
		Description: "The user flash freezes chunks of ice and hurls them at the target. This move always goes first.",
	},
	{
		Name:            "Icicle Crash",
		functionCode:    "00F",
		Power:           85,
		Type:            ICE,
		Category:        physical,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks by harshly dropping an icicle onto the target. It may also make the target flinch.",
	},
	{
		Name:         "Icicle Spear",
		functionCode: "0C0",
		Power:        25,
		Type:         ICE,
		Category:     physical,
		Accuracy:     100,
		PP:           30,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user launches sharp icicles at the target. It strikes two to five times in a row.",
	},
	{
		Name: "Icy Wind",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           55,
		Type:            ICE,
		Category:        special,
		Accuracy:        95,
		PP:              15,
		AddEffectChance: 100,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks with a gust of chilled air. It also lowers the targets' Speed stat.",
	},
	{
		Name:         "Imprison",
		functionCode: "0B8",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "If the opponents know any move also known by the user, the opponents are prevented from using it.",
	},
	{
		Name:         "Incinerate",
		functionCode: "0F5",
		Power:        30,
		Type:         FIRE,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "be",
		Description:  "The user attacks the target with fire. If the target is holding a Berry, the Berry becomes burnt up and unusable.",
	},
	{
		Name: "Inferno",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           100,
		Type:            FIRE,
		Category:        special,
		Accuracy:        50,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks by engulfing the target in an intense fire. It leaves the target with a burn.",
	},
	{
		Name:         "Ingrain",
		functionCode: "0DB",
		Type:         GRASS,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user lays roots that restore its HP on every turn. Because it is rooted, it can't switch out.",
	},
	{
		Name: "Iron Defense",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +2})
		},
		Type:        STEEL,
		Category:    statusEffect,
		PP:          15,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user hardens its body's surface like iron, sharply raising its Defense stat.",
	},
	{
		Name:            "Iron Head",
		functionCode:    "00F",
		Power:           80,
		Type:            STEEL,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The foe slams the target with its steel-hard head. It may also make the target flinch.",
	},
	{
		Name: "Iron Tail",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Power:           100,
		Type:            STEEL,
		Category:        physical,
		Accuracy:        75,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is slammed with a steel-hard tail. It may also lower the target's Defense stat.",
	},
	{
		Name:         "Judgment",
		functionCode: "09F",
		Power:        100,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user releases countless shots of light at the target. Its type varies with the kind of Plate the user is holding.",
	},
	{
		Name:         "Jump Kick",
		functionCode: "10B",
		Power:        100,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     95,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefl",
		Description:  "The user jumps up high, then strikes with a kick. If the kick misses, the user hurts itself.",
	},
	{
		Name:        "Karate Chop",
		Power:       50,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          25,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The target is attacked with a sharp chop. Critical hits land more easily.",
	},
	{
		Name: "Kinesis",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		Accuracy:    80,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user distracts the target by bending a spoon. It lowers the target's accuracy.",
	},
	{
		Name:         "Knock Off",
		functionCode: "0F0",
		Power:        20,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "The user slaps down the target's held item, preventing that item from being used in the battle.",
	},
	{
		Name:         "Last Resort",
		functionCode: "125",
		Power:        140,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "This move can be used only after the user has used all the other moves it knows in the battle.",
	},
	{
		Name: "Lava Plume",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           80,
		Type:            FIRE,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          allButUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "An inferno of scarlet flames torches everything around the user. It may leave targets with a burn.",
	},
	{
		Name:        "Leaf Blade",
		Power:       90,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The user handles a sharp leaf like a sword and attacks by cutting its target. Critical hits land more easily.",
	},
	{
		Name: "Leaf Storm",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: -2})
		},
		Power:           140,
		Type:            GRASS,
		Category:        special,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user whips up a storm of leaves around the target. The attack's recoil harshly reduces the user's Sp. Atk stat.",
	},
	{
		Name: "Leaf Tornado",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           65,
		Type:            GRASS,
		Category:        special,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks its target by encircling it in sharp leaves. This attack may also lower the foe's accuracy.",
	},
	{
		Name: "Leech Life",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       20,
		Type:        BUG,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abe",
		Description: "The user drains the target's blood. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name:         "Leech Seed",
		functionCode: "0DC",
		Type:         GRASS,
		Category:     statusEffect,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "A seed is planted on the target. It steals some HP from the target every turn.",
	},
	{
		Name: "Leer",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          30,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "The opposing team gains an intimidating leer with sharp eyes. The opposing team's Defense stats are reduced.",
	},
	{
		Name: "Lick",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           20,
		Type:            GHOST,
		Category:        physical,
		Accuracy:        100,
		PP:              30,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is licked with a long tongue, causing damage. It may also leave the target with paralysis.",
	},
	{
		Name:         "Light Screen",
		functionCode: "0A3",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           30,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "A wondrous wall of light is put up to suppress damage from special attacks for five turns.",
	},
	{
		Name:         "Lock-On",
		functionCode: "0A6",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user takes sure aim at the target. It ensures the next attack does not fail to hit the target.",
	},
	{
		Name: "Lovely Kiss",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "With a scary face, the user tries to force a kiss on the target. If it suceeds, the target falls asleep.",
	},
	{
		Name:         "Low Kick",
		functionCode: "09A",
		Power:        1,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A powerful low kick that makes the target fall over. It inflicts greater damage on heavier targets.",
	},
	{
		Name: "Low Sweep",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:       60,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user attacks the target's legs swiftly, reducing the target's Speed stat.",
	},
	{
		Name:         "Lucky Chant",
		functionCode: "0A1",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           30,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "The user chants an incantation toward the sky, preventing opposing Pokémon from landing critical hits.",
	},
	{
		Name:         "Lunar Dance",
		functionCode: "0E4",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user faints. In return, the Pokémon taking its place will have its status and HP fully restored.",
	},
	{
		Name: "Luster Purge",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           70,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user lets loose a damaging burst of light. It may also reduce the target's Sp. Def stat.",
	},
	{
		Name:        "Mach Punch",
		Power:       40,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "abefj",
		Description: "The user throws a punch at blinding speed. It is certain to strike first.",
	},
	{
		Name:         "Magic Coat",
		functionCode: "0B1",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           15,
		Target:       user,
		Priority:     4,
		Flags:        "",
		Description:  "A barrier reflects back to the target moves like Leech Seed and moves that damage status.",
	},
	{
		Name:         "Magic Room",
		functionCode: "0F9",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       usersSide,
		Priority:     -7,
		Flags:        "e",
		Description:  "The user creates a bizarre area in which Pokémon's held items lose their effects for five turns.",
	},
	{
		Name:        "Magical Leaf",
		Power:       60,
		Type:        GRASS,
		Category:    special,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user scatters curious leaves that chase the target. This attack will not miss.",
	},
	{
		Name:         "Magma Storm",
		functionCode: "0CF",
		Power:        120,
		Type:         FIRE,
		Category:     special,
		Accuracy:     75,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target becomes trapped within a maelstrom of fire that rages for four to five turns.",
	},
	{
		Name:        "Magnet Bomb",
		Power:       60,
		Type:        STEEL,
		Category:    physical,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user launches steel bombs that stick to the target. This attack will not miss.",
	},
	{
		Name:         "Magnet Rise",
		functionCode: "119",
		Type:         ELECTRIC,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "dl",
		Description:  "The user levitates using electrically generated magnetism for five turns.",
	},
	{
		Name:         "Magnitude",
		functionCode: "095",
		Power:        1,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     100,
		PP:           30,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user looses a ground-shaking quake affecting everyone around the user. Its power varies.",
	},
	{
		Name:         "Me First",
		functionCode: "0B0",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       singleOpposingSide,
		Priority:     0,
		Flags:        "b",
		Description:  "The user tries to cut ahead of the target to steal and use the target's intended move with greater power.",
	},
	{
		Name:         "Mean Look",
		functionCode: "0EF",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user pins the target with a dark, arresting look. The target becomes unable to flee.",
	},
	{
		Name: "Meditate",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1})
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user meditates to awaken the power deep within its body and raise its Attack stat.",
	},
	{
		Name: "Mega Drain",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			drain(log, s, si, dmg, t.Name)
		},
		Power:       40,
		Type:        GRASS,
		Category:    special,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "A nutrient-draining attack. The user's HP is restored by half the damage taken by the target.",
	},
	{
		Name:        "Mega Kick",
		Power:       120,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    75,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is attacked by a kick launched with muscle-packed power.",
	},
	{
		Name:        "Mega Punch",
		Power:       80,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    85,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefj",
		Description: "The target is slugged by a punch thrown with muscle-packed power.",
	},
	{
		Name:        "Megahorn",
		Power:       120,
		Type:        BUG,
		Category:    physical,
		Accuracy:    85,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "Using its tough and impressive horn, the user rams into the target with no letup.",
	},
	{
		Name:         "Memento",
		functionCode: "0E2",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "e",
		Description:  "The user faints when using this move. In return, it harshly lowers the target's Attack and Sp. Atk.",
	},
	{
		Name:         "Metal Burst",
		functionCode: "073",
		Power:        1,
		Type:         STEEL,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       none,
		Priority:     0,
		Flags:        "be",
		Description:  "The user retaliates with much greater power against the target that last inflicted damage on it.",
	},
	{
		Name: "Metal Claw",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: +1})
		},
		Power:           50,
		Type:            STEEL,
		Category:        physical,
		Accuracy:        95,
		PP:              35,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is raked with steel claws. It may also raise the user's Attack stat.",
	},
	{
		Name: "Metal Sound",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -2})
		},
		Type:        STEEL,
		Category:    statusEffect,
		Accuracy:    85,
		PP:          40,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bcek",
		Description: "A horrible sound like scraping metal harshly reduces the target's Sp. Def stat.",
	},
	{
		Name: "Meteor Mash",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: +1})
		},
		Power:           100,
		Type:            STEEL,
		Category:        physical,
		Accuracy:        85,
		PP:              10,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abefj",
		Description:     "The target is hit with a hard punch fired like a meteor. It may also raise the user's Attack.",
	},
	{
		Name:         "Metronome",
		functionCode: "0B6",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       none,
		Priority:     0,
		Flags:        "",
		Description:  "The user waggles a finger and stimulates its brain into randomly using nearly any move.",
	},
	{
		Name:         "Milk Drink",
		functionCode: "0D5",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp/2)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user restores its own HP by up to half of its maximum HP. May also be used in the field to heal HP.",
	},
	{
		Name:         "Mimic",
		functionCode: "05C",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "b",
		Description:  "The user copies the target's last move. The move can be used during battle until the Pokémon is switched out.",
	},
	{
		Name:         "Mind Reader",
		functionCode: "0A6",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user senses the target's movements with its mind to ensure its next attack does not miss the target.",
	},
	{
		Name:         "Minimize",
		functionCode: "034",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeEvasion(log, t, ti, +2)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user compresses its body to make itself look smaller, which sharply raises its evasiveness.",
	},
	{
		Name:         "Miracle Eye",
		functionCode: "0A8",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "Enables a Dark-type target to be hit by Psychic-type attacks. It also enables an evasive target to be hit.",
	},
	{
		Name:         "Mirror Coat",
		functionCode: "072",
		Power:        1,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           20,
		Target:       none,
		Priority:     -5,
		Flags:        "b",
		Description:  "A retaliation move that counters any special attack, inflicting double the damage taken.",
	},
	{
		Name:         "Mirror Move",
		functionCode: "0AE",
		Type:         FLYING,
		Category:     statusEffect,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "The user counters the target by mimicking the target's last move.",
	},
	{
		Name: "Mirror Shot",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           65,
		Type:            STEEL,
		Category:        special,
		Accuracy:        85,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user looses a flash of energy at the target from its polished body. It may also lower the target's accuracy.",
	},
	{
		Name: "Mist Ball",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: -1})
		},
		Power:           70,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "A mistlike flurry of down envelops and damages the target. It may also lower the target's Sp. Atk.",
	},
	{
		Name:         "Mist",
		functionCode: "056",
		Type:         ICE,
		Category:     statusEffect,
		PP:           30,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "The user cloaks its body with a white mist that prevents any of its stats from being cut for five turns.",
	},
	{
		Name:         "Moonlight",
		functionCode: "0D8",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user restores its own HP. The amount of HP regained varies with the weather.",
	},
	{
		Name:         "Morning Sun",
		functionCode: "0D8",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user restores its own HP. The amount of HP regained varies with the weather.",
	},
	{
		Name: "Mud Bomb",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           65,
		Type:            GROUND,
		Category:        special,
		Accuracy:        85,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user launches a hard-packed mud ball to attack. It may also lower the target's accuracy.",
	},
	{
		Name: "Mud Shot",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           55,
		Type:            GROUND,
		Category:        special,
		Accuracy:        95,
		PP:              15,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks by hurling a blob of mud at the target. It also reduces the target's Speed.",
	},
	{
		Name:         "Mud Sport",
		functionCode: "09D",
		Type:         GROUND,
		Category:     statusEffect,
		PP:           15,
		Target:       bothSides,
		Priority:     0,
		Flags:        "",
		Description:  "The user covers itself with mud. It weakens Electric-type moves while the user is in the battle.",
	},
	{
		Name: "Mud-Slap",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           20,
		Type:            GROUND,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user hurls mud in the target's face to inflict damage and lower its accuracy.",
	},
	{
		Name: "Muddy Water",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           95,
		Type:            WATER,
		Category:        special,
		Accuracy:        85,
		PP:              10,
		AddEffectChance: 30,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks by shooting muddy water at the opposing team. It may also lower the target's accuracy.",
	},
	{
		Name: "Nasty Plot",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: +2})
		},
		Type:        DARK,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user stimulates its brain by thinking bad thoughts. It sharply raises the user's Sp. Atk.",
	},
	{
		Name:         "Natural Gift",
		functionCode: "096",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user draws power to attack by using its held Berry. The Berry determines its type and power.",
	},
	{
		Name:         "Nature Power",
		functionCode: "0B3",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       none,
		Priority:     0,
		Flags:        "",
		Description:  "An attack that makes use of nature's power. Its effects vary depending on the user's environment.",
	},
	{
		Name:            "Needle Arm",
		functionCode:    "00F",
		Power:           60,
		Type:            GRASS,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user attacks by wildly swinging its thorny arms. It may also make the target flinch.",
	},
	{
		Name: "Night Daze",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           85,
		Type:            DARK,
		Category:        special,
		Accuracy:        95,
		PP:              10,
		AddEffectChance: 40,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user lets loose a pitch-black shock wave at its target. It may also lower the target's accuracy.",
	},
	{
		Name: "Night Shade",
		damageFunction: func(source, target *Pokemon) (int, float64, float64) {
			return fixedDamage(source.Level, GHOST, target)
		},
		Power:       1,
		Type:        GHOST,
		Category:    special,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user makes the target see a frightening mirage. It inflicts damage matching the user's level.",
	},
	{
		Name:        "Night Slash",
		Power:       70,
		Type:        DARK,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The user slashes the target the instant an opportunity arises. Critical hits land more easily.",
	},
	{
		Name:         "Nightmare",
		functionCode: "10F",
		Type:         GHOST,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "A sleeping target sees a nightmare that inflicts some damage every turn.",
	},
	{
		Name: "Octazooka",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Power:           65,
		Type:            WATER,
		Category:        special,
		Accuracy:        85,
		PP:              10,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks by spraying ink in the target's face or eyes. It may also lower the target's accuracy.",
	},
	{
		Name:         "Odor Sleuth",
		functionCode: "0A7",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "Enables a Ghost-type target to be hit with Normal- and Fighting-type attacks. It also enables an evasive target to be hit.",
	},
	{
		Name: "Ominous Wind",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: +1, defense: +1, speed: +1, spattack: +1, spdefense: +1})
		},
		Power:           60,
		Type:            GHOST,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user blasts the target with a gust of repulsive wind. It may also raise all the user's stats at once.",
	},
	{
		Name:         "Outrage",
		functionCode: "0D2",
		Power:        120,
		Type:         DRAGON,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleRandomOpposing,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user rampages and attacks for two to three turns. It then becomes confused, however.",
	},
	{
		Name: "Overheat",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: -2})
		},
		Power:           140,
		Type:            FIRE,
		Category:        special,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks the target at full power. The attack's recoil sharply reduces the user's Sp. Atk stat.",
	},
	{
		Name: "Pain Split",
		// TODO: Gen III-IV: Pain Split now bypasses accuracy checks to always hit,
		// unless the opponent is in the semi-invulnerable turn of a move such as Dig or Fly.
		// It will fail if the target is behind a substitute.
		// TODO: Gen V: Pain Split can now affect a target's substitute,
		// and the user of the move will gain HP depending on the HP the substitute lost.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			hp := (s.currentHP + t.currentHP) / 2
			if s.currentHP > hp {
				s.TakeDamage(s.currentHP - hp)
			} else {
				s.Heal(hp - s.currentHP)
			}
			if t.currentHP > hp {
				t.TakeDamage(t.currentHP - hp)
			} else {
				t.Heal(hp - t.currentHP)
			}
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user adds its HP to the target's HP, then equally shares the combined HP with the target.",
	},
	{
		Name:         "Pay Day",
		functionCode: "109",
		Power:        40,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Numerous coins are hurled at the target to inflict damage. Money is earned after battle.",
	},
	{
		Name:         "Payback",
		functionCode: "084",
		Power:        50,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "If the user moves after the target, this attack's power will be doubled.",
	},
	{
		Name:        "Peck",
		Power:       35,
		Type:        FLYING,
		Category:    physical,
		Accuracy:    100,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is jabbed with a sharply pointed beak or horn.",
	},
	{
		Name:         "Perish Song",
		functionCode: "0E5",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           5,
		Target:       usersSide,
		Priority:     0,
		Flags:        "k",
		Description:  "Any Pokémon that hears this song faints in three turns, unless it switches out of battle.",
	},
	{
		Name:         "Petal Dance",
		functionCode: "0D2",
		Power:        120,
		Type:         GRASS,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleRandomOpposing,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks the target by scattering petals for two to three turns. The user then becomes confused.",
	},
	{
		Name:         "Pin Missile",
		functionCode: "0C0",
		Power:        14,
		Type:         BUG,
		Category:     physical,
		Accuracy:     85,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Sharp spikes are shot at the target in rapid succession. They hit two to five times in a row.",
	},
	{
		Name:         "Pluck",
		functionCode: "0F4",
		Power:        60,
		Type:         FLYING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user pecks the target. If the target is holding a Berry, the user eats it and gains its effect.",
	},
	{
		Name: "Poison Fang",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewBadPoison())
		},
		Power:           50,
		Type:            POISON,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30, // TODO: 50% since Gen VI
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user bites the target with toxic fangs. It may also leave the target badly poisoned.",
	},
	{
		Name: "Poison Gas",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Type:        POISON,
		Category:    statusEffect,
		Accuracy:    80,
		PP:          40,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "A cloud of poison gas is sprayed in the face of opposing Pokémon. It may poison those hit.",
	},
	{
		Name: "Poison Jab",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           80,
		Type:            POISON,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The target is stabbed with a tentacle or arm seeped with poison. It may also poison the target.",
	},
	{
		Name: "Poison Sting",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           15,
		Type:            POISON,
		Category:        physical,
		Accuracy:        100,
		PP:              35,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user stabs the target with a poisonous stinger. This may also poison the target.",
	},
	{
		Name: "Poison Tail",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           50,
		Type:            POISON,
		Category:        physical,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abefh",
		Description:     "The user hits the target with its tail. It may also poison the target. Critical hits land more easily.",
	},
	{
		Name: "Poison Powder",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Type:        POISON,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user scatters a cloud of poisonous dust on the target. It may poison the target.",
	},
	{
		Name:        "Pound",
		Power:       40,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is physically pounded with a long tail or a foreleg, etc.",
	},
	{
		Name: "Powder Snow",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Freeze{})
		},
		Power:           40,
		Type:            ICE,
		Category:        special,
		Accuracy:        100,
		PP:              25,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks with a chilling gust of powdery snow. It may also freeze the targets.",
	},
	{
		Name:        "Power Gem",
		Power:       70,
		Type:        ROCK,
		Category:    special,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user attacks with a ray of light that sparkles as if it were made of gemstones.",
	},
	{
		Name:         "Power Split",
		functionCode: "058",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "b",
		Description:  "The user employs its psychic power to average its Attack and Sp. Atk stats with those of the target's.",
	},
	{
		Name: "Power Swap",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			s.statStages.attack, t.statStages.attack = t.statStages.attack, s.statStages.attack
			s.statStages.spattack, t.statStages.spattack = t.statStages.spattack, s.statStages.spattack
			log.add(GenericUpdateLog{Index: si, StatStages: s.statStages})
			log.add(GenericUpdateLog{Index: ti, StatStages: t.statStages})
			log.f("%s switched all changes to its Attack and Sp. Atk with the target!", s.Name)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user employs its psychic power to switch changes to its Attack and Sp. Atk with the target.",
	},
	{
		Name:         "Power Trick",
		functionCode: "057",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user employs its psychic power to switch its Attack with its Defense stat.",
	},
	{
		Name:        "Power Whip",
		Power:       120,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    85,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user violently whirls its vines or tentacles to harshly lash the target.",
	},
	{
		Name:         "Present",
		functionCode: "094",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     90,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user attacks by giving the target a gift with a hidden trap. It restores HP sometimes, however.",
	},
	{
		Name:         "Protect",
		functionCode: "0AA",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     4,
		Flags:        "",
		Description:  "It enables the user to evade all attacks. Its chance of failing rises if it is used in succession.",
	},
	{
		Name: "Psybeam",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           65,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is attacked with a peculiar ray. It may also cause confusion.",
	},
	{
		Name: "Psych Up",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			s.statStages = t.statStages
			log.add(GenericUpdateLog{Index: si, StatStages: s.statStages})
			log.f("%s copied %s's stat changes!", s.Name, t.Name)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "",
		Description: "The user hypnotizes itself into copying any stat change made by the target.",
	},
	{
		Name: "Psychic",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           90,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is hit by a strong telekinetic force. It may also reduce the target's Sp. Def stat.",
	},
	{
		Name: "Psycho Boost",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{spattack: -2})
		},
		Power:           140,
		Type:            PSYCHIC,
		Category:        special,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks the target at full power. The attack's recoil harshly reduces the user's Sp. Atk stat.",
	},
	{
		Name:        "Psycho Cut",
		Power:       70,
		Type:        PSYCHIC,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "befh",
		Description: "The user tears at the target with blades formed by psychic power. Critical hits land more easily.",
	},
	{
		Name:         "Psycho Shift",
		functionCode: "01B",
		Type:         PSYCHIC,
		Category:     statusEffect,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "Using its psychic power of suggestion, the user transfers its status problems to the target.",
	},
	{
		Name:         "Psyshock",
		functionCode: "122",
		Power:        80,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user materializes an odd psychic wave to attack the target. This attack does physical damage.",
	},
	{
		Name:         "Psystrike",
		functionCode: "122",
		Power:        100,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user materializes an odd psychic wave to attack the target. This attack does physical damage.",
	},
	{
		Name:         "Psywave",
		functionCode: "06F",
		Power:        1,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     80,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The target is attacked with an odd psychic wave. The attack varies in intensity.",
	},
	{
		Name:         "Punishment",
		functionCode: "08F",
		Power:        1,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "This attack's power increases the more the target has powered up with stat changes.",
	},
	{
		Name:         "Pursuit",
		functionCode: "088",
		Power:        40,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "An attack move that inflicts double damage if used on a target that is switching out of battle.",
	},
	{
		Name:         "Quash",
		functionCode: "11E",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user suppresses the target and makes its move go last.",
	},
	{
		Name:        "Quick Attack",
		Power:       40,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "abef",
		Description: "The user lunges at the target at a speed that makes it almost invisible. It is sure to strike first.",
	},
	{
		Name:         "Quick Guard",
		functionCode: "0AB",
		Type:         FIGHTING,
		Category:     statusEffect,
		PP:           15,
		Target:       bothSides,
		Priority:     3,
		Flags:        "d",
		Description:  "The user protects itself and its allies from priority moves. If used in succession, its chance of failing rises.",
	},
	{
		Name: "Quiver Dance",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: +1, spdefense: +1, speed: +1})
		},
		Type:        BUG,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user lightly performs a beautiful, mystic dance. It boosts the user's Sp. Atk, Sp. Def, and Speed stats.",
	},
	{
		Name:         "Rage Powder",
		functionCode: "117",
		Type:         BUG,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     3,
		Flags:        "",
		Description:  "The user scatters a cloud of irritating powder to draw attention to itself. Opponents aim only at the user.",
	},
	{
		Name:         "Rage",
		functionCode: "093",
		Power:        20,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "As long as this move is in use, the power of rage raises the Attack stat each time the user is hit in battle.",
	},
	{
		Name:         "Rain Dance",
		functionCode: "100",
		Type:         WATER,
		Category:     statusEffect,
		PP:           5,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user summons a heavy rain that falls for five turns, Powering up Water-type moves.",
	},
	{
		Name:         "Rapid Spin",
		functionCode: "110",
		Power:        20,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           40,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A spin attack that can also eliminate such moves as Bind, Wrap, Leech Seed, and Spikes.",
	},
	{
		Name:        "Razor Leaf",
		Power:       55,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    95,
		PP:          25,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "befh",
		Description: "Sharp-edged leaves are launched to slash at the opposing team. Critical hits land more easily.",
	},
	{
		Name: "Razor Shell",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Power:           75,
		Type:            WATER,
		Category:        physical,
		Accuracy:        95,
		PP:              10,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user cuts its target with sharp shells. This attack may also lower the target's Defense stat.",
	},
	{
		Name:         "Razor Wind",
		functionCode: "0C3",
		Power:        80,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "befh",
		Description:  "A two-turn attack. Blades of wind hit opposing Pokémon on the second turn. Critical hits land more easily.",
	},
	{
		Name:         "Recover",
		functionCode: "0D5",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp/2)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "Restoring its own cells, the user restores its own HP by half of its max HP.",
	},
	{
		Name:         "Recycle",
		functionCode: "0F6",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user recycles a held item that has been used in battle so it can be used again.",
	},
	{
		Name:         "Reflect Type",
		functionCode: "062",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "b",
		Description:  "The user reflects the target's type, making it the same type as the target.",
	},
	{
		Name:         "Reflect",
		functionCode: "0A2",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           20,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "A wondrous wall of light is put up to suppress damage from physical attacks for five turns.",
	},
	{
		Name: "Refresh",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			if s.NonVolatileCondition == nil {
				return
			}
			switch s.NonVolatileCondition.(type) {
			case Poison, *BadPoison, Burn, Paralysis:
				log.f("%s's status returned to normal!", s.Name)
				s.clearNonVolatile()
			default:
				return
			}
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user rests to cure itself of a poisoning, burn, or paralysis.",
	},
	{
		Name:         "Relic Song",
		functionCode: "003",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Power:           75,
		Type:            NORMAL,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "bek",
		Description:     "The user sings an ancient song and attacks by appealing to the hearts of those listening. It may also induce sleep.",
	},
	{
		Name:         "Rest",
		functionCode: "0D9",
		// TODO: this will not trigger as it will fail because of Sleep condition first
		// I guess mark some moves as applicable when asleep (and then let this one fail)
		applicable: func(s, t *Pokemon) bool {
			switch s.NonVolatileCondition.(type) {
			case *Sleep:
				return false
			}
			return true
		},
		// TODO: Fails if the user's HP is already full, if the user is already asleep,
		// or if the user cannot fall asleep. Cannot be used if the user is affected by Heal Block.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp)
			sleep := &Sleep{counter: 2}
			t.NonVolatileCondition = sleep
			log.nonVolatileCondition(t.Name, ti, true, sleep)
		},
		Type:        PSYCHIC,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user goes to sleep for two turns. It fully restores the user's HP and heals any status problem.",
	},
	{
		Name:         "Retaliate",
		functionCode: "085",
		Power:        70,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user gets revenge for a fainted ally. If an ally fainted in the previous turn, this attack's damage increases.",
	},
	{
		Name:         "Return",
		functionCode: "089",
		Power:        1,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A full-power attack that grows more powerful the more the user likes its Trainer.",
	},
	{
		Name:         "Revenge",
		functionCode: "081",
		Power:        60,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     -4,
		Flags:        "abef",
		Description:  "An attack move that inflicts double the damage if the user has been hurt by the opponent in the same turn.",
	},
	{
		Name:         "Reversal",
		functionCode: "098",
		Power:        1,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "An all-out attack that becomes more powerful the less HP the user has.",
	},
	{
		Name:         "Roar of Time",
		functionCode: "0C2",
		Power:        150,
		Type:         DRAGON,
		Category:     special,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user blasts the target with power that distorts even time. The user must rest on the next turn.",
	},
	{
		Name:         "Roar",
		functionCode: "0EB",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     -6,
		Flags:        "bcek",
		Description:  "The target is scared off and replaced by another Pokémon in its party. In the wild, the battle ends.",
	},
	{
		Name:         "Rock Blast",
		functionCode: "0C0",
		Power:        25,
		Type:         ROCK,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user hurls hard rocks at the target. Two to five rocks are launched in quick succession.",
	},
	{
		Name: "Rock Climb",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           90,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        85,
		PP:              20,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user attacks the target by smashing into it with incredible force. It may also confuse the target.",
	},
	{
		Name: "Rock Polish",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: +2})
		},
		Type:        ROCK,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user polishes its body to reduce drag. It can sharply raise the Speed stat.",
	},
	{
		Name:            "Rock Slide",
		functionCode:    "00F",
		Power:           75,
		Type:            ROCK,
		Category:        physical,
		Accuracy:        90,
		PP:              10,
		AddEffectChance: 30,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "be",
		Description:     "Large boulders are hurled at the opposing team to inflict damage. It may also make the targets flinch.",
	},
	{
		Name: "Rock Smash",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Power:           40,
		Type:            FIGHTING,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user attacks with a punch that can shatter a rock. It may also lower the target's Defense stat.",
	},
	{
		Name:        "Rock Throw",
		Power:       50,
		Type:        ROCK,
		Category:    physical,
		Accuracy:    90,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user picks up and throws a small rock at the target to attack.",
	},
	{
		Name: "Rock Tomb",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Power:           50,
		Type:            ROCK,
		Category:        physical,
		Accuracy:        80,
		PP:              10,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "Boulders are hurled at the target. It also lowers the target's Speed by preventing its movement.",
	},
	{
		Name:         "Rock Wrecker",
		functionCode: "0C2",
		Power:        150,
		Type:         ROCK,
		Category:     physical,
		Accuracy:     90,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user launches a huge boulder at the target to attack. It must rest on the next turn, however.",
	},
	{
		Name:         "Role Play",
		functionCode: "065",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "The user mimics the target completely, copying the target's natural Ability.",
	},
	{
		Name:            "Rolling Kick",
		functionCode:    "00F",
		Power:           60,
		Type:            FIGHTING,
		Category:        physical,
		Accuracy:        85,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user lashes out with a quick, spinning kick. It may also make the target flinch.",
	},
	{
		Name:         "Rollout",
		functionCode: "0D3",
		Power:        30,
		Type:         ROCK,
		Category:     physical,
		Accuracy:     90,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user continually rolls into the target over five turns. It becomes stronger each time it hits.",
	},
	{
		Name:         "Roost",
		functionCode: "0D6",
		// TODO: Fails if its HP is full. Cannot be used if the user is affected by Heal Block.
		// If the user gains HP, then until the end of the same turn, its Flying type
		// is ignored for attacks used against the user (or treated as Normal if the user is pure Flying).
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			s.Heal(s.Stats.hp / 2)
		},
		Type:        FLYING,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user lands and rests its body. It restores the user's HP by up to half of its max HP.",
	},
	{
		Name:         "Round",
		functionCode: "083",
		Power:        60,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "befk",
		Description:  "The user attacks the target with a song. Others can join in the Round and make the attack do greater damage.",
	},
	{
		Name: "Sacred Fire",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           100,
		Type:            FIRE,
		Category:        physical,
		Accuracy:        95,
		PP:              5,
		AddEffectChance: 50,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "beg",
		Description:     "The target is razed with a mystical fire of great intensity. It may also leave the target with a burn.",
	},
	{
		Name:         "Sacred Sword",
		functionCode: "0A9",
		Power:        90,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks by slicing with its long horns. The target's stat changes don't affect this attack's damage.",
	},
	{
		Name:         "Safeguard",
		functionCode: "01A",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           25,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "The user creates a protective field that prevents status problems for five turns.",
	},
	{
		Name:         "Sand Tomb",
		functionCode: "0CF",
		Power:        35,
		Type:         GROUND,
		Category:     physical,
		Accuracy:     85,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user traps the target inside a harshly raging sandstorm for four to five turns.",
	},
	{
		Name: "Sand-Attack",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Type:        GROUND,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "Sand is hurled in the target's face, reducing its accuracy.",
	},
	{
		Name:         "Sandstorm",
		functionCode: "101",
		Type:         ROCK,
		Category:     statusEffect,
		PP:           10,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "A five-turn sandstorm is summoned to hurt all combatants except the Rock, Ground, and Steel types.",
	},
	{
		Name:         "Scald",
		functionCode: "00A",
		// TODO: Gen 6: For Scald only, thaws the target if the move hits and the target is frozen.
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           80,
		Type:            WATER,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "befg",
		Description:     "The user shoots boiling hot water at its target. It may also leave the target with a burn.",
	},
	{
		Name: "Scary Face",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -2})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user frightens the target with a scary face to harshly reduce its Speed stat.",
	},
	{
		Name:        "Scratch",
		Power:       40,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "Hard, pointed, and sharp claws rake the target to inflict damage.",
	},
	{
		Name: "Screech",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -2})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    85,
		PP:          40,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bcek",
		Description: "An earsplitting screech harshly reduces the target's Defense stat.",
	},
	{
		Name: "Searing Shot",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Power:           100,
		Type:            FIRE,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 30,
		Target:          allButUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "An inferno of scarlet flames torches everything around the user. It may leave the target with a burn.",
	},
	{
		Name:            "Secret Power",
		functionCode:    "0A4",
		Power:           70,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user attacks the target with a secret power. Its added effects vary depending on the user's environment.",
	},
	{
		Name:         "Secret Sword",
		functionCode: "122",
		Power:        85,
		Type:         FIGHTING,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user cuts with its long horn. The odd power contained in the horn does physical damage to the target.",
	},
	{
		Name:        "Seed Bomb",
		Power:       80,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user slams a barrage of hard-shelled seeds down on the target from above.",
	},
	{
		Name: "Seed Flare",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -2})
		},
		Power:           120,
		Type:            GRASS,
		Category:        special,
		Accuracy:        85,
		PP:              5,
		AddEffectChance: 40,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user emits a shock wave from its body to attack its target. It may harshly lower the target's Sp. Def.",
	},
	{
		Name: "Seismic Toss",
		damageFunction: func(source, target *Pokemon) (int, float64, float64) {
			return fixedDamage(source.Level, FIGHTING, target)
		},
		Power:       1,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is thrown using the power of gravity. It inflicts damage equal to the user's level.",
	},
	{
		Name:         "Selfdestruct",
		functionCode: "0E0",
		Power:        200,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user attacks everything around it by causing an explosion. The user faints upon using this move.",
	},
	{
		Name: "Shadow Ball",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spdefense: -1})
		},
		Power:           80,
		Type:            GHOST,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user hurls a shadowy blob at the target. It may also lower the target's Sp. Def stat.",
	},
	{
		Name:        "Shadow Claw",
		Power:       70,
		Type:        GHOST,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The user slashes with a sharp claw made from shadows. Critical hits land more easily.",
	},
	{
		Name:         "Shadow Force",
		functionCode: "0CD",
		Power:        120,
		Type:         GHOST,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "aef",
		Description:  "The user disappears, then strikes the target on the second turn. It hits even if the target protects itself.",
	},
	{
		Name:        "Shadow Punch",
		Power:       60,
		Type:        GHOST,
		Category:    physical,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefj",
		Description: "The user throws a punch from the shadows. The punch lands without fail.",
	},
	{
		Name:        "Shadow Sneak",
		Power:       40,
		Type:        GHOST,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "abef",
		Description: "The user extends its shadow and attacks the target from behind. This move always goes first.",
	},
	{
		Name: "Sharpen",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user reduces its polygon count to make itself more jagged, raising the Attack stat.",
	},
	{
		Name:         "Sheer Cold",
		functionCode: "070",
		Power:        1,
		Type:         ICE,
		Category:     special,
		Accuracy:     30,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The target is attacked with a blast of absolute-zero cold. The target instantly faints if it hits.",
	},
	{
		Name: "Shell Smash",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +2, defense: -1, spattack: +2, spdefense: -1, speed: +2})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          15,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user breaks its shell, lowering its Defense and Sp. Def stats but sharply raising Attack, Sp. Atk, and Speed stats.",
	},
	{
		Name: "Shift Gear",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, speed: +2})
		},
		Type:        STEEL,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user rotates its gears, raising its Attack and sharply raising its Speed.",
	},
	{
		Name:        "Shock Wave",
		Power:       60,
		Type:        ELECTRIC,
		Category:    special,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The user strikes the target with a quick jolt of electricity. This attack cannot be evaded.",
	},
	{
		Name: "Signal Beam",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           75,
		Type:            BUG,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks with a sinister beam of light. It may also confuse the target.",
	},
	{
		Name: "Silver Wind",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: +1, defense: +1, speed: +1, spattack: +1, spdefense: +1})
		},
		Power:           60,
		Type:            BUG,
		Category:        special,
		Accuracy:        100,
		PP:              5,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The target is attacked with powdery scales blown by wind. It may also raise all the user's stats.",
	},
	{
		Name:         "Simple Beam",
		functionCode: "063",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     10,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user's mysterious psychic wave changes the target's Ability to Simple.",
	},
	{
		Name: "Sing",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    55,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bcek",
		Description: "A soothing lullaby is sung in a calming voice that puts the target into a deep slumber.",
	},
	{
		Name:         "Sketch",
		functionCode: "05D",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           1,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "It enables the user to permanently learn the move last used by the target. Once used, Sketch disappears.",
	},
	{
		Name:         "Skill Swap",
		functionCode: "067",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user employs its psychic power to exchange Abilities with the target.",
	},
	{
		Name:            "Skull Bash",
		functionCode:    "0C8",
		Power:           100,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user tucks in its head to raise its Defense in the first turn, then rams the target on the next turn.",
	},
	{
		Name:            "Sky Attack",
		functionCode:    "0C7",
		Power:           140,
		Type:            FLYING,
		Category:        physical,
		Accuracy:        90,
		PP:              5,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "befh",
		Description:     "A second-turn attack move where critical hits land more easily. It may also make the target flinch.",
	},
	{
		Name:         "Sky Drop",
		functionCode: "0CE",
		Power:        60,
		Type:         FLYING,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefl",
		Description:  "The user takes the target into the sky, then drops it during the next turn. The target cannot attack while in the sky.",
	},
	{
		Name:         "Sky Uppercut",
		functionCode: "11B",
		Power:        85,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     90,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abefj",
		Description:  "The user attacks the target with an uppercut thrown skyward with force.",
	},
	{
		Name:         "Slack Off",
		functionCode: "0D5",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp/2)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user slacks off, restoring its own HP by up to half of its maximum HP.",
	},
	{
		Name:        "Slam",
		Power:       80,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    75,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is slammed with a long tail, vines, etc., to inflict damage.",
	},
	{
		Name:        "Slash",
		Power:       70,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abefh",
		Description: "The target is attacked with a slash of claws or blades. Critical hits land more easily.",
	},
	{
		Name: "Sleep Powder",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        GRASS,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user scatters a big cloud of sleep-inducing dust around the target.",
	},
	{
		Name:         "Sleep Talk",
		functionCode: "0B4",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       none,
		Priority:     0,
		Flags:        "",
		Description:  "While it is asleep, the user randomly uses one of the moves it knows.",
	},
	{
		Name: "Sludge Bomb",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           90,
		Type:            POISON,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "Unsanitary sludge is hurled at the target. It may also poison the target.",
	},
	{
		Name: "Sludge Wave",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           95,
		Type:            POISON,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 10,
		Target:          allButUser,
		Priority:        0,
		Flags:           "be",
		Description:     "It swamps the area around the user with a giant sludge wave. It may also poison those hit.",
	},
	{
		Name: "Sludge",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           65,
		Type:            POISON,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "Unsanitary sludge is hurled at the target. It may also poison the target.",
	},
	{
		Name:            "Smack Down",
		functionCode:    "11C",
		Power:           50,
		Type:            ROCK,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user throws a stone or projectile to attack an opponent. A flying Pokémon will fall to the ground when it's hit.",
	},
	{
		Name:         "Smelling Salts",
		functionCode: "07C",
		Power:        60,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "This attack inflicts double damage on a target with paralysis. It also cures the target's paralysis, however.",
	},
	{
		Name: "Smog",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Poison{})
		},
		Power:           20,
		Type:            POISON,
		Category:        special,
		Accuracy:        70,
		PP:              20,
		AddEffectChance: 40,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The target is attacked with a discharge of filthy gases. It may also poison the target.",
	},
	{
		Name: "Smokescreen",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeAccuracy(log, t, ti, -1)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user releases an obscuring cloud of smoke or ink. It reduces the target's accuracy.",
	},
	{
		Name: "Snarl",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: -1})
		},
		Power:       55,
		Type:        DARK,
		Category:    physical,
		Accuracy:    95,
		PP:          15,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "befk",
		Description: "The user yells as if it is ranting about something, making the target's Sp. Atk stat decrease.",
	},
	{
		Name:         "Snatch",
		functionCode: "0B2",
		Type:         DARK,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     4,
		Flags:        "",
		Description:  "The user steals the effects of any healing or stat-changing move the opponent attempts to use.",
	},
	{
		Name:            "Snore",
		functionCode:    "011",
		Power:           40,
		Type:            NORMAL,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "befk",
		Description:     "An attack that can be used only if the user is asleep. The harsh noise may also make the target flinch.",
	},
	{
		Name:         "Soak",
		functionCode: "061",
		Type:         WATER,
		Category:     statusEffect,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user shoots a torrent of water at the target and changes the target's type to Water.",
	},
	{
		Name:         "Soft-boiled",
		functionCode: "0D5",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			healWithFail(log, s, si, s.Stats.hp/2)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          10,
		Target:      user,
		Priority:    0,
		Flags:       "di",
		Description: "The user restores its own HP by up to half of its maximum HP. May also be used in the field to heal HP.",
	},
	{
		Name:         "Solar Beam",
		functionCode: "0C4",
		Power:        120,
		Type:         GRASS,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A two-turn attack. The user gathers light, then blasts a bundled beam on the second turn.",
	},
	{
		Name: "Sonic Boom",
		damageFunction: func(source, target *Pokemon) (int, float64, float64) {
			return fixedDamage(20, NORMAL, target)
		},
		Power:       1,
		Type:        NORMAL,
		Category:    special,
		Accuracy:    90,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The target is hit with a destructive shock wave that always inflicts 20 HP damage.",
	},
	{
		Name:        "Spacial Rend",
		Power:       100,
		Type:        DRAGON,
		Category:    special,
		Accuracy:    95,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "befh",
		Description: "The user tears the target along with the space around it. Critical hits land more easily.",
	},
	{
		Name: "Spark",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           65,
		Type:            ELECTRIC,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user throws an electrically charged tackle at the target. It may also leave the target with paralysis.",
	},
	{
		Name:         "Spider Web",
		functionCode: "0EF",
		Type:         BUG,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user ensnares the target with thin, gooey silk so it can't flee from battle.",
	},
	{
		Name:         "Spike Cannon",
		functionCode: "0C0",
		Power:        20,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Sharp spikes are shot at the target in rapid succession. They hit two to five times in a row.",
	},
	{
		Name:         "Spikes",
		functionCode: "103",
		Type:         GROUND,
		Category:     statusEffect,
		PP:           20,
		Target:       opposingSide,
		Priority:     0,
		Flags:        "c",
		Description:  "The user lays a trap of spikes at the opposing team's feet. The trap hurts Pokémon that switch into battle.",
	},
	{
		Name:         "Spit Up",
		functionCode: "113",
		Power:        1,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bf",
		Description:  "The power stored using the move Stockpile is released at once in an attack. The more power is stored, the greater the damage.",
	},
	{
		Name:         "Spite",
		functionCode: "10E",
		Type:         GHOST,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user unleashes its grudge on the move last used by the target by cutting 4,PP from it.",
	},
	{
		Name: "Splash",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			log.f("But nothing happened!")
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "l",
		Description: "The user just flops and splashes around to no effect at all…",
	},
	{
		Name: "Spore",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewSleep())
		},
		Type:        GRASS,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user scatters bursts of spores that induce sleep.",
	},
	{
		Name:         "Stealth Rock",
		functionCode: "105",
		Type:         ROCK,
		Category:     statusEffect,
		PP:           20,
		Target:       opposingSide,
		Priority:     0,
		Flags:        "c",
		Description:  "The user lays a trap of levitating stones around the opponent's team. The trap hurts opponents that switch into battle.",
	},
	{
		Name:            "Steamroller",
		functionCode:    "010",
		Power:           65,
		Type:            BUG,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The user crushes its targets by rolling over them with its rolled-up body. This attack may make the target flinch.",
	},
	{
		Name: "Steel Wing",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{defense: +1})
		},
		Power:           70,
		Type:            STEEL,
		Category:        physical,
		Accuracy:        90,
		PP:              25,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The target is hit with wings of steel. It may also raise the user's Defense stat.",
	},
	{
		Name:         "Stockpile",
		functionCode: "112",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user charges up power and raises both its Defense and Sp. Def. The move can be used three times.",
	},
	{
		Name:            "Stomp",
		functionCode:    "010",
		Power:           65,
		Type:            NORMAL,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abe",
		Description:     "The target is stomped with a big foot. It may also make the target flinch.",
	},
	{
		Name:        "Stone Edge",
		Power:       100,
		Type:        ROCK,
		Category:    physical,
		Accuracy:    80,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "befh",
		Description: "The user stabs the foe with sharpened stones from below. It has a high critical-hit ratio.",
	},
	{
		Name:         "Stored Power",
		functionCode: "08E",
		Power:        1,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user attacks the target with stored power. The more the user's stats are raised, the greater the damage.",
	},
	{
		Name:         "Storm Throw",
		functionCode: "0A0",
		Power:        40,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user strikes the target with a fierce blow. This attack always results in a critical hit.",
	},
	{
		Name:        "Strength",
		Power:       80,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is slugged with a punch thrown at maximum power. It can also be used to move heavy boulders.",
	},
	{
		Name: "String Shot",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{speed: -1})
		},
		Type:        BUG,
		Category:    statusEffect,
		Accuracy:    95,
		PP:          40,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "The targets are bound with silk blown from the user's mouth. This silk reduces the targets' Speed stat.",
	},
	{
		Name: "Struggle Bug",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: -1})
		},
		Power:       30,
		Type:        BUG,
		Category:    special,
		Accuracy:    100,
		PP:          20,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "be",
		Description: "While resisting, the user attacks the opposing Pokémon. The targets' Sp. Atk stat is reduced.",
	},
	{
		Name:         "Struggle",
		functionCode: "002",
		Power:        50,
		Type:         NORMAL,
		Category:     physical,
		PP:           1,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abf",
		Description:  "An attack that is used in desperation only if the user has no PP. It also hurts the user slightly.",
	},
	{
		Name: "Stun Spore",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Type:        GRASS,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          30,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user scatters a cloud of paralyzing powder. It may leave the target with paralysis.",
	},
	{
		Name: "Submission",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 4)
		},
		Power:       80,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    80,
		PP:          25,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user grabs the target and recklessly dives for the ground. It also hurts the user slightly.",
	},
	{
		Name:         "Substitute",
		functionCode: "10C",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "d",
		Description:  "The user makes a copy of itself using some of its HP. The copy serves as the user's decoy.",
	},
	{
		Name:         "Sucker Punch",
		functionCode: "116",
		Power:        80,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     1,
		Flags:        "abef",
		Description:  "This move enables the user to attack first. It fails if the target is not readying an attack, however.",
	},
	{
		Name:         "Sunny Day",
		functionCode: "0FF",
		Type:         FIRE,
		Category:     statusEffect,
		PP:           5,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user intensifies the sun for five turns, Powering up Fire-type moves.",
	},
	{
		Name: "Super Fang",
		damageFunction: func(source, target *Pokemon) (int, float64, float64) {
			return fixedDamage(target.currentHP/2, NORMAL, target)
		},
		Power:       1,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    90,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abe",
		Description: "The user chomps hard on the target with its sharp front fangs. It cuts the target's HP to half.",
	},
	{
		Name: "Superpower",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{attack: -1, defense: -1})
		},
		Power:       120,
		Type:        FIGHTING,
		Category:    physical,
		Accuracy:    100,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abe",
		Description: "The user attacks the target with great power. However, it also lowers the user's Attack and Defense.",
	},
	{
		Name: "Supersonic",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    55,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bcek",
		Description: "The user generates odd sound waves from its body. It may confuse the target.",
	},
	{
		Name:         "Surf",
		functionCode: "075",
		Power:        95,
		Type:         WATER,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "It swamps the area around the user with a giant wave. It can also be used for crossing water.",
	},
	{
		Name: "Swagger",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(log, t, ti, Stats{attack: +2})
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    90,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user enrages and confuses the target. However, it also sharply raises the target's Attack stat.",
	},
	{
		Name:         "Swallow",
		functionCode: "114",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The power stored using the move Stockpile is absorbed by the user to heal its HP. Storing more power heals more HP.",
	},
	{
		Name: "Sweet Kiss",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user kisses the target with a sweet, angelic cuteness that causes confusion.",
	},
	{
		Name: "Sweet Scent",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			// TODO: Gen 6: Decreases the target's evasiveness by 2 stages instead.
			changeEvasion(log, t, ti, -1)
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "A sweet scent that lowers the opposing team's evasiveness. It also lures wild Pokémon if used in grass, etc.",
	},
	{
		Name:        "Swift",
		Power:       60,
		Type:        NORMAL,
		Category:    special,
		PP:          20,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bef",
		Description: "Star-shaped rays are shot at the opposing team. This attack never misses.",
	},
	{
		Name:         "Switcheroo",
		functionCode: "0F2",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user trades held items with the target faster than the eye can follow.",
	},
	{
		Name: "Swords Dance",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +2})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "A frenetic dance to uplift the fighting spirit. It sharply raises the user's Attack stat.",
	},
	{
		Name:         "Synchronoise",
		functionCode: "123",
		Power:        70,
		Type:         PSYCHIC,
		Category:     special,
		Accuracy:     100,
		PP:           15,
		Target:       allButUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Using an odd shock wave, the user inflicts damage on any Pokémon of the same type in the area around it.",
	},
	{
		Name:         "Synthesis",
		functionCode: "0D8",
		Type:         GRASS,
		Category:     statusEffect,
		PP:           5,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "The user restores its own HP. The amount of HP regained varies with the weather.",
	},
	{
		Name:        "Tackle",
		Power:       50,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "A physical attack in which the user charges and slams into the target with its whole body.",
	},
	{
		Name: "Tail Glow",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{spattack: +3})
		},
		Type:        BUG,
		Category:    statusEffect,
		PP:          20,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user stares at flashing lights to focus its mind, drastically raising its Sp. Atk stat.",
	},
	{
		Name:         "Tail Slap",
		functionCode: "0C0",
		Power:        25,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     85,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user attacks by striking the target with its hard tail. It hits the Pokémon two to five times in a row.",
	},
	{
		Name: "Tail Whip",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: -1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          30,
		Target:      allOpposing,
		Priority:    0,
		Flags:       "bce",
		Description: "The user wags its tail cutely, making opposing Pokémon less wary and lowering their Defense stat.",
	},
	{
		Name:         "Tailwind",
		functionCode: "05B",
		Type:         FLYING,
		Category:     statusEffect,
		PP:           30,
		Target:       bothSides,
		Priority:     0,
		Flags:        "d",
		Description:  "The user whips up a turbulent whirlwind that ups the Speed of all party Pokémon for four turns.",
	},
	{
		Name: "Take Down",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 4)
		},
		Power:       90,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    85,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "A reckless, full-body charge attack for slamming into the target. It also damages the user a little.",
	},
	{
		Name:         "Taunt",
		functionCode: "0BA",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The target is taunted into a rage that allows it to use only attack moves for three turns.",
	},
	{
		Name:         "Techno Blast",
		functionCode: "09F",
		Power:        85,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user fires a beam of light at its target. The type changes depending on the Drive the user holds.",
	},
	{
		Name: "Teeter Dance",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      allButUser,
		Priority:    0,
		Flags:       "be",
		Description: "The user performs a wobbly dance that confuses the Pokémon around it.",
	},
	{
		Name:         "Telekinesis",
		functionCode: "11A",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bcel",
		Description:  "The user makes the target float with its psychic power. The target is easier to hit for three turns.",
	},
	{
		Name:         "Teleport",
		functionCode: "0EA",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           20,
		Target:       user,
		Priority:     0,
		Flags:        "",
		Description:  "Use it to flee from any wild Pokémon. It can also warp to the last Pokémon Center visited.",
	},
	{
		Name:         "Thief",
		functionCode: "0F1",
		Power:        40,
		Type:         DARK,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abe",
		Description:  "The user attacks and steals the target's held item simultaneously. It can't steal if the user holds an item.",
	},
	{
		Name:         "Thrash",
		functionCode: "0D2",
		Power:        120,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleRandomOpposing,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user rampages and attacks for two to three turns. It then becomes confused, however.",
	},
	{
		Name:            "Thunder Fang",
		functionCode:    "009",
		Power:           65,
		Type:            ELECTRIC,
		Category:        physical,
		Accuracy:        95,
		PP:              15,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user bites with electrified fangs. It may also make the target flinch or leave it with paralysis.",
	},
	{
		Name: "Thunder Wave",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Type:        ELECTRIC,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "A weak electric charge is launched at the target. It causes paralysis if it hits.",
	},
	{
		Name:         "Thunder",
		functionCode: "008",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           120,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        70,
		PP:              10,
		AddEffectChance: 30,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "A wicked thunderbolt is dropped on the target to inflict damage. It may also leave the target with paralysis.",
	},
	{
		Name: "Thunderbolt",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           95,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "A strong electric blast is loosed at the target. It may also leave the target with paralysis.",
	},
	{
		Name: "Thunder Punch",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           75,
		Type:            ELECTRIC,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abej",
		Description:     "The target is punched with an electrified fist. It may also leave the target with paralysis.",
	},
	{
		Name: "Thunder Shock",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           40,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        100,
		PP:              30,
		AddEffectChance: 10,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "A jolt of electricity is hurled at the target to inflict damage. It may also leave the target with paralysis.",
	},
	{
		Name: "Tickle",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: -1, defense: -1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		Accuracy:    100,
		PP:          20,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user tickles the target into laughing, reducing its Attack and Defense stats.",
	},
	{
		Name:         "Torment",
		functionCode: "0B7",
		Type:         DARK,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user torments and enrages the target, making it incapable of using the same move twice in a row.",
	},
	{
		Name:         "Toxic Spikes",
		functionCode: "104",
		Type:         POISON,
		Category:     statusEffect,
		PP:           20,
		Target:       opposingSide,
		Priority:     0,
		Flags:        "c",
		Description:  "The user lays a trap of poison spikes at the opponent's feet. They poison opponents that switch into battle.",
	},
	{
		Name: "Toxic",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, NewBadPoison())
		},
		canNotMiss: func(s *Pokemon) bool {
			return typeContains(POISON, s.getSpecies().Types)
		},
		Type:        POISON,
		Category:    statusEffect,
		Accuracy:    90,
		PP:          10,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "A move that leaves the target badly poisoned. Its poison damage worsens every turn.",
	},
	{
		Name:         "Transform",
		functionCode: "069",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "",
		Description:  "The user transforms into a copy of the target right down to having the same move set.",
	},
	{
		Name: "Tri Attack",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			switch random.Intn(3) {
			case 0:
				inflictNonVolatileCondition(log, t, ti, Burn{})
			case 1:
				inflictNonVolatileCondition(log, t, ti, Freeze{})
			case 2:
				inflictNonVolatileCondition(log, t, ti, Paralysis{})
			}
		},
		Power:           80,
		Type:            NORMAL,
		Category:        special,
		Accuracy:        100,
		PP:              10,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user strikes with a simultaneous three-beam attack. May also burn, freeze, or leave the target with paralysis.",
	},
	{
		Name:         "Trick Room",
		functionCode: "11F",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           5,
		Target:       usersSide,
		Priority:     -7,
		Flags:        "e",
		Description:  "The user creates a bizarre area in which slower Pokémon get to move first for five turns.",
	},
	{
		Name:         "Trick",
		functionCode: "0F2",
		Type:         PSYCHIC,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "be",
		Description:  "The user catches the target off guard and swaps its held item with its own.",
	},
	{
		Name:         "Triple Kick",
		functionCode: "0BF",
		Power:        10,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     90,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A consecutive three-kick attack that becomes more powerful with each successive hit.",
	},
	{
		Name:         "Trump Card",
		functionCode: "097",
		Power:        1,
		Type:         NORMAL,
		Category:     special,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The fewer PP this move has, the greater its attack power.",
	},
	{
		Name:            "Twineedle",
		functionCode:    "0BE",
		Power:           25,
		Type:            BUG,
		Category:        physical,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user damages the target twice in succession by jabbing it with two spikes. It may also poison the target.",
	},
	{
		Name:            "Twister",
		functionCode:    "078",
		Power:           40,
		Type:            DRAGON,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 20,
		Target:          allOpposing,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user whips up a vicious tornado to tear at the opposing team. It may also make targets flinch.",
	},
	{
		Name:         "U-turn",
		functionCode: "0EE",
		Power:        70,
		Type:         BUG,
		Category:     physical,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "After making its attack, the user rushes back to switch places with a party Pokémon in waiting.",
	},
	{
		Name:         "Uproar",
		functionCode: "0D1",
		Power:        90,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleRandomOpposing,
		Priority:     0,
		Flags:        "befk",
		Description:  "The user attacks in an uproar for three turns. Over that time, no one can fall asleep.",
	},
	{
		Name:         "V-create",
		functionCode: "03D",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, s, si, Stats{defense: -1, spdefense: -1, speed: -1})
		},
		Power:       180,
		Type:        FIRE,
		Category:    physical,
		Accuracy:    95,
		PP:          5,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "With a hot flame on its forehead, the user hurls itself at its target. It lowers the user's Defense, Sp. Def, and Speed stats.",
	},
	{
		Name:        "Vacuum Wave",
		Power:       40,
		Type:        FIGHTING,
		Category:    special,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    1,
		Flags:       "bef",
		Description: "The user whirls its fists to send a wave of pure vacuum at the target. This move always goes first.",
	},
	{
		Name:         "Venoshock",
		functionCode: "07B",
		Power:        65,
		Type:         POISON,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "The user drenches the target in a special poisonous liquid. Its power is doubled if the target is poisoned.",
	},
	{
		Name:        "Vice Grip",
		Power:       55,
		Type:        NORMAL,
		Category:    physical,
		Accuracy:    100,
		PP:          30,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is gripped and squeezed from both sides to inflict damage.",
	},
	{
		Name:        "Vine Whip",
		Power:       35,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is struck with slender, whiplike vines to inflict damage.",
	},
	{
		Name:        "Vital Throw",
		Power:       70,
		Type:        FIGHTING,
		Category:    physical,
		PP:          10,
		Target:      singleNotUser,
		Priority:    -1,
		Flags:       "abef",
		Description: "The user attacks last. In return, this throw move is guaranteed not to miss.",
	},
	{
		Name:         "Volt Switch",
		functionCode: "0EE",
		Power:        70,
		Type:         ELECTRIC,
		Category:     special,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "After making its attack, the user rushes back to switch places with a party Pokémon in waiting.",
	},
	{
		Name: "Volt Tackle",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 3)
			// 10% chance to paralyze. not using AddEffectChance because recoil always happens
			if random.Float64() < 0.1 {
				inflictNonVolatileCondition(l, t, ti, Paralysis{})
			}
		},
		Power:       120,
		Type:        ELECTRIC,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user electrifies itself, then charges. It causes considerable damage to the user and may leave the target with paralysis.",
	},
	{
		Name:         "Wake-Up Slap",
		functionCode: "07D",
		Power:        60,
		Type:         FIGHTING,
		Category:     physical,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "This attack inflicts big damage on a sleeping target. It also wakes the target up, however.",
	},
	{
		Name:        "Water Gun",
		Power:       40,
		Type:        WATER,
		Category:    special,
		Accuracy:    100,
		PP:          25,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bef",
		Description: "The target is blasted with a forceful shot of water.",
	},
	{
		Name:         "Water Pledge",
		functionCode: "108",
		Power:        50,
		Type:         WATER,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "A column of water strikes the target. When combined with its fire equivalent, the damage increases and a rainbow appears.",
	},
	{
		Name: "Water Pulse",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictVolatileCondition(log, t, ti, NewConfusion())
		},
		Power:           60,
		Type:            WATER,
		Category:        special,
		Accuracy:        100,
		PP:              20,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "bef",
		Description:     "The user attacks the target with a pulsing blast of water. It may also confuse the target.",
	},
	{
		Name:         "Water Sport",
		functionCode: "09E",
		Type:         WATER,
		Category:     statusEffect,
		PP:           15,
		Target:       usersSide,
		Priority:     0,
		Flags:        "",
		Description:  "The user soaks itself with water. The move weakens Fire-type moves while the user is in the battle.",
	},
	{
		Name:         "Water Spout",
		functionCode: "08B",
		Power:        150,
		Type:         WATER,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       allOpposing,
		Priority:     0,
		Flags:        "be",
		Description:  "The user spouts water to damage the opposing team. The lower the user's HP, the less powerful it becomes.",
	},
	{
		Name:            "Waterfall",
		functionCode:    "00F",
		Power:           80,
		Type:            WATER,
		Category:        physical,
		Accuracy:        100,
		PP:              15,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user charges at the target and may make it flinch. It can also be used to climb a waterfall.",
	},
	{
		Name:         "Weather Ball",
		functionCode: "087",
		Power:        50,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "An attack move that varies in power and type depending on the weather.",
	},
	{
		Name:         "Whirlpool",
		functionCode: "0D0",
		Power:        35,
		Type:         WATER,
		Category:     special,
		Accuracy:     85,
		PP:           15,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bef",
		Description:  "Traps foes in a violent swirling whirlpool for four to five turns.",
	},
	{
		Name:         "Whirlwind",
		functionCode: "0EB",
		Type:         NORMAL,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           20,
		Target:       singleNotUser,
		Priority:     -6,
		Flags:        "bce",
		Description:  "The foe is blown away, to be replaced by another Pokémon in its party. In the wild, the battle ends.",
	},
	{
		Name:         "Wide Guard",
		functionCode: "0AC",
		Type:         ROCK,
		Category:     statusEffect,
		PP:           10,
		Target:       bothSides,
		Priority:     3,
		Flags:        "d",
		Description:  "The user and its allies are protected from wide-ranging attacks for one turn. If used in succession, its chances of failing rises.",
	},
	{
		Name: "Wild Charge",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 4)
		},
		Power:       90,
		Type:        ELECTRIC,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user shrouds itself in electricity and smashes into its target. It also damages the user a little.",
	},
	{
		Name: "Will-O-Wisp",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Burn{})
		},
		Type:        FIRE,
		Category:    statusEffect,
		Accuracy:    75,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "bce",
		Description: "The user shoots a sinister, bluish-white flame at the target to inflict a burn.",
	},
	{
		Name:        "Wing Attack",
		Power:       60,
		Type:        FLYING,
		Category:    physical,
		Accuracy:    100,
		PP:          35,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The target is struck with large, imposing wings spread wide to inflict damage.",
	},
	{
		Name:         "Wish",
		functionCode: "0D7",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       user,
		Priority:     0,
		Flags:        "di",
		Description:  "One turn after this move is used, the target's HP is restored by half the user's maximum HP.",
	},
	{
		Name: "Withdraw",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{defense: +1})
		},
		Type:        WATER,
		Category:    statusEffect,
		PP:          40,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user withdraws its body into its hard shell, raising its Defense stat.",
	},
	{
		Name:         "Wonder Room",
		functionCode: "124",
		Type:         PSYCHIC,
		Category:     statusEffect,
		PP:           10,
		Target:       usersSide,
		Priority:     -7,
		Flags:        "e",
		Description:  "The user creates a bizarre area in which Pokémon's Defense and Sp. Def stats are swapped for five turns.",
	},
	{
		Name: "Wood Hammer",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			recoil(l, s, si, dmg, 3)
		},
		Power:       120,
		Type:        GRASS,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user slams its rugged body into the target to attack. The user also sustains serious damage.",
	},
	{
		Name: "Work Up",
		effect: func(l *Logger, s, t *Pokemon, si, ti, dmg int) {
			changeStatStages(l, t, ti, Stats{attack: +1, spattack: +1})
		},
		Type:        NORMAL,
		Category:    statusEffect,
		PP:          30,
		Target:      user,
		Priority:    0,
		Flags:       "d",
		Description: "The user is roused, and its Attack and Sp. Atk stats increase.",
	},
	{
		Name:         "Worry Seed",
		functionCode: "064",
		Type:         GRASS,
		Category:     statusEffect,
		Accuracy:     100,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "A seed that causes worry is planted on the target. It prevents sleep by making its Ability Insomnia.",
	},
	{
		Name:         "Wrap",
		functionCode: "0CF",
		Power:        15,
		Type:         NORMAL,
		Category:     physical,
		Accuracy:     90,
		PP:           20,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "A long body or vines are used to wrap and squeeze the target for four to five turns.",
	},
	{
		Name:         "Wring Out",
		functionCode: "08C",
		Power:        1,
		Type:         NORMAL,
		Category:     special,
		Accuracy:     100,
		PP:           5,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "abef",
		Description:  "The user powerfully wrings the target. The more HP the target has, the greater this attack's power.",
	},
	{
		Name:        "X-Scissor",
		Power:       80,
		Type:        BUG,
		Category:    physical,
		Accuracy:    100,
		PP:          15,
		Target:      singleNotUser,
		Priority:    0,
		Flags:       "abef",
		Description: "The user slashes at the target by crossing its scythes or claws as if they were a pair of scissors.",
	},
	{
		Name:         "Yawn",
		functionCode: "004",
		Type:         NORMAL,
		Category:     statusEffect,
		PP:           10,
		Target:       singleNotUser,
		Priority:     0,
		Flags:        "bce",
		Description:  "The user lets loose a huge yawn that lulls the target into falling asleep on the next turn.",
	},
	{
		Name: "Zap Cannon",
		effect: func(log *Logger, s, t *Pokemon, si, ti, dmg int) {
			inflictNonVolatileCondition(log, t, ti, Paralysis{})
		},
		Power:           120,
		Type:            ELECTRIC,
		Category:        special,
		Accuracy:        50,
		PP:              5,
		AddEffectChance: 100,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "be",
		Description:     "The user fires an electric blast like a cannon to inflict damage and cause paralysis.",
	},
	{
		Name:            "Zen Headbutt",
		functionCode:    "00F",
		Power:           80,
		Type:            PSYCHIC,
		Category:        physical,
		Accuracy:        90,
		PP:              15,
		AddEffectChance: 20,
		Target:          singleNotUser,
		Priority:        0,
		Flags:           "abef",
		Description:     "The user focuses its willpower to its head and attacks the target. It may also make the target flinch.",
	},
}
