package main

import (
	"github.com/SanjeevChoubey/testgrpc/pb"
)

var employees = []pb.Employee{
	pb.Employee{
		Id:          1,
		BadgeNumber: 2080,
		Firstname:   "Sanjeev",
		Lastname:    "Choubey",
	},

	pb.Employee{
		Id:          2,
		BadgeNumber: 2081,
		Firstname:   "kanchan",
		Lastname:    "Choubey",
	},
	pb.Employee{
		Id:          3,
		BadgeNumber: 2082,
		Firstname:   "Sangam",
		Lastname:    "Choubey",
	},
	pb.Employee{
		Id:          4,
		BadgeNumber: 2083,
		Firstname:   "Samik",
		Lastname:    "Choubey",
	},
}
