package git

import "github.com/devstream-io/devstream/pkg/util/log"

var (
	platformMap map[PlatformType]Platform
)

func RegisterPlatform(platform Platform) {
	if platformMap == nil {
		platformMap = make(map[PlatformType]Platform)
	}
	if _, ok := platformMap[platform.GetType()]; ok {
		log.Fatalf("git platform of the same type [%v] already registered.", platform.GetType())
	}
	platformMap[platform.GetType()] = platform
	log.Infof("git platform [%v] registered successfully.", platform.GetType())
}

func GetPlatformByType(platformType PlatformType) Platform {
	return platformMap[platformType]
}

func GetPlatformMap() map[PlatformType]Platform {
	return platformMap
}
