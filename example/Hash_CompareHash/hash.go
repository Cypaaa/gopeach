package main

import (
	"fmt"

	"github.com/Cypaaa/gopeach"
)

func main() {
	// generate hash
	hash, err := gopeach.Hash("test")
	if err != nil {
		panic(err)
	}

	// compare good hash
	okg, err := gopeach.CompareHash("test", hash)
	if err != nil {
		panic(err)
	}

	// are they the same?
	fmt.Println(okg) // true

	// compare wrong hash
	okw, err := gopeach.CompareHash("something different", hash)
	if err != nil {
		panic(err)
	}

	// are they the same?
	fmt.Println(okw) // false
}
