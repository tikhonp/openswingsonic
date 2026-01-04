package util

import "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"

var AppVersion = "dev"

func GetBaseResponse() models.SubsonicBase {
	return models.SubsonicBase{
		Status:        "ok",
		Version:       "1.16.1",
		Type:          "openswingsonic",
		ServerVersion: AppVersion,
		OpenSubsonic:  true,
	}
}
