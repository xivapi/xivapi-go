// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ClassJobCategory struct {
	PLD int `json:"PLD,omitempty"`
	SAM int `json:"SAM,omitempty"`
	Url string `json:"Url,omitempty"`
	ALC int `json:"ALC,omitempty"`
	BRD int `json:"BRD,omitempty"`
	LTW int `json:"LTW,omitempty"`
	MRD int `json:"MRD,omitempty"`
	Name_ja string `json:"Name_ja,omitempty"`
	WHM int `json:"WHM,omitempty"`
	ADV int `json:"ADV,omitempty"`
	DRG int `json:"DRG,omitempty"`
	GameContentLinks struct {
		Action struct {
			ClassJobCategory []interface{} `json:"ClassJobCategory,omitempty"`
		}
		Item struct {
			ClassJobCategory []interface{} `json:"ClassJobCategory,omitempty"`
		}
		Quest struct {
			ClassJobCategory1 []interface{} `json:"ClassJobCategory1,omitempty"`
			ClassJobCategory0 []interface{} `json:"ClassJobCategory0,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name_en string `json:"Name_en,omitempty"`
	Name_fr string `json:"Name_fr,omitempty"`
	RDM int `json:"RDM,omitempty"`
	ROG int `json:"ROG,omitempty"`
	ACN int `json:"ACN,omitempty"`
	ARM int `json:"ARM,omitempty"`
	DRK int `json:"DRK,omitempty"`
	ARC int `json:"ARC,omitempty"`
	WVR int `json:"WVR,omitempty"`
	THM int `json:"THM,omitempty"`
	MCH int `json:"MCH,omitempty"`
	MNK int `json:"MNK,omitempty"`
	Name string `json:"Name,omitempty"`
	MIN int `json:"MIN,omitempty"`
	PGL int `json:"PGL,omitempty"`
	SMN int `json:"SMN,omitempty"`
	AST int `json:"AST,omitempty"`
	BLM int `json:"BLM,omitempty"`
	GLA int `json:"GLA,omitempty"`
	GSM int `json:"GSM,omitempty"`
	NIN int `json:"NIN,omitempty"`
	SCH int `json:"SCH,omitempty"`
	WAR int `json:"WAR,omitempty"`
	CRP int `json:"CRP,omitempty"`
	CUL int `json:"CUL,omitempty"`
	FSH int `json:"FSH,omitempty"`
	LNC int `json:"LNC,omitempty"`
	Name_de string `json:"Name_de,omitempty"`
	BSM int `json:"BSM,omitempty"`
	BTN int `json:"BTN,omitempty"`
	CNJ int `json:"CNJ,omitempty"`
}
