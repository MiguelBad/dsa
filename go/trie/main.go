package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type LanguageData struct {
	Id          int     `json:"id"`
	English     *string `json:"English"`
	EnlighGloss *string `json:"Gloss (english)"`
	DharugGloss *string `json:"Dharug(Gloss)"`
	Dharug      *string `json:"Dharug"`
	Topic       *string `json:"Topic"`
	ImageName   *string `json:"Image Name (optional)"`
	Recording   *string `json:"recording"`
	Completed   bool    `json:"completed"`
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func ReadJson() []LanguageData {
	var data []LanguageData

	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(fmt.Sprintf("Error reading json file\n%v\n", err))
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(fmt.Sprintf("Error decoding json file\n%v\n", err))
	}

	return data
}

func CreateNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

func CreateTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isWord:   false,
		},
	}
}

func (t *Trie) Insert(phrase string) {
	node := t.root

	for _, char := range strings.ToLower(phrase) {
		if _, ok := node.children[char]; !ok {
			node.children[char] = CreateNode()
		}

		node = node.children[char]
	}

	node.isWord = true
}

func (t *Trie) Search(searchPhrase string) bool {
	node := t.root

	for _, char := range strings.ToLower(searchPhrase) {
		if _, ok := node.children[char]; !ok {
			return false
		}

		node = node.children[char]
	}

	return node.isWord
}

func (t *Trie) checkChild(node *TrieNode, prefix []rune, potential *[]string) {
    if potential != nil {
        if len(*potential) == 5 {
            return
        }
    }

	if node.isWord {
		*potential = append(*potential, string(prefix))
	}

	for char, child := range node.children {
		t.checkChild(child, append(prefix, char), potential)
	}
}

func (t *Trie) PrefixOf(searchPrefix string) []string {
	var potential []string
	var prefix []rune
	node := t.root

	for _, char := range strings.ToLower(searchPrefix) {
		prefix = append(prefix, char)

		if _, ok := node.children[char]; !ok {
			return []string{}
		}

		node = node.children[char]
	}

	t.checkChild(node, prefix, &potential)

	for _, p := range potential {
		fmt.Println(p)
	}

	return potential
}

func main() {
	trie := CreateTrie()
	data := ReadJson()

	for _, phrase := range data {
		if phrase.Topic != nil {
			trie.Insert(*phrase.Topic)
		}

		if phrase.English != nil {
			trie.Insert(*phrase.English)
		}

		if phrase.EnlighGloss != nil {
			trie.Insert(*phrase.EnlighGloss)
		}

		if phrase.Dharug != nil {
			trie.Insert(*phrase.Dharug)
		}

		if phrase.DharugGloss != nil {
			trie.Insert(*phrase.DharugGloss)
		}
	}

	trie.PrefixOf("you")
}
