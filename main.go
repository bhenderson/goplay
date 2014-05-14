package main

import (
	"fmt"
	term "github.com/bhenderson/terminalgo"
	"io/ioutil"
	"net/http"
	"os"
)

// Inspired from a play.golang I found online, but I can't find it at the moment.
func main() {
	url := "http://play.golang.org"
	resp, err := http.Post(url+"/share", "text/plain", os.Stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	body := string(b)

	if resp.StatusCode != 200 {
		fmt.Println(body)
		os.Exit(1)
	}

	end := ""
	if term.IsTerminal(os.Stdout.Fd()) {
		end = "\n"
	}

	fmt.Printf("%s/p/%s%s", url, body, end)
}
