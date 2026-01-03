// Package middleware provides middleware functions for handling OpenSubsonic API errors.
package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/util"
)

type OpenSubsonicError = models.SubsonicError

var (
	GenericError                             = OpenSubsonicError{Code: 0, Message: "Generic error"}
	RequiredParametrIsMissing                = OpenSubsonicError{Code: 10, Message: "Required parameter is missing."}
	IncompatibleProtocolVersionClientUpgrade = OpenSubsonicError{Code: 20, Message: "Incompatible Subsonic REST protocol version. Client must upgrade."}
	IncompatibleProtocolVersionServerUpgrade = OpenSubsonicError{Code: 30, Message: "Incompatible Subsonic REST protocol version. Server must upgrade."}
	WrongUsernameOrPassword                  = OpenSubsonicError{Code: 40, Message: "Wrong username or password."}
	TokenAuthNotSupported                    = OpenSubsonicError{Code: 41, Message: "Token authentication not supported for LDAP users."}
	ProvidedAuthMechanisomNotSupported       = OpenSubsonicError{Code: 0, Message: "Provided authentication mechanism not supported."}
	MultipleConflictingAuthMechanisms        = OpenSubsonicError{Code: 43, Message: "Multiple conflicting authentication mechanisms provided."}
	InvalidAPIKey                            = OpenSubsonicError{Code: 44, Message: "Invalid API key."}
	UserIsNotAuthorized                      = OpenSubsonicError{Code: 50, Message: "User is not authorized for the given operation."}
	TrialIsOver                              = OpenSubsonicError{Code: 60, Message: "The trial period for the Subsonic server is over. Please upgrade to Subsonic Premium. Visit subsonic.org for details."}
	TheRequestedDataWasNotFound              = OpenSubsonicError{Code: 70, Message: "The requested data was not found."}
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		// Prevent double-handling if response committed
		if c.Response().Committed {
			return nil
		}

		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code == http.StatusNotFound {
				err = TheRequestedDataWasNotFound
			}
		}

		errSubsonic, ok := err.(OpenSubsonicError)
		if !ok {
			c.Logger().Errorf("Unhandled error type: %T, error: %v", err, err)
			errSubsonic = GenericError
		}

		base := util.GetBaseResponse()
		response := map[string]any{
			"subsonic-response": map[string]any{
				"status":        "failed",
				"version":       base.Version,
				"type":          base.Type,
				"serverVersion": base.Version,
				"openSubsonic":  base.OpenSubsonic,
				"error":         errSubsonic,
			},
		}
		err = c.JSON(http.StatusOK, response)
		if err != nil {
			c.Logger().Error(err)
		}

		return nil
	}
}
