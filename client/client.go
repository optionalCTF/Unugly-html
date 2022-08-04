package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yosssi/gohtml"
)

func Pull(url string) {
	fmt.Printf("[+] Pulling HTML from: %s", url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println(gohtml.FormatWithLineNo(string(data)))
}
