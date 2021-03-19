package generatetinyurl

import (
	"testing"
)

var urlset []string

func Test_codec_encode(t *testing.T) {
	table := map[string]string{
		"api/v2/customer": "http://tinyurl.com/0",
	}
	codec := constructor()
	for k, v := range table {
		if result := codec.encode(k); result != v {
			t.Errorf("testing: %s, want: %v, got: %v", k, v, result)
		}
	}
}
