package community

import (
	"fmt"
)

type user struct {
	userID    string
	community Community
}

func (u *user) GetUserID() string {
	return u.userID
}

func (u *user) OfCommunity() Community {
	return u.community
}

var users map[Type]map[string]User

func init() {
	users = make(map[Type]map[string]User)
}

// Deprecated: use union.registerUsersFromConfig instead
//func RegisterUsersFromConfig() (errors []error) {
//	usersFromConfig := config.GetUsers()
//	for _, u := range usersFromConfig {
//		for c, id := range u.Community {
//			if _, err := RegisterUser(c, id); err != nil {
//				errors = append(errors, err)
//			}
//		}
//	}
//	return
//}

func RegisterUser(c Type, userID string) (User, error) {
	if _, ok := communityMap[c]; !ok {
		return nil, fmt.Errorf("user register err: user [%v] belong to un-registered community [%v]", userID, c)
	}
	if users[c] == nil {
		users[c] = make(map[string]User)
	}
	u := &user{
		userID:    userID,
		community: communityMap[c],
	}
	users[c][userID] = u
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
