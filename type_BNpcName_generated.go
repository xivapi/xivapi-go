// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type BNpcName struct {
	GamePatch interface{} `json:"GamePatch,omitempty"`
	ID int `json:"ID,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	Url string `json:"Url,omitempty"`
	Name string `json:"Name,omitempty"`
	Plural_fr string `json:"Plural_fr,omitempty"`
	StartsWithVowel int `json:"StartsWithVowel,omitempty"`
	GameContentLinks struct {
		BattleLeve struct {
			BNpcName1 []interface{} `json:"BNpcName1,omitempty"`
			BNpcName2 []interface{} `json:"BNpcName2,omitempty"`
			BNpcName3 []interface{} `json:"BNpcName3,omitempty"`
			BNpcName4 []interface{} `json:"BNpcName4,omitempty"`
			BNpcName5 []interface{} `json:"BNpcName5,omitempty"`
			BNpcName6 []interface{} `json:"BNpcName6,omitempty"`
			BNpcName0 []interface{} `json:"BNpcName0,omitempty"`
		}
	}
	Plural_en string `json:"Plural_en,omitempty"`
	Icon string `json:"Icon,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	Plural string `json:"Plural,omitempty"`
	Plural_de string `json:"Plural_de,omitempty"`
	Plural_ja string `json:"Plural_ja,omitempty"`
}
