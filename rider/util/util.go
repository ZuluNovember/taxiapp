package util

import "github.com/ZuluNovember/taxiapp/rider/pkg/settings"

func Setup() {
	jwtSecret = []byte(settings.Configuration.Server.JwtSecret)
}
