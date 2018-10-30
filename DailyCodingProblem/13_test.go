package dailycodingproblem

import (
	"testing"
)

/*
This problem was asked by Amazon.

Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.

For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb", so your function should return 3.
*/

// TestDCP13 : Runs test for DailyCodingProblem.com Problem 15
func TestDCP13(t *testing.T) {
	tables := []struct {
		k int
		s string
		e string
	}{
		{2, "abcba", "bcb"},
	}

	for _, table := range tables {
		res := find(table.k, table.s)
		if res != table.e {
			t.Errorf("Substring was incorrect (%d::%s), got: %s, want: %s.", table.k, table.s, res, table.e)
		}
	}
}
