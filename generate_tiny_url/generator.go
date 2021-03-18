package generatetinyurl

import (
	"strconv"
	"strings"
)

type Codec struct {
	urls []string
}

func Constructor() Codec {
	return Codec{[]string{}}
}

func (this *Codec) encode(longUrl string) string {
	this.urls = append(this.urls, longUrl)
	return "http://tinyurl.com/" + string(len(this.urls)-1)
}

func (this *Codec) decode(shortUrl string) string {
	tmp := strings.Split(shortUrl, "/")
	i, _ := strconv.Atoi(tmp[len(tmp)-1])
	return this.urls[i]
}
