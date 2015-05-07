package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Result struct {
	Basic struct {
		Explains []string
		phonetic string
	}
	ErrorCode   int
	Query       string
	Translation []string
	Web         []struct {
		Key   string
		Value []string
	}
}

func translate(word string) {

	base_url := "http://fanyi.youdao.com/openapi.do?" +
		"keyfrom=smarterhjw&" +
		"key=1800385601&" +
		"type=data&" +
		"doctype=json&" +
		"version=1.1&" +
		"q="

	queryURL := buildQueryURL(base_url, word)

	res, _ := http.Get(queryURL)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	decodeJson(body)

}

func buildQueryURL(base_url string, word string) string {
	u, err := url.Parse(base_url)
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "http"
	u.Host = "fanyi.youdao.com"
	q := u.Query()
	q.Set("q", word)
	u.RawQuery = q.Encode()
	return u.String()
}

func decodeJson(res []byte) {

	var r Result

	err := json.Unmarshal(res, &r)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%v", r)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter only one word to translate")
		os.Exit(1)
	}

	translate(args[1])
}
