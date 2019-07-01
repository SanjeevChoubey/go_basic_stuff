package service

type Employee struct {
	Id          int
	BadgeNumber int
	FirstName   string
	LastName    string
}

var employeeData []*Employee

func GetAllEmployees() ([]*Employee, error) {
	return employeeData, nil
}
