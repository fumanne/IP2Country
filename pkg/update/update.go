package update

import (
	"github.com/fumanne/IP2Country/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

const (
	AFRINIC = "http://ftp.apnic.net/stats/afrinic/delegated-afrinic-latest"
	APNIC   = "http://ftp.apnic.net/stats/apnic/delegated-apnic-latest"
	LACNIC  = "http://ftp.apnic.net/stats/lacnic/delegated-lacnic-latest"
	RIPENCC = "http://ftp.apnic.net/stats/ripe-ncc/delegated-ripencc-latest"
	ARIN    = "http://ftp.apnic.net/stats/arin/delegated-arin-extended-latest"
)

type Region struct {
	name string
	url  string
}

func (r *Region) file_v4() string {
	return filepath.Join(utils.GetHome(), utils.DOWNLOAD, r.filename_v4())
}

func (r *Region) file_v6() string {
	return filepath.Join(utils.GetHome(), utils.DOWNLOAD, r.filename_v6())
}

func (r *Region) filename_v4() string {
	return utils.IPv4Prefix + r.name + ".tsv"
}

func (r *Region) filename_v6() string {
	return utils.IPv6Prefix + r.name + ".tsv"
}

func (r *Region) stream() []byte {
	response, err := http.Get(r.url)
	defer response.Body.Close()
	utils.CheckErr(err)
	body, err := ioutil.ReadAll(response.Body)
	utils.CheckErr(err)
	return body
}

func (r *Region) generate(wg *sync.WaitGroup) {
	defer wg.Done()
	f4, _ := os.OpenFile(r.file_v4(), os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	f6, _ := os.OpenFile(r.file_v6(), os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer f4.Close()
	defer f6.Close()
	for _, record := range  strings.Split(string(r.stream()), "\n") {
		if ! isSkip(record) && isIPFlag(record) {
			s, e, c := utils.ParseIPInt(record)
			line := s.String() + "\t" + e.String() +  "\t" + c + "\n"
			if isV4record(record) {
				makeFile(f4, line)
			}
			if isV6record(record) {
				makeFile(f6, line)
			}

		}
	}
}

func makeFile(file *os.File, s string) {
	_, err := file.WriteString(s)
	utils.CheckErr(err)
}



func mkdir(d string) {
	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		panic(err)
	}
}

func NewRegion(name, url string) *Region {
	return &Region{
		name: name,
		url:  url,
	}
}


func isSkip(record string) bool {
	words := strings.Split(record, "|")
	if len(words) < 7 {
		return true
	} else {
		return words[6] != "assigned" && words[6] != "allocated"
	}
}

func isIPFlag(record string) bool {
	words := strings.Split(record, "|")
	ok, err := regexp.MatchString("^ipv.*$", words[2])
	if err != nil {
		panic(err)
	}
	if ! ok {
		return false
	}

	return true
}


func isV4record(record string) bool {
	words := strings.Split(record, "|")
	ok, _ := regexp.MatchString("^ipv4$", words[2])
	if ok {
		return true
	} else {
		return false
	}
}

func isV6record(record string) bool {
	words := strings.Split(record, "|")
	ok, _ := regexp.MatchString("^ipv6$", words[2])
	if ok {
		return true
	} else {
		return false
	}
}



func Do() {
	mkdir(utils.Locate(utils.DOWNLOAD))
	wg := &sync.WaitGroup{}
	elem := map[string]string{
		"afrinic": AFRINIC,
		"apnic":   APNIC,
		"arin":    ARIN,
		"lacnic":  LACNIC,
		"ripencc": RIPENCC,
	}
	for k, v := range elem {
		wg.Add(1)
		r := NewRegion(k, v)
		go r.generate(wg)
	}
	wg.Wait()

}
