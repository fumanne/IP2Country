package utils

import (
	"github.com/mitchellh/go-homedir"
	"math/big"
	"net"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	DOWNLOAD   = ".IP2Country"
	IPv4Prefix = "IPv4_"
	IPv6Prefix = "IPv6_"
)


var PrivateIPRegexp = "^10.*$|^172.16.*$|^192.168.*$|^127.*$"


var IPv4Regexp = "^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
var IPV6Regexp = `^((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|` +
	`(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
	`(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
	`(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
	`(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))$`

func Ip2long(ip string) *big.Int {
	ipInt := big.NewInt(0)
	ipAddress := net.ParseIP(ip)
	if IsIPv4(ip) {
		ipInt.SetBytes(ipAddress.To4())
	}
	if IsIPv6(ip) {
		ipInt.SetBytes(ipAddress.To16())
	}
	return ipInt
}

//func Long2ip(ipInt int64) string {
//	// todo: ipv6 is not supported
//	// help to offer one function which can convert int to ipv6 address string
//	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
//	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
//	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
//	b3 := strconv.FormatInt((ipInt & 0xff), 10)
//	return b0 + "." + b1 + "." + b2 + "." + b3
//}

func Str2BigInt(ipint string) *big.Int {
	x := big.NewInt(0)
	x, ok := x.SetString(ipint, 10)
	if ! ok {
		panic("Set Ip to Big Int Error")
	}
	return x
}

func ParseIPInt(record string) (*big.Int, *big.Int, string) {
	words := strings.Split(record, "|")
	country := words[1]
	startInt := Ip2long(words[3])
	length, _ := strconv.Atoi(words[4])
	endInt := big.NewInt(0)
	offset := big.NewInt(0)
	offset = offset.Sub(big.NewInt(int64(length)), big.NewInt(1))
	endInt = endInt.Add(Ip2long(words[3]), offset)
	return startInt, endInt, country

}

func IsIPv4(ip string) bool {
	ok, _ := regexp.MatchString(IPv4Regexp, ip)
	if ok {
		return true
	} else {
		return false
	}
}

func IsIPv6(ip string) bool {
	ok, _ := regexp.MatchString(IPV6Regexp, ip)
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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func IsPrivate(ip string) bool {
	if IsIPv4(ip) {
		if ok, _ := regexp.MatchString(PrivateIPRegexp, ip); ok {
			return true
		}
	}
	return false
}

func IsIP(ip string) bool {
	if ! IsIPv4(ip) && ! IsIPv6(ip) {
		return false
	}
	return true
}