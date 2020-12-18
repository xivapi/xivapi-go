// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ActionComboRoute struct {
	Action1TargetID int `json:"Action1TargetID,omitempty"`
	Action3Target string `json:"Action3Target,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Action0TargetID int `json:"Action0TargetID,omitempty"`
	Action1Target string `json:"Action1Target,omitempty"`
	Action3 struct {
		ClassJobTarget string `json:"ClassJobTarget,omitempty"`
		ClassJobTargetID int `json:"ClassJobTargetID,omitempty"`
		TargetArea int `json:"TargetArea,omitempty"`
		ActionTimelineHitTargetID int `json:"ActionTimelineHitTargetID,omitempty"`
		AffectsPosition int `json:"AffectsPosition,omitempty"`
		CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
		ClassJobCategoryTargetID int `json:"ClassJobCategoryTargetID,omitempty"`
		ActionCategory struct {
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
		}
		ActionTimelineHit struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		AnimationEnd struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		Recast100ms int `json:"Recast100ms,omitempty"`
		EffectRange int `json:"EffectRange,omitempty"`
		IsRoleAction int `json:"IsRoleAction,omitempty"`
		Name string `json:"Name,omitempty"`
		Omen int `json:"Omen,omitempty"`
		Cost int `json:"Cost,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		ActionData int `json:"ActionData,omitempty"`
		VFX int `json:"VFX,omitempty"`
		CooldownGroup int `json:"CooldownGroup,omitempty"`
		CostType int `json:"CostType,omitempty"`
		XAxisModifier int `json:"XAxisModifier,omitempty"`
		AnimationEndTarget string `json:"AnimationEndTarget,omitempty"`
		Cast100ms int `json:"Cast100ms,omitempty"`
		IsPvP int `json:"IsPvP,omitempty"`
		PreservesCombo int `json:"PreservesCombo,omitempty"`
		Aspect int `json:"Aspect,omitempty"`
		ClassJob struct {
			NameEnglish_fr string `json:"NameEnglish_fr,omitempty"`
			Abbreviation_fr string `json:"Abbreviation_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			ItemSoulCrystal int `json:"ItemSoulCrystal,omitempty"`
			ItemStartingWeapon int `json:"ItemStartingWeapon,omitempty"`
			ModifierHitPoints int `json:"ModifierHitPoints,omitempty"`
			Name string `json:"Name,omitempty"`
			Abbreviation_en string `json:"Abbreviation_en,omitempty"`
			Abbreviation_ja string `json:"Abbreviation_ja,omitempty"`
			ModifierIntelligence int `json:"ModifierIntelligence,omitempty"`
			ModifierStrength int `json:"ModifierStrength,omitempty"`
			NameEnglish_ja string `json:"NameEnglish_ja,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ClassJobParent int `json:"ClassJobParent,omitempty"`
			LimitBreak1 int `json:"LimitBreak1,omitempty"`
			StartingLevel int `json:"StartingLevel,omitempty"`
			Abbreviation_de string `json:"Abbreviation_de,omitempty"`
			Prerequisite int `json:"Prerequisite,omitempty"`
			RelicQuest int `json:"RelicQuest,omitempty"`
			NameEnglish string `json:"NameEnglish,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			LimitBreak3 int `json:"LimitBreak3,omitempty"`
			ModifierPiety int `json:"ModifierPiety,omitempty"`
			Abbreviation string `json:"Abbreviation,omitempty"`
			ExpArrayIndex int `json:"ExpArrayIndex,omitempty"`
			ModifierDexterity int `json:"ModifierDexterity,omitempty"`
			ModifierMind int `json:"ModifierMind,omitempty"`
			ModifierVitality int `json:"ModifierVitality,omitempty"`
			NameEnglish_de string `json:"NameEnglish_de,omitempty"`
			NameEnglish_en string `json:"NameEnglish_en,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Icon string `json:"Icon,omitempty"`
			LimitBreak2 int `json:"LimitBreak2,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			UnlockQuest int `json:"UnlockQuest,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ModifierManaPoints int `json:"ModifierManaPoints,omitempty"`
		}
		StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		CanTargetDead int `json:"CanTargetDead,omitempty"`
		ActionComboTargetID int `json:"ActionComboTargetID,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Range int `json:"Range,omitempty"`
		ActionTimelineHitTarget string `json:"ActionTimelineHitTarget,omitempty"`
		ActionComboTarget string `json:"ActionComboTarget,omitempty"`
		CanTargetHostile int `json:"CanTargetHostile,omitempty"`
		CanTargetSelf int `json:"CanTargetSelf,omitempty"`
		ID int `json:"ID,omitempty"`
		ActionProcStatus int `json:"ActionProcStatus,omitempty"`
		CastType int `json:"CastType,omitempty"`
		CanTargetParty int `json:"CanTargetParty,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		AnimationStart int `json:"AnimationStart,omitempty"`
		ActionCategoryTargetID int `json:"ActionCategoryTargetID,omitempty"`
		ActionCombo struct {
			Aspect int `json:"Aspect,omitempty"`
			Cast100ms int `json:"Cast100ms,omitempty"`
			CastType int `json:"CastType,omitempty"`
			Icon string `json:"Icon,omitempty"`
			VFX int `json:"VFX,omitempty"`
			XAxisModifier int `json:"XAxisModifier,omitempty"`
			ActionCombo int `json:"ActionCombo,omitempty"`
			AnimationStart int `json:"AnimationStart,omitempty"`
			ClassJob int `json:"ClassJob,omitempty"`
			CooldownGroup int `json:"CooldownGroup,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			StatusGainSelf int `json:"StatusGainSelf,omitempty"`
			ActionCategory int `json:"ActionCategory,omitempty"`
			ActionProcStatus int `json:"ActionProcStatus,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			PreservesCombo int `json:"PreservesCombo,omitempty"`
			EffectRange int `json:"EffectRange,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			Omen int `json:"Omen,omitempty"`
			Range int `json:"Range,omitempty"`
			Recast100ms int `json:"Recast100ms,omitempty"`
			TargetArea int `json:"TargetArea,omitempty"`
			AnimationEnd int `json:"AnimationEnd,omitempty"`
			ID int `json:"ID,omitempty"`
			IsRoleAction int `json:"IsRoleAction,omitempty"`
			Name string `json:"Name,omitempty"`
			IconID int `json:"IconID,omitempty"`
			ActionTimelineHit int `json:"ActionTimelineHit,omitempty"`
			AffectsPosition int `json:"AffectsPosition,omitempty"`
			CanTargetHostile int `json:"CanTargetHostile,omitempty"`
			CanTargetParty int `json:"CanTargetParty,omitempty"`
			CanTargetSelf int `json:"CanTargetSelf,omitempty"`
			ClassJobLevel int `json:"ClassJobLevel,omitempty"`
			CostType int `json:"CostType,omitempty"`
			ActionData int `json:"ActionData,omitempty"`
			CanTargetDead int `json:"CanTargetDead,omitempty"`
			CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			Cost int `json:"Cost,omitempty"`
			AttackType int `json:"AttackType,omitempty"`
		}
		AnimationEndTargetID int `json:"AnimationEndTargetID,omitempty"`
		AttackType int `json:"AttackType,omitempty"`
		ClassJobCategory struct {
			CUL int `json:"CUL,omitempty"`
			Name string `json:"Name,omitempty"`
			RDM int `json:"RDM,omitempty"`
			SMN int `json:"SMN,omitempty"`
			WVR int `json:"WVR,omitempty"`
			ADV int `json:"ADV,omitempty"`
			ARC int `json:"ARC,omitempty"`
			BTN int `json:"BTN,omitempty"`
			NIN int `json:"NIN,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ACN int `json:"ACN,omitempty"`
			ALC int `json:"ALC,omitempty"`
			CNJ int `json:"CNJ,omitempty"`
			ID int `json:"ID,omitempty"`
			PGL int `json:"PGL,omitempty"`
			WHM int `json:"WHM,omitempty"`
			ARM int `json:"ARM,omitempty"`
			BSM int `json:"BSM,omitempty"`
			MIN int `json:"MIN,omitempty"`
			GLA int `json:"GLA,omitempty"`
			GSM int `json:"GSM,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			SAM int `json:"SAM,omitempty"`
			BRD int `json:"BRD,omitempty"`
			CRP int `json:"CRP,omitempty"`
			FSH int `json:"FSH,omitempty"`
			MCH int `json:"MCH,omitempty"`
			MNK int `json:"MNK,omitempty"`
			MRD int `json:"MRD,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			SCH int `json:"SCH,omitempty"`
			WAR int `json:"WAR,omitempty"`
			DRG int `json:"DRG,omitempty"`
			LNC int `json:"LNC,omitempty"`
			LTW int `json:"LTW,omitempty"`
			PLD int `json:"PLD,omitempty"`
			AST int `json:"AST,omitempty"`
			BLM int `json:"BLM,omitempty"`
			DRK int `json:"DRK,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			ROG int `json:"ROG,omitempty"`
			THM int `json:"THM,omitempty"`
		}
		ClassJobCategoryTarget string `json:"ClassJobCategoryTarget,omitempty"`
		Icon string `json:"Icon,omitempty"`
		ActionCategoryTarget string `json:"ActionCategoryTarget,omitempty"`
		IconID int `json:"IconID,omitempty"`
		ClassJobLevel int `json:"ClassJobLevel,omitempty"`
	}
	ID int `json:"ID,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	Action2 struct {
		AffectsPosition int `json:"AffectsPosition,omitempty"`
		ClassJobLevel int `json:"ClassJobLevel,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		XAxisModifier int `json:"XAxisModifier,omitempty"`
		ActionCategory struct {
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		ActionCategoryTargetID int `json:"ActionCategoryTargetID,omitempty"`
		ActionProcStatus int `json:"ActionProcStatus,omitempty"`
		CostType int `json:"CostType,omitempty"`
		ActionComboTarget string `json:"ActionComboTarget,omitempty"`
		ActionTimelineHitTargetID int `json:"ActionTimelineHitTargetID,omitempty"`
		ClassJobCategoryTarget string `json:"ClassJobCategoryTarget,omitempty"`
		PreservesCombo int `json:"PreservesCombo,omitempty"`
		Range int `json:"Range,omitempty"`
		Cast100ms int `json:"Cast100ms,omitempty"`
		Cost int `json:"Cost,omitempty"`
		ClassJob struct {
			LimitBreak1 int `json:"LimitBreak1,omitempty"`
			ModifierIntelligence int `json:"ModifierIntelligence,omitempty"`
			ModifierPiety int `json:"ModifierPiety,omitempty"`
			Abbreviation_de string `json:"Abbreviation_de,omitempty"`
			Abbreviation_en string `json:"Abbreviation_en,omitempty"`
			Abbreviation_ja string `json:"Abbreviation_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			LimitBreak2 int `json:"LimitBreak2,omitempty"`
			ModifierStrength int `json:"ModifierStrength,omitempty"`
			NameEnglish_ja string `json:"NameEnglish_ja,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ClassJobParent int `json:"ClassJobParent,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ModifierHitPoints int `json:"ModifierHitPoints,omitempty"`
			ModifierManaPoints int `json:"ModifierManaPoints,omitempty"`
			NameEnglish string `json:"NameEnglish,omitempty"`
			Abbreviation_fr string `json:"Abbreviation_fr,omitempty"`
			ExpArrayIndex int `json:"ExpArrayIndex,omitempty"`
			ModifierMind int `json:"ModifierMind,omitempty"`
			NameEnglish_de string `json:"NameEnglish_de,omitempty"`
			NameEnglish_fr string `json:"NameEnglish_fr,omitempty"`
			Prerequisite int `json:"Prerequisite,omitempty"`
			Abbreviation string `json:"Abbreviation,omitempty"`
			ModifierDexterity int `json:"ModifierDexterity,omitempty"`
			ModifierVitality int `json:"ModifierVitality,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Icon string `json:"Icon,omitempty"`
			ItemStartingWeapon int `json:"ItemStartingWeapon,omitempty"`
			StartingLevel int `json:"StartingLevel,omitempty"`
			UnlockQuest int `json:"UnlockQuest,omitempty"`
			ItemSoulCrystal int `json:"ItemSoulCrystal,omitempty"`
			LimitBreak3 int `json:"LimitBreak3,omitempty"`
			NameEnglish_en string `json:"NameEnglish_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			RelicQuest int `json:"RelicQuest,omitempty"`
		}
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		IsPvP int `json:"IsPvP,omitempty"`
		AnimationEnd struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		Aspect int `json:"Aspect,omitempty"`
		ActionCombo struct {
			EffectRange int `json:"EffectRange,omitempty"`
			IconID int `json:"IconID,omitempty"`
			CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
			CanTargetSelf int `json:"CanTargetSelf,omitempty"`
			ClassJobLevel int `json:"ClassJobLevel,omitempty"`
			IsRoleAction int `json:"IsRoleAction,omitempty"`
			ActionCategory int `json:"ActionCategory,omitempty"`
			ActionData int `json:"ActionData,omitempty"`
			AnimationStart int `json:"AnimationStart,omitempty"`
			Aspect int `json:"Aspect,omitempty"`
			CanTargetDead int `json:"CanTargetDead,omitempty"`
			CastType int `json:"CastType,omitempty"`
			ActionProcStatus int `json:"ActionProcStatus,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			AttackType int `json:"AttackType,omitempty"`
			CostType int `json:"CostType,omitempty"`
			PreservesCombo int `json:"PreservesCombo,omitempty"`
			Recast100ms int `json:"Recast100ms,omitempty"`
			XAxisModifier int `json:"XAxisModifier,omitempty"`
			CanTargetParty int `json:"CanTargetParty,omitempty"`
			CooldownGroup int `json:"CooldownGroup,omitempty"`
			Cost int `json:"Cost,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ActionTimelineHit int `json:"ActionTimelineHit,omitempty"`
			AffectsPosition int `json:"AffectsPosition,omitempty"`
			Cast100ms int `json:"Cast100ms,omitempty"`
			ID int `json:"ID,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			Name string `json:"Name,omitempty"`
			Omen int `json:"Omen,omitempty"`
			Range int `json:"Range,omitempty"`
			TargetArea int `json:"TargetArea,omitempty"`
			VFX int `json:"VFX,omitempty"`
			ActionCombo int `json:"ActionCombo,omitempty"`
			AnimationEnd int `json:"AnimationEnd,omitempty"`
			CanTargetHostile int `json:"CanTargetHostile,omitempty"`
			ClassJob int `json:"ClassJob,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		}
		ActionComboTargetID int `json:"ActionComboTargetID,omitempty"`
		ClassJobCategory struct {
			BRD int `json:"BRD,omitempty"`
			ROG int `json:"ROG,omitempty"`
			SAM int `json:"SAM,omitempty"`
			WVR int `json:"WVR,omitempty"`
			ARM int `json:"ARM,omitempty"`
			CRP int `json:"CRP,omitempty"`
			NIN int `json:"NIN,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			ADV int `json:"ADV,omitempty"`
			GLA int `json:"GLA,omitempty"`
			GSM int `json:"GSM,omitempty"`
			LNC int `json:"LNC,omitempty"`
			MRD int `json:"MRD,omitempty"`
			SMN int `json:"SMN,omitempty"`
			WAR int `json:"WAR,omitempty"`
			DRG int `json:"DRG,omitempty"`
			DRK int `json:"DRK,omitempty"`
			MCH int `json:"MCH,omitempty"`
			PLD int `json:"PLD,omitempty"`
			WHM int `json:"WHM,omitempty"`
			ALC int `json:"ALC,omitempty"`
			BLM int `json:"BLM,omitempty"`
			BSM int `json:"BSM,omitempty"`
			ID int `json:"ID,omitempty"`
			MNK int `json:"MNK,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			AST int `json:"AST,omitempty"`
			CUL int `json:"CUL,omitempty"`
			MIN int `json:"MIN,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			CNJ int `json:"CNJ,omitempty"`
			ARC int `json:"ARC,omitempty"`
			BTN int `json:"BTN,omitempty"`
			FSH int `json:"FSH,omitempty"`
			LTW int `json:"LTW,omitempty"`
			RDM int `json:"RDM,omitempty"`
			SCH int `json:"SCH,omitempty"`
			ACN int `json:"ACN,omitempty"`
			THM int `json:"THM,omitempty"`
			PGL int `json:"PGL,omitempty"`
		}
		EffectRange int `json:"EffectRange,omitempty"`
		ID int `json:"ID,omitempty"`
		CanTargetSelf int `json:"CanTargetSelf,omitempty"`
		CanTargetDead int `json:"CanTargetDead,omitempty"`
		IsRoleAction int `json:"IsRoleAction,omitempty"`
		CastType int `json:"CastType,omitempty"`
		VFX int `json:"VFX,omitempty"`
		AnimationEndTarget string `json:"AnimationEndTarget,omitempty"`
		AttackType int `json:"AttackType,omitempty"`
		CanTargetHostile int `json:"CanTargetHostile,omitempty"`
		AnimationStart int `json:"AnimationStart,omitempty"`
		Omen int `json:"Omen,omitempty"`
		AnimationEndTargetID int `json:"AnimationEndTargetID,omitempty"`
		CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
		ClassJobTargetID int `json:"ClassJobTargetID,omitempty"`
		CooldownGroup int `json:"CooldownGroup,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Recast100ms int `json:"Recast100ms,omitempty"`
		ActionTimelineHit struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		ActionTimelineHitTarget string `json:"ActionTimelineHitTarget,omitempty"`
		ClassJobTarget string `json:"ClassJobTarget,omitempty"`
		StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		TargetArea int `json:"TargetArea,omitempty"`
		ActionData int `json:"ActionData,omitempty"`
		CanTargetParty int `json:"CanTargetParty,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		ActionCategoryTarget string `json:"ActionCategoryTarget,omitempty"`
		ClassJobCategoryTargetID int `json:"ClassJobCategoryTargetID,omitempty"`
	}
	Action2Target string `json:"Action2Target,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Action0 struct {
		CastType int `json:"CastType,omitempty"`
		CooldownGroup int `json:"CooldownGroup,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		ActionData int `json:"ActionData,omitempty"`
		CanTargetSelf int `json:"CanTargetSelf,omitempty"`
		Aspect int `json:"Aspect,omitempty"`
		Cast100ms int `json:"Cast100ms,omitempty"`
		Omen int `json:"Omen,omitempty"`
		Range int `json:"Range,omitempty"`
		VFX int `json:"VFX,omitempty"`
		XAxisModifier int `json:"XAxisModifier,omitempty"`
		ActionProcStatus int `json:"ActionProcStatus,omitempty"`
		AffectsPosition int `json:"AffectsPosition,omitempty"`
		IsPvP int `json:"IsPvP,omitempty"`
		PreservesCombo int `json:"PreservesCombo,omitempty"`
		ClassJobCategory struct {
			BLM int `json:"BLM,omitempty"`
			BRD int `json:"BRD,omitempty"`
			PLD int `json:"PLD,omitempty"`
			CUL int `json:"CUL,omitempty"`
			DRG int `json:"DRG,omitempty"`
			FSH int `json:"FSH,omitempty"`
			ADV int `json:"ADV,omitempty"`
			BSM int `json:"BSM,omitempty"`
			ID int `json:"ID,omitempty"`
			MIN int `json:"MIN,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			WHM int `json:"WHM,omitempty"`
			WVR int `json:"WVR,omitempty"`
			ALC int `json:"ALC,omitempty"`
			ARM int `json:"ARM,omitempty"`
			BTN int `json:"BTN,omitempty"`
			GSM int `json:"GSM,omitempty"`
			MCH int `json:"MCH,omitempty"`
			MNK int `json:"MNK,omitempty"`
			PGL int `json:"PGL,omitempty"`
			THM int `json:"THM,omitempty"`
			ACN int `json:"ACN,omitempty"`
			ARC int `json:"ARC,omitempty"`
			CNJ int `json:"CNJ,omitempty"`
			DRK int `json:"DRK,omitempty"`
			GLA int `json:"GLA,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			SMN int `json:"SMN,omitempty"`
			LNC int `json:"LNC,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ROG int `json:"ROG,omitempty"`
			AST int `json:"AST,omitempty"`
			NIN int `json:"NIN,omitempty"`
			SCH int `json:"SCH,omitempty"`
			WAR int `json:"WAR,omitempty"`
			CRP int `json:"CRP,omitempty"`
			LTW int `json:"LTW,omitempty"`
			MRD int `json:"MRD,omitempty"`
			RDM int `json:"RDM,omitempty"`
			SAM int `json:"SAM,omitempty"`
		}
		ClassJobCategoryTargetID int `json:"ClassJobCategoryTargetID,omitempty"`
		Cost int `json:"Cost,omitempty"`
		CostType int `json:"CostType,omitempty"`
		Icon string `json:"Icon,omitempty"`
		IsRoleAction int `json:"IsRoleAction,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		CanTargetDead int `json:"CanTargetDead,omitempty"`
		ClassJob struct {
			ModifierIntelligence int `json:"ModifierIntelligence,omitempty"`
			Name string `json:"Name,omitempty"`
			StartingLevel int `json:"StartingLevel,omitempty"`
			Abbreviation_ja string `json:"Abbreviation_ja,omitempty"`
			ItemStartingWeapon int `json:"ItemStartingWeapon,omitempty"`
			LimitBreak2 int `json:"LimitBreak2,omitempty"`
			ModifierMind int `json:"ModifierMind,omitempty"`
			LimitBreak3 int `json:"LimitBreak3,omitempty"`
			NameEnglish_fr string `json:"NameEnglish_fr,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Prerequisite int `json:"Prerequisite,omitempty"`
			UnlockQuest int `json:"UnlockQuest,omitempty"`
			Abbreviation_de string `json:"Abbreviation_de,omitempty"`
			Abbreviation_fr string `json:"Abbreviation_fr,omitempty"`
			ClassJobParent int `json:"ClassJobParent,omitempty"`
			Icon string `json:"Icon,omitempty"`
			ModifierDexterity int `json:"ModifierDexterity,omitempty"`
			NameEnglish string `json:"NameEnglish,omitempty"`
			ItemSoulCrystal int `json:"ItemSoulCrystal,omitempty"`
			ModifierHitPoints int `json:"ModifierHitPoints,omitempty"`
			ModifierManaPoints int `json:"ModifierManaPoints,omitempty"`
			ModifierVitality int `json:"ModifierVitality,omitempty"`
			NameEnglish_de string `json:"NameEnglish_de,omitempty"`
			Abbreviation string `json:"Abbreviation,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ID int `json:"ID,omitempty"`
			NameEnglish_ja string `json:"NameEnglish_ja,omitempty"`
			NameEnglish_en string `json:"NameEnglish_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Abbreviation_en string `json:"Abbreviation_en,omitempty"`
			LimitBreak1 int `json:"LimitBreak1,omitempty"`
			ModifierPiety int `json:"ModifierPiety,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			RelicQuest int `json:"RelicQuest,omitempty"`
			ExpArrayIndex int `json:"ExpArrayIndex,omitempty"`
			ModifierStrength int `json:"ModifierStrength,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
		}
		CanTargetParty int `json:"CanTargetParty,omitempty"`
		ClassJobCategoryTarget string `json:"ClassJobCategoryTarget,omitempty"`
		ClassJobLevel int `json:"ClassJobLevel,omitempty"`
		ID int `json:"ID,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		ActionTimelineHitTarget string `json:"ActionTimelineHitTarget,omitempty"`
		ActionTimelineHitTargetID int `json:"ActionTimelineHitTargetID,omitempty"`
		AnimationEnd struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		AttackType int `json:"AttackType,omitempty"`
		CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
		Recast100ms int `json:"Recast100ms,omitempty"`
		TargetArea int `json:"TargetArea,omitempty"`
		ActionCategoryTargetID int `json:"ActionCategoryTargetID,omitempty"`
		ActionTimelineHit struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		AnimationStart int `json:"AnimationStart,omitempty"`
		ClassJobTarget string `json:"ClassJobTarget,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		ActionCategory struct {
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
		}
		ActionCategoryTarget string `json:"ActionCategoryTarget,omitempty"`
		AnimationEndTargetID int `json:"AnimationEndTargetID,omitempty"`
		CanTargetHostile int `json:"CanTargetHostile,omitempty"`
		ClassJobTargetID int `json:"ClassJobTargetID,omitempty"`
		EffectRange int `json:"EffectRange,omitempty"`
		ActionCombo int `json:"ActionCombo,omitempty"`
		AnimationEndTarget string `json:"AnimationEndTarget,omitempty"`
	}
	Action0Target string `json:"Action0Target,omitempty"`
	Action3TargetID int `json:"Action3TargetID,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Action1 struct {
		ActionTimelineHit struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		AffectsPosition int `json:"AffectsPosition,omitempty"`
		ClassJobCategoryTargetID int `json:"ClassJobCategoryTargetID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		ActionData int `json:"ActionData,omitempty"`
		CanTargetDead int `json:"CanTargetDead,omitempty"`
		CastType int `json:"CastType,omitempty"`
		CanTargetSelf int `json:"CanTargetSelf,omitempty"`
		ClassJob struct {
			Abbreviation_ja string `json:"Abbreviation_ja,omitempty"`
			ItemStartingWeapon int `json:"ItemStartingWeapon,omitempty"`
			LimitBreak3 int `json:"LimitBreak3,omitempty"`
			ModifierMind int `json:"ModifierMind,omitempty"`
			ModifierStrength int `json:"ModifierStrength,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Icon string `json:"Icon,omitempty"`
			ModifierVitality int `json:"ModifierVitality,omitempty"`
			NameEnglish_en string `json:"NameEnglish_en,omitempty"`
			Abbreviation string `json:"Abbreviation,omitempty"`
			ClassJobParent int `json:"ClassJobParent,omitempty"`
			ModifierHitPoints int `json:"ModifierHitPoints,omitempty"`
			ModifierManaPoints int `json:"ModifierManaPoints,omitempty"`
			Name string `json:"Name,omitempty"`
			StartingLevel int `json:"StartingLevel,omitempty"`
			NameEnglish string `json:"NameEnglish,omitempty"`
			Prerequisite int `json:"Prerequisite,omitempty"`
			ItemSoulCrystal int `json:"ItemSoulCrystal,omitempty"`
			ModifierIntelligence int `json:"ModifierIntelligence,omitempty"`
			NameEnglish_de string `json:"NameEnglish_de,omitempty"`
			Abbreviation_en string `json:"Abbreviation_en,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			RelicQuest int `json:"RelicQuest,omitempty"`
			UnlockQuest int `json:"UnlockQuest,omitempty"`
			ExpArrayIndex int `json:"ExpArrayIndex,omitempty"`
			ModifierDexterity int `json:"ModifierDexterity,omitempty"`
			ModifierPiety int `json:"ModifierPiety,omitempty"`
			NameEnglish_fr string `json:"NameEnglish_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Abbreviation_de string `json:"Abbreviation_de,omitempty"`
			Abbreviation_fr string `json:"Abbreviation_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			LimitBreak1 int `json:"LimitBreak1,omitempty"`
			LimitBreak2 int `json:"LimitBreak2,omitempty"`
			NameEnglish_ja string `json:"NameEnglish_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
		}
		ClassJobTargetID int `json:"ClassJobTargetID,omitempty"`
		ActionComboTargetID int `json:"ActionComboTargetID,omitempty"`
		Cost int `json:"Cost,omitempty"`
		Omen int `json:"Omen,omitempty"`
		AttackType int `json:"AttackType,omitempty"`
		PreservesCombo int `json:"PreservesCombo,omitempty"`
		TargetArea int `json:"TargetArea,omitempty"`
		Aspect int `json:"Aspect,omitempty"`
		ClassJobCategory struct {
			BRD int `json:"BRD,omitempty"`
			CRP int `json:"CRP,omitempty"`
			MRD int `json:"MRD,omitempty"`
			NIN int `json:"NIN,omitempty"`
			PLD int `json:"PLD,omitempty"`
			ROG int `json:"ROG,omitempty"`
			ACN int `json:"ACN,omitempty"`
			BSM int `json:"BSM,omitempty"`
			BTN int `json:"BTN,omitempty"`
			CUL int `json:"CUL,omitempty"`
			GSM int `json:"GSM,omitempty"`
			DRK int `json:"DRK,omitempty"`
			FSH int `json:"FSH,omitempty"`
			GLA int `json:"GLA,omitempty"`
			RDM int `json:"RDM,omitempty"`
			MCH int `json:"MCH,omitempty"`
			Name string `json:"Name,omitempty"`
			ARM int `json:"ARM,omitempty"`
			AST int `json:"AST,omitempty"`
			DRG int `json:"DRG,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			PGL int `json:"PGL,omitempty"`
			SCH int `json:"SCH,omitempty"`
			ARC int `json:"ARC,omitempty"`
			ID int `json:"ID,omitempty"`
			LNC int `json:"LNC,omitempty"`
			LTW int `json:"LTW,omitempty"`
			MIN int `json:"MIN,omitempty"`
			WAR int `json:"WAR,omitempty"`
			THM int `json:"THM,omitempty"`
			WVR int `json:"WVR,omitempty"`
			ADV int `json:"ADV,omitempty"`
			ALC int `json:"ALC,omitempty"`
			BLM int `json:"BLM,omitempty"`
			CNJ int `json:"CNJ,omitempty"`
			MNK int `json:"MNK,omitempty"`
			WHM int `json:"WHM,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			SAM int `json:"SAM,omitempty"`
			SMN int `json:"SMN,omitempty"`
		}
		Name string `json:"Name,omitempty"`
		ActionComboTarget string `json:"ActionComboTarget,omitempty"`
		ActionTimelineHitTargetID int `json:"ActionTimelineHitTargetID,omitempty"`
		AnimationEndTarget string `json:"AnimationEndTarget,omitempty"`
		CooldownGroup int `json:"CooldownGroup,omitempty"`
		EffectRange int `json:"EffectRange,omitempty"`
		CanTargetHostile int `json:"CanTargetHostile,omitempty"`
		IconID int `json:"IconID,omitempty"`
		IsRoleAction int `json:"IsRoleAction,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Range int `json:"Range,omitempty"`
		ActionCombo struct {
			StatusGainSelf int `json:"StatusGainSelf,omitempty"`
			AffectsPosition int `json:"AffectsPosition,omitempty"`
			CostType int `json:"CostType,omitempty"`
			ID int `json:"ID,omitempty"`
			Range int `json:"Range,omitempty"`
			Recast100ms int `json:"Recast100ms,omitempty"`
			VFX int `json:"VFX,omitempty"`
			CanTargetHostile int `json:"CanTargetHostile,omitempty"`
			CanTargetSelf int `json:"CanTargetSelf,omitempty"`
			ClassJob int `json:"ClassJob,omitempty"`
			Cost int `json:"Cost,omitempty"`
			CanTargetDead int `json:"CanTargetDead,omitempty"`
			Cast100ms int `json:"Cast100ms,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			Icon string `json:"Icon,omitempty"`
			ActionCombo int `json:"ActionCombo,omitempty"`
			ActionData int `json:"ActionData,omitempty"`
			ActionProcStatus int `json:"ActionProcStatus,omitempty"`
			CanTargetParty int `json:"CanTargetParty,omitempty"`
			Omen int `json:"Omen,omitempty"`
			Aspect int `json:"Aspect,omitempty"`
			AttackType int `json:"AttackType,omitempty"`
			ClassJobLevel int `json:"ClassJobLevel,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			EffectRange int `json:"EffectRange,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			XAxisModifier int `json:"XAxisModifier,omitempty"`
			CooldownGroup int `json:"CooldownGroup,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			ActionCategory int `json:"ActionCategory,omitempty"`
			ActionTimelineHit int `json:"ActionTimelineHit,omitempty"`
			AnimationStart int `json:"AnimationStart,omitempty"`
			CastType int `json:"CastType,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			PreservesCombo int `json:"PreservesCombo,omitempty"`
			TargetArea int `json:"TargetArea,omitempty"`
			AnimationEnd int `json:"AnimationEnd,omitempty"`
			CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
			IconID int `json:"IconID,omitempty"`
			IsRoleAction int `json:"IsRoleAction,omitempty"`
		}
		Recast100ms int `json:"Recast100ms,omitempty"`
		AnimationStart int `json:"AnimationStart,omitempty"`
		ClassJobTarget string `json:"ClassJobTarget,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		ActionCategory struct {
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		Cast100ms int `json:"Cast100ms,omitempty"`
		ActionProcStatus int `json:"ActionProcStatus,omitempty"`
		AnimationEnd struct {
			ID int `json:"ID,omitempty"`
			Key string `json:"Key,omitempty"`
			Key_en string `json:"Key_en,omitempty"`
		}
		ActionCategoryTarget string `json:"ActionCategoryTarget,omitempty"`
		AnimationEndTargetID int `json:"AnimationEndTargetID,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		ActionCategoryTargetID int `json:"ActionCategoryTargetID,omitempty"`
		ActionTimelineHitTarget string `json:"ActionTimelineHitTarget,omitempty"`
		IsPvP int `json:"IsPvP,omitempty"`
		StatusGainSelf int `json:"StatusGainSelf,omitempty"`
		XAxisModifier int `json:"XAxisModifier,omitempty"`
		CanTargetFriendly int `json:"CanTargetFriendly,omitempty"`
		CanTargetParty int `json:"CanTargetParty,omitempty"`
		ClassJobCategoryTarget string `json:"ClassJobCategoryTarget,omitempty"`
		ClassJobLevel int `json:"ClassJobLevel,omitempty"`
		CostType int `json:"CostType,omitempty"`
		ID int `json:"ID,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		VFX int `json:"VFX,omitempty"`
	}
	Action2TargetID int `json:"Action2TargetID,omitempty"`
}