// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type MinionSkillType struct {
	GameContentLinks struct {
		CompanionTransient struct {
			MinionSkillType []interface{} `json:"MinionSkillType,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
}