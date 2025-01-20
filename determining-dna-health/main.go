// https://github.com/runtimeninja
// https://www.hackerrank.com/profile/toufiq_py
// https://www.hackerrank.com/challenges/determining-dna-health/problem?isFullScreen=true
// https://toufiq.info

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// TrieNode represents a node in the Trie structure.
type TrieNode struct {
	children map[rune]*TrieNode // Child nodes for each character in the gene
	health   int32              // Total health value for this node
	indexes  []int32            // List of indexes where this gene appears
}

// Trie represents the Trie data structure.
type Trie struct {
	root *TrieNode
}

// newTrie initializes and returns a new Trie.
func newTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// insert adds a gene to the Trie, associating it with health and index.
func (t *Trie) insert(gene string, health int32, index int32) {
	current := t.root
	for _, char := range gene {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[char]
	}
	current.health += health
	current.indexes = append(current.indexes, index)
}

// calculateHealth computes the total health of a DNA strand within a given range.
func (t *Trie) calculateHealth(dna string, start int32, end int32) int32 {
	totalHealth := int32(0)
	for i := 0; i < len(dna); i++ {
		current := t.root
		for j := i; j < len(dna); j++ {
			char := rune(dna[j])
			if _, exists := current.children[char]; !exists {
				break
			}
			current = current.children[char]
			for _, idx := range current.indexes {
				if idx >= start && idx <= end {
					totalHealth += current.health
				}
			}
		}
	}
	return totalHealth
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	genesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var genes []string
	for i := 0; i < int(n); i++ {
		genes = append(genes, genesTemp[i])
	}

	healthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var health []int32
	for i := 0; i < int(n); i++ {
		healthItemTemp, err := strconv.ParseInt(healthTemp[i], 10, 64)
		checkError(err)
		health = append(health, int32(healthItemTemp))
	}

	sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	s := int32(sTemp)

	trie := newTrie()
	for i := int32(0); i < n; i++ {
		trie.insert(genes[i], health[i], i)
	}

	minHealth := int32(1<<31 - 1)
	maxHealth := int32(0)

	for sItr := 0; sItr < int(s); sItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		first := int32(firstTemp)

		lastTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		last := int32(lastTemp)

		d := firstMultipleInput[2]

		totalHealth := trie.calculateHealth(d, first, last)

		if totalHealth < minHealth {
			minHealth = totalHealth
		}
		if totalHealth > maxHealth {
			maxHealth = totalHealth
		}
	}

	fmt.Printf("%d %d\n", minHealth, maxHealth)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
