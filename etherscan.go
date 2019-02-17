package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// etherscan doesn't have an endpoint for raw tx's so we need to crawl their website
func etherscanCrawlRaw(tx string) string {

	resp, err := http.Get("https://etherscan.io/getRawTx?tx=" + tx)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	str := string(body)
	//arr := make([]string, 0)

	for strings.Index(str, "0x") > 0 {
		i := strings.Index(str, "0x")
		//fmt.Println("i:", i)
		res := ""
		for str[i] != '\n' {
			res += string(str[i])
			//fmt.Println(res)
			i++
		}
		//fmt.Println(res)
		str = str[i:]
		//arr = append(arr, res)
		if len(res) > 100 && len(res) < 500 {
			_, err := hex.DecodeString(strings.TrimSpace(res[2:]))
			if err == nil {
				return strings.TrimSpace(res)
			}
		}
	}
	return ""
}
