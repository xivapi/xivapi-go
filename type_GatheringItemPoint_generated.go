// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type GatheringItemPoint struct {
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	GatheringPoint struct {
		GatheringPointBase struct {
			GatheringLevel int `json:"GatheringLevel,omitempty"`
			GatheringType int `json:"GatheringType,omitempty"`
			ID int `json:"ID,omitempty"`
			Item2 int `json:"Item2,omitempty"`
			Item3 int `json:"Item3,omitempty"`
			Item6 int `json:"Item6,omitempty"`
			Item7 int `json:"Item7,omitempty"`
			IsLimited int `json:"IsLimited,omitempty"`
			Item0 int `json:"Item0,omitempty"`
			Item1 int `json:"Item1,omitempty"`
			Item4 int `json:"Item4,omitempty"`
			Item5 int `json:"Item5,omitempty"`
		}
		TerritoryType struct {
			PlaceNameZone int `json:"PlaceNameZone,omitempty"`
			ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
			Bg string `json:"Bg,omitempty"`
			Bg_en string `json:"Bg_en,omitempty"`
			Map int `json:"Map,omitempty"`
			Name string `json:"Name,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
			Aetheryte int `json:"Aetheryte,omitempty"`
			ID int `json:"ID,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			WeatherRate int `json:"WeatherRate,omitempty"`
		}
		TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
		GatheringPointBonus0 struct {
			BonusType int `json:"BonusType,omitempty"`
			BonusValue int `json:"BonusValue,omitempty"`
			Condition int `json:"Condition,omitempty"`
			ConditionValue int `json:"ConditionValue,omitempty"`
			ID int `json:"ID,omitempty"`
		}
		GatheringPointBonus0Target string `json:"GatheringPointBonus0Target,omitempty"`
		GatheringPointBonus1 int `json:"GatheringPointBonus1,omitempty"`
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		GatheringPointBonus0TargetID int `json:"GatheringPointBonus0TargetID,omitempty"`
		GatheringSubCategory int `json:"GatheringSubCategory,omitempty"`
		TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
		GatheringPointBaseTarget string `json:"GatheringPointBaseTarget,omitempty"`
		GatheringPointBaseTargetID int `json:"GatheringPointBaseTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		PlaceName struct {
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
		}
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
	}
	GatheringPointTarget string `json:"GatheringPointTarget,omitempty"`
	GatheringPointTargetID int `json:"GatheringPointTargetID,omitempty"`
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
}
