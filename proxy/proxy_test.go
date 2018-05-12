package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)

	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}
	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}
	knownIDs := [3]int32{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser with empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])

		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user ID doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be 1")
		}
		if proxy.DidLastSearchUsedCache != false {
			t.Error("No user can be returned fron an empty cache")
		}
	})

	t.Run("FindUser with one user in cache and ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])

		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for a object that is stored on it")
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cahce")
		}
	})

	t.Run("FindUser with overflow cache should apply FIFO to remove first user from cache", func(t *testing.T) {
		user1, _ := proxy.FindUser(knownIDs[0])
		user2, _ := proxy.FindUser(knownIDs[1])
		user3, _ := proxy.FindUser(knownIDs[2])

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}
		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow more than 2")
		}
		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
