package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
		`)
	re := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	var res []byte
	for _, s := range re.FindAllSubmatchIndex(src, -1) {
		res = re.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
