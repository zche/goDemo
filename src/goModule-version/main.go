package main

import (
	"fmt"
	"github.com/zche/stardust/stringsx"
	stringsxV2 "github.com/zche/stardust/v2/stringsx"
	)

func main() {
	fmt.Println(stringsx.Hello("check"))
	if greet, err := stringsxV2.Hello("check", "zh"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greet)
	}
}
