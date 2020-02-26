package db

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

// User represents a real person.
type User struct {
	Id         bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	Deleted    bool            `json:"-"`
	Username   string          `json:"username" bson:"username" binding:"required"`
	Email      string          `json:"email" bson:"email" binding:"required"`
	Password   string          `json:"-" bson:"password"`
	Admin      bool            `json:"admin" bson:"admin"`
	Verified   bool            `json:"verified" bson:"verified"`
	Mmr        int64           `json:"mmr" bson:"mmr"`
	FriendList []bson.ObjectId `json:"friend_list" bson:"friend_list"`
}

type UpdateUser struct {
	Username   string          `json:"username" bson:"username" binding:"required"`
	Email      string          `json:"email" bson:"email" binding:"required"`
	Password   string          `json:"password" bson:"password"`
	Mmr        int64           `json:"mmr" bson:"mmr"`
	FriendList []bson.ObjectId `json:"friend_list" bson:"friend_list"`
	Admin      bool            `json:"admin" bson:"admin"`
	Verified   bool            `json:"verified" bson:"verified"`
}

func (user *User) Validate(nb int, nbUsername int, forbiddenMailAddresses []*ForbiddenMailAddress, err error) (bool, []error) {
	var validationErrors []error
	valid := true
	mailHostname := strings.SplitN(user.Email, "@", 2)
	if len(mailHostname) == 1 {
		validationErrors = append(validationErrors, errors.New("forbidden email address"))
	} else {
		println(mailHostname[1], forbiddenMailAddresses)
		for _, forbiddenMailAddress := range forbiddenMailAddresses {
			println(mailHostname[1], forbiddenMailAddress.Name)
			if forbiddenMailAddress.Name == mailHostname[1] {
				validationErrors = append(validationErrors, errors.New("forbidden email address"))
				break
			}
		}
	}

	if nb > 0 || err != nil || nbUsername > 0 {
		valid = false
		if err != nil {
			validationErrors = append(validationErrors, err)
		} else if nbUsername > 0 {
			validationErrors = append(validationErrors, errors.New("username already used"))
		} else {
			validationErrors = append(validationErrors, errors.New("email already used"))
		}
	}

	if len(validationErrors) > 0 {
		valid = false
	}

	return valid, validationErrors
}
