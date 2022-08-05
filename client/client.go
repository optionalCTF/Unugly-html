package client

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/yosssi/gohtml"
)

var (
	lmao string
)

func Pull(url string, path string) {
	fmt.Printf("[+] Pulling HTML from: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	misc, _ := ioutil.ReadAll(res.Body)

	lmao = gohtml.Format(string(misc))
	fmt.Println(lmao)

	if path != "" {
		fileWrite(path)
	}
}

func fileWrite(path string) error {
	if _, err := os.Stat(path); err == nil {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("[-] %s", err.Error())
		}
		defer f.Close()

		if _, err := f.WriteString(lmao); err != nil {
			fmt.Println(err.Error())
		}
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintln(w, lmao)
		return w.Flush()
	} else {
		fmt.Printf("Schrodingers error... How did you do this?")
		return nil
	}
}
