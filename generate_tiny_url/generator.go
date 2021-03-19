package generatetinyurl

import (
	"fmt"
	"strconv"
	"strings"
)

type codec struct {
	urls []string
}

func constructor() codec {
	return codec{[]string{}}
}

func (c *codec) encode(longURL string) string {
	c.urls = append(c.urls, longURL)
	return "http://tinyurl.com/" + fmt.Sprintf("%d", len(c.urls)-1)
}

func (c *codec) decode(shortURL string) string {
	tmp := strings.Split(shortURL, "/")
	i, _ := strconv.Atoi(tmp[len(tmp)-1])
	return c.urls[i]
}
