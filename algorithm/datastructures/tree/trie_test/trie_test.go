package trie

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func getNames() []string {
	file, err := os.Open("names.txt")
	if err != nil {
		panic(err)
	}
	var ret []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t := scanner.Text()
		ret = append(ret, t)
	}
	return ret
}

func TestTrie_Add(t *testing.T) {
	tr := NewTrie()
	names := getNames()
	for _, name := range names {
		if len(name) > 0 {
			tr.Add(name, name)
		}
	}
	fmt.Println("数量:", tr.Size())
	fmt.Println(tr.SearchPrefix("约翰"))
}
