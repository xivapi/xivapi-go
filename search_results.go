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

// SearchResultInstanceContent holds all fields for a InstanceContent search result
type SearchResultInstanceContent struct {
	*SearchResultCommon
	URL                                         string `json:"Url"`
	GamePatchID                                 int    `json:"GamePatch.ID"`
	BossCurrencyA0                              int
	BossCurrencyA1                              int
	BossCurrencyA2                              int
	BossCurrencyA3                              int
	BossCurrencyA4                              int
	BossCurrencyB0                              int
	BossCurrencyB1                              int
	BossCurrencyB2                              int
	BossCurrencyB3                              int
	BossCurrencyB4                              int
	BossCurrencyC0                              int
	BossCurrencyC1                              int
	BossCurrencyC2                              int
	BossCurrencyC3                              int
	BossCurrencyC4                              int
	FinalBossCurrencyA                          int
	FinalBossCurrencyB                          int
	FinalBossCurrencyC                          int
	NewPlayerBonusA                             int
	NewPlayerBonusB                             int
	PartyCondition                              int // TODO: correct type?
	SortKey                                     int // TODO: correct type?
	TimeLimitMin                                int
	WeekRestriction                             JSONBoolNumber
	InstanceContentTypeID                       uint   `json:"InstanceContentType.ID"`
	ContentTypeID                               uint   `json:"ContentType.ID"`
	ContentTypeName                             string `json:"ContentType.Name"`
	ContentMemberTypeHealersPerParty            int    `json:"ContentMemberType.HealersPerParty"`
	ContentMemberTypeMeleesPerParty             int    `json:"ContentMemberType.MeleesPerParty"`
	ContentMemberTypeRangedPerParty             int    `json:"ContentMemberType.RangedPerParty"`
	ContentMemberTypeTanksPerParty              int    `json:"ContentMemberType.TanksPerParty"`
	ContentFinderConditionClassJobLevelRequired int    `json:"ContentFinderCondition.ClassJobLevelRequired"`
	ContentFinderConditionClassJobLevelSync     int    `json:"ContentFinderCondition.ClassJobLevelSync"`
	ContentFinderConditionItemLevelRequired     int    `json:"ContentFinderCondition.ItemLevelRequired"`
	ContentFinderConditionItemLevelSync         int    `json:"ContentFinderCondition.ItemLevelSync"`
}

// GetInstanceContent returns the InstanceContent details for a result
func (e *SearchResultEntry) GetInstanceContent() (*SearchResultInstanceContent, error) {
	if e.Type != IndexInstanceContent {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultInstanceContent)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultItem holds all fields for a Item search result
type SearchResultItem struct {
	*SearchResultCommon
	URL                        string `json:"Url"`
	GamePatchID                int    `json:"GamePatch.ID"`
	BaseParam0ID               uint   `json:"BaseParam0.ID"`
	BaseParam1ID               uint   `json:"BaseParam1.ID"`
	BaseParam2ID               uint   `json:"BaseParam2.ID"`
	BaseParam3ID               uint   `json:"BaseParam3.ID"`
	BaseParam4ID               uint   `json:"BaseParam4.ID"`
	BaseParam5ID               uint   `json:"BaseParam5.ID"`
	BaseParamSpecial0ID        uint   `json:"BaseParamSpecial0.ID"`
	BaseParamSpecial1ID        uint   `json:"BaseParamSpecial1.ID"`
	BaseParamSpecial2ID        uint   `json:"BaseParamSpecial2.ID"`
	BaseParamSpecial3ID        uint   `json:"BaseParamSpecial3.ID"`
	BaseParamSpecial4ID        uint   `json:"BaseParamSpecial4.ID"`
	BaseParamSpecial5ID        uint   `json:"BaseParamSpecial5.ID"`
	BaseParamValue0            int
	BaseParamValue1            int
	BaseParamValue2            int
	BaseParamValue3            int
	BaseParamValue4            int
	BaseParamValue5            int
	BaseParamValueSpecial0     int
	BaseParamValueSpecial1     int
	BaseParamValueSpecial2     int
	BaseParamValueSpecial3     int
	BaseParamValueSpecial4     int
	BaseParamValueSpecial5     int
	Block                      int
	BlockRate                  int
	CanBeHq                    JSONBoolNumber
	Description                string
	IsAdvancedMeldingPermitted JSONBoolNumber
	IsCollectable              JSONBoolNumber
	IsCrestWorthy              JSONBoolNumber
	IsDyeable                  JSONBoolNumber
	IsIndisposable             JSONBoolNumber
	IsPvP                      JSONBoolNumber
	IsUnique                   JSONBoolNumber
	IsUntradable               JSONBoolNumber
	LevelEquip                 int
	LevelItem                  int
	MateriaSlotCount           int
	MaterializeType            int
	ModelMain                  string
	ModelSub                   string
	Rarity                     int
	StackSize                  int
	ClassJobCategoryACN        JSONBoolNumber `json:"ClassJobCategory.ACN"`
	ClassJobCategoryADV        JSONBoolNumber `json:"ClassJobCategory.ADV"`
	ClassJobCategoryALC        JSONBoolNumber `json:"ClassJobCategory.ALC"`
	ClassJobCategoryARC        JSONBoolNumber `json:"ClassJobCategory.ARC"`
	ClassJobCategoryARM        JSONBoolNumber `json:"ClassJobCategory.ARM"`
	ClassJobCategoryAST        JSONBoolNumber `json:"ClassJobCategory.AST"`
	ClassJobCategoryBLM        JSONBoolNumber `json:"ClassJobCategory.BLM"`
	ClassJobCategoryBRD        JSONBoolNumber `json:"ClassJobCategory.BRD"`
	ClassJobCategoryBSM        JSONBoolNumber `json:"ClassJobCategory.BSM"`
	ClassJobCategoryBTN        JSONBoolNumber `json:"ClassJobCategory.BTN"`
	ClassJobCategoryCNJ        JSONBoolNumber `json:"ClassJobCategory.CNJ"`
	ClassJobCategoryCRP        JSONBoolNumber `json:"ClassJobCategory.CRP"`
	ClassJobCategoryCUL        JSONBoolNumber `json:"ClassJobCategory.CUL"`
	ClassJobCategoryDRG        JSONBoolNumber `json:"ClassJobCategory.DRG"`
	ClassJobCategoryDRK        JSONBoolNumber `json:"ClassJobCategory.DRK"`
	ClassJobCategoryFSH        JSONBoolNumber `json:"ClassJobCategory.FSH"`
	ClassJobCategoryGLA        JSONBoolNumber `json:"ClassJobCategory.GLA"`
	ClassJobCategoryGSM        JSONBoolNumber `json:"ClassJobCategory.GSM"`
	ClassJobCategoryID         uint           `json:"ClassJobCategory.ID"`
	ClassJobCategoryLNC        JSONBoolNumber `json:"ClassJobCategory.LNC"`
	ClassJobCategoryLTW        JSONBoolNumber `json:"ClassJobCategory.LTW"`
	ClassJobCategoryMCH        JSONBoolNumber `json:"ClassJobCategory.MCH"`
	ClassJobCategoryMIN        JSONBoolNumber `json:"ClassJobCategory.MIN"`
	ClassJobCategoryMNK        JSONBoolNumber `json:"ClassJobCategory.MNK"`
	ClassJobCategoryMRD        JSONBoolNumber `json:"ClassJobCategory.MRD"`
	ClassJobCategoryNIN        JSONBoolNumber `json:"ClassJobCategory.NIN"`
	ClassJobCategoryName       string         `json:"ClassJobCategory.Name"`
	ClassJobCategoryPGL        JSONBoolNumber `json:"ClassJobCategory.PGL"`
	ClassJobCategoryPLD        JSONBoolNumber `json:"ClassJobCategory.PLD"`
	ClassJobCategoryRDM        JSONBoolNumber `json:"ClassJobCategory.RDM"`
	ClassJobCategoryROG        JSONBoolNumber `json:"ClassJobCategory.ROG"`
	ClassJobCategorySAM        JSONBoolNumber `json:"ClassJobCategory.SAM"`
	ClassJobCategorySCH        JSONBoolNumber `json:"ClassJobCategory.SCH"`
	ClassJobCategorySMN        JSONBoolNumber `json:"ClassJobCategory.SMN"`
	ClassJobCategoryTHM        JSONBoolNumber `json:"ClassJobCategory.THM"`
	ClassJobCategoryWAR        JSONBoolNumber `json:"ClassJobCategory.WAR"`
	ClassJobCategoryWHM        JSONBoolNumber `json:"ClassJobCategory.WHM"`
	ClassJobCategoryWVR        JSONBoolNumber `json:"ClassJobCategory.WVR"`
	ClassJobRepairID           uint           `json:"ClassJobRepair.ID"`
	ClassJobUseID              uint           `json:"ClassJobUse.ID"`
	CooldownSeconds            int            `json:"CooldownS"`
	DamageMagical              int            `json:"DamageMag"`
	DamagePhysical             int            `json:"DamagePhys"`
	DefenseMagical             int            `json:"DefenseMag"`
	DefensePhysical            int            `json:"DefensePhys"`
	DelayMilliseconds          int            `json:"DelayMs"`
	EquipSlotCategoryID        uint           `json:"EquipSlotCategory.ID"`
	ItemActionID               uint           `json:"ItemAction.ID"`
	ItemGlamourID              uint           `json:"ItemGlamour.ID"`
	ItemKindID                 uint           `json:"ItemKind.ID"`
	ItemKindName               string         `json:"ItemKind.Name"`
	ItemRepairID               uint           `json:"ItemRepair.ID"`
	ItemSearchCategoryID       uint           `json:"ItemSearchCategory.ID"`
	ItemSearchCategoryName     string         `json:"ItemSearchCategory.Name"`
	ItemSeriesID               uint           `json:"ItemSeries.ID"`
	ItemSpecialBonusID         uint           `json:"ItemSpecialBonus.ID"`
	ItemSpecialBonusParamID    uint           `json:"ItemSpecialBonusParam.ID"`
	ItemUICategoryID           uint           `json:"ItemUICategory.ID"`
	ItemUICategoryName         string         `json:"ItemUICategory.Name"`
	StainID                    uint           `json:"Stain.ID"`
}

// GetItem returns the Item details for a result
func (e *SearchResultEntry) GetItem() (*SearchResultItem, error) {
	if e.Type != IndexItem {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultItem)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultLeve holds all fields for a Leve search result
type SearchResultLeve struct {
	*SearchResultCommon
	URL                             string `json:"Url"`
	GamePatchID                     int    `json:"GamePatch.ID"`
	AllowanceCost                   int
	ClassJobLevel                   int
	Description                     string
	ExpReward                       int
	GilReward                       int
	ClassJobCategoryName            string `json:"ClassJobCategory.Name"`
	JournalGenreID                  uint   `json:"JournalGenre.ID"`
	JournalGenreJournalCategoryID   uint   `json:"JournalGenre.JournalCategory.ID"`
	JournalGenreJournalCategoryName string `json:"JournalGenre.JournalCategory.Name"`
	JournalGenreName                string `json:"JournalGenre.Name"`
	LeveAssignmentTypeID            uint   `json:"LeveAssignmentType.ID"`
	LeveClientID                    uint   `json:"LeveClient.ID"`
	PlaceNameIssuedID               uint   `json:"PlaceNameIssued.ID"`
	PlaceNameStartID                uint   `json:"PlaceNameStart.ID"`
	PlaceNameStartZoneID            uint   `json:"PlaceNameStartZone.ID"`
}

// GetLeve returns the Leve details for a result
func (e *SearchResultEntry) GetLeve() (*SearchResultLeve, error) {
	if e.Type != IndexLeve {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultLeve)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultMount holds all fields for a Mount search result
type SearchResultMount struct {
	*SearchResultCommon
	URL             string `json:"Url"`
	GamePatchID     int    `json:"GamePatch.ID"`
	FlyingCondition JSONBoolNumber
	IsFlying        JSONBoolNumber
}

// GetMount returns the Mount details for a result
func (e *SearchResultEntry) GetMount() (*SearchResultMount, error) {
	if e.Type != IndexMount {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultMount)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultPlaceName holds all fields for a PlaceName search result
type SearchResultPlaceName struct {
	*SearchResultCommon
	URL         string `json:"Url"`
	GamePatchID int    `json:"GamePatch.ID"`
}

// GetPlaceName returns the PlaceName details for a result
func (e *SearchResultEntry) GetPlaceName() (*SearchResultPlaceName, error) {
	if e.Type != IndexPlaceName {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultPlaceName)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultQuest holds all fields for a Quest search result
type SearchResultQuest struct {
	*SearchResultCommon
	URL                             string `json:"Url"`
	GamePatchID                     int    `json:"GamePatch.ID"`
	ClassJobLevel0                  int
	ClassJobLevel1                  int
	ExperiencePoints                int
	GilReward                       int
	IconSpecial                     string
	TomestoneCountReward            int
	ClassJobCategory0Name           string `json:"ClassJobCategory0.Name"`
	ClassJobCategory1Name           string `json:"ClassJobCategory1.Name"`
	EmoteRewardID                   uint   `json:"EmoteReward.ID"`
	JournalGenreJournalCategoryName string `json:"JournalGenre.JournalCategory.Name"`
	JournalGenreName                string `json:"JournalGenre.Name"`
	QuestID                         uint   `json:"Quest.ID"`
	TomestoneRewardID               uint   `json:"TomestoneReward.ID"`
}

// GetQuest returns the Quest details for a result
func (e *SearchResultEntry) GetQuest() (*SearchResultQuest, error) {
	if e.Type != IndexQuest {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultQuest)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultRecipe holds all fields for a Recipe search result
type SearchResultRecipe struct {
	*SearchResultCommon
	URL                           string         `json:"Url"`
	GamePatchID                   int            `json:"GamePatch.ID"`
	AmountResult                  int            `json:"AmountResult"`
	CanHq                         JSONBoolNumber `json:"CanHq"`
	CanQuickSynth                 JSONBoolNumber `json:"CanQuickSynth"`
	ClassJobID                    uint           `json:"ClassJob.ID"`
	ClassJobName                  string         `json:"ClassJob.Name"`
	DifficultyFactor              int            `json:"DifficultyFactor"`
	DurabilityFactor              int            `json:"DurabilityFactor"`
	ExpRewarded                   int            `json:"ExpRewarded"`
	IsSecondary                   JSONBoolNumber `json:"IsSecondary"`
	IsSpecializationRequired      JSONBoolNumber `json:"IsSpecializationRequired"`
	ItemResultID                  uint           `json:"ItemResult.ID"`
	QualityFactor                 int            `json:"QualityFactor"`
	QuickSynthControl             int            `json:"QuickSynthControl"`
	QuickSynthCraftsmanship       int            `json:"QuickSynthCraftsmanship"`
	RecipeLevelTableClassJobLevel int            `json:"RecipeLevelTable.ClassJobLevel"`
	RecipeLevelTableDifficulty    int            `json:"RecipeLevelTable.Difficulty"`
	RecipeLevelTableDurability    int            `json:"RecipeLevelTable.Durability"`
	RecipeLevelTableQuality       int            `json:"RecipeLevelTable.Quality"`
	RecipeLevelTableStars         int            `json:"RecipeLevelTable.Stars"`
	RequiredControl               int            `json:"RequiredControl"`
	RequiredCraftsmanship         int            `json:"RequiredCraftsmanship"`
	SecretRecipeBookID            uint           `json:"SecretRecipeBook.ID"`
	SecretRecipeBookName          string         `json:"SecretRecipeBook.Name"`
	StatusRequiredID              uint           `json:"StatusRequired.ID"`
}

// GetRecipe returns the Recipe details for a result
func (e *SearchResultEntry) GetRecipe() (*SearchResultRecipe, error) {
	if e.Type != IndexRecipe {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultRecipe)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultStatus holds all fields for a Status search result
type SearchResultStatus struct {
	*SearchResultCommon
	URL              string         `json:"Url"`
	GamePatchID      int            `json:"GamePatch.ID"`
	CanDispel        JSONBoolNumber `json:"CanDispel"`
	Category         int            `json:"Category"`
	Description      string         `json:"Description"`
	InflictedByActor JSONBoolNumber `json:"InflictedByActor"`
	Invisibility     JSONBoolNumber `json:"Invisibility"`
	IsFcBuff         JSONBoolNumber `json:"IsFcBuff"`
	IsPermanent      JSONBoolNumber `json:"IsPermanent"`
	LockActions      JSONBoolNumber `json:"LockActions"`
	LockControl      JSONBoolNumber `json:"LockControl"`
	LockMovement     JSONBoolNumber `json:"LockMovement"`
	MaxStacks        int            `json:"MaxStacks"`
	Name             string         `json:"Name"`
	Transfiguration  int            `json:"Transfiguration"`
}

// GetStatus returns the Status details for a result
func (e *SearchResultEntry) GetStatus() (*SearchResultStatus, error) {
	if e.Type != IndexStatus {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultStatus)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultTitle holds all fields for a Title search result
type SearchResultTitle struct {
	*SearchResultCommon
	URL         string `json:"Url"`
	GamePatchID int    `json:"GamePatch.ID"`
	NameFemale  string `json:"NameFemale"`
}

// GetTitle returns the Title details for a result
func (e *SearchResultEntry) GetTitle() (*SearchResultTitle, error) {
	if e.Type != IndexTitle {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultTitle)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultWeather holds all fields for a Weather search result
type SearchResultWeather struct {
	*SearchResultCommon
	URL         string `json:"Url"`
	GamePatchID int    `json:"GamePatch.ID"`
	Description string `json:"Description"`
}

// GetWeather returns the Weather details for a result
func (e *SearchResultEntry) GetWeather() (*SearchResultWeather, error) {
	if e.Type != IndexWeather {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultWeather)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// SearchResultBuddyEquip holds all fields for a BuddyEquip search result
type SearchResultBuddyEquip struct {
	*SearchResultCommon
	URL             string         `json:"Url"`
	GamePatchID     int            `json:"GamePatch.ID"`
	GrandCompanyID  uint           `json:"GrandCompany.ID"`
	IconBody        string         `json:"IconBody"`
	IconHead        string         `json:"IconHead"`
	IconLegs        string         `json:"IconLegs"`
	ModelBody       int            `json:"ModelBody"`
	ModelLegs       int            `json:"ModelLegs"`
	ModelTop        int            `json:"ModelTop"`
	Rarity          int            `json:"Rarity"`
	StartsWithVowel JSONBoolNumber `json:"StartsWithVowel"`
}

// GetBuddyEquip returns the BuddyEquip details for a result
func (e *SearchResultEntry) GetBuddyEquip() (*SearchResultBuddyEquip, error) {
	if e.Type != IndexBuddyEquip {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultBuddyEquip)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}

// GetChocoboGear is an alias for GetBuddyEquip
func (e *SearchResultEntry) GetChocoboGear() (*SearchResultBuddyEquip, error) {
	return e.GetBuddyEquip()
}

// SearchResultOrchestrion holds all fields for a Orchestrion search result
type SearchResultOrchestrion struct {
	*SearchResultCommon
	URL                                       string `json:"Url"`
	GamePatchID                               int    `json:"GamePatch.ID"`
	Description                               string `json:"Description"`
	OrchestrionUiparamID                      uint   `json:"OrchestrionUiparam.ID"`
	OrchestrionUiparamOrchestrionCategoryID   uint   `json:"OrchestrionUiparam.OrchestrionCategory.ID"`
	OrchestrionUiparamOrchestrionCategoryName string `json:"OrchestrionUiparam.OrchestrionCategory.Name"`
	OrchestrionUiparamOrder                   int    `json:"OrchestrionUiparam.Order"`
}

// GetOrchestrion returns the Orchestrion details for a result
func (e *SearchResultEntry) GetOrchestrion() (*SearchResultOrchestrion, error) {
	if e.Type != IndexOrchestrion {
		return nil, ErrUnexpectedType
	}

	v := new(SearchResultOrchestrion)
	if err := json.Unmarshal(e.raw, v); err != nil {
		return nil, err
	}
	return v, nil
}
