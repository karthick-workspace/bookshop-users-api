package users

import (
	"fmt"
	"github.com/karthick-workspace/bookshop-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("User id %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DataCreated = result.DataCreated

	return nil
}

func (user *User) Save() *errors.RestErr {

	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email id %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	usersDB[user.Id] = user

	return nil
}
