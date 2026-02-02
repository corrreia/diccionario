package wordlist

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// WordList contains a list of words and the supported list operations.
type WordList interface {
	// AddWord persists a new word to the existing list.
	AddWord(word string) (err error)

	// GetWords returns all of the words in the existing list.
	GetWords() (words []string, err error)

	// WordExists checks if a word exists in the list (case insensitive).
	WordExists(word string) (exists bool, err error)
}

type wordListImpl struct {
	filename string
}

// New instantiates a new WordList.
func New(filename string) WordList {
	return &wordListImpl{filename: filename}
}

// AddWord persists a new word to the existing list.
func (w *wordListImpl) AddWord(word string) (err error) {
	// todo: add some sort of  locking feature (mutex)
	var f *os.File
	if f, err = os.OpenFile(w.filename, os.O_APPEND, 0644); err != nil {
		return
	}
	defer f.Close()

	if _, err = f.Write([]byte(word)); err != nil {
		return
	}

	return
}

// GetWords returns all of the words in the existing list.
func (w *wordListImpl) GetWords() (words []string, err error) {
	var f *os.File
	if f, err = os.Open(w.filename); err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		var s string
		s, err = r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// Add the last word if there's content (file may not end with newline)
				if s = strings.TrimSpace(s); s != "" {
					words = append(words, s)
				}
				err = nil // EOF is not an error condition
				return words, err
			}
			return words, err
		}

		words = append(words, strings.TrimSpace(s))
	}
}

// WordExists checks if a word exists in the list (case insensitive).
// It builds a map for O(1) lookup.
func (w *wordListImpl) WordExists(word string) (exists bool, err error) {
	words, err := w.GetWords()
	if err != nil {
		return false, err
	}

	// Build a map with lowercase keys for case-insensitive lookup
	wordMap := make(map[string]struct{}, len(words))
	for _, w := range words {
		wordMap[strings.ToLower(w)] = struct{}{}
	}

	_, exists = wordMap[strings.ToLower(word)]
	return exists, nil
}
