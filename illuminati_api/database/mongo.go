package database

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// Connect will establish a new session
func (m *DatabaseService) Connect() {
	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.database)
	fmt.Printf("[MONGODB] Connected to the database - %s | %s | [%s, %s]\n", m.host, m.database, m.collections["users"], m.collections["forbidden_mail_addresses"])
}

// Config is a function that is used to config the database
func (m *DatabaseService) Config(host string, database string, collections map[string]string) {
	m.host = host
	m.database = database
	m.collections = collections
}

func (m *DatabaseService) getUserCollection() *mgo.Collection {
	return m.db.C(m.collections["users"])
}

func (m *DatabaseService) getForbiddenMailAddressCollection() *mgo.Collection {
	return m.db.C(m.collections["forbidden_mail_addresses"])
}

// GetDatabase return the current session
func (m *DatabaseService) GetDatabase() *mgo.Session {
	return m.db.Session
}
