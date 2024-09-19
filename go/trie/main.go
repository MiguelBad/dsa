package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataType struct {
	ID           int     `json:"id"`
	English      *string `json:"English"`
	GlossEnglish *string `json:"Gloss (english)"`
	GlossDharug  *string `json:"Dharug(Gloss)"`
	Dharug       *string `json:"Dharug"`
	Topic        *string `json:"Topic"`
	ImageName    *string `json:"Image Name (optional)"`
	Recording    *string `json:"recording"`
	Completed    bool    `json:"completed"`
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type TrieTree struct {
	root *TrieNode
}

func readJson() []DataType {
	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(fmt.Sprintf("Error opening json file\n%v\n", err))
	}

	var data []DataType
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(fmt.Sprintf("Error decoding json file\n%v\n", err))
	}

	return data
}

func (t *TrieTree) Insert(phrase string) {
	node := t.root

	test := &TrieNode{children: make(map[rune]*TrieNode), isWord: false}
	node.children[200] = test

	for _, char := range phrase {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isWord:   false,
			}
		}

		node = node.children[char]
	}

	node.isWord = true
}

func (t *TrieTree) Search(phrase string) bool {
	node := t.root

	for _, char := range phrase {
		if _, ok := node.children[char]; !ok {
			return false
		}

        node = node.children[char]
        fmt.Println(string(char))
	}

    return node.isWord
}

func main() {
	trie := &TrieTree{root: &TrieNode{children: make(map[rune]*TrieNode), isWord: false}}
	data := readJson()

	for _, phrase := range data {
		if phrase.English != nil {
			trie.Insert(*phrase.English)
		}

		if phrase.Dharug != nil {
			trie.Insert(*phrase.Dharug)
		}

		if phrase.GlossEnglish != nil {
			trie.Insert(*phrase.GlossEnglish)
		}

		if phrase.GlossDharug != nil {
			trie.Insert(*phrase.GlossDharug)
		}
	}

    fmt.Println(trie.Search("you"))

}
