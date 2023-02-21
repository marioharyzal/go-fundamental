package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	pl := fmt.Println
	sWord := "A word"
	sCat := "A cat"
	// string NewReplace and replace
	replacer := strings.NewReplacer("A", "Another")
	sWord2 := replacer.Replace(sWord)
	sNewCat := replacer.Replace(sCat)
	pl(sWord2)
	pl(sNewCat)

	// length string
	pl("length sWord :", len(sWord2))
	pl("length sNewCat :", len(sNewCat))

	// string contains
	bIsContains := strings.Contains(sWord2, "word")
	bIsContains2 := strings.Contains(sNewCat, "cat")
	pl("contains word :", bIsContains)
	pl("contains cat :", bIsContains2)

	// replace
	pl(strings.Replace(sNewCat, "c", "C", 1))

	// split
	sSplitString := strings.Split(sNewCat, "")
	pl(reflect.TypeOf(sSplitString))
}
