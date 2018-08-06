package xivapi

import (
	"fmt"
)

// SearchResultDefaultSubtype holds the common info of all search result entries
type SearchResultDefaultSubtype struct {
	Type SearchIndex `mapstructure:"_"`
	Name string
}

// SearchResultURLSubtype holds all the common URL fields for a search result
type SearchResultURLSubtype struct {
	URL     string
	SiteURL string
}

// SearchResultIconSubtype holds all the common icon fields for a search result
type SearchResultIconSubtype struct {
	Icon string
}

// FullIcon returns the full url for an icon
func (i *SearchResultIconSubtype) FullIcon() string {
	return BaseURL + i.Icon
}

// SearchResultCompanion holds all fields for a companion search result
type SearchResultCompanion struct {
	SearchResultDefaultSubtype `mapstructure:",squash"`
	SearchResultURLSubtype     `mapstructure:",squash"`
	SearchResultIconSubtype    `mapstructure:",squash"`
	MinionRaceName             string `mapstructure:"MinionRace.Name"`
	BehaviorName               string `mapstructure:"Behavior.Name"`
	ID                         int
	GameType                   string
}

func (c SearchResultCompanion) String() string {
	return fmt.Sprintf("%v { [%v] %v (%v) <%v> <%v> }", c.Type, c.ID, c.Name, c.BehaviorName, c.SiteURL, c.FullIcon())
}
