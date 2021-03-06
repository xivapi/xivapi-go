// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ExVersion struct {
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		Quest struct {
			Expansion []interface{} `json:"Expansion,omitempty"`
		}
		ScenarioTreeTipsClassQuest struct {
			RequiredExpansion []interface{} `json:"RequiredExpansion,omitempty"`
		}
		BeastTribe struct {
			Expansion []interface{} `json:"Expansion,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
}
