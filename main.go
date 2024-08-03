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
		tokenizeByKagome(text, tokenizer.Normal)
		println("------- Kagome ipa mode=Search --------")
		tokenizeByKagome(text, tokenizer.Search)
	}

}

func tokenizeByGse(text string, seg gse.Segmenter) error {
	fmt.Println(seg.Cut(text))
	return nil
}

func tokenizeByGseCutAll(text string, seg gse.Segmenter) error {
	fmt.Println(seg.CutAll(text))
	return nil
}

func tokenizeByGseCutSearch(text string, seg gse.Segmenter) error {
	fmt.Println(seg.CutSearch(text))
	return nil
}

func tokenizeByKagome(text string, mode tokenizer.TokenizeMode) error {
	t := tokenizer.New()
	tokens := t.Analyze(text, mode)
	print("[")
	for _, token := range tokens {
		if token.Class != tokenizer.DUMMY {
			print(token.Surface)
			print(" ")
		}
	}
	println("]")

	return nil
}
