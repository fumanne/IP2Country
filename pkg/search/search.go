package search

import (
	"fmt"
	"github.com/fumanne/IP2Country/pkg/update"
	"github.com/fumanne/IP2Country/pkg/utils"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Do(ipAddress string) string {
	if utils.IsPrivate(ipAddress) {
		fmt.Printf("Private Address Find: %s\n", ipAddress)
		os.Exit(0)
	}

	if ! utils.IsIP(ipAddress) {
		fmt.Printf("Not A IPaddress Format: %s\n", ipAddress)
		os.Exit(0)
	}
	targetInt := utils.Ip2long(ipAddress)
	ch := make(chan string, 1)
	for _, f := range setFiles(ipAddress) {
		go search(f, targetInt, ch)
	}
	return <-ch
}

func setFiles(IpAddress string) []string {

	if len(findV4File()) == 0 || len(findV6File()) == 0 {
		update.Do()
	}
	if utils.IsIPv4(IpAddress) {
		return findV4File()
	}

	if utils.IsIPv6(IpAddress) {
		return findV6File()
	}
	return nil
}

func findV4File() []string {
	files := []string{}
	Dir := utils.Locate(utils.DOWNLOAD)
	fs, err := ioutil.ReadDir(Dir)
	utils.CheckErr(err)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		ok, _ := regexp.MatchString("^IPv4.*.tsv$", f.Name())
		if ok {
			files = append(files, f.Name())
		}
	}
	return files
}

func findV6File() []string {
	files := []string{}
	Dir := utils.Locate(utils.DOWNLOAD)
	fs, err := ioutil.ReadDir(Dir)
	utils.CheckErr(err)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		ok, _ := regexp.MatchString("^IPv6.*.tsv$", f.Name())
		if ok {
			files = append(files, f.Name())
		}
	}
	return files
}

func search(filename string, target *big.Int, ch chan string) {
	fd, err := os.Open(filepath.Join(utils.Locate(utils.DOWNLOAD), filename))
	utils.CheckErr(err)
	content, _ := ioutil.ReadAll(fd)
	for _, record := range strings.Split(string(content), "\n") {
		//todo: why the end  of reader is none?
		if record != "" {
			ss := strings.Split(record, "\t")
			startInt := utils.Str2BigInt(ss[0])
			endInt := utils.Str2BigInt(ss[1])
			if target.Cmp(startInt) >= 0 && target.Cmp(endInt) == -1 {
				ch <- ss[2]
				break
			}
		}
	}

}
