// Package middleware provides middleware functions for handling OpenSubsonic API errors.
package middleware

import (
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
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

// errorResponseXML defines the XML structure for error responses.
type errorResponseXML struct {
	XMLName       xml.Name        `xml:"subsonic-response"`
	Status        string          `xml:"status,attr"`
	Version       string          `xml:"version,attr"`
	Type          string          `xml:"type,attr,omitempty"`
	ServerVersion string          `xml:"serverVersion,attr,omitempty"`
	OpenSubsonic  bool            `xml:"openSubsonic,attr,omitempty"`
	Error         errorElementXML `xml:"error"`
}

type errorElementXML struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:"message,attr"`
}

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

		var errMessage string
		var errSubsonic OpenSubsonicError

		if httpErr, ok := errors.AsType[*echo.HTTPError](err); ok {
			if httpErr.Code == http.StatusNotFound {
				errSubsonic = TheRequestedDataWasNotFound
				errMessage = err.Error()
			} else {
				errSubsonic = GenericError
				errMessage = GenericError.Message
			}
		} else if errors.Is(err, swingmusic.ErrNotFound) {
			errSubsonic = TheRequestedDataWasNotFound
			errMessage = TheRequestedDataWasNotFound.Message
		} else if invalidValidationError, ok := errors.AsType[*validator.InvalidValidationError](err); ok {
			errSubsonic = RequiredParametrIsMissing
			errMessage = invalidValidationError.Error()
		} else if validationErrors, ok := errors.AsType[validator.ValidationErrors](err); ok {
			errSubsonic = RequiredParametrIsMissing
			errMessage = validationErrors.Error()
		} else if errors.As(err, &errSubsonic) {
			errMessage = errSubsonic.Message
		} else {
			c.Logger().Errorf("Unhandled error type: %T, error: %v", err, err)
			errSubsonic = GenericError
			errMessage = GenericError.Message
		}

		base := util.GetBaseResponse()
		jsonResponse := c.QueryParam("f") == "json"

		if jsonResponse {
			// JSON response format
			response := map[string]any{
				"subsonic-response": map[string]any{
					"status":        "failed",
					"version":       base.Version,
					"type":          base.Type,
					"serverVersion": base.ServerVersion,
					"openSubsonic":  base.OpenSubsonic,
					"error": map[string]any{
						"code":    errSubsonic.Code,
						"message": errMessage,
					},
				},
			}
			err = c.JSON(http.StatusOK, response)
		} else {
			// XML response format
			xmlResp := errorResponseXML{
				Status:        "failed",
				Version:       base.Version,
				Type:          base.Type,
				ServerVersion: base.ServerVersion,
				OpenSubsonic:  base.OpenSubsonic,
				Error: errorElementXML{
					Code:    errSubsonic.Code,
					Message: errMessage,
				},
			}
			err = c.XML(http.StatusOK, xmlResp)
		}

		if err != nil {
			c.Logger().Error(err)
		}

		return nil
	}
}
