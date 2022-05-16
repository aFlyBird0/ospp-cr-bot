package git

var (
	platformMap map[PlatformType]Platform
)

func RegisterPlatform(platform Platform) {
	if platformMap == nil {
		platformMap = make(map[PlatformType]Platform)
	}
	platformMap[platform.GetType()] = platform
}

func GetPlatformByType(platformType PlatformType) Platform {
	return platformMap[platformType]
}

func GetPlatformMap() map[PlatformType]Platform {
	return platformMap
}
