// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type PublicContent struct {
	ID int `json:"ID,omitempty"`
	MapIcon string `json:"MapIcon,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	TerritoryTypeTargetID int `json:"TerritoryTypeTargetID,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	MapIconID int `json:"MapIconID,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	TerritoryType struct {
		ArrayEventHandler int `json:"ArrayEventHandler,omitempty"`
		Bg string `json:"Bg,omitempty"`
		MapTarget string `json:"MapTarget,omitempty"`
		WeatherRate int `json:"WeatherRate,omitempty"`
		TerritoryIntendedUse int `json:"TerritoryIntendedUse,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		PlaceNameRegionTarget string `json:"PlaceNameRegionTarget,omitempty"`
		PlaceNameZone struct {
			ID int `json:"ID,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
		}
		PlaceNameZoneTarget string `json:"PlaceNameZoneTarget,omitempty"`
		Aetheryte struct {
			Level0 int `json:"Level0,omitempty"`
			Level2 int `json:"Level2,omitempty"`
			Territory int `json:"Territory,omitempty"`
			AethernetName int `json:"AethernetName,omitempty"`
			IsAetheryte int `json:"IsAetheryte,omitempty"`
			AetherstreamY int `json:"AetherstreamY,omitempty"`
			ID int `json:"ID,omitempty"`
			Level1 int `json:"Level1,omitempty"`
			Level3 int `json:"Level3,omitempty"`
			Map int `json:"Map,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			AethernetGroup int `json:"AethernetGroup,omitempty"`
			AetherstreamX int `json:"AetherstreamX,omitempty"`
		}
		Bg_en string `json:"Bg_en,omitempty"`
		Map struct {
			MapMarkerRange int `json:"MapMarkerRange,omitempty"`
			OffsetY int `json:"OffsetY,omitempty"`
			TerritoryType int `json:"TerritoryType,omitempty"`
			MapFilenameId string `json:"MapFilenameId,omitempty"`
			PlaceNameRegion int `json:"PlaceNameRegion,omitempty"`
			PlaceNameSub int `json:"PlaceNameSub,omitempty"`
			DiscoveryArrayByte int `json:"DiscoveryArrayByte,omitempty"`
			DiscoveryIndex int `json:"DiscoveryIndex,omitempty"`
			Hierarchy int `json:"Hierarchy,omitempty"`
			ID int `json:"ID,omitempty"`
			MapFilename string `json:"MapFilename,omitempty"`
			OffsetX int `json:"OffsetX,omitempty"`
			PlaceName int `json:"PlaceName,omitempty"`
			SizeFactor int `json:"SizeFactor,omitempty"`
		}
		PlaceNameRegion struct {
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			ID int `json:"ID,omitempty"`
			Icon string `json:"Icon,omitempty"`
			Name string `json:"Name,omitempty"`
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
		}
		PlaceNameRegionTargetID int `json:"PlaceNameRegionTargetID,omitempty"`
		PlaceNameTarget string `json:"PlaceNameTarget,omitempty"`
		PlaceNameTargetID int `json:"PlaceNameTargetID,omitempty"`
		AetheryteTarget string `json:"AetheryteTarget,omitempty"`
		AetheryteTargetID int `json:"AetheryteTargetID,omitempty"`
		ID int `json:"ID,omitempty"`
		MapTargetID int `json:"MapTargetID,omitempty"`
		PlaceName struct {
			NameNoArticle_de string `json:"NameNoArticle_de,omitempty"`
			NameNoArticle_en string `json:"NameNoArticle_en,omitempty"`
			NameNoArticle_fr string `json:"NameNoArticle_fr,omitempty"`
			NameNoArticle_ja string `json:"NameNoArticle_ja,omitempty"`
			Name_ja string `json:"Name_ja,omitempty"`
			Icon string `json:"Icon,omitempty"`
			NameNoArticle string `json:"NameNoArticle,omitempty"`
			Name_de string `json:"Name_de,omitempty"`
			Name_en string `json:"Name_en,omitempty"`
			Name_fr string `json:"Name_fr,omitempty"`
			ID int `json:"ID,omitempty"`
			Name string `json:"Name,omitempty"`
		}
		PlaceNameZoneTargetID int `json:"PlaceNameZoneTargetID,omitempty"`
	}
	TimeLimit int `json:"TimeLimit,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	TerritoryTypeTarget string `json:"TerritoryTypeTarget,omitempty"`
	Url string `json:"Url,omitempty"`
}
