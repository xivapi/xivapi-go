// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type Tomestones struct {
	GameContentLinks struct {
		TomestonesItem struct {
			Tomestones []interface{} `json:"Tomestones,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	WeeklyLimit int `json:"WeeklyLimit,omitempty"`
}
