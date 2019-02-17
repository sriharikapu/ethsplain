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
		res := ""
		for str[i] != '\n' {
			res += string(str[i])
			i++
		}
		str = str[i:]
		if len(res) > 100 {
			_, err := hex.DecodeString(strings.TrimSpace(res[2:]))
			if err == nil {
				return strings.TrimSpace(res)
			}
		}
	}
	return ""
}
