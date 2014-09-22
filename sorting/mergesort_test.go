package sorting

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	test := []int{9, 1, 2, 4, 5, 3, 7, 8, 2, 2, 10, 1, 12}

	if sorted := fmt.Sprintf("%v", MergeSort(test)); sorted != "[1 1 2 2 2 3 4 5 7 8 9 10 12]" {
		t.Errorf("Mergesort did not work quite well, obtained %v", sorted)
	}
}
