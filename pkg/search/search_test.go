package search

import (
	"testing"
)

func TestDo(t *testing.T) {
	ipaddress := "196.223.39.2"
	expected := "GA"

	if Do(ipaddress) == expected {
		t.Logf("ipaddress (%s) is expected\n", ipaddress)
	} else {
		t.Errorf("Not Match")
	}

}
