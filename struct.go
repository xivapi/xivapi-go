package xivapi

// TranslationLanguage is the preferred translation language
type TranslationLanguage string

// The different available translations
const (
	TranslationEnglish  TranslationLanguage = "en"
	TranslationJapanese                     = "ja"
	TranslationGerman                       = "de"
	TranslationFrench                       = "fr"
	TranslationChinese                      = "cn"
	TranslationKorean                       = "kr"
)

// GlobalQueryParameters define the query parameters valid for all queries
type GlobalQueryParameters struct {
	Language TranslationLanguage `url:"language,omitempty"`
	// Pretty bool `url:"pretty"` // Pretty is used to prettify the json output
	// TODO: columns

	Tags string `url:"tags,omitempty"`
}
