package database

import (
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

var mockedUsers []*db.User

func mockFindUsersByQuery(research filters.ResearchData) ([]*db.User, error) {
	return mockedUsers, nil
}

func mockFindUserByID(id string) (db.User, error) {
	for _, user := range mockedUsers {
		if bson.ObjectIdHex(id) == user.Id {
			return *user, nil
		}
	}
	return db.User{}, errors.New("user not found")
}

func mockFindUserByKey(key, data string) (*db.User, error) {
	for _, user := range mockedUsers {
		if key == "email" && user.Email == data {
			return user, nil
		} else if key == "username" && user.Username == data {
			return user, nil
		}
	}
	return &db.User{}, errors.New("user not found")
}

func mockCountUserByKey(key, data string) (int, error) {
	var nb = 0
	for _, user := range mockedUsers {
		if key == "email" && user.Email == data {
			nb++
		} else if key == "username" && user.Username == data {
			nb++
		}
	}
	return nb, nil
}

func mockInsertUser(user *db.User) error {
	mockedUsers = append(mockedUsers, user)
	return nil
}

func mockDeleteUserByID(id string) error {
	var newMockedUsers []*db.User
	for _, user := range mockedUsers {
		if user.Id != bson.ObjectIdHex(id) {
			newMockedUsers = append(newMockedUsers, user)
		}
	}
	if len(newMockedUsers) == len(mockedUsers) {
		return errors.New("user not found")
	}
	mockedUsers = newMockedUsers
	return nil
}

func mockUpdateUser(user *db.User) error {
	for i := range mockedUsers {
		if mockedUsers[i].Id == user.Id {
			mockedUsers[i] = user
		}
	}
	return nil
}

func mockFindOrCreate(username string, password string, email string, isAdmin bool) error {
	var user db.User
	_, err := mockFindUserByKey("username", username)
	if err != nil {
		user.Id = bson.NewObjectId()
		user.Username = username
		user.Password = password
		user.Email = email
		user.Verified = true
		if isAdmin {
			user.Admin = true
		} else {
			user.Admin = false
		}
		err := mockInsertUser(&user)
		return err
	}
	return nil
}
