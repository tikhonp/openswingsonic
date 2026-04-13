package models

import "time"

// https://opensubsonic.netlify.app/docs/responses/indexes/

type Indexes struct {
	Shortcut        []Artist `json:"shortcut,omitempty" xml:"shortcut,omitempty"`
	Index           []Index  `json:"index" xml:"index"`
	LastModified    *int64   `json:"lastModified,omitempty" xml:"lastModified,attr,omitempty"`
	IgnoredArticles *string  `json:"ignoredArticles,omitempty" xml:"ignoredArticles,attr,omitempty"`
}

type Index struct {
	Name   string   `json:"name" xml:"name,attr"`
	Artist []Artist `json:"artist" xml:"artist"`
}

type Artist struct {
	ID      string     `json:"id" xml:"id,attr"`
	Name    string     `json:"name" xml:"name,attr"`
	Starred *time.Time `json:"starred,omitempty" xml:"starred,attr,omitempty"`
}
