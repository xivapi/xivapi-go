// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type PvPAction struct {
	Action struct {
		AnimationEndTargetID int `json:"AnimationEndTargetID,omitempty"`
		IconID int `json:"IconID,omitempty"`
		PreservesCombo int `json:"PreservesCombo,omitempty"`
		ActionCategory struct {
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		AttackTypeTargetID int `json:"AttackTypeTargetID,omitempty"`
		ClassJobTarget string `json:"ClassJobTarget,omitempty"`
		XAxisModifier int `json:"XAxisModifier,omitempty"`
		Name string `json:"Name,omitempty"`
		ClassJobTargetID int `json:"ClassJobTargetID,omitempty"`
		ActionCategoryTarget string `json:"ActionCategoryTarget,omitempty"`
		ActionData int `json:"ActionData,omitempty"`
		Cost int `json:"Cost,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		CanTargetDead int `json:"CanTargetDead,omitempty"`
		CanTargetSelf int `json:"CanTargetSelf,omitempty"`
		ClassJobLevel int `json:"ClassJobLevel,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Range int `json:"Range,omitempty"`
		ActionCategoryTargetID int `json:"ActionCategoryTargetID,omitempty"`
		AnimationEnd struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		CanTargetParty int `json:"CanTargetParty,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		ActionTimelineHit struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		CanTargetHostile int `json:"CanTargetHostile,omitempty"`
		ClassJobCategoryTarget string `json:"ClassJobCategoryTarget,omitempty"`
		Omen int `json:"Omen,omitempty"`
		AttackType struct {
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
		}
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		ActionCombo int `json:"ActionCombo,omitempty"`
		ActionTimelineHitTarget string `json:"ActionTimelineHitTarget,omitempty"`
		AnimationEndTarget string `json:"AnimationEndTarget,omitempty"`
		ClassJobCategoryTargetID int `json:"ClassJobCategoryTargetID,omitempty"`
		ActionProcStatus int `json:"ActionProcStatus,omitempty"`
		ActionTimelineHitTargetID int `json:"ActionTimelineHitTargetID,omitempty"`
		AffectsPosition int `json:"AffectsPosition,omitempty"`
		AnimationStart int `json:"AnimationStart,omitempty"`
		CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		AttackTypeTarget string `json:"AttackTypeTarget,omitempty"`
		CooldownGroup int `json:"CooldownGroup,omitempty"`
		IsPvP int `json:"IsPvP,omitempty"`
		Recast100ms int `json:"Recast100ms,omitempty"`
		StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		TargetArea int `json:"TargetArea,omitempty"`
		CostType int `json:"CostType,omitempty"`
		EffectRange int `json:"EffectRange,omitempty"`
		IsRoleAction int `json:"IsRoleAction,omitempty"`
		CastType int `json:"CastType,omitempty"`
		ClassJob struct {
			Abbreviation_de string `json:"Abbreviation_de,omitempty"`
			ClassJobParent int `json:"ClassJobParent,omitempty"`
			Icon string `json:"Icon,omitempty"`
			LimitBreak1 int `json:"LimitBreak1,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			RelicQuest int `json:"RelicQuest,omitempty"`
			Abbreviation string `json:"Abbreviation,omitempty"`
			LimitBreak3 int `json:"LimitBreak3,omitempty"`
			ModifierHitPoints int `json:"ModifierHitPoints,omitempty"`
			ModifierManaPoints int `json:"ModifierManaPoints,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			UnlockQuest int `json:"UnlockQuest,omitempty"`
			LimitBreak2 int `json:"LimitBreak2,omitempty"`
			ModifierMind int `json:"ModifierMind,omitempty"`
			ExpArrayIndex int `json:"ExpArrayIndex,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Prerequisite int `json:"Prerequisite,omitempty"`
			ModifierVitality int `json:"ModifierVitality,omitempty"`
			NameEnglish string `json:"NameEnglish,omitempty"`
			StartingLevel int `json:"StartingLevel,omitempty"`
			ItemSoulCrystal int `json:"ItemSoulCrystal,omitempty"`
			ItemStartingWeapon int `json:"ItemStartingWeapon,omitempty"`
			ModifierPiety int `json:"ModifierPiety,omitempty"`
			ModifierStrength int `json:"ModifierStrength,omitempty"`
			NameEnglish_de string `json:"NameEnglish_de,omitempty"`
			NameEnglish_ja string `json:"NameEnglish_ja,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Abbreviation_ja string `json:"Abbreviation_ja,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ModifierIntelligence int `json:"ModifierIntelligence,omitempty"`
			NameEnglish_fr string `json:"NameEnglish_fr,omitempty"`
			Abbreviation_en string `json:"Abbreviation_en,omitempty"`
			Abbreviation_fr string `json:"Abbreviation_fr,omitempty"`
			ModifierDexterity int `json:"ModifierDexterity,omitempty"`
			NameEnglish_en string `json:"NameEnglish_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		ClassJobCategory struct {
			ARC int `json:"ARC,omitempty"`
			BLM int `json:"BLM,omitempty"`
			LTW int `json:"LTW,omitempty"`
			NIN int `json:"NIN,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			WVR int `json:"WVR,omitempty"`
			ADV int `json:"ADV,omitempty"`
			BRD int `json:"BRD,omitempty"`
			MNK int `json:"MNK,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			PLD int `json:"PLD,omitempty"`
			THM int `json:"THM,omitempty"`
			WAR int `json:"WAR,omitempty"`
			AST int `json:"AST,omitempty"`
			Name string `json:"Name,omitempty"`
			ROG int `json:"ROG,omitempty"`
			ACN int `json:"ACN,omitempty"`
			GSM int `json:"GSM,omitempty"`
			MIN int `json:"MIN,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			RDM int `json:"RDM,omitempty"`
			SAM int `json:"SAM,omitempty"`
			FSH int `json:"FSH,omitempty"`
			LNC int `json:"LNC,omitempty"`
			MRD int `json:"MRD,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			SCH int `json:"SCH,omitempty"`
			ALC int `json:"ALC,omitempty"`
			BSM int `json:"BSM,omitempty"`
			BTN int `json:"BTN,omitempty"`
			CNJ int `json:"CNJ,omitempty"`
			CRP int `json:"CRP,omitempty"`
			ID int `json:"ID,omitempty"`
			MCH int `json:"MCH,omitempty"`
			ARM int `json:"ARM,omitempty"`
			DRK int `json:"DRK,omitempty"`
			GLA int `json:"GLA,omitempty"`
			PGL int `json:"PGL,omitempty"`
			SMN int `json:"SMN,omitempty"`
			CUL int `json:"CUL,omitempty"`
			WHM int `json:"WHM,omitempty"`
			DRG int `json:"DRG,omitempty"`
		}
		Aspect int `json:"Aspect,omitempty"`
		Cast100ms int `json:"Cast100ms,omitempty"`
		VFX int `json:"VFX,omitempty"`
	}
	ActionTarget string `json:"ActionTarget,omitempty"`
	ActionTargetID int `json:"ActionTargetID,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
}