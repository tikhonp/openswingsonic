package models

import "time"

// License https://opensubsonic.netlify.app/docs/responses/license/
type License struct {
	Valid          bool       `json:"valid" xml:"valid,attr"`
	Email          *string    `json:"email,omitempty" xml:"email,attr,omitempty"`
	LicenseExpires *time.Time `json:"licenseExpires,omitempty" xml:"licenseExpires,attr,omitempty"`
	TrialExpires   *time.Time `json:"trialExpires,omitempty" xml:"trialExpires,attr,omitempty"`
}

// OpenSubsonicExtension https://opensubsonic.netlify.app/docs/responses/opensubsonicextension/
type OpenSubsonicExtension struct {
	Name     string  `json:"name" xml:"name"`
	Versions []int64 `json:"versions" xml:"versions"`
}
