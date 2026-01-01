package allsum

import (
	"testing"
    "slices"
)

func TestAllSum(t *testing.T){
	t.Run("testing all sum",func(t *testing.T) {
		got := SumAll([]int {},[]int {1,2})
		want :=[]int{12,3}

		if slices.Equal(got,want) {
			t.Errorf("expexted %q , got %q",want,got)
		}
	})
}