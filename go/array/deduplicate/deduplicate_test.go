package deduplicate

import (
	"reflect"
	"testing"
)

type Case struct {
	Inputs   []int
	Expected int
}

func Test1(t *testing.T) {
	cases := make([]Case, 5)
	cases[0] = Case{[]int{1}, 1}
	cases[1] = Case{[]int{1, 1}, 1}
	cases[2] = Case{[]int{1, 2}, 2}
	cases[3] = Case{[]int{1, 1, 2, 2, 3}, 3}
	cases[4] = Case{[]int{1, 2, 2, 2, 2, 3, 3, 3, 4}, 4}

	for _, testCase := range cases {
		nums := testCase.Inputs
		result := removeDuplicates(nums)
		expected := testCase.Expected

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v; got %v", expected, nums)
		}

		// Make sure the sequence is always rising
		for n := 1; n < result; n++ {
			if nums[n-1] >= nums[n] {
				t.Errorf("Expected %v; to be greater than %v", nums[n], nums[n-1])
			}
		}
	}
}

func Benchmark1(b *testing.B) {
	cases := make([]Case, 5)
	cases[0] = Case{[]int{1}, 1}
	cases[1] = Case{[]int{1, 1}, 1}
	cases[2] = Case{[]int{1, 2}, 2}
	cases[3] = Case{[]int{1, 1, 2, 2, 3}, 3}
	cases[4] = Case{[]int{1, 2, 2, 2, 2, 3, 3, 3, 4}, 4}

	for n := 0; n < b.N; n++ {
		for _, testCase := range cases {
			removeDuplicates(testCase.Inputs)
		}

	}
}
