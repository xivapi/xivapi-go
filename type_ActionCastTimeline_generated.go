// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ActionCastTimeline struct {
	GameContentLinks struct {
		Action struct {
			AnimationStart []interface{} `json:"AnimationStart,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Name struct {
		ID int `json:"ID,omitempty"`
		Key string `json:"Key,omitempty"`
		Key_en string `json:"Key_en,omitempty"`
	}
	NameTarget string `json:"NameTarget,omitempty"`
	NameTargetID int `json:"NameTargetID,omitempty"`
	Url string `json:"Url,omitempty"`
	VFX int `json:"VFX,omitempty"`
}
