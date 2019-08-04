package main

import (
	"database/sql"

	pb "github.com/sanjeevchoubey/shippy/user-service/proto/user"
)

type Repository interface {
	Create(*pb.User) error
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(in *pb.User) error {
	name := (*pb.User).GetName(in)
	email := (*pb.User).GetEmail(in)
	password := (*pb.User).GetPassword(in)
	company := (*pb.User).GetCompany(in)
	id := (*pb.User).GetId(in)
	err := db.QueryRow("Insert into SHIPPY_user (name,email,password,company) values($1,$2,$3,$4) returning id",
		name, email, password, company).Scan(id)
	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	rows := db.QueryRow("Select * from shippy_user  where user_id = $1", id)
	err := rows.Scan(&user.Name, &user.Password, &user.Email, &user.Company)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]*pb.User, error) {

	// Test cooments for git
	var users []*pb.User
	rows, err := db.Query("Select * from shippy_user")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	

	return users, nil
}

func (r *UserRepository) GetByEmailAndPassword(*pb.User) (*pb.User, error) {
	var user *pb.User
	if err := r.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}
