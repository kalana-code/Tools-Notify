package app

type Application struct {
	ApplicationName string `json:"ApplicationName"`

}

type ApplicationInfo struct {
	ApplicationName string `json:"ApplicationName"`
	ApplicationId string `json:"ApplicationId"`
}

type ApplicationInfoResponse struct {
	ApplicationName string `json:"application_name"`
	ApplicationId string `json:"application_id"`
	ApplicationAccessToken string `json:"application_access_token"`
}

type ApplicationLevels struct {
	LevelName string
	LevelId     int
}

