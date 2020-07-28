package services

import (
	"github.com/alvarezcarlos/bookstore_users-api/domain/users"
	"github.com/alvarezcarlos/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil{
		return nil, err
	}

	if err := user.Save(); err != nil{
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr){
	result := &users.User{Id:userId}

	if getErr := result.Get(); getErr	!= nil {
		return nil, getErr
	}

	return result, nil
}
