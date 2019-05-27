package creational_test

import (
	"testing"

)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer after calling getInstance() is nil")
	}
	expectedcounter := counter1

	currentcount := counter1.AddOne()

	if currentcount != 1 {
		t.Errorf("After calling first time counter should be 1 but is is %d", currentcount)
	}

	// if user call again  the get instance he should get the same instance
	counter2 := GetInstance()
	if counter2 != expectedcounter {
		t.Errorf("Expected same instance but got the different one")
	}

	// Check for count is increasing correctly with correct instance
	currentcount = counter1.AddOne()
	if currentcount != 2 {
		t.Errorf("Counter is not increamenting correctly, count should be 2 at this point")
	}
}
