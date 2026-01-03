package models

import "time"

// License https://opensubsonic.netlify.app/docs/responses/license/
type License struct {
	Valid          bool       `json:"valid"`
	Email          *string    `json:"email,omitempty"`
	LicenseExpires *time.Time `json:"licenseExpires,omitempty"`
	TrialExpires   *time.Time `json:"trialExpires,omitempty"`
}

// OpenSubsonicExtension https://opensubsonic.netlify.app/docs/responses/opensubsonicextension/
type OpenSubsonicExtension struct {
	Name     string  `json:"name"`
	Versions []int64 `json:"versions"`
}
