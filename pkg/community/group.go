package community

import "fmt"

var groups map[Type]map[string]Group

func init() {
	groups = make(map[Type]map[string]Group)
}

func RegisterGroup(t Type, groupID string) (Group, error) {
	community, ok := communityMap[t]
	if !ok {
		return nil, fmt.Errorf("user register err: group [%v] belong to un-registered community [%v]", groupID, t)
	}
	g, err := community.GetGroupByID(groupID)
	if err != nil {
		return nil, fmt.Errorf("user register err: failed to fetch group [%v] of community [%v], err: %v", groupID, t, err)
	}
	if groups[t] == nil {
		groups[t] = make(map[string]Group)
	}
	groups[t][groupID] = g
	return g, nil
}

func GetGroupMap() map[Type]map[string]Group {
	return groups
}

func GetGroupByTypeAndID(p Type, id string) Group {
	return groups[p][id]
}
