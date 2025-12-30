package sum

import ("testing")

func TestSum(t *testing.T){
	numbers := [5]int{2,3,5,1,1}
	sum := Sum(numbers)
	expected := 12
	if sum != expected {
		t.Errorf("sum is %d, but expecteted is %d",sum,expected)

	}
}