package iterator

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		user1 := &User{
			name: "a",
			age:  30,
		}
		user2 := &User{
			name: "b",
			age:  20,
		}

		userCollection := &UserCollection{
			users: []*User{user1, user2},
		}

		iterator := userCollection.createIterator()

		for iterator.hasNext() {
			user := iterator.next()
			fmt.Printf("User is %+v\n", user)
		}
	})
}
