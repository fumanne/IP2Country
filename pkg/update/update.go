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
	"time"
)

const (
	AFRINIC  = "http://ftp.apnic.net/stats/afrinic/delegated-afrinic-latest"
	APNIC    = "http://ftp.apnic.net/stats/apnic/delegated-apnic-latest"
	LACNIC   = "http://ftp.apnic.net/stats/lacnic/delegated-lacnic-latest"
	RIPENCC  = "http://ftp.apnic.net/stats/ripe-ncc/delegated-ripencc-latest"
	ARIN     = "http://ftp.apnic.net/stats/arin/delegated-arin-extended-latest"
	DOWNLOAD = ".IP2Country"
)

type Region struct {
	name string
	url  string
}

func (r *Region) file() string {
	home := utils.GetHome()
	return filepath.Join(home, DOWNLOAD, r.filename())
}

func (r *Region) filename() string {
	return r.name + ".tsv"
}

func (r *Region) download(w *sync.WaitGroup) {
	defer w.Done()
	response, err := http.Get(r.url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	home := utils.GetHome()
	_d := filepath.Join(home, DOWNLOAD)
	mkdir(_d)

	generate(r.file(), body)
	//if err := ioutil.WriteFile(r.file(), body, 0644); err != nil {
	//	panic(err)
	//}

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

func merge(total string, file ...string) {
	fd, err := os.OpenFile(total, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	for _, f := range file {
		rf, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}
		if _, err:=fd.Write(rf); err != nil {
			panic(err)
		}
	}
}


func generate(total string, s []byte) {
	fd, err := os.OpenFile(total, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	ss := strings.Split(string(s), "\n")
	for _, v := range ss {
		if ! isSkip(v) && isIPFlag(v) {
			s, e, c := utils.ParseIP(v)
			// todo: ugly code
			one := s + "\t" + e + "\t" + c + "\n"
			if _, err := fd.Write([]byte(one)); err != nil {
				panic(err)
			}
		}
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


func isIPFlag(record string)  bool {
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


func Do() {
	wg := &sync.WaitGroup{}
	// todo: ugly code
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
		go r.download(wg)
	}
	//fmt.Println("Before Wait")
	wg.Wait()
	//fmt.Println("ALL Done")
	time.Sleep(time.Second * 3)

	// todo:  ugly code
	home := utils.GetHome()
	_d := filepath.Join(home, DOWNLOAD)
	keys := []string{}
	for k := range elem {
		keys = append(keys, filepath.Join(_d, k+".tsv"))
	}
	merge(filepath.Join(_d, "ip.tsv"), keys...)
}
