package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	h1 := New16()
	h1.Add([]byte("Apple"))
	fmt.Println("apple", hex.EncodeToString(h1.Sum(nil)))
	h1.Add([]byte("Banana"))
	h1.Add([]byte("Orange"))
	fmt.Println("apple, banana, orange", hex.EncodeToString(h1.Sum(nil)))

	h2 := New16()
	h2.Add([]byte("Banana"))
	h2.Add([]byte("Orange"))
	h2sum := h2.Sum(nil)

	var h2sumHash [2048]byte
	copy(h2sumHash[:], h2sum[:])
	h1.RemoveHash(&h2sumHash)

	fmt.Println("total", hex.EncodeToString(h1.Sum(nil)))
}
