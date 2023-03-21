package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SearchResult struct {
	Result struct {
		Hits []struct {
			IP string `json:"ip"`
		} `json:"hits"`
	} `json:"result"`
}

func main() {

	search := flag.String("search", "", "the search query")
	flag.Parse()

	// ensure that the search flag is set
	if *search == "" {
		fmt.Println("missing required flag: -search")
		flag.Usage()
		os.Exit(1)
	}

	url := fmt.Sprintf("https://search.censys.io/api/v2/hosts/search?q=(%s)&per_page=100&virtual_hosts=INCLUDE", *search)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Basic MYKEY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var searchResult SearchResult
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		log.Fatal(err)
	}

	for _, hit := range searchResult.Result.Hits {
		fmt.Println(hit.IP)
	}

}
