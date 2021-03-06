// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type SatisfactionSupplyReward struct {
	GilMid int `json:"GilMid,omitempty"`
	QuantityHigh0 int `json:"QuantityHigh0,omitempty"`
	RewardCurrency0 struct {
		Limit int `json:"Limit,omitempty"`
		ID int `json:"ID,omitempty"`
		Item struct {
			Description string `json:"Description,omitempty"`
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			IconID int `json:"IconID,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
			PriceLow int `json:"PriceLow,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			Block int `json:"Block,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			ID int `json:"ID,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			Stain int `json:"Stain,omitempty"`
			BaseParam3 int `json:"BaseParam3,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			Singular string `json:"Singular,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
		}
		ItemTarget string `json:"ItemTarget,omitempty"`
		ItemTargetID int `json:"ItemTargetID,omitempty"`
	}
	GameContentLinks struct {
		SatisfactionSupply struct {
			Reward []interface{} `json:"Reward,omitempty"`
		}
	}
	GilLow int `json:"GilLow,omitempty"`
	QuantityMid0 int `json:"QuantityMid0,omitempty"`
	QuantityMid1 int `json:"QuantityMid1,omitempty"`
	SatisfactionLow int `json:"SatisfactionLow,omitempty"`
	QuantityHigh1 int `json:"QuantityHigh1,omitempty"`
	QuantityLow0 int `json:"QuantityLow0,omitempty"`
	RewardCurrency0Target string `json:"RewardCurrency0Target,omitempty"`
	RewardCurrency1 struct {
		ID int `json:"ID,omitempty"`
		Item struct {
			ItemSpecialBonusParam int `json:"ItemSpecialBonusParam,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			BaseParam4 int `json:"BaseParam4,omitempty"`
			BaseParamSpecial5 int `json:"BaseParamSpecial5,omitempty"`
			BaseParamValueSpecial3 int `json:"BaseParamValueSpecial3,omitempty"`
			Description_de string `json:"Description_de,omitempty"`
			Description_en string `json:"Description_en,omitempty"`
			Plural string `json:"Plural,omitempty"`
			Singular_en string `json:"Singular_en,omitempty"`
			BaseParamSpecial1 int `json:"BaseParamSpecial1,omitempty"`
			BaseParamValue2 int `json:"BaseParamValue2,omitempty"`
			BaseParamValue4 int `json:"BaseParamValue4,omitempty"`
			ClassJobUse int `json:"ClassJobUse,omitempty"`
			EquipSlotCategory int `json:"EquipSlotCategory,omitempty"`
			Plural_de string `json:"Plural_de,omitempty"`
			BaseParamSpecial0 int `json:"BaseParamSpecial0,omitempty"`
			BaseParamValueSpecial1 int `json:"BaseParamValueSpecial1,omitempty"`
			CooldownS int `json:"CooldownS,omitempty"`
			Description_fr string `json:"Description_fr,omitempty"`
			GrandCompany int `json:"GrandCompany,omitempty"`
			IconID int `json:"IconID,omitempty"`
			Stain int `json:"Stain,omitempty"`
			IsGlamourous int `json:"IsGlamourous,omitempty"`
			ItemGlamour int `json:"ItemGlamour,omitempty"`
			StackSize int `json:"StackSize,omitempty"`
			AetherialReduce int `json:"AetherialReduce,omitempty"`
			BaseParamValueSpecial0 int `json:"BaseParamValueSpecial0,omitempty"`
			CanBeHq int `json:"CanBeHq,omitempty"`
			Description string `json:"Description,omitempty"`
			ItemSpecialBonus int `json:"ItemSpecialBonus,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Salvage int `json:"Salvage,omitempty"`
			Singular_fr string `json:"Singular_fr,omitempty"`
			BaseParam1 int `json:"BaseParam1,omitempty"`
			BaseParamValue0 int `json:"BaseParamValue0,omitempty"`
			BaseParamValue1 int `json:"BaseParamValue1,omitempty"`
			ItemAction int `json:"ItemAction,omitempty"`
			Singular_ja string `json:"Singular_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			IsDyeable int `json:"IsDyeable,omitempty"`
			ItemRepair int `json:"ItemRepair,omitempty"`
			ItemUICategory int `json:"ItemUICategory,omitempty"`
			DefensePhys int `json:"DefensePhys,omitempty"`
			ID int `json:"ID,omitempty"`
			IsAdvancedMeldingPermitted int `json:"IsAdvancedMeldingPermitted,omitempty"`
			IsCollectable int `json:"IsCollectable,omitempty"`
			BaseParam5 int `json:"BaseParam5,omitempty"`
			BaseParamSpecial2 int `json:"BaseParamSpecial2,omitempty"`
			BlockRate int `json:"BlockRate,omitempty"`
			ClassJobCategory int `json:"ClassJobCategory,omitempty"`
			Plural_ja string `json:"Plural_ja,omitempty"`
			PriceLow int `json:"PriceLow,omitempty"`
			Singular string `json:"Singular,omitempty"`
			ItemSearchCategory int `json:"ItemSearchCategory,omitempty"`
			MateriaSlotCount int `json:"MateriaSlotCount,omitempty"`
			ModelMain string `json:"ModelMain,omitempty"`
			Plural_en string `json:"Plural_en,omitempty"`
			FilterGroup int `json:"FilterGroup,omitempty"`
			IsPvP int `json:"IsPvP,omitempty"`
			ItemSeries int `json:"ItemSeries,omitempty"`
			LevelItem int `json:"LevelItem,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			StartsWithVowel int `json:"StartsWithVowel,omitempty"`
			DefenseMag int `json:"DefenseMag,omitempty"`
			IsUntradable int `json:"IsUntradable,omitempty"`
			LevelEquip int `json:"LevelEquip,omitempty"`
			MaterializeType int `json:"MaterializeType,omitempty"`
			Rarity int `json:"Rarity,omitempty"`
			Block int `json:"Block,omitempty"`
			DamagePhys int `json:"DamagePhys,omitempty"`
			DelayMs int `json:"DelayMs,omitempty"`
			IsIndisposable int `json:"IsIndisposable,omitempty"`
			IsUnique int `json:"IsUnique,omitempty"`
			PriceMid int `json:"PriceMid,omitempty"`
			BaseParamModifier int `json:"BaseParamModifier,omitempty"`
			BaseParamValue5 int `json:"BaseParamValue5,omitempty"`
			ClassJobRepair int `json:"ClassJobRepair,omitempty"`
			Description_ja string `json:"Description_ja,omitempty"`
			BaseParamValueSpecial2 int `json:"BaseParamValueSpecial2,omitempty"`
			IsCrestWorthy int `json:"IsCrestWorthy,omitempty"`
			ModelSub string `json:"ModelSub,omitempty"`
			Plural_fr string `json:"Plural_fr,omitempty"`
			BaseParam0 int `json:"BaseParam0,omitempty"`
			BaseParam3 int `json:"BaseParam3,omitempty"`
			BaseParamSpecial3 int `json:"BaseParamSpecial3,omitempty"`
			BaseParamSpecial4 int `json:"BaseParamSpecial4,omitempty"`
			Singular_de string `json:"Singular_de,omitempty"`
			BaseParam2 int `json:"BaseParam2,omitempty"`
			BaseParamValueSpecial4 int `json:"BaseParamValueSpecial4,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			BaseParamValue3 int `json:"BaseParamValue3,omitempty"`
			BaseParamValueSpecial5 int `json:"BaseParamValueSpecial5,omitempty"`
			DamageMag int `json:"DamageMag,omitempty"`
			EquipRestriction int `json:"EquipRestriction,omitempty"`
		}
		ItemTarget string `json:"ItemTarget,omitempty"`
		ItemTargetID int `json:"ItemTargetID,omitempty"`
		Limit int `json:"Limit,omitempty"`
	}
	RewardCurrency1Target string `json:"RewardCurrency1Target,omitempty"`
	SatisfactionHigh int `json:"SatisfactionHigh,omitempty"`
	ID int `json:"ID,omitempty"`
	QuantityLow1 int `json:"QuantityLow1,omitempty"`
	RewardCurrency1TargetID int `json:"RewardCurrency1TargetID,omitempty"`
	SatisfactionMid int `json:"SatisfactionMid,omitempty"`
	Url string `json:"Url,omitempty"`
	GilHigh int `json:"GilHigh,omitempty"`
	RewardCurrency0TargetID int `json:"RewardCurrency0TargetID,omitempty"`
}
