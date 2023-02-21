package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	pl := fmt.Println
	vNumberFloat := 1.5
	vNumber := int(vNumberFloat)
	pl(vNumber)

	vMyMoney := "50000000"
	vMyMoneyNumber, err := strconv.Atoi(vMyMoney)
	pl(vMyMoneyNumber, err, reflect.TypeOf(vMyMoneyNumber))
}
