import jsonData from "../go/trie/data.json"

type LanguageType = {
    "id": number
    "English": string | null
    "Gloss (english)": string | null
    "Dharug(Gloss)": string | null
    "Dharug": string | null
    "Topic": string | null
    "Image Name (optional)": string | null
    "recording": string | null
    "completed": boolean
}

type TrieNode = {
    children: Map<string, TrieNode>
    isWord: boolean
}

function newTrieNode(): TrieNode {
    return {
        children: new Map<string, TrieNode>(),
        isWord: false
    };
}

class Trie {
    root: TrieNode;

    constructor() {
        this.root = newTrieNode();
    }

    insert(phrase: string): void {
        let node = this.root;

        for (let char of phrase.toLowerCase()) {
            if (!node.children.has(char)) {
                node.children.set(char, newTrieNode());
            }

            node = node.children.get(char)!;
        }

        node.isWord = true;
    }

    prefixOf(phrase: string): string[] {
        const potential: string[] = [];
        let prefix: string = "";
        let node = this.root;

        for (let char of phrase.toLowerCase()) {
            if (node && !node.children.has(char)) {
                return [];
            }

            prefix += char;
            node = node.children.get(char)!;
        }

        this.checkChild(node, prefix, potential)
        console.log(potential)
        return potential;
    }

    private checkChild(node: TrieNode, prefix: string, potential: string[]) {
        if (potential.length > 5) {
            return;
        }

        if (node.isWord) {
            potential.push(prefix);
        }

        for (let char of node.children.keys()) {
            this.checkChild(node.children.get(char), prefix + char, potential);
        }
    }
}

function trieTest(): void {
    const trie = new Trie();
    const data: LanguageType[] = jsonData;

    for (let phrase of data) {
        if (phrase.Topic) {
            trie.insert(phrase.Topic);
        }
        if (phrase.Dharug) {
            trie.insert(phrase.Dharug);
        }
        if (phrase["Dharug(Gloss)"]) {
            trie.insert(phrase["Dharug(Gloss)"]);
        }
        if (phrase.English) {
            trie.insert(phrase.English);
        }
        if (phrase["Gloss (english)"]) {
            trie.insert(phrase["Gloss (english)"]);
        }
    }

    trie.prefixOf("you")
}

trieTest();
