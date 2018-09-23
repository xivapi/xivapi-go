package xivapi

import (
	"encoding/json"
	"fmt"
)

// SearchResultCommon holds the common info of all search result entries
type SearchResultCommon struct {
	Type SearchIndex `json:"_"`
	Name string
	ID   uint
	Icon string
}

// SearchResultAchievement holds the achievement details
type SearchResultAchievement struct {
	*SearchResultCommon
	Description                 string
	Order                       int
	Points                      int
	Type                        int
	URL                         string `json:"Url"`
	GamePatchID                 int    `json:"GamePatch.ID"`
	AchievementCategoryID       uint   `json:"AchievementCategory.ID"`
	AchievementCategoryName     string `json:"AchievementCategory.Name"`
	AchievementCategoryKindID   uint   `json:"AchievementCategory.AchievementKind.ID"`
	AchievementCategoryKindName string `json:"AchievementCategory.AchievementKind.Name"`
}

// GetAchievement returns the achievement details for a result
func (e *SearchResultEntry) GetAchievement() (*SearchResultAchievement, error) {
	if e.Type != IndexAchievement {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultAchievement)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultAction holds the Action details
type SearchResultAction struct {
	*SearchResultCommon
	URL                  string `json:"Url"`
	GamePatchID          int    `json:"GamePatch.ID"`
	ActionCategoryID     uint   `json:"ActionCategory.ID"`
	ActionCategoryName   string `json:"ActionCategory.Name"`
	ClassJobID           uint   `json:"ClassJob.ID"`
	ClassJobName         string `json:"ClassJob.Name"`
	ClassJobCategoryID   uint   `json:"ClassJobCategory.ID"`
	ClassJobCategoryName string `json:"ClassJobCategory.Name"`
	ClassJobLevel        uint
	CanTargetDead        JSONBoolNumber
	CanTargetFriendly    JSONBoolNumber
	CanTargetHostile     JSONBoolNumber
	CanTargetParty       JSONBoolNumber
	CanTargetSelf        JSONBoolNumber
	IsPvP                JSONBoolNumber
	IsRoleAction         JSONBoolNumber
	PreservesCombo       JSONBoolNumber
	Range                uint
	Recast100ms          uint
	TargetArea           uint
}

// GetAction returns the Action details for a result
func (e *SearchResultEntry) GetAction() (*SearchResultAction, error) {
	if e.Type != IndexAction {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultAction)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
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
	SearchResultCommon      `mapstructure:",squash"`
	SearchResultURLSubtype  `mapstructure:",squash"`
	SearchResultIconSubtype `mapstructure:",squash"`
	MinionRaceName          string `mapstructure:"MinionRace.Name"`
	BehaviorName            string `mapstructure:"Behavior.Name"`
	ID                      int
	GameType                string
}

func (c SearchResultCompanion) String() string {
	return fmt.Sprintf("%v { [%v] %v (%v) <%v> <%v> }", c.Type, c.ID, c.Name, c.BehaviorName, c.SiteURL, c.FullIcon())
}
