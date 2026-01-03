package util

import "github.com/tikhonp/openswingsonic/internal/models"

func GetBaseResponse() models.SubsonicBase {
	return models.SubsonicBase{
		Status:        "ok",
		Version:       "1.16.1",
		Type:          "openwingsonic",
		ServerVersion: "1.0.0", // TODO: set actual version
		OpenSubsonic:  true,
	}
}
