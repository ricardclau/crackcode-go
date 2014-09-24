package stringarrays

import (
	"fmt"
	"testing"
)

func TestRotate904x4(t *testing.T) {
	test := [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}

	if rotated := fmt.Sprintf("%v", rotate90Matrix(test)); rotated != "[[12 8 4 0] [13 9 5 1] [14 10 6 2] [15 11 7 3]]" {
		t.Errorf("Rotation did not work quite well, obtained %v", rotated)
	}

}
