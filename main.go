package main

import (
	"fmt"
	"github.com/go-ego/gse"
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
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
		"早稲田大学",
		"早稲田大学（基幹理工学部・創造理工学部・先進理工学部） (2021年版大学入試シリーズ)",
		"ｍillefiori",
		"ｺｽﾄｺ",
		"サーキユﾚｰﾀｰ",
		"ペンギン",
		"㌔",
		"\u30DB\u309A",
	}
	var seg gse.Segmenter
	err := seg.LoadDict("ja")
	if err != nil {
		panic(err)
	}
	normalizer := transform.Chain(width.Fold, norm.NFC)

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
		println("------- text/unicode/norm string --------")
		normalize(text)
		println("------- text/width string --------")
		folding(text)
		println("------- Fold + NFC string --------")
		transforming(normalizer, text)
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

func transforming(t transform.Transformer, text string) {
	result, n, err := transform.String(t, text)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("[%s]/ n:[%d] len[%d] runes[%d]", result, n, len(result), len([]rune(result))))
}

func folding(text string) {
	widen := width.Widen.String(text)
	narrow := width.Narrow.String(text)
	fold := width.Fold.String(text)
	fmt.Print(fmt.Sprintf("Widen[%s]\nNarrow:[%s]\nFold:[%s]", widen, narrow, fold))
	fmt.Println(fmt.Sprintf("runes:[%d, %d, %d]", len([]rune(widen)), len([]rune(narrow)), len([]rune(fold))))
}

func normalize(text string) {
	fmt.Println(fmt.Sprintf("orig: len[%d], runes[%d]", len(text), len([]rune(text))))
	fmt.Println(fmt.Sprintf("NFKC:[%s]\nNFC :[%s]\nNFD :[%s],\nNFKD:[%s]", norm.NFKC.String(text), norm.NFC.String(text), norm.NFD.String(text), norm.NFKD.String(text)))
	normalized := norm.NFKC.String(text)
	println(normalized)
	fmt.Println(fmt.Sprintf("norm: len[%d], runes[%d]", len(normalized), len([]rune(normalized))))
}
