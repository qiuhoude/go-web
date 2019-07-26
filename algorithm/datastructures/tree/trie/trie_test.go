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

func TestTrie(t *testing.T) {
	tr := NewTrie()
	names := getNames()
	for _, name := range names {
		if len(name) > 0 {
			tr.Add(name, name)
		}
	}
	fmt.Println("数量:", tr.Size())
	fmt.Println(tr.SearchPrefix("约翰"))
	tr.Remove("约翰顿")
	fmt.Println(tr.SearchPrefix("约翰"))
}

func TestTrie_Remove(t *testing.T) {
	tr := NewTrie()
	tr.Add("小", "小")
	tr.Add("小米啦", "小米啦")
	tr.Add("小洪啦", "小洪啦")
	fmt.Println(tr.SearchPrefix("小"))
	tr.Remove("小")
	fmt.Println(tr.SearchPrefix("小"))
}
