package community

import "github.com/devstream-io/devstream/pkg/util/log"

var (
	communityMap map[Type]Community
)

//const (
//	CommunityTypeDingtalk Type = "dingtalk"
//)

func RegisterCommunity(p Community) {
	if communityMap == nil {
		communityMap = make(map[Type]Community)
	}
	if _, ok := communityMap[p.GetType()]; ok {
		log.Fatalf("community of the same type [%s] already registered.", p.GetType())
	}
	communityMap[p.GetType()] = p
	log.Infof("community [%s] registered successfully.", p.GetType())
}

func GetCommunityMap() map[Type]Community {
	return communityMap
}

func GetCommunityByType(p Type) Community {
	return communityMap[p]
}
