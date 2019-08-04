package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/sanjeevchoubey/shippy/user-service/proto/user"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	*pb.User
}

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}
