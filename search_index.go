package xivapi

import (
	"net/url"
	"strings"
)

// SearchIndex specifies the index to search in
type SearchIndex string

// The different indexes
const (
	IndexAchievement     SearchIndex = "achievement"
	IndexAction                      = "action"
	IndexBNPCName                    = "bnpcname"
	IndexCompanion                   = "companion"
	IndexENPCResident                = "enpcresident"
	IndexEmote                       = "emote"
	IndexFate                        = "fate"
	IndexInstanceContent             = "instancecontent"
	IndexItem                        = "item"
	IndexLeve                        = "leve"
	IndexMount                       = "mount"
	IndexPlaceName                   = "placename"
	IndexQuest                       = "quest"
	IndexRecipe                      = "recipe"
	IndexStatus                      = "status"
	IndexTitle                       = "title"
	IndexWeather                     = "weather"
	IndexBuddyEquip                  = "buddyequip"
	IndexOrchestrion                 = "orchestrion"

	// IndexEnemy is an alias for IndexBNPCName
	IndexEnemy = "bnpcname"
	// IndexMinion is an alias for IndexCompanion
	IndexMinion = "companion"
	// IndexNPC is an alias for IndexENPCResident
	IndexNPC = "enpcresident"
	// IndexChocoboGear is an alias for IndexBuddyEquip
	IndexChocoboGear = "buddyequip"
)

// SearchIndexes is the list of indexes to search in
type SearchIndexes []SearchIndex

func (si SearchIndexes) String() string {
	indexesMap := make(map[string]struct{})
	for _, index := range si {
		indexesMap[string(index)] = struct{}{}
	}

	indexes := make([]string, len(indexesMap))
	i := 0
	for index := range indexesMap {
		indexes[i] = index
		i++
	}

	return strings.Join(indexes, ",")
}

// EncodeValues encodes the list into a custom list of values
// This is a custom Encoder for go-querystring https://godoc.org/github.com/google/go-querystring/query
func (si SearchIndexes) EncodeValues(key string, v *url.Values) error {
	v.Set(key, si.String())
	return nil
}
