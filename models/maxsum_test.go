package models

import (
	"testing"
)

func TestKadanes(t *testing.T){

	//table test
	tests := []struct {
		list []float64
		result int
	}{
		{[]float64{-2,-3,4,-1,-2,1,5,-3},7},
		{[]float64{1,2,3,-2,5},9},
		{[]float64{-1,-2,-3,-4},-1},
		{[]float64{-2,3,5,-1,4,-5}, 11},
	//	{[]float64{2,2,2,2}, 7}, //Proposital error triggers
	}

	for indx, test := range tests {
		result := Kadanes(test.list)
		if result != test.result{
			t.Errorf("Test for Kadanes algorithm failed. Expected %v and got %v. Test Table index %v", test.result, result, indx)
		}
	}

}