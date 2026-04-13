// Package utils provides utility functions for the OpenSubsonic API handlers.
package utils

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/util"
)

// subsonicResponseXML is the root structure for XML responses.
type subsonicResponseXML struct {
	XMLName       xml.Name `xml:"http://subsonic.org/restapi subsonic-response"`
	Status        string   `xml:"status,attr"`
	Version       string   `xml:"version,attr"`
	Type          string   `xml:"type,attr,omitempty"`
	ServerVersion string   `xml:"serverVersion,attr,omitempty"`
	OpenSubsonic  bool     `xml:"openSubsonic,attr,omitempty"`
	InnerXML      string   `xml:",innerxml"` // optional child elements
}

// marshalToInnerXML converts a value to an XML fragment with the given element name.
// Returns something like "<key>value</key>" or "<key><nested>...</nested></key>".
func marshalToInnerXML(key string, value any) (string, error) {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	start := xml.StartElement{Name: xml.Name{Local: key}}
	if err := enc.EncodeElement(value, start); err != nil {
		return "", err
	}
	if err := enc.Flush(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderEmptyResponse returns a successful Subsonic response with no additional data.
// Handles both JSON (?f=json) and XML (default) formats.
func RenderEmptyResponse(c echo.Context) error {
	base := util.GetBaseResponse()

	jsonResponse := c.QueryParam("f") == "json"
	if jsonResponse {
		// JSON format: wrap inside "subsonic-response" object
		response := map[string]any{
			"subsonic-response": map[string]any{
				"status":        "ok",
				"version":       base.Version,
				"type":          base.Type,
				"serverVersion": base.ServerVersion,
				"openSubsonic":  base.OpenSubsonic,
			},
		}
		return c.JSON(http.StatusOK, response)
	}

	// XML format: root element is <subsonic-response>
	resp := subsonicResponseXML{
		Status:        "ok",
		Version:       base.Version,
		Type:          base.Type,
		ServerVersion: base.ServerVersion,
		OpenSubsonic:  base.OpenSubsonic,
		InnerXML:      "", // no child elements
	}
	return c.XML(http.StatusOK, resp)
}

// RenderResponse returns a successful Subsonic response with one additional child element.
// The child element is named by the provided key and contains the value i.
// Handles both JSON (?f=json) and XML (default) formats.
func RenderResponse(c echo.Context, key string, i any) error {
	base := util.GetBaseResponse()

	jsonResponse := c.QueryParam("f") == "json"
	if jsonResponse {
		// JSON format: wrap inside "subsonic-response" object
		response := map[string]any{
			"subsonic-response": map[string]any{
				"status":        "ok",
				"version":       base.Version,
				"type":          base.Type,
				"serverVersion": base.ServerVersion,
				"openSubsonic":  base.OpenSubsonic,
				key:             i,
			},
		}
		return c.JSON(http.StatusOK, response)
	}

	// XML format: marshal the child value to an XML fragment
	innerXML, err := marshalToInnerXML(key, i)
	if err != nil {
		// Fallback or return an error response
		return errors.Join(errors.New("failed to marshal response"), err)
	}

	resp := subsonicResponseXML{
		Status:        "ok",
		Version:       base.Version,
		Type:          base.Type,
		ServerVersion: base.ServerVersion,
		OpenSubsonic:  base.OpenSubsonic,
		InnerXML:      innerXML,
	}
	return c.XML(http.StatusOK, resp)
}
