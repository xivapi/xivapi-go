// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type TutorialTank struct {
	ObjectiveTargetID int `json:"ObjectiveTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Objective struct {
		Objective struct {
			ID int `json:"ID,omitempty"`
			Text string `json:"Text,omitempty"`
			Text_de string `json:"Text_de,omitempty"`
			Text_en string `json:"Text_en,omitempty"`
			Text_fr string `json:"Text_fr,omitempty"`
			Text_ja string `json:"Text_ja,omitempty"`
		}
		RewardMeleeTarget string `json:"RewardMeleeTarget,omitempty"`
		RewardTank struct {
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			Block int `json:"Block,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			Name string `json:"Name,omitempty"`
			Singular string `json:"Singular,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			Icon string `json:"Icon,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			ID int `json:"ID,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			BaseParam3 int `json:"BaseParam3,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			Stain int `json:"Stain,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			Description string `json:"Description,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			IconID int `json:"IconID,omitempty"`
			PriceLow int `json:"PriceLow,omitempty"`
		}
		RewardTankTargetID int `json:"RewardTankTargetID,omitempty"`
		ObjectiveTargetID int `json:"ObjectiveTargetID,omitempty"`
		RewardTankTarget string `json:"RewardTankTarget,omitempty"`
		Exp int `json:"Exp,omitempty"`
		Gil int `json:"Gil,omitempty"`
		ObjectiveTarget string `json:"ObjectiveTarget,omitempty"`
		RewardMeleeTargetID int `json:"RewardMeleeTargetID,omitempty"`
		RewardRangedTargetID int `json:"RewardRangedTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		RewardMelee struct {
			BaseParam3 int `json:"BaseParam3,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			Description string `json:"Description,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			ID int `json:"ID,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
			PriceLow int `json:"PriceLow,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			Stain int `json:"Stain,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			Singular string `json:"Singular,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			IconID int `json:"IconID,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			Block int `json:"Block,omitempty"`
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
			Name string `json:"Name,omitempty"`
		}
		RewardRanged struct {
			PriceLow int `json:"PriceLow,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			IconID int `json:"IconID,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			Name string `json:"Name,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			Description string `json:"Description,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Block int `json:"Block,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			Singular string `json:"Singular,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			BaseParam3 int `json:"BaseParam3,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			Icon string `json:"Icon,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			Stain int `json:"Stain,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			ID int `json:"ID,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
		}
		RewardRangedTarget string `json:"RewardRangedTarget,omitempty"`
	}
	ObjectiveTarget string `json:"ObjectiveTarget,omitempty"`
}