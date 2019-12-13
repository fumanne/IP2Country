package update

import (
	"testing"
)

func TestIsSkip(t *testing.T) {
	correct1 := "a|b|c|d|e|f|assigned|xxxx"
	correct2 := "a|b|c|d|e|f|allocated|yyyy"
	wrong1 := "a||b|C"
	wrong2 := "a|b|c|d|e|f|Test|yyyy"

	if isSkip(correct1) {
		t.Errorf("%s expected not skip\n", correct1)
	}

	if isSkip(correct2) {
		t.Errorf("%s expected not skip\n", correct2)
	}

	if isSkip(wrong1) {
		t.Logf("%s expected skip\n", wrong1)
	}

	if isSkip(wrong2) {
		t.Logf("%s expected skip\n", wrong2)
	}
}


func TestIsIPFlag(t *testing.T) {
	correct := "afrinic|GA|ipv4|196.223.39.0|256|20140923|assigned"
	wrong := "afrinic|GA|Now|196.223.39.0|256|20140923|assigned"

	if isIPFlag(correct) {
		t.Logf("%s is a IPFlag\n", correct)
	}

	if isIPFlag(wrong) {
		t.Errorf("%s expect is not a IPFlag\n", wrong)
	}

}



