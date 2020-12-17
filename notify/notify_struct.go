package notify

// ################
// # DAILY NOTIFY #
// ################

type DailyEntry struct {
	Greeting string
	Time string
	Content interface{}
}

type Daily7h struct {
	NetherworldGate bool
	DemonEncounter string
}

type Daily8h struct {
	GuildRaid string
}

type Daily12h struct {
	Maintenance bool
	SoulZone10 []string
	SoulZone11 string
	TotemZone string
}

type Daily18h struct {
	DemonEncounter string
	NetherworldGate bool
}

type Daily21h struct {
	GuildFeast bool
}
