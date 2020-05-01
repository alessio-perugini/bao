package bao

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	geoIpDb        = "GeoLite2-City.mmdb"
	linuxLog       = "./out/failed.log"
	onlyIpFile     = "./out/blacklist.log"
	detailedIpFile = "./out/detailed.json"
	listIp         []IpInfo
)

type IpInfo struct {
	Ip      string
	Country *geoip2.Country
}

func IsIpv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

func GeoIpSearch(ip string, db *geoip2.Reader) bool {
	parsedIp := net.ParseIP(ip)
	record, err := db.Country(parsedIp)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Printf("[%s] nation: %s \n", ip, record.Country.IsoCode)

	infoIp := IpInfo{Ip: ip, Country: record}
	listIp = append(listIp, infoIp)
	return record.Country.IsoCode != "IT" //TODO map?
}

func GetIpFromLog() {
	file, err := os.Open(linuxLog) //Read raw linux log
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(onlyIpFile)
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	db, err := geoip2.Open(geoIpDb) //Open geoip for filtering purpose
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ipVisited := make(map[string]bool, 1000) //used to guarantee duplicate-free ip's
	rawLogScanner := bufio.NewScanner(file)

	for rawLogScanner.Scan() {
		ip := ExtractNIpFromString(rawLogScanner.Text(), 1)

		for i := 0; i < len(ip); i++ {
			if !ipVisited[ip[i]] {
				ipVisited[ip[i]] = true
				if GeoIpSearch(ip[i], db) {
					f.WriteString(ip[i] + "\n")
				}
			}
		}
	}

	WriteObjectToJsonFile()

	fmt.Println("done")
}

func ExtractNIpFromString(v string, n int) []string {
	rawLogParts := strings.Split(v, " ")
	ipSlice := make([]string, 2)
	var wg sync.WaitGroup
	maxIp := 0
	valueLock := &sync.Mutex{}

	for i := 0; i < len(rawLogParts); i++ {
		wg.Add(1)

		go func(part string, arr *[]string) {
			defer wg.Done()
			valueLock.Lock()
			defer valueLock.Unlock()

			if maxIp == n {
				return
			}

			if IsIpv4Net(part) {
				*arr = append(*arr, part)
				maxIp++
				return
			}
		}(rawLogParts[i], &ipSlice)
	}

	wg.Wait()

	return ipSlice
}

func WriteObjectToJsonFile() {
	fileDetail, err := os.Create(detailedIpFile)
	defer func() {
		err := fileDetail.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(fileDetail)
	err = encoder.Encode(listIp)
	if err != nil {
		log.Fatal(err)
	}
}
