// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ScreenImage struct {
	Url string `json:"Url,omitempty"`
	GameContentLinks struct {
		ScenarioTree struct {
			Image []interface{} `json:"Image,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Image string `json:"Image,omitempty"`
	ImageID int `json:"ImageID,omitempty"`
}
