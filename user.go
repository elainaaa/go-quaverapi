package quaverapi

import (
	"github.com/elainaaa/go-osuapi/structs"
	"fmt"
)

func GetUser(id int) (*User, error) {
	return nil, nil
}

func SearchUsers(username string) (*Users, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/search/%s", username))
	if err != nil {
		return nil, err
	}
	
	users := structs.Users{}
	err = json.Unmarshal(rawData, &users)
	if err != nil {
		return nil, err
	}
	
	return &users, nil
}