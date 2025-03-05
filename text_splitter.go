package main

import (
	"github.com/pkoukk/tiktoken-go"
	"strings"
)

// Tokenizer for OpenAI
var enc, _ = tiktoken.GetEncoding("cl100k_base")

type Splitter interface {
	Split(text string) []string
}

// Character-based splitter
type CharacterSplitter struct {
	ChunkSize    int
	ChunkOverlap int
}

func (cs *CharacterSplitter) Split(text string) []string {
	// Implementation here
}

// Recursive splitter (similar logic adapted from Python)
type RecursiveSplitter struct {
	ChunkSize    int
	ChunkOverlap int
}

func (rs *RecursiveSplitter) Split(text string) []string {
	// Implementation here
}

// Token counting function
func countTokens(text string) int {
	return len(enc.Encode(text, nil, nil))
}