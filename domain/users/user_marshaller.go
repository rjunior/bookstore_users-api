package users

import "encoding/json"

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"data_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"data_created"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for key, user := range users {
		result[key] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	} else {
		var privateUser PrivateUser
		json.Unmarshal(userJson, &privateUser)
		return privateUser
	}
}
