// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type ContentTalk struct {
	ContentTalkParamTarget string `json:"ContentTalkParamTarget,omitempty"`
	ContentTalkParamTargetID int `json:"ContentTalkParamTargetID,omitempty"`
	Text string `json:"Text,omitempty"`
	Text_en string `json:"Text_en,omitempty"`
	Url string `json:"Url,omitempty"`
	ContentTalkParam struct {
		ID int `json:"ID,omitempty"`
		Param int `json:"Param,omitempty"`
		TestAction int `json:"TestAction,omitempty"`
	}
	GameContentLinks struct {
		ContentNpcTalk struct {
			ContentTalk3 []interface{} `json:"ContentTalk3,omitempty"`
			ContentTalk1 []interface{} `json:"ContentTalk1,omitempty"`
			ContentTalk2 []interface{} `json:"ContentTalk2,omitempty"`
		}
	}
	ID int `json:"ID,omitempty"`
	Text_de string `json:"Text_de,omitempty"`
	Text_fr string `json:"Text_fr,omitempty"`
	Text_ja string `json:"Text_ja,omitempty"`
}
