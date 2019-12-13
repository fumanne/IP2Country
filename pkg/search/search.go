package search

import (
	"IP2Country/pkg/update"
	"IP2Country/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Do(ipaddress string) string {
	home := utils.GetHome()
	iptsv := filepath.Join(home, update.DOWNLOAD, "ip.tsv")
	if _, err := os.Stat(iptsv); os.IsNotExist(err) {
		update.Do()
	}

	res, err := ioutil.ReadFile(iptsv)
	if err != nil {
		panic(err)
	}

	target := utils.Ip2long(ipaddress)
	country := ""
	for _, record := range strings.Split(string(res), "\n") {
		sr := strings.Split(record, "\t")
		if len(sr) != 3 {
			continue
		}
		start := utils.Ip2long(sr[0])
		end := utils.Ip2long(sr[1])
		if target >= start && target <= end {
			country = sr[2]
			break
		}
	}
	return country
}
