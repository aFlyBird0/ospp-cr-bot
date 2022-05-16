package git

import (
	"fmt"
)

type user struct {
	userID   string
	platform Platform
}

var users map[PlatformType]map[string]User

func init() {
	users = make(map[PlatformType]map[string]User)
}

func (u *user) GetUserID() string {
	return u.userID
}

func (u *user) OfPlatform() Platform {
	return u.platform
}

// Deprecated: use union.registerUsersFromConfig instead
//func RegisterUsersFromConfig() (errors []error) {
//	usersFromConfig := config.GetUsers()
//	for _, u := range usersFromConfig {
//		for p, id := range u.Git {
//			if _, err := RegisterUser(PlatformType(p), id); err != nil {
//				errors = append(errors, err)
//			}
//		}
//	}
//	return
//}

func RegisterUser(p PlatformType, userID string) (User, error) {
	if _, ok := platformMap[p]; !ok {
		return nil, fmt.Errorf("user register err: user [%v] belong to un-registered git platform [%v]", userID, p)
	}
	if _, ok := users[p]; !ok {
		users[p] = make(map[string]User)
	}
	u := &user{userID: userID, platform: platformMap[p]}
	users[p][userID] = u
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
