package dailycodingproblem

import "testing"

/*
Given an array of integers and a number k, where 1 <= k <= length of the array, compute the maximum values of each subarray of length k.

For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:

10 = max(10, 5, 2)
7 = max(5, 2, 7)
8 = max(2, 7, 8)
8 = max(7, 8, 7)
Do this in O(n) time and O(k) space. You can modify the input array in-place and you do not need to store the results. You can simply print them out as you compute them.
*/

func TestDCP18(t *testing.T) {
	tables := []struct {
		arr []int
		k   int
		e   []int
	}{
		{[]int{10, 5, 2, 7, 8, 7}, 3, []int{10, 7, 8, 8}},
	}

	for _, table := range tables {
		res := maxSubArray(table.arr, table.k)
		if len(res) != len(table.e) {
			t.Errorf("Array length was incorrect (%v), got: %d, want: %d.", table.arr, len(res), len(table.e))
		}
		for i := range res {
			if res[i] != table.e[i] {
				t.Errorf("Value was incorrect (%v), got: %d, want: %d.", table.arr, res[i], table.e[i])
			}
		}
	}
}
