package xivapi

import (
	"fmt"
)

type SearchResultDefaultSubtype struct {
	Type SearchIndex `mapstructure:"_"`
	Name string
}

type SearchResultUrlSubtype struct {
	URL     string
	SiteURL string
}

type SearchResultIconSubtype struct {
	Icon string
}

func (i *SearchResultIconSubtype) FullIcon() string {
	return BaseURL + i.Icon
}

type SearchResultCompanion struct {
	SearchResultDefaultSubtype `mapstructure:",squash"`
	SearchResultUrlSubtype     `mapstructure:",squash"`
	SearchResultIconSubtype    `mapstructure:",squash"`
	MinionRaceName             string `mapstructure:"MinionRace.Name"`
	BehaviorName               string `mapstructure:"Behavior.Name"`
	ID                         int
	GameType                   string
}

func (c SearchResultCompanion) String() string {
	return fmt.Sprintf("%v { [%v] %v (%v) <%v> <%v> }", c.Type, c.ID, c.Name, c.BehaviorName, c.SiteURL, c.FullIcon())
}
