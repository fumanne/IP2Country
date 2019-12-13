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


func TestLong2ip(t *testing.T) {
	longip := int64(15191967)
	v := reflect.TypeOf(Long2ip(longip))
	if v.String() == "string" {
		t.Logf("%d convert ip type is OK", longip)
	}

}