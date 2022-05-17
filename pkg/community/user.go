package community

import (
	"fmt"
)

var users map[Type]map[string]User

func init() {
	users = make(map[Type]map[string]User)
}

func RegisterUser(t Type, userID string) (User, error) {
	community, ok := communityMap[t]
	if !ok {
		return nil, fmt.Errorf("user register err: user [%v] belong to un-registered community [%v]", userID, t)
	}
	u, err := community.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user register err: failed to fetch user [%v] of community [%v], err: %v", userID, t, err)
	}
	if users[t] == nil {
		users[t] = make(map[string]User)
	}
	users[t][userID] = u
	return u, nil
}

func GetAllUsers() map[Type]map[string]User {
	return users
}

func GetUserByTypeAndID(p Type, id string) (User, bool) {
	if userMap, ok := users[p]; ok {
		if user, ok := userMap[id]; ok {
			return user, true
		}
	}
	return nil, false
}
