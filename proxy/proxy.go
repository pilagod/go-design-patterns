package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found", id)
}

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)

	if err == nil {
		u.DidLastSearchUsedCache = true
	} else {
		user, err = u.SomeDatabase.FindUser(id)
		if err != nil {
			return User{}, err
		}
		u.DidLastSearchUsedCache = false
		u.addUserToStack(user)
	}
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache = append(u.StackCache, user)
	}
}
