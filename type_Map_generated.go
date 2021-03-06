// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Map struct {
	MapMarkerRange int `json:"MapMarkerRange,omitempty"`
	OffsetY int `json:"OffsetY,omitempty"`
	PlaceName struct {
		NameNoArticle string `json:"NameNoArticle,omitempty"`
		NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
		NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name string `json:"Name,omitempty"`
		NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		ID int `json:"ID,omitempty"`
	}
	PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
	PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
	ID int `json:"ID,omitempty"`
	MapFilenameId string `json:"MapFilenameId,omitempty"`
	PlaceNameRegion struct {
		NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
		NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
		NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Name string `json:"Name,omitempty"`
		NameNoArticle string `json:"NameNoArticle,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
	}
	TerritoryType struct {
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
		Bg_en string `json:"Bg_en,omitempty"`
		ArrayEventHandlerTargetID int `json:"ArrayEventHandlerTargetID,omitempty"`
		PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
		PlaceNameZone struct {
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		}
		PlaceNameZoneTarget string `json:"PlaceNameZoneTarget,omitempty"`
		PlaceNameZoneTargetID int `json:"PlaceNameZoneTargetID,omitempty"`
		AetheryteTarget string `json:"AetheryteTarget,omitempty"`
		ArrayEventHandler struct {
			MULTIREF14 int `json:"MULTIREF14,omitempty"`
			MULTIREF3 int `json:"MULTIREF3,omitempty"`
			MULTIREF6 int `json:"MULTIREF6,omitempty"`
			MULTIREF7 int `json:"MULTIREF7,omitempty"`
			MULTIREF10 int `json:"MULTIREF10,omitempty"`
			MULTIREF1 int `json:"MULTIREF1,omitempty"`
			MULTIREF11 int `json:"MULTIREF11,omitempty"`
			MULTIREF12 int `json:"MULTIREF12,omitempty"`
			MULTIREF15 int `json:"MULTIREF15,omitempty"`
			MULTIREF2 int `json:"MULTIREF2,omitempty"`
			ID int `json:"ID,omitempty"`
			MULTIREF5 int `json:"MULTIREF5,omitempty"`
			MULTIREF0 int `json:"MULTIREF0,omitempty"`
			MULTIREF4 int `json:"MULTIREF4,omitempty"`
			MULTIREF8 int `json:"MULTIREF8,omitempty"`
			MULTIREF9 int `json:"MULTIREF9,omitempty"`
			MULTIREF13 int `json:"MULTIREF13,omitempty"`
		}
		ArrayEventHandlerTarget string `json:"ArrayEventHandlerTarget,omitempty"`
		ID int `json:"ID,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
		Aetheryte struct {
			ID int `json:"ID,omitempty"`
			IsAetheryte int `json:"IsAetheryte,omitempty"`
			Level1 int `json:"Level1,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			Level0 int `json:"Level0,omitempty"`
			Level2 int `json:"Level2,omitempty"`
			Level3 int `json:"Level3,omitempty"`
			Map int `json:"Map,omitempty"`
			AethernetGroup int `json:"AethernetGroup,omitempty"`
			AethernetName int `json:"AethernetName,omitempty"`
			AetherstreamX int `json:"AetherstreamX,omitempty"`
			AetherstreamY int `json:"AetherstreamY,omitempty"`
			Territory int `json:"Territory,omitempty"`
		}
		Bg string `json:"Bg,omitempty"`
		Map struct {
			DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
			OffsetY int `json:"OffsetY,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			TerritoryType int `json:"TerritoryType,omitempty"`
			ID int `json:"ID,omitempty"`
			MapFilename string `json:"MapFilename,omitempty"`
			MapMarkerRange int `json:"MapMarkerRange,omitempty"`
			OffsetX int `json:"OffsetX,omitempty"`
			PlaceNameSub int `json:"PlaceNameSub,omitempty"`
			SizeFactor int `json:"SizeFactor,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
			Hierarchy int `json:"Hierarchy,omitempty"`
			MapFilenameId string `json:"MapFilenameId,omitempty"`
		}
		MapTarget string `json:"MapTarget,omitempty"`
		MapTargetID int `json:"MapTargetID,omitempty"`
		PlaceName struct {
			ID int `json:"ID,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
		}
		PlaceNameRegion struct {
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
		}
		PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
		AetheryteTargetID int `json:"AetheryteTargetID,omitempty"`
		WeatherRate int `json:"WeatherRate,omitempty"`
	}
	TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
	MapFilename string `json:"MapFilename,omitempty"`
	OffsetX int `json:"OffsetX,omitempty"`
	PlaceNameSub int `json:"PlaceNameSub,omitempty"`
	PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
	TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
	Url string `json:"Url,omitempty"`
	DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
	DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
	GameContentLinks struct {
		ENpcResident struct {
			Map []interface{} `json:"Map,omitempty"`
		}
		Level struct {
			Map []interface{} `json:"Map,omitempty"`
		}
		TerritoryType struct {
			Map []interface{} `json:"Map,omitempty"`
		}
	}
	Hierarchy int `json:"Hierarchy,omitempty"`
	PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
	SizeFactor int `json:"SizeFactor,omitempty"`
}
