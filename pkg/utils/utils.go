package utils

import (
	"github.com/mitchellh/go-homedir"
	"math/big"
	"net"
	"strconv"
	"strings"
)

func Ip2long(ip string) int64 {
	// todo: ipv6 is not supported now
	ipInt := big.NewInt(0)
	ipAddress := net.ParseIP(ip)
	ipInt.SetBytes(ipAddress.To4())
	return ipInt.Int64()
}

func Long2ip(ipInt int64) string {
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt&0xff), 10)
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

func GetHome() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home
}