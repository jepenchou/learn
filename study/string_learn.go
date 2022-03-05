package study

import (
	"fmt"
	"strconv"
)

// first demo
func main() {
	fmt.Println(addBinary("11", "10"))
}
func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)

	n := lenA
	if n < lenB {
		n = lenB
	}
	res := ""
	t := 0
	for i := 0; i < n; i++ {
		if i < lenA {
			t += (int)(a[lenA-i-1] - '0')
		}
		if i < lenB {
			t += (int)(b[lenB-i-1] - '0')
		}
		res = strconv.Itoa(t%2) + res
		if t >= 2 {
			t = 1
		} else {
			t = 0
		}
	}
	if t == 1 {
		res = "1" + res
	}
	return res
}
