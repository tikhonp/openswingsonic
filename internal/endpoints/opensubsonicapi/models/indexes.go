package models

import "time"

// https://opensubsonic.netlify.app/docs/responses/indexes/

type Indexes struct {
	Index           []Index `json:"index" xml:"index"`
	LastModified    *int64  `json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	IgnoredArticles *string `json:"ignoredArticles,omitempty" xml:"ignoredArticles,omitempty"`
}

type Index struct {
	Name   string   `json:"name" xml:"name"`
	Artist []Artist `json:"artist" xml:"artist"`
}

type Artist struct {
	ID      string     `json:"id" xml:"id"`
	Name    string     `json:"name" xml:"name"`
	Starred *time.Time `json:"starred,omitempty" xml:"starred,omitempty"`
}
