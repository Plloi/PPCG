package dailycodingproblem

import (
	"testing"
)

/*
This problem was asked by Airbnb.

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?
*/

// TestDCP9 : Runs test for DailyCodingProblem.com Problem 9
func TestDCP9(t *testing.T) {
	tables := []struct {
		li []int
		e  int
	}{
		{[]int{2, 4, 6, 2, 5}, 13},
		{[]int{5, 1, 1, 5}, 10},
	}

	for _, table := range tables {
		res := sumNonAdjacent(table.li)
		if res != table.e {
			t.Errorf("Sum was incorrect (%v), got: %d, want: %d.", table.li, res, table.e)
		}
	}
}
