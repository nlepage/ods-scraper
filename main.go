package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var words []string

func main() {
	if err := scrape("https://www.listesdemots.net/touslesmots.htm"); err != nil {
		panic(err)
	}
	for i := 2; i <= 918; i++ {
		if err := scrape(fmt.Sprintf("https://www.listesdemots.net/touslesmotspage%d.htm", i)); err != nil {
			panic(err)
		}
	}
	for _, word := range words {
		fmt.Println(word)
	}
}

func scrape(url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	re, err := regexp.Compile("<span class=mot>([^<]*)")
	if err != nil {
		return err
	}

	m := re.FindSubmatch(b)
	if m == nil {
		return errors.New("no match")
	}

	words = append(words, strings.Split(string(m[1]), " ")...)

	return nil
}
