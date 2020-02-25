package utils

import (
	"reflect"
	"testing"
)

func TestIP2long(t *testing.T) {
	ip := "12.12.12.12"
	v := reflect.TypeOf(Ip2long(ip))
	if v.String() == "int64" {
		t.Logf("%s convert long is OK", ip)
	}

}

func TestIsIPv4(t *testing.T) {
	right := "255.255.252.252"
	wrong := "fe80::21b:77ff:fbd6:7860"
	if IsIPv4(right) {
		t.Logf("%s is ipv4", right)
	}

	if IsIPv4(wrong) {
		t.Errorf("%s should not be ipv4", wrong)
	}
}

func TestIsIPv6(t *testing.T) {
	wrong := "255.255.252.252"
	right := "fe80::21b:77ff:fbd6:7860"
	if IsIPv6(wrong) {
		t.Errorf("%s should not be ipv6", wrong)
	}

	if IsIPv6(right) {
		t.Logf("%s is ipv6", right)
	}
}
