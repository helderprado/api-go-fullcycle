package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" valid:"notnull"`
	Email     string    `json:"email" valid:"notnull"`
	Password  string    `json:"password" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func NewUser(name string, email string, password string) (*User, error) {
	user := User{
		ID:        uuid.NewV4().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
