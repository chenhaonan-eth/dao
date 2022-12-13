package utils

import "testing"

func TestLastDayOfLastMonth(t *testing.T) {
	s := LastDayOfLastMonth()
	t.Log(s)
}
