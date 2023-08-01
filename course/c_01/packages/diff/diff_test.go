package diff_test

import (
	"github.com/dev2019zheng/MyGoStudy/course/c_01/packages/diff"
	"testing"
)

func TestDiff(t *testing.T) {
	if diff.Diff() != 1 {
		t.Error("Diff() != 1")
	}
}
