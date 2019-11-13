package ac

import (
	"testing"
)

func TestAC_Match(t *testing.T) {

	acAuto := NewAc()
	dirtWorld := []string{
		"好屌",
		"妈的",
		"fuck",
	}
	acAuto.AddWorlds(dirtWorld)
	t.Log(acAuto.Contains("好屌"))

	acAuto.BuildFailurePointer()

	acAuto.Remove("fuck")
	acAuto.BuildFailurePointer()

	text := "妈的-我看他说话的语气，好屌啊 fuck"
	afterText := []rune(text)
	acAuto.Match(text, func(start, end int) {
		//t.Logf("%v [%v-%v]\n", text[start:end+1], start, end)
		for i := start; i <= end; i++ {
			afterText[i] = '*'
		}
	})
	t.Log(string(afterText))
}
