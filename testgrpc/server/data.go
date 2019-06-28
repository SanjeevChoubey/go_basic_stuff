package server

import (
	"testgrpc/pb/message"
)

var employees = []pb.Employee{
	pb.Employee{
		Id:           1,
		BandgeNumber: 2080,
		Firstname:    "Sanjeev",
		Lastname:     "Choubey",
	},

	pb.Employee{
		Id:           2,
		BandgeNumber: 2081,
		Firstname:    "kanchan",
		Lastname:     "Choubey",
	},
	pb.Employee{
		Id:           3,
		BandgeNumber: 2082,
		Firstname:    "Sangam",
		Lastname:     "Choubey",
	},
	pb.Employee{
		Id:           4,
		BandgeNumber: 2083,
		Firstname:    "Samik",
		Lastname:     "Choubey",
	},
}
