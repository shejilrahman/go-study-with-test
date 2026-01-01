package sum

import (
	"testing"
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
}