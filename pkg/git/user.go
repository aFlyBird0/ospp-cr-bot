package git

import (
	"fmt"
)

var users map[PlatformType]map[string]User

func init() {
	users = make(map[PlatformType]map[string]User)
}

func RegisterUser(t PlatformType, userID string) (User, error) {
	p, ok := platformMap[t]
	if !ok {
		return nil, fmt.Errorf("user register err: user [%v] belongs to un-registered git platform [%v]", userID, t)
	}
	u, err := p.GetUserInfoByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user register err: failed to fetch user  [%v] of [%v], err: [%v]", userID, t, err)
	}
	if _, ok := users[t]; !ok {
		users[t] = make(map[string]User)
	}
	users[t][userID] = u
	return u, nil
}

func GetAllUsers() map[PlatformType]map[string]User {
	return users
}

func GetUserByTypeAndID(p PlatformType, id string) (User, bool) {
	if userMap, ok := users[p]; ok {
		if user, ok := userMap[id]; ok {
			return user, true
		}
	}
	return nil, false
}
