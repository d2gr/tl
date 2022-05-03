package iter

import (
	"fmt"
	"testing"
)

func TestWindow(t *testing.T) {
	iter := ToSlice(WindowCopy(Slice([]int{1, 2, 3, 4, 5, 6, 7}), 2))

	fmt.Println(iter)
}
