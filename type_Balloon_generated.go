// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Balloon struct {
	GameContentLinks struct {
		Behavior struct {
			Balloon []interface{} `json:"Balloon,omitempty"`
		}
	}
	GamePatch interface{} `json:"GamePatch,omitempty"`
	Dialogue_en string `json:"Dialogue_en,omitempty"`
	Dialogue_fr string `json:"Dialogue_fr,omitempty"`
	Dialogue_ja string `json:"Dialogue_ja,omitempty"`
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	Dialogue string `json:"Dialogue,omitempty"`
	Dialogue_de string `json:"Dialogue_de,omitempty"`
}