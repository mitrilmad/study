package main

import (
	_ "bytes"
	_ "bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	_"strings"
	"encoding/json"
	_ "bytes"
	"strings"
	"bytes"
)

var (
	seed2 = flag.String("seed", "http://a0.awsstatic.com/pricing/1/ec2/linux-od.min.js", "seed")
)

type AwsPrice struct {
	Vers   float64 `json:"vers"`
	Config struct {
		Rate         string   `json:"rate"`
		ValueColumns []string `json:"valueColumns"`
		Currencies   []string `json:"currencies"`
		Regions      []struct {
			Region        string `json:"region"`
			InstanceTypes []struct {
				Type  string `json:"type"`
				Sizes []struct {
					Size         string `json:"size"`
					VCPU         string `json:"vCPU"`
					ECU          string `json:"ECU"`
					MemoryGiB    string `json:"memoryGiB"`
					StorageGB    string `json:"storageGB"`
					ValueColumns []struct {
						Name   string `json:"name"`
						Prices struct {
							USD string `json:"USD"`
						} `json:"prices"`
					} `json:"valueColumns"`
				} `json:"sizes"`
			} `json:"instanceTypes"`
		} `json:"regions"`
	} `json:"config"`
}

func main() {
	flag.Parse()

	u, err := url.Parse(*seed2)
	if err != nil {
		log.Fatal(err)
	}

	//resp, err := http.Get(u.String())
	//defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	//b := bytes.NewBuffer(body).String()

	//temp[1]

	//fmt.Println(strings.SplitAfter(b, "callback")[1])
	//str := strings.SplitAfter(b,"callback")[1]
	//fmt.Println(str)
	//
	//
	//res := AwsPrice{}
	//json.Unmarshal([]byte(str),&res)
	//fmt.Println(res.Vers)
}
