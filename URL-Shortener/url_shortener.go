package main

import (
	"errors"
	"fmt"
	"strconv"
)

type LinkShortener interface {
	Shorten(string) (string, error)
	Resolve(string) (string, error)
}

type URLShortener struct {
	store   map[string]string
	counter int // simple counter to generate unique hashes
}

var URLNotFoundError = errors.New("URL not found")

// Constructor for URLShortener
func NewURLShortener() *URLShortener {
	return &URLShortener{
		store:   make(map[string]string),
		counter: 0,
	}
}

func (u *URLShortener) Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL cannot be empty")
	}

	hash := strconv.Itoa(u.counter)
	u.counter++
	u.store[hash] = url

	return hash, nil
}

func (u *URLShortener) Resolve(hash string) (string, error) {
	if hash == "" {
		return "", errors.New("hash cannot be empty.")
	}
	if val, ok := u.store[hash]; ok {
		return val, nil
	}
	return "", URLNotFoundError

}

func main() {
	store := NewURLShortener()

	code, _ := store.Shorten("https://google.com")
	fmt.Println("Short Code:", code)

	original, _ := store.Resolve(code)
	fmt.Println("Original:", original)

	code2, _ := store.Shorten("https://googlee.com")
	fmt.Println("Short Code:", code2)

	original2, _ := store.Resolve(code2)
	fmt.Println("Original:", original2)
}
