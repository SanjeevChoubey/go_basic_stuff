package main

import (
	"context"

	pb "github.com/sanjeevchoubey/shippy/user-service/proto/user"
)

type UserService struct {
	repo Repository
}

func (h *UserService) Create(ctx context.Context, in *pb.User, out *pb.Response) error {
	if err := h.repo.Create(in); err != nil {
		return err
	}
	out.User = in
	return nil
}

func (h *UserService) Get(ctx context.Context, in *pb.User, out *pb.Response) error {
	var user *pb.User
	user, err := h.repo.Get(in.Id)
	if err != nil {
		return err
	}
	out.User = user
	return nil
}

func (h *UserService) GetAll(ctx context.Context, in *pb.User, out *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	out.Users = users
	return nil
}

func (h *UserService) Auth(ctx context.Context, in *pb.User, out *pb.Token) error {

	_, err := h.repo.GetByEmailAndPassword(in)
	if err != nil {
		return err
	}
	out.Token = "Just Testing"
	return nil
}

func (h *UserService) ValidateToken(ctx context.Context, in *pb.Token, out *pb.Token) error {
	return nil
}
