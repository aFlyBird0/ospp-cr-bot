package community

var (
	communityMap map[Type]Community
)

const (
	CommunityTypeFeishu   Type = "feishu"
	CommunityTypeDingtalk Type = "dingtalk"
)

func RegisterCommunity(p Community) {
	if communityMap == nil {
		communityMap = make(map[Type]Community)
	}
	communityMap[p.GetType()] = p
}

func GetCommunityMap() map[Type]Community {
	return communityMap
}

func GetCommunityByType(p Type) Community {
	return communityMap[p]
}
