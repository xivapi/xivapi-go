package xivapi

import (
	"encoding/json"
)

// SearchResultCommon holds the common info of all search result entries
type SearchResultCommon struct {
	Type SearchIndex `json:"_"`
	Name string
	ID   uint
	Icon string
}

// FullIcon returns the full url for an icon
func (c *SearchResultCommon) FullIcon() string {
	return BaseURL + c.Icon
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

// SearchResultBNPCName holds the bnnpcname details
type SearchResultBNPCName struct {
	*SearchResultCommon
	URL string `json:"Url"`
}

// GetBNPCName returns the BNPCName details for a result
func (e *SearchResultEntry) GetBNPCName() (*SearchResultBNPCName, error) {
	if e.Type != IndexBNPCName {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultBNPCName)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetEnemy is an alias for GetBNPCName
func (e *SearchResultEntry) GetEnemy() (*SearchResultBNPCName, error) {
	return e.GetBNPCName()
}

// SearchResultCompanion holds all fields for a companion search result
type SearchResultCompanion struct {
	*SearchResultCommon
	URL            string `json:"Url"`
	BehaviorID     uint   `json:"Behaviour.ID"`
	BehaviorName   string `json:"Behaviour.Name"`
	GamePatchID    int    `json:"GamePatch.ID"`
	MinionRaceID   uint   `json:"MinionRace.ID"`
	MinionRaceName string `json:"MinionRace.Name"`
	Cost           int
	HP             int
	SkillAngle     uint
	SkillCost      uint
}

// GetCompanion returns the Companion details for a result
func (e *SearchResultEntry) GetCompanion() (*SearchResultCompanion, error) {
	if e.Type != IndexCompanion {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultCompanion)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetMinion is an alias for GetCompanion
func (e *SearchResultEntry) GetMinion() (*SearchResultCompanion, error) {
	return e.GetCompanion()
}

// SearchResultENPCResident holds all fields for a enpcresident search result
type SearchResultENPCResident struct {
	*SearchResultCommon
	URL   string `json:"Url"`
	Title string
}

// GetENPCResident returns the enpcresident details for a result
func (e *SearchResultEntry) GetENPCResident() (*SearchResultENPCResident, error) {
	if e.Type != IndexENPCResident {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultENPCResident)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetNPC is an alias for GetENPCResident
func (e *SearchResultEntry) GetNPC() (*SearchResultENPCResident, error) {
	return e.GetENPCResident()
}

// SearchResultEmote holds all fields for a emote search result
type SearchResultEmote struct {
	*SearchResultCommon
	URL                string `json:"Url"`
	GamePatchID        int    `json:"GamePatch.ID"`
	EmoteCategoryID    uint   `json:"EmoteCategory.ID"`
	EmoteCategoryName  string `json:"EmoteCategory.Name"`
	TextCommandID      uint   `json:"TextCommand.ID"`
	TextCommandCommand string `json:"TextCommand.Command"`
}

// GetEmote returns the emote details for a result
func (e *SearchResultEntry) GetEmote() (*SearchResultEmote, error) {
	if e.Type != IndexEmote {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultEmote)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultFate holds all fields for a Fate search result
type SearchResultFate struct {
	*SearchResultCommon
	URL              string `json:"Url"`
	GamePatchID      int    `json:"GamePatch.ID"`
	Description      string
	ClassJobLevel    int
	ClassJobLevelMax int
}

// GetFate returns the Fate details for a result
func (e *SearchResultEntry) GetFate() (*SearchResultFate, error) {
	if e.Type != IndexFate {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultFate)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}
