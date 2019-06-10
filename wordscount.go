package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/loremipsum.v1"
)

func TopWords(text string, count int) []string {

	if count <= 0 {
		panic("Count is non positive!")
	}

	result := make([]string, 0, count)
	wordcountmap := make(map[string]int)

	CheckWordDelimeter := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c)
	}
	words := strings.FieldsFunc(text, CheckWordDelimeter)
	for _, v := range words {
		wordcountmap[strings.ToLower(v)]++
	}

	type wordcount struct {
		word  string
		count int
	}
	wordcountlist := make([]wordcount, 0, len(wordcountmap))
	for k, v := range wordcountmap {
		wordcountlist = append(wordcountlist, wordcount{k, v})
	}
	sort.Slice(wordcountlist, func(i, j int) bool {
		return wordcountlist[i].count > wordcountlist[j].count
	})

	for i := 0; i < len(wordcountlist) && i < count; i++ {
		result = append(result, wordcountlist[i].word)
	}
	return result

}

func main() {

	loremIpsumGenerator := loremipsum.New()

	text := loremIpsumGenerator.Sentences(100)
	count := 10

	fmt.Println("Text:")
	fmt.Println(text)
	fmt.Printf("Top %v words  by count:\n", count)
	fmt.Println(TopWords(text, count))

}
