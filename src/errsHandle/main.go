package main

import (
	"errors"
	"fmt"
)

type Divide struct {
	dividend int
	divider  int
}

func (d *Divide) Error() string {
	strFormat := `
	Can't proceed,the divider is zero.
	dividend: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, d.dividend)
}

func computeDivi(d *Divide) (result int, err error) {
	if d.divider == 0 {
		err = d
	} else {
		result = d.dividend / d.divider
	}
	return
}

func main() {
	err := errors.New("a new err object")
	fmt.Printf("%v\n", err)

	err = fmt.Errorf("a fmt err format object: %s", err.Error())
	fmt.Printf("%v\n", err)

	de := Divide{100, 0}
	if result, err := computeDivi(&de); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
