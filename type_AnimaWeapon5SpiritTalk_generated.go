// Code generated by generate_structs - DO NOT EDIT.

package xivapi

type AnimaWeapon5SpiritTalk struct {
	DialogueTargetID int `json:"DialogueTargetID,omitempty"`
	GameContentLinks []interface{} `json:"GameContentLinks,omitempty"`
	ID int `json:"ID,omitempty"`
	Url string `json:"Url,omitempty"`
	Dialogue struct {
		Epilogue_fr string `json:"Epilogue_fr,omitempty"`
		Prologue_fr string `json:"Prologue_fr,omitempty"`
		Prologue_ja string `json:"Prologue_ja,omitempty"`
		Epilogue_de string `json:"Epilogue_de,omitempty"`
		Epilogue_en string `json:"Epilogue_en,omitempty"`
		ID int `json:"ID,omitempty"`
		Prologue string `json:"Prologue,omitempty"`
		Prologue_de string `json:"Prologue_de,omitempty"`
		Prologue_en string `json:"Prologue_en,omitempty"`
		Epilogue string `json:"Epilogue,omitempty"`
		Epilogue_ja string `json:"Epilogue_ja,omitempty"`
	}
	DialogueTarget string `json:"DialogueTarget,omitempty"`
}