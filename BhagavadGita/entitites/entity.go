package entities

type ChapterSchema struct {
	ChapterNumber     			int32 	`json:"chapter_number"`
	ChapterSummary             	string 	`json:"chapter_summary"`
	Name   						string 	`json:"name"`
	NameMeaning                	string  `json:"name_meaning"`
	NameTranslation				string 	`json:"name_translation"`
	name_transliterated  		string 	`json:"name_transliterated"`
	VersesCount 				int32   `json:"verses_count"`
}

type ServiceTokenDetailObject struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn int64 `json:"expires_in,omitempty"`
	ExtExpiresIn int64 `json:"ext_expires_in,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}