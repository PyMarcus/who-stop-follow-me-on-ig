package internal

import (
	"fmt"

	"github.com/ahmdrz/goinsta/v2"
)

func LoginService(userName, password string) (*goinsta.Instagram, error) {
	insta := goinsta.New(userName, password)
	if err := insta.Login(); err != nil {
		return nil, err
	}
	return insta, nil
}

// GetFollowers get followers users by name
func GetFollowers(insta *goinsta.Instagram) []string {
	fmt.Println("Get followers...")
	followers := []string{}

	me := insta.Account.Followers()
	for me.Next() {
		for _, user := range me.Users {
			followers = append(followers, user.Username)
		}
	}
	return followers
}

// GetFollowing get followers users by name
func GetFollowing(insta *goinsta.Instagram) []string {
	fmt.Println("Get following...")
	followings := []string{}

	me := insta.Account.Following()
	for me.Next() {
		for _, user := range me.Users {
			followings = append(followings, user.Username)
		}
	}
	return followings
}
