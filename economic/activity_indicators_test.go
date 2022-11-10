package economic

import (
	"testing"
)

func TestEquityBbondYieldSpreads(t *testing.T) {
	v, err := EquityBbondYieldSpreads()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}
