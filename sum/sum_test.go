package sum

import (
	"testing"
	"slices"
)

func TestSum(t *testing.T){
	// numbers := [5]int{2,3,5,1,1}
	// sum := Sum(numbers)
	// expected := 12
	// if sum != expected {
	// 	t.Errorf("sum is %d, but expecteted is %d",sum,expected)

	// }
	t.Run("for fixed numbers sum",func(t *testing.T){
			numbers := []int{2,3,5,1,1}
			sum := Sum(numbers)
			expected := 12
			if sum != expected {
				t.Errorf("sum is %d, but expecteted is %d",sum,expected)

			}
	})

	t.Run("for any numbers sum",func (t *testing.T)  {
		numbers := []int{3,5,8,4}
		sum := Sum(numbers)
			expected := 20
			if sum != expected {
				t.Errorf("sum is %d, but expecteted is %d",sum,expected)

			}
	})

	t.Run("testing all sum",func(t *testing.T) {
		got := SumAll([]int {3,9},[]int {1,2})
		want :=[]int{12,3}

		if !slices.Equal(got,want) {
			t.Errorf("expexted %v , got %v",want,got)
		}
	})

	t.Run("testing all tails",func(t *testing.T) {
		got := SumAllTails([]int {},[]int {1,2,5})
		want := []int{0,7}
		if !slices.Equal(got,want) {
			t.Errorf("expexted %v , got %v",want,got)
		}
	})
}