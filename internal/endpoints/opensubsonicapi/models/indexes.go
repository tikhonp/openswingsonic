package models

// https://opensubsonic.netlify.app/docs/responses/indexes/

type Indexes struct {
	Index           []Index `json:"index"`
	LastModified    *int64  `json:"lastModified,omitempty"`
	IgnoredArticles *string `json:"ignoredArticles,omitempty"`
}

type Index struct {
	Name   string   `json:"name"`
	Artist []Artist `json:"artist"`
}

type Artist struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Starred *string `json:"starred,omitempty"`
}
