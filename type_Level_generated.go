// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Level struct {
	ID int `json:"ID,omitempty"`
	X int `json:"X,omitempty"`
	Y int `json:"Y,omitempty"`
	GameContentLinks struct {
		Leve struct {
			LevelLevemete []interface{} `json:"LevelLevemete,omitempty"`
		}
	}
	MapTarget string `json:"MapTarget,omitempty"`
	TerritoryTarget string `json:"TerritoryTarget,omitempty"`
	Url string `json:"Url,omitempty"`
	Z int `json:"Z,omitempty"`
	Map struct {
		Hierarchy int `json:"Hierarchy,omitempty"`
		ID int `json:"ID,omitempty"`
		MapMarkerRange int `json:"MapMarkerRange,omitempty"`
		OffsetX int `json:"OffsetX,omitempty"`
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
		SizeFactor int `json:"SizeFactor,omitempty"`
		PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
		PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
		PlaceNameSub int `json:"PlaceNameSub,omitempty"`
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
		DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
		DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
		MapFilename string `json:"MapFilename,omitempty"`
		MapFilenameId string `json:"MapFilenameId,omitempty"`
		TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
		OffsetY int `json:"OffsetY,omitempty"`
		PlaceName struct {
			ID int `json:"ID,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		PlaceNameRegion struct {
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
		}
		TerritoryType struct {
			Aetheryte int `json:"Aetheryte,omitempty"`
			ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
			Bg_en string `json:"Bg_en,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
			Bg string `json:"Bg,omitempty"`
			Map int `json:"Map,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			PlaceNameZone int `json:"PlaceNameZone,omitempty"`
			WeatherRate int `json:"WeatherRate,omitempty"`
		}
	}
	ObjectKey int `json:"ObjectKey,omitempty"`
	Radius int `json:"Radius,omitempty"`
	TerritoryTargetID int `json:"TerritoryTargetID,omitempty"`
	Type int `json:"Type,omitempty"`
	MapTargetID int `json:"MapTargetID,omitempty"`
	Territory struct {
		PlaceNameZone struct {
			Icon string `json:"Icon,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		ArrayEventHandler struct {
			MULTIREF13 int `json:"MULTIREF13,omitempty"`
			MULTIREF4 int `json:"MULTIREF4,omitempty"`
			MULTIREF8 int `json:"MULTIREF8,omitempty"`
			MULTIREF0 int `json:"MULTIREF0,omitempty"`
			MULTIREF1 int `json:"MULTIREF1,omitempty"`
			MULTIREF3 int `json:"MULTIREF3,omitempty"`
			MULTIREF5 int `json:"MULTIREF5,omitempty"`
			MULTIREF7 int `json:"MULTIREF7,omitempty"`
			MULTIREF15 int `json:"MULTIREF15,omitempty"`
			MULTIREF6 int `json:"MULTIREF6,omitempty"`
			ID int `json:"ID,omitempty"`
			MULTIREF10 int `json:"MULTIREF10,omitempty"`
			MULTIREF11 int `json:"MULTIREF11,omitempty"`
			MULTIREF12 int `json:"MULTIREF12,omitempty"`
			MULTIREF14 int `json:"MULTIREF14,omitempty"`
			MULTIREF2 int `json:"MULTIREF2,omitempty"`
			MULTIREF9 int `json:"MULTIREF9,omitempty"`
		}
		PlaceName struct {
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		PlaceNameZoneTarget string `json:"PlaceNameZoneTarget,omitempty"`
		AetheryteTargetID int `json:"AetheryteTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		ArrayEventHandlerTargetID int `json:"ArrayEventHandlerTargetID,omitempty"`
		Bg string `json:"Bg,omitempty"`
		Bg_en string `json:"Bg_en,omitempty"`
		MapTargetID int `json:"MapTargetID,omitempty"`
		Name string `json:"Name,omitempty"`
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
		Aetheryte struct {
			AethernetGroup int `json:"AethernetGroup,omitempty"`
			AetherstreamX int `json:"AetherstreamX,omitempty"`
			IsAetheryte int `json:"IsAetheryte,omitempty"`
			Level0 int `json:"Level0,omitempty"`
			Level1 int `json:"Level1,omitempty"`
			Level2 int `json:"Level2,omitempty"`
			Map int `json:"Map,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			AethernetName int `json:"AethernetName,omitempty"`
			AetherstreamY int `json:"AetherstreamY,omitempty"`
			ID int `json:"ID,omitempty"`
			Level3 int `json:"Level3,omitempty"`
			Territory int `json:"Territory,omitempty"`
		}
		AetheryteTarget string `json:"AetheryteTarget,omitempty"`
		TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
		WeatherRate int `json:"WeatherRate,omitempty"`
		MapTarget string `json:"MapTarget,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		PlaceNameRegion struct {
			ID int `json:"ID,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
		}
		PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
		PlaceNameZoneTargetID int `json:"PlaceNameZoneTargetID,omitempty"`
		ArrayEventHandlerTarget string `json:"ArrayEventHandlerTarget,omitempty"`
		Map struct {
			MapFilename string `json:"MapFilename,omitempty"`
			MapMarkerRange int `json:"MapMarkerRange,omitempty"`
			DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
			DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
			ID int `json:"ID,omitempty"`
			MapFilenameId string `json:"MapFilenameId,omitempty"`
			OffsetY int `json:"OffsetY,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			SizeFactor int `json:"SizeFactor,omitempty"`
			OffsetX int `json:"OffsetX,omitempty"`
			PlaceNameSub int `json:"PlaceNameSub,omitempty"`
			Hierarchy int `json:"Hierarchy,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			TerritoryType int `json:"TerritoryType,omitempty"`
		}
	}
	Yaw int `json:"Yaw,omitempty"`
	EventId int `json:"EventId,omitempty"`
}
