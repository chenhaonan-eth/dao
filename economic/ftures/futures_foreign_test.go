package ftures

import (
	"testing"
)

func TestFuturesForeignHist(t *testing.T) {
	v, err := FuturesForeignHist("CAD")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", v)
}
