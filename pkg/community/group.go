package community

import "fmt"

var groups map[Type]map[string]Group

func init() {
	groups = make(map[Type]map[string]Group)
}

type group struct {
	id        string
	community Community
}

func (g *group) GetGroupID() string {
	return g.id
}

func (g group) OfCommunity() Community {
	return g.community
}

func RegisterGroup(c Type, groupID string) (Group, error) {
	if _, ok := communityMap[c]; !ok {
		return nil, fmt.Errorf("user register err: group [%v] belong to un-registered community [%v]", groupID, c)
	}
	if groups[c] == nil {
		groups[c] = make(map[string]Group)
	}
	g := &group{
		id:        groupID,
		community: communityMap[c],
	}
	groups[c][groupID] = g
	return g, nil
}

func GetGroupMap() map[Type]map[string]Group {
	return groups
}

func GetGroupByTypeAndID(p Type, id string) Group {
	return groups[p][id]
}
