package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var pl = fmt.Println
	pl("hello go!")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("hello,", name)
	} else {
		log.Fatal(err)
	}
}
