package action

import "fingerprintClient/service"

func LoadProjectConfig() {
	service.GetServer().GetConfig("projectName")
}