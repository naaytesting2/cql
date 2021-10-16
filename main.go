package main

import (
	"errors"
	"fmt"
)

func b() error { return errors.New("B") }

func a() (err error) {
	resolve := func() {
		err = b()
		if err != nil {
			err = fmt.Errorf("error: %v", err)
			return
		}
	}

	resolve()
	return
}

func main() {
	fmt.Println(a())
}
