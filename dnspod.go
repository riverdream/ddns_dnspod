package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	loginToken := flag.String("login_token", "", "format id,token")
	subDomain := flag.String("sub_domain", "www", "sub domain")
	domainID := flag.String("domain_id", "", "domain id")
	recordID := flag.String("record_id", "", "record id")
	flag.Parse()

	if *loginToken == "" || *subDomain == "" || *domainID == "" || *recordID == "" {
		fmt.Println("dnspod -login_token \"id,token\" -sub_domain \"www\" -domain_id \"1234567\" -record_id \"1234567\"")
		return
	}

	v := url.Values{}
	v.Set("login_token", *loginToken)
	v.Set("sub_domain", *subDomain)
	v.Set("format", "json")
	v.Set("record_line", "默认")
	v.Set("domain_id", *domainID)
	v.Set("record_id", *recordID)

	var ip string
	for {
		tmpIP := getIP()
		if tmpIP == ip {
			time.Sleep(5 * time.Minute)
			continue
		}

		ip = tmpIP
		v.Set("value", tmpIP)

		ddns(v.Encode())

		time.Sleep(3 * time.Second)
	}
}

func getIP() string {
	resp, err := http.Get("http://ip.cip.cc")
	if err != nil {
		fmt.Println("http get error:", err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil read error:", err)
		return ""
	}

	return string(body)

}

func ddns(params string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", "https://dnsapi.cn/Record.Ddns", bytes.NewReader([]byte(params)))
	req.Header.Set("UserAgent", "pengjx DDNS Client/1.0.0 (yourmail@yeah.net)")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "text/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client do error:", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil read error:", err)
		return
	}

	fmt.Println(string(body))
}
