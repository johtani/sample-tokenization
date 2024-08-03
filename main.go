package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/ikawaha/kagome/tokenizer"
)

func main() {

	texts := []string{
		"私はその人を常に先生と呼んでいた。",
		"だからここでもただ先生と書くだけで本名は打ち明けない。",
		"これは世間を憚かる遠慮というよりも、その方が私にとって自然だからである。",
		"素早い茶色の狐が怠けた犬を飛び越えた",
		"すばやいちゃいろのきつねがなまけたいぬをとびこえた",
	}
	var seg gse.Segmenter
	err := seg.LoadDict("jp")
	if err != nil {
		panic(err)
	}

	for _, text := range texts {
		println("===========================================")
		println("------- GSE cut --------")
		tokenizeByGse(text, seg)
		println("------- GSE cut all --------")
		tokenizeByGseCutAll(text, seg)
		println("------- GSE cut search --------")
		tokenizeByGseCutSearch(text, seg)
		println("------- Kagome ipa mode=Normal --------")
		tokenizeByKagome(text, tokenizer.SysDicIPA(), tokenizer.Normal)
		println("------- Kagome ipa mode=Search --------")
		tokenizeByKagome(text, tokenizer.SysDicIPA(), tokenizer.Search)
		println("------- Kagome uni mode=Normal --------")
		tokenizeByKagome(text, tokenizer.SysDicUni(), tokenizer.Normal)
	}

}

func tokenizeByGse(text string, seg gse.Segmenter) {
	fmt.Println(seg.Cut(text))
}

func tokenizeByGseCutAll(text string, seg gse.Segmenter) {
	fmt.Println(seg.CutAll(text))
}

func tokenizeByGseCutSearch(text string, seg gse.Segmenter) {
	fmt.Println(seg.CutSearch(text))
}

func tokenizeByKagome(text string, dict tokenizer.Dic, mode tokenizer.TokenizeMode) {
	t := tokenizer.NewWithDic(dict)
	tokens := t.Analyze(text, mode)
	print("[")
	for _, token := range tokens {
		if token.Class != tokenizer.DUMMY {
			print(token.Surface)
			print(" ")
		}
	}
	println("]")
}
