package utils

import (
	"github.com/mitchellh/go-homedir"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const DOWNLOAD = ".IP2Country"

var DBFile = filepath.Join(Locate(DOWNLOAD), "ip.db")

var IPv4Regexp = "^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
var IPV6Regexp = `^((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|` +
	`(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
	`(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
	`(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))$`

func Ip2long(ip string) int64 {
	ipInt := big.NewInt(0)
	ipAddress := net.ParseIP(ip)
	// todo: need optimize
	if IsIPv4(ipAddress) {
		ipInt.SetBytes(ipAddress.To4())
	}
	if IsIPv6(ipAddress) {
		ipInt.SetBytes(ipAddress.To16())
	}
	return ipInt.Int64()
}

func Long2ip(ipInt int64) string {
	// todo: ipv6 is not supported
	// help to offer one function which can convert int to ipv6 address string
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func ParseIP(record string) (string, string, string) {
	words := strings.Split(record, "|")
	country := words[1]
	start_ip := words[3]
	length, _ := strconv.Atoi(words[4])
	end_ip := Long2ip(Ip2long(start_ip) + int64(length) - 1)
	return start_ip, end_ip, country
}

func ParseIPInt(record string) (int64, int64, string) {
	words := strings.Split(record, "|")
	country := words[1]
	startInt := Ip2long(words[3])
	length, _ := strconv.Atoi(words[4])
	endInt := Ip2long(words[3]) + int64(length) - 1
	return startInt, endInt, country

}

func IsIPv4(ip net.IP) bool {
	ok, _ := regexp.MatchString(IPv4Regexp, ip.String())
	if ok {
		return true
	} else {
		return false
	}

}

func IsIPv6(ip net.IP) bool {
	ok, _ := regexp.MatchString(IPV6Regexp, ip.String())
	if ok {
		return true
	} else {
		return false
	}
}

func GetHome() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home
}

func Locate(str string) string {
	home := GetHome()
	return filepath.Join(home, str)
}

func Checkerr(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func IsExist(file string) bool {
	if _, err:= os.Stat(file); err != nil {
		return false
	}
	return true
}