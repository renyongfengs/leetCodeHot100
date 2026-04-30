package hash

import (
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	result := longestConsecutive(nums)
	if result != expected {
		t.Errorf("Test failed: expected %d, got %d", expected, result)
	}
}
