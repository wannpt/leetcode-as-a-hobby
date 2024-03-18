package zigzag

import (
	"testing"
)

func Test_zigzag(t *testing.T) {
	expect := "PINALSIGYAHRPI"
	actual := Convert("PAYPALISHIRING", 4)

	if expect != actual {
		t.Errorf("expected %v, got %v", expect, actual)
	}
}
