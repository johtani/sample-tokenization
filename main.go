package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func main() {

	texts := []string{
		"私はその人を常に先生と呼んでいた。",
		"だからここでもただ先生と書くだけで本名は打ち明けない。",
		"これは世間を憚かる遠慮というよりも、その方が私にとって自然だからである。",
		"素早い茶色の狐が怠けた犬を飛び越えた",
		"すばやいちゃいろのきつねがなまけたいぬをとびこえた",
		"きつね",
		"関西国際空港は、日本の主要な国際空港の一つで、多くの外国人観光客が利用しています。",
		"関西国際空港",
		"隣の客は何をよく食べますか？",
		"客",
		`春の夜の夢はうつつよりもかなしき
	夏の夜の夢はうつつに似たり
	秋の夜の夢はうつつを超え
	冬の夜の夢は心に響く

	山のあなたに小さな村が見える
	川の音が静かに耳に届く
	風が木々を通り抜ける音
	星空の下、すべてが平和である`,
		"素早い茶色の狐が怠けた犬を飛び越えた",
		"すばやいちゃいろのきつねがなまけたいぬをとびこえた",
		"スバヤイチャイロノキツネガナマケタイヌヲトビコエタ",
		"The quick brown fox jumps over the lazy dog",
	}
	var seg gse.Segmenter
	err := seg.LoadDict("ja")
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
		tokenizeByKagome(text, ipa.Dict(), tokenizer.Normal)
		println("------- Kagome ipa mode=Search --------")
		tokenizeByKagome(text, ipa.Dict(), tokenizer.Search)
		println("------- Kagome uni mode=Normal --------")
		tokenizeByKagome(text, uni.Dict(), tokenizer.Normal)
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

func tokenizeByKagome(text string, dict *dict.Dict, mode tokenizer.TokenizeMode) {
	t, err := tokenizer.New(dict)
	if err != nil {
		panic(err)
	}
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
