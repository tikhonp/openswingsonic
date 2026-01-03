package opensubsonicauth

// authParams represents the authentication parameters
// provided by the client in a open subsonic request.
//
// https://opensubsonic.netlify.app/docs/opensubsonic-api/
type authParams struct {

	// The username.
	U string `query:"u" json:"u" form:"u"`

	// The password, either in clear text
	// or hex-encoded with a “enc:” prefix.
	// Since 1.13.0 this should only be used for testing purposes.
	P string `query:"p" json:"p" form:"p"`

	// (Since 1.13.0) The authentication token
	// computed as md5(password + salt).
	T string `query:"t" json:"t" form:"t"`

	// (Since 1.13.0) A random string (“salt”)
	// used as input for computing the password hash.
	S string `query:"s" json:"s" form:"s"`

	// [OS] An API key used for authentication
	APIKey string `query:"apiKey" json:"apiKey" form:"apiKey"`

	// The protocol version implemented by the client,
	// i.e., the version of the subsonic-rest-api.xsd schema used
	V string `query:"v" json:"v" form:"v"`

	// I do not want to parse "f" and "callback" parameters here
	// because they are not related to authentication.
}
