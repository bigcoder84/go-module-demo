package httputil

import (
	"fmt"
	"testing"
)

func Test_get(t *testing.T) {
	code, body := get("https://blog.bigcoder.cn")
	fmt.Println(code)
	fmt.Println(body)
}
