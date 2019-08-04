package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/sanjeevchoubey/shippy/user-service/proto/user"
)

func main() {

	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}

	defer db.Close()

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted
	//db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	srv := micro.NewService(
		micro.Name("shippy_service_user"),
		micro.Version("latest"),
	)
	// Init will parse the command line flag
	srv.Init()

	//Register Handler
	pb.RegisterUserServiceHandler(srv.Server(), &UserService{repo})

	// run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
