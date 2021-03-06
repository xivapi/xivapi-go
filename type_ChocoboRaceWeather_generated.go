// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ChocoboRaceWeather struct {
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	WeatherType2Target string `json:"WeatherType2Target,omitempty"`
	WeatherType2TargetID int `json:"WeatherType2TargetID,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	WeatherType1Target string `json:"WeatherType1Target,omitempty"`
	WeatherType1TargetID int `json:"WeatherType1TargetID,omitempty"`
	WeatherType2 struct {
		Description_en string `json:"Description_en,omitempty"`
		Description_ja string `json:"Description_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Description_de string `json:"Description_de,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Description_fr string `json:"Description_fr,omitempty"`
		Icon string `json:"Icon,omitempty"`
		Description string `json:"Description,omitempty"`
	}
	WeatherType1 struct {
		Description_de string `json:"Description_de,omitempty"`
		Description_en string `json:"Description_en,omitempty"`
		Description_fr string `json:"Description_fr,omitempty"`
		Description_ja string `json:"Description_ja,omitempty"`
		ID int `json:"ID,omitempty"`
		Icon string `json:"Icon,omitempty"`
		IconID int `json:"IconID,omitempty"`
		Description string `json:"Description,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
		Name string `json:"Name,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
	}
}
