package database

import (
	"gopkg.in/mgo.v2/bson"
	"343-Group-K-Illuminati/illuminati_api/config"
	"343-Group-K-Illuminati/illuminati_api/models/db"
)

func (m *DatabaseService) FindAllForbiddenMailAddress() ([]*db.ForbiddenMailAddress, error) {
	var forbiddenMailAddress []*db.ForbiddenMailAddress
	if config.Config.IsUnitTest {
		return mockFindAllForbiddenMailAddress()
	}
	err := m.getForbiddenMailAddressCollection().Find(bson.M{}).All(&forbiddenMailAddress)
	return forbiddenMailAddress, err
}
