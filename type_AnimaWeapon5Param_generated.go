// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type AnimaWeapon5Param struct {
	BaseParamTargetID int `json:"BaseParamTargetID,omitempty"`
	ID int `json:"ID,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Url string `json:"Url,omitempty"`
	BaseParam struct {
		Description_fr string `json:"Description_fr,omitempty"`
		Description_ja string `json:"Description_ja,omitempty"`
		HeadChestHandsLegsFeet% int `json:"HeadChestHandsLegsFeet%,omitempty"`
		Name_de string `json:"Name_de,omitempty"`
		OH% int `json:"OH%,omitempty"`
		Chest% int `json:"Chest%,omitempty"`
		ChestHead% int `json:"ChestHead%,omitempty"`
		ChestLegsFeet% int `json:"ChestLegsFeet%,omitempty"`
		ChestLegsGloves% int `json:"ChestLegsGloves%,omitempty"`
		Name_fr string `json:"Name_fr,omitempty"`
		Description_de string `json:"Description_de,omitempty"`
		Hands% int `json:"Hands%,omitempty"`
		Necklace% int `json:"Necklace%,omitempty"`
		ChestHeadLegsFeet% int `json:"ChestHeadLegsFeet%,omitempty"`
		Description string `json:"Description,omitempty"`
		Name_en string `json:"Name_en,omitempty"`
		Waist% int `json:"Waist%,omitempty"`
		Description_en string `json:"Description_en,omitempty"`
		Legs% int `json:"Legs%,omitempty"`
		Bracelet% int `json:"Bracelet%,omitempty"`
		Head% int `json:"Head%,omitempty"`
		Name string `json:"Name,omitempty"`
		2HWpn% int `json:"2HWpn%,omitempty"`
		ID int `json:"ID,omitempty"`
		LegsFeet% int `json:"LegsFeet%,omitempty"`
		Ring% int `json:"Ring%,omitempty"`
		1HWpn% int `json:"1HWpn%,omitempty"`
		Earring% int `json:"Earring%,omitempty"`
		Feet% int `json:"Feet%,omitempty"`
		Name_ja string `json:"Name_ja,omitempty"`
	}
	BaseParamTarget string `json:"BaseParamTarget,omitempty"`
	GameContentLinks struct {
		AnimaWeapon5 struct {
			Parameter1 []interface{} `json:"Parameter1,omitempty"`
		}
	}
	Name string `json:"Name,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
}