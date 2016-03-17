package qsort

import "testing"

func TestQsort( t *testing.T ) {
	values := []int{ 2, 1, 5, 8, 3 }
	Qsort( values )
	if values[0]!=1 || values[1]!=2 || values[2]!=3 || values[3]!=5 || values[4]!=8 {
		t.Error( "Bubble sort error, Got ", values, "Except 1, 2, 3, 5, 8" )
	}
}