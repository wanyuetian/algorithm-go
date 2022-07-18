package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Interface interface {
	Insert(word string)
	Delete(word string)
	Search(word string) int
	PrefixCount(word string) int
}

type Node struct {
	Pass      int
	End       int
	NextNodes []*Node
}

func NewNode() *Node {
	return &Node{
		NextNodes: make([]*Node, 26),
	}
}

type TrieTree struct {
	Root *Node
}

func NewTrieTree() *TrieTree {
	return &TrieTree{Root: NewNode()}
}

func (t *TrieTree) Insert(word string) {
	curNode := t.Root
	curNode.Pass++
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if curNode.NextNodes[path] == nil {
			curNode.NextNodes[path] = NewNode()
		}
		curNode = curNode.NextNodes[path]
		curNode.Pass++
	}
	curNode.End++
}

func (t *TrieTree) Delete(word string) {
	if t.Search(word) != 0 {
		curNode := t.Root
		curNode.Pass--
		for i := 0; i < len(word); i++ {
			path := word[i] - 'a'
			curNode.NextNodes[path].Pass--
			if curNode.NextNodes[path].Pass == 0 {
				curNode.NextNodes[path] = nil
				return
			}
			curNode = curNode.NextNodes[path]
		}
		curNode.End--
	}
}

func (t *TrieTree) Search(word string) int {
	curNode := t.Root
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if curNode.NextNodes[path] == nil {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.Pass
}

func (t *TrieTree) PrefixCount(prefix string) int {
	curNode := t.Root
	for i := 0; i < len(prefix); i++ {
		path := prefix[i] - 'a'
		if curNode.NextNodes[path] == nil {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.Pass
}

type Node2 struct {
	Pass      int
	End       int
	NextNodes map[uint8]*Node2
}

func NewNode2() *Node2 {
	return &Node2{
		NextNodes: make(map[uint8]*Node2),
	}
}

type TrieTree2 struct {
	Root *Node2
}

func (t *TrieTree2) Insert(word string) {
	curNode := t.Root
	curNode.Pass++
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if _, ok := curNode.NextNodes[path]; !ok {
			curNode.NextNodes[path] = NewNode2()
		}
		curNode = curNode.NextNodes[path]
		curNode.Pass++
	}
	curNode.End++
}

func (t *TrieTree2) Delete(word string) {
	if t.Search(word) != 0 {
		curNode := t.Root
		curNode.Pass--
		for i := 0; i < len(word); i++ {
			path := word[i] - 'a'
			curNode.NextNodes[path].Pass--
			if curNode.NextNodes[path].Pass == 0 {
				delete(curNode.NextNodes, path)
				return
			}
			curNode = curNode.NextNodes[path]
		}
		curNode.End--
	}

}

func (t *TrieTree2) Search(word string) int {
	curNode := t.Root
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if _, ok := curNode.NextNodes[path]; !ok {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.End
}

func (t *TrieTree2) PrefixCount(word string) int {
	curNode := t.Root
	for i := 0; i < len(word); i++ {
		path := word[i] - 'a'
		if _, ok := curNode.NextNodes[path]; !ok {
			return 0
		}
		curNode = curNode.NextNodes[path]
	}
	return curNode.Pass
}

func NewTrieTree2() *TrieTree2 {
	return &TrieTree2{
		Root: NewNode2(),
	}
}

type Right struct {
	Box map[string]int
}

func (r *Right) Insert(word string) {
	if _, ok := r.Box[word]; !ok {
		r.Box[word] = 1
	} else {
		r.Box[word]++
	}
}

func (r *Right) Delete(word string) {
	if _, ok := r.Box[word]; !ok {
		return
	} else {
		r.Box[word]--
		if r.Box[word] == 0 {
			delete(r.Box, word)
		}
	}
}

func (r *Right) Search(word string) int {
	return r.Box[word]
}

func (r *Right) PrefixCount(word string) int {
	count := 0
	for k, v := range r.Box {
		if strings.HasPrefix(k, word) {
			count += v
		}
	}
	return count
}

func NewRight() *Right {
	return &Right{Box: make(map[string]int)}
}

func generateRandomString(strLen int) string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < strLen; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func generateRandomStringArray(arrLen int, strLen int) []string {
	result := make([]string, arrLen)
	for i := 0; i < arrLen; i++ {
		result[i] = generateRandomString(strLen)
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arrLen := 100
	strLen := 20
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		arr := generateRandomStringArray(arrLen, strLen)
		trie := NewTrieTree()
		trie2 := NewTrieTree2()
		right := NewRight()
		for j := 0; j < len(arr); j++ {
			decide := rand.Float64()
			if decide < 0.25 {
				trie.Insert(arr[j])
				trie2.Insert(arr[j])
				right.Insert(arr[j])
			} else if decide < 0.5 {
				trie.Delete(arr[j])
				trie2.Delete(arr[j])
				right.Delete(arr[j])
			} else if decide < 0.75 {
				ans1 := trie.Search(arr[j])
				ans2 := trie2.Search(arr[j])
				ans3 := right.Search(arr[j])
				if ans1 != ans2 || ans2 != ans3 {
					fmt.Println("Oops")
					return
				}
			} else {
				ans1 := trie.PrefixCount(arr[j])
				ans2 := trie2.PrefixCount(arr[j])
				ans3 := right.PrefixCount(arr[j])
				if ans1 != ans2 || ans2 != ans3 {
					fmt.Println("Oops")
					return
				}
			}
		}
	}
	fmt.Println("Nice")
}
