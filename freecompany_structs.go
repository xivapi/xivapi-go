package xivapi

import "time"

type SearchFreeCompanyResult struct {
	Pagination Pagination                    `json:"pagination"`
	Results    []SearchFreeCompanyResultItem `json:"results"`
}

type SearchFreeCompanyResultItem struct {
	Crest  []string `json:"Crest"`
	ID     string   `json:"ID"`
	Name   string   `json:"Name"`
	Server string   `json:"Server"`
}

//PlayerContentQueryOptions provide search parameters for all /*/Search endpoints (PvPTeam, Character, FreeCompany, LinkShell)
type PlayerContentQueryOptions struct {
	Name   string `url:"name,omitempty"`
	Server string `url:"server,omitempty"`
	Page   int    `url:"page,omitempty"`
}

type FreeCompanyResult struct {
	FreeCompany FreeCompany         `json:"FreeCompany"`
	Members     []FreeCompanyMember `json:"FreeCompanyMembers"`
	Info        FreeCompanyInfo     `json:"Info"`
}

type FreeCompanyInfo struct {
	FreeCompanyState   FreeCompanyState `json:"FreeCompany"`
	FreeCompanyMembers FreeCompanyState `json:"FreeCompanyMembers"`
}

type FreeCompanyState struct {
	State            State  `json:"State"`
	Updated          string `json:"Updated"`
	UpdatedTimestamp time.Time
}

type FreeCompany struct {
	Active             string
	ActiveMemberCount  int
	Crest              []string
	Estate             FreeCompanyEstate
	Focus              []FreeCompanyFocus
	Formed             string
	FormedTimestamp    time.Time
	GrandCompany       string
	ID                 string
	Name               string
	ParseDate          string
	ParseDateTimestamp time.Time
	Rank               int
	Ranking            FreeCompanyRanking
	Recruitment        string                  `json:"Recruitment"`
	Reputation         []FreeCompanyReputation `json:"Reputation"`
	Seeking            []FreeCompanySeeking
	Server             string
	Slogon             string
	Tag                string
}

type FreeCompanySeeking struct {
	Icon   string
	Name   string
	Status bool
}

type FreeCompanyReputation struct {
	Name     string
	Progress int
	Rank     string
}

type FreeCompanyRanking struct {
	Weekly  string `json:"Weekly"`
	Monthly string `json:"Monthly"`
}

type FreeCompanyFocus struct {
	Icon   string
	Name   string
	Status bool
}

type FreeCompanyEstate struct {
	Greeting string
	Name     string
	Plot     string
}

type FreeCompanyMember struct {
	Avatar       string
	FeastMatches int
	ID           int64
	Name         string
	Rank         string
	RankIcon     string
	Server       string
}

type FreeCompanyQuery struct {
	Data    string `url:"data,omitempty"`
	Columns string `url:"columns,omitempty"`
}
