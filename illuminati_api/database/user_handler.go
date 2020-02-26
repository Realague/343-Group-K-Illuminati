package database

import (
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
	"343-Group-K-Illuminati/illuminati_api/models/filters"
	"343-Group-K-Illuminati/illuminati_api/utils"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// FindUsersByQuery return all users that match the query
func (m *DatabaseService) FindUsersByQuery(research filters.ResearchData) ([]*db.User, error) {
	var users []*db.User
	if config.Config.IsUnitTest {
		return mockFindUsersByQuery(research)
	}
	query := m.getUserCollection().Find(research.QueryString).Skip(research.Pagination.Offset).Limit(research.Pagination.Limit)
	query = utils.Sort(query, research.Sorting)
	err := query.All(&users)
	return users, err
}

// FindUserByID return a specific user found by his ID
func (m *DatabaseService) FindUserByID(id string) (db.User, error) {
	var user db.User
	if valid := bson.IsObjectIdHex(id); valid == false {
		return db.User{}, errors.New("user not found")
	}
	if config.Config.IsUnitTest {
		return mockFindUserByID(id)
	}
	err := m.getUserCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	return user, err
}

// FindUserByKey return a specific user depending on the key used
func (m *DatabaseService) FindUserByKey(key, data string) (*db.User, error) {
	var user db.User
	if config.Config.IsUnitTest {
		return mockFindUserByKey(key, data)
	}
	err := m.getUserCollection().Find(bson.M{key: data}).One(&user)
	return &user, err
}

// CountUserByKey return a specific user depending on the key used
func (m *DatabaseService) CountUserByKey(key, data string) (int, error) {
	if config.Config.IsUnitTest {
		return mockCountUserByKey(key, data)
	}
	nb, err := m.getUserCollection().Find(bson.M{key: data}).Count()
	return nb, err
}

// InsertUser create a new user
func (m *DatabaseService) InsertUser(user *db.User) error {
	if config.Config.IsUnitTest {
		return mockInsertUser(user)
	}
	err := m.getUserCollection().Insert(user)
	return err
}

// DeleteUserByID remove a user
func (m *DatabaseService) DeleteUserByID(id string) error {
	if config.Config.IsUnitTest {
		return mockDeleteUserByID(id)
	}
	err := m.getUserCollection().RemoveId(bson.ObjectIdHex(id))
	return err
}

// UpdateUser update a user
func (m *DatabaseService) UpdateUser(id string, user *db.User) error {
	if valid := bson.IsObjectIdHex(id); valid == false {
		return errors.New("invalid id")
	}
	if config.Config.IsUnitTest {
		return mockUpdateUser(user)
	}
	err := m.getUserCollection().UpdateId(bson.ObjectIdHex(id), &user)
	return err
}

// FindOrCreate create a new admin user if he does not exist
func (m *DatabaseService) FindOrCreate(username string, password string, email string, isAdmin bool) error {
	var user db.User
	if config.Config.IsUnitTest {
		return mockFindOrCreate(username, password, email, isAdmin)
	}
	err := m.getUserCollection().Find(bson.M{"username": username}).One(&user)
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
		err := m.getUserCollection().Insert(&user)
		return err
	}
	return nil
}
