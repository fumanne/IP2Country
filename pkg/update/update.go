package update

import (
	"database/sql"
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

func (r *Region) download(db *sql.DB, w *sync.WaitGroup) {

	defer w.Done()
	response, err := http.Get(r.url)
	defer response.Body.Close()
	utils.Checkerr(err)

	body, err := ioutil.ReadAll(response.Body)
	utils.Checkerr(err)

	//home := utils.GetHome()
	//_d := filepath.Join(home, DOWNLOAD)
	mkdir(utils.Locate(DOWNLOAD))

	generate(db, body)

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

//func merge(total string, file ...string) {
//	fd, err := os.OpenFile(total, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
//	defer fd.Close()
//	if err != nil {
//		panic(err)
//	}
//	for _, f := range file {
//		rf, err := ioutil.ReadFile(f)
//		if err != nil {
//			panic(err)
//		}
//		if _, err:=fd.Write(rf); err != nil {
//			panic(err)
//		}
//	}
//}

func generate(db *sql.DB, s []byte) {
	//mu.Lock()
	//defer mu.Unlock()
	stmt, err := db.Prepare("Insert into ip2country (start, end, country) values (?, ?, ?);")
	utils.Checkerr(err)
	ss := strings.Split(string(s), "\n")
	for _, v := range ss {
		if ! isSkip(v) && isIPFlag(v) {
			s, e, c := utils.ParseIPInt(v)
			_, err := stmt.Exec(s, e, c)
			utils.Checkerr(err)
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

func Do() {
	wg := &sync.WaitGroup{}
	//mutex := &sync.Mutex{}
	Clean()
	Prepare()
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
		go r.download(DB, wg)
	}
	wg.Wait()
	time.Sleep(time.Second * 3)

	//// todo:  ugly code
	//home := utils.GetHome()
	//_d := filepath.Join(home, DOWNLOAD)
	//keys := []string{}
	//for k := range elem {
	//	keys = append(keys, filepath.Join(_d, k+".tsv"))
	//}
	//merge(filepath.Join(_d, "ip.tsv"), keys...)
}
